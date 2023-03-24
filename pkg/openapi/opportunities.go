package openapi

import (
	"context"
	"encoding/json"
	"fmt"
)

func DoCreateOpportunity(ctx context.Context, client *ClientWithResponses, companyPk string, opportunity CreateOpportunity) error {

	params := CreateOpportunityApiCompaniesCompanyPkOpportunitiesPostParams{}
	resp, err := client.CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost(
		ctx,
		companyPk,
		&params,
		CreateOpportunityApiCompaniesCompanyPkOpportunitiesPostJSONRequestBody(opportunity),
	)
	if err != nil {
		return err
	}

	obj, err := ParseCreateOpportunityApiCompaniesCompanyPkOpportunitiesPostResponse(resp)

	if err != nil {
		return err
	}

	if obj.JSON422 != nil {
		repr, err := json.MarshalIndent(*obj.JSON422.Detail, "", "\t")

		if err != nil {
			return err
		}
		return fmt.Errorf("validation error: %s", repr)
	}

	return CheckStatus(resp)
}
