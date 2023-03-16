# IndividualScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | [**ScoreLevel**](ScoreLevel.md) |  | 
**Reason** | **string** |  | 

## Methods

### NewIndividualScore

`func NewIndividualScore(score ScoreLevel, reason string, ) *IndividualScore`

NewIndividualScore instantiates a new IndividualScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualScoreWithDefaults

`func NewIndividualScoreWithDefaults() *IndividualScore`

NewIndividualScoreWithDefaults instantiates a new IndividualScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *IndividualScore) GetScore() ScoreLevel`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IndividualScore) GetScoreOk() (*ScoreLevel, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IndividualScore) SetScore(v ScoreLevel)`

SetScore sets Score field to given value.


### GetReason

`func (o *IndividualScore) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IndividualScore) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IndividualScore) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


