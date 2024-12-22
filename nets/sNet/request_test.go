package sNet

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
)

type mockClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}

func TestRequest_Call(t *testing.T) {
	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message": "success"}`)),
	}

	client := &mockClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			return mockResp, nil
		},
	}

	service := &Service{
		url:    "http://example.com",
		secure: false,
		port:   80,
	}

	req := NewRequest().MethodGet().SetEndPoint("/test")

	data, err := req.Call(context.Background(), service, client)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{"message": "success"}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}
