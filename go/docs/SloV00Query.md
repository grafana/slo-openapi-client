# SloV00Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Freeform** | Pointer to [**SloV00FreeformQuery**](SloV00FreeformQuery.md) |  | [optional] 
**Histogram** | Pointer to [**SloV00HistogramQuery**](SloV00HistogramQuery.md) |  | [optional] 
**Ratio** | Pointer to [**SloV00RatioQuery**](SloV00RatioQuery.md) |  | [optional] 
**Threshold** | Pointer to [**SloV00ThresholdQuery**](SloV00ThresholdQuery.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewSloV00Query

`func NewSloV00Query(type_ string, ) *SloV00Query`

NewSloV00Query instantiates a new SloV00Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00QueryWithDefaults

`func NewSloV00QueryWithDefaults() *SloV00Query`

NewSloV00QueryWithDefaults instantiates a new SloV00Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeform

`func (o *SloV00Query) GetFreeform() SloV00FreeformQuery`

GetFreeform returns the Freeform field if non-nil, zero value otherwise.

### GetFreeformOk

`func (o *SloV00Query) GetFreeformOk() (*SloV00FreeformQuery, bool)`

GetFreeformOk returns a tuple with the Freeform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeform

`func (o *SloV00Query) SetFreeform(v SloV00FreeformQuery)`

SetFreeform sets Freeform field to given value.

### HasFreeform

`func (o *SloV00Query) HasFreeform() bool`

HasFreeform returns a boolean if a field has been set.

### GetHistogram

`func (o *SloV00Query) GetHistogram() SloV00HistogramQuery`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *SloV00Query) GetHistogramOk() (*SloV00HistogramQuery, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *SloV00Query) SetHistogram(v SloV00HistogramQuery)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *SloV00Query) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetRatio

`func (o *SloV00Query) GetRatio() SloV00RatioQuery`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *SloV00Query) GetRatioOk() (*SloV00RatioQuery, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *SloV00Query) SetRatio(v SloV00RatioQuery)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *SloV00Query) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetThreshold

`func (o *SloV00Query) GetThreshold() SloV00ThresholdQuery`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SloV00Query) GetThresholdOk() (*SloV00ThresholdQuery, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SloV00Query) SetThreshold(v SloV00ThresholdQuery)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *SloV00Query) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetType

`func (o *SloV00Query) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SloV00Query) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SloV00Query) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


