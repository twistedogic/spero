package poll

import (
	"time"
)

type Poller interface {
	Poll(chan interface{}) error
}

type PollService struct {
	Poller
	Interval time.Duration
	done     chan struct{}
}

func New(poller Poller, interval time.Duration) *PollService {
	done := make(chan struct{})
	return &PollService{
		poller,
		interval,
		done,
	}
}

func (p *PollService) Start(ch chan interface{}) error {
	ticker := time.NewTicker(p.Interval)
	errCh := make(chan error)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := p.Poll(ch); err != nil {
					errCh <- err
				}
			case <-p.done:
				ticker.Stop()
				errCh <- nil
				break
			}
		}
	}()
	return <-errCh
}

func (p *PollService) Stop() {
	p.done <- struct{}{}
}
