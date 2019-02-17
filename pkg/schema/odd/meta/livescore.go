package meta

import (
	"encoding/json"
	"strconv"
)

type Livescore struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

func (l *Livescore) UnmarshalJSON(b []byte) error {
	container := make(map[string]string)
	if err := json.Unmarshal(b, &container); err != nil {
		return err
	}
	for k, v := range container {
		val, _ := strconv.Atoi(v)
		switch k {
		case "home":
			l.Home = val
		case "away":
			l.Away = val
		}
	}
	return nil
}
