package jc

import (
	"io/ioutil"
	"testing"
)

func Test_parseResponse(t *testing.T) {
	cases := map[string]struct {
		file string
	}{
		"had": {
			file: "testdata/had.json",
		},
		"historical": {
			file: "testdata/historical.json",
		},
	}
	for name := range cases {
		tc := cases[name]
		t.Run(name, func(t *testing.T) {
			b, err := ioutil.ReadFile(tc.file)
			if err != nil {
				t.Fatal(err)
			}
			got, err := parseResponse(b)
			if err != nil {
				t.Fatal(err)
			}
			if len(got.Matches) == 0 {
				t.Fatal("'matches' is empty")
			}
			if len(got.Odds) == 0 {
				t.Fatal("'odds' is empty")
			}
		})
	}
}
