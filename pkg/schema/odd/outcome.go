package odd

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/twistedogic/spero/pkg/tag"
)

type OddEnum uint

func (o OddEnum) Key() string {
	switch {
	case o == HAD:
		return "had"
	case o == HFT:
		return "hft"
	case o == FHA:
		return "fha"
	case o == HHA:
		return "hha"
	}
	return ""
}

func (o *OddEnum) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(o.Key())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func ToBetTypeEnum(key string) (OddEnum, bool) {
	v, ok := mapping[key]
	return v, ok
}

const (
	HAD OddEnum = iota
	HFT
	FHA
	HHA
)

var mapping = map[string]OddEnum{
	"had": HAD,
	"hft": HFT,
	"fha": FHA,
	"hha": HAD,
}

const TAG_NAME = "odd"

type Outcome struct {
	MatchID    string
	OddID      string
	Type       OddEnum
	Outcome    string
	Odd        float64
	LastUpdate time.Time
}

func parseOdd(odd string) float64 {
	val := odd
	if strings.Contains(odd, "@") {
		val = strings.Split(odd, "@")[1]
	}
	f, _ := strconv.ParseFloat(val, 64)
	return f
}

func ParseOutcome(v interface{}) []Outcome {
	out := make([]Outcome, 0)
	fields := tag.GetTaggedFields(v, TAG_NAME)
	var ID string
	for _, f := range fields {
		if strings.Contains(f.Tag, "id") {
			ID = f.Value.String()
			continue
		}
	}
	for _, f := range fields {
		if t, ok := tag.ParseTag(f.Tag)["type"]; ok {
			t = strings.ToLower(t)
			if enum, ok := mapping[t]; ok {
				val := parseOdd(f.Value.String())
				out = append(out, Outcome{
					Type:    enum,
					Outcome: f.Name,
					Odd:     val,
					OddID:   ID,
				})
			}
		}
	}
	return out
}
