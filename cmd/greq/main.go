package main

import (
	"fmt"
	"os"

	"greq/internal/httpclient"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: greq <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	
	body, err := httpclient.SimpleGet(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(body)
}