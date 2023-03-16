# AccountServicesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "account-services-summary"]
**ServiceCounts** | **map[string]int32** |  | 

## Methods

### NewAccountServicesSummary

`func NewAccountServicesSummary(pk string, serviceCounts map[string]int32, ) *AccountServicesSummary`

NewAccountServicesSummary instantiates a new AccountServicesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountServicesSummaryWithDefaults

`func NewAccountServicesSummaryWithDefaults() *AccountServicesSummary`

NewAccountServicesSummaryWithDefaults instantiates a new AccountServicesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AccountServicesSummary) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AccountServicesSummary) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AccountServicesSummary) SetPk(v string)`

SetPk sets Pk field to given value.


### GetUpdated

`func (o *AccountServicesSummary) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AccountServicesSummary) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AccountServicesSummary) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AccountServicesSummary) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetType

`func (o *AccountServicesSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountServicesSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountServicesSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountServicesSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServiceCounts

`func (o *AccountServicesSummary) GetServiceCounts() map[string]int32`

GetServiceCounts returns the ServiceCounts field if non-nil, zero value otherwise.

### GetServiceCountsOk

`func (o *AccountServicesSummary) GetServiceCountsOk() (*map[string]int32, bool)`

GetServiceCountsOk returns a tuple with the ServiceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCounts

`func (o *AccountServicesSummary) SetServiceCounts(v map[string]int32)`

SetServiceCounts sets ServiceCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


