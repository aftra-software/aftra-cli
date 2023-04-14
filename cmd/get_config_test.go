package cmd

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteGetTokenConfig(t *testing.T) {
	type test struct {
		serverResponse     int
		serverResponseData string
		expectedOutput     string
	}

	tests := []test{
		{
			serverResponse:     200,
			serverResponseData: `{"ranges":"0/1","type":"INTERNAL"}`,
			expectedOutput:     `{"ranges":"0/1","type":"INTERNAL"}`,
		},
	}

	for _, tc := range tests {
		header := make(http.Header, 1)
		header.Set("Content-Type", "application/json")

		mockDoer := &MockHTTP{
			Response: http.Response{
				StatusCode: tc.serverResponse,
				Status:     "",
				Body:       ioutil.NopCloser(bytes.NewBufferString(tc.serverResponseData)),
				Header:     header,
			},
			ResponseError: nil,
		}
		actual := new(bytes.Buffer)
		rootCmd.SetOut(actual)
		rootCmd.SetErr(actual)
		rootCmd.SetArgs([]string{"get", "config", "syndis", "scan-name"})

		ctx := context.WithValue(context.Background(), doerKey, mockDoer)
		getConfigCmd.SetContext(ctx)

		err := rootCmd.ExecuteContext(ctx)

		assert.Equal(t, nil, err)
		assert.Equal(t, len(mockDoer.Requests), 1)
		assert.Equal(t, tc.expectedOutput, actual.String())
	}

}
