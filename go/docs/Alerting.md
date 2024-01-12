# Alerting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**FastBurn** | Pointer to [**AlertingMetadata**](AlertingMetadata.md) |  | [optional] 
**Labels** | Pointer to **interface{}** |  | [optional] 
**SlowBurn** | Pointer to [**AlertingMetadata**](AlertingMetadata.md) |  | [optional] 

## Methods

### NewAlerting

`func NewAlerting() *Alerting`

NewAlerting instantiates a new Alerting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingWithDefaults

`func NewAlertingWithDefaults() *Alerting`

NewAlertingWithDefaults instantiates a new Alerting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *Alerting) GetAnnotations() []Label`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Alerting) GetAnnotationsOk() (*[]Label, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Alerting) SetAnnotations(v []Label)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Alerting) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFastBurn

`func (o *Alerting) GetFastBurn() AlertingMetadata`

GetFastBurn returns the FastBurn field if non-nil, zero value otherwise.

### GetFastBurnOk

`func (o *Alerting) GetFastBurnOk() (*AlertingMetadata, bool)`

GetFastBurnOk returns a tuple with the FastBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastBurn

`func (o *Alerting) SetFastBurn(v AlertingMetadata)`

SetFastBurn sets FastBurn field to given value.

### HasFastBurn

`func (o *Alerting) HasFastBurn() bool`

HasFastBurn returns a boolean if a field has been set.

### GetLabels

`func (o *Alerting) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Alerting) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Alerting) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Alerting) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *Alerting) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Alerting) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetSlowBurn

`func (o *Alerting) GetSlowBurn() AlertingMetadata`

GetSlowBurn returns the SlowBurn field if non-nil, zero value otherwise.

### GetSlowBurnOk

`func (o *Alerting) GetSlowBurnOk() (*AlertingMetadata, bool)`

GetSlowBurnOk returns a tuple with the SlowBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowBurn

`func (o *Alerting) SetSlowBurn(v AlertingMetadata)`

SetSlowBurn sets SlowBurn field to given value.

### HasSlowBurn

`func (o *Alerting) HasSlowBurn() bool`

HasSlowBurn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


