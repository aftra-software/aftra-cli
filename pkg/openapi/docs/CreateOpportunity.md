# CreateOpportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Name** | **string** |  | 
**Details** | **map[string]string** |  | 
**Score** | [**OpportunityScore**](OpportunityScore.md) |  | 

## Methods

### NewCreateOpportunity

`func NewCreateOpportunity(uid string, name string, details map[string]string, score OpportunityScore, ) *CreateOpportunity`

NewCreateOpportunity instantiates a new CreateOpportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOpportunityWithDefaults

`func NewCreateOpportunityWithDefaults() *CreateOpportunity`

NewCreateOpportunityWithDefaults instantiates a new CreateOpportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *CreateOpportunity) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CreateOpportunity) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CreateOpportunity) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *CreateOpportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOpportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOpportunity) SetName(v string)`

SetName sets Name field to given value.


### GetDetails

`func (o *CreateOpportunity) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CreateOpportunity) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CreateOpportunity) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.


### GetScore

`func (o *CreateOpportunity) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateOpportunity) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateOpportunity) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


