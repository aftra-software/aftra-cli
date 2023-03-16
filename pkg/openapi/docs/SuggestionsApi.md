# \SuggestionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost**](SuggestionsApi.md#AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) | **Post** /api/companies/{company_pk}/suggestions/{suggestion_id} | Accept Suggestion
[**GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet**](SuggestionsApi.md#GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet) | **Get** /api/companies/{company_pk}/suggestions/{suggestion_id} | Get Suggestion
[**ListSuggestionsApiCompaniesCompanyPkSuggestionsGet**](SuggestionsApi.md#ListSuggestionsApiCompaniesCompanyPkSuggestionsGet) | **Get** /api/companies/{company_pk}/suggestions/ | List Suggestions
[**RejectSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdDelete**](SuggestionsApi.md#RejectSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdDelete) | **Delete** /api/companies/{company_pk}/suggestions/{suggestion_id} | Reject Suggestion



## AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost

> ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost(ctx, companyPk, suggestionId).ParentPk(parentPk).Identity(identity).Execute()

Accept Suggestion

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
    suggestionId := "suggestionId_example" // string | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestionsApi.AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost(context.Background(), companyPk, suggestionId).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsApi.AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost`: ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost
    fmt.Fprintf(os.Stdout, "Response from `SuggestionsApi.AcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**suggestionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost**](ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet

> Suggestion GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet(ctx, companyPk, suggestionId).ParentPk(parentPk).Identity(identity).Execute()

Get Suggestion

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
    suggestionId := "suggestionId_example" // string | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestionsApi.GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet(context.Background(), companyPk, suggestionId).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsApi.GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet`: Suggestion
    fmt.Fprintf(os.Stdout, "Response from `SuggestionsApi.GetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**suggestionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**Suggestion**](Suggestion.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSuggestionsApiCompaniesCompanyPkSuggestionsGet

> []Suggestion ListSuggestionsApiCompaniesCompanyPkSuggestionsGet(ctx, companyPk).EntityType(entityType).ParentPk(parentPk).Identity(identity).Execute()

List Suggestions

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
    entityType := openapiclient.SuggestionEntityTypes("Domain") // SuggestionEntityTypes |  (optional)
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestionsApi.ListSuggestionsApiCompaniesCompanyPkSuggestionsGet(context.Background(), companyPk).EntityType(entityType).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsApi.ListSuggestionsApiCompaniesCompanyPkSuggestionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSuggestionsApiCompaniesCompanyPkSuggestionsGet`: []Suggestion
    fmt.Fprintf(os.Stdout, "Response from `SuggestionsApi.ListSuggestionsApiCompaniesCompanyPkSuggestionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSuggestionsApiCompaniesCompanyPkSuggestionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityType** | [**SuggestionEntityTypes**](SuggestionEntityTypes.md) |  | 
 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**[]Suggestion**](Suggestion.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdDelete

> RejectSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdDelete(ctx, companyPk, suggestionId).ParentPk(parentPk).Identity(identity).Execute()

Reject Suggestion

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
    suggestionId := "suggestionId_example" // string | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestionsApi.RejectSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdDelete(context.Background(), companyPk, suggestionId).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsApi.RejectSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**suggestionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

