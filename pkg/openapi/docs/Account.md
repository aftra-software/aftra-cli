# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FullName** | Pointer to **string** |  | [optional] [default to ""]
**Status** | Pointer to [**AccountStatusEnum**](AccountStatusEnum.md) |  | [optional] 
**CampaignOnNextScan** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAccount

`func NewAccount(name string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *Account) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Account) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Account) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Account) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetStatus

`func (o *Account) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Account) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCampaignOnNextScan

`func (o *Account) GetCampaignOnNextScan() bool`

GetCampaignOnNextScan returns the CampaignOnNextScan field if non-nil, zero value otherwise.

### GetCampaignOnNextScanOk

`func (o *Account) GetCampaignOnNextScanOk() (*bool, bool)`

GetCampaignOnNextScanOk returns a tuple with the CampaignOnNextScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignOnNextScan

`func (o *Account) SetCampaignOnNextScan(v bool)`

SetCampaignOnNextScan sets CampaignOnNextScan field to given value.

### HasCampaignOnNextScan

`func (o *Account) HasCampaignOnNextScan() bool`

HasCampaignOnNextScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


