# Company

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

## Methods

### NewCompany

`func NewCompany(name string, mainContact string, website string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.


### GetMainContact

`func (o *Company) GetMainContact() string`

GetMainContact returns the MainContact field if non-nil, zero value otherwise.

### GetMainContactOk

`func (o *Company) GetMainContactOk() (*string, bool)`

GetMainContactOk returns a tuple with the MainContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContact

`func (o *Company) SetMainContact(v string)`

SetMainContact sets MainContact field to given value.


### GetWebsite

`func (o *Company) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Company) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Company) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetEmployeeCount

`func (o *Company) GetEmployeeCount() int32`

GetEmployeeCount returns the EmployeeCount field if non-nil, zero value otherwise.

### GetEmployeeCountOk

`func (o *Company) GetEmployeeCountOk() (*int32, bool)`

GetEmployeeCountOk returns a tuple with the EmployeeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCount

`func (o *Company) SetEmployeeCount(v int32)`

SetEmployeeCount sets EmployeeCount field to given value.

### HasEmployeeCount

`func (o *Company) HasEmployeeCount() bool`

HasEmployeeCount returns a boolean if a field has been set.

### GetEmployeeCountIt

`func (o *Company) GetEmployeeCountIt() int32`

GetEmployeeCountIt returns the EmployeeCountIt field if non-nil, zero value otherwise.

### GetEmployeeCountItOk

`func (o *Company) GetEmployeeCountItOk() (*int32, bool)`

GetEmployeeCountItOk returns a tuple with the EmployeeCountIt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCountIt

`func (o *Company) SetEmployeeCountIt(v int32)`

SetEmployeeCountIt sets EmployeeCountIt field to given value.

### HasEmployeeCountIt

`func (o *Company) HasEmployeeCountIt() bool`

HasEmployeeCountIt returns a boolean if a field has been set.

### GetSector

`func (o *Company) GetSector() CompanySector`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *Company) GetSectorOk() (*CompanySector, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *Company) SetSector(v CompanySector)`

SetSector sets Sector field to given value.

### HasSector

`func (o *Company) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetVulnScanFrequency

`func (o *Company) GetVulnScanFrequency() ScanFrequency`

GetVulnScanFrequency returns the VulnScanFrequency field if non-nil, zero value otherwise.

### GetVulnScanFrequencyOk

`func (o *Company) GetVulnScanFrequencyOk() (*ScanFrequency, bool)`

GetVulnScanFrequencyOk returns a tuple with the VulnScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnScanFrequency

`func (o *Company) SetVulnScanFrequency(v ScanFrequency)`

SetVulnScanFrequency sets VulnScanFrequency field to given value.

### HasVulnScanFrequency

`func (o *Company) HasVulnScanFrequency() bool`

HasVulnScanFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


