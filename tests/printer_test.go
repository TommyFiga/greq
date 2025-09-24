package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/TommyFiga/greq/internal/printer"
)

func TestFormatResponse_BodyOnly(t *testing.T) {
	resp := &http.Response{
		Proto:      "HTTP/1.1",
		StatusCode: 200,
		Status:     "200 OK",
		Header:     http.Header{},
		Body:       io.NopCloser(strings.NewReader("Hello World")),
	}

	output := printer.FormatResponse(resp, false, false)

	if !strings.Contains(output, "Hello World") {
		t.Errorf("Expected body in output, got: %s", output)
	}
}

func TestFormatResponse_WithHeaders(t *testing.T) {
	resp := &http.Response{
		Proto:      "HTTP/1.1",
		StatusCode: 200,
		Status:     "200 OK",
		Header:     http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(strings.NewReader(`{"msg":"ok"}`)),
	}

	output := printer.FormatResponse(resp, true, false)

	if !strings.Contains(output, "HTTP/1.1 200 OK") {
		t.Errorf("Expected status line in output, got: %s", output)
	}

	if !strings.Contains(output, "Content-Type: application/json") {
		t.Errorf("Expected header in output, got: %s", output)
	}
}

func TestFormatResponse_PrettyJSON(t *testing.T) {
	resp := &http.Response{
		Proto:      "HTTP/1.1",
		StatusCode: 200,
		Status:     "200 OK",
		Header:     http.Header{},
		Body:       io.NopCloser(strings.NewReader(`{"name":"Alice","age":30}`)),
	}

	output := printer.FormatResponse(resp, false, true)

	if !strings.Contains(output, "\"name\": \"Alice\"") {
		t.Errorf("Expected pretty-printed JSON, got: %s", output)
	}
}
