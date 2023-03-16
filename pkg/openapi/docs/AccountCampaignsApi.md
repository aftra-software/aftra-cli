# \AccountCampaignsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet**](AccountCampaignsApi.md#GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet) | **Get** /api/account-campaigns/{campaign_pk}/campaign/{campaign_sk} | Get Account Campaign
[**GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet**](AccountCampaignsApi.md#GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet) | **Get** /api/account-campaigns/{campaign_pk}/campaign/{campaign_sk}/leaks-info | Get Account Campaign Leaks Info
[**ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost**](AccountCampaignsApi.md#ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost) | **Post** /api/account-campaigns/{campaign_pk}/campaign/{campaign_sk} | Resolve Account Campaign
[**ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet**](AccountCampaignsApi.md#ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet) | **Get** /api/account-campaigns/{campaign_pk}/campaign/{campaign_sk}/view | View Account Campaign



## GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet

> AccountCampaignEntity GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet(ctx, campaignPk, campaignSk).Execute()

Get Account Campaign

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
    campaignPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    campaignSk := "AccountCampaign-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountCampaignsApi.GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet(context.Background(), campaignPk, campaignSk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountCampaignsApi.GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet`: AccountCampaignEntity
    fmt.Fprintf(os.Stdout, "Response from `AccountCampaignsApi.GetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignPk** | **string** |  | 
**campaignSk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountCampaignEntity**](AccountCampaignEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet

> LeakInfo GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet(ctx, campaignPk, campaignSk).Execute()

Get Account Campaign Leaks Info

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
    campaignPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    campaignSk := "AccountCampaign-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountCampaignsApi.GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet(context.Background(), campaignPk, campaignSk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountCampaignsApi.GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet`: LeakInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountCampaignsApi.GetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignPk** | **string** |  | 
**campaignSk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCampaignLeaksInfoApiAccountCampaignsCampaignPkCampaignCampaignSkLeaksInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LeakInfo**](LeakInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost

> interface{} ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost(ctx, campaignPk, campaignSk).AccountResolvePayload(accountResolvePayload).Execute()

Resolve Account Campaign

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
    campaignPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    campaignSk := "AccountCampaign-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    accountResolvePayload := *openapiclient.NewAccountResolvePayload(false) // AccountResolvePayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountCampaignsApi.ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost(context.Background(), campaignPk, campaignSk).AccountResolvePayload(accountResolvePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountCampaignsApi.ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountCampaignsApi.ResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignPk** | **string** |  | 
**campaignSk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountResolvePayload** | [**AccountResolvePayload**](AccountResolvePayload.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet

> interface{} ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet(ctx, campaignPk, campaignSk).Execute()

View Account Campaign

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
    campaignPk := "Account-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    campaignSk := "AccountCampaign-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountCampaignsApi.ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet(context.Background(), campaignPk, campaignSk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountCampaignsApi.ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountCampaignsApi.ViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignPk** | **string** |  | 
**campaignSk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewAccountCampaignApiAccountCampaignsCampaignPkCampaignCampaignSkViewGetRequest struct via the builder pattern


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

