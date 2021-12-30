package jc

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Jeffail/benthos/v3/public/service"

	"github.com/twistedogic/spero/pkg/client"
	"github.com/twistedogic/spero/pkg/message"
	"github.com/twistedogic/spero/pkg/poll"
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

type data struct {
	Matches []json.RawMessage `json:"matches"`
}

func flattenOddResponse(b []byte) ([]*service.Message, error) {
	items := make([]data, 0)
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}
	msgs := make([]*service.Message, 0)
	for _, item := range items {
		for _, match := range item.Matches {
			msgs = append(msgs, message.NewContentHashMessage(match))
		}
	}
	return msgs, nil
}

func toQueryString(kv map[string]string) string {
	terms := make([]string, 0, len(kv))
	for k, v := range kv {
		terms = append(terms, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(terms)
	return strings.Join(terms, "&")
}

func getResultURL(base string, start, end time.Time) string {
	terms := make(map[string]string)
	terms[typeKey] = resultQuery
	terms[startDateKey] = start.Format(resultDateFormat)
	terms[endDateKey] = end.Format(resultDateFormat)
	return fmt.Sprintf("%s?%s&teamid=default", base, toQueryString(terms))

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

func (c Client) getOddTable(ctx context.Context, bettype string) ([]byte, error) {
	u := fmt.Sprintf("%s?jsontype=odds_%s.aspx", c.BaseURL, strings.ToLower(bettype))
	return client.GetByte(ctx, u)
}

func (c Client) PollOdd(bettype string) poll.PollFunc {
	return func(ctx context.Context) ([]*service.Message, error) {
		b, err := c.getOddTable(ctx, bettype)
		if err != nil {
			return nil, err
		}
		return flattenOddResponse(b)
	}
}

func (c Client) getResults(ctx context.Context, start, end time.Time) ([]byte, error) {
	u := getResultURL(c.BaseURL, start, end)
	return client.PostByte(ctx, u, nil)
}

func (c Client) PollResult(start, end time.Time) poll.PollFunc {
	timeCh := make(chan time.Time)
	ackCh := make(chan bool)
	go func() {
		current := start
		for current.Before(end) {
			timeCh <- current
			if ack := <-ackCh; ack {
				current = current.Add(MONTH)
			}
		}
	}()
	return func(ctx context.Context) (msgs []*service.Message, err error) {
		defer func() {
			ackCh <- err == nil
		}()
		current := <-timeCh
		b, err := c.getResults(ctx, current, current.Add(MONTH))
		if err != nil {
			return nil, err
		}
		msgs, err = flattenOddResponse(b)
		if err != nil {
			return nil, err
		}
		return msgs, nil
	}
}
