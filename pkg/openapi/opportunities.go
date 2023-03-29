package openapi

import (
	"context"
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

	return CheckStatus(resp)
}
