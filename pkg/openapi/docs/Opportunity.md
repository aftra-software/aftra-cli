# Opportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | **string** |  | 
**Type** | [**ModelType**](ModelType.md) |  | 
**Group** | [**OpportunityGroup**](OpportunityGroup.md) |  | 
**Entity** | Pointer to **string** |  | [optional] 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewOpportunity

`func NewOpportunity(company string, uid string, type_ ModelType, group OpportunityGroup, ) *Opportunity`

NewOpportunity instantiates a new Opportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityWithDefaults

`func NewOpportunityWithDefaults() *Opportunity`

NewOpportunityWithDefaults instantiates a new Opportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *Opportunity) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Opportunity) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Opportunity) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *Opportunity) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Opportunity) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Opportunity) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *Opportunity) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Opportunity) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Opportunity) SetType(v ModelType)`

SetType sets Type field to given value.


### GetGroup

`func (o *Opportunity) GetGroup() OpportunityGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Opportunity) GetGroupOk() (*OpportunityGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Opportunity) SetGroup(v OpportunityGroup)`

SetGroup sets Group field to given value.


### GetEntity

`func (o *Opportunity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Opportunity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Opportunity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Opportunity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetScore

`func (o *Opportunity) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Opportunity) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Opportunity) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *Opportunity) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *Opportunity) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Opportunity) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Opportunity) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Opportunity) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *Opportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Opportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Opportunity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Opportunity) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


