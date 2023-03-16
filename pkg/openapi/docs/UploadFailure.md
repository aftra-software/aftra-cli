# UploadFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | **string** |  | 
**Reason** | [**UploadFailureReasonEnum**](UploadFailureReasonEnum.md) |  | 

## Methods

### NewUploadFailure

`func NewUploadFailure(line string, reason UploadFailureReasonEnum, ) *UploadFailure`

NewUploadFailure instantiates a new UploadFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFailureWithDefaults

`func NewUploadFailureWithDefaults() *UploadFailure`

NewUploadFailureWithDefaults instantiates a new UploadFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *UploadFailure) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *UploadFailure) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *UploadFailure) SetLine(v string)`

SetLine sets Line field to given value.


### GetReason

`func (o *UploadFailure) GetReason() UploadFailureReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UploadFailure) GetReasonOk() (*UploadFailureReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UploadFailure) SetReason(v UploadFailureReasonEnum)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


