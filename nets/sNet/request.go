package sNet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type secType int

const (
	secTypeNone   secType = iota
	secTypeSigned         // private Request
)

type Params map[string]interface{}

func NewParams() Params {
	return make(Params)
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

// Request define an API Request
type Request struct {
	method     string
	endpoint   string
	query      url.Values
	recvWindow string
	secType    secType
	header     http.Header
	params     []byte
	fullURL    string
	body       io.Reader
}

func (r *Request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	return nil
}

func (r *Request) Get() *Request {
	r.method = http.MethodGet
	return r
}

func (r *Request) Post() *Request {
	r.method = http.MethodPost
	return r
}

func (r *Request) Signed() *Request {
	r.secType = secTypeSigned
	return r
}

func (r *Request) EndPoint(endpoint string) *Request {
	r.endpoint = endpoint
	return r
}

// setParam set param with key/value to query string
func (r *Request) setParam(key string, value interface{}) {
	r.validate()
	r.query.Set(key, fmt.Sprintf("%v", value))
}

func (r *Request) SetParams(params Params) error {
	switch r.method {
	case http.MethodGet:
		for k, v := range params {
			r.setParam(k, v)
		}
	case http.MethodPost:
		jsonData, err := json.Marshal(params)
		if err != nil {
			return fmt.Errorf("error marshalling query: %s", err)
		}
		r.params = jsonData
	}

	return nil
}
