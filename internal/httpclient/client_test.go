package httpclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSendRequest_BasicGET(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}))
	defer ts.Close()

	resp, err := SendRequest(ts.URL, "GET", nil, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Status != "200 OK" {
		t.Errorf("got %s, want 200 OK", resp.Status)
	}
	if string(resp.Body) != "hello world" {
		t.Errorf("got body %q, want %q", resp.Body, "hello world")
	}
}

func TestSendRequest_CustomMethodAndBody(t *testing.T) {
	expectedBody := "foo=bar"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != expectedBody {
			t.Errorf("got body %q, want %q", body, expectedBody)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	resp, err := SendRequest(ts.URL, "POST", nil, expectedBody)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Status != "201 Created" {
		t.Errorf("got %s, want 201 Created", resp.Status)
	}
}

func TestSendRequest_CustomHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test") != "value" {
			t.Errorf("expected header X-Test=value, got %s", r.Header.Get("X-Test"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	headers := map[string][]string{"X-Test" : {"value"}}
	_, err := SendRequest(ts.URL, "GET", headers, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSendRequest_MultipleResponseHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "a=1")
		w.Header().Add("Set-Cookie", "b=2")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	resp, err := SendRequest(ts.URL, "GET", nil, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cookies := resp.Headers["Set-Cookie"]
	if len(cookies) != 2 {
		t.Errorf("expected 2 cookies, got %d", len(cookies))
	}
}

func TestSendRequest_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(12 * time.Second) // longer than client timeout
	}))
	defer ts.Close()

	clientTimeout := 10 * time.Second
	oldTimeout := clientTimeout
	defer func() { clientTimeout = oldTimeout }()

	_, err := SendRequest(ts.URL, "GET", nil, "")
	if err == nil {
		t.Fatalf("expected timeout error, got nil")
	}
}

func TestSendRequest_GETWithQueryParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "foo=bar&baz=qux" {
			t.Errorf("unexpected query: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	_, err := SendRequest(ts.URL+"?foo=bar&baz=qux", "GET", nil, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSendRequest_POSTWithMultipleHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test1") != "val1" || r.Header.Get("X-Test2") != "val2" {
			t.Errorf("headers missing or incorrect: %v", r.Header)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	headers := map[string][]string{
		"X-Test1" : {"val1"},
		"X-Test2" : {"val2"},
	}

	_, err := SendRequest(ts.URL, "POST", headers, "payload")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSendRequest_ResponseMultipleValuesSameHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Set-Cookie", "a=1")
		w.Header().Add("Set-Cookie", "b=2")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	resp, err := SendRequest(ts.URL, "GET", nil, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	values := resp.Headers["Set-Cookie"]
	if len(values) != 2 || values[0] != "a=1" || values[1] != "b=2" {
		t.Errorf("expected 2 Set-Cookie values [a=1, b=2], got %v", values)
	}
}

func TestSendRequest_InvalidURL(t *testing.T) {
	_, err := SendRequest("http://\x7f", "GET", nil, "")
	if err == nil {
		t.Fatalf("expected error for invalid URL, got nil")
	}
}

func TestSendRequest_POSTEmptyBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		if len(body) != 0 {
			t.Errorf("expected empty body, got %q", body)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	_, err := SendRequest(ts.URL, "POST", nil, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}