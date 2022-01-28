package jc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupServer(t *testing.T, qs map[string]string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.RawQuery
		if f, ok := qs[q]; ok {
			b, err := ioutil.ReadFile(f)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Fprintln(w, b)
			return
		}
		http.Error(w, "not found", 400)
	}))
}

func Test_Client(t *testing.T) {
	cases := map[string]struct {
		start, end time.Time
		bettype    string
		qs         map[string]string
	}{
		"base": {
			start:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			end:     time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
			bettype: "HAD",
			qs: map[string]string{
				"jsontype=odds_had.aspx": "testdata/had.json",
				"enddate=20220201&jsontype=search_result.aspx&startdate=20220101&teamid=default": "testdata/historical.json",
			},
		},
	}
	for name := range cases {
		tc := cases[name]
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			ts := setupServer(t, tc.qs)
			defer ts.Close()
			c := New(ts.URL)
			if _, err := c.getInstant(ctx, tc.bettype); err != nil {
				t.Fatal(err)
			}
			if _, err := c.getHistorical(ctx, tc.start, tc.end); err != nil {
				t.Fatal(err)
			}
		})
	}
}
