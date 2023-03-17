package cmd

import (
	"fmt"
	"net/http"
)

type MockHTTP struct {
	Response      http.Response
	ResponseError error
}

func (c MockHTTP) Do(req *http.Request) (*http.Response, error) {
	fmt.Printf("Returning response %v", c.Response)
	return &c.Response, c.ResponseError
}
