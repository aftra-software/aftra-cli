# GlobalStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "GLOBAL_STATS"]
**Company** | Pointer to **string** |  | [optional] [default to "GLOBAL-STATS"]
**Date** | **string** |  | 
**GlobalMean** | **int32** |  | 
**BySector** | [**map[string]CompanySectorScore**](CompanySectorScore.md) |  | 

## Methods

### NewGlobalStats

`func NewGlobalStats(date string, globalMean int32, bySector map[string]CompanySectorScore, ) *GlobalStats`

NewGlobalStats instantiates a new GlobalStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalStatsWithDefaults

`func NewGlobalStatsWithDefaults() *GlobalStats`

NewGlobalStatsWithDefaults instantiates a new GlobalStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GlobalStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GlobalStats) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCompany

`func (o *GlobalStats) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GlobalStats) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GlobalStats) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GlobalStats) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDate

`func (o *GlobalStats) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GlobalStats) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GlobalStats) SetDate(v string)`

SetDate sets Date field to given value.


### GetGlobalMean

`func (o *GlobalStats) GetGlobalMean() int32`

GetGlobalMean returns the GlobalMean field if non-nil, zero value otherwise.

### GetGlobalMeanOk

`func (o *GlobalStats) GetGlobalMeanOk() (*int32, bool)`

GetGlobalMeanOk returns a tuple with the GlobalMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalMean

`func (o *GlobalStats) SetGlobalMean(v int32)`

SetGlobalMean sets GlobalMean field to given value.


### GetBySector

`func (o *GlobalStats) GetBySector() map[string]CompanySectorScore`

GetBySector returns the BySector field if non-nil, zero value otherwise.

### GetBySectorOk

`func (o *GlobalStats) GetBySectorOk() (*map[string]CompanySectorScore, bool)`

GetBySectorOk returns a tuple with the BySector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySector

`func (o *GlobalStats) SetBySector(v map[string]CompanySectorScore)`

SetBySector sets BySector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


