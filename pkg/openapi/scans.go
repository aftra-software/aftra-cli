package openapi

import (
	"context"
	"encoding/json"
)

//  change this to be getconfig, switch on scan-type
func DoGetScanConfig(ctx context.Context, scanType string, configName string, client *ClientWithResponses) (string, error) {

	resp, err := client.GetSyndisConfigInfo(ctx, configName)

	if err != nil {
		return "", err
	}

	err = CheckStatus(resp)

	if err != nil {
		return "", err
	}

	config, err := ParseGetSyndisConfigInfoResponse(resp)

	if err != nil {
		return "", err
	}
	s, _ := json.Marshal(config.JSON200)
	return string(s), nil

}
