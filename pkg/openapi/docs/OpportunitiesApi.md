# \OpportunitiesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost**](OpportunitiesApi.md#CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost) | **Post** /api/companies/{company_pk}/opportunities/ | Create Opportunity
[**GetOpportunityDetailsObject**](OpportunitiesApi.md#GetOpportunityDetailsObject) | **Get** /api/companies/{company_pk}/opportunities/{opportunity_uid}/details/ | Get Opportunity Details
[**ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet**](OpportunitiesApi.md#ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet) | **Get** /api/companies/{company_pk}/opportunities/ | List Opportunities
[**PostUpdateOpportunityResolution**](OpportunitiesApi.md#PostUpdateOpportunityResolution) | **Post** /api/companies/{company_pk}/opportunities/{opportunity_uid}/ | Post Update Resolution



## CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost

> interface{} CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost(ctx, companyPk).CreateOpportunity(createOpportunity).ParentPk(parentPk).Identity(identity).Execute()

Create Opportunity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    createOpportunity := *openapiclient.NewCreateOpportunity("Uid_example", "Name_example", map[string]string{"key": "Inner_example"}, openapiclient.OpportunityScore("none")) // CreateOpportunity | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpportunitiesApi.CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost(context.Background(), companyPk).CreateOpportunity(createOpportunity).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesApi.CreateOpportunityApiCompaniesCompanyPkOpportunitiesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOpportunityApiCompaniesCompanyPkOpportunitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOpportunity** | [**CreateOpportunity**](CreateOpportunity.md) |  | 
 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpportunityDetailsObject

> OpportunityDetails GetOpportunityDetailsObject(ctx, companyPk, opportunityUid).ParentPk(parentPk).Identity(identity).Execute()

Get Opportunity Details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    opportunityUid := "opportunityUid_example" // string | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpportunitiesApi.GetOpportunityDetailsObject(context.Background(), companyPk, opportunityUid).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.GetOpportunityDetailsObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOpportunityDetailsObject`: OpportunityDetails
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesApi.GetOpportunityDetailsObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**opportunityUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpportunityDetailsObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**OpportunityDetails**](OpportunityDetails.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet

> []Opportunity ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet(ctx, companyPk).Resolution(resolution).ParentPk(parentPk).Identity(identity).Execute()

List Opportunities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    resolution := []openapiclient.OpportunityResolution{openapiclient.OpportunityResolution("unacknowledged")} // []OpportunityResolution |  (optional)
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpportunitiesApi.ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet(context.Background(), companyPk).Resolution(resolution).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet`: []Opportunity
    fmt.Fprintf(os.Stdout, "Response from `OpportunitiesApi.ListOpportunitiesApiCompaniesCompanyPkOpportunitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOpportunitiesApiCompaniesCompanyPkOpportunitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolution** | [**[]OpportunityResolution**](OpportunityResolution.md) |  | 
 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**[]Opportunity**](Opportunity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUpdateOpportunityResolution

> PostUpdateOpportunityResolution(ctx, companyPk, opportunityUid).ResolutionUpdate(resolutionUpdate).ParentPk(parentPk).Identity(identity).Execute()

Post Update Resolution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    opportunityUid := "opportunityUid_example" // string | 
    resolutionUpdate := *openapiclient.NewResolutionUpdate("Resolution_example") // ResolutionUpdate | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpportunitiesApi.PostUpdateOpportunityResolution(context.Background(), companyPk, opportunityUid).ResolutionUpdate(resolutionUpdate).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesApi.PostUpdateOpportunityResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**opportunityUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUpdateOpportunityResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolutionUpdate** | [**ResolutionUpdate**](ResolutionUpdate.md) |  | 
 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

