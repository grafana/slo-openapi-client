# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Freeform** | Pointer to [**FreeformQuery**](FreeformQuery.md) |  | [optional] 
**Histogram** | Pointer to [**HistogramQuery**](HistogramQuery.md) |  | [optional] 
**Ratio** | Pointer to [**RatioQuery**](RatioQuery.md) |  | [optional] 
**Threshold** | Pointer to [**ThresholdQuery**](ThresholdQuery.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewQuery

`func NewQuery(type_ string, ) *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeform

`func (o *Query) GetFreeform() FreeformQuery`

GetFreeform returns the Freeform field if non-nil, zero value otherwise.

### GetFreeformOk

`func (o *Query) GetFreeformOk() (*FreeformQuery, bool)`

GetFreeformOk returns a tuple with the Freeform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeform

`func (o *Query) SetFreeform(v FreeformQuery)`

SetFreeform sets Freeform field to given value.

### HasFreeform

`func (o *Query) HasFreeform() bool`

HasFreeform returns a boolean if a field has been set.

### GetHistogram

`func (o *Query) GetHistogram() HistogramQuery`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *Query) GetHistogramOk() (*HistogramQuery, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *Query) SetHistogram(v HistogramQuery)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *Query) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetRatio

`func (o *Query) GetRatio() RatioQuery`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *Query) GetRatioOk() (*RatioQuery, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *Query) SetRatio(v RatioQuery)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *Query) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetThreshold

`func (o *Query) GetThreshold() ThresholdQuery`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Query) GetThresholdOk() (*ThresholdQuery, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Query) SetThreshold(v ThresholdQuery)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Query) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetType

`func (o *Query) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Query) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Query) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


