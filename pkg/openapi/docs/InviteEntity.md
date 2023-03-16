# InviteEntity

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

## Methods

### NewInviteEntity

`func NewInviteEntity(pk string, sk string, name string, accessType CompanyAccessEnum, ) *InviteEntity`

NewInviteEntity instantiates a new InviteEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteEntityWithDefaults

`func NewInviteEntityWithDefaults() *InviteEntity`

NewInviteEntityWithDefaults instantiates a new InviteEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *InviteEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *InviteEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *InviteEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *InviteEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *InviteEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *InviteEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *InviteEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *InviteEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *InviteEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *InviteEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *InviteEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InviteEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InviteEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InviteEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *InviteEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *InviteEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *InviteEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *InviteEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetName

`func (o *InviteEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InviteEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InviteEntity) SetName(v string)`

SetName sets Name field to given value.


### GetAccessType

`func (o *InviteEntity) GetAccessType() CompanyAccessEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *InviteEntity) GetAccessTypeOk() (*CompanyAccessEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *InviteEntity) SetAccessType(v CompanyAccessEnum)`

SetAccessType sets AccessType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


