package poll

import (
	"log"
	"time"

	pb "github.com/twistedogic/spero/pb"
	"github.com/twistedogic/spero/pkg/client"
)

type Poller struct {
	client *client.Client
	ticker *time.Ticker
	done   chan struct{}
}

func New(base string, interval time.Duration) *Poller {
	ticker := time.NewTicker(interval)
	c := client.New(base)
	return &Poller{
		client: c,
		ticker: ticker,
		done:   make(chan struct{}),
	}
}

func (p *Poller) Start(out chan *pb.Odd, bets ...string) {
	log.Println("Start Polling...")
	go func() {
		<-p.done
		p.ticker.Stop()
		close(out)
	}()
	go func() {
		for {
			select {
			case <-p.ticker.C:
				log.Println("poll")
				for _, bet := range bets {
					if res, err := p.client.GetMatchesByType(bet); err == nil {
						for _, m := range res {
							for _, odd := range m.ToProto() {
								out <- odd
							}
						}
					}
				}
			}
		}
	}()
}

func (p *Poller) Stop() error {
	p.done <- struct{}{}
	log.Println("Stopping Poller")
	return nil
}
