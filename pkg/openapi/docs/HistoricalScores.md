# HistoricalScores

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyScores** | [**[]CompanyScore**](CompanyScore.md) |  | 
**GlobalScores** | [**[]GlobalStats**](GlobalStats.md) |  | 

## Methods

### NewHistoricalScores

`func NewHistoricalScores(companyScores []CompanyScore, globalScores []GlobalStats, ) *HistoricalScores`

NewHistoricalScores instantiates a new HistoricalScores object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalScoresWithDefaults

`func NewHistoricalScoresWithDefaults() *HistoricalScores`

NewHistoricalScoresWithDefaults instantiates a new HistoricalScores object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyScores

`func (o *HistoricalScores) GetCompanyScores() []CompanyScore`

GetCompanyScores returns the CompanyScores field if non-nil, zero value otherwise.

### GetCompanyScoresOk

`func (o *HistoricalScores) GetCompanyScoresOk() (*[]CompanyScore, bool)`

GetCompanyScoresOk returns a tuple with the CompanyScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyScores

`func (o *HistoricalScores) SetCompanyScores(v []CompanyScore)`

SetCompanyScores sets CompanyScores field to given value.


### GetGlobalScores

`func (o *HistoricalScores) GetGlobalScores() []GlobalStats`

GetGlobalScores returns the GlobalScores field if non-nil, zero value otherwise.

### GetGlobalScoresOk

`func (o *HistoricalScores) GetGlobalScoresOk() (*[]GlobalStats, bool)`

GetGlobalScoresOk returns a tuple with the GlobalScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalScores

`func (o *HistoricalScores) SetGlobalScores(v []GlobalStats)`

SetGlobalScores sets GlobalScores field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


