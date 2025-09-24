package main

import (
	"fmt"
	"os"

	"github.com/TommyFiga/greq/internal/httpclient"
	"github.com/TommyFiga/greq/internal/parser"
	"github.com/TommyFiga/greq/internal/printer"
)


func main() {
	opts,err := parser.ParseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	resp, err := httpclient.SendRequest(opts.URL, opts.Method, opts.Headers, opts.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	res := printer.FormatResponse(resp, opts.IncludeHeaders, opts.PrettyJSON)

	fmt.Println(res)
}