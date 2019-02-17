package oddjob

import (
	"log"
	"time"

	"github.com/twistedogic/spero/pkg/metric"
	"github.com/twistedogic/spero/pkg/poll"
	"github.com/twistedogic/spero/pkg/schema/odd"
	"github.com/twistedogic/spero/pkg/scraper/oddscraper"
	"github.com/twistedogic/spero/pkg/storage"
)

type Job struct {
	service *poll.PollService
	store   *storage.Storage
}

func New(interval time.Duration, oddURL string, types []string, store *storage.Storage) *Job {
  bettype := make([]odd.OddEnum, 0)
  for _, t := range types {
    if v, ok := odd.ToBetTypeEnum(t); ok {
      bettype = append(bettype, v)
    }
  }
	poller := oddscraper.New(oddURL, bettype...)
	service := poll.New(poller, interval)
	return &Job{service, store}
}

func (j *Job) Start() error {
	ch := make(chan interface{}, 1)
	go func() {
		for v := range ch {
			if match, ok := v.(odd.Match); ok {
				for _, outcome := range match.GetOutcome() {
					metric.UpdateOddMetric(outcome)
				}
      }
      if err := j.store.WriteToDB(v); err != nil {
        log.Println(err)
      }
		}
	}()
	return j.service.Start(ch)
}

func (j *Job) Stop() {
	j.service.Stop()
}
