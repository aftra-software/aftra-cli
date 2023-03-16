# OpportunityDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextDescription** | **string** |  | 
**Opportunity** | [**Opportunity1**](Opportunity1.md) |  | 

## Methods

### NewOpportunityDetails

`func NewOpportunityDetails(contextDescription string, opportunity Opportunity1, ) *OpportunityDetails`

NewOpportunityDetails instantiates a new OpportunityDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityDetailsWithDefaults

`func NewOpportunityDetailsWithDefaults() *OpportunityDetails`

NewOpportunityDetailsWithDefaults instantiates a new OpportunityDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextDescription

`func (o *OpportunityDetails) GetContextDescription() string`

GetContextDescription returns the ContextDescription field if non-nil, zero value otherwise.

### GetContextDescriptionOk

`func (o *OpportunityDetails) GetContextDescriptionOk() (*string, bool)`

GetContextDescriptionOk returns a tuple with the ContextDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextDescription

`func (o *OpportunityDetails) SetContextDescription(v string)`

SetContextDescription sets ContextDescription field to given value.


### GetOpportunity

`func (o *OpportunityDetails) GetOpportunity() Opportunity1`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *OpportunityDetails) GetOpportunityOk() (*Opportunity1, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *OpportunityDetails) SetOpportunity(v Opportunity1)`

SetOpportunity sets Opportunity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


