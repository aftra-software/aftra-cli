# IntegrationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallationId** | **string** |  | 
**IntegrationType** | [**IntegrationTypeEnum**](IntegrationTypeEnum.md) |  | 
**User** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] [default to ""]
**Pk** | **string** |  | 
**Sk** | **string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIntegrationEntity

`func NewIntegrationEntity(installationId string, integrationType IntegrationTypeEnum, user string, pk string, sk string, ) *IntegrationEntity`

NewIntegrationEntity instantiates a new IntegrationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEntityWithDefaults

`func NewIntegrationEntityWithDefaults() *IntegrationEntity`

NewIntegrationEntityWithDefaults instantiates a new IntegrationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallationId

`func (o *IntegrationEntity) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *IntegrationEntity) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *IntegrationEntity) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetIntegrationType

`func (o *IntegrationEntity) GetIntegrationType() IntegrationTypeEnum`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *IntegrationEntity) GetIntegrationTypeOk() (*IntegrationTypeEnum, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *IntegrationEntity) SetIntegrationType(v IntegrationTypeEnum)`

SetIntegrationType sets IntegrationType field to given value.


### GetUser

`func (o *IntegrationEntity) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IntegrationEntity) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IntegrationEntity) SetUser(v string)`

SetUser sets User field to given value.


### GetName

`func (o *IntegrationEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPk

`func (o *IntegrationEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *IntegrationEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *IntegrationEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *IntegrationEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *IntegrationEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *IntegrationEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *IntegrationEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *IntegrationEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *IntegrationEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *IntegrationEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *IntegrationEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IntegrationEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IntegrationEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IntegrationEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *IntegrationEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *IntegrationEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *IntegrationEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *IntegrationEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


