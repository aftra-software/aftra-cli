package cmd

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteGetOpportunities(t *testing.T) {
	type test struct {
		serverResponse     int
		serverResponseData string
		expectedOutput     string
	}

	tests := []test{
		{
			serverResponse:     200,
			serverResponseData: `{"opportunities": [{"a": 1}, {"b": 2}], "total": 10}`,
			expectedOutput:     "{\"a\":1}\n{\"b\":2}\n",
		},
	}

	for _, tc := range tests {
		header := make(http.Header, 1)
		header.Set("Content-Type", "application/json")

		m := make(map[string]Response)

		m["/api/companies//opportunities/v3"] = Response{
			Response: http.Response{
				StatusCode: tc.serverResponse,
				Status:     "200",
				Body:       ioutil.NopCloser(bytes.NewBufferString(tc.serverResponseData)),
				Header:     header,
			},
			ResponseError: nil,
		}
		mockDoer := &MockHTTP{Responses: m}
		actual := new(bytes.Buffer)
		rootCmd.SetOut(actual)
		rootCmd.SetErr(actual)
		rootCmd.SetArgs([]string{"get", "opportunities", "--updated-since", "2024-01-01T00:00:00Z", "--limit", "100"})

		ctx := context.WithValue(context.Background(), doerKey, mockDoer)
		getConfigCmd.SetContext(ctx)

		err := rootCmd.ExecuteContext(ctx)

		assert.Equal(t, nil, err)
		assert.Equal(t, len(mockDoer.Requests), 1)
		assert.Equal(t, tc.expectedOutput, actual.String())
	}

}

func Test_ValidationGetOpportunities(t *testing.T) {
	type test struct {
		updatedSinceArg       string
		limitArg              string
		expectedOutputPartial string
	}

	tests := []test{
		{
			updatedSinceArg:       "wrong-format",
			limitArg:              "100",
			expectedOutputPartial: "Error: invalid time format",
		},
		{
			updatedSinceArg:       "2024-01-01T00:00:00Z",
			limitArg:              "abc",
			expectedOutputPartial: "Error: invalid argument \"abc\" for \"--limit\" flag",
		},
	}

	for _, tc := range tests {

		mockDoer := &MockHTTP{}
		actual := new(bytes.Buffer)
		rootCmd.SetOut(actual)
		rootCmd.SetErr(actual)
		rootCmd.SetArgs([]string{"get", "opportunities", "--updated-since", tc.updatedSinceArg, "--limit", tc.limitArg})

		ctx := context.WithValue(context.Background(), doerKey, mockDoer)
		getConfigCmd.SetContext(ctx)

		rootCmd.ExecuteContext(ctx)

		assert.Contains(t, actual.String(), tc.expectedOutputPartial)
	}

}
