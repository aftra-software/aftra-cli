# UserCompanyAccessEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Sk** | **string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**AccessType** | [**CompanyAccessEnum**](CompanyAccessEnum.md) |  | 
**Email** | **string** |  | 

## Methods

### NewUserCompanyAccessEntity

`func NewUserCompanyAccessEntity(pk string, sk string, name string, accessType CompanyAccessEnum, email string, ) *UserCompanyAccessEntity`

NewUserCompanyAccessEntity instantiates a new UserCompanyAccessEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCompanyAccessEntityWithDefaults

`func NewUserCompanyAccessEntityWithDefaults() *UserCompanyAccessEntity`

NewUserCompanyAccessEntityWithDefaults instantiates a new UserCompanyAccessEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserCompanyAccessEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserCompanyAccessEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserCompanyAccessEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *UserCompanyAccessEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *UserCompanyAccessEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *UserCompanyAccessEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *UserCompanyAccessEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UserCompanyAccessEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UserCompanyAccessEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *UserCompanyAccessEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *UserCompanyAccessEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserCompanyAccessEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserCompanyAccessEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserCompanyAccessEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *UserCompanyAccessEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UserCompanyAccessEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UserCompanyAccessEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UserCompanyAccessEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetName

`func (o *UserCompanyAccessEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCompanyAccessEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCompanyAccessEntity) SetName(v string)`

SetName sets Name field to given value.


### GetAccessType

`func (o *UserCompanyAccessEntity) GetAccessType() CompanyAccessEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UserCompanyAccessEntity) GetAccessTypeOk() (*CompanyAccessEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UserCompanyAccessEntity) SetAccessType(v CompanyAccessEnum)`

SetAccessType sets AccessType field to given value.


### GetEmail

`func (o *UserCompanyAccessEntity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCompanyAccessEntity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCompanyAccessEntity) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


