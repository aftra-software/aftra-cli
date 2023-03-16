# ResetPasswordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewResetPasswordResponse

`func NewResetPasswordResponse(success bool, ) *ResetPasswordResponse`

NewResetPasswordResponse instantiates a new ResetPasswordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordResponseWithDefaults

`func NewResetPasswordResponseWithDefaults() *ResetPasswordResponse`

NewResetPasswordResponseWithDefaults instantiates a new ResetPasswordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResetPasswordResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResetPasswordResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResetPasswordResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetReason

`func (o *ResetPasswordResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResetPasswordResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResetPasswordResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResetPasswordResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


