# ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost

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
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**TxtRecord** | Pointer to **string** |  | [optional] 
**SourceInformation** | [**Sourceinformation**](Sourceinformation.md) |  | 
**FullName** | Pointer to **string** |  | [optional] [default to ""]
**Status** | Pointer to [**AccountStatusEnum**](AccountStatusEnum.md) |  | [optional] 
**CampaignOnNextScan** | Pointer to **bool** |  | [optional] [default to false]
**Verified** | Pointer to **bool** |  | [optional] [default to false]
**Installation** | **string** |  | 

## Methods

### NewResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost

`func NewResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost(name string, pk string, sk string, sourceInformation Sourceinformation, installation string, ) *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost`

NewResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost instantiates a new ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPostWithDefaults

`func NewResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPostWithDefaults() *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost`

NewResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPostWithDefaults instantiates a new ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetName(v string)`

SetName sets Name field to given value.


### GetActiveScanFrequency

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetActiveScanFrequency() ScanFrequency`

GetActiveScanFrequency returns the ActiveScanFrequency field if non-nil, zero value otherwise.

### GetActiveScanFrequencyOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetActiveScanFrequencyOk() (*ScanFrequency, bool)`

GetActiveScanFrequencyOk returns a tuple with the ActiveScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveScanFrequency

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetActiveScanFrequency(v ScanFrequency)`

SetActiveScanFrequency sets ActiveScanFrequency field to given value.

### HasActiveScanFrequency

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasActiveScanFrequency() bool`

HasActiveScanFrequency returns a boolean if a field has been set.

### GetPk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetDnsVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetDnsVerified() bool`

GetDnsVerified returns the DnsVerified field if non-nil, zero value otherwise.

### GetDnsVerifiedOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetDnsVerifiedOk() (*bool, bool)`

GetDnsVerifiedOk returns a tuple with the DnsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetDnsVerified(v bool)`

SetDnsVerified sets DnsVerified field to given value.

### HasDnsVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasDnsVerified() bool`

HasDnsVerified returns a boolean if a field has been set.

### GetAdminVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetAdminVerified() bool`

GetAdminVerified returns the AdminVerified field if non-nil, zero value otherwise.

### GetAdminVerifiedOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetAdminVerifiedOk() (*bool, bool)`

GetAdminVerifiedOk returns a tuple with the AdminVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetAdminVerified(v bool)`

SetAdminVerified sets AdminVerified field to given value.

### HasAdminVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasAdminVerified() bool`

HasAdminVerified returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTxtRecord

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetTxtRecord() string`

GetTxtRecord returns the TxtRecord field if non-nil, zero value otherwise.

### GetTxtRecordOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetTxtRecordOk() (*string, bool)`

GetTxtRecordOk returns a tuple with the TxtRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtRecord

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetTxtRecord(v string)`

SetTxtRecord sets TxtRecord field to given value.

### HasTxtRecord

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasTxtRecord() bool`

HasTxtRecord returns a boolean if a field has been set.

### GetSourceInformation

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetSourceInformation() Sourceinformation`

GetSourceInformation returns the SourceInformation field if non-nil, zero value otherwise.

### GetSourceInformationOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetSourceInformationOk() (*Sourceinformation, bool)`

GetSourceInformationOk returns a tuple with the SourceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInformation

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetSourceInformation(v Sourceinformation)`

SetSourceInformation sets SourceInformation field to given value.


### GetFullName

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCampaignOnNextScan

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetCampaignOnNextScan() bool`

GetCampaignOnNextScan returns the CampaignOnNextScan field if non-nil, zero value otherwise.

### GetCampaignOnNextScanOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetCampaignOnNextScanOk() (*bool, bool)`

GetCampaignOnNextScanOk returns a tuple with the CampaignOnNextScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignOnNextScan

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetCampaignOnNextScan(v bool)`

SetCampaignOnNextScan sets CampaignOnNextScan field to given value.

### HasCampaignOnNextScan

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasCampaignOnNextScan() bool`

HasCampaignOnNextScan returns a boolean if a field has been set.

### GetVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetInstallation

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetInstallation() string`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) GetInstallationOk() (*string, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *ResponseAcceptSuggestionApiCompaniesCompanyPkSuggestionsSuggestionIdPost) SetInstallation(v string)`

SetInstallation sets Installation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


