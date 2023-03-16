# DNSScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLookup** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "dns.whois"]

## Methods

### NewDNSScan

`func NewDNSScan(sourceLookup string, ) *DNSScan`

NewDNSScan instantiates a new DNSScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSScanWithDefaults

`func NewDNSScanWithDefaults() *DNSScan`

NewDNSScanWithDefaults instantiates a new DNSScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLookup

`func (o *DNSScan) GetSourceLookup() string`

GetSourceLookup returns the SourceLookup field if non-nil, zero value otherwise.

### GetSourceLookupOk

`func (o *DNSScan) GetSourceLookupOk() (*string, bool)`

GetSourceLookupOk returns a tuple with the SourceLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLookup

`func (o *DNSScan) SetSourceLookup(v string)`

SetSourceLookup sets SourceLookup field to given value.


### GetType

`func (o *DNSScan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSScan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSScan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSScan) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


