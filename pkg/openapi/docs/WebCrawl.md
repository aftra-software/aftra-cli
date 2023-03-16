# WebCrawl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLookup** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "scan.webpage"]

## Methods

### NewWebCrawl

`func NewWebCrawl(sourceLookup string, ) *WebCrawl`

NewWebCrawl instantiates a new WebCrawl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebCrawlWithDefaults

`func NewWebCrawlWithDefaults() *WebCrawl`

NewWebCrawlWithDefaults instantiates a new WebCrawl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLookup

`func (o *WebCrawl) GetSourceLookup() string`

GetSourceLookup returns the SourceLookup field if non-nil, zero value otherwise.

### GetSourceLookupOk

`func (o *WebCrawl) GetSourceLookupOk() (*string, bool)`

GetSourceLookupOk returns a tuple with the SourceLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLookup

`func (o *WebCrawl) SetSourceLookup(v string)`

SetSourceLookup sets SourceLookup field to given value.


### GetType

`func (o *WebCrawl) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebCrawl) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebCrawl) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebCrawl) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


