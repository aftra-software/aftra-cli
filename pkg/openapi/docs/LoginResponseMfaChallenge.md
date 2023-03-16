# LoginResponseMfaChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | [**ChallengeTypes**](ChallengeTypes.md) |  | 
**SetupSecret** | Pointer to **string** |  | [optional] 
**Session** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginResponseMfaChallenge

`func NewLoginResponseMfaChallenge(challenge ChallengeTypes, ) *LoginResponseMfaChallenge`

NewLoginResponseMfaChallenge instantiates a new LoginResponseMfaChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginResponseMfaChallengeWithDefaults

`func NewLoginResponseMfaChallengeWithDefaults() *LoginResponseMfaChallenge`

NewLoginResponseMfaChallengeWithDefaults instantiates a new LoginResponseMfaChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *LoginResponseMfaChallenge) GetChallenge() ChallengeTypes`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *LoginResponseMfaChallenge) GetChallengeOk() (*ChallengeTypes, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *LoginResponseMfaChallenge) SetChallenge(v ChallengeTypes)`

SetChallenge sets Challenge field to given value.


### GetSetupSecret

`func (o *LoginResponseMfaChallenge) GetSetupSecret() string`

GetSetupSecret returns the SetupSecret field if non-nil, zero value otherwise.

### GetSetupSecretOk

`func (o *LoginResponseMfaChallenge) GetSetupSecretOk() (*string, bool)`

GetSetupSecretOk returns a tuple with the SetupSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupSecret

`func (o *LoginResponseMfaChallenge) SetSetupSecret(v string)`

SetSetupSecret sets SetupSecret field to given value.

### HasSetupSecret

`func (o *LoginResponseMfaChallenge) HasSetupSecret() bool`

HasSetupSecret returns a boolean if a field has been set.

### GetSession

`func (o *LoginResponseMfaChallenge) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *LoginResponseMfaChallenge) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *LoginResponseMfaChallenge) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *LoginResponseMfaChallenge) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


