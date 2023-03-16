# CompanyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainContact** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**EmployeeCount** | Pointer to **int32** |  | [optional] 
**EmployeeCountIt** | Pointer to **int32** |  | [optional] 
**Sector** | Pointer to [**CompanySector**](CompanySector.md) |  | [optional] 
**VulnScanFrequency** | Pointer to [**ScanFrequency**](ScanFrequency.md) |  | [optional] 

## Methods

### NewCompanyUpdate

`func NewCompanyUpdate() *CompanyUpdate`

NewCompanyUpdate instantiates a new CompanyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyUpdateWithDefaults

`func NewCompanyUpdateWithDefaults() *CompanyUpdate`

NewCompanyUpdateWithDefaults instantiates a new CompanyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMainContact

`func (o *CompanyUpdate) GetMainContact() string`

GetMainContact returns the MainContact field if non-nil, zero value otherwise.

### GetMainContactOk

`func (o *CompanyUpdate) GetMainContactOk() (*string, bool)`

GetMainContactOk returns a tuple with the MainContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContact

`func (o *CompanyUpdate) SetMainContact(v string)`

SetMainContact sets MainContact field to given value.

### HasMainContact

`func (o *CompanyUpdate) HasMainContact() bool`

HasMainContact returns a boolean if a field has been set.

### GetWebsite

`func (o *CompanyUpdate) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CompanyUpdate) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CompanyUpdate) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CompanyUpdate) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmployeeCount

`func (o *CompanyUpdate) GetEmployeeCount() int32`

GetEmployeeCount returns the EmployeeCount field if non-nil, zero value otherwise.

### GetEmployeeCountOk

`func (o *CompanyUpdate) GetEmployeeCountOk() (*int32, bool)`

GetEmployeeCountOk returns a tuple with the EmployeeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCount

`func (o *CompanyUpdate) SetEmployeeCount(v int32)`

SetEmployeeCount sets EmployeeCount field to given value.

### HasEmployeeCount

`func (o *CompanyUpdate) HasEmployeeCount() bool`

HasEmployeeCount returns a boolean if a field has been set.

### GetEmployeeCountIt

`func (o *CompanyUpdate) GetEmployeeCountIt() int32`

GetEmployeeCountIt returns the EmployeeCountIt field if non-nil, zero value otherwise.

### GetEmployeeCountItOk

`func (o *CompanyUpdate) GetEmployeeCountItOk() (*int32, bool)`

GetEmployeeCountItOk returns a tuple with the EmployeeCountIt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCountIt

`func (o *CompanyUpdate) SetEmployeeCountIt(v int32)`

SetEmployeeCountIt sets EmployeeCountIt field to given value.

### HasEmployeeCountIt

`func (o *CompanyUpdate) HasEmployeeCountIt() bool`

HasEmployeeCountIt returns a boolean if a field has been set.

### GetSector

`func (o *CompanyUpdate) GetSector() CompanySector`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *CompanyUpdate) GetSectorOk() (*CompanySector, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *CompanyUpdate) SetSector(v CompanySector)`

SetSector sets Sector field to given value.

### HasSector

`func (o *CompanyUpdate) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetVulnScanFrequency

`func (o *CompanyUpdate) GetVulnScanFrequency() ScanFrequency`

GetVulnScanFrequency returns the VulnScanFrequency field if non-nil, zero value otherwise.

### GetVulnScanFrequencyOk

`func (o *CompanyUpdate) GetVulnScanFrequencyOk() (*ScanFrequency, bool)`

GetVulnScanFrequencyOk returns a tuple with the VulnScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnScanFrequency

`func (o *CompanyUpdate) SetVulnScanFrequency(v ScanFrequency)`

SetVulnScanFrequency sets VulnScanFrequency field to given value.

### HasVulnScanFrequency

`func (o *CompanyUpdate) HasVulnScanFrequency() bool`

HasVulnScanFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


