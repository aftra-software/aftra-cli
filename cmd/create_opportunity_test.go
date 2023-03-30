package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

func Test_ExecuteCreateOpportunity(t *testing.T) {
	type test struct {
		serverResponse int
		expectedOutput string
		errorExpected  bool
	}

	tests := []test{
		{serverResponse: 200, expectedOutput: "unique101 created\n", errorExpected: false},
		{serverResponse: 401, expectedOutput: "Error: unauthorized\n", errorExpected: true},
		{serverResponse: 403, expectedOutput: "Error: forbidden\n", errorExpected: true},
		{serverResponse: 490, expectedOutput: "Error: unrecognized status code 490\n", errorExpected: true},
		{serverResponse: 503, expectedOutput: "Error: server error: 503\n", errorExpected: true},
	}

	for _, tc := range tests {
		mockDoer := &MockHTTP{
			Response: http.Response{
				StatusCode: tc.serverResponse,
				Status:     "",
				Body:       ioutil.NopCloser(bytes.NewBufferString("null")),
			},
			ResponseError: nil,
		}
		actual := new(bytes.Buffer)
		rootCmd.SetOut(actual)
		rootCmd.SetErr(actual)
		rootCmd.SetArgs([]string{"create", "opportunity", "--uid=unique101", "--name=foo", "--score=medium", "--details=a=1,b=2"})

		ctx := context.WithValue(context.Background(), doerKey, mockDoer)
		opportunityCmd.SetContext(ctx)

		err := rootCmd.ExecuteContext(ctx)

		if tc.errorExpected {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, nil, err)
		}
		assert.Equal(t, len(mockDoer.Requests), 1)
		assert.Equal(t, tc.expectedOutput, actual.String())
	}

}

func Test_ExecuteCreateOpportunityDetails(t *testing.T) {
	type test struct {
		details                  string
		expectedDetailsSubmitted map[string]string
	}

	tests := map[string]test{
		"notgiven":  {details: "", expectedDetailsSubmitted: map[string]string{}},
		"empty":     {details: "--details=", expectedDetailsSubmitted: map[string]string{}},
		"single":    {details: "--details=a=1", expectedDetailsSubmitted: map[string]string{"a": "1"}},
		"double":    {details: "--details=a=1,b=2", expectedDetailsSubmitted: map[string]string{"a": "1", "b": "2"}},
		"duplicate": {details: "--details=a=1,a=2", expectedDetailsSubmitted: map[string]string{"a": "2"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			detailsStr = ""
			mockDoer := &MockHTTP{
				Response: http.Response{
					StatusCode: 200,
					Status:     "",
					Body:       ioutil.NopCloser(bytes.NewBufferString("null")),
				},
				ResponseError: nil,
			}
			actual := new(bytes.Buffer)
			rootCmd.SetOut(actual)
			rootCmd.SetErr(actual)
			rootCmd.SetArgs([]string{"create", "opportunity", "--uid=unique101", "--name=foo", "--score=medium", tc.details})

			ctx := context.WithValue(context.Background(), doerKey, mockDoer)
			opportunityCmd.SetContext(ctx)

			err := rootCmd.ExecuteContext(ctx)

			assert.Equal(t, nil, err)
			assert.Equal(t, len(mockDoer.Requests), 1)

			body, _ := ioutil.ReadAll(mockDoer.Requests[0].Body)
			var submitted openapi.CreateOpportunity
			_ = json.Unmarshal(body, &submitted)

			assert.Equal(t, tc.expectedDetailsSubmitted, submitted.Details)
		})
	}

}
