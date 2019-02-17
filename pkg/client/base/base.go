package base

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/twistedogic/spero/pkg/errors"
	"github.com/twistedogic/spero/pkg/metric"
)

type Client struct {
	*http.Client
}

func New() *Client {
	client := &http.Client{}
	return &Client{client}
}

func (c *Client) GetJSON(u string, value interface{}) error {
	log.Printf("making request to %s", u)
	statusCode := "-1"
	res, err := c.Get(u)
	if res != nil {
		statusCode = res.Status
		defer res.Body.Close()
	}
	metric.UpdateClientMetric(statusCode)
	if err != nil {
		return errors.NewError(errors.REQ_ERROR, err)
	}
	if err := json.NewDecoder(res.Body).Decode(value); err != nil {
		return errors.NewError(errors.JSON_ERROR, err)
	}
	return nil
}
