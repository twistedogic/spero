package base

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func New() Client {
	client := &http.Client{}
	return Client{Client: client}
}

func (c Client) GetByte(u string) ([]byte, error) {
	res, err := c.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
