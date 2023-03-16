# ResponseWhoisDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] [default to false]
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
**Error** | [**Array**](array.md) |  | 

## Methods

### NewResponseWhoisDomain

`func NewResponseWhoisDomain(rawText string, error_ Array, ) *ResponseWhoisDomain`

NewResponseWhoisDomain instantiates a new ResponseWhoisDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWhoisDomainWithDefaults

`func NewResponseWhoisDomainWithDefaults() *ResponseWhoisDomain`

NewResponseWhoisDomainWithDefaults instantiates a new ResponseWhoisDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResponseWhoisDomain) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseWhoisDomain) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseWhoisDomain) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseWhoisDomain) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAddress

`func (o *ResponseWhoisDomain) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponseWhoisDomain) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponseWhoisDomain) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ResponseWhoisDomain) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreationDate

`func (o *ResponseWhoisDomain) GetCreationDate() []time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ResponseWhoisDomain) GetCreationDateOk() (*[]time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ResponseWhoisDomain) SetCreationDate(v []time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ResponseWhoisDomain) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDnssec

`func (o *ResponseWhoisDomain) GetDnssec() []string`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *ResponseWhoisDomain) GetDnssecOk() (*[]string, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *ResponseWhoisDomain) SetDnssec(v []string)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *ResponseWhoisDomain) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.

### GetDomainName

`func (o *ResponseWhoisDomain) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ResponseWhoisDomain) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ResponseWhoisDomain) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ResponseWhoisDomain) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEmail

`func (o *ResponseWhoisDomain) GetEmail() []string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResponseWhoisDomain) GetEmailOk() (*[]string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResponseWhoisDomain) SetEmail(v []string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResponseWhoisDomain) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ResponseWhoisDomain) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResponseWhoisDomain) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResponseWhoisDomain) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ResponseWhoisDomain) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetName

`func (o *ResponseWhoisDomain) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseWhoisDomain) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseWhoisDomain) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseWhoisDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameServers

`func (o *ResponseWhoisDomain) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *ResponseWhoisDomain) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *ResponseWhoisDomain) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *ResponseWhoisDomain) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetRegistrantName

`func (o *ResponseWhoisDomain) GetRegistrantName() []string`

GetRegistrantName returns the RegistrantName field if non-nil, zero value otherwise.

### GetRegistrantNameOk

`func (o *ResponseWhoisDomain) GetRegistrantNameOk() (*[]string, bool)`

GetRegistrantNameOk returns a tuple with the RegistrantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantName

`func (o *ResponseWhoisDomain) SetRegistrantName(v []string)`

SetRegistrantName sets RegistrantName field to given value.

### HasRegistrantName

`func (o *ResponseWhoisDomain) HasRegistrantName() bool`

HasRegistrantName returns a boolean if a field has been set.

### GetRawText

`func (o *ResponseWhoisDomain) GetRawText() string`

GetRawText returns the RawText field if non-nil, zero value otherwise.

### GetRawTextOk

`func (o *ResponseWhoisDomain) GetRawTextOk() (*string, bool)`

GetRawTextOk returns a tuple with the RawText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawText

`func (o *ResponseWhoisDomain) SetRawText(v string)`

SetRawText sets RawText field to given value.


### GetError

`func (o *ResponseWhoisDomain) GetError() Array`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseWhoisDomain) GetErrorOk() (*Array, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseWhoisDomain) SetError(v Array)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


