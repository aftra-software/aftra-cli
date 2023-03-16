# RepositoryEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Installation** | **string** |  | 
**Pk** | **string** |  | 
**Sk** | **string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRepositoryEntity

`func NewRepositoryEntity(name string, installation string, pk string, sk string, ) *RepositoryEntity`

NewRepositoryEntity instantiates a new RepositoryEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryEntityWithDefaults

`func NewRepositoryEntityWithDefaults() *RepositoryEntity`

NewRepositoryEntityWithDefaults instantiates a new RepositoryEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryEntity) SetName(v string)`

SetName sets Name field to given value.


### GetInstallation

`func (o *RepositoryEntity) GetInstallation() string`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *RepositoryEntity) GetInstallationOk() (*string, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *RepositoryEntity) SetInstallation(v string)`

SetInstallation sets Installation field to given value.


### GetPk

`func (o *RepositoryEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *RepositoryEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *RepositoryEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *RepositoryEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *RepositoryEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *RepositoryEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *RepositoryEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *RepositoryEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *RepositoryEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *RepositoryEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *RepositoryEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RepositoryEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RepositoryEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RepositoryEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RepositoryEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RepositoryEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RepositoryEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RepositoryEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetMetadata

`func (o *RepositoryEntity) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RepositoryEntity) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RepositoryEntity) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RepositoryEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


