# ReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrillDownDashboardRef** | Pointer to [**DashboardRef**](DashboardRef.md) |  | [optional] 
**Provenance** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewReadOnly

`func NewReadOnly() *ReadOnly`

NewReadOnly instantiates a new ReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadOnlyWithDefaults

`func NewReadOnlyWithDefaults() *ReadOnly`

NewReadOnlyWithDefaults instantiates a new ReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrillDownDashboardRef

`func (o *ReadOnly) GetDrillDownDashboardRef() DashboardRef`

GetDrillDownDashboardRef returns the DrillDownDashboardRef field if non-nil, zero value otherwise.

### GetDrillDownDashboardRefOk

`func (o *ReadOnly) GetDrillDownDashboardRefOk() (*DashboardRef, bool)`

GetDrillDownDashboardRefOk returns a tuple with the DrillDownDashboardRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrillDownDashboardRef

`func (o *ReadOnly) SetDrillDownDashboardRef(v DashboardRef)`

SetDrillDownDashboardRef sets DrillDownDashboardRef field to given value.

### HasDrillDownDashboardRef

`func (o *ReadOnly) HasDrillDownDashboardRef() bool`

HasDrillDownDashboardRef returns a boolean if a field has been set.

### GetProvenance

`func (o *ReadOnly) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *ReadOnly) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *ReadOnly) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *ReadOnly) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetStatus

`func (o *ReadOnly) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReadOnly) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReadOnly) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReadOnly) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


