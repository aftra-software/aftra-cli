# SubdomainEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ActiveScanFrequency** | Pointer to [**ScanFrequency**](ScanFrequency.md) |  | [optional] 
**Pk** | **string** |  | 
**Sk** | **string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**SourceInformation** | [**Sourceinformation**](Sourceinformation.md) |  | 

## Methods

### NewSubdomainEntity

`func NewSubdomainEntity(name string, pk string, sk string, sourceInformation Sourceinformation, ) *SubdomainEntity`

NewSubdomainEntity instantiates a new SubdomainEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubdomainEntityWithDefaults

`func NewSubdomainEntityWithDefaults() *SubdomainEntity`

NewSubdomainEntityWithDefaults instantiates a new SubdomainEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubdomainEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubdomainEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubdomainEntity) SetName(v string)`

SetName sets Name field to given value.


### GetActiveScanFrequency

`func (o *SubdomainEntity) GetActiveScanFrequency() ScanFrequency`

GetActiveScanFrequency returns the ActiveScanFrequency field if non-nil, zero value otherwise.

### GetActiveScanFrequencyOk

`func (o *SubdomainEntity) GetActiveScanFrequencyOk() (*ScanFrequency, bool)`

GetActiveScanFrequencyOk returns a tuple with the ActiveScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveScanFrequency

`func (o *SubdomainEntity) SetActiveScanFrequency(v ScanFrequency)`

SetActiveScanFrequency sets ActiveScanFrequency field to given value.

### HasActiveScanFrequency

`func (o *SubdomainEntity) HasActiveScanFrequency() bool`

HasActiveScanFrequency returns a boolean if a field has been set.

### GetPk

`func (o *SubdomainEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SubdomainEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SubdomainEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *SubdomainEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *SubdomainEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *SubdomainEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *SubdomainEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SubdomainEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SubdomainEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *SubdomainEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *SubdomainEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubdomainEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubdomainEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SubdomainEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *SubdomainEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SubdomainEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SubdomainEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SubdomainEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetMetadata

`func (o *SubdomainEntity) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubdomainEntity) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubdomainEntity) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubdomainEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSourceInformation

`func (o *SubdomainEntity) GetSourceInformation() Sourceinformation`

GetSourceInformation returns the SourceInformation field if non-nil, zero value otherwise.

### GetSourceInformationOk

`func (o *SubdomainEntity) GetSourceInformationOk() (*Sourceinformation, bool)`

GetSourceInformationOk returns a tuple with the SourceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInformation

`func (o *SubdomainEntity) SetSourceInformation(v Sourceinformation)`

SetSourceInformation sets SourceInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


