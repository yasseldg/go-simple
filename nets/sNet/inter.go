package sNet

import (
	"io"
	"net/http"
)

type InterService interface {
	Log()

	Port() int
	GetUri() string
	GetUrl() string

	HandlePath(handle string) string

	Call(method, action string, body io.Reader) (resp *http.Response, err error)
}
