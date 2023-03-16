# LeakedPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "leaked_password"]
**Group** | Pointer to **string** |  | [optional] [default to "human_exposure"]
**Entity** | **string** |  | 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | **string** |  | 
**PasswordHint** | **string** |  | 
**LeakName** | **string** |  | 
**LeakYear** | **int32** |  | 

## Methods

### NewLeakedPassword

`func NewLeakedPassword(company string, uid string, entity string, name string, passwordHint string, leakName string, leakYear int32, ) *LeakedPassword`

NewLeakedPassword instantiates a new LeakedPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeakedPasswordWithDefaults

`func NewLeakedPasswordWithDefaults() *LeakedPassword`

NewLeakedPasswordWithDefaults instantiates a new LeakedPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *LeakedPassword) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *LeakedPassword) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *LeakedPassword) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *LeakedPassword) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *LeakedPassword) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *LeakedPassword) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *LeakedPassword) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LeakedPassword) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LeakedPassword) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LeakedPassword) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *LeakedPassword) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *LeakedPassword) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *LeakedPassword) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *LeakedPassword) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEntity

`func (o *LeakedPassword) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *LeakedPassword) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *LeakedPassword) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetScore

`func (o *LeakedPassword) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *LeakedPassword) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *LeakedPassword) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *LeakedPassword) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *LeakedPassword) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *LeakedPassword) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *LeakedPassword) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *LeakedPassword) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *LeakedPassword) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LeakedPassword) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LeakedPassword) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordHint

`func (o *LeakedPassword) GetPasswordHint() string`

GetPasswordHint returns the PasswordHint field if non-nil, zero value otherwise.

### GetPasswordHintOk

`func (o *LeakedPassword) GetPasswordHintOk() (*string, bool)`

GetPasswordHintOk returns a tuple with the PasswordHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHint

`func (o *LeakedPassword) SetPasswordHint(v string)`

SetPasswordHint sets PasswordHint field to given value.


### GetLeakName

`func (o *LeakedPassword) GetLeakName() string`

GetLeakName returns the LeakName field if non-nil, zero value otherwise.

### GetLeakNameOk

`func (o *LeakedPassword) GetLeakNameOk() (*string, bool)`

GetLeakNameOk returns a tuple with the LeakName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeakName

`func (o *LeakedPassword) SetLeakName(v string)`

SetLeakName sets LeakName field to given value.


### GetLeakYear

`func (o *LeakedPassword) GetLeakYear() int32`

GetLeakYear returns the LeakYear field if non-nil, zero value otherwise.

### GetLeakYearOk

`func (o *LeakedPassword) GetLeakYearOk() (*int32, bool)`

GetLeakYearOk returns a tuple with the LeakYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeakYear

`func (o *LeakedPassword) SetLeakYear(v int32)`

SetLeakYear sets LeakYear field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


