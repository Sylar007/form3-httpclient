package customhttp

import (
	"net/http"
	"time"
)

type HttpBuilder interface {
	SetHeaders(headers http.Header) HttpBuilder
	SetConnectionTimeout(timeout time.Duration) HttpBuilder
	SetResponseTimeout(timeout time.Duration) HttpBuilder
	SetMaxIdleConnections(i int) HttpBuilder
	SetHttpClient(c *http.Client) HttpBuilder
	Build() CustomClient
}

type httpBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	client             *http.Client
}

func NewBuilder() HttpBuilder {
	builder := &httpBuilder{}
	return builder
}

func (c *httpBuilder) Build() CustomClient {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *httpBuilder) SetHeaders(headers http.Header) HttpBuilder {
	c.headers = headers
	return c
}

func (c *httpBuilder) SetConnectionTimeout(timeout time.Duration) HttpBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *httpBuilder) SetResponseTimeout(timeout time.Duration) HttpBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *httpBuilder) SetMaxIdleConnections(i int) HttpBuilder {
	c.maxIdleConnections = i
	return c
}

func (c *httpBuilder) SetHttpClient(client *http.Client) HttpBuilder {
	c.client = client
	return c
}
