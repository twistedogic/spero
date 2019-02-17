package oddclient

import (
	"fmt"
	"strings"

	"github.com/twistedogic/spero/pkg/client/base"
)

type Client struct {
	*base.Client
	BaseURL string
}

func New(oddURL string) *Client {
	client := base.New()
	return &Client{client, oddURL}
}

//https://bet.hkjc.com/football/getJSON.aspx?jsontype=odds_fha.aspx
func (c *Client) GetOddTable(bettype string, value interface{}) error {
	u := fmt.Sprintf("%s?jsontype=odds_%s.aspx", c.BaseURL, strings.ToLower(bettype))
	return c.GetJSON(u, value)
}
