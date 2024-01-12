# MetricDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrometheusMetric** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricDef

`func NewMetricDef(prometheusMetric string, ) *MetricDef`

NewMetricDef instantiates a new MetricDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDefWithDefaults

`func NewMetricDefWithDefaults() *MetricDef`

NewMetricDefWithDefaults instantiates a new MetricDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrometheusMetric

`func (o *MetricDef) GetPrometheusMetric() string`

GetPrometheusMetric returns the PrometheusMetric field if non-nil, zero value otherwise.

### GetPrometheusMetricOk

`func (o *MetricDef) GetPrometheusMetricOk() (*string, bool)`

GetPrometheusMetricOk returns a tuple with the PrometheusMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusMetric

`func (o *MetricDef) SetPrometheusMetric(v string)`

SetPrometheusMetric sets PrometheusMetric field to given value.


### GetType

`func (o *MetricDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricDef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricDef) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetricDef) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


