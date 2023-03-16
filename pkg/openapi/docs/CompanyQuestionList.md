# CompanyQuestionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "company-questions"]
**Mfa** | Pointer to [**MFAAnswer**](MFAAnswer.md) |  | [optional] 
**PasswordLength** | Pointer to [**PasswordLengthAnswer**](PasswordLengthAnswer.md) |  | [optional] 
**PasswordManager** | Pointer to [**PasswordManagerAnswer**](PasswordManagerAnswer.md) |  | [optional] 
**PasswordScreen** | Pointer to [**PasswordScreenAnswer**](PasswordScreenAnswer.md) |  | [optional] 
**DisclosurePolicy** | Pointer to [**DisclosurePolicyAnswer**](DisclosurePolicyAnswer.md) |  | [optional] 
**Vdp** | Pointer to [**VDPAnswer**](VDPAnswer.md) |  | [optional] 
**Vip** | Pointer to [**VIPAnswer**](VIPAnswer.md) |  | [optional] 
**VulnScan** | Pointer to [**VulnScanAnswer**](VulnScanAnswer.md) |  | [optional] 

## Methods

### NewCompanyQuestionList

`func NewCompanyQuestionList(pk string, ) *CompanyQuestionList`

NewCompanyQuestionList instantiates a new CompanyQuestionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyQuestionListWithDefaults

`func NewCompanyQuestionListWithDefaults() *CompanyQuestionList`

NewCompanyQuestionListWithDefaults instantiates a new CompanyQuestionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *CompanyQuestionList) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *CompanyQuestionList) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *CompanyQuestionList) SetPk(v string)`

SetPk sets Pk field to given value.


### GetUpdated

`func (o *CompanyQuestionList) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CompanyQuestionList) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CompanyQuestionList) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CompanyQuestionList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetType

`func (o *CompanyQuestionList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyQuestionList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyQuestionList) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyQuestionList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMfa

`func (o *CompanyQuestionList) GetMfa() MFAAnswer`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *CompanyQuestionList) GetMfaOk() (*MFAAnswer, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *CompanyQuestionList) SetMfa(v MFAAnswer)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *CompanyQuestionList) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetPasswordLength

`func (o *CompanyQuestionList) GetPasswordLength() PasswordLengthAnswer`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *CompanyQuestionList) GetPasswordLengthOk() (*PasswordLengthAnswer, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *CompanyQuestionList) SetPasswordLength(v PasswordLengthAnswer)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *CompanyQuestionList) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetPasswordManager

`func (o *CompanyQuestionList) GetPasswordManager() PasswordManagerAnswer`

GetPasswordManager returns the PasswordManager field if non-nil, zero value otherwise.

### GetPasswordManagerOk

`func (o *CompanyQuestionList) GetPasswordManagerOk() (*PasswordManagerAnswer, bool)`

GetPasswordManagerOk returns a tuple with the PasswordManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManager

`func (o *CompanyQuestionList) SetPasswordManager(v PasswordManagerAnswer)`

SetPasswordManager sets PasswordManager field to given value.

### HasPasswordManager

`func (o *CompanyQuestionList) HasPasswordManager() bool`

HasPasswordManager returns a boolean if a field has been set.

### GetPasswordScreen

`func (o *CompanyQuestionList) GetPasswordScreen() PasswordScreenAnswer`

GetPasswordScreen returns the PasswordScreen field if non-nil, zero value otherwise.

### GetPasswordScreenOk

`func (o *CompanyQuestionList) GetPasswordScreenOk() (*PasswordScreenAnswer, bool)`

GetPasswordScreenOk returns a tuple with the PasswordScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordScreen

`func (o *CompanyQuestionList) SetPasswordScreen(v PasswordScreenAnswer)`

SetPasswordScreen sets PasswordScreen field to given value.

### HasPasswordScreen

`func (o *CompanyQuestionList) HasPasswordScreen() bool`

HasPasswordScreen returns a boolean if a field has been set.

### GetDisclosurePolicy

`func (o *CompanyQuestionList) GetDisclosurePolicy() DisclosurePolicyAnswer`

GetDisclosurePolicy returns the DisclosurePolicy field if non-nil, zero value otherwise.

### GetDisclosurePolicyOk

`func (o *CompanyQuestionList) GetDisclosurePolicyOk() (*DisclosurePolicyAnswer, bool)`

GetDisclosurePolicyOk returns a tuple with the DisclosurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosurePolicy

`func (o *CompanyQuestionList) SetDisclosurePolicy(v DisclosurePolicyAnswer)`

SetDisclosurePolicy sets DisclosurePolicy field to given value.

### HasDisclosurePolicy

`func (o *CompanyQuestionList) HasDisclosurePolicy() bool`

HasDisclosurePolicy returns a boolean if a field has been set.

### GetVdp

`func (o *CompanyQuestionList) GetVdp() VDPAnswer`

GetVdp returns the Vdp field if non-nil, zero value otherwise.

### GetVdpOk

`func (o *CompanyQuestionList) GetVdpOk() (*VDPAnswer, bool)`

GetVdpOk returns a tuple with the Vdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdp

`func (o *CompanyQuestionList) SetVdp(v VDPAnswer)`

SetVdp sets Vdp field to given value.

### HasVdp

`func (o *CompanyQuestionList) HasVdp() bool`

HasVdp returns a boolean if a field has been set.

### GetVip

`func (o *CompanyQuestionList) GetVip() VIPAnswer`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *CompanyQuestionList) GetVipOk() (*VIPAnswer, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *CompanyQuestionList) SetVip(v VIPAnswer)`

SetVip sets Vip field to given value.

### HasVip

`func (o *CompanyQuestionList) HasVip() bool`

HasVip returns a boolean if a field has been set.

### GetVulnScan

`func (o *CompanyQuestionList) GetVulnScan() VulnScanAnswer`

GetVulnScan returns the VulnScan field if non-nil, zero value otherwise.

### GetVulnScanOk

`func (o *CompanyQuestionList) GetVulnScanOk() (*VulnScanAnswer, bool)`

GetVulnScanOk returns a tuple with the VulnScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnScan

`func (o *CompanyQuestionList) SetVulnScan(v VulnScanAnswer)`

SetVulnScan sets VulnScan field to given value.

### HasVulnScan

`func (o *CompanyQuestionList) HasVulnScan() bool`

HasVulnScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


