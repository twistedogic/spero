package jc

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/twistedogic/spero/proto/model"
)

const (
	hadodds = "hadodds"
	fhaodds = "fhaodds"
)

type MatchOdds struct {
	Odds []*model.MatchOdd
}

func parseHadOdd(key string, i interface{}, mm *model.MatchOdd) error {
	m, ok := i.(map[string]string)
	if !ok {
		return errors.Errorf("input is not type of map[string]string")
	}
	return nil
}

func (m *MatchOdds) UnmarshalJSON(b []byte) error {
	matches := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	for _, m := range matches {
		for key, val := range m {
			switch key {
			case hadodds:
			}
		}
	}
	return nil
}
