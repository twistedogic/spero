package poll

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/twistedogic/spero/pkg/client"
	"github.com/twistedogic/spero/pkg/schema"
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

type Poller struct {
	client *client.Client
	ticker *time.Ticker
	done   chan struct{}
}

func New(base, interval string) (*Poller, error) {
	period, err := time.ParseDuration(interval)
	if err != nil {
		return nil, err
	}
	ticker := time.NewTicker(period)
	c := client.New(base)
	return &Poller{
		client: c,
		ticker: ticker,
		done:   make(chan struct{}),
	}, nil
}

func (p *Poller) Start(bets ...string) chan schema.Match {
	out := make(chan schema.Match)
	go func() {
		<-p.done
		p.ticker.Stop()
		close(out)
	}()
	go func() {
		for {
			select {
			case <-p.ticker.C:
				for _, bet := range bets {
					res, err := p.client.GetMatchesByType(bet)
					if err != nil {
						v := err.(client.Error)
						ClientMetric.WithLabelValues(v.Type, string(v.StatusCode)).Inc()
					}
					for _, m := range res {
						for _, entry := range m.ToProm() {
							OddMetric.WithLabelValues(entry.Labels...).Set(entry.Value)
						}
						out <- m
					}
				}
			}
		}
	}()
	return out
}

func (p *Poller) Stop() {
	p.done <- struct{}{}
	log.Println("Stopping Poller")
}
