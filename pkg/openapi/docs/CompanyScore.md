# CompanyScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "COMPANY_SCORE"]
**Company** | **string** |  | 
**Date** | Pointer to **string** |  | [optional] 
**Total** | **int32** |  | 
**Breakdown** | **map[string]int32** |  | 
**Ttl** | Pointer to **float32** |  | [optional] 

## Methods

### NewCompanyScore

`func NewCompanyScore(company string, total int32, breakdown map[string]int32, ) *CompanyScore`

NewCompanyScore instantiates a new CompanyScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyScoreWithDefaults

`func NewCompanyScoreWithDefaults() *CompanyScore`

NewCompanyScoreWithDefaults instantiates a new CompanyScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompanyScore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyScore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyScore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyScore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyScore) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyScore) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyScore) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetDate

`func (o *CompanyScore) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CompanyScore) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CompanyScore) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CompanyScore) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTotal

`func (o *CompanyScore) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CompanyScore) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CompanyScore) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetBreakdown

`func (o *CompanyScore) GetBreakdown() map[string]int32`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *CompanyScore) GetBreakdownOk() (*map[string]int32, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *CompanyScore) SetBreakdown(v map[string]int32)`

SetBreakdown sets Breakdown field to given value.


### GetTtl

`func (o *CompanyScore) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CompanyScore) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CompanyScore) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CompanyScore) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


