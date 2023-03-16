# UploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessCount** | **int32** |  | 
**FailedCount** | **int32** |  | 
**Failures** | [**[]UploadFailure**](UploadFailure.md) |  | 

## Methods

### NewUploadResponse

`func NewUploadResponse(successCount int32, failedCount int32, failures []UploadFailure, ) *UploadResponse`

NewUploadResponse instantiates a new UploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadResponseWithDefaults

`func NewUploadResponseWithDefaults() *UploadResponse`

NewUploadResponseWithDefaults instantiates a new UploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessCount

`func (o *UploadResponse) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UploadResponse) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UploadResponse) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.


### GetFailedCount

`func (o *UploadResponse) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *UploadResponse) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *UploadResponse) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.


### GetFailures

`func (o *UploadResponse) GetFailures() []UploadFailure`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *UploadResponse) GetFailuresOk() (*[]UploadFailure, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *UploadResponse) SetFailures(v []UploadFailure)`

SetFailures sets Failures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


