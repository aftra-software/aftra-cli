# AggregatedSaasEnumerationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "aggregated-saas-enumeration"]
**Services** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewAggregatedSaasEnumerationReport

`func NewAggregatedSaasEnumerationReport(pk string, ) *AggregatedSaasEnumerationReport`

NewAggregatedSaasEnumerationReport instantiates a new AggregatedSaasEnumerationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedSaasEnumerationReportWithDefaults

`func NewAggregatedSaasEnumerationReportWithDefaults() *AggregatedSaasEnumerationReport`

NewAggregatedSaasEnumerationReportWithDefaults instantiates a new AggregatedSaasEnumerationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AggregatedSaasEnumerationReport) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AggregatedSaasEnumerationReport) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AggregatedSaasEnumerationReport) SetPk(v string)`

SetPk sets Pk field to given value.


### GetUpdated

`func (o *AggregatedSaasEnumerationReport) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AggregatedSaasEnumerationReport) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AggregatedSaasEnumerationReport) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AggregatedSaasEnumerationReport) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetType

`func (o *AggregatedSaasEnumerationReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AggregatedSaasEnumerationReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AggregatedSaasEnumerationReport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AggregatedSaasEnumerationReport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServices

`func (o *AggregatedSaasEnumerationReport) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AggregatedSaasEnumerationReport) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AggregatedSaasEnumerationReport) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *AggregatedSaasEnumerationReport) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


