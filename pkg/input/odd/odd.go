package odd

import (
	"time"

	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/pkg/errors"

	"github.com/twistedogic/spero/pkg/client/oddclient"
	"github.com/twistedogic/spero/pkg/poll"
)

const (
	InputName = "odd"

	typeField     = "odd_type"
	intervalField = "interval"
	urlField      = "base_url"

	typeFieldDefault     = "had"
	intervalFieldDefault = "15m"
	urlFieldDefault      = oddclient.DefaultURL
)

func Register() error {
	configSpec := service.NewConfigSpec().
		Summary("Type of odd to monitor").
		Field(service.NewStringField(typeField).Default(typeFieldDefault)).
		Field(service.NewStringField(intervalField).Default(intervalFieldDefault)).
		Field(service.NewStringField(urlField).Default(urlFieldDefault))

	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.Input, error) {
		oddType, err := conf.FieldString(typeField)
		if err != nil {
			return nil, err
		}
		intervalStr, err := conf.FieldString(intervalField)
		if err != nil {
			return nil, err
		}
		interval, err := time.ParseDuration(intervalStr)
		if err != nil {
			return nil, err
		}
		baseURL, err := conf.FieldString(urlField)
		client := oddclient.New(baseURL)
		return poll.New(client.PollOdd(oddType), interval), nil
	}
	if err := service.RegisterInput(InputName, configSpec, constructor); err != nil {
		return errors.Wrapf(err, "register input %s", InputName)
	}
	return nil
}
