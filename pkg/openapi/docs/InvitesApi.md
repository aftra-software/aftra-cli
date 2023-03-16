# \InvitesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInviteToCompany**](InvitesApi.md#AddInviteToCompany) | **Post** /api/companies/{parent_pk}/invites | Add Invite To Company
[**AddUserToCompany**](InvitesApi.md#AddUserToCompany) | **Get** /api/companies/{parent_pk}/invites/{invite_pk}/ | Add User To Company
[**GetCompanyInvite**](InvitesApi.md#GetCompanyInvite) | **Get** /api/companies/{parent_pk}/invites/{entity_pk} | Get Company Invite
[**ListCompanyInvites**](InvitesApi.md#ListCompanyInvites) | **Get** /api/companies/{parent_pk}/invites | List Company Invites
[**RemoveCompanyInvite**](InvitesApi.md#RemoveCompanyInvite) | **Delete** /api/companies/{parent_pk}/invites/{entity_pk} | Remove Company Invite
[**UpdateCompanyInvite**](InvitesApi.md#UpdateCompanyInvite) | **Put** /api/companies/{parent_pk}/invites/{entity_pk} | Update Company Invite



## AddInviteToCompany

> InviteEntity AddInviteToCompany(ctx, parentPk).Invite(invite).Identity(identity).Execute()

Add Invite To Company

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
    invite := *openapiclient.NewInvite("Name_example", openapiclient.CompanyAccessEnum("admin")) // Invite | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitesApi.AddInviteToCompany(context.Background(), parentPk).Invite(invite).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitesApi.AddInviteToCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInviteToCompany`: InviteEntity
    fmt.Fprintf(os.Stdout, "Response from `InvitesApi.AddInviteToCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInviteToCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invite** | [**Invite**](Invite.md) |  | 
 **identity** | **string** |  | 

### Return type

[**InviteEntity**](InviteEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToCompany

> interface{} AddUserToCompany(ctx, parentPk, invitePk).Identity(identity).Execute()

Add User To Company

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
    invitePk := "Invite-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitesApi.AddUserToCompany(context.Background(), parentPk, invitePk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitesApi.AddUserToCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToCompany`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `InvitesApi.AddUserToCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**invitePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyInvite

> InviteWithAssets GetCompanyInvite(ctx, parentPk, entityPk).Identity(identity).Execute()

Get Company Invite

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
    entityPk := "Invite-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitesApi.GetCompanyInvite(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitesApi.GetCompanyInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyInvite`: InviteWithAssets
    fmt.Fprintf(os.Stdout, "Response from `InvitesApi.GetCompanyInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**InviteWithAssets**](InviteWithAssets.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyInvites

> PaginatedEntityCollectionInviteEntity ListCompanyInvites(ctx, parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()

List Company Invites

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
    resp, r, err := apiClient.InvitesApi.ListCompanyInvites(context.Background(), parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitesApi.ListCompanyInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyInvites`: PaginatedEntityCollectionInviteEntity
    fmt.Fprintf(os.Stdout, "Response from `InvitesApi.ListCompanyInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **marker** | **string** |  | 
 **limit** | **int32** |  | [default to 1000]
 **identity** | **string** |  | 

### Return type

[**PaginatedEntityCollectionInviteEntity**](PaginatedEntityCollectionInviteEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCompanyInvite

> RemoveCompanyInvite(ctx, parentPk, entityPk).Identity(identity).Execute()

Remove Company Invite

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
    entityPk := "Invite-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitesApi.RemoveCompanyInvite(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitesApi.RemoveCompanyInvite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveCompanyInviteRequest struct via the builder pattern


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


## UpdateCompanyInvite

> UpdateCompanyInvite(ctx, parentPk, entityPk).Invite(invite).Identity(identity).Execute()

Update Company Invite

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
    entityPk := "Invite-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    invite := *openapiclient.NewInvite("Name_example", openapiclient.CompanyAccessEnum("admin")) // Invite | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitesApi.UpdateCompanyInvite(context.Background(), parentPk, entityPk).Invite(invite).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitesApi.UpdateCompanyInvite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCompanyInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invite** | [**Invite**](Invite.md) |  | 
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

