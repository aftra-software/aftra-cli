# Opportunity1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Uid** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "internal"]
**Group** | Pointer to **string** |  | [optional] [default to "internal"]
**Entity** | **string** |  | 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 
**Resolution** | Pointer to [**OpportunityResolution**](OpportunityResolution.md) |  | [optional] 
**Name** | **string** |  | 
**AccountsCount** | **int32** |  | 
**AccountsRatio** | **float32** |  | 
**AccountsBySource** | **map[string]int32** |  | 
**TypeIndex** | **string** |  | 
**Description** | **[]string** |  | 
**Origin** | **string** |  | 
**Path** | **string** |  | 
**Background** | **string** |  | 
**Remediation** | Pointer to **string** |  | [optional] 
**Filename** | **string** |  | 
**Secrets** | [**[]Secret**](Secret.md) |  | 
**Exposure** | [**[]Exposure**](Exposure.md) |  | 
**Affected** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**SolutionType** | Pointer to **string** |  | [optional] 
**Insight** | Pointer to **string** |  | [optional] 
**References** | **[]string** |  | 
**Host** | **string** |  | 
**Hostname** | Pointer to **string** |  | [optional] 
**PasswordHint** | **string** |  | 
**LeakName** | **string** |  | 
**LeakYear** | **int32** |  | 
**NucleiTemplateId** | **string** |  | 
**ScoreDetails** | [**map[string][]IndividualScore**](array.md) |  | 
**Details** | **map[string]string** |  | 

## Methods

### NewOpportunity1

`func NewOpportunity1(company string, uid string, entity string, name string, accountsCount int32, accountsRatio float32, accountsBySource map[string]int32, typeIndex string, description []string, origin string, path string, background string, filename string, secrets []Secret, exposure []Exposure, references []string, host string, passwordHint string, leakName string, leakYear int32, nucleiTemplateId string, scoreDetails map[string][]IndividualScore, details map[string]string, ) *Opportunity1`

NewOpportunity1 instantiates a new Opportunity1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunity1WithDefaults

`func NewOpportunity1WithDefaults() *Opportunity1`

NewOpportunity1WithDefaults instantiates a new Opportunity1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *Opportunity1) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Opportunity1) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Opportunity1) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetUid

`func (o *Opportunity1) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Opportunity1) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Opportunity1) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *Opportunity1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Opportunity1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Opportunity1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Opportunity1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *Opportunity1) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Opportunity1) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Opportunity1) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Opportunity1) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEntity

`func (o *Opportunity1) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Opportunity1) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Opportunity1) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetScore

`func (o *Opportunity1) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Opportunity1) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Opportunity1) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *Opportunity1) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetResolution

`func (o *Opportunity1) GetResolution() OpportunityResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Opportunity1) GetResolutionOk() (*OpportunityResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Opportunity1) SetResolution(v OpportunityResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Opportunity1) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetName

`func (o *Opportunity1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Opportunity1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Opportunity1) SetName(v string)`

SetName sets Name field to given value.


### GetAccountsCount

`func (o *Opportunity1) GetAccountsCount() int32`

GetAccountsCount returns the AccountsCount field if non-nil, zero value otherwise.

### GetAccountsCountOk

`func (o *Opportunity1) GetAccountsCountOk() (*int32, bool)`

GetAccountsCountOk returns a tuple with the AccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsCount

`func (o *Opportunity1) SetAccountsCount(v int32)`

SetAccountsCount sets AccountsCount field to given value.


### GetAccountsRatio

`func (o *Opportunity1) GetAccountsRatio() float32`

GetAccountsRatio returns the AccountsRatio field if non-nil, zero value otherwise.

### GetAccountsRatioOk

`func (o *Opportunity1) GetAccountsRatioOk() (*float32, bool)`

GetAccountsRatioOk returns a tuple with the AccountsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsRatio

`func (o *Opportunity1) SetAccountsRatio(v float32)`

SetAccountsRatio sets AccountsRatio field to given value.


### GetAccountsBySource

`func (o *Opportunity1) GetAccountsBySource() map[string]int32`

GetAccountsBySource returns the AccountsBySource field if non-nil, zero value otherwise.

### GetAccountsBySourceOk

`func (o *Opportunity1) GetAccountsBySourceOk() (*map[string]int32, bool)`

GetAccountsBySourceOk returns a tuple with the AccountsBySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsBySource

`func (o *Opportunity1) SetAccountsBySource(v map[string]int32)`

SetAccountsBySource sets AccountsBySource field to given value.


### GetTypeIndex

`func (o *Opportunity1) GetTypeIndex() string`

GetTypeIndex returns the TypeIndex field if non-nil, zero value otherwise.

### GetTypeIndexOk

`func (o *Opportunity1) GetTypeIndexOk() (*string, bool)`

GetTypeIndexOk returns a tuple with the TypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIndex

`func (o *Opportunity1) SetTypeIndex(v string)`

SetTypeIndex sets TypeIndex field to given value.


### GetDescription

`func (o *Opportunity1) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Opportunity1) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Opportunity1) SetDescription(v []string)`

SetDescription sets Description field to given value.


### GetOrigin

`func (o *Opportunity1) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Opportunity1) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Opportunity1) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetPath

`func (o *Opportunity1) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Opportunity1) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Opportunity1) SetPath(v string)`

SetPath sets Path field to given value.


### GetBackground

`func (o *Opportunity1) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *Opportunity1) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *Opportunity1) SetBackground(v string)`

SetBackground sets Background field to given value.


### GetRemediation

`func (o *Opportunity1) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *Opportunity1) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *Opportunity1) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *Opportunity1) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetFilename

`func (o *Opportunity1) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Opportunity1) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Opportunity1) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSecrets

`func (o *Opportunity1) GetSecrets() []Secret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *Opportunity1) GetSecretsOk() (*[]Secret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *Opportunity1) SetSecrets(v []Secret)`

SetSecrets sets Secrets field to given value.


### GetExposure

`func (o *Opportunity1) GetExposure() []Exposure`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *Opportunity1) GetExposureOk() (*[]Exposure, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *Opportunity1) SetExposure(v []Exposure)`

SetExposure sets Exposure field to given value.


### GetAffected

`func (o *Opportunity1) GetAffected() string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *Opportunity1) GetAffectedOk() (*string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *Opportunity1) SetAffected(v string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *Opportunity1) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetSummary

`func (o *Opportunity1) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Opportunity1) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Opportunity1) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Opportunity1) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSolution

`func (o *Opportunity1) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *Opportunity1) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *Opportunity1) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *Opportunity1) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetSolutionType

`func (o *Opportunity1) GetSolutionType() string`

GetSolutionType returns the SolutionType field if non-nil, zero value otherwise.

### GetSolutionTypeOk

`func (o *Opportunity1) GetSolutionTypeOk() (*string, bool)`

GetSolutionTypeOk returns a tuple with the SolutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionType

`func (o *Opportunity1) SetSolutionType(v string)`

SetSolutionType sets SolutionType field to given value.

### HasSolutionType

`func (o *Opportunity1) HasSolutionType() bool`

HasSolutionType returns a boolean if a field has been set.

### GetInsight

`func (o *Opportunity1) GetInsight() string`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *Opportunity1) GetInsightOk() (*string, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *Opportunity1) SetInsight(v string)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *Opportunity1) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### GetReferences

`func (o *Opportunity1) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Opportunity1) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Opportunity1) SetReferences(v []string)`

SetReferences sets References field to given value.


### GetHost

`func (o *Opportunity1) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Opportunity1) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Opportunity1) SetHost(v string)`

SetHost sets Host field to given value.


### GetHostname

`func (o *Opportunity1) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Opportunity1) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Opportunity1) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Opportunity1) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPasswordHint

`func (o *Opportunity1) GetPasswordHint() string`

GetPasswordHint returns the PasswordHint field if non-nil, zero value otherwise.

### GetPasswordHintOk

`func (o *Opportunity1) GetPasswordHintOk() (*string, bool)`

GetPasswordHintOk returns a tuple with the PasswordHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHint

`func (o *Opportunity1) SetPasswordHint(v string)`

SetPasswordHint sets PasswordHint field to given value.


### GetLeakName

`func (o *Opportunity1) GetLeakName() string`

GetLeakName returns the LeakName field if non-nil, zero value otherwise.

### GetLeakNameOk

`func (o *Opportunity1) GetLeakNameOk() (*string, bool)`

GetLeakNameOk returns a tuple with the LeakName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeakName

`func (o *Opportunity1) SetLeakName(v string)`

SetLeakName sets LeakName field to given value.


### GetLeakYear

`func (o *Opportunity1) GetLeakYear() int32`

GetLeakYear returns the LeakYear field if non-nil, zero value otherwise.

### GetLeakYearOk

`func (o *Opportunity1) GetLeakYearOk() (*int32, bool)`

GetLeakYearOk returns a tuple with the LeakYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeakYear

`func (o *Opportunity1) SetLeakYear(v int32)`

SetLeakYear sets LeakYear field to given value.


### GetNucleiTemplateId

`func (o *Opportunity1) GetNucleiTemplateId() string`

GetNucleiTemplateId returns the NucleiTemplateId field if non-nil, zero value otherwise.

### GetNucleiTemplateIdOk

`func (o *Opportunity1) GetNucleiTemplateIdOk() (*string, bool)`

GetNucleiTemplateIdOk returns a tuple with the NucleiTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleiTemplateId

`func (o *Opportunity1) SetNucleiTemplateId(v string)`

SetNucleiTemplateId sets NucleiTemplateId field to given value.


### GetScoreDetails

`func (o *Opportunity1) GetScoreDetails() map[string][]IndividualScore`

GetScoreDetails returns the ScoreDetails field if non-nil, zero value otherwise.

### GetScoreDetailsOk

`func (o *Opportunity1) GetScoreDetailsOk() (*map[string][]IndividualScore, bool)`

GetScoreDetailsOk returns a tuple with the ScoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreDetails

`func (o *Opportunity1) SetScoreDetails(v map[string][]IndividualScore)`

SetScoreDetails sets ScoreDetails field to given value.


### GetDetails

`func (o *Opportunity1) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Opportunity1) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Opportunity1) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


