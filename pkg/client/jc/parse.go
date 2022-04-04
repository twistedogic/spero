package jc

import (
	"encoding/json"
	"fmt"

	"github.com/twistedogic/spero/proto/model"
)

type Matches []Content

type Content struct {
	Matches []Match `json:"matches"`
}

type League struct {
	ID   string `json:"leagueID"`
	Name string `json:"leagueNameEN"`
}

func (l League) toModel() *model.League {
	return &model.League{Id: l.ID, Name: l.Name}
}

type Team struct {
	ID   string `json:"teamID"`
	Name string `json:"teamNameEN"`
}

type Score struct {
	Home string `json:"home"`
	Away string `json:"away"`
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

type Odds []*model.Odd

func (o Odds) UnmarshalJSON(b []byte) error {
	i := make(map[string]interface{})

}

type result struct {
	Matches []*model.Match
	Odds    []*model.Odd
}

func (r *result) UnmarshalJSON(b []byte) error {
	data := make([]json.RawMessage, 0)
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	return nil
}

func parseResponse(b []byte) (result, error) {
	m, res := make(Matches, 0), result{}
	if err := json.Unmarshal(b, &m); err != nil {
		return res, err
	}
	mBytes, err := json.MarshalIndent(&m, "", "  ")
	if err != nil {
		return res, err
	}
	fmt.Println(string(mBytes))
	err = json.Unmarshal(mBytes, &res)
	return res, err
}
