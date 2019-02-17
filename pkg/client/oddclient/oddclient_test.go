package oddclient

import (
	"reflect"
	"testing"

	"github.com/twistedogic/spero/pkg/schema/odd"
)

func TestNew(t *testing.T) {
	oddURL := "http://localhost"
	c := New(oddURL)
	if c.BaseURL != oddURL {
		t.Fail()
	}
}

func TestGetOddTable(t *testing.T) {
	oddURL := "https://bet.hkjc.com/football/getJSON.aspx"
	c := New(oddURL)
	var oddtable odd.OddTable
	if err := c.GetOddTable("had", &oddtable); err != nil {
		t.Error(err)
	}
	for _, m := range oddtable.Matches {
		for _, outcome := range m.GetOutcome() {
			if reflect.DeepEqual(outcome, odd.Outcome{}) {
				t.Fail()
			}
		}
	}
}
