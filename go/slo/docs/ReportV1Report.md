# ReportV1Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Labels** | Pointer to [**[]ReportV1Label**](ReportV1Label.md) |  | [optional] 
**Name** | **string** |  | 
**ReportDefinition** | [**ReportV1ReportDefinitionSloList**](ReportV1ReportDefinitionSloList.md) |  | 
**TimeSpan** | **string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewReportV1Report

`func NewReportV1Report(description string, name string, reportDefinition ReportV1ReportDefinitionSloList, timeSpan string, uuid string, ) *ReportV1Report`

NewReportV1Report instantiates a new ReportV1Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportV1ReportWithDefaults

`func NewReportV1ReportWithDefaults() *ReportV1Report`

NewReportV1ReportWithDefaults instantiates a new ReportV1Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReportV1Report) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportV1Report) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportV1Report) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabels

`func (o *ReportV1Report) GetLabels() []ReportV1Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ReportV1Report) GetLabelsOk() (*[]ReportV1Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ReportV1Report) SetLabels(v []ReportV1Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ReportV1Report) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *ReportV1Report) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportV1Report) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportV1Report) SetName(v string)`

SetName sets Name field to given value.


### GetReportDefinition

`func (o *ReportV1Report) GetReportDefinition() ReportV1ReportDefinitionSloList`

GetReportDefinition returns the ReportDefinition field if non-nil, zero value otherwise.

### GetReportDefinitionOk

`func (o *ReportV1Report) GetReportDefinitionOk() (*ReportV1ReportDefinitionSloList, bool)`

GetReportDefinitionOk returns a tuple with the ReportDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDefinition

`func (o *ReportV1Report) SetReportDefinition(v ReportV1ReportDefinitionSloList)`

SetReportDefinition sets ReportDefinition field to given value.


### GetTimeSpan

`func (o *ReportV1Report) GetTimeSpan() string`

GetTimeSpan returns the TimeSpan field if non-nil, zero value otherwise.

### GetTimeSpanOk

`func (o *ReportV1Report) GetTimeSpanOk() (*string, bool)`

GetTimeSpanOk returns a tuple with the TimeSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpan

`func (o *ReportV1Report) SetTimeSpan(v string)`

SetTimeSpan sets TimeSpan field to given value.


### GetUuid

`func (o *ReportV1Report) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ReportV1Report) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ReportV1Report) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


