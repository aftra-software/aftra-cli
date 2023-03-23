package openapi

import (
	"context"
	"net/http"
)

func DoFetchToken(ctx context.Context, client *ClientWithResponses) (*GetTokenInfoResponse, error) {
	return client.GetTokenInfoWithResponse(ctx)

}

func DoSubmitLog(ctx context.Context, client *ClientWithResponses, logs []SubmitLogEvent) (*http.Response, error) {
	return client.SubmitLogsForToken(ctx, SubmitLogsForTokenJSONRequestBody(logs))
}
