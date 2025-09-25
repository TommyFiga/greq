package main

import (
	"fmt"
	"os"

	"github.com/TommyFiga/greq/internal/httpclient"
	"github.com/TommyFiga/greq/internal/output"
	"github.com/TommyFiga/greq/internal/parser"
	"github.com/TommyFiga/greq/internal/printer"
)

func main() {
	opts, err := parser.ParseArgs()
	if err != nil {
		errorExit(fmt.Errorf("parser: %w", err))
	}

	respData, err := httpclient.SendRequest(opts.URL, opts.Method, opts.Headers, opts.Body)
	if err != nil {
		errorExit(fmt.Errorf("client: %w", err))
	}

	fmtResp := printer.FormatResponse(respData, opts.IncludeHeaders, opts.PrettyJSON)

	if len(opts.OutputFile) > 0 {
		err := output.WriteResponseContentToFile(fmtResp, opts.OutputFile)
		if err != nil {
			errorExit(fmt.Errorf("output file: %w", err))
		}
	} else {
		fmt.Println(fmtResp)
	}
}

func errorExit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
