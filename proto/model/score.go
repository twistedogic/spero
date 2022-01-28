package model

func (s *Score) Add(o *Score) *Score {
	return &Score{
		Home: s.Home + o.Home,
		Away: s.Away + o.Away,
	}
}
