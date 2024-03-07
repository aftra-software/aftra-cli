package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteUpdateResolution(t *testing.T) {
	type test struct {
		serverResponse int
		expectedOutput string
		uid            string
	}

	tests := []test{
		{
			uid:            "123",
			serverResponse: 200,
			expectedOutput: "",
		},
	}

	for _, tc := range tests {
		t.Run(name, func(to *testing.T) {
			header := make(http.Header, 1)
			header.Set("Content-Type", "application/json")

			m := make(map[string][]Response)
			mockDoer := &MockHTTP{Responses: m}
			mockDoer.AddResponse(fmt.Sprintf("/api/companies//opportunities/%s/", tc.uid), Response{
				Response: http.Response{
					StatusCode: tc.serverResponse,
					Status:     fmt.Sprint(tc.serverResponse),
					Body:       ioutil.NopCloser(bytes.NewBufferString("")),
					Header:     header,
				},
				ResponseError: nil,
			})
			actual := new(bytes.Buffer)
			rootCmd.SetOut(actual)
			rootCmd.SetErr(actual)
			rootCmd.SetArgs([]string{"update", "resolution", tc.uid, "RESOLVED", "--comment", "Worked all night"})

			ctx := context.WithValue(context.Background(), doerKey, mockDoer)
			updateResolutionsCmd.SetContext(ctx)

			err := rootCmd.ExecuteContext(ctx)

			assert.Equal(to, nil, err)
			assert.Equal(to, tc.expectedOutput, actual.String())
		})
	}

}
