package metric

import (
	"reflect"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/twistedogic/spero/pkg/errors"
	"github.com/twistedogic/spero/pkg/schema/odd"
)

var (
	OddMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "odd",
		Help: "current odd",
	}, []string{
		"matchID",
		"oddID",
		"type",
		"outcome",
	})
	ClientMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "client_request_count",
		Help: "client request count",
	}, []string{
		"status",
	})
	PollMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "poll_count",
		Help: "poller poll count",
	}, []string{
		"type",
		"status",
	})
	ScrapeMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "scrape_count",
		Help: "scraper scrape count",
	}, []string{
		"type",
		"status",
	})
	DBMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "db_write_count",
		Help: "db write count",
	}, []string{
		"type",
		"status",
	})
)

func UpdateOddMetric(o odd.Outcome) {
	OddMetric.WithLabelValues(
		o.MatchID,
		o.OddID,
		o.Type.Key(),
		o.Outcome,
	).Set(o.Odd)
}

func UpdateClientMetric(status string) { ClientMetric.WithLabelValues(status).Inc() }

func UpdatePollMetric(bettype odd.OddEnum, err error) {
	PollMetric.WithLabelValues(bettype.Key(), errors.ParseError(err)).Inc()
}

func UpdateScrapeMetric(v interface{}, err error) {
	structName := reflect.TypeOf(v).Name()
	ScrapeMetric.WithLabelValues(structName, errors.ParseError(err)).Inc()
}

func UpdateDBMetric(v interface{}, err error) {
	structName := reflect.TypeOf(v).Name()
	DBMetric.WithLabelValues(structName, errors.ParseError(err)).Inc()
}
