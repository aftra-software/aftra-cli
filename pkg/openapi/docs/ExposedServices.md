# ExposedServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "exposed_service"]
**Group** | Pointer to **string** |  | [optional] [default to "human_exposure"]
**Entity** | **string** |  | 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | **string** |  | 
**Exposure** | [**[]Exposure**](Exposure.md) |  | 

## Methods

### NewExposedServices

`func NewExposedServices(company string, uid string, entity string, name string, exposure []Exposure, ) *ExposedServices`

NewExposedServices instantiates a new ExposedServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExposedServicesWithDefaults

`func NewExposedServicesWithDefaults() *ExposedServices`

NewExposedServicesWithDefaults instantiates a new ExposedServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *ExposedServices) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ExposedServices) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ExposedServices) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *ExposedServices) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ExposedServices) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ExposedServices) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *ExposedServices) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExposedServices) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExposedServices) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExposedServices) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *ExposedServices) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ExposedServices) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ExposedServices) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ExposedServices) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEntity

`func (o *ExposedServices) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ExposedServices) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ExposedServices) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetScore

`func (o *ExposedServices) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ExposedServices) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ExposedServices) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *ExposedServices) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *ExposedServices) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ExposedServices) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ExposedServices) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ExposedServices) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *ExposedServices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExposedServices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExposedServices) SetName(v string)`

SetName sets Name field to given value.


### GetExposure

`func (o *ExposedServices) GetExposure() []Exposure`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *ExposedServices) GetExposureOk() (*[]Exposure, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *ExposedServices) SetExposure(v []Exposure)`

SetExposure sets Exposure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


