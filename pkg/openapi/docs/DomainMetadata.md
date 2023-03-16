# DomainMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubdomainCount** | Pointer to **int32** |  | [optional] [default to 0]
**SaasServices** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewDomainMetadata

`func NewDomainMetadata() *DomainMetadata`

NewDomainMetadata instantiates a new DomainMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainMetadataWithDefaults

`func NewDomainMetadataWithDefaults() *DomainMetadata`

NewDomainMetadataWithDefaults instantiates a new DomainMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubdomainCount

`func (o *DomainMetadata) GetSubdomainCount() int32`

GetSubdomainCount returns the SubdomainCount field if non-nil, zero value otherwise.

### GetSubdomainCountOk

`func (o *DomainMetadata) GetSubdomainCountOk() (*int32, bool)`

GetSubdomainCountOk returns a tuple with the SubdomainCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainCount

`func (o *DomainMetadata) SetSubdomainCount(v int32)`

SetSubdomainCount sets SubdomainCount field to given value.

### HasSubdomainCount

`func (o *DomainMetadata) HasSubdomainCount() bool`

HasSubdomainCount returns a boolean if a field has been set.

### GetSaasServices

`func (o *DomainMetadata) GetSaasServices() []string`

GetSaasServices returns the SaasServices field if non-nil, zero value otherwise.

### GetSaasServicesOk

`func (o *DomainMetadata) GetSaasServicesOk() (*[]string, bool)`

GetSaasServicesOk returns a tuple with the SaasServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasServices

`func (o *DomainMetadata) SetSaasServices(v []string)`

SetSaasServices sets SaasServices field to given value.

### HasSaasServices

`func (o *DomainMetadata) HasSaasServices() bool`

HasSaasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


