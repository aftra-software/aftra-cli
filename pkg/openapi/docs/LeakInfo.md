# LeakInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leaks** | [**[]LeakData**](LeakData.md) |  | 

## Methods

### NewLeakInfo

`func NewLeakInfo(leaks []LeakData, ) *LeakInfo`

NewLeakInfo instantiates a new LeakInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeakInfoWithDefaults

`func NewLeakInfoWithDefaults() *LeakInfo`

NewLeakInfoWithDefaults instantiates a new LeakInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaks

`func (o *LeakInfo) GetLeaks() []LeakData`

GetLeaks returns the Leaks field if non-nil, zero value otherwise.

### GetLeaksOk

`func (o *LeakInfo) GetLeaksOk() (*[]LeakData, bool)`

GetLeaksOk returns a tuple with the Leaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaks

`func (o *LeakInfo) SetLeaks(v []LeakData)`

SetLeaks sets Leaks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


