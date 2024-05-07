# SloV00HistogramQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**Metric** | [**SloV00MetricDef**](SloV00MetricDef.md) |  | 
**Percentile** | **float64** |  | 
**Threshold** | [**SloV00Threshold**](SloV00Threshold.md) |  | 

## Methods

### NewSloV00HistogramQuery

`func NewSloV00HistogramQuery(metric SloV00MetricDef, percentile float64, threshold SloV00Threshold, ) *SloV00HistogramQuery`

NewSloV00HistogramQuery instantiates a new SloV00HistogramQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00HistogramQueryWithDefaults

`func NewSloV00HistogramQueryWithDefaults() *SloV00HistogramQuery`

NewSloV00HistogramQueryWithDefaults instantiates a new SloV00HistogramQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByLabels

`func (o *SloV00HistogramQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *SloV00HistogramQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *SloV00HistogramQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *SloV00HistogramQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetMetric

`func (o *SloV00HistogramQuery) GetMetric() SloV00MetricDef`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SloV00HistogramQuery) GetMetricOk() (*SloV00MetricDef, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SloV00HistogramQuery) SetMetric(v SloV00MetricDef)`

SetMetric sets Metric field to given value.


### GetPercentile

`func (o *SloV00HistogramQuery) GetPercentile() float64`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *SloV00HistogramQuery) GetPercentileOk() (*float64, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *SloV00HistogramQuery) SetPercentile(v float64)`

SetPercentile sets Percentile field to given value.


### GetThreshold

`func (o *SloV00HistogramQuery) GetThreshold() SloV00Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SloV00HistogramQuery) GetThresholdOk() (*SloV00Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SloV00HistogramQuery) SetThreshold(v SloV00Threshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


