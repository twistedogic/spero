package jc

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/twistedogic/spero/pkg/client"
	"github.com/twistedogic/spero/proto/model"
)

type response []content

type content struct {
	Matches []json.RawMessage `json:"matches"`
}

func (r response) flatten() []json.RawMessage {
	out := make([]json.RawMessage, 0)
	for _, content := range r {
		out = append(out, content.Matches...)
	}
	return out
}

func (r response) unmarshal(i interface{}) error {
	b, err := json.Marshal(r.flatten())
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &i)
}

func (r response) toMatches() ([]*model.Match, error) {
	var matches []Match
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

func (r response) toOdds() ([]*model.MatchOdd, error) {
	var odds matchOdds
	err := r.unmarshal(&odds)
	return odds.Odds, err
}

func parseResponse(b []byte) (client.Result, error) {
	var r response
	var res client.Result
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
