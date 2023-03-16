# PaginatedEntityCollectionAccountEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]AccountEntity**](AccountEntity.md) |  | 
**Marker** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginatedEntityCollectionAccountEntity

`func NewPaginatedEntityCollectionAccountEntity(entities []AccountEntity, ) *PaginatedEntityCollectionAccountEntity`

NewPaginatedEntityCollectionAccountEntity instantiates a new PaginatedEntityCollectionAccountEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEntityCollectionAccountEntityWithDefaults

`func NewPaginatedEntityCollectionAccountEntityWithDefaults() *PaginatedEntityCollectionAccountEntity`

NewPaginatedEntityCollectionAccountEntityWithDefaults instantiates a new PaginatedEntityCollectionAccountEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *PaginatedEntityCollectionAccountEntity) GetEntities() []AccountEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PaginatedEntityCollectionAccountEntity) GetEntitiesOk() (*[]AccountEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PaginatedEntityCollectionAccountEntity) SetEntities(v []AccountEntity)`

SetEntities sets Entities field to given value.


### GetMarker

`func (o *PaginatedEntityCollectionAccountEntity) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *PaginatedEntityCollectionAccountEntity) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *PaginatedEntityCollectionAccountEntity) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *PaginatedEntityCollectionAccountEntity) HasMarker() bool`

HasMarker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


