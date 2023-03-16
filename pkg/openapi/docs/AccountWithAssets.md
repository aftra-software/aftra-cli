# AccountWithAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**PaginatedEntityCollectionMetaEntity**](PaginatedEntityCollectionMetaEntity.md) |  | 
**Entity** | [**AccountEntity**](AccountEntity.md) |  | 

## Methods

### NewAccountWithAssets

`func NewAccountWithAssets(assets PaginatedEntityCollectionMetaEntity, entity AccountEntity, ) *AccountWithAssets`

NewAccountWithAssets instantiates a new AccountWithAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithAssetsWithDefaults

`func NewAccountWithAssetsWithDefaults() *AccountWithAssets`

NewAccountWithAssetsWithDefaults instantiates a new AccountWithAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *AccountWithAssets) GetAssets() PaginatedEntityCollectionMetaEntity`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *AccountWithAssets) GetAssetsOk() (*PaginatedEntityCollectionMetaEntity, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *AccountWithAssets) SetAssets(v PaginatedEntityCollectionMetaEntity)`

SetAssets sets Assets field to given value.


### GetEntity

`func (o *AccountWithAssets) GetEntity() AccountEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AccountWithAssets) GetEntityOk() (*AccountEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AccountWithAssets) SetEntity(v AccountEntity)`

SetEntity sets Entity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


