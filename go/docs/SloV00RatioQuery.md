# SloV00RatioQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**SuccessMetric** | [**SloV00MetricDef**](SloV00MetricDef.md) |  | 
**TotalMetric** | [**SloV00MetricDef**](SloV00MetricDef.md) |  | 

## Methods

### NewSloV00RatioQuery

`func NewSloV00RatioQuery(successMetric SloV00MetricDef, totalMetric SloV00MetricDef, ) *SloV00RatioQuery`

NewSloV00RatioQuery instantiates a new SloV00RatioQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00RatioQueryWithDefaults

`func NewSloV00RatioQueryWithDefaults() *SloV00RatioQuery`

NewSloV00RatioQueryWithDefaults instantiates a new SloV00RatioQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByLabels

`func (o *SloV00RatioQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *SloV00RatioQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *SloV00RatioQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *SloV00RatioQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetSuccessMetric

`func (o *SloV00RatioQuery) GetSuccessMetric() SloV00MetricDef`

GetSuccessMetric returns the SuccessMetric field if non-nil, zero value otherwise.

### GetSuccessMetricOk

`func (o *SloV00RatioQuery) GetSuccessMetricOk() (*SloV00MetricDef, bool)`

GetSuccessMetricOk returns a tuple with the SuccessMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMetric

`func (o *SloV00RatioQuery) SetSuccessMetric(v SloV00MetricDef)`

SetSuccessMetric sets SuccessMetric field to given value.


### GetTotalMetric

`func (o *SloV00RatioQuery) GetTotalMetric() SloV00MetricDef`

GetTotalMetric returns the TotalMetric field if non-nil, zero value otherwise.

### GetTotalMetricOk

`func (o *SloV00RatioQuery) GetTotalMetricOk() (*SloV00MetricDef, bool)`

GetTotalMetricOk returns a tuple with the TotalMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMetric

`func (o *SloV00RatioQuery) SetTotalMetric(v SloV00MetricDef)`

SetTotalMetric sets TotalMetric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


