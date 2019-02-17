package metajob

import (
	"log"
	"time"

	"github.com/twistedogic/spero/pkg/poll"
	"github.com/twistedogic/spero/pkg/scraper/metascraper"
	"github.com/twistedogic/spero/pkg/storage"
)

type Job struct {
	service *poll.PollService
	store   *storage.Storage
}

func New(interval time.Duration, eventURL string, rate int, store *storage.Storage) *Job {
	poller := metascraper.New(eventURL, rate)
	service := poll.New(poller, interval)
	return &Job{service, store}
}

func (j *Job) Start() error {
	ch := make(chan interface{})
	go func() {
		for v := range ch {
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
