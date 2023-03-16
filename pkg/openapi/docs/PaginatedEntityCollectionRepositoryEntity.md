# PaginatedEntityCollectionRepositoryEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]RepositoryEntity**](RepositoryEntity.md) |  | 
**Marker** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginatedEntityCollectionRepositoryEntity

`func NewPaginatedEntityCollectionRepositoryEntity(entities []RepositoryEntity, ) *PaginatedEntityCollectionRepositoryEntity`

NewPaginatedEntityCollectionRepositoryEntity instantiates a new PaginatedEntityCollectionRepositoryEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEntityCollectionRepositoryEntityWithDefaults

`func NewPaginatedEntityCollectionRepositoryEntityWithDefaults() *PaginatedEntityCollectionRepositoryEntity`

NewPaginatedEntityCollectionRepositoryEntityWithDefaults instantiates a new PaginatedEntityCollectionRepositoryEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *PaginatedEntityCollectionRepositoryEntity) GetEntities() []RepositoryEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PaginatedEntityCollectionRepositoryEntity) GetEntitiesOk() (*[]RepositoryEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PaginatedEntityCollectionRepositoryEntity) SetEntities(v []RepositoryEntity)`

SetEntities sets Entities field to given value.


### GetMarker

`func (o *PaginatedEntityCollectionRepositoryEntity) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *PaginatedEntityCollectionRepositoryEntity) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *PaginatedEntityCollectionRepositoryEntity) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *PaginatedEntityCollectionRepositoryEntity) HasMarker() bool`

HasMarker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


