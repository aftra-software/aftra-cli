package openapi

import (
	"context"
	"net/http"
)

func DoCreateOpportunity(ctx context.Context, client *ClientWithResponses, companyPk string, opportunity CreateOpportunity) (*http.Response, error) {

	params := CreateOpportunityApiCompaniesCompanyPkOpportunitiesPostParams{}
	return client.CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost(
		ctx,
		companyPk,
		&params,
		CreateOpportunityApiCompaniesCompanyPkOpportunitiesPostJSONRequestBody(opportunity),
	)
}
