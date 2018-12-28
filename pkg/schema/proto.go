package schema

import (
	"strconv"
	"strings"

	pb "github.com/twistedogic/spero/pb"
	"github.com/twistedogic/spero/pkg/metric"
	"github.com/twistedogic/spero/pkg/tag"
)

const (
	oddTag = "odd"
)

func ParsePayout(pay string) (int64, float64) {
	tokens := strings.Split(pay, "@")
	if len(tokens) == 2 {
		minBuy, _ := strconv.ParseInt(tokens[0], 10, 64)
		payout, _ := strconv.ParseFloat(tokens[1], 64)
		return minBuy, payout
	}
	payout, _ := strconv.ParseFloat(tokens[0], 64)
	return 0, payout
}

func (m Match) ToProto() []*pb.Odd {
	fields := tag.GetTaggedFields(m, oddTag)
	odds := make([]*pb.Odd, 0)
	var matchID string
	var league League
	var home, away Team
	for _, field := range fields {
		tagged, value := field.Tag, field.Value
		switch tagged {
		case "id":
			matchID = value.String()
		case "league":
			if v, ok := value.Interface().(League); ok {
				league = v
			}
		case "home":
			if v, ok := value.Interface().(Team); ok {
				home = v
			}
		case "away":
			if v, ok := value.Interface().(Team); ok {
				away = v
			}
		}
	}
	for _, field := range fields {
		if strings.Contains(field.Tag, "value") {
			kv := tag.ParseTag(field.Tag)
			if bettype, ok := kv["type"]; ok {
				minbet, payout := ParsePayout(field.Value.String())
				outcome := field.Name
				metric.OddMetric.WithLabelValues(matchID, bettype, home.TeamNameEN, away.TeamNameEN, outcome).Set(payout)
				odd := &pb.Odd{
					Id:      matchID,
					Home:    home.ToProto(),
					Away:    away.ToProto(),
					League:  league.ToProto(),
					Outcome: field.Name,
					BetType: bettype,
					MinBet:  minbet,
					Payout:  payout,
				}
				odds = append(odds, odd)
			}
		}
	}
	return odds
}
