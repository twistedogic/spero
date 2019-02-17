package odd

func SliceEqual(a, b []Match) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Statuslastupdated != b[i].Statuslastupdated {
			return false
		}
	}
	return true
}
