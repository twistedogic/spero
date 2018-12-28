package errors

import (
	"github.com/twistedogic/spero/pkg/metric"
)

type Error struct {
	error
	Type       string
	StatusCode int
}

func NewClientError(t string, status int, err error) Error {
	metric.ClientMetric.WithLabelValues(t, string(status)).Inc()
	return Error{
		err, t, status,
	}
}
