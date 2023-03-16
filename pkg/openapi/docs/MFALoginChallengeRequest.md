# MFALoginChallengeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Token** | **string** |  | 
**Session** | **string** |  | 

## Methods

### NewMFALoginChallengeRequest

`func NewMFALoginChallengeRequest(email string, token string, session string, ) *MFALoginChallengeRequest`

NewMFALoginChallengeRequest instantiates a new MFALoginChallengeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFALoginChallengeRequestWithDefaults

`func NewMFALoginChallengeRequestWithDefaults() *MFALoginChallengeRequest`

NewMFALoginChallengeRequestWithDefaults instantiates a new MFALoginChallengeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MFALoginChallengeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MFALoginChallengeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MFALoginChallengeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetToken

`func (o *MFALoginChallengeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MFALoginChallengeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MFALoginChallengeRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetSession

`func (o *MFALoginChallengeRequest) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *MFALoginChallengeRequest) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *MFALoginChallengeRequest) SetSession(v string)`

SetSession sets Session field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


