# PaginatedEntityCollectionInviteEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]InviteEntity**](InviteEntity.md) |  | 
**Marker** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginatedEntityCollectionInviteEntity

`func NewPaginatedEntityCollectionInviteEntity(entities []InviteEntity, ) *PaginatedEntityCollectionInviteEntity`

NewPaginatedEntityCollectionInviteEntity instantiates a new PaginatedEntityCollectionInviteEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEntityCollectionInviteEntityWithDefaults

`func NewPaginatedEntityCollectionInviteEntityWithDefaults() *PaginatedEntityCollectionInviteEntity`

NewPaginatedEntityCollectionInviteEntityWithDefaults instantiates a new PaginatedEntityCollectionInviteEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *PaginatedEntityCollectionInviteEntity) GetEntities() []InviteEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PaginatedEntityCollectionInviteEntity) GetEntitiesOk() (*[]InviteEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PaginatedEntityCollectionInviteEntity) SetEntities(v []InviteEntity)`

SetEntities sets Entities field to given value.


### GetMarker

`func (o *PaginatedEntityCollectionInviteEntity) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *PaginatedEntityCollectionInviteEntity) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *PaginatedEntityCollectionInviteEntity) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *PaginatedEntityCollectionInviteEntity) HasMarker() bool`

HasMarker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


