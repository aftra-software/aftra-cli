package cmd

import (
	"errors"
	"fmt"
	"net/http"
)

type MockHTTP struct {
	Requests  []http.Request
	Responses map[string]Response
}
type Response struct {
	Response      http.Response
	ResponseError error
}

func (c *MockHTTP) Do(req *http.Request) (*http.Response, error) {
	c.Requests = append(c.Requests, *req)
	fmt.Println("IN DO---------")
	fmt.Println(req.URL.Path)
	response, ok := c.Responses[req.URL.Path]
	if !ok {
		return nil, errors.New(fmt.Sprintf("URL %s not found in mapping", req.URL.Path))
	}
	return &response.Response, response.ResponseError
}

func (c *MockHTTP) CountRequests(path string) int {
	count := 0
	for _, req := range c.Requests {
		if req.URL.Path == path {
			count += 1
		}
	}
	return count
}
