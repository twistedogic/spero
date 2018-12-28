package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	OddMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "odd",
		Help: "current odd",
	}, []string{
		"matchID",
		"oddtype",
		"home",
		"away",
		"outcome",
	})
	ClientMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "client_failure_count",
		Help: "client request failure count",
	}, []string{
		"error_type",
		"statuscode",
	})
)
