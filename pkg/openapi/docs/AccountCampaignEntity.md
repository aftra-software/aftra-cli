# AccountCampaignEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Sk** | **string** |  | 
**EntityType** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Services** | **[]string** |  | 
**RiskScore** | **float32** |  | 
**Conclusion** | Pointer to [**AccountCampaignConclusion**](AccountCampaignConclusion.md) |  | [optional] 

## Methods

### NewAccountCampaignEntity

`func NewAccountCampaignEntity(pk string, sk string, name string, services []string, riskScore float32, ) *AccountCampaignEntity`

NewAccountCampaignEntity instantiates a new AccountCampaignEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCampaignEntityWithDefaults

`func NewAccountCampaignEntityWithDefaults() *AccountCampaignEntity`

NewAccountCampaignEntityWithDefaults instantiates a new AccountCampaignEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AccountCampaignEntity) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AccountCampaignEntity) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AccountCampaignEntity) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSk

`func (o *AccountCampaignEntity) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *AccountCampaignEntity) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *AccountCampaignEntity) SetSk(v string)`

SetSk sets Sk field to given value.


### GetEntityType

`func (o *AccountCampaignEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AccountCampaignEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AccountCampaignEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AccountCampaignEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCreated

`func (o *AccountCampaignEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountCampaignEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountCampaignEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountCampaignEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AccountCampaignEntity) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AccountCampaignEntity) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AccountCampaignEntity) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AccountCampaignEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetName

`func (o *AccountCampaignEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCampaignEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCampaignEntity) SetName(v string)`

SetName sets Name field to given value.


### GetServices

`func (o *AccountCampaignEntity) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AccountCampaignEntity) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AccountCampaignEntity) SetServices(v []string)`

SetServices sets Services field to given value.


### GetRiskScore

`func (o *AccountCampaignEntity) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *AccountCampaignEntity) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *AccountCampaignEntity) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.


### GetConclusion

`func (o *AccountCampaignEntity) GetConclusion() AccountCampaignConclusion`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *AccountCampaignEntity) GetConclusionOk() (*AccountCampaignConclusion, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *AccountCampaignEntity) SetConclusion(v AccountCampaignConclusion)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *AccountCampaignEntity) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


