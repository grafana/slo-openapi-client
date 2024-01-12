# RatioQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**SuccessMetric** | [**MetricDef**](MetricDef.md) |  | 
**TotalMetric** | [**MetricDef**](MetricDef.md) |  | 

## Methods

### NewRatioQuery

`func NewRatioQuery(successMetric MetricDef, totalMetric MetricDef, ) *RatioQuery`

NewRatioQuery instantiates a new RatioQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatioQueryWithDefaults

`func NewRatioQueryWithDefaults() *RatioQuery`

NewRatioQueryWithDefaults instantiates a new RatioQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByLabels

`func (o *RatioQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *RatioQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *RatioQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *RatioQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetSuccessMetric

`func (o *RatioQuery) GetSuccessMetric() MetricDef`

GetSuccessMetric returns the SuccessMetric field if non-nil, zero value otherwise.

### GetSuccessMetricOk

`func (o *RatioQuery) GetSuccessMetricOk() (*MetricDef, bool)`

GetSuccessMetricOk returns a tuple with the SuccessMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMetric

`func (o *RatioQuery) SetSuccessMetric(v MetricDef)`

SetSuccessMetric sets SuccessMetric field to given value.


### GetTotalMetric

`func (o *RatioQuery) GetTotalMetric() MetricDef`

GetTotalMetric returns the TotalMetric field if non-nil, zero value otherwise.

### GetTotalMetricOk

`func (o *RatioQuery) GetTotalMetricOk() (*MetricDef, bool)`

GetTotalMetricOk returns a tuple with the TotalMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMetric

`func (o *RatioQuery) SetTotalMetric(v MetricDef)`

SetTotalMetric sets TotalMetric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


