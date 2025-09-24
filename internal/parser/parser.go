package parser

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	URL            string
	Method         string
	Headers        map[string]string
	Body           string
	IncludeHeaders bool
	PrettyJSON     bool
	OutputFile     string
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *arrayFlags) Set(val string) error {
	*i = append(*i, val)
	return nil
}

func ParseArgs() (*Options, error) {
	var methodFlag string
	var rawHeaders arrayFlags
	var bodyFlag string
	var includeHeaders bool
	var prettyJSON bool
	var outputFile string

	flag.StringVar(&methodFlag, "X", "GET", "HTTP method to use")
	flag.Var(&rawHeaders, "H", "Custom header to include in the request")
	flag.StringVar(&bodyFlag, "d", "", "Data to include in the request body")
	flag.BoolVar(&includeHeaders, "i", false, "Include response headers in the output")
	flag.BoolVar(&prettyJSON, "json", false, "Pretty-print JSON responses")
	flag.StringVar(&outputFile, "o", "", "Output response body to a file")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: greq [options] <url>")
		os.Exit(1)
	}
	url := flag.Arg(0)

	headers := make(map[string]string)
	for _, h := range rawHeaders {
		splitHeaders := strings.Split(h, ":")

		if len(splitHeaders) != 2 {
			fmt.Fprintf(os.Stderr, "WARNING: invalid header format, skipping: %s", h)
			continue
		}

		key := strings.TrimSpace(splitHeaders[0])
		val := strings.TrimSpace(splitHeaders[1])
		headers[key] = val
	}

	return &Options{
		URL:            url,
		Method:         methodFlag,
		Headers:        headers,
		Body:           bodyFlag,
		IncludeHeaders: includeHeaders,
		PrettyJSON:     prettyJSON,
		OutputFile:     outputFile,
	}, nil
}
