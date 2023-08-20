package httpclient

import (
	"context"
	"kmf-proxy/pkg/domain"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewRepository() *Client {
	return &Client{http.DefaultClient}
}

func (c *Client) Do(ctx context.Context, request domain.Request) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, request.Method, request.URL, nil)
	if err != nil {
		return nil, err
	}

	for header, value := range request.Headers {
		req.Header.Add(header, value)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
