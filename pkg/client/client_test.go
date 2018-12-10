package client

import (
	"testing"
)

func TestNew(t *testing.T) {
	base := "http://localhost"
	period := "1s"
	c, err := New(base, period)
	if err != nil {
		t.Error(err)
	}
	if c.BaseURL != base {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	base := "https://bet.hkjc.com/football/getJSON.aspx"
	period := "1s"
	c, err := New(base, period)
	if err != nil {
		t.Error(err)
	}
	matches := c.get("had")
	if len(matches) == 0 {
		t.Fail()
	}
	t.Log(matches)
}
