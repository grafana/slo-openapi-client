# AlertingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Labels** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAlertingMetadata

`func NewAlertingMetadata() *AlertingMetadata`

NewAlertingMetadata instantiates a new AlertingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingMetadataWithDefaults

`func NewAlertingMetadataWithDefaults() *AlertingMetadata`

NewAlertingMetadataWithDefaults instantiates a new AlertingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *AlertingMetadata) GetAnnotations() []Label`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AlertingMetadata) GetAnnotationsOk() (*[]Label, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AlertingMetadata) SetAnnotations(v []Label)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AlertingMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLabels

`func (o *AlertingMetadata) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertingMetadata) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertingMetadata) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AlertingMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AlertingMetadata) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AlertingMetadata) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


