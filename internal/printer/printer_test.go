package printer

import(
	"strings"
	"testing"

	"github.com/TommyFiga/greq/internal/types"
)

func TestFormatResponse_Basic(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "200 OK",
		Headers:  map[string][]string{},
		Body:     []byte("Hello World"),
	}

	out := FormatResponse(resp, false, false)
	if !strings.Contains(out, "HTTP/1.1 200 OK") {
		t.Errorf("expected status in output, got %s", out)
	}
	if !strings.Contains(out, "Hello World") {
		t.Errorf("expected body in output, got %s", out)
	}
}

func TestFormatResponse_WithHeaders(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/2",
		Status:   "404 Not Found",
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
			"Set-Cookie":   {"a=1", "b=2"},
		},
		Body: []byte("Not Found"),
	}

	out := FormatResponse(resp, true, false)

	if !strings.Contains(out, "HTTP/2 404 Not Found") {
		t.Errorf("status missing: %s", out)
	}
	if !strings.Contains(out, "Content-Type: application/json") {
		t.Errorf("header missing: %s", out)
	}
	if !strings.Contains(out, "Set-Cookie: a=1, b=2") {
		t.Errorf("multi-value header missing or wrong: %s", out)
	}
	if !strings.Contains(out, "Not Found") {
		t.Errorf("body missing: %s", out)
	}
}

func TestFormatResponse_PrettyJSON(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "200 OK",
		Headers:  map[string][]string{},
		Body:     []byte(`{"name":"greq","version":1}`),
	}

	out := FormatResponse(resp, false, true)

	// Should contain indentation/newlines
	if !strings.Contains(out, "\n  ") {
		t.Errorf("expected pretty JSON, got: %s", out)
	}
}

func TestFormatResponse_InvalidJSON(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "200 OK",
		Headers:  map[string][]string{},
		Body:     []byte(`{invalid-json}`),
	}

	out := FormatResponse(resp, false, true)

	// Should still include raw body
	if !strings.Contains(out, "{invalid-json}") {
		t.Errorf("expected raw body on invalid JSON, got %s", out)
	}
}

func TestFormatResponse_EmptyBody(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "204 No Content",
		Headers:  map[string][]string{},
		Body:     []byte{},
	}

	out := FormatResponse(resp, false, false)

	if !strings.Contains(out, "HTTP/1.1 204 No Content") {
		t.Errorf("status missing: %s", out)
	}
	if !strings.Contains(out, "Data:") {
		t.Errorf("expected Data header in output, got %s", out)
	}
}

func TestFormatResponse_MultiValueHeaders(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "200 OK",
		Headers: map[string][]string{
			"Set-Cookie": {"a=1", "b=2", "c=3"},
		},
		Body: []byte("OK"),
	}

	out := FormatResponse(resp, true, false)

	if !strings.Contains(out, "Set-Cookie: a=1, b=2, c=3") {
		t.Errorf("expected multi-value headers, got %s", out)
	}
}

func TestFormatResponse_PrettyJSON_Nested(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "200 OK",
		Headers:  map[string][]string{},
		Body:     []byte(`{"user":{"name":"Alice","age":30},"active":true}`),
	}

	out := FormatResponse(resp, false, true)

	if !strings.Contains(out, "\n  \"user\": {") || !strings.Contains(out, "\n  \"active\": true") {
		t.Errorf("expected pretty-printed nested JSON, got %s", out)
	}
}

func TestFormatResponse_EmptyHeaders_JSONFalse(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "200 OK",
		Headers:  map[string][]string{},
		Body:     []byte("Hello World"),
	}

	out := FormatResponse(resp, false, false)

	if !strings.Contains(out, "HTTP/1.1 200 OK") || !strings.Contains(out, "Hello World") {
		t.Errorf("unexpected output: %s", out)
	}
}

func TestFormatResponse_Headers_EmptyBody(t *testing.T) {
	resp := &types.ResponseData{
		Protocol: "HTTP/1.1",
		Status:   "204 No Content",
		Headers: map[string][]string{
			"X-Test": {"value"},
		},
		Body: []byte{},
	}

	out := FormatResponse(resp, true, false)

	if !strings.Contains(out, "X-Test: value") {
		t.Errorf("expected header in output, got %s", out)
	}
	if !strings.Contains(out, "204 No Content") {
		t.Errorf("expected status in output, got %s", out)
	}
}
