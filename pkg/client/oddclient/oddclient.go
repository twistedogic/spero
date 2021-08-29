package oddclient

import (
	"fmt"
	"strings"

	"github.com/twistedogic/spero/pkg/client/base"
)

const (
	DefaultURL = "https://bet.hkjc.com/football/getJSON.aspx"
)

type Client struct {
	base.Client
	BaseURL string
}

func New(oddURL string) Client {
	client := base.New()
	return Client{client, oddURL}
}

func NewWithDefault() Client {
	return New(DefaultURL)
}

func (c *Client) GetOddTable(bettype string) ([]byte, error) {
	u := fmt.Sprintf("%s?jsontype=odds_%s.aspx", c.BaseURL, strings.ToLower(bettype))
	return c.GetByte(u)
}
