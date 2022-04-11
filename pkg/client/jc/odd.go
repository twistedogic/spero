package jc

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/twistedogic/spero/proto/model"
)

const (
	hadodds = "hadodds"
	fhaodds = "fhaodds"
)

type matchOdds struct {
	Odds []*model.MatchOdd
}

func parseOdd(input string) (float32, error) {
	tokens := strings.Split(input, "@")
	if len(tokens) != 2 {
		return 0, errors.Errorf("odd %s is invalid", input)
	}
	odd, err := strconv.ParseFloat(tokens[1], 32)
	return float32(odd), err
}

func parseHadOdd(t model.MatchOdd_Type, matchId string, i interface{}) (*model.MatchOdd, error) {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("input is not type of map[string]interface{}")
	}
	odds := make([]*model.Odd, 0, len(model.Outcome_value))
	for key, outcome := range model.Outcome_value {
		val, ok := m[key]
		if !ok {
			return nil, errors.Errorf("key %s not found", key)
		}
		odd, err := parseOdd(val.(string))
		if err != nil {
			return nil, err
		}
		odds = append(odds, &model.Odd{Outcome: model.Outcome(outcome), Odd: odd})
	}
	id, ok := m["ID"]
	if !ok {
		return nil, errors.Errorf("id not found")
	}
	return &model.MatchOdd{
		Id:      id.(string),
		MatchId: matchId,
		Odds:    odds,
		Type:    t,
	}, nil
}

func (m *matchOdds) parseMatchOdd(t model.MatchOdd_Type, matchId string, i interface{}) error {
	if m.Odds == nil {
		m.Odds = make([]*model.MatchOdd, 0)
	}
	mm, err := parseHadOdd(t, matchId, i)
	if err != nil {
		return err
	}
	m.Odds = append(m.Odds, mm)
	return nil
}

func (m *matchOdds) UnmarshalJSON(b []byte) error {
	matches := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(b, &matches); err != nil {
		return err
	}
	for _, match := range matches {
		matchId := match["matchID"].(string)
		for key, val := range match {
			switch key {
			case hadodds:
				if err := m.parseMatchOdd(model.MatchOdd_HAD, matchId, val); err != nil {
					return errors.Wrapf(err, "parse %s odd for match %s", hadodds, matchId)
				}
			case fhaodds:
				if err := m.parseMatchOdd(model.MatchOdd_FHA, matchId, val); err != nil {
					return errors.Wrapf(err, "parse %s odd for match %s", fhaodds, matchId)
				}
			}
		}
	}
	return nil
}
