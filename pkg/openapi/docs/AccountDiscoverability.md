# AccountDiscoverability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | Pointer to **string** |  | [optional] [default to "account_discoverability"]
**Type** | Pointer to **string** |  | [optional] [default to "account_discoverability"]
**Group** | Pointer to **string** |  | [optional] [default to "human_exposure"]
**Entity** | Pointer to **string** |  | [optional] 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountsCount** | **int32** |  | 
**AccountsRatio** | **float32** |  | 
**AccountsBySource** | **map[string]int32** |  | 

## Methods

### NewAccountDiscoverability

`func NewAccountDiscoverability(company string, accountsCount int32, accountsRatio float32, accountsBySource map[string]int32, ) *AccountDiscoverability`

NewAccountDiscoverability instantiates a new AccountDiscoverability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDiscoverabilityWithDefaults

`func NewAccountDiscoverabilityWithDefaults() *AccountDiscoverability`

NewAccountDiscoverabilityWithDefaults instantiates a new AccountDiscoverability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *AccountDiscoverability) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AccountDiscoverability) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AccountDiscoverability) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *AccountDiscoverability) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AccountDiscoverability) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AccountDiscoverability) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AccountDiscoverability) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *AccountDiscoverability) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountDiscoverability) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountDiscoverability) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountDiscoverability) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *AccountDiscoverability) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AccountDiscoverability) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AccountDiscoverability) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AccountDiscoverability) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEntity

`func (o *AccountDiscoverability) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AccountDiscoverability) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AccountDiscoverability) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *AccountDiscoverability) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetScore

`func (o *AccountDiscoverability) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AccountDiscoverability) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AccountDiscoverability) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *AccountDiscoverability) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *AccountDiscoverability) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *AccountDiscoverability) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *AccountDiscoverability) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *AccountDiscoverability) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *AccountDiscoverability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDiscoverability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDiscoverability) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountDiscoverability) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountsCount

`func (o *AccountDiscoverability) GetAccountsCount() int32`

GetAccountsCount returns the AccountsCount field if non-nil, zero value otherwise.

### GetAccountsCountOk

`func (o *AccountDiscoverability) GetAccountsCountOk() (*int32, bool)`

GetAccountsCountOk returns a tuple with the AccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsCount

`func (o *AccountDiscoverability) SetAccountsCount(v int32)`

SetAccountsCount sets AccountsCount field to given value.


### GetAccountsRatio

`func (o *AccountDiscoverability) GetAccountsRatio() float32`

GetAccountsRatio returns the AccountsRatio field if non-nil, zero value otherwise.

### GetAccountsRatioOk

`func (o *AccountDiscoverability) GetAccountsRatioOk() (*float32, bool)`

GetAccountsRatioOk returns a tuple with the AccountsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsRatio

`func (o *AccountDiscoverability) SetAccountsRatio(v float32)`

SetAccountsRatio sets AccountsRatio field to given value.


### GetAccountsBySource

`func (o *AccountDiscoverability) GetAccountsBySource() map[string]int32`

GetAccountsBySource returns the AccountsBySource field if non-nil, zero value otherwise.

### GetAccountsBySourceOk

`func (o *AccountDiscoverability) GetAccountsBySourceOk() (*map[string]int32, bool)`

GetAccountsBySourceOk returns a tuple with the AccountsBySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsBySource

`func (o *AccountDiscoverability) SetAccountsBySource(v map[string]int32)`

SetAccountsBySource sets AccountsBySource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


