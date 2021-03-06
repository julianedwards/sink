package rest

import (
	"context"

	"github.com/evergreen-ci/gimlet"
	"github.com/evergreen-ci/sink"
	"github.com/mongodb/amboy"
	"github.com/mongodb/grip"
	"github.com/pkg/errors"
)

type Service struct {
	Port int

	// internal settings
	queue amboy.Queue
	app   *gimlet.APIApp
	env   sink.Environment
}

func (s *Service) Validate() error {
	var err error

	if s.env == nil {
		s.env = sink.GetEnvironment()
	}

	if s.queue == nil {
		s.queue, err = s.env.GetQueue()
		if err != nil {
			return errors.Wrap(err, "problem getting queue")
		}
		if s.queue == nil {
			return errors.New("no queue defined")
		}
	}

	if s.app == nil {
		s.app = gimlet.NewApp()
	}

	if s.Port == 0 {
		s.Port = 3000
	}

	if err := s.app.SetPort(s.Port); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Service) Start(ctx context.Context) error {
	if s.queue == nil || s.app == nil {
		return errors.New("application is not valid")
	}

	s.addRoutes()

	if err := s.queue.Start(ctx); err != nil {
		return errors.Wrap(err, "problem starting queue")
	}

	if err := s.app.Resolve(); err != nil {
		return errors.Wrap(err, "problem resolving routes")
	}

	grip.Noticef("completed sink service; shutting down")

	return nil
}

// Run starts the REST service. All errors are logged.
func (s *Service) Run(ctx context.Context) {
	grip.CatchAlert(s.app.Run(ctx))
}

func (s *Service) addRoutes() {
	s.app.AddRoute("/status").Version(1).Get().Handler(s.statusHandler)
	s.app.AddRoute("/status/event/{id}").Version(1).Get().Handler(s.getSystemEvent)
	s.app.AddRoute("/status/event/{id}/acknowledge").Version(1).Get().Handler(s.acknowledgeSystemEvent)
	s.app.AddRoute("/status/events/{level}").Version(1).Get().Handler(s.getSystemEvents)
	s.app.AddRoute("/simple_log/{id}").Version(1).Post().Handler(s.simpleLogInjestion)
	s.app.AddRoute("/simple_log/{id}").Version(1).Get().Handler(s.simpleLogRetrieval)
	s.app.AddRoute("/simple_log/{id}/text").Version(1).Get().Handler(s.simpleLogGetText)
	s.app.AddRoute("/system_info").Version(1).Post().Handler(s.recieveSystemInfo)
	s.app.AddRoute("/system_info/host/{host}").Version(1).Post().Handler(s.fetchSystemInfo)

	s.app.AddRoute("/depgraph/{id}").Version(1).Post().Handler(s.createDepGraph)
	s.app.AddRoute("/depgraph/{id}").Version(1).Get().Handler(s.resolveDepGraph)
	s.app.AddRoute("/depgraph/{id}/nodes").Version(1).Post().Handler(s.addDepGraphNodes)
	s.app.AddRoute("/depgraph/{id}/nodes").Version(1).Get().Handler(s.getDepGraphNodes)
	s.app.AddRoute("/depgraph/{id}/edges").Version(1).Post().Handler(s.addDepGraphEdges)
	s.app.AddRoute("/depgraph/{id}/edges").Version(1).Get().Handler(s.getDepGraphEdges)
}
