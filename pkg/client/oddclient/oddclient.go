package oddclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Jeffail/benthos/v3/public/service"

	"github.com/twistedogic/spero/pkg/client"
	"github.com/twistedogic/spero/pkg/message"
	"github.com/twistedogic/spero/pkg/poll"
)

const (
	DefaultURL = "https://bet.hkjc.com/football/getJSON.aspx"
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
