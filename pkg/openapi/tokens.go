package openapi

import (
	"context"
	"net/http"
)

func DoFetchToken(ctx context.Context, client *Client) (*http.Response, error) {
	return client.GetTokenInfo(ctx)

}

func DoSubmitLog(ctx context.Context, client *Client, logs []SubmitLogEvent) (*http.Response, error) {
	return client.SubmitLogsForToken(ctx, SubmitLogsForTokenJSONRequestBody(logs))
}
