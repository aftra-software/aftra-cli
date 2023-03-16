# DomainEntity

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
**DnsVerified** | Pointer to **bool** |  | [optional] [default to false]
**AdminVerified** | Pointer to **bool** |  | [optional] [default to false]
**Metadata** | Pointer to [**DomainMetadata**](DomainMetadata.md) |  | [optional] 
**TxtRecord** | Pointer to **string** |  | [optional] 
**SourceInformation** | [**Sourceinformation**](Sourceinformation.md) |  | 

## Methods

### NewDomainEntity

`func NewDomainEntity(name string, pk string, sk string, sourceInformation Sourceinformation, ) *DomainEntity`

NewDomainEntity instantiates a new DomainEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainEntityWithDefaults

`func NewDomainEntityWithDefaults() *DomainEntity`

NewDomainEntityWithDefaults instantiates a new DomainEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainEntity) SetName(v string)`

SetName sets Name field to given value.


### GetActiveScanFrequency

`func (o *DomainEntity) GetActiveScanFrequency() ScanFrequency`

GetActiveScanFrequency returns the ActiveScanFrequency field if non-nil, zero value otherwise.

### GetActiveScanFrequencyOk

`func (o *DomainEntity) GetActiveScanFrequencyOk() (*ScanFrequency, bool)`

GetActiveScanFrequencyOk returns a tuple with the ActiveScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveScanFrequency

`func (o *DomainEntity) SetActiveScanFrequency(v ScanFrequency)`

SetActiveScanFrequency sets ActiveScanFrequency field to given value.

### HasActiveScanFrequency

`func (o *DomainEntity) HasActiveScanFrequency() bool`

HasActiveScanFrequency returns a boolean if a field has been set.

### GetPk

`func (o *DomainEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *DomainEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *DomainEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *DomainEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *DomainEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *DomainEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *DomainEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DomainEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DomainEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DomainEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *DomainEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DomainEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DomainEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DomainEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *DomainEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DomainEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DomainEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DomainEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetDnsVerified

`func (o *DomainEntity) GetDnsVerified() bool`

GetDnsVerified returns the DnsVerified field if non-nil, zero value otherwise.

### GetDnsVerifiedOk

`func (o *DomainEntity) GetDnsVerifiedOk() (*bool, bool)`

GetDnsVerifiedOk returns a tuple with the DnsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerified

`func (o *DomainEntity) SetDnsVerified(v bool)`

SetDnsVerified sets DnsVerified field to given value.

### HasDnsVerified

`func (o *DomainEntity) HasDnsVerified() bool`

HasDnsVerified returns a boolean if a field has been set.

### GetAdminVerified

`func (o *DomainEntity) GetAdminVerified() bool`

GetAdminVerified returns the AdminVerified field if non-nil, zero value otherwise.

### GetAdminVerifiedOk

`func (o *DomainEntity) GetAdminVerifiedOk() (*bool, bool)`

GetAdminVerifiedOk returns a tuple with the AdminVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminVerified

`func (o *DomainEntity) SetAdminVerified(v bool)`

SetAdminVerified sets AdminVerified field to given value.

### HasAdminVerified

`func (o *DomainEntity) HasAdminVerified() bool`

HasAdminVerified returns a boolean if a field has been set.

### GetMetadata

`func (o *DomainEntity) GetMetadata() DomainMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DomainEntity) GetMetadataOk() (*DomainMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DomainEntity) SetMetadata(v DomainMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DomainEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTxtRecord

`func (o *DomainEntity) GetTxtRecord() string`

GetTxtRecord returns the TxtRecord field if non-nil, zero value otherwise.

### GetTxtRecordOk

`func (o *DomainEntity) GetTxtRecordOk() (*string, bool)`

GetTxtRecordOk returns a tuple with the TxtRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtRecord

`func (o *DomainEntity) SetTxtRecord(v string)`

SetTxtRecord sets TxtRecord field to given value.

### HasTxtRecord

`func (o *DomainEntity) HasTxtRecord() bool`

HasTxtRecord returns a boolean if a field has been set.

### GetSourceInformation

`func (o *DomainEntity) GetSourceInformation() Sourceinformation`

GetSourceInformation returns the SourceInformation field if non-nil, zero value otherwise.

### GetSourceInformationOk

`func (o *DomainEntity) GetSourceInformationOk() (*Sourceinformation, bool)`

GetSourceInformationOk returns a tuple with the SourceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInformation

`func (o *DomainEntity) SetSourceInformation(v Sourceinformation)`

SetSourceInformation sets SourceInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


