# \DomainsSubdomainsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSubdomainToDomain**](DomainsSubdomainsApi.md#AddSubdomainToDomain) | **Post** /api/domains/{parent_pk}/subdomains | Add Subdomain To Domain
[**GetDomainSubdomain**](DomainsSubdomainsApi.md#GetDomainSubdomain) | **Get** /api/domains/{parent_pk}/subdomains/{entity_pk} | Get Domain Subdomain
[**ListDomainSubdomains**](DomainsSubdomainsApi.md#ListDomainSubdomains) | **Get** /api/domains/{parent_pk}/subdomains | List Domain Subdomains
[**RemoveDomainSubdomain**](DomainsSubdomainsApi.md#RemoveDomainSubdomain) | **Delete** /api/domains/{parent_pk}/subdomains/{entity_pk} | Remove Domain Subdomain
[**UpdateDomainSubdomain**](DomainsSubdomainsApi.md#UpdateDomainSubdomain) | **Put** /api/domains/{parent_pk}/subdomains/{entity_pk} | Update Domain Subdomain



## AddSubdomainToDomain

> SubdomainEntity AddSubdomainToDomain(ctx, parentPk).Domain(domain).Identity(identity).Execute()

Add Subdomain To Domain

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
    parentPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    domain := *openapiclient.NewDomain("Name_example") // Domain | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsSubdomainsApi.AddSubdomainToDomain(context.Background(), parentPk).Domain(domain).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsSubdomainsApi.AddSubdomainToDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSubdomainToDomain`: SubdomainEntity
    fmt.Fprintf(os.Stdout, "Response from `DomainsSubdomainsApi.AddSubdomainToDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSubdomainToDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | [**Domain**](Domain.md) |  | 
 **identity** | **string** |  | 

### Return type

[**SubdomainEntity**](SubdomainEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainSubdomain

> SubdomainWithAssets GetDomainSubdomain(ctx, parentPk, entityPk).Identity(identity).Execute()

Get Domain Subdomain

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
    parentPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    entityPk := "Subdomain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsSubdomainsApi.GetDomainSubdomain(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsSubdomainsApi.GetDomainSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainSubdomain`: SubdomainWithAssets
    fmt.Fprintf(os.Stdout, "Response from `DomainsSubdomainsApi.GetDomainSubdomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**SubdomainWithAssets**](SubdomainWithAssets.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomainSubdomains

> PaginatedEntityCollectionSubdomainEntity ListDomainSubdomains(ctx, parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()

List Domain Subdomains

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
    parentPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    search := "search_example" // string |  (optional)
    marker := "marker_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 1000)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsSubdomainsApi.ListDomainSubdomains(context.Background(), parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsSubdomainsApi.ListDomainSubdomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomainSubdomains`: PaginatedEntityCollectionSubdomainEntity
    fmt.Fprintf(os.Stdout, "Response from `DomainsSubdomainsApi.ListDomainSubdomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **marker** | **string** |  | 
 **limit** | **int32** |  | [default to 1000]
 **identity** | **string** |  | 

### Return type

[**PaginatedEntityCollectionSubdomainEntity**](PaginatedEntityCollectionSubdomainEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDomainSubdomain

> RemoveDomainSubdomain(ctx, parentPk, entityPk).Identity(identity).Execute()

Remove Domain Subdomain

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
    parentPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    entityPk := "Subdomain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsSubdomainsApi.RemoveDomainSubdomain(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsSubdomainsApi.RemoveDomainSubdomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveDomainSubdomainRequest struct via the builder pattern


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


## UpdateDomainSubdomain

> UpdateDomainSubdomain(ctx, parentPk, entityPk).Domain(domain).Identity(identity).Execute()

Update Domain Subdomain

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
    parentPk := "Domain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    entityPk := "Subdomain-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    domain := *openapiclient.NewDomain("Name_example") // Domain | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsSubdomainsApi.UpdateDomainSubdomain(context.Background(), parentPk, entityPk).Domain(domain).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsSubdomainsApi.UpdateDomainSubdomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateDomainSubdomainRequest struct via the builder pattern


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

