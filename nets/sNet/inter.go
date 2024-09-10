package sNet

import (
	"context"
	"io"
	"net/http"
)

type InterService interface {
	Log()

	Port() int
	GetUri() string
	GetUrl() string
	LocalAddr() string

	HandlePath(string) string
}

type InterRequest interface {
	String() string

	// Options
	MethodGet() InterRequest
	MethodPost() InterRequest
	SetEndPoint(string) InterRequest

	SetParam(string, string)
	AddParam(string, string)
	DelParam(string)

	SetHeader(string, string)
	AddHeader(string, string)
	DelHeader(string)

	SetBody(io.Reader)

	Call(context.Context, InterService, InterClient) ([]byte, error)
}

type InterClient interface {
	Do(*http.Request) (*http.Response, error)
}
