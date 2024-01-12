# HistogramQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**Metric** | [**MetricDef**](MetricDef.md) |  | 
**Percentile** | **float64** |  | 
**Threshold** | [**Threshold**](Threshold.md) |  | 

## Methods

### NewHistogramQuery

`func NewHistogramQuery(metric MetricDef, percentile float64, threshold Threshold, ) *HistogramQuery`

NewHistogramQuery instantiates a new HistogramQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistogramQueryWithDefaults

`func NewHistogramQueryWithDefaults() *HistogramQuery`

NewHistogramQueryWithDefaults instantiates a new HistogramQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByLabels

`func (o *HistogramQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *HistogramQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *HistogramQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *HistogramQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetMetric

`func (o *HistogramQuery) GetMetric() MetricDef`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *HistogramQuery) GetMetricOk() (*MetricDef, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *HistogramQuery) SetMetric(v MetricDef)`

SetMetric sets Metric field to given value.


### GetPercentile

`func (o *HistogramQuery) GetPercentile() float64`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *HistogramQuery) GetPercentileOk() (*float64, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *HistogramQuery) SetPercentile(v float64)`

SetPercentile sets Percentile field to given value.


### GetThreshold

`func (o *HistogramQuery) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *HistogramQuery) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *HistogramQuery) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


