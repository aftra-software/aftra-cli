# CognitoIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewCognitoIdentity

`func NewCognitoIdentity(sub string, name string, email string, ) *CognitoIdentity`

NewCognitoIdentity instantiates a new CognitoIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitoIdentityWithDefaults

`func NewCognitoIdentityWithDefaults() *CognitoIdentity`

NewCognitoIdentityWithDefaults instantiates a new CognitoIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *CognitoIdentity) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *CognitoIdentity) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *CognitoIdentity) SetSub(v string)`

SetSub sets Sub field to given value.


### GetName

`func (o *CognitoIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CognitoIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CognitoIdentity) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CognitoIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CognitoIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CognitoIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


