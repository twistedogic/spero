package metascraper

import (
	"reflect"
	"testing"

	"github.com/twistedogic/spero/pkg/schema/match"
)

func setup() *MetaScraper {
	eventURL := "https://lsc.fn.sportradar.com/hkjc/en"
	return New(eventURL, 5)
}

func TestMetaScraper_GetMatchDetail(t *testing.T) {
	m := setup()
	id := 14728777
	detail, err := m.GetMatchDetail(id)
	if err != nil {
		t.Error(err)
	}
	if reflect.DeepEqual(detail, match.Detail{}) {
		t.Fail()
	}
}

func TestMetaScraper_GetMatchEvents(t *testing.T) {
	m := setup()
	id := 14971653
	events, err := m.GetMatchEvents(id)
	if err != nil {
		t.Error(err)
	}
	if len(events) == 0 {
		t.Fail()
	}
}

func TestMetaScraper_GetMatchSituations(t *testing.T) {
	m := setup()
	id := 14971653
	situations, err := m.GetMatchSituations(id)
	if err != nil {
		t.Error(err)
	}
	if len(situations) == 0 {
		t.Fail()
	}
}
func TestMetaScraper_GetMatchMeta(t *testing.T) {
	m := setup()
	id := 14971653
	input := match.Match{ID: id}
	output, err := m.GetMatchMeta(input)
	if err != nil {
		t.Error(err)
	}
	detail := output.Detail
	data := output.Match
	events := output.Events
	situations := output.Situations
	switch {
	case reflect.DeepEqual(detail, match.Detail{}):
		t.Log(detail)
		t.Error("empty detail")
	case !reflect.DeepEqual(data, input):
		t.Error("match not equal to input")
	case len(events) == 0:
		t.Error("empty events")
	case len(situations) == 0:
		t.Error("empty situations")
	}
}

func TestMetaScraper_GetMatches(t *testing.T) {
	m := setup()
	offset := -1
	matches, err := m.GetMatches(offset)
	if err != nil {
		t.Error(err)
	}
	if len(matches) == 0 {
		t.Fail()
	}
}
