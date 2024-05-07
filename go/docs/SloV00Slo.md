# SloV00Slo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerting** | Pointer to [**SloV00Alerting**](SloV00Alerting.md) |  | [optional] 
**Description** | **string** |  | 
**DestinationDatasource** | Pointer to [**SloV00DestinationDatasource**](SloV00DestinationDatasource.md) |  | [optional] 
**Folder** | Pointer to [**SloV00Folder**](SloV00Folder.md) |  | [optional] 
**Labels** | Pointer to [**[]SloV00Label**](SloV00Label.md) |  | [optional] 
**Name** | **string** |  | 
**Objectives** | [**[]SloV00Objective**](SloV00Objective.md) |  | 
**Query** | [**SloV00Query**](SloV00Query.md) |  | 
**ReadOnly** | Pointer to [**SloV00ReadOnly**](SloV00ReadOnly.md) |  | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewSloV00Slo

`func NewSloV00Slo(description string, name string, objectives []SloV00Objective, query SloV00Query, uuid string, ) *SloV00Slo`

NewSloV00Slo instantiates a new SloV00Slo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloV00SloWithDefaults

`func NewSloV00SloWithDefaults() *SloV00Slo`

NewSloV00SloWithDefaults instantiates a new SloV00Slo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerting

`func (o *SloV00Slo) GetAlerting() SloV00Alerting`

GetAlerting returns the Alerting field if non-nil, zero value otherwise.

### GetAlertingOk

`func (o *SloV00Slo) GetAlertingOk() (*SloV00Alerting, bool)`

GetAlertingOk returns a tuple with the Alerting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerting

`func (o *SloV00Slo) SetAlerting(v SloV00Alerting)`

SetAlerting sets Alerting field to given value.

### HasAlerting

`func (o *SloV00Slo) HasAlerting() bool`

HasAlerting returns a boolean if a field has been set.

### GetDescription

`func (o *SloV00Slo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SloV00Slo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SloV00Slo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationDatasource

`func (o *SloV00Slo) GetDestinationDatasource() SloV00DestinationDatasource`

GetDestinationDatasource returns the DestinationDatasource field if non-nil, zero value otherwise.

### GetDestinationDatasourceOk

`func (o *SloV00Slo) GetDestinationDatasourceOk() (*SloV00DestinationDatasource, bool)`

GetDestinationDatasourceOk returns a tuple with the DestinationDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDatasource

`func (o *SloV00Slo) SetDestinationDatasource(v SloV00DestinationDatasource)`

SetDestinationDatasource sets DestinationDatasource field to given value.

### HasDestinationDatasource

`func (o *SloV00Slo) HasDestinationDatasource() bool`

HasDestinationDatasource returns a boolean if a field has been set.

### GetFolder

`func (o *SloV00Slo) GetFolder() SloV00Folder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *SloV00Slo) GetFolderOk() (*SloV00Folder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *SloV00Slo) SetFolder(v SloV00Folder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *SloV00Slo) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetLabels

`func (o *SloV00Slo) GetLabels() []SloV00Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SloV00Slo) GetLabelsOk() (*[]SloV00Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SloV00Slo) SetLabels(v []SloV00Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SloV00Slo) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *SloV00Slo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SloV00Slo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SloV00Slo) SetName(v string)`

SetName sets Name field to given value.


### GetObjectives

`func (o *SloV00Slo) GetObjectives() []SloV00Objective`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *SloV00Slo) GetObjectivesOk() (*[]SloV00Objective, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *SloV00Slo) SetObjectives(v []SloV00Objective)`

SetObjectives sets Objectives field to given value.


### GetQuery

`func (o *SloV00Slo) GetQuery() SloV00Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SloV00Slo) GetQueryOk() (*SloV00Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SloV00Slo) SetQuery(v SloV00Query)`

SetQuery sets Query field to given value.


### GetReadOnly

`func (o *SloV00Slo) GetReadOnly() SloV00ReadOnly`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *SloV00Slo) GetReadOnlyOk() (*SloV00ReadOnly, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *SloV00Slo) SetReadOnly(v SloV00ReadOnly)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *SloV00Slo) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetUuid

`func (o *SloV00Slo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SloV00Slo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SloV00Slo) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


