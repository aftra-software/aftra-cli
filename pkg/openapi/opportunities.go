package openapi

import (
	"context"
	"time"
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

func DoGetOpportunities(ctx context.Context, client *ClientWithResponses, companyPk string, lastUpdatedGte *time.Time, limit *int) (*SearchedOpportunitiesResponse, error) {

	var order SearchOpportunitiesApiCompaniesCompanyPkOpportunitiesV3GetParamsOrder = "asc"
	var sort SortOptions = "timestamp_last_updated"

	params := SearchOpportunitiesApiCompaniesCompanyPkOpportunitiesV3GetParams{
		TimestampLastUpdatedGte: lastUpdatedGte,
		Sort:                    &sort,
		Order:                   &order,
		Limit:                   limit,
	}
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
