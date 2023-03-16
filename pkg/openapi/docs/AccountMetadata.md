# AccountMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceRunningCount** | Pointer to [**ModelMap**](map.md) |  | [optional] 
**RiskScore** | Pointer to **float32** |  | [optional] [default to 0]
**CampaignConclusion** | Pointer to [**AccountCampaignConclusion**](AccountCampaignConclusion.md) |  | [optional] 

## Methods

### NewAccountMetadata

`func NewAccountMetadata() *AccountMetadata`

NewAccountMetadata instantiates a new AccountMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountMetadataWithDefaults

`func NewAccountMetadataWithDefaults() *AccountMetadata`

NewAccountMetadataWithDefaults instantiates a new AccountMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceRunningCount

`func (o *AccountMetadata) GetServiceRunningCount() ModelMap`

GetServiceRunningCount returns the ServiceRunningCount field if non-nil, zero value otherwise.

### GetServiceRunningCountOk

`func (o *AccountMetadata) GetServiceRunningCountOk() (*ModelMap, bool)`

GetServiceRunningCountOk returns a tuple with the ServiceRunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRunningCount

`func (o *AccountMetadata) SetServiceRunningCount(v ModelMap)`

SetServiceRunningCount sets ServiceRunningCount field to given value.

### HasServiceRunningCount

`func (o *AccountMetadata) HasServiceRunningCount() bool`

HasServiceRunningCount returns a boolean if a field has been set.

### GetRiskScore

`func (o *AccountMetadata) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *AccountMetadata) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *AccountMetadata) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *AccountMetadata) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetCampaignConclusion

`func (o *AccountMetadata) GetCampaignConclusion() AccountCampaignConclusion`

GetCampaignConclusion returns the CampaignConclusion field if non-nil, zero value otherwise.

### GetCampaignConclusionOk

`func (o *AccountMetadata) GetCampaignConclusionOk() (*AccountCampaignConclusion, bool)`

GetCampaignConclusionOk returns a tuple with the CampaignConclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignConclusion

`func (o *AccountMetadata) SetCampaignConclusion(v AccountCampaignConclusion)`

SetCampaignConclusion sets CampaignConclusion field to given value.

### HasCampaignConclusion

`func (o *AccountMetadata) HasCampaignConclusion() bool`

HasCampaignConclusion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


