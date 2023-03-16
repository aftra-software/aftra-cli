# PaginatedEntityCollectionSubdomainEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]SubdomainEntity**](SubdomainEntity.md) |  | 
**Marker** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginatedEntityCollectionSubdomainEntity

`func NewPaginatedEntityCollectionSubdomainEntity(entities []SubdomainEntity, ) *PaginatedEntityCollectionSubdomainEntity`

NewPaginatedEntityCollectionSubdomainEntity instantiates a new PaginatedEntityCollectionSubdomainEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEntityCollectionSubdomainEntityWithDefaults

`func NewPaginatedEntityCollectionSubdomainEntityWithDefaults() *PaginatedEntityCollectionSubdomainEntity`

NewPaginatedEntityCollectionSubdomainEntityWithDefaults instantiates a new PaginatedEntityCollectionSubdomainEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *PaginatedEntityCollectionSubdomainEntity) GetEntities() []SubdomainEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PaginatedEntityCollectionSubdomainEntity) GetEntitiesOk() (*[]SubdomainEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PaginatedEntityCollectionSubdomainEntity) SetEntities(v []SubdomainEntity)`

SetEntities sets Entities field to given value.


### GetMarker

`func (o *PaginatedEntityCollectionSubdomainEntity) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *PaginatedEntityCollectionSubdomainEntity) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *PaginatedEntityCollectionSubdomainEntity) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *PaginatedEntityCollectionSubdomainEntity) HasMarker() bool`

HasMarker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


