package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syndis-software/iris-api/pkg/openapi"
)

func Test_ExecuteLog_Single(t *testing.T) {
	type test struct {
		serverResponse int
		expectedOutput string
		errorExpected  bool
	}

	tests := map[string]test{
		"success": {serverResponse: 200, expectedOutput: "", errorExpected: false},
		"401":     {serverResponse: 401, expectedOutput: "Error: unauthorized\n", errorExpected: true},
		"403":     {serverResponse: 403, expectedOutput: "Error: forbidden\n", errorExpected: true},
		"500":     {serverResponse: 500, expectedOutput: "Error: server error: 500\n", errorExpected: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mockDoer := &MockHTTP{
				Response: http.Response{
					StatusCode: tc.serverResponse,
					Status:     "",
				},
				ResponseError: nil,
			}
			actual := new(bytes.Buffer)
			rootCmd.SetOut(actual)
			rootCmd.SetErr(actual)
			rootCmd.SetArgs([]string{"log", "My log message"})

			ctx := context.WithValue(context.Background(), doerKey, mockDoer)
			logCmd.SetContext(ctx)

			err := rootCmd.ExecuteContext(ctx)

			if tc.errorExpected {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, nil, err)
			}

			assert.Equal(t, len(mockDoer.Requests), 1)

			body, _ := ioutil.ReadAll(mockDoer.Requests[0].Body)
			var submitted []openapi.SubmitLogEvent
			_ = json.Unmarshal(body, &submitted)

			assert.Equal(t, len(submitted), 1)
			assert.Equal(t, "My log message", submitted[0].Message)
			assert.Equal(t, tc.expectedOutput, actual.String())
		})
	}

}

func Test_ExecuteLog_Stdin(t *testing.T) {
	stdinInput := strings.NewReader("abcde\nfoobar")

	mockDoer := &MockHTTP{
		Response: http.Response{
			StatusCode: 200,
			Status:     "",
		},
		ResponseError: nil,
	}
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"log"})

	ctx := context.WithValue(context.Background(), doerKey, mockDoer)
	ctx = context.WithValue(ctx, stdInKey, stdinInput)
	logCmd.SetContext(ctx)

	err := rootCmd.ExecuteContext(ctx)

	assert.Equal(t, nil, err)
	assert.Equal(t, len(mockDoer.Requests), 1)

	body, _ := ioutil.ReadAll(mockDoer.Requests[0].Body)
	var submitted []openapi.SubmitLogEvent
	_ = json.Unmarshal(body, &submitted)

	assert.Equal(t, len(submitted), 2)
	assert.Equal(t, "abcde", submitted[0].Message)
	assert.Equal(t, "foobar", submitted[1].Message)

}
