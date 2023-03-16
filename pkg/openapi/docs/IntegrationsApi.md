# \IntegrationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwsApiIntegrationsAwsCompanyPkGet**](IntegrationsApi.md#AwsApiIntegrationsAwsCompanyPkGet) | **Get** /api/integrations/aws/{company_pk} | Aws
[**AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost**](IntegrationsApi.md#AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost) | **Post** /api/integrations/aws/{company_pk}/{user_pk} | Aws Callback
[**BitbucketApiIntegrationsBitbucketCompanyPkGet**](IntegrationsApi.md#BitbucketApiIntegrationsBitbucketCompanyPkGet) | **Get** /api/integrations/bitbucket/{company_pk} | Bitbucket
[**BitbucketCallbackApiIntegrationsBitbucketGet**](IntegrationsApi.md#BitbucketCallbackApiIntegrationsBitbucketGet) | **Get** /api/integrations/bitbucket | Bitbucket Callback
[**GetCompanyIntegration**](IntegrationsApi.md#GetCompanyIntegration) | **Get** /api/integrations/crud/{parent_pk}/{entity_pk} | Get Company Integration
[**GithubApiIntegrationsGithubCompanyPkGet**](IntegrationsApi.md#GithubApiIntegrationsGithubCompanyPkGet) | **Get** /api/integrations/github/{company_pk} | Github
[**GithubCallbackApiIntegrationsGithubGet**](IntegrationsApi.md#GithubCallbackApiIntegrationsGithubGet) | **Get** /api/integrations/github | Github Callback
[**GitlabApiIntegrationsGitlabCompanyPkGet**](IntegrationsApi.md#GitlabApiIntegrationsGitlabCompanyPkGet) | **Get** /api/integrations/gitlab/{company_pk} | Gitlab
[**GitlabCallbackApiIntegrationsGitlabGet**](IntegrationsApi.md#GitlabCallbackApiIntegrationsGitlabGet) | **Get** /api/integrations/gitlab | Gitlab Callback
[**GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet**](IntegrationsApi.md#GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet) | **Get** /api/integrations/google-workspaces/{company_pk} | Google Workspaces
[**GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet**](IntegrationsApi.md#GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet) | **Get** /api/integrations/google-workspaces | Google Workspaces Callback
[**ListCompanyIntegrations**](IntegrationsApi.md#ListCompanyIntegrations) | **Get** /api/integrations/crud/{parent_pk} | List Company Integrations
[**MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet**](IntegrationsApi.md#MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet) | **Get** /api/integrations/microsoft-azure/{company_pk} | Microsoft Azure Ad
[**MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet**](IntegrationsApi.md#MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet) | **Get** /api/integrations/microsoft-azure | Microsoft Azure Ad Callback
[**RemoveCompanyIntegration**](IntegrationsApi.md#RemoveCompanyIntegration) | **Delete** /api/integrations/crud/{parent_pk}/{entity_pk} | Remove Company Integration



## AwsApiIntegrationsAwsCompanyPkGet

> interface{} AwsApiIntegrationsAwsCompanyPkGet(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Aws



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
    resp, r, err := apiClient.IntegrationsApi.AwsApiIntegrationsAwsCompanyPkGet(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.AwsApiIntegrationsAwsCompanyPkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsApiIntegrationsAwsCompanyPkGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.AwsApiIntegrationsAwsCompanyPkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsApiIntegrationsAwsCompanyPkGetRequest struct via the builder pattern


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


## AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost

> interface{} AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost(ctx, companyPk, userPk).AWSRoute53IntegrationCallbackData(aWSRoute53IntegrationCallbackData).Execute()

Aws Callback



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
    userPk := "User-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    aWSRoute53IntegrationCallbackData := *openapiclient.NewAWSRoute53IntegrationCallbackData("RoleArn_example") // AWSRoute53IntegrationCallbackData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost(context.Background(), companyPk, userPk).AWSRoute53IntegrationCallbackData(aWSRoute53IntegrationCallbackData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.AwsCallbackApiIntegrationsAwsCompanyPkUserPkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**userPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsCallbackApiIntegrationsAwsCompanyPkUserPkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aWSRoute53IntegrationCallbackData** | [**AWSRoute53IntegrationCallbackData**](AWSRoute53IntegrationCallbackData.md) |  | 

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


## BitbucketApiIntegrationsBitbucketCompanyPkGet

> interface{} BitbucketApiIntegrationsBitbucketCompanyPkGet(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Bitbucket



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
    resp, r, err := apiClient.IntegrationsApi.BitbucketApiIntegrationsBitbucketCompanyPkGet(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.BitbucketApiIntegrationsBitbucketCompanyPkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BitbucketApiIntegrationsBitbucketCompanyPkGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.BitbucketApiIntegrationsBitbucketCompanyPkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBitbucketApiIntegrationsBitbucketCompanyPkGetRequest struct via the builder pattern


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


## BitbucketCallbackApiIntegrationsBitbucketGet

> interface{} BitbucketCallbackApiIntegrationsBitbucketGet(ctx).State(state).BitbucketState(bitbucketState).BitbucketCompany(bitbucketCompany).Identity(identity).Execute()

Bitbucket Callback



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
    state := "state_example" // string | 
    bitbucketState := "bitbucketState_example" // string | 
    bitbucketCompany := "bitbucketCompany_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.BitbucketCallbackApiIntegrationsBitbucketGet(context.Background()).State(state).BitbucketState(bitbucketState).BitbucketCompany(bitbucketCompany).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.BitbucketCallbackApiIntegrationsBitbucketGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BitbucketCallbackApiIntegrationsBitbucketGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.BitbucketCallbackApiIntegrationsBitbucketGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBitbucketCallbackApiIntegrationsBitbucketGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 
 **bitbucketState** | **string** |  | 
 **bitbucketCompany** | **string** |  | 
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


## GetCompanyIntegration

> IntegrationWithAssets GetCompanyIntegration(ctx, parentPk, entityPk).Identity(identity).Execute()

Get Company Integration

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
    entityPk := "Integration-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GetCompanyIntegration(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetCompanyIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyIntegration`: IntegrationWithAssets
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetCompanyIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 
**entityPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **string** |  | 

### Return type

[**IntegrationWithAssets**](IntegrationWithAssets.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GithubApiIntegrationsGithubCompanyPkGet

> interface{} GithubApiIntegrationsGithubCompanyPkGet(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Github



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
    resp, r, err := apiClient.IntegrationsApi.GithubApiIntegrationsGithubCompanyPkGet(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GithubApiIntegrationsGithubCompanyPkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GithubApiIntegrationsGithubCompanyPkGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GithubApiIntegrationsGithubCompanyPkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGithubApiIntegrationsGithubCompanyPkGetRequest struct via the builder pattern


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


## GithubCallbackApiIntegrationsGithubGet

> interface{} GithubCallbackApiIntegrationsGithubGet(ctx).InstallationId(installationId).GithubCompany(githubCompany).Identity(identity).Execute()

Github Callback



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
    installationId := "installationId_example" // string | 
    githubCompany := "githubCompany_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GithubCallbackApiIntegrationsGithubGet(context.Background()).InstallationId(installationId).GithubCompany(githubCompany).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GithubCallbackApiIntegrationsGithubGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GithubCallbackApiIntegrationsGithubGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GithubCallbackApiIntegrationsGithubGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGithubCallbackApiIntegrationsGithubGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installationId** | **string** |  | 
 **githubCompany** | **string** |  | 
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


## GitlabApiIntegrationsGitlabCompanyPkGet

> interface{} GitlabApiIntegrationsGitlabCompanyPkGet(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Gitlab



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
    resp, r, err := apiClient.IntegrationsApi.GitlabApiIntegrationsGitlabCompanyPkGet(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GitlabApiIntegrationsGitlabCompanyPkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitlabApiIntegrationsGitlabCompanyPkGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GitlabApiIntegrationsGitlabCompanyPkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitlabApiIntegrationsGitlabCompanyPkGetRequest struct via the builder pattern


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


## GitlabCallbackApiIntegrationsGitlabGet

> interface{} GitlabCallbackApiIntegrationsGitlabGet(ctx).State(state).GitlabState(gitlabState).GitlabCompany(gitlabCompany).Identity(identity).Execute()

Gitlab Callback



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
    state := "state_example" // string | 
    gitlabState := "gitlabState_example" // string | 
    gitlabCompany := "gitlabCompany_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GitlabCallbackApiIntegrationsGitlabGet(context.Background()).State(state).GitlabState(gitlabState).GitlabCompany(gitlabCompany).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GitlabCallbackApiIntegrationsGitlabGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitlabCallbackApiIntegrationsGitlabGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GitlabCallbackApiIntegrationsGitlabGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGitlabCallbackApiIntegrationsGitlabGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 
 **gitlabState** | **string** |  | 
 **gitlabCompany** | **string** |  | 
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


## GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet

> interface{} GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Google Workspaces



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
    resp, r, err := apiClient.IntegrationsApi.GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGoogleWorkspacesApiIntegrationsGoogleWorkspacesCompanyPkGetRequest struct via the builder pattern


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


## GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet

> interface{} GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet(ctx).State(state).Code(code).GoogleWorkspacesState(googleWorkspacesState).GoogleWorkspacesCompany(googleWorkspacesCompany).Identity(identity).Execute()

Google Workspaces Callback



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
    state := "state_example" // string | 
    code := "code_example" // string | 
    googleWorkspacesState := "googleWorkspacesState_example" // string | 
    googleWorkspacesCompany := "googleWorkspacesCompany_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet(context.Background()).State(state).Code(code).GoogleWorkspacesState(googleWorkspacesState).GoogleWorkspacesCompany(googleWorkspacesCompany).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGoogleWorkspacesCallbackApiIntegrationsGoogleWorkspacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 
 **code** | **string** |  | 
 **googleWorkspacesState** | **string** |  | 
 **googleWorkspacesCompany** | **string** |  | 
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


## ListCompanyIntegrations

> PaginatedEntityCollectionIntegrationEntity ListCompanyIntegrations(ctx, parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()

List Company Integrations

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
    resp, r, err := apiClient.IntegrationsApi.ListCompanyIntegrations(context.Background(), parentPk).Search(search).Marker(marker).Limit(limit).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListCompanyIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyIntegrations`: PaginatedEntityCollectionIntegrationEntity
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListCompanyIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **marker** | **string** |  | 
 **limit** | **int32** |  | [default to 1000]
 **identity** | **string** |  | 

### Return type

[**PaginatedEntityCollectionIntegrationEntity**](PaginatedEntityCollectionIntegrationEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet

> interface{} MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Microsoft Azure Ad



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
    resp, r, err := apiClient.IntegrationsApi.MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.MicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMicrosoftAzureAdApiIntegrationsMicrosoftAzureCompanyPkGetRequest struct via the builder pattern


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


## MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet

> interface{} MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet(ctx).MicrosoftAzureAdFlow(microsoftAzureAdFlow).MicrosoftAzureAdCompany(microsoftAzureAdCompany).Identity(identity).Execute()

Microsoft Azure Ad Callback



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
    microsoftAzureAdFlow := "microsoftAzureAdFlow_example" // string | 
    microsoftAzureAdCompany := "microsoftAzureAdCompany_example" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet(context.Background()).MicrosoftAzureAdFlow(microsoftAzureAdFlow).MicrosoftAzureAdCompany(microsoftAzureAdCompany).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.MicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMicrosoftAzureAdCallbackApiIntegrationsMicrosoftAzureGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftAzureAdFlow** | **string** |  | 
 **microsoftAzureAdCompany** | **string** |  | 
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


## RemoveCompanyIntegration

> RemoveCompanyIntegration(ctx, parentPk, entityPk).Identity(identity).Execute()

Remove Company Integration

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
    entityPk := "Integration-8ccfcc01-da99-4f5a-8ad7-946d6a6f48f5" // string | 
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.RemoveCompanyIntegration(context.Background(), parentPk, entityPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.RemoveCompanyIntegration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveCompanyIntegrationRequest struct via the builder pattern


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

