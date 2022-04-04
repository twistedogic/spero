package jc

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twistedogic/spero/pkg/client"
)

const (
	DefaultURL = "https://bet.hkjc.com/football/getJSON.aspx"

	startDateKey = "startdate"
	endDateKey   = "enddate"
	typeKey      = "jsontype"

	resultQuery      = "search_result.aspx"
	resultDateFormat = "20060102"

	MONTH = 28 * 24 * time.Hour
)

func chunkByDuration(dur time.Duration, start, end time.Time) []time.Time {
	if start.After(end) {
		start, end = end, start
	}
	out := make([]time.Time, 0)
	current := start
	out = append(out, current)
	for current.Before(end) {
		current = current.Add(dur)
		switch {
		case current.After(end), current.Equal(end):
			out = append(out, end)
		default:
			out = append(out, current)
		}
	}
	return out
}

func getResultURL(base string, start, end time.Time) string {
	terms := make(url.Values)
	terms.Add("teamid", "default")
	terms.Add(typeKey, resultQuery)
	terms.Add(startDateKey, start.Format(resultDateFormat))
	terms.Add(endDateKey, end.Format(resultDateFormat))
	return fmt.Sprintf("%s?%s", base, terms.Encode())
}

type Client struct {
	BaseURL string
}

func New(oddURL string) Client {
	return Client{BaseURL: oddURL}
}

func NewWithDefault() Client {
	return New(DefaultURL)
}

func (c Client) getInstant(ctx context.Context, bettype string) ([]byte, error) {
	u := fmt.Sprintf("%s?jsontype=odds_%s.aspx", c.BaseURL, strings.ToLower(bettype))
	return client.GetByte(ctx, u)
}

func (c Client) getHistorical(ctx context.Context, start, end time.Time) ([]byte, error) {
	u := getResultURL(c.BaseURL, start, end)
	return client.PostByte(ctx, u, nil)
}
