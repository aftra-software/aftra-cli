# \DomainsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDomainToCompany**](DomainsApi.md#AddDomainToCompany) | **Post** /api/companies/{parent_pk}/domains | Add Domain To Company
[**CheckDomainVerified**](DomainsApi.md#CheckDomainVerified) | **Get** /api/companies/{parent_pk}/domains/{entity_pk}/verified | Check Domain Verified
[**DomainScreenshot**](DomainsApi.md#DomainScreenshot) | **Get** /api/companies/domains/screenshot | Domain Screenshot
[**GetCompanyDomain**](DomainsApi.md#GetCompanyDomain) | **Get** /api/companies/{parent_pk}/domains/{entity_pk} | Get Company Domain
[**ListCompanyDomains**](DomainsApi.md#ListCompanyDomains) | **Get** /api/companies/{parent_pk}/domains | List Company Domains
[**RemoveCompanyDomain**](DomainsApi.md#RemoveCompanyDomain) | **Delete** /api/companies/{parent_pk}/domains/{entity_pk} | Remove Company Domain
[**UpdateCompanyDomain**](DomainsApi.md#UpdateCompanyDomain) | **Put** /api/companies/{parent_pk}/domains/{entity_pk} | Update Company Domain
[**UploadDomainListApiCompaniesParentPkDomainsUploadPost**](DomainsApi.md#UploadDomainListApiCompaniesParentPkDomainsUploadPost) | **Post** /api/companies/{parent_pk}/domains/upload | Upload Domain List
[**WhoisDomain**](DomainsApi.md#WhoisDomain) | **Get** /api/companies/domains/whois | Whois Domain



## AddDomainToCompany

> DomainEntity AddDomainToCompany(ctx, parentPk).Domain(domain).Identity(identity).Execute()

Add Domain To Company

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
    parentPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    domain := *openapiclient.NewDomain("Name_example") // Domain | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.AddDomainToCompany(context.Background(), parentPk).Domain(domain).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.AddDomainToCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDomainToCompany`: DomainEntity
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.AddDomainToCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDomainToCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | [**Domain**](Domain.md) |  | 
 **identity** | **string** |  | 

### Return type

[**DomainEntity**](DomainEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckDomainVerified

> VerificationResult CheckDomainVerified(ctx, parentPk, entityPk).Identity(identity).Execute()

Check Domain Verified

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
    parentPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    entityPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.CheckDomainVerified(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.CheckDomainVerified``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDomainVerified`: VerificationResult
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.CheckDomainVerified`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDomainVerifiedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**VerificationResult**](VerificationResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainScreenshot

> DomainScreenshot DomainScreenshot(ctx).DomainName(domainName).Identity(identity).Execute()

Domain Screenshot

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
    domainName := "domainName_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.DomainScreenshot(context.Background()).DomainName(domainName).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.DomainScreenshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainScreenshot`: DomainScreenshot
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.DomainScreenshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainScreenshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**DomainScreenshot**](DomainScreenshot.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyDomain

> DomainWithAssets GetCompanyDomain(ctx, parentPk, entityPk).Identity(identity).Execute()

Get Company Domain

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
    parentPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    entityPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.GetCompanyDomain(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetCompanyDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyDomain`: DomainWithAssets
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetCompanyDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**DomainWithAssets**](DomainWithAssets.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyDomains

> PaginatedEntityCollectionDomainEntity ListCompanyDomains(ctx, parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()

List Company Domains

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
    parentPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    search := "search_example" // string |  (optional)
    marker := "marker_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 1000)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.ListCompanyDomains(context.Background(), parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.ListCompanyDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyDomains`: PaginatedEntityCollectionDomainEntity
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.ListCompanyDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **marker** | **string** |  | 
 **limit** | **int32** |  | [default to 1000]
 **identity** | **string** |  | 

### Return type

[**PaginatedEntityCollectionDomainEntity**](PaginatedEntityCollectionDomainEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCompanyDomain

> RemoveCompanyDomain(ctx, parentPk, entityPk).Identity(identity).Execute()

Remove Company Domain

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
    parentPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    entityPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.RemoveCompanyDomain(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.RemoveCompanyDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCompanyDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateCompanyDomain

> UpdateCompanyDomain(ctx, parentPk, entityPk).Domain(domain).Identity(identity).Execute()

Update Company Domain

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
    parentPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    entityPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    domain := *openapiclient.NewDomain("Name_example") // Domain | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.UpdateCompanyDomain(context.Background(), parentPk, entityPk).Domain(domain).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.UpdateCompanyDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domain** | [**Domain**](Domain.md) |  | 
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


## UploadDomainListApiCompaniesParentPkDomainsUploadPost

> UploadResponse UploadDomainListApiCompaniesParentPkDomainsUploadPost(ctx, parentPk).DomainList(domainList).Identity(identity).Execute()

Upload Domain List

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
    parentPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    domainList := os.NewFile(1234, "some_file") // *os.File | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.UploadDomainListApiCompaniesParentPkDomainsUploadPost(context.Background(), parentPk).DomainList(domainList).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.UploadDomainListApiCompaniesParentPkDomainsUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadDomainListApiCompaniesParentPkDomainsUploadPost`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.UploadDomainListApiCompaniesParentPkDomainsUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadDomainListApiCompaniesParentPkDomainsUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainList** | ***os.File** |  | 
 **identity** | **string** |  | 

### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WhoisDomain

> ResponseWhoisDomain WhoisDomain(ctx).DomainName(domainName).Identity(identity).Execute()

Whois Domain

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
    domainName := "domainName_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.WhoisDomain(context.Background()).DomainName(domainName).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.WhoisDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WhoisDomain`: ResponseWhoisDomain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.WhoisDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWhoisDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**ResponseWhoisDomain**](ResponseWhoisDomain.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

