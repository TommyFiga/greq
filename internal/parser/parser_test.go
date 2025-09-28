package parser

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}


func TestParseArgs_Defaults(t *testing.T) {
	resetFlags()
	os.Args = []string{"greq", "https://example.com"}

	opts, err := ParseArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if opts.URL != "https://example.com" {
		t.Errorf("got URL %q, want %q", opts.URL, "https://example.com")
	}
	if opts.Method != "GET" {
		t.Errorf("got method %q, want %q", opts.Method, "GET")
	}
}

func TestParseArgs_CustomHeaders_Single(t *testing.T) {
	resetFlags()
	os.Args = []string{"greq", "-H", "Content-Type: application/json", "https://api.test"}

	opts, err := ParseArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	values, ok := opts.Headers["Content-Type"]
	if !ok {
		t.Fatalf("expected header Content-Type to be set")
	}
	if len(values) != 1 || values[0] != "application/json" {
		t.Errorf("got %v, want [application/json]", values)
	}
}

func TestParseArgs_CustomHeaders_Multiple(t *testing.T) {
	resetFlags()
	os.Args = []string{
		"greq",
		"-H", "Accept: application/json",
		"-H", "Accept: text/html",
		"https://api.test",
	}

	opts, err := ParseArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	values, ok := opts.Headers["Accept"]
	fmt.Printf("%v\n", opts.Headers)
	if !ok {
		t.Fatalf("expected header Accept to be set")
	}
	if len(values) != 2 {
		t.Errorf("expected 2 values for Accept, got %v", values)
	}
}

func TestParseArgs_InvalidHeader(t *testing.T) {
	resetFlags()
	os.Args = []string{"greq", "-H", "InvalidHeader", "https://example.com"}

	// Should not crash, just skip invalid header
	opts, err := ParseArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(opts.Headers) != 0 {
		t.Errorf("expected no headers, got %v", opts.Headers)
	}
}

func TestParseArgs_Flags(t *testing.T) {
	resetFlags()
	os.Args = []string{
		"greq",
		"-X", "POST",
		"-d", `{"name":"greq"}`,
		"-i",
		"-json",
		"-o", "out.txt",
		"https://httpbin.org/post",
	}

	opts, err := ParseArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if opts.Method != "POST" {
		t.Errorf("got method %q, want POST", opts.Method)
	}
	if opts.Body != `{"name":"greq"}` {
		t.Errorf("got body %q, want %q", opts.Body, `{"name":"greq"}`)
	}
	if !opts.IncludeHeaders {
		t.Errorf("expected IncludeHeaders to be true")
	}
	if !opts.PrettyJSON {
		t.Errorf("expected PrettyJSON to be true")
	}
	if opts.OutputFile != "out.txt" {
		t.Errorf("got output file %q, want out.txt", opts.OutputFile)
	}
}