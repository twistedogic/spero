package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/twistedogic/spero/pkg/schema"
)

type Error struct {
	error
	Type       string
	StatusCode int
}

func newError(t string, status int, err error) Error {
	return Error{
		err, t, status,
	}
}

type Client struct {
	*http.Client
	BaseURL string
}

//https://bet.hkjc.com/football/getJSON.aspx?jsontype=odds_fha.aspx
func New(baseURL string) *Client {
	client := &http.Client{}
	return &Client{client, baseURL}
}

func (c *Client) GetMatchesByType(bettype string) ([]schema.Match, error) {
	matches := []schema.Match{}
	var container []struct {
		Matches []schema.Match `json:"matches"`
	}
	u := fmt.Sprintf("%s?jsontype=odds_%s.aspx", c.BaseURL, strings.ToLower(bettype))
	res, err := c.Get(u)
	if err != nil {
		return matches, newError("http", res.StatusCode, err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(b, &container); err != nil {
		return matches, newError("json", 0, err)
	}
	for _, c := range container {
		matches = append(matches, c.Matches...)
	}
	return matches, nil
}
