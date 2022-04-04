package jc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
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

func Test_chunkByDuration(t *testing.T) {
	cases := map[string]struct {
		start, end time.Time
		dur        time.Duration
		want       []time.Time
	}{
		"base": {
			start: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC),
			dur:   24 * time.Hour,
			want: []time.Time{
				time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC),
			},
		},
		"cap at end": {
			start: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 1, 2, 10, 0, 0, 0, time.UTC),
			dur:   24 * time.Hour,
			want: []time.Time{
				time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 2, 10, 0, 0, 0, time.UTC),
			},
		},
		"reverse": {
			start: time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			dur:   24 * time.Hour,
			want: []time.Time{
				time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for name := range cases {
		tc := cases[name]
		t.Run(name, func(t *testing.T) {
			got := chunkByDuration(tc.dur, tc.start, tc.end)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}
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
