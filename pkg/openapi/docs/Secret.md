# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Secret** | **string** |  | 
**Filename** | **string** |  | 
**CommitHash** | **string** |  | 
**LineNumber** | **string** |  | 
**RepoUrl** | Pointer to **string** |  | [optional] [default to ""]
**Branch** | **string** |  | 
**Score** | Pointer to [**OpportunityScore**](OpportunityScore.md) |  | [optional] 

## Methods

### NewSecret

`func NewSecret(type_ string, secret string, filename string, commitHash string, lineNumber string, branch string, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Secret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Secret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Secret) SetType(v string)`

SetType sets Type field to given value.


### GetSecret

`func (o *Secret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Secret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Secret) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetFilename

`func (o *Secret) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Secret) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Secret) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetCommitHash

`func (o *Secret) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *Secret) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *Secret) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetLineNumber

`func (o *Secret) GetLineNumber() string`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *Secret) GetLineNumberOk() (*string, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *Secret) SetLineNumber(v string)`

SetLineNumber sets LineNumber field to given value.


### GetRepoUrl

`func (o *Secret) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *Secret) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *Secret) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *Secret) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetBranch

`func (o *Secret) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Secret) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Secret) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetScore

`func (o *Secret) GetScore() OpportunityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Secret) GetScoreOk() (*OpportunityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Secret) SetScore(v OpportunityScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *Secret) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


