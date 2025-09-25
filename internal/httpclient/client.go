package httpclient

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/TommyFiga/greq/internal/types"
)

func SendRequest(url string, method string, headers map[string][]string, body string) (*types.ResponseData, error) {
	if method == "" {
		return SendRequest(url, "GET", nil, "")
	}

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, vals := range headers {
		for _, v := range vals {
			req.Header.Set(key, v)
		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return getResponseData(resp)
}

func getResponseData(resp *http.Response) (*types.ResponseData, error) {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	headers := make(map[string][]string)
	for key, val := range resp.Header {
		headers[key] = append([]string(nil), val...)
	}

	return &types.ResponseData{
		Protocol: resp.Proto,
		Status:   resp.Status,
		Headers:  headers,
		Body:     bodyBytes,
	}, nil
}
