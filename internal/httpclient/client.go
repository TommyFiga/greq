package httpclient

import (
	"net/http"
	"strings"
	"time"
)

func SimpleGet(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SendRequest(url string, method string, headers map[string]string, body string) (*http.Response, error) {
	if method == "" {
		return SimpleGet(url)
	}

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
