# CompanySectorScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectorMean** | **int32** |  | 
**ByGroup** | **map[string]int32** |  | 

## Methods

### NewCompanySectorScore

`func NewCompanySectorScore(sectorMean int32, byGroup map[string]int32, ) *CompanySectorScore`

NewCompanySectorScore instantiates a new CompanySectorScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanySectorScoreWithDefaults

`func NewCompanySectorScoreWithDefaults() *CompanySectorScore`

NewCompanySectorScoreWithDefaults instantiates a new CompanySectorScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectorMean

`func (o *CompanySectorScore) GetSectorMean() int32`

GetSectorMean returns the SectorMean field if non-nil, zero value otherwise.

### GetSectorMeanOk

`func (o *CompanySectorScore) GetSectorMeanOk() (*int32, bool)`

GetSectorMeanOk returns a tuple with the SectorMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorMean

`func (o *CompanySectorScore) SetSectorMean(v int32)`

SetSectorMean sets SectorMean field to given value.


### GetByGroup

`func (o *CompanySectorScore) GetByGroup() map[string]int32`

GetByGroup returns the ByGroup field if non-nil, zero value otherwise.

### GetByGroupOk

`func (o *CompanySectorScore) GetByGroupOk() (*map[string]int32, bool)`

GetByGroupOk returns a tuple with the ByGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByGroup

`func (o *CompanySectorScore) SetByGroup(v map[string]int32)`

SetByGroup sets ByGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


