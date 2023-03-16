# ExposedFileSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "exposed_secret"]
**Group** | Pointer to **string** |  | [optional] [default to "vulnerabilities"]
**Entity** | **string** |  | 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | **string** |  | 
**Filename** | **string** |  | 
**Secrets** | [**[]Secret**](Secret.md) |  | 

## Methods

### NewExposedFileSecrets

`func NewExposedFileSecrets(company string, uid string, entity string, name string, filename string, secrets []Secret, ) *ExposedFileSecrets`

NewExposedFileSecrets instantiates a new ExposedFileSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExposedFileSecretsWithDefaults

`func NewExposedFileSecretsWithDefaults() *ExposedFileSecrets`

NewExposedFileSecretsWithDefaults instantiates a new ExposedFileSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *ExposedFileSecrets) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ExposedFileSecrets) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ExposedFileSecrets) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *ExposedFileSecrets) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ExposedFileSecrets) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ExposedFileSecrets) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *ExposedFileSecrets) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExposedFileSecrets) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExposedFileSecrets) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExposedFileSecrets) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *ExposedFileSecrets) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ExposedFileSecrets) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ExposedFileSecrets) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ExposedFileSecrets) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEntity

`func (o *ExposedFileSecrets) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ExposedFileSecrets) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ExposedFileSecrets) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetScore

`func (o *ExposedFileSecrets) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ExposedFileSecrets) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ExposedFileSecrets) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *ExposedFileSecrets) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *ExposedFileSecrets) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ExposedFileSecrets) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ExposedFileSecrets) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ExposedFileSecrets) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *ExposedFileSecrets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExposedFileSecrets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExposedFileSecrets) SetName(v string)`

SetName sets Name field to given value.


### GetFilename

`func (o *ExposedFileSecrets) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ExposedFileSecrets) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ExposedFileSecrets) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSecrets

`func (o *ExposedFileSecrets) GetSecrets() []Secret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ExposedFileSecrets) GetSecretsOk() (*[]Secret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ExposedFileSecrets) SetSecrets(v []Secret)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


