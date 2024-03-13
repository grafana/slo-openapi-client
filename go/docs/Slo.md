# Slo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerting** | Pointer to [**Alerting**](Alerting.md) |  | [optional] 
**Description** | **string** |  | 
**DestinationDatasource** | Pointer to [**DestinationDatasource**](DestinationDatasource.md) |  | [optional] 
**Folder** | Pointer to [**Folder**](Folder.md) |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Name** | **string** |  | 
**Objectives** | [**[]Objective**](Objective.md) |  | 
**Query** | [**Query**](Query.md) |  | 
**ReadOnly** | Pointer to [**ReadOnly**](ReadOnly.md) |  | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewSlo

`func NewSlo(description string, name string, objectives []Objective, query Query, uuid string, ) *Slo`

NewSlo instantiates a new Slo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloWithDefaults

`func NewSloWithDefaults() *Slo`

NewSloWithDefaults instantiates a new Slo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerting

`func (o *Slo) GetAlerting() Alerting`

GetAlerting returns the Alerting field if non-nil, zero value otherwise.

### GetAlertingOk

`func (o *Slo) GetAlertingOk() (*Alerting, bool)`

GetAlertingOk returns a tuple with the Alerting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerting

`func (o *Slo) SetAlerting(v Alerting)`

SetAlerting sets Alerting field to given value.

### HasAlerting

`func (o *Slo) HasAlerting() bool`

HasAlerting returns a boolean if a field has been set.

### GetDescription

`func (o *Slo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Slo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Slo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationDatasource

`func (o *Slo) GetDestinationDatasource() DestinationDatasource`

GetDestinationDatasource returns the DestinationDatasource field if non-nil, zero value otherwise.

### GetDestinationDatasourceOk

`func (o *Slo) GetDestinationDatasourceOk() (*DestinationDatasource, bool)`

GetDestinationDatasourceOk returns a tuple with the DestinationDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDatasource

`func (o *Slo) SetDestinationDatasource(v DestinationDatasource)`

SetDestinationDatasource sets DestinationDatasource field to given value.

### HasDestinationDatasource

`func (o *Slo) HasDestinationDatasource() bool`

HasDestinationDatasource returns a boolean if a field has been set.

### GetFolder

`func (o *Slo) GetFolder() Folder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *Slo) GetFolderOk() (*Folder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *Slo) SetFolder(v Folder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *Slo) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetLabels

`func (o *Slo) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Slo) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Slo) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Slo) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *Slo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Slo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Slo) SetName(v string)`

SetName sets Name field to given value.


### GetObjectives

`func (o *Slo) GetObjectives() []Objective`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *Slo) GetObjectivesOk() (*[]Objective, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *Slo) SetObjectives(v []Objective)`

SetObjectives sets Objectives field to given value.


### GetQuery

`func (o *Slo) GetQuery() Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Slo) GetQueryOk() (*Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Slo) SetQuery(v Query)`

SetQuery sets Query field to given value.


### GetReadOnly

`func (o *Slo) GetReadOnly() ReadOnly`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *Slo) GetReadOnlyOk() (*ReadOnly, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *Slo) SetReadOnly(v ReadOnly)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *Slo) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetUuid

`func (o *Slo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Slo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Slo) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


