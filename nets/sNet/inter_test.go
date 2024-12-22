package sNet

import (
	"testing"
)

type MockService struct{}

func (m *MockService) String() string {
	return "MockService"
}

func (m *MockService) Log() {}

func (m *MockService) SetDebug(bool) {}

func (m *MockService) Debug() bool {
	return false
}

func (m *MockService) Port() int {
	return 8080
}

func (m *MockService) GetUri() string {
	return "http://mockservice"
}

func (m *MockService) GetUrl() string {
	return "http://mockservice"
}

func (m *MockService) LocalAddr() string {
	return "localhost:8080"
}

func (m *MockService) Secret() string {
	return "secret"
}

func (m *MockService) HandlePath(path string) string {
	return "/mockservice" + path
}

func (m *MockService) SendObj(end_point string, obj interface{}) error {
	return nil
}

type MockRequest struct{}

func (m *MockRequest) String() string {
	return "MockRequest"
}

func (m *MockRequest) MethodGet() InterRequest {
	return m
}

func (m *MockRequest) MethodPost() InterRequest {
	return m
}

func (m *MockRequest) SetEndPoint(string) InterRequest {
	return m
}

func (m *MockRequest) SetParam(string, string) {}

func (m *MockRequest) AddParam(string, string) {}

func (m *MockRequest) DelParam(string) {}

func (m *MockRequest) SetHeader(string, string) {}

func (m *MockRequest) AddHeader(string, string) {}

func (m *MockRequest) DelHeader(string) {}

func (m *MockRequest) SetBody(io.Reader) {}

func (m *MockRequest) Call(context.Context, InterService, InterClient) ([]byte, error) {
	return []byte("mock response"), nil
}

type MockClient struct{}

func (m *MockClient) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("mock response")),
	}, nil
}

func TestMockService(t *testing.T) {
	service := &MockService{}

	if service.String() != "MockService" {
		t.Errorf("Expected MockService, got %s", service.String())
	}

	if service.Port() != 8080 {
		t.Errorf("Expected 8080, got %d", service.Port())
	}

	if service.GetUri() != "http://mockservice" {
		t.Errorf("Expected http://mockservice, got %s", service.GetUri())
	}

	if service.GetUrl() != "http://mockservice" {
		t.Errorf("Expected http://mockservice, got %s", service.GetUrl())
	}

	if service.LocalAddr() != "localhost:8080" {
		t.Errorf("Expected localhost:8080, got %s", service.LocalAddr())
	}

	if service.Secret() != "secret" {
		t.Errorf("Expected secret, got %s", service.Secret())
	}

	if service.HandlePath("/test") != "/mockservice/test" {
		t.Errorf("Expected /mockservice/test, got %s", service.HandlePath("/test"))
	}
}

func TestMockRequest(t *testing.T) {
	request := &MockRequest{}

	if request.String() != "MockRequest" {
		t.Errorf("Expected MockRequest, got %s", request.String())
	}

	response, err := request.Call(context.Background(), &MockService{}, &MockClient{})
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if string(response) != "mock response" {
		t.Errorf("Expected mock response, got %s", string(response))
	}
}
