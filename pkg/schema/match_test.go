package schema

import (
	"sort"
	"testing"
)

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
func TestByLastUpdate(t *testing.T) {
	first := Match{Statuslastupdated: "2018-01-02T15:04:05-07:00"}
	second := Match{Statuslastupdated: "2017-01-02T15:04:05-07:00"}
	third := Match{Statuslastupdated: "2016-01-02T15:04:05-07:00"}
	input := []Match{third, second, first}
	expect := []Match{first, second, third}
	sort.Sort(ByLastUpdate{input})
	if !SliceEqual(expect, input) {
		t.Fail()
	}

}
