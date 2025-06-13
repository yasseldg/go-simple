package sNet

import (
	"context"
	"io"
	"net/http"
)

type InterService interface {
	String() string
	Log()

	Clone() InterService

	SetPathPrefix(string)

	SetDebug(bool)
	Debug() bool

	Url() string
	Port() int
	GetUri() string
	GetUrl() string
	LocalAddr() string
	User() string
	Secret() string

	HandlePath(string) string

	SendObj(end_point string, obj any) error
}

type InterRequest interface {
	String() string

	Clone() InterRequest

	Method() string
	EndPoint() string

	// Options
	MethodGet() InterRequest
	MethodPost() InterRequest
	MethodDelete() InterRequest
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
