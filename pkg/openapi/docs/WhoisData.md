# WhoisData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] [default to true]
**Address** | Pointer to **[]string** |  | [optional] 
**CreationDate** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Dnssec** | Pointer to **[]string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **[]string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**RegistrantName** | Pointer to **[]string** |  | [optional] 
**RawText** | **string** |  | 

## Methods

### NewWhoisData

`func NewWhoisData(rawText string, ) *WhoisData`

NewWhoisData instantiates a new WhoisData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoisDataWithDefaults

`func NewWhoisDataWithDefaults() *WhoisData`

NewWhoisDataWithDefaults instantiates a new WhoisData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WhoisData) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WhoisData) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WhoisData) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *WhoisData) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAddress

`func (o *WhoisData) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WhoisData) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WhoisData) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WhoisData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreationDate

`func (o *WhoisData) GetCreationDate() []time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *WhoisData) GetCreationDateOk() (*[]time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *WhoisData) SetCreationDate(v []time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *WhoisData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDnssec

`func (o *WhoisData) GetDnssec() []string`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *WhoisData) GetDnssecOk() (*[]string, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *WhoisData) SetDnssec(v []string)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *WhoisData) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.

### GetDomainName

`func (o *WhoisData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *WhoisData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *WhoisData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *WhoisData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEmail

`func (o *WhoisData) GetEmail() []string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WhoisData) GetEmailOk() (*[]string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WhoisData) SetEmail(v []string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WhoisData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpirationDate

`func (o *WhoisData) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *WhoisData) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *WhoisData) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *WhoisData) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetName

`func (o *WhoisData) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhoisData) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhoisData) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *WhoisData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameServers

`func (o *WhoisData) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *WhoisData) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *WhoisData) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *WhoisData) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetRegistrantName

`func (o *WhoisData) GetRegistrantName() []string`

GetRegistrantName returns the RegistrantName field if non-nil, zero value otherwise.

### GetRegistrantNameOk

`func (o *WhoisData) GetRegistrantNameOk() (*[]string, bool)`

GetRegistrantNameOk returns a tuple with the RegistrantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantName

`func (o *WhoisData) SetRegistrantName(v []string)`

SetRegistrantName sets RegistrantName field to given value.

### HasRegistrantName

`func (o *WhoisData) HasRegistrantName() bool`

HasRegistrantName returns a boolean if a field has been set.

### GetRawText

`func (o *WhoisData) GetRawText() string`

GetRawText returns the RawText field if non-nil, zero value otherwise.

### GetRawTextOk

`func (o *WhoisData) GetRawTextOk() (*string, bool)`

GetRawTextOk returns a tuple with the RawText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawText

`func (o *WhoisData) SetRawText(v string)`

SetRawText sets RawText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


