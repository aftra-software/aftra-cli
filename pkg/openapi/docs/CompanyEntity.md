# CompanyEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MainContact** | **string** |  | 
**Website** | **string** |  | 
**EmployeeCount** | Pointer to **int32** |  | [optional] [default to 1]
**EmployeeCountIt** | Pointer to **int32** |  | [optional] [default to 1]
**Sector** | Pointer to [**CompanySector**](CompanySector.md) |  | [optional] 
**VulnScanFrequency** | Pointer to [**ScanFrequency**](ScanFrequency.md) |  | [optional] 
**Pk** | **string** |  | 
**Sk** | **string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCompanyEntity

`func NewCompanyEntity(name string, mainContact string, website string, pk string, sk string, ) *CompanyEntity`

NewCompanyEntity instantiates a new CompanyEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyEntityWithDefaults

`func NewCompanyEntityWithDefaults() *CompanyEntity`

NewCompanyEntityWithDefaults instantiates a new CompanyEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompanyEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyEntity) SetName(v string)`

SetName sets Name field to given value.


### GetMainContact

`func (o *CompanyEntity) GetMainContact() string`

GetMainContact returns the MainContact field if non-nil, zero value otherwise.

### GetMainContactOk

`func (o *CompanyEntity) GetMainContactOk() (*string, bool)`

GetMainContactOk returns a tuple with the MainContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContact

`func (o *CompanyEntity) SetMainContact(v string)`

SetMainContact sets MainContact field to given value.


### GetWebsite

`func (o *CompanyEntity) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CompanyEntity) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CompanyEntity) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetEmployeeCount

`func (o *CompanyEntity) GetEmployeeCount() int32`

GetEmployeeCount returns the EmployeeCount field if non-nil, zero value otherwise.

### GetEmployeeCountOk

`func (o *CompanyEntity) GetEmployeeCountOk() (*int32, bool)`

GetEmployeeCountOk returns a tuple with the EmployeeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCount

`func (o *CompanyEntity) SetEmployeeCount(v int32)`

SetEmployeeCount sets EmployeeCount field to given value.

### HasEmployeeCount

`func (o *CompanyEntity) HasEmployeeCount() bool`

HasEmployeeCount returns a boolean if a field has been set.

### GetEmployeeCountIt

`func (o *CompanyEntity) GetEmployeeCountIt() int32`

GetEmployeeCountIt returns the EmployeeCountIt field if non-nil, zero value otherwise.

### GetEmployeeCountItOk

`func (o *CompanyEntity) GetEmployeeCountItOk() (*int32, bool)`

GetEmployeeCountItOk returns a tuple with the EmployeeCountIt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCountIt

`func (o *CompanyEntity) SetEmployeeCountIt(v int32)`

SetEmployeeCountIt sets EmployeeCountIt field to given value.

### HasEmployeeCountIt

`func (o *CompanyEntity) HasEmployeeCountIt() bool`

HasEmployeeCountIt returns a boolean if a field has been set.

### GetSector

`func (o *CompanyEntity) GetSector() CompanySector`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *CompanyEntity) GetSectorOk() (*CompanySector, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *CompanyEntity) SetSector(v CompanySector)`

SetSector sets Sector field to given value.

### HasSector

`func (o *CompanyEntity) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetVulnScanFrequency

`func (o *CompanyEntity) GetVulnScanFrequency() ScanFrequency`

GetVulnScanFrequency returns the VulnScanFrequency field if non-nil, zero value otherwise.

### GetVulnScanFrequencyOk

`func (o *CompanyEntity) GetVulnScanFrequencyOk() (*ScanFrequency, bool)`

GetVulnScanFrequencyOk returns a tuple with the VulnScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnScanFrequency

`func (o *CompanyEntity) SetVulnScanFrequency(v ScanFrequency)`

SetVulnScanFrequency sets VulnScanFrequency field to given value.

### HasVulnScanFrequency

`func (o *CompanyEntity) HasVulnScanFrequency() bool`

HasVulnScanFrequency returns a boolean if a field has been set.

### GetPk

`func (o *CompanyEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *CompanyEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *CompanyEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *CompanyEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *CompanyEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *CompanyEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *CompanyEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CompanyEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CompanyEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CompanyEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *CompanyEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CompanyEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CompanyEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CompanyEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *CompanyEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CompanyEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CompanyEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CompanyEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetMetadata

`func (o *CompanyEntity) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CompanyEntity) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CompanyEntity) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CompanyEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


