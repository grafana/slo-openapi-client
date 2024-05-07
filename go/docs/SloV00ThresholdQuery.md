# SloV00ThresholdQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**Metric** | [**SloV00MetricDef**](SloV00MetricDef.md) |  | 
**Threshold** | [**SloV00Threshold**](SloV00Threshold.md) |  | 

## Methods

### NewSloV00ThresholdQuery

`func NewSloV00ThresholdQuery(metric SloV00MetricDef, threshold SloV00Threshold, ) *SloV00ThresholdQuery`

NewSloV00ThresholdQuery instantiates a new SloV00ThresholdQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00ThresholdQueryWithDefaults

`func NewSloV00ThresholdQueryWithDefaults() *SloV00ThresholdQuery`

NewSloV00ThresholdQueryWithDefaults instantiates a new SloV00ThresholdQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByLabels

`func (o *SloV00ThresholdQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *SloV00ThresholdQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *SloV00ThresholdQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *SloV00ThresholdQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetMetric

`func (o *SloV00ThresholdQuery) GetMetric() SloV00MetricDef`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SloV00ThresholdQuery) GetMetricOk() (*SloV00MetricDef, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SloV00ThresholdQuery) SetMetric(v SloV00MetricDef)`

SetMetric sets Metric field to given value.


### GetThreshold

`func (o *SloV00ThresholdQuery) GetThreshold() SloV00Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SloV00ThresholdQuery) GetThresholdOk() (*SloV00Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SloV00ThresholdQuery) SetThreshold(v SloV00Threshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

