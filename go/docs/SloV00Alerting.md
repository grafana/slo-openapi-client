# SloV00Alerting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to [**[]SloV00Label**](SloV00Label.md) |  | [optional] 
**FastBurn** | Pointer to [**SloV00AlertingMetadata**](SloV00AlertingMetadata.md) |  | [optional] 
**Labels** | Pointer to **interface{}** |  | [optional] 
**SlowBurn** | Pointer to [**SloV00AlertingMetadata**](SloV00AlertingMetadata.md) |  | [optional] 

## Methods

### NewSloV00Alerting

`func NewSloV00Alerting() *SloV00Alerting`

NewSloV00Alerting instantiates a new SloV00Alerting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00AlertingWithDefaults

`func NewSloV00AlertingWithDefaults() *SloV00Alerting`

NewSloV00AlertingWithDefaults instantiates a new SloV00Alerting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *SloV00Alerting) GetAnnotations() []SloV00Label`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SloV00Alerting) GetAnnotationsOk() (*[]SloV00Label, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SloV00Alerting) SetAnnotations(v []SloV00Label)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SloV00Alerting) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFastBurn

`func (o *SloV00Alerting) GetFastBurn() SloV00AlertingMetadata`

GetFastBurn returns the FastBurn field if non-nil, zero value otherwise.

### GetFastBurnOk

`func (o *SloV00Alerting) GetFastBurnOk() (*SloV00AlertingMetadata, bool)`

GetFastBurnOk returns a tuple with the FastBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastBurn

`func (o *SloV00Alerting) SetFastBurn(v SloV00AlertingMetadata)`

SetFastBurn sets FastBurn field to given value.

### HasFastBurn

`func (o *SloV00Alerting) HasFastBurn() bool`

HasFastBurn returns a boolean if a field has been set.

### GetLabels

`func (o *SloV00Alerting) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SloV00Alerting) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SloV00Alerting) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SloV00Alerting) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *SloV00Alerting) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *SloV00Alerting) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetSlowBurn

`func (o *SloV00Alerting) GetSlowBurn() SloV00AlertingMetadata`

GetSlowBurn returns the SlowBurn field if non-nil, zero value otherwise.

### GetSlowBurnOk

`func (o *SloV00Alerting) GetSlowBurnOk() (*SloV00AlertingMetadata, bool)`

GetSlowBurnOk returns a tuple with the SlowBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowBurn

`func (o *SloV00Alerting) SetSlowBurn(v SloV00AlertingMetadata)`

SetSlowBurn sets SlowBurn field to given value.

### HasSlowBurn

`func (o *SloV00Alerting) HasSlowBurn() bool`

HasSlowBurn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

