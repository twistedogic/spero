package result

import (
	"time"

	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/pkg/errors"

	"github.com/twistedogic/spero/pkg/client/jc"
	"github.com/twistedogic/spero/pkg/input/inpututil"
	"github.com/twistedogic/spero/pkg/poll"
)

const (
	InputName = "result"

	intervalField = "interval"
	urlField      = "base_url"
	startField    = "start"
	endField      = "end"

	intervalFieldDefault = "15m"
	urlFieldDefault      = jc.DefaultURL
)

var (
	startFieldDefault = time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC).Format(inpututil.TimeFormat)
	endFieldDefault   = time.Now().UTC().Format(inpututil.TimeFormat)
)

func Register() error {
	configSpec := service.NewConfigSpec().
		Summary("match and odd result to poll").
		Field(service.NewStringField(intervalField).Default(intervalFieldDefault)).
		Field(service.NewStringField(urlField).Default(urlFieldDefault)).
		Field(service.NewStringField(startField).Default(startFieldDefault)).
		Field(service.NewStringField(endField).Default(endFieldDefault))

	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.Input, error) {
		interval, err := inpututil.ParseIntervalField(conf, intervalField)
		if err != nil {
			return nil, err
		}
		start, err := inpututil.ParseTimeField(conf, startField)
		if err != nil {
			return nil, err
		}
		end, err := inpututil.ParseTimeField(conf, endField)
		if err != nil {
			return nil, err
		}
		baseURL, err := conf.FieldString(urlField)
		if err != nil {
			return nil, err
		}
		client := jc.New(baseURL)
		return poll.New(client.PollResult(start, end), interval), nil
	}
	if err := service.RegisterInput(InputName, configSpec, constructor); err != nil {
		return errors.Wrapf(err, "register input %s", InputName)
	}
	return nil
}
