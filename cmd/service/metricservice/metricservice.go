package metricservice

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/twistedogic/spero/pkg/metric"
)

var (
	registry = prometheus.NewRegistry()
)

type Service struct {
	registry *prometheus.Registry
}

func New() *Service {
	registry.MustRegister(
		metric.OddMetric,
		metric.ClientMetric,
		metric.PollMetric,
		metric.ScrapeMetric,
	)
	return &Service{registry}
}

func (s *Service) Handler() http.Handler {
	return promhttp.HandlerFor(s.registry, promhttp.HandlerOpts{})
}

func (s *Service) Route() string { return "/metrics" }
