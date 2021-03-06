package model

import (
	"github.com/evergreen-ci/sink"
	"github.com/evergreen-ci/sink/util"
	"github.com/mongodb/anser/bsonutil"
	"github.com/mongodb/anser/db"
	"github.com/mongodb/anser/model"
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/message"
	"github.com/mongodb/grip/send"
	"github.com/pkg/errors"
)

const sinkConfigurationID = "sink-system-configuration"

type SinkConfig struct {
	ID     string                    `bson:"_id" json:"id" yaml:"id"`
	Splunk send.SplunkConnectionInfo `bson:"splunk" json:"splunk" yaml:"splunk"`
	Slack  SlackConfig               `bson:"slack" json:"slack" yaml:"slack"`

	populated bool
	env       sink.Environment
}

var (
	sinkConfigurationIDKey     = bsonutil.MustHaveTag(SinkConfig{}, "ID")
	sinkConfigurationSplunkKey = bsonutil.MustHaveTag(SinkConfig{}, "Splunk")
	sinkConfigurationSlackKey  = bsonutil.MustHaveTag(SinkConfig{}, "Slack")
)

type SlackConfig struct {
	Options *send.SlackOptions `bson:"options" json:"options" yaml:"options"`
	Token   string             `bson:"token" json:"token" yaml:"token"`
	Level   string             `bson:"level" json:"level" yaml:"level"`
}

var (
	sinkSlackConfigOptionsKey = bsonutil.MustHaveTag(SlackConfig{}, "Options")
	sinkSlackConfigTokenKey   = bsonutil.MustHaveTag(SlackConfig{}, "Token")
	sinkSlackConfigLevelKey   = bsonutil.MustHaveTag(SlackConfig{}, "Level")
)

func (c *SinkConfig) Setup(e sink.Environment) { c.env = e }
func (c *SinkConfig) IsNil() bool              { return !c.populated }
func (c *SinkConfig) Find() error {
	conf, session, err := sink.GetSessionWithConfig(c.env)
	if err != nil {
		return errors.WithStack(err)
	}
	defer session.Close()

	c.populated = false
	err = session.DB(conf.DatabaseName).C(configurationCollection).FindId(sinkConfigurationID).One(c)
	if db.ResultsNotFound(err) {
		return errors.New("could not find application configuration in the database")
	} else if err != nil {
		return errors.Wrap(err, "problem finding app config document")
	}

	c.populated = true
	return nil
}

func (c *SinkConfig) Save() error {
	// TODO: validate here when that's possible

	if !c.populated {
		return errors.New("cannot save a non-populated app configuration")
	}

	c.ID = sinkConfigurationID

	conf, session, err := sink.GetSessionWithConfig(c.env)
	if err != nil {
		return errors.WithStack(err)
	}
	defer session.Close()

	changeInfo, err := session.DB(conf.DatabaseName).C(configurationCollection).UpsertId(sinkConfigurationID, c)
	grip.Debug(message.Fields{
		"ns":          model.Namespace{DB: conf.DatabaseName, Collection: configurationCollection},
		"id":          sinkConfigurationID,
		"operation":   "save build cost reporting configuration",
		"change_info": changeInfo,
	})

	return errors.Wrap(err, "problem saving application configuration")
}

func LoadSinkConfig(file string) (*SinkConfig, error) {
	newConfig := &SinkConfig{}

	if err := util.ReadFileYAML(file, newConfig); err != nil {
		return nil, errors.WithStack(err)
	}

	// TODO: validate here (?)

	newConfig.populated = true

	return newConfig, nil
}
