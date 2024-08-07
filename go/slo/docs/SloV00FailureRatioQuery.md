# SloV00FailureRatioQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureMetric** | [**SloV00MetricDef**](SloV00MetricDef.md) |  | 
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**TotalMetric** | [**SloV00MetricDef**](SloV00MetricDef.md) |  | 

## Methods

### NewSloV00FailureRatioQuery

`func NewSloV00FailureRatioQuery(failureMetric SloV00MetricDef, totalMetric SloV00MetricDef, ) *SloV00FailureRatioQuery`

NewSloV00FailureRatioQuery instantiates a new SloV00FailureRatioQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00FailureRatioQueryWithDefaults

`func NewSloV00FailureRatioQueryWithDefaults() *SloV00FailureRatioQuery`

NewSloV00FailureRatioQueryWithDefaults instantiates a new SloV00FailureRatioQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureMetric

`func (o *SloV00FailureRatioQuery) GetFailureMetric() SloV00MetricDef`

GetFailureMetric returns the FailureMetric field if non-nil, zero value otherwise.

### GetFailureMetricOk

`func (o *SloV00FailureRatioQuery) GetFailureMetricOk() (*SloV00MetricDef, bool)`

GetFailureMetricOk returns a tuple with the FailureMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMetric

`func (o *SloV00FailureRatioQuery) SetFailureMetric(v SloV00MetricDef)`

SetFailureMetric sets FailureMetric field to given value.


### GetGroupByLabels

`func (o *SloV00FailureRatioQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *SloV00FailureRatioQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *SloV00FailureRatioQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *SloV00FailureRatioQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetTotalMetric

`func (o *SloV00FailureRatioQuery) GetTotalMetric() SloV00MetricDef`

GetTotalMetric returns the TotalMetric field if non-nil, zero value otherwise.

### GetTotalMetricOk

`func (o *SloV00FailureRatioQuery) GetTotalMetricOk() (*SloV00MetricDef, bool)`

GetTotalMetricOk returns a tuple with the TotalMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMetric

`func (o *SloV00FailureRatioQuery) SetTotalMetric(v SloV00MetricDef)`

SetTotalMetric sets TotalMetric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


