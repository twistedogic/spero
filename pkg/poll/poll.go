package poll

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/pkg/errors"

	"github.com/twistedogic/spero/pkg/message"
)

type PollFunc func(context.Context) ([]*service.Message, error)

type Poller struct {
	*sync.RWMutex
	pollFunc PollFunc
	interval time.Duration
	buffer   chan *service.Message
	acked    string
	stopFunc context.CancelFunc
}

func New(poll PollFunc, interval time.Duration) *Poller {
	p := &Poller{
		RWMutex:  new(sync.RWMutex),
		pollFunc: poll,
		buffer:   make(chan *service.Message, 1),
		interval: interval,
	}
	return p
}

func (p *Poller) updateAcked(hash string) {
	p.Lock()
	defer p.Unlock()
	p.acked = hash
}

func (p *Poller) Acked() string {
	p.RLock()
	defer p.RUnlock()
	return p.acked
}

func (p *Poller) Ack(hash string) service.AckFunc {
	return func(ctx context.Context, err error) error {
		if err == nil {
			p.updateAcked(hash)
		}
		return err
	}
}

func (p *Poller) addToBuffer(msg *service.Message) {
	hash, ok := message.GetContentHash(msg)
	switch {
	case !ok:
		p.buffer <- msg
	case hash != p.Acked():
		p.buffer <- msg
	}
}

func (p *Poller) poll(ctx context.Context) {
	msgs, err := p.pollFunc(ctx)
	if err != nil {
		log.Println(errors.Wrap(err, "poll error"))
	}
	for _, msg := range msgs {
		p.addToBuffer(msg)
	}
}

func (p *Poller) Poll(ctx context.Context) {
	for {
		timeoutCtx, _ := context.WithTimeout(ctx, p.interval)
		select {
		case <-ctx.Done():
			return
		default:
			go p.poll(timeoutCtx)
		}
		<-time.After(p.interval)
	}
}

func (p *Poller) Connect(ctx context.Context) error {
	doneCtx, cancel := context.WithCancel(context.Background())
	p.stopFunc = cancel
	go p.Poll(doneCtx)
	for {
		if len(p.buffer) != 0 {
			break
		}
	}
	return ctx.Err()
}

func (p *Poller) Close(ctx context.Context) error {
	defer p.stopFunc()
	return nil
}

func (p *Poller) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	ack := func(c context.Context, e error) error { return e }
	msg := <-p.buffer
	hash, ok := message.GetContentHash(msg)
	if !ok {
		return msg, ack, nil
	}
	return msg, p.Ack(hash), nil
}
