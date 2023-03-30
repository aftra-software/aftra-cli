package cmd

import (
	"net/http"
)

type MockHTTP struct {
	Requests      []http.Request
	Response      http.Response
	ResponseError error
}

func (c *MockHTTP) Do(req *http.Request) (*http.Response, error) {
	c.Requests = append(c.Requests, *req)
	return &c.Response, c.ResponseError
}
