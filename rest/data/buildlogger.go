package data

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	dbModel "github.com/evergreen-ci/cedar/model"
	"github.com/evergreen-ci/cedar/rest/model"
	"github.com/evergreen-ci/cedar/util"
	"github.com/evergreen-ci/gimlet"
	"github.com/evergreen-ci/pail"
	"github.com/mongodb/anser/db"
	"github.com/pkg/errors"
)

/////////////////////////////
// DBConnector Implementation
/////////////////////////////

func (dbc *DBConnector) FindLogByID(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	log := dbModel.Log{ID: opts.ID}
	log.Setup(dbc.env)
	if err := log.Find(ctx); db.ResultsNotFound(err) {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("log with id '%s' not found", opts.ID),
		}
	} else if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("database error"),
		}
	}

	log.Setup(dbc.env)
	it, err := log.Download(ctx, opts.TimeRange)
	if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("%s", errors.Wrap(err, "problem downloading log")),
		}
	}

	readerOpts := dbModel.LogIteratorReaderOptions{PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, it, readerOpts), nil
}

func (dbc *DBConnector) FindLogMetadataByID(ctx context.Context, id string) (*model.APILog, error) {
	log := dbModel.Log{ID: id}
	log.Setup(dbc.env)
	if err := log.Find(ctx); db.ResultsNotFound(err) {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("log with id '%s' not found", id),
		}
	} else if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("database error"),
		}
	}

	apiLog := &model.APILog{}
	if err := apiLog.Import(log); err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "corrupt data",
		}
	}

	return apiLog, nil
}

func (dbc *DBConnector) FindLogsByTaskID(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	dbOpts := dbModel.LogFindOptions{
		TimeRange: opts.TimeRange,
		Info: dbModel.LogInfo{
			TaskID:      opts.TaskID,
			ProcessName: opts.ProcessName,
			Tags:        opts.Tags,
		},
	}
	logs := dbModel.Logs{}
	logs.Setup(dbc.env)
	if err := logs.Find(ctx, dbOpts); db.ResultsNotFound(err) {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' not found", opts.TaskID),
		}
	} else if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("database error"),
		}
	}

	logs.Setup(dbc.env)
	it, err := logs.Merge(ctx)
	if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("%s", errors.Wrap(err, "problem downloading log")),
		}
	}

	readerOpts := dbModel.LogIteratorReaderOptions{TailN: opts.Tail, PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, it, readerOpts), nil
}

func (dbc *DBConnector) FindLogMetadataByTaskID(ctx context.Context, opts BuildloggerOptions) ([]model.APILog, error) {
	dbOpts := dbModel.LogFindOptions{
		TimeRange: util.TimeRange{EndAt: time.Now()},
		Info: dbModel.LogInfo{
			TaskID: opts.TaskID,
			Tags:   opts.Tags,
		},
	}
	logs := dbModel.Logs{}
	logs.Setup(dbc.env)
	if err := logs.Find(ctx, dbOpts); db.ResultsNotFound(err) {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' not found", opts.TaskID),
		}
	} else if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("database error"),
		}
	}

	apiLogs := make([]model.APILog, len(logs.Logs))
	for i, log := range logs.Logs {
		if err := apiLogs[i].Import(log); err != nil {
			return nil, gimlet.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "corrupt data",
			}
		}
	}

	return apiLogs, nil
}

func (dbc *DBConnector) FindLogsByTestName(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	it, err := dbc.findLogsByTestName(ctx, opts)
	if err != nil {
		return nil, err
	}
	readerOpts := dbModel.LogIteratorReaderOptions{PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, it, readerOpts), nil
}

func (dbc *DBConnector) FindLogMetadataByTestName(ctx context.Context, opts BuildloggerOptions) ([]model.APILog, error) {
	dbOpts := dbModel.LogFindOptions{
		TimeRange: util.TimeRange{EndAt: time.Now()},
		Info: dbModel.LogInfo{
			TaskID: opts.TaskID,
			Tags:   opts.Tags,
		},
	}
	if opts.TestName != "" {
		dbOpts.Info.TestName = opts.TestName
	} else {
		dbOpts.Empty = dbModel.EmptyLogInfo{TestName: true}
	}
	logs := dbModel.Logs{}
	logs.Setup(dbc.env)
	if err := logs.Find(ctx, dbOpts); db.ResultsNotFound(err) {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' and test name '%s' not found", opts.TaskID, opts.TestName),
		}
	} else if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("database error"),
		}
	}

	apiLogs := make([]model.APILog, len(logs.Logs))
	for i, log := range logs.Logs {
		if err := apiLogs[i].Import(log); err != nil {
			return nil, gimlet.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "corrupt data",
			}
		}
	}

	return apiLogs, nil
}

func (dbc *DBConnector) FindGroupedLogs(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	its := []dbModel.LogIterator{}
	it, err := dbc.findLogsByTestName(ctx, opts)
	if err != nil {
		return nil, err
	}
	its = append(its, it)

	opts.TestName = ""
	it, err = dbc.findLogsByTestName(ctx, opts)
	if err == nil {
		its = append(its, it)
	} else if errResp, ok := err.(gimlet.ErrorResponse); !ok || errResp.StatusCode != http.StatusNotFound {
		return nil, err
	}

	readerOpts := dbModel.LogIteratorReaderOptions{PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, dbModel.NewMergingIterator(its...), readerOpts), nil
}

func (dbc *DBConnector) findLogsByTestName(ctx context.Context, opts BuildloggerOptions) (dbModel.LogIterator, error) {
	dbOpts := dbModel.LogFindOptions{
		TimeRange: opts.TimeRange,
		Info: dbModel.LogInfo{
			TaskID: opts.TaskID,
			Tags:   opts.Tags,
		},
	}
	if opts.TestName != "" {
		dbOpts.Info.TestName = opts.TestName
	} else {
		dbOpts.Empty = dbModel.EmptyLogInfo{TestName: true}
	}
	logs := dbModel.Logs{}
	logs.Setup(dbc.env)
	if err := logs.Find(ctx, dbOpts); db.ResultsNotFound(err) {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' and test name '%s' not found", opts.TaskID, opts.TestName),
		}
	} else if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("database error"),
		}
	}

	logs.Setup(dbc.env)
	it, err := logs.Merge(ctx)
	if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("%s", errors.Wrap(err, "problem downloading log")),
		}
	}

	return it, nil
}

///////////////////////////////
// MockConnector Implementation
///////////////////////////////

func (mc *MockConnector) FindLogByID(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	log, ok := mc.CachedLogs[opts.ID]
	if !ok {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("log with id '%s' not found", opts.ID),
		}
	}

	bucketOpts := pail.LocalOptions{
		Path:   mc.Bucket,
		Prefix: log.Artifact.Prefix,
	}
	bucket, err := pail.NewLocalBucket(bucketOpts)
	if err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("%s", errors.Wrap(err, "problem creating bucket")),
		}
	}

	readerOpts := dbModel.LogIteratorReaderOptions{PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, dbModel.NewBatchedLogIterator(bucket, log.Artifact.Chunks, 2, opts.TimeRange), readerOpts), ctx.Err()
}

func (mc *MockConnector) FindLogMetadataByID(ctx context.Context, id string) (*model.APILog, error) {
	log, ok := mc.CachedLogs[id]
	if !ok {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("log with id '%s' not found", id),
		}
	}

	apiLog := &model.APILog{}
	if err := apiLog.Import(log); err != nil {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "corrupt data",
		}
	}

	return apiLog, ctx.Err()
}

func (mc *MockConnector) FindLogsByTaskID(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	logs := []dbModel.Log{}
	for _, log := range mc.CachedLogs {
		if log.Info.TaskID == opts.TaskID {
			logs = append(logs, log)
		}
	}
	if len(logs) == 0 {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' not found", opts.TaskID),
		}
	}

	sort.Slice(logs, func(i, j int) bool { return logs[i].CreatedAt.After(logs[j].CreatedAt) })

	its := []dbModel.LogIterator{}
	for _, log := range logs {
		if opts.ProcessName != "" && opts.ProcessName != log.Info.ProcessName {
			continue
		}
		if !containsTags(opts.Tags, log.Info.Tags) {
			continue
		}

		bucketOpts := pail.LocalOptions{
			Path:   mc.Bucket,
			Prefix: log.Artifact.Prefix,
		}
		bucket, err := pail.NewLocalBucket(bucketOpts)
		if err != nil {
			return nil, gimlet.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("%s", errors.Wrap(err, "problem creating bucket")),
			}
		}

		its = append(its, dbModel.NewBatchedLogIterator(bucket, log.Artifact.Chunks, 2, opts.TimeRange))
	}

	it := dbModel.NewMergingIterator(its...)
	readerOpts := dbModel.LogIteratorReaderOptions{TailN: opts.Tail, PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, it, readerOpts), ctx.Err()
}

func (mc *MockConnector) FindLogMetadataByTaskID(ctx context.Context, opts BuildloggerOptions) ([]model.APILog, error) {
	logs := []dbModel.Log{}
	for _, log := range mc.CachedLogs {
		if log.Info.TaskID == opts.TaskID {
			logs = append(logs, log)
		}
	}
	if len(logs) == 0 {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' not found", opts.TaskID),
		}
	}

	sort.Slice(logs, func(i, j int) bool { return logs[i].CreatedAt.After(logs[j].CreatedAt) })

	apiLogs := []model.APILog{}
	for _, log := range logs {
		if !containsTags(opts.Tags, log.Info.Tags) {
			continue
		}

		apiLogs = append(apiLogs, model.APILog{})
		if err := apiLogs[len(apiLogs)-1].Import(log); err != nil {
			return nil, gimlet.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "corrupt data",
			}
		}
	}

	return apiLogs, ctx.Err()
}

func (mc *MockConnector) FindLogsByTestName(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	it, err := mc.findLogsByTestName(ctx, opts)
	if err != nil {
		return nil, err
	}
	readerOpts := dbModel.LogIteratorReaderOptions{PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, it, readerOpts), nil
}

func (mc *MockConnector) FindLogMetadataByTestName(ctx context.Context, opts BuildloggerOptions) ([]model.APILog, error) {
	logs := []dbModel.Log{}
	for _, log := range mc.CachedLogs {
		if log.Info.TaskID == opts.TaskID && log.Info.TestName == opts.TestName {
			logs = append(logs, log)
		}
	}
	if len(logs) == 0 {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' and test name '%s' not found", opts.TaskID, opts.TestName),
		}
	}

	sort.Slice(logs, func(i, j int) bool { return logs[i].CreatedAt.After(logs[j].CreatedAt) })

	apiLogs := []model.APILog{}
	for _, log := range logs {
		if !containsTags(opts.Tags, log.Info.Tags) {
			continue
		}

		apiLogs = append(apiLogs, model.APILog{})
		if err := apiLogs[len(apiLogs)-1].Import(log); err != nil {
			return nil, gimlet.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "corrupt data",
			}
		}

	}

	return apiLogs, ctx.Err()
}

func (mc *MockConnector) FindGroupedLogs(ctx context.Context, opts BuildloggerOptions) (io.Reader, error) {
	its := []dbModel.LogIterator{}
	it, err := mc.findLogsByTestName(ctx, opts)
	if err != nil {
		return nil, err
	}
	its = append(its, it)

	opts.TestName = ""
	it, err = mc.findLogsByTestName(ctx, opts)
	if err == nil {
		its = append(its, it)
	} else if errResp, ok := err.(gimlet.ErrorResponse); !ok || errResp.StatusCode != http.StatusNotFound {
		return nil, err
	}

	readerOpts := dbModel.LogIteratorReaderOptions{PrintTime: opts.PrintTime}
	return dbModel.NewLogIteratorReader(ctx, dbModel.NewMergingIterator(its...), readerOpts), ctx.Err()
}

func (mc *MockConnector) findLogsByTestName(ctx context.Context, opts BuildloggerOptions) (dbModel.LogIterator, error) {
	logs := []dbModel.Log{}
	for _, log := range mc.CachedLogs {
		if log.Info.TaskID == opts.TaskID && log.Info.TestName == opts.TestName && containsTags(opts.Tags, log.Info.Tags) {
			logs = append(logs, log)
		}
	}

	if len(logs) == 0 {
		return nil, gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("logs with task id '%s' and test name '%s' not found", opts.TaskID, opts.TestName),
		}
	}

	sort.Slice(logs, func(i, j int) bool { return logs[i].CreatedAt.After(logs[j].CreatedAt) })

	its := []dbModel.LogIterator{}
	for _, log := range logs {
		bucketOpts := pail.LocalOptions{
			Path:   mc.Bucket,
			Prefix: log.Artifact.Prefix,
		}
		bucket, err := pail.NewLocalBucket(bucketOpts)
		if err != nil {
			return nil, gimlet.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("%s", errors.Wrap(err, "problem creating bucket")),
			}
		}

		its = append(its, dbModel.NewBatchedLogIterator(bucket, log.Artifact.Chunks, 2, opts.TimeRange))
	}

	return dbModel.NewMergingIterator(its...), ctx.Err()
}