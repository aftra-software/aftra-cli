# \CompaniesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompany**](CompaniesApi.md#CreateCompany) | **Post** /api/companies/ | Create Company
[**GetCompany**](CompaniesApi.md#GetCompany) | **Get** /api/companies/{company_pk} | Get Company
[**GetCompanyScore**](CompaniesApi.md#GetCompanyScore) | **Get** /api/companies/{company_pk}/score/{back_date} | Get Score
[**GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet**](CompaniesApi.md#GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet) | **Get** /api/companies/{company_pk}/services-summary | Get Saas Services
[**ListCompanies**](CompaniesApi.md#ListCompanies) | **Get** /api/companies/ | List Companies
[**ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost**](CompaniesApi.md#ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost) | **Post** /api/companies/{company_pk}/scan/{asset_type} | Scan Company Asset
[**UpdateCompany**](CompaniesApi.md#UpdateCompany) | **Patch** /api/companies/{company_pk}/ | Update Company



## CreateCompany

> CompanyEntity CreateCompany(ctx).Company(company).Identity(identity).Execute()

Create Company

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
    company := *openapiclient.NewCompany("Name_example", "MainContact_example", "Website_example") // Company | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.CreateCompany(context.Background()).Company(company).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.CreateCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompany`: CompanyEntity
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.CreateCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | [**Company**](Company.md) |  | 
 **identity** | **string** |  | 

### Return type

[**CompanyEntity**](CompanyEntity.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> CompanyEntity GetCompany(ctx, companyPk).Identity(identity).Execute()

Get Company

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
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.GetCompany(context.Background(), companyPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: CompanyEntity
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identity** | **string** |  | 

### Return type

[**CompanyEntity**](CompanyEntity.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyScore

> HistoricalScores GetCompanyScore(ctx, companyPk, backDate).Identity(identity).Execute()

Get Score

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
    backDate := openapiclient.BackDates("1w") // BackDates | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.GetCompanyScore(context.Background(), companyPk, backDate).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetCompanyScore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyScore`: HistoricalScores
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetCompanyScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**backDate** | [**BackDates**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**HistoricalScores**](HistoricalScores.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet

> AggregatedSaasEnumerationReport GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Get Saas Services

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
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet`: AggregatedSaasEnumerationReport
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetSaasServicesApiCompaniesCompanyPkServicesSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSaasServicesApiCompaniesCompanyPkServicesSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**AggregatedSaasEnumerationReport**](AggregatedSaasEnumerationReport.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanies

> []CompanyWithUserAccess ListCompanies(ctx).Identity(identity).Execute()

List Companies

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
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.ListCompanies(context.Background()).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanies`: []CompanyWithUserAccess
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identity** | **string** |  | 

### Return type

[**[]CompanyWithUserAccess**](CompanyWithUserAccess.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost

> CompanyEntity ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost(ctx, companyPk, assetType).Identity(identity).Execute()

Scan Company Asset

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
    assetType := "assetType_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost(context.Background(), companyPk, assetType).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost`: CompanyEntity
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**assetType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanCompanyAssetApiCompaniesCompanyPkScanAssetTypePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**CompanyEntity**](CompanyEntity.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompany

> CompanyEntity UpdateCompany(ctx, companyPk).CompanyUpdate(companyUpdate).ParentPk(parentPk).Identity(identity).Execute()

Update Company

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
    companyUpdate := *openapiclient.NewCompanyUpdate() // CompanyUpdate | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.UpdateCompany(context.Background(), companyPk).CompanyUpdate(companyUpdate).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.UpdateCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompany`: CompanyEntity
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.UpdateCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyUpdate** | [**CompanyUpdate**](CompanyUpdate.md) |  | 
 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**CompanyEntity**](CompanyEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

