package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/twistedogic/spero/proto/model"
)

type Result struct {
	Matches []*model.Match
	Odds    []*model.MatchOdd
}

func (r Result) Merge(o Result) Result {
	return Result{
		Matches: append(r.Matches, o.Matches...),
		Odds:    append(r.Odds, o.Odds...),
	}
}

type ByInstant interface {
	GetCurrentMatches(context.Context, model.MatchOdd_Type) (Result, error)
}

type ByDates interface {
	GetMatchesByDates(context.Context, time.Time, time.Time) (Result, error)
}

type Source interface {
	ByDates
	ByInstant
}

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
)

var client = new(http.Client)

func requestBytes(ctx context.Context, method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got response %v from %s", res.StatusCode, url)
	}
	return ioutil.ReadAll(res.Body)
}

func GetByte(ctx context.Context, url string) ([]byte, error) {
	return requestBytes(ctx, http.MethodGet, url, nil)
}

func PostByte(ctx context.Context, url string, body io.Reader) ([]byte, error) {
	return requestBytes(ctx, http.MethodPost, url, body)
}
