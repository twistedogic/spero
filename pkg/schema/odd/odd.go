package odd

import (
	"encoding/json"
	"time"

	"github.com/twistedogic/spero/pkg/schema/odd/meta"
	"github.com/twistedogic/spero/pkg/schema/odd/oddtype"
)

const (
	TIMESTAMP_FORAMT = "2006-01-02T15:04:05-07:00"
	DATE_FORMAT      = "2006-01-02-07:00"
)

type Match struct {
	MatchID           string         `json:"matchID" odd:"id"`
	MatchIDinofficial string         `json:"matchIDinofficial"`
	MatchDate         string         `json:"matchDate"`
	League            meta.League    `json:"league,omitempty" odd:"league"`
	HomeTeam          meta.Team      `json:"homeTeam" odd:"home"`
	AwayTeam          meta.Team      `json:"awayTeam" odd:"away"`
	MatchStatus       string         `json:"matchStatus,omitempty"`
	MatchTime         string         `json:"matchTime" odd:"match_time"`
	Statuslastupdated string         `json:"statuslastupdated,omitempty"`
	Livescore         meta.Livescore `json:"livescore,omitempty"`
	Hadodds           oddtype.Had    `json:"hadodds,omitempty" odd:"odd"`
	Fhaodds           oddtype.Fha    `json:"fhaodds,omitempty" odd:"odd"`
	Hftodds           oddtype.Hft    `json:"hftodds,omitempty" odd:"odd"`
	Hhaodds           oddtype.Hha    `json:"hhaodds,omitempty" odd:"odd"`
	HasExtraTimePools bool           `json:"hasExtraTimePools"`
}

func (m Match) GetMatchTime() (time.Time, error) {
	return time.Parse(TIMESTAMP_FORAMT, m.MatchTime)
}

func (m Match) GetUpdateTimestamp() (time.Time, error) {
	return time.Parse(TIMESTAMP_FORAMT, m.Statuslastupdated)
}

func (m Match) GetOutcome() []Outcome {
	out := ParseOutcome(m)
	for i, v := range out {
		v.MatchID = m.MatchID
		v.LastUpdate, _ = m.GetUpdateTimestamp()
		out[i] = v
	}
	return out
}

type OddTable struct {
	Matches []Match `json:"matches,omitempty`
}

func (o *OddTable) UnmarshalJSON(b []byte) error {
	type Container struct {
		Matches []Match `json:"matches,omitempty"`
	}
	var response []Container
	if err := json.Unmarshal(b, &response); err != nil {
		return err
	}
	out := make([]Match, 0)
	for _, r := range response {
		out = append(out, r.Matches...)
	}
	o.Matches = out
	return nil
}
