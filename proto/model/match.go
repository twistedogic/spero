package model

func (m *Match) Final() *Score {
	score := new(Score)
	if first := m.GetFirstHalf(); first != nil {
		score = score.Add(first)
	}
	if second := m.GetSecondHalf(); second != nil {
		score = score.Add(second)
	}
	if et := m.GetExtraTime(); et != nil {
		score = score.Add(et)
	}
	return score
}
