# SloV00ReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrillDownDashboardRef** | Pointer to [**SloV00DashboardRef**](SloV00DashboardRef.md) |  | [optional] 
**Provenance** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**SloV00Status**](SloV00Status.md) |  | [optional] 

## Methods

### NewSloV00ReadOnly

`func NewSloV00ReadOnly() *SloV00ReadOnly`

NewSloV00ReadOnly instantiates a new SloV00ReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00ReadOnlyWithDefaults

`func NewSloV00ReadOnlyWithDefaults() *SloV00ReadOnly`

NewSloV00ReadOnlyWithDefaults instantiates a new SloV00ReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrillDownDashboardRef

`func (o *SloV00ReadOnly) GetDrillDownDashboardRef() SloV00DashboardRef`

GetDrillDownDashboardRef returns the DrillDownDashboardRef field if non-nil, zero value otherwise.

### GetDrillDownDashboardRefOk

`func (o *SloV00ReadOnly) GetDrillDownDashboardRefOk() (*SloV00DashboardRef, bool)`

GetDrillDownDashboardRefOk returns a tuple with the DrillDownDashboardRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrillDownDashboardRef

`func (o *SloV00ReadOnly) SetDrillDownDashboardRef(v SloV00DashboardRef)`

SetDrillDownDashboardRef sets DrillDownDashboardRef field to given value.

### HasDrillDownDashboardRef

`func (o *SloV00ReadOnly) HasDrillDownDashboardRef() bool`

HasDrillDownDashboardRef returns a boolean if a field has been set.

### GetProvenance

`func (o *SloV00ReadOnly) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *SloV00ReadOnly) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *SloV00ReadOnly) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *SloV00ReadOnly) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetStatus

`func (o *SloV00ReadOnly) GetStatus() SloV00Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SloV00ReadOnly) GetStatusOk() (*SloV00Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SloV00ReadOnly) SetStatus(v SloV00Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SloV00ReadOnly) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


