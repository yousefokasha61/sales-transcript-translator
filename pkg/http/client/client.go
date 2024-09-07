package client

import (
	"fmt"
	"io"
	"net/http"
)

const (
	contentTypeHeader         = "Content-Type"
	HTTPHeaderContentTypeJSON = "application/json; charset=utf-8"
)

type HTTPClient struct {
	httpClient *http.Client
}

func NewHTTPClient() *HTTPClient {
	httpClient := &http.Client{}
	return &HTTPClient{
		httpClient: httpClient,
	}
}

func (c *HTTPClient) Post(url, contentType string, body io.Reader, httpHeaders map[string]string) (*http.Response, error) {
	op := "HTTPClient.Post"
	httpReq, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	httpReq.Header.Set(contentTypeHeader, contentType)
	for k, v := range httpHeaders {
		httpReq.Header.Set(k, v)
	}

	return c.httpClient.Do(httpReq)
}
