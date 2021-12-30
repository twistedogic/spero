package inpututil

import (
	"time"

	"github.com/Jeffail/benthos/v3/public/service"
)

const TimeFormat = "2006-01-02"

func ParseTimeField(conf *service.ParsedConfig, field string) (time.Time, error) {
	timeFieldStr, err := conf.FieldString(field)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(TimeFormat, timeFieldStr)
}

func ParseIntervalField(conf *service.ParsedConfig, field string) (time.Duration, error) {
	intervalStr, err := conf.FieldString(field)
	if err != nil {
		return 0, err
	}
	return time.ParseDuration(intervalStr)
}
