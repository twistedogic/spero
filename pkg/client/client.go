package client

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

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
	return ioutil.ReadAll(res.Body)
}

func GetByte(ctx context.Context, url string) ([]byte, error) {
	return requestBytes(ctx, http.MethodGet, url, nil)
}

func PostByte(ctx context.Context, url string, body io.Reader) ([]byte, error) {
	return requestBytes(ctx, http.MethodPost, url, body)
}
