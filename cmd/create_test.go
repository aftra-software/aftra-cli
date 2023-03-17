package cmd

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteCreate(t *testing.T) {

	mockDoer := MockHTTP{
		Response:      http.Response{StatusCode: 200, Status: "OK"},
		ResponseError: nil,
	}

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"create", "--uid=unique101", "--name=foo", "--score=medium", "--details=a=1,b=2"})

	ctx := context.WithValue(context.Background(), "doer", mockDoer) // TODO: move to use special type
	err := rootCmd.ExecuteContext(ctx)

	assert.Equal(t, nil, err)
	assert.Equal(t, "unique101 created", actual.String())
}

func Test_ExecuteCreate_401(t *testing.T) {

	mockDoer := MockHTTP{
		Response:      http.Response{StatusCode: 401, Status: "Unauthorized"},
		ResponseError: nil,
	}

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"create", "--uid=unique101", "--name=foo", "--score=medium", "--details=a=1,b=2"})

	ctx := context.WithValue(context.Background(), "doer", mockDoer) // TODO: move to use special type
	err := rootCmd.ExecuteContext(ctx)

	assert.Equal(t, nil, err)
	assert.Equal(t, "Unauthorized", actual.String())
}
