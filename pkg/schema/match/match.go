package match

type MatchData struct {
	ID         int `gorm:"PRIMARY_KEY"`
	Match      Match
	Detail     Detail
	Events     []Event
	Situations []Situation
}

func (m MatchData) GetHome() Team {
	return m.Match.Teams.Home
}

func (m MatchData) GetAway() Team {
	return m.Match.Teams.Away
}
