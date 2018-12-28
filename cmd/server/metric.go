package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricServer struct {
	instance *http.Server
}

func NewMetricServer(registry *prometheus.Registry, port int) *MetricServer {
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	http.Handle("/metrics", handler)
	return &MetricServer{instance: httpServer}
}

func (m *MetricServer) Start() {
	log.Println("Starting Metric Server")
	if err := m.instance.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (m *MetricServer) Stop() error {
	log.Println("Stoping Metric Server")
	if err := m.instance.Shutdown(nil); err != nil {
		return err
	}
	return nil
}
