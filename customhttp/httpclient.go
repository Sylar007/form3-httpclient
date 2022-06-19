package customhttp

import (
	"net/http"
	"sync"
)

type httpClient struct {
	builder    *httpBuilder
	client     *http.Client
	clientOnce sync.Once
}

type CustomClient interface {
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, body interface{}, headers http.Header) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}

func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, body interface{}, headers http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
