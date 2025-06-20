package sNet

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"
)

// Request define an API Request
type Request struct {
	method   string
	endpoint string
	query    url.Values
	header   http.Header
	body     io.Reader
}

func NewRequest() *Request {
	return &Request{
		query:  url.Values{},
		header: http.Header{},
	}
}

func (r *Request) Clone() InterRequest {

	query := make(url.Values, len(r.query))
	for k, v := range r.query {
		query[k] = append([]string{}, v...)
	}

	return &Request{
		method:   r.method,
		endpoint: r.endpoint,
		query:    query,
		header:   r.header.Clone(),
		body:     nil,
	}
}

func (r *Request) String() string {
	return fmt.Sprintf("Request: method: %s  ..  endpoint: %s  ..  query: %s  ..  header: %s  ..  body: %s",
		r.method, r.endpoint, r.query, r.header, r.body)
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) EndPoint() string {
	return r.endpoint
}

func (r *Request) MethodGet() InterRequest {
	r.method = http.MethodGet
	return r
}

func (r *Request) MethodPost() InterRequest {
	r.method = http.MethodPost
	return r
}

func (r *Request) MethodDelete() InterRequest {
	r.method = http.MethodDelete
	return r
}

func (r *Request) SetEndPoint(endpoint string) InterRequest {
	r.endpoint = endpoint
	return r
}

func (r *Request) SetParam(key, value string) {
	r.query.Set(key, value)
}

func (r *Request) AddParam(key, value string) {
	r.query.Add(key, value)
}

func (r *Request) DelParam(key string) {
	r.query.Del(key)
}

func (r *Request) SetHeader(key, value string) {
	r.header.Set(key, value)
}

func (r *Request) AddHeader(key, value string) {
	r.header.Add(key, value)
}

func (r *Request) DelHeader(key string) {
	r.header.Del(key)
}

func (r *Request) SetBody(body io.Reader) {
	r.body = body
}

func (r *Request) Call(ctx context.Context, service InterService, client InterClient) (data []byte, err error) {
	if service == nil {
		return nil, fmt.Errorf("service is nil")
	}

	rawURL, err := url.JoinPath(service.GetUrl(), r.endpoint)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath(): %s ", err)
	}

	// Build the final URL securely
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %s", err)
	}

	// Add correctly coded parameters to the URL
	u.RawQuery = r.query.Encode()

	if service.Debug() {
		sLog.Debug("%s .. %s", u.String(), r.String())
	}

	request, err := http.NewRequest(r.method, u.String(), r.body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(): %s ", err)
	}

	if ctx == nil {
		ctx_to, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ctx = ctx_to
	}
	request = request.WithContext(ctx)

	request.Header = r.header

	if client == nil {
		client = NewClient(nil, nil)
	}

	response, err := client.Do(request)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, err
		}
		return nil, fmt.Errorf("client.Do(): %s ", err)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll(): %s ", err)
	}

	defer func() {
		close_err := response.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && close_err != nil {
			err = close_err
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d", response.StatusCode)
	}

	return data, nil
}
