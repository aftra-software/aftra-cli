# PaginatedEntityCollectionIntegrationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]IntegrationEntity**](IntegrationEntity.md) |  | 
**Marker** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginatedEntityCollectionIntegrationEntity

`func NewPaginatedEntityCollectionIntegrationEntity(entities []IntegrationEntity, ) *PaginatedEntityCollectionIntegrationEntity`

NewPaginatedEntityCollectionIntegrationEntity instantiates a new PaginatedEntityCollectionIntegrationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEntityCollectionIntegrationEntityWithDefaults

`func NewPaginatedEntityCollectionIntegrationEntityWithDefaults() *PaginatedEntityCollectionIntegrationEntity`

NewPaginatedEntityCollectionIntegrationEntityWithDefaults instantiates a new PaginatedEntityCollectionIntegrationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *PaginatedEntityCollectionIntegrationEntity) GetEntities() []IntegrationEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PaginatedEntityCollectionIntegrationEntity) GetEntitiesOk() (*[]IntegrationEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PaginatedEntityCollectionIntegrationEntity) SetEntities(v []IntegrationEntity)`

SetEntities sets Entities field to given value.


### GetMarker

`func (o *PaginatedEntityCollectionIntegrationEntity) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *PaginatedEntityCollectionIntegrationEntity) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *PaginatedEntityCollectionIntegrationEntity) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *PaginatedEntityCollectionIntegrationEntity) HasMarker() bool`

HasMarker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


