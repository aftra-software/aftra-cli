# InternalOpportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "internal"]
**Group** | Pointer to **string** |  | [optional] [default to "internal"]
**Entity** | Pointer to **string** |  | [optional] 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | **string** |  | 
**Details** | **map[string]string** |  | 

## Methods

### NewInternalOpportunity

`func NewInternalOpportunity(company string, uid string, name string, details map[string]string, ) *InternalOpportunity`

NewInternalOpportunity instantiates a new InternalOpportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalOpportunityWithDefaults

`func NewInternalOpportunityWithDefaults() *InternalOpportunity`

NewInternalOpportunityWithDefaults instantiates a new InternalOpportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *InternalOpportunity) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InternalOpportunity) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InternalOpportunity) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *InternalOpportunity) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *InternalOpportunity) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *InternalOpportunity) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *InternalOpportunity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalOpportunity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalOpportunity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InternalOpportunity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *InternalOpportunity) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InternalOpportunity) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InternalOpportunity) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InternalOpportunity) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEntity

`func (o *InternalOpportunity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *InternalOpportunity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *InternalOpportunity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *InternalOpportunity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetScore

`func (o *InternalOpportunity) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *InternalOpportunity) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *InternalOpportunity) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *InternalOpportunity) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *InternalOpportunity) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *InternalOpportunity) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *InternalOpportunity) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *InternalOpportunity) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *InternalOpportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalOpportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalOpportunity) SetName(v string)`

SetName sets Name field to given value.


### GetDetails

`func (o *InternalOpportunity) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InternalOpportunity) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InternalOpportunity) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


