package cmd

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteCreate(t *testing.T) {
	type test struct {
		serverResponse int
		expectedOutput string
	}

	tests := []test{
		{serverResponse: 200, expectedOutput: "unique101 created"},
		{serverResponse: 401, expectedOutput: "Unauthorized"},
		{serverResponse: 403, expectedOutput: "Forbidden"},
		{serverResponse: 490, expectedOutput: "Unrecognized status code 490"},
		{serverResponse: 503, expectedOutput: "Server Error: 503"},
	}

	for _, tc := range tests {
		mockDoer := &MockHTTP{
			Response:      http.Response{StatusCode: tc.serverResponse, Status: "Unauthorized"},
			ResponseError: nil,
		}
		actual := new(bytes.Buffer)
		rootCmd.SetOut(actual)
		rootCmd.SetErr(actual)
		rootCmd.SetArgs([]string{"create", "--uid=unique101", "--name=foo", "--score=medium", "--details=a=1,b=2"})

		ctx := context.WithValue(context.Background(), "doer", mockDoer) // TODO: move to use special type
		createCmd.SetContext(ctx)

		err := rootCmd.ExecuteContext(ctx)

		assert.Equal(t, nil, err)
		assert.Equal(t, len(mockDoer.Requests), 1)
		assert.Equal(t, tc.expectedOutput, actual.String())
	}

}
