# SuspiciousDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "suspicious_domain"]
**Group** | Pointer to **string** |  | [optional] [default to "domains"]
**Entity** | **string** |  | 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | **string** |  | 
**ScoreDetails** | [**map[string][]IndividualScore**](array.md) |  | 

## Methods

### NewSuspiciousDomain

`func NewSuspiciousDomain(company string, uid string, entity string, name string, scoreDetails map[string][]IndividualScore, ) *SuspiciousDomain`

NewSuspiciousDomain instantiates a new SuspiciousDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspiciousDomainWithDefaults

`func NewSuspiciousDomainWithDefaults() *SuspiciousDomain`

NewSuspiciousDomainWithDefaults instantiates a new SuspiciousDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *SuspiciousDomain) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SuspiciousDomain) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SuspiciousDomain) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *SuspiciousDomain) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SuspiciousDomain) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SuspiciousDomain) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *SuspiciousDomain) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SuspiciousDomain) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SuspiciousDomain) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SuspiciousDomain) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *SuspiciousDomain) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SuspiciousDomain) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SuspiciousDomain) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SuspiciousDomain) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEntity

`func (o *SuspiciousDomain) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SuspiciousDomain) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SuspiciousDomain) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetScore

`func (o *SuspiciousDomain) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SuspiciousDomain) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SuspiciousDomain) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *SuspiciousDomain) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *SuspiciousDomain) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *SuspiciousDomain) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *SuspiciousDomain) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *SuspiciousDomain) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *SuspiciousDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuspiciousDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuspiciousDomain) SetName(v string)`

SetName sets Name field to given value.


### GetScoreDetails

`func (o *SuspiciousDomain) GetScoreDetails() map[string][]IndividualScore`

GetScoreDetails returns the ScoreDetails field if non-nil, zero value otherwise.

### GetScoreDetailsOk

`func (o *SuspiciousDomain) GetScoreDetailsOk() (*map[string][]IndividualScore, bool)`

GetScoreDetailsOk returns a tuple with the ScoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreDetails

`func (o *SuspiciousDomain) SetScoreDetails(v map[string][]IndividualScore)`

SetScoreDetails sets ScoreDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


