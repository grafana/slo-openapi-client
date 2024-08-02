# SloV00FailureThresholdQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureThresholdExpression** | **string** |  | 
**GroupByLabels** | Pointer to **[]string** |  | [optional] 
**Threshold** | [**SloV00Threshold**](SloV00Threshold.md) |  | 

## Methods

### NewSloV00FailureThresholdQuery

`func NewSloV00FailureThresholdQuery(failureThresholdExpression string, threshold SloV00Threshold, ) *SloV00FailureThresholdQuery`

NewSloV00FailureThresholdQuery instantiates a new SloV00FailureThresholdQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00FailureThresholdQueryWithDefaults

`func NewSloV00FailureThresholdQueryWithDefaults() *SloV00FailureThresholdQuery`

NewSloV00FailureThresholdQueryWithDefaults instantiates a new SloV00FailureThresholdQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureThresholdExpression

`func (o *SloV00FailureThresholdQuery) GetFailureThresholdExpression() string`

GetFailureThresholdExpression returns the FailureThresholdExpression field if non-nil, zero value otherwise.

### GetFailureThresholdExpressionOk

`func (o *SloV00FailureThresholdQuery) GetFailureThresholdExpressionOk() (*string, bool)`

GetFailureThresholdExpressionOk returns a tuple with the FailureThresholdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThresholdExpression

`func (o *SloV00FailureThresholdQuery) SetFailureThresholdExpression(v string)`

SetFailureThresholdExpression sets FailureThresholdExpression field to given value.


### GetGroupByLabels

`func (o *SloV00FailureThresholdQuery) GetGroupByLabels() []string`

GetGroupByLabels returns the GroupByLabels field if non-nil, zero value otherwise.

### GetGroupByLabelsOk

`func (o *SloV00FailureThresholdQuery) GetGroupByLabelsOk() (*[]string, bool)`

GetGroupByLabelsOk returns a tuple with the GroupByLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByLabels

`func (o *SloV00FailureThresholdQuery) SetGroupByLabels(v []string)`

SetGroupByLabels sets GroupByLabels field to given value.

### HasGroupByLabels

`func (o *SloV00FailureThresholdQuery) HasGroupByLabels() bool`

HasGroupByLabels returns a boolean if a field has been set.

### GetThreshold

`func (o *SloV00FailureThresholdQuery) GetThreshold() SloV00Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SloV00FailureThresholdQuery) GetThresholdOk() (*SloV00Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SloV00FailureThresholdQuery) SetThreshold(v SloV00Threshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


