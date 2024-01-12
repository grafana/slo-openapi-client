# ThresholdQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**Metric** | [**MetricDef**](MetricDef.md) |  | 
**Threshold** | [**Threshold**](Threshold.md) |  | 

## Methods

### NewThresholdQuery

`func NewThresholdQuery(metric MetricDef, threshold Threshold, ) *ThresholdQuery`

NewThresholdQuery instantiates a new ThresholdQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdQueryWithDefaults

`func NewThresholdQueryWithDefaults() *ThresholdQuery`

NewThresholdQueryWithDefaults instantiates a new ThresholdQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByLabels

`func (o *ThresholdQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *ThresholdQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *ThresholdQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *ThresholdQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetMetric

`func (o *ThresholdQuery) GetMetric() MetricDef`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ThresholdQuery) GetMetricOk() (*MetricDef, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ThresholdQuery) SetMetric(v MetricDef)`

SetMetric sets Metric field to given value.


### GetThreshold

`func (o *ThresholdQuery) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ThresholdQuery) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ThresholdQuery) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


