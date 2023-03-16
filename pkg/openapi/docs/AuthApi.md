# \AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackApiAuthCallbackGet**](AuthApi.md#CallbackApiAuthCallbackGet) | **Get** /api/auth/callback | Callback
[**ConfirmSignupApiAuthConfirmEmailCodeGet**](AuthApi.md#ConfirmSignupApiAuthConfirmEmailCodeGet) | **Get** /api/auth/confirm/{email}/{code} | Confirm Signup
[**ForgotPasswordApiAuthForgotPasswordPost**](AuthApi.md#ForgotPasswordApiAuthForgotPasswordPost) | **Post** /api/auth/forgot-password | Forgot Password
[**GoogleRedirectApiAuthGoogleGet**](AuthApi.md#GoogleRedirectApiAuthGoogleGet) | **Get** /api/auth/google | Google Redirect
[**LoginApiAuthLoginPost**](AuthApi.md#LoginApiAuthLoginPost) | **Post** /api/auth/login | Login
[**LogoutApiAuthLogoutGet**](AuthApi.md#LogoutApiAuthLogoutGet) | **Get** /api/auth/logout | Logout
[**MfaChallengeApiAuthMfaChallengePost**](AuthApi.md#MfaChallengeApiAuthMfaChallengePost) | **Post** /api/auth/mfa-challenge | Mfa Challenge
[**MyselfApiAuthMyselfGet**](AuthApi.md#MyselfApiAuthMyselfGet) | **Get** /api/auth/myself | Myself
[**ProfileApiAuthProfileGet**](AuthApi.md#ProfileApiAuthProfileGet) | **Get** /api/auth/profile | Profile
[**ResetPasswordApiAuthResetPasswordPost**](AuthApi.md#ResetPasswordApiAuthResetPasswordPost) | **Post** /api/auth/reset-password | Reset Password
[**SignupApiAuthSignupPost**](AuthApi.md#SignupApiAuthSignupPost) | **Post** /api/auth/signup | Signup
[**SoftwareMfaSetupApiAuthSoftwareMfaSetupPost**](AuthApi.md#SoftwareMfaSetupApiAuthSoftwareMfaSetupPost) | **Post** /api/auth/software-mfa-setup | Software Mfa Setup



## CallbackApiAuthCallbackGet

> interface{} CallbackApiAuthCallbackGet(ctx).State(state).PostLogin(postLogin).Execute()

Callback

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
    postLogin := "postLogin_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.CallbackApiAuthCallbackGet(context.Background()).State(state).PostLogin(postLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.CallbackApiAuthCallbackGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallbackApiAuthCallbackGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.CallbackApiAuthCallbackGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallbackApiAuthCallbackGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 
 **postLogin** | **string** |  | [default to &quot;&quot;]

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


## ConfirmSignupApiAuthConfirmEmailCodeGet

> interface{} ConfirmSignupApiAuthConfirmEmailCodeGet(ctx, email, code).Execute()

Confirm Signup

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
    email := "email_example" // string | 
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ConfirmSignupApiAuthConfirmEmailCodeGet(context.Background(), email, code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ConfirmSignupApiAuthConfirmEmailCodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSignupApiAuthConfirmEmailCodeGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ConfirmSignupApiAuthConfirmEmailCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSignupApiAuthConfirmEmailCodeGetRequest struct via the builder pattern


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


## ForgotPasswordApiAuthForgotPasswordPost

> SuccessResponse ForgotPasswordApiAuthForgotPasswordPost(ctx).EmailBody(emailBody).Execute()

Forgot Password

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
    emailBody := *openapiclient.NewEmailBody("Email_example") // EmailBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ForgotPasswordApiAuthForgotPasswordPost(context.Background()).EmailBody(emailBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ForgotPasswordApiAuthForgotPasswordPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForgotPasswordApiAuthForgotPasswordPost`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ForgotPasswordApiAuthForgotPasswordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordApiAuthForgotPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailBody** | [**EmailBody**](EmailBody.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GoogleRedirectApiAuthGoogleGet

> interface{} GoogleRedirectApiAuthGoogleGet(ctx).Execute()

Google Redirect

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GoogleRedirectApiAuthGoogleGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GoogleRedirectApiAuthGoogleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GoogleRedirectApiAuthGoogleGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.GoogleRedirectApiAuthGoogleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGoogleRedirectApiAuthGoogleGetRequest struct via the builder pattern


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


## LoginApiAuthLoginPost

> LoginResponseMfaChallenge LoginApiAuthLoginPost(ctx).LoginRequest(loginRequest).Execute()

Login

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
    loginRequest := *openapiclient.NewLoginRequest("Email_example", "Password_example") // LoginRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.LoginApiAuthLoginPost(context.Background()).LoginRequest(loginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.LoginApiAuthLoginPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoginApiAuthLoginPost`: LoginResponseMfaChallenge
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.LoginApiAuthLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginApiAuthLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequest** | [**LoginRequest**](LoginRequest.md) |  | 

### Return type

[**LoginResponseMfaChallenge**](LoginResponseMfaChallenge.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutApiAuthLogoutGet

> interface{} LogoutApiAuthLogoutGet(ctx).Execute()

Logout

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.LogoutApiAuthLogoutGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.LogoutApiAuthLogoutGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogoutApiAuthLogoutGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.LogoutApiAuthLogoutGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutApiAuthLogoutGetRequest struct via the builder pattern


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


## MfaChallengeApiAuthMfaChallengePost

> SuccessResponse MfaChallengeApiAuthMfaChallengePost(ctx).MFALoginChallengeRequest(mFALoginChallengeRequest).Execute()

Mfa Challenge

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
    mFALoginChallengeRequest := *openapiclient.NewMFALoginChallengeRequest("Email_example", "Token_example", "Session_example") // MFALoginChallengeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.MfaChallengeApiAuthMfaChallengePost(context.Background()).MFALoginChallengeRequest(mFALoginChallengeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.MfaChallengeApiAuthMfaChallengePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MfaChallengeApiAuthMfaChallengePost`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.MfaChallengeApiAuthMfaChallengePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMfaChallengeApiAuthMfaChallengePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mFALoginChallengeRequest** | [**MFALoginChallengeRequest**](MFALoginChallengeRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyselfApiAuthMyselfGet

> CognitoIdentity MyselfApiAuthMyselfGet(ctx).Identity(identity).Execute()

Myself

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
    identity := "identity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.MyselfApiAuthMyselfGet(context.Background()).Identity(identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.MyselfApiAuthMyselfGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyselfApiAuthMyselfGet`: CognitoIdentity
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.MyselfApiAuthMyselfGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMyselfApiAuthMyselfGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identity** | **string** |  | 

### Return type

[**CognitoIdentity**](CognitoIdentity.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfileApiAuthProfileGet

> AccessUser ProfileApiAuthProfileGet(ctx).Authorization(authorization).Execute()

Profile

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
    authorization := "authorization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ProfileApiAuthProfileGet(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ProfileApiAuthProfileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfileApiAuthProfileGet`: AccessUser
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ProfileApiAuthProfileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProfileApiAuthProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**AccessUser**](AccessUser.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordApiAuthResetPasswordPost

> ResetPasswordResponse ResetPasswordApiAuthResetPasswordPost(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()

Reset Password

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
    resetPasswordRequest := *openapiclient.NewResetPasswordRequest("Email_example", "Code_example", "Password_example") // ResetPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ResetPasswordApiAuthResetPasswordPost(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ResetPasswordApiAuthResetPasswordPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPasswordApiAuthResetPasswordPost`: ResetPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ResetPasswordApiAuthResetPasswordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordApiAuthResetPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

### Return type

[**ResetPasswordResponse**](ResetPasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignupApiAuthSignupPost

> SignupResponse SignupApiAuthSignupPost(ctx).SignupRequest(signupRequest).Execute()

Signup

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
    signupRequest := *openapiclient.NewSignupRequest("Email_example", "Password_example", "Name_example") // SignupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.SignupApiAuthSignupPost(context.Background()).SignupRequest(signupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SignupApiAuthSignupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignupApiAuthSignupPost`: SignupResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SignupApiAuthSignupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupApiAuthSignupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupRequest** | [**SignupRequest**](SignupRequest.md) |  | 

### Return type

[**SignupResponse**](SignupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareMfaSetupApiAuthSoftwareMfaSetupPost

> SuccessResponse SoftwareMfaSetupApiAuthSoftwareMfaSetupPost(ctx).Authorization(authorization).MFASetupChallengeRequest(mFASetupChallengeRequest).Execute()

Software Mfa Setup

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
    authorization := "authorization_example" // string | 
    mFASetupChallengeRequest := *openapiclient.NewMFASetupChallengeRequest("Token_example") // MFASetupChallengeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.SoftwareMfaSetupApiAuthSoftwareMfaSetupPost(context.Background()).Authorization(authorization).MFASetupChallengeRequest(mFASetupChallengeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SoftwareMfaSetupApiAuthSoftwareMfaSetupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareMfaSetupApiAuthSoftwareMfaSetupPost`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SoftwareMfaSetupApiAuthSoftwareMfaSetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareMfaSetupApiAuthSoftwareMfaSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **mFASetupChallengeRequest** | [**MFASetupChallengeRequest**](MFASetupChallengeRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

