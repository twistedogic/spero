package jc

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/twistedogic/spero/proto/model"
)

type Response []Content

type Content struct {
	Matches []json.RawMessage `json:"matches"`
}

func (r Response) flatten() []json.RawMessage {
	out := make([]json.RawMessage, 0)
	for _, content := range r {
		out = append(out, content.Matches...)
	}
	return out
}

func (r Response) unmarshal(i interface{}) error {
	b, err := json.Marshal(r.flatten())
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &i)
}

func (r Response) toMatches() ([]*model.Match, error) {
	matches := make([]Match, 0)
	if err := r.unmarshal(&matches); err != nil {
		return nil, err
	}
	out := make([]*model.Match, len(matches))
	for i, m := range matches {
		mm, err := m.toModel()
		if err != nil {
			return nil, errors.Wrapf(err, "convert match %s", m.ID)
		}
		out[i] = mm
	}
	return out, nil
}

func (r Response) toOdds() ([]*model.MatchOdd, error) {
	odds := MatchOdds{}
	err := r.unmarshal(&odds)
	return odds.Odds, err
}

type result struct {
	Matches []*model.Match
	Odds    []*model.MatchOdd
}

func (r *result) UnmarshalJSON(b []byte) error {
	data := make([]json.RawMessage, 0)
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	return nil
}

func parseResponse(b []byte) (result, error) {
	r, res := make(Response, 0), result{}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return res, err
	}
	res.Matches, err = r.toMatches()
	if err != nil {
		return res, err
	}
	res.Odds, err = r.toOdds()
	return res, err
}
