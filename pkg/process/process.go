package runner

import (
	"context"
	"time"

	"github.com/twistedogic/store"
	"google.golang.org/protobuf/proto"

	"github.com/twistedogic/spero/pkg/client"
	"github.com/twistedogic/spero/proto/model"
)

type Writer struct {
	store.Store
}

func (w Writer) writeMatch(ctx context.Context, m *model.Match) error {
	b, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	item := store.Item{Key: []byte(m.GetId()), Data: b}
	return w.Set(ctx, item)
}

func (w Writer) writeOdd(ctx context.Context, m *model.MatchOdd) error {
	b, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	item := store.Item{Key: []byte(m.GetId()), Data: b}
	return w.Set(ctx, item)
}

func (w Writer) WriteMatches(ctx context.Context, matches *model.Match) error {
	for _, m := range matches {
		if err := w.writeMatch(m); err != nil {
			return err
		}
	}
	return err
}

func (w Writer) WriteOdds(ctx context.Context, odds *model.MatchOdd) error {
	for _, m := range odds {
		if err := w.writeMatch(m); err != nil {
			return err
		}
	}
	return err
}

type Runner struct {
	client.Source
	Writer
}

func New(c client.Source, s store.Store) Runner {
	return Runner{Source: c, Writer: Writer{Store: s}}
}

func (r Runner) store(ctx context.Context, res client.Result) error {
	if err := r.WriteMatches(ctx, res.Matches); err != nil {
		return err
	}
	if err := r.WriteOdds(ctx, res.Odds); err != nil {
		return err
	}
	return nil
}

func (r Runner) StoreRange(ctx context.Context, start, end time.Time) error {
	res, err := r.GetMatchesByDates(ctx, start, end)
	if err != nil {
		return err
	}
	return r.store(ctx, res)
}

func (r Runner) StoreInstant(ctx context.Context, t string) error {
	res, err := r.GetCurrentMatches(ctx)
	if err != nil {
		return err
	}
	return r.store(ctx, res)
}
