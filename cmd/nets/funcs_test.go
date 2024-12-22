package nets

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestPrueba(t *testing.T) {
	// Set up a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/a" {
			t.Errorf("Expected request to '/a', got '%s'", r.URL.EscapedPath())
		}
		body := new(strings.Builder)
		_, err := body.ReadFrom(r.Body)
		if err != nil {
			t.Errorf("Error reading request body: %v", err)
		}
		expectedBody := "aqqbPIycNs+ToAXQIX3TInl/ArkJcqbnhmZQpneqxL6iHg+nynYVzxusUuv5Z7qfRkSVmiIXeFSTRL6FcwDiyUJBMU1r3NxarRFjvMTlTGIsbpDwwjNyqYtMfUlkDX/iEmzpnZvxri0fJ2sipA8ouQ=="
		if body.String() != expectedBody {
			t.Errorf("Expected body '%s', got '%s'", expectedBody, body.String())
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer ts.Close()

	// Override the environment variable for testing
	oldEnv := os.Getenv("SERV_Own")
	defer os.Setenv("SERV_Own", oldEnv)
	os.Setenv("SERV_Own", ts.URL)

	Prueba()
}

func TestHandler1(t *testing.T) {
	req, err := http.NewRequest("GET", "/server1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler1)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler1 returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello from server 1!"
	if rr.Body.String() != expected {
		t.Errorf("handler1 returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandler2(t *testing.T) {
	req, err := http.NewRequest("GET", "/server2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler2)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler2 returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello from server 2!"
	if rr.Body.String() != expected {
		t.Errorf("handler2 returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestMain(t *testing.T) {
	// Create test servers
	server1 := httptest.NewServer(http.HandlerFunc(handler1))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(handler2))
	defer server2.Close()

	// Override the default servers with test servers
	go func() {
		if err := http.ListenAndServe(":8080", http.HandlerFunc(handler1)); err != nil {
			t.Errorf("Server 1 Error: %v", err)
		}
	}()

	go func() {
		if err := http.ListenAndServe(":8081", http.HandlerFunc(handler2)); err != nil {
			t.Errorf("Server 2 Error: %v", err)
		}
	}()

	// Allow some time for servers to start
	time.Sleep(1 * time.Second)

	// Test server 1
	resp1, err := http.Get(server1.URL + "/server1")
	if err != nil {
		t.Fatalf("Failed to send request to server 1: %v", err)
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200 from server 1, got %v", resp1.StatusCode)
	}

	// Test server 2
	resp2, err := http.Get(server2.URL + "/server2")
	if err != nil {
		t.Fatalf("Failed to send request to server 2: %v", err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200 from server 2, got %v", resp2.StatusCode)
	}
}
