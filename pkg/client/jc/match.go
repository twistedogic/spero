package jc

import (
	"strconv"
	"time"

	"github.com/twistedogic/spero/proto/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type League struct {
	ID   string `json:"leagueID"`
	Name string `json:"leagueNameEN"`
}

func (l League) toModel() *model.League { return &model.League{Id: l.ID, Name: l.Name} }

type Team struct {
	ID   string `json:"teamID"`
	Name string `json:"teamNameEN"`
}

func (t Team) toModel() *model.Team { return &model.Team{Id: t.ID, Name: t.Name} }

type Score struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

func toInt32(s string) (int32, error) {
	if len(s) == 0 {
		return 0, nil
	}
	i, err := strconv.Atoi(s)
	return int32(i), err
}

func (s Score) toModel() (*model.Score, error) {
	h, err := toInt32(s.Home)
	if err != nil {
		return nil, err
	}
	a, err := toInt32(s.Away)
	if err != nil {
		return nil, err
	}
	return &model.Score{Home: h, Away: a}, nil
}

type Match struct {
	ID     string  `json:"matchID"`
	Date   string  `json:"matchTime"`
	Status string  `json:"matchStatus"`
	League League  `json:"league"`
	Home   Team    `json:"homeTeam"`
	Away   Team    `json:"awayTeam"`
	Scores []Score `json:"accumulatedscore"`
}

func (m Match) toModel() (mm *model.Match, err error) {
	mm = new(model.Match)
	mm.Id = m.ID
	date, err := time.Parse(time.RFC3339, m.Date)
	if err != nil {
		return
	}
	mm.Date = timestamppb.New(date)
	mm.League = m.League.toModel()
	mm.Home = m.Home.toModel()
	mm.Away = m.Away.toModel()
	mm.Scores = make([]*model.Score, len(m.Scores))
	for i := range mm.Scores {
		score, err := m.Scores[i].toModel()
		if err != nil {
			return mm, err
		}
		mm.Scores[i] = score
	}
	return
}
