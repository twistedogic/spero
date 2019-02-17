package match

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/twistedogic/spero/pkg/schema"
)

type Detail struct {
	YellowCards     ValueInt
	RedCards        ValueInt
	Substitutions   ValueInt
	BallPossession  ValueInt
	FreeKicks       ValueInt
	GoalKicks       ValueInt
	ThrowIns        ValueInt
	Offsides        ValueInt
	CornerKicks     ValueInt
	ShotsOnTarget   ValueInt
	ShotsOffTarget  ValueInt
	Saves           ValueInt
	Fouls           ValueInt
	Injuries        ValueInt
	Penalties       ValueInt
	ShotsBlocked    ValueInt
	DangerousAttack ValueInt
	BallSafe        ValueInt
	Attack          ValueInt
	GoalAttempts    ValueInt
}

func penaltiesString(s string) int {
	p := 0
	for _, v := range strings.Split(s, "/") {
		i, _ := strconv.Atoi(v)
		p += i
	}
	return p
}

func (d *Detail) UnmarshalJSON(b []byte) error {
	var result struct {
		Doc []struct {
			Data struct {
				Values map[string]struct {
					Name  string          `json:"name"`
					Value json.RawMessage `json:"value"`
				} `json:"values`
			} `json:"data"`
		} `json:"doc"`
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return err
	}
	for _, doc := range result.Doc {
		for _, v := range doc.Data.Values {
			var container map[string]interface{}
			key := strcase.ToCamel(v.Name)
			if err := json.Unmarshal(v.Value, &container); err != nil {
				return err
			}
			var home int
			var away int
			for k, val := range container {
				var intVal int
				switch v := val.(type) {
				case string:
					intVal = penaltiesString(v)
				case int:
					intVal = v
				case float64:
					intVal = int(v)
				}
				switch k {
				case "home":
					home = intVal
				case "away":
					away = intVal
				}
			}
			value := ValueInt{Home: home, Away: away}
			if err := schema.SetField(d, key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

type ValueInt struct {
	Home int `json:"home"`
	Away int `json:"away"`
}
