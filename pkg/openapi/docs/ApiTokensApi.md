# \ApiTokensApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIToken**](ApiTokensApi.md#CreateAPIToken) | **Post** /api/companies/{company_pk}/api-tokens/ | Create
[**DeleteAPIToken**](ApiTokensApi.md#DeleteAPIToken) | **Delete** /api/companies/{company_pk}/api-tokens/{token_uid}/ | Delete Token
[**ListAPITokens**](ApiTokensApi.md#ListAPITokens) | **Get** /api/companies/{company_pk}/api-tokens/ | List Tokens



## CreateAPIToken

> Token CreateAPIToken(ctx, companyPk).CreateToken(createToken).ParentPk(parentPk).Identity(identity).Execute()

Create

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
    createToken := *openapiclient.NewCreateToken("Name_example") // CreateToken | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiTokensApi.CreateAPIToken(context.Background(), companyPk).CreateToken(createToken).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokensApi.CreateAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAPIToken`: Token
    fmt.Fprintf(os.Stdout, "Response from `ApiTokensApi.CreateAPIToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createToken** | [**CreateToken**](CreateToken.md) |  | 
 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**Token**](Token.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIToken

> interface{} DeleteAPIToken(ctx, companyPk, tokenUid).ParentPk(parentPk).Identity(identity).Execute()

Delete Token

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
    tokenUid := "tokenUid_example" // string | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiTokensApi.DeleteAPIToken(context.Background(), companyPk, tokenUid).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokensApi.DeleteAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAPIToken`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApiTokensApi.DeleteAPIToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**tokenUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentPk** | **string** |  | 
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


## ListAPITokens

> []MaskedToken ListAPITokens(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

List Tokens

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
    resp, r, err := apiClient.ApiTokensApi.ListAPITokens(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokensApi.ListAPITokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAPITokens`: []MaskedToken
    fmt.Fprintf(os.Stdout, "Response from `ApiTokensApi.ListAPITokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAPITokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**[]MaskedToken**](MaskedToken.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

