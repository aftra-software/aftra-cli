# CompanyWithUserAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | [**CompanyEntity**](CompanyEntity.md) |  | 
**AccessType** | [**CompanyAccessEnum**](CompanyAccessEnum.md) |  | 

## Methods

### NewCompanyWithUserAccess

`func NewCompanyWithUserAccess(company CompanyEntity, accessType CompanyAccessEnum, ) *CompanyWithUserAccess`

NewCompanyWithUserAccess instantiates a new CompanyWithUserAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithUserAccessWithDefaults

`func NewCompanyWithUserAccessWithDefaults() *CompanyWithUserAccess`

NewCompanyWithUserAccessWithDefaults instantiates a new CompanyWithUserAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *CompanyWithUserAccess) GetCompany() CompanyEntity`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyWithUserAccess) GetCompanyOk() (*CompanyEntity, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyWithUserAccess) SetCompany(v CompanyEntity)`

SetCompany sets Company field to given value.


### GetAccessType

`func (o *CompanyWithUserAccess) GetAccessType() CompanyAccessEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CompanyWithUserAccess) GetAccessTypeOk() (*CompanyAccessEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CompanyWithUserAccess) SetAccessType(v CompanyAccessEnum)`

SetAccessType sets AccessType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


