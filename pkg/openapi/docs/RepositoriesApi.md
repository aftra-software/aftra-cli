# \RepositoriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyRepository**](RepositoriesApi.md#GetCompanyRepository) | **Get** /api/companies/{parent_pk}/repositories/{entity_pk} | Get Company Repository
[**ListCompanyRepositories**](RepositoriesApi.md#ListCompanyRepositories) | **Get** /api/companies/{parent_pk}/repositories | List Company Repositories
[**RemoveCompanyRepository**](RepositoriesApi.md#RemoveCompanyRepository) | **Delete** /api/companies/{parent_pk}/repositories/{entity_pk} | Remove Company Repository



## GetCompanyRepository

> RepositoryWithAssets GetCompanyRepository(ctx, parentPk, entityPk).Identity(identity).Execute()

Get Company Repository

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
    entityPk := "Repository-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetCompanyRepository(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetCompanyRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyRepository`: RepositoryWithAssets
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetCompanyRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**RepositoryWithAssets**](RepositoryWithAssets.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyRepositories

> PaginatedEntityCollectionRepositoryEntity ListCompanyRepositories(ctx, parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()

List Company Repositories

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
    resp, r, err := apiClient.RepositoriesApi.ListCompanyRepositories(context.Background(), parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.ListCompanyRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyRepositories`: PaginatedEntityCollectionRepositoryEntity
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.ListCompanyRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **marker** | **string** |  | 
 **limit** | **int32** |  | [default to 1000]
 **identity** | **string** |  | 

### Return type

[**PaginatedEntityCollectionRepositoryEntity**](PaginatedEntityCollectionRepositoryEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCompanyRepository

> RemoveCompanyRepository(ctx, parentPk, entityPk).Identity(identity).Execute()

Remove Company Repository

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
    entityPk := "Repository-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.RemoveCompanyRepository(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RemoveCompanyRepository``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveCompanyRepositoryRequest struct via the builder pattern


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

