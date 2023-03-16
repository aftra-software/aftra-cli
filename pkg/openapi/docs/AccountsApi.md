# \AccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountToCompany**](AccountsApi.md#AddAccountToCompany) | **Post** /api/companies/{parent_pk}/accounts | Add Account To Company
[**CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost**](AccountsApi.md#CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost) | **Post** /api/companies/{parent_pk}/accounts/{entity_pk}/create-campaign | Create Account Campaign
[**GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet**](AccountsApi.md#GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet) | **Get** /api/companies/{parent_pk}/accounts/{entity_pk}/services | Get Account Exposed Services
[**GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet**](AccountsApi.md#GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet) | **Get** /api/companies/{parent_pk}/accounts/services-summary | Get Account Exposed Services
[**GetCompanyAccount**](AccountsApi.md#GetCompanyAccount) | **Get** /api/companies/{parent_pk}/accounts/{entity_pk} | Get Company Account
[**ListCompanyAccounts**](AccountsApi.md#ListCompanyAccounts) | **Get** /api/companies/{parent_pk}/accounts | List Company Accounts
[**RemoveCompanyAccount**](AccountsApi.md#RemoveCompanyAccount) | **Delete** /api/companies/{parent_pk}/accounts/{entity_pk} | Remove Company Account
[**SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost**](AccountsApi.md#SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost) | **Post** /api/companies/{parent_pk}/accounts/{entity_pk}/verify | Send Email Verification
[**UpdateCompanyAccount**](AccountsApi.md#UpdateCompanyAccount) | **Put** /api/companies/{parent_pk}/accounts/{entity_pk} | Update Company Account
[**UploadAccountListApiCompaniesParentPkAccountsUploadPost**](AccountsApi.md#UploadAccountListApiCompaniesParentPkAccountsUploadPost) | **Post** /api/companies/{parent_pk}/accounts/upload | Upload Account List
[**VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet**](AccountsApi.md#VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet) | **Get** /api/companies/{parent_pk}/accounts/{entity_pk}/verify/{verification_secret} | Verify Account



## AddAccountToCompany

> AccountEntity AddAccountToCompany(ctx, parentPk).Account(account).Identity(identity).Execute()

Add Account To Company

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
    account := *openapiclient.NewAccount("Name_example") // Account | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.AddAccountToCompany(context.Background(), parentPk).Account(account).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AddAccountToCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccountToCompany`: AccountEntity
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AddAccountToCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountToCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**Account**](Account.md) |  | 
 **identity** | **string** |  | 

### Return type

[**AccountEntity**](AccountEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost

> interface{} CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost(ctx, parentPk, entityPk).Identity(identity).Execute()

Create Account Campaign

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
    entityPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCampaignApiCompaniesParentPkAccountsEntityPkCreateCampaignPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet

> []string GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet(ctx, parentPk, entityPk).Identity(identity).Execute()

Get Account Exposed Services

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
    entityPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountExposedServicesApiCompaniesParentPkAccountsEntityPkServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

**[]string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet

> AccountServicesSummary GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet(ctx, parentPk).Identity(identity).Execute()

Get Account Exposed Services

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
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet(context.Background(), parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet`: AccountServicesSummary
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountExposedServicesApiCompaniesParentPkAccountsServicesSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identity** | **string** |  | 

### Return type

[**AccountServicesSummary**](AccountServicesSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyAccount

> AccountWithAssets GetCompanyAccount(ctx, parentPk, entityPk).Identity(identity).Execute()

Get Company Account

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
    entityPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetCompanyAccount(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetCompanyAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyAccount`: AccountWithAssets
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetCompanyAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**AccountWithAssets**](AccountWithAssets.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyAccounts

> PaginatedEntityCollectionAccountEntity ListCompanyAccounts(ctx, parentPk).Search(search).PresentInService(presentInService).Marker(marker).Limit(limit).Identity(identity).Execute()

List Company Accounts

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
    presentInService := "presentInService_example" // string |  (optional)
    marker := "marker_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 1000)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ListCompanyAccounts(context.Background(), parentPk).Search(search).PresentInService(presentInService).Marker(marker).Limit(limit).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListCompanyAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyAccounts`: PaginatedEntityCollectionAccountEntity
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListCompanyAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **presentInService** | **string** |  | 
 **marker** | **string** |  | 
 **limit** | **int32** |  | [default to 1000]
 **identity** | **string** |  | 

### Return type

[**PaginatedEntityCollectionAccountEntity**](PaginatedEntityCollectionAccountEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCompanyAccount

> RemoveCompanyAccount(ctx, parentPk, entityPk).Identity(identity).Execute()

Remove Company Account

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
    entityPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.RemoveCompanyAccount(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.RemoveCompanyAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveCompanyAccountRequest struct via the builder pattern


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


## SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost

> interface{} SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost(ctx, parentPk, entityPk).Identity(identity).Execute()

Send Email Verification

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
    entityPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.SendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailVerificationApiCompaniesParentPkAccountsEntityPkVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyAccount

> UpdateCompanyAccount(ctx, parentPk, entityPk).Account(account).Identity(identity).Execute()

Update Company Account

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
    entityPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    account := *openapiclient.NewAccount("Name_example") // Account | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UpdateCompanyAccount(context.Background(), parentPk, entityPk).Account(account).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateCompanyAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCompanyAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **account** | [**Account**](Account.md) |  | 
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


## UploadAccountListApiCompaniesParentPkAccountsUploadPost

> UploadResponse UploadAccountListApiCompaniesParentPkAccountsUploadPost(ctx, parentPk).AccountList(accountList).Identity(identity).Execute()

Upload Account List

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
    accountList := os.NewFile(1234, "some_file") // *os.File | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UploadAccountListApiCompaniesParentPkAccountsUploadPost(context.Background(), parentPk).AccountList(accountList).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UploadAccountListApiCompaniesParentPkAccountsUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadAccountListApiCompaniesParentPkAccountsUploadPost`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UploadAccountListApiCompaniesParentPkAccountsUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAccountListApiCompaniesParentPkAccountsUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountList** | ***os.File** |  | 
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


## VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet

> interface{} VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet(ctx, parentPk, entityPk, verificationSecret).Execute()

Verify Account

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
    entityPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    verificationSecret := "verificationSecret_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet(context.Background(), parentPk, entityPk, verificationSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.VerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 
**verificationSecret** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAccountApiCompaniesParentPkAccountsEntityPkVerifyVerificationSecretGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

