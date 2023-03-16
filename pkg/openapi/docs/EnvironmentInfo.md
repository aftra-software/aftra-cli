# EnvironmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIps** | **[]string** |  | 
**WasCidrs** | **[]string** |  | 
**VulnCidrs** | **[]string** |  | 
**Env** | **string** |  | 

## Methods

### NewEnvironmentInfo

`func NewEnvironmentInfo(externalIps []string, wasCidrs []string, vulnCidrs []string, env string, ) *EnvironmentInfo`

NewEnvironmentInfo instantiates a new EnvironmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentInfoWithDefaults

`func NewEnvironmentInfoWithDefaults() *EnvironmentInfo`

NewEnvironmentInfoWithDefaults instantiates a new EnvironmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIps

`func (o *EnvironmentInfo) GetExternalIps() []string`

GetExternalIps returns the ExternalIps field if non-nil, zero value otherwise.

### GetExternalIpsOk

`func (o *EnvironmentInfo) GetExternalIpsOk() (*[]string, bool)`

GetExternalIpsOk returns a tuple with the ExternalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIps

`func (o *EnvironmentInfo) SetExternalIps(v []string)`

SetExternalIps sets ExternalIps field to given value.


### GetWasCidrs

`func (o *EnvironmentInfo) GetWasCidrs() []string`

GetWasCidrs returns the WasCidrs field if non-nil, zero value otherwise.

### GetWasCidrsOk

`func (o *EnvironmentInfo) GetWasCidrsOk() (*[]string, bool)`

GetWasCidrsOk returns a tuple with the WasCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasCidrs

`func (o *EnvironmentInfo) SetWasCidrs(v []string)`

SetWasCidrs sets WasCidrs field to given value.


### GetVulnCidrs

`func (o *EnvironmentInfo) GetVulnCidrs() []string`

GetVulnCidrs returns the VulnCidrs field if non-nil, zero value otherwise.

### GetVulnCidrsOk

`func (o *EnvironmentInfo) GetVulnCidrsOk() (*[]string, bool)`

GetVulnCidrsOk returns a tuple with the VulnCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnCidrs

`func (o *EnvironmentInfo) SetVulnCidrs(v []string)`

SetVulnCidrs sets VulnCidrs field to given value.


### GetEnv

`func (o *EnvironmentInfo) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EnvironmentInfo) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EnvironmentInfo) SetEnv(v string)`

SetEnv sets Env field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


