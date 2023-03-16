# AccountEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Sk** | **string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**FullName** | Pointer to **string** |  | [optional] [default to ""]
**Status** | Pointer to [**AccountStatusEnum**](AccountStatusEnum.md) |  | [optional] 
**CampaignOnNextScan** | Pointer to **bool** |  | [optional] [default to false]
**Verified** | Pointer to **bool** |  | [optional] [default to false]
**SourceInformation** | [**Sourceinformation**](Sourceinformation.md) |  | 
**Metadata** | Pointer to [**AccountMetadata**](AccountMetadata.md) |  | [optional] 

## Methods

### NewAccountEntity

`func NewAccountEntity(pk string, sk string, name string, sourceInformation Sourceinformation, ) *AccountEntity`

NewAccountEntity instantiates a new AccountEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountEntityWithDefaults

`func NewAccountEntityWithDefaults() *AccountEntity`

NewAccountEntityWithDefaults instantiates a new AccountEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AccountEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AccountEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AccountEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *AccountEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *AccountEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *AccountEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *AccountEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AccountEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AccountEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AccountEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *AccountEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AccountEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AccountEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AccountEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AccountEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetName

`func (o *AccountEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountEntity) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *AccountEntity) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AccountEntity) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AccountEntity) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AccountEntity) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetStatus

`func (o *AccountEntity) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountEntity) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountEntity) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCampaignOnNextScan

`func (o *AccountEntity) GetCampaignOnNextScan() bool`

GetCampaignOnNextScan returns the CampaignOnNextScan field if non-nil, zero value otherwise.

### GetCampaignOnNextScanOk

`func (o *AccountEntity) GetCampaignOnNextScanOk() (*bool, bool)`

GetCampaignOnNextScanOk returns a tuple with the CampaignOnNextScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignOnNextScan

`func (o *AccountEntity) SetCampaignOnNextScan(v bool)`

SetCampaignOnNextScan sets CampaignOnNextScan field to given value.

### HasCampaignOnNextScan

`func (o *AccountEntity) HasCampaignOnNextScan() bool`

HasCampaignOnNextScan returns a boolean if a field has been set.

### GetVerified

`func (o *AccountEntity) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AccountEntity) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AccountEntity) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AccountEntity) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetSourceInformation

`func (o *AccountEntity) GetSourceInformation() Sourceinformation`

GetSourceInformation returns the SourceInformation field if non-nil, zero value otherwise.

### GetSourceInformationOk

`func (o *AccountEntity) GetSourceInformationOk() (*Sourceinformation, bool)`

GetSourceInformationOk returns a tuple with the SourceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInformation

`func (o *AccountEntity) SetSourceInformation(v Sourceinformation)`

SetSourceInformation sets SourceInformation field to given value.


### GetMetadata

`func (o *AccountEntity) GetMetadata() AccountMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountEntity) GetMetadataOk() (*AccountMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountEntity) SetMetadata(v AccountMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


