package meta

import "time"

type Meta struct {
	LeagueID  string
	MatchID   string
	HomeID    string
	AwayID    string
	OddID     []string
	MatchTime time.Time
}
