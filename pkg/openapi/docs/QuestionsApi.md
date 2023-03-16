# \QuestionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityQuestions**](QuestionsApi.md#GetSecurityQuestions) | **Get** /api/companies/{company_pk}/questions/ | Get Questions
[**PostUpdateAnswer**](QuestionsApi.md#PostUpdateAnswer) | **Post** /api/companies/{company_pk}/questions/{question_uid}/ | Post Update Answer



## GetSecurityQuestions

> CompanyQuestionList GetSecurityQuestions(ctx, companyPk).ParentPk(parentPk).Identity(identity).Execute()

Get Questions

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
    resp, r, err := apiClient.QuestionsApi.GetSecurityQuestions(context.Background(), companyPk).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionsApi.GetSecurityQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityQuestions`: CompanyQuestionList
    fmt.Fprintf(os.Stdout, "Response from `QuestionsApi.GetSecurityQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**CompanyQuestionList**](CompanyQuestionList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUpdateAnswer

> CompanyQuestionList PostUpdateAnswer(ctx, companyPk, questionUid).AnswerUpdate(answerUpdate).ParentPk(parentPk).Identity(identity).Execute()

Post Update Answer

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
    questionUid := "questionUid_example" // string | 
    answerUpdate := *openapiclient.NewAnswerUpdate("Answer_example") // AnswerUpdate | 
    parentPk := "parentPk_example" // string |  (optional)
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuestionsApi.PostUpdateAnswer(context.Background(), companyPk, questionUid).AnswerUpdate(answerUpdate).ParentPk(parentPk).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionsApi.PostUpdateAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUpdateAnswer`: CompanyQuestionList
    fmt.Fprintf(os.Stdout, "Response from `QuestionsApi.PostUpdateAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyPk** | **string** |  | 
**questionUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUpdateAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **answerUpdate** | [**AnswerUpdate**](AnswerUpdate.md) |  | 
 **parentPk** | **string** |  | 
 **identity** | **string** |  | 

### Return type

[**CompanyQuestionList**](CompanyQuestionList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

