package match

import "encoding/json"

type Fullfeed struct {
	Doc []Doc `json:"doc"`
}

func (f Fullfeed) GetTournaments() []Tournament {
	out := make([]Tournament, 0)
	for _, doc := range f.Doc {
		for _, data := range doc.Data {
			for _, rc := range data.Realcategories {
				out = append(out, rc.Tournaments...)
			}
		}
	}
	return out
}

func (f Fullfeed) GetMatches() []Match {
	out := make([]Match, 0)
	for _, tournament := range f.GetTournaments() {
		out = append(out, tournament.Matches...)
	}
	return out
}

type Doc struct {
	Data []Data `json:"data"`
}

type Data struct {
	Realcategories []struct {
		Doc         string       `json:"_doc"`
		ID          int          `json:"_id"`
		Sid         int          `json:"_sid"`
		Rcid        int          `json:"_rcid"`
		Name        string       `json:"name"`
		Tournaments []Tournament `json:"tournaments"`
	} `json:"realcategories"`
}

type Team struct {
	ID         int    `json:"_id"`
	UID        int    `json:"uid"`
	Virtual    bool   `json:"virtual"`
	Name       string `json:"name"`
	Mediumname string `json:"mediumname"`
	Abbr       string `json:"abbr"`
	Iscountry  bool   `json:"iscountry"`
	Haslogo    bool   `json:"haslogo"`
}

type CardCount struct {
	YellowCount int `json:"yellow_count"`
	RedCount    int `json:"red_count"`
}

type Score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type Periods struct {
	P1 Score `json:"p1"`
	Ft Score `json:"ft"`
}

type Match struct {
	ID       int `json:"_id"`
	Seasonid int `json:"_seasonid"`
	Dt       struct {
		Tzoffset int `json:"tzoffset"`
		Uts      int `json:"uts"`
	} `json:"_dt"`
	Round   int             `json:"round"`
	Week    int             `json:"week"`
	Result  Score           `json:"result"`
	Periods json.RawMessage `json:"periods"`
	Teams   struct {
		Home Team `json:"home"`
		Away Team `json:"away"`
	} `json:"teams"`
	Status struct {
		Name string `json:"name"`
	} `json:"status"`
	Cards struct {
		Home CardCount `json:"home"`
		Away CardCount `json:"away"`
	} `json:"cards"`
}

type Tournament struct {
	ID                   int     `json:"_id"`
	Name                 string  `json:"name"`
	Friendly             bool    `json:"friendly"`
	Seasonid             int     `json:"seasonid"`
	Currentseason        int     `json:"currentseason"`
	Year                 string  `json:"year"`
	Tournamentlevelorder int     `json:"tournamentlevelorder"`
	Outdated             bool    `json:"outdated"`
	Matches              []Match `json:"matches"`
}
