package sNet

import (
	"net/http"
)

// Client define API client
type Client struct {
	httpClient *http.Client
	do         doFunc
}

type doFunc func(request *http.Request) (*http.Response, error)

// NewClient Create client function for initialising new Bybit client
func NewClient(httpClient *http.Client, do doFunc) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if do == nil {
		do = httpClient.Do
	}

	return &Client{
		httpClient: httpClient,
		do:         do,
	}
}

// Do execute the request
func (c *Client) Do(request *http.Request) (*http.Response, error) {
	return c.do(request)
}
