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

func DoGetOpportunities(ctx context.Context, client *ClientWithResponses, companyPk string, params SearchOpportunitiesApiCompaniesCompanyPkOpportunitiesV3GetParams) (*SearchedOpportunitiesResponse, error) {

	resp, err := client.SearchOpportunitiesApiCompaniesCompanyPkOpportunitiesV3Get(ctx, companyPk, &params)
	if err != nil {
		return nil, err
	}

	err = CheckStatus(resp)

	if err != nil {
		return nil, err
	}
	opportunities, err := ParseSearchOpportunitiesApiCompaniesCompanyPkOpportunitiesV3GetResponse(resp)

	if err != nil {
		return nil, err
	}
	return opportunities.JSON200, nil

}
