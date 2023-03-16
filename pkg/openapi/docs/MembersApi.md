# \MembersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyMember**](MembersApi.md#DeleteCompanyMember) | **Delete** /api/companies/{parent_pk}/members/{user_pk}/ | Delete Company Member
[**GetCompanyMembers**](MembersApi.md#GetCompanyMembers) | **Get** /api/companies/{parent_pk}/members/ | Get Company Members
[**SendCompanyInvite**](MembersApi.md#SendCompanyInvite) | **Post** /api/companies/{parent_pk}/members/ | Create And Send Invite
[**UpdateMemberPermission**](MembersApi.md#UpdateMemberPermission) | **Put** /api/companies/{parent_pk}/members/{user_pk}/ | Update Member Permission



## DeleteCompanyMember

> DeleteCompanyMember(ctx, parentPk, userPk).CompanyPk(companyPk).Identity(identity).Execute()

Delete Company Member

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
    parentPk := "parentPk_example" // string | 
    userPk := "User-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersApi.DeleteCompanyMember(context.Background(), parentPk, userPk).CompanyPk(companyPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.DeleteCompanyMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**userPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **companyPk** | **string** |  | 
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


## GetCompanyMembers

> []UserCompanyAccessEntity GetCompanyMembers(ctx, parentPk).CompanyPk(companyPk).Identity(identity).Execute()

Get Company Members

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
    parentPk := "parentPk_example" // string | 
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersApi.GetCompanyMembers(context.Background(), parentPk).CompanyPk(companyPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.GetCompanyMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyMembers`: []UserCompanyAccessEntity
    fmt.Fprintf(os.Stdout, "Response from `MembersApi.GetCompanyMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**[]UserCompanyAccessEntity**](UserCompanyAccessEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCompanyInvite

> interface{} SendCompanyInvite(ctx, parentPk).Invite(invite).CompanyPk(companyPk).Identity(identity).Execute()

Create And Send Invite

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
    parentPk := "parentPk_example" // string | 
    invite := *openapiclient.NewInvite("Name_example", openapiclient.CompanyAccessEnum("admin")) // Invite | 
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersApi.SendCompanyInvite(context.Background(), parentPk).Invite(invite).CompanyPk(companyPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.SendCompanyInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendCompanyInvite`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MembersApi.SendCompanyInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCompanyInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invite** | [**Invite**](Invite.md) |  | 
 **companyPk** | **string** |  | 
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


## UpdateMemberPermission

> UpdateMemberPermission(ctx, parentPk, userPk).AccessTypeUpdate(accessTypeUpdate).CompanyPk(companyPk).Identity(identity).Execute()

Update Member Permission

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
    parentPk := "parentPk_example" // string | 
    userPk := "User-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    accessTypeUpdate := *openapiclient.NewAccessTypeUpdate(openapiclient.CompanyAccessEnum("admin")) // AccessTypeUpdate | 
    companyPk := "Company-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersApi.UpdateMemberPermission(context.Background(), parentPk, userPk).AccessTypeUpdate(accessTypeUpdate).CompanyPk(companyPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.UpdateMemberPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**userPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessTypeUpdate** | [**AccessTypeUpdate**](AccessTypeUpdate.md) |  | 
 **companyPk** | **string** |  | 
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

