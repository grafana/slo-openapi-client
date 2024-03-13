/*
Grafana SLO API

This API CRUDs SLO objects for the Grafana plugin.  Modifying an SLO object will create or update recording and alerting rules in a connected Prometheus instance and create or update dashboards in Grafana.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Slo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Slo{}

// Slo struct for Slo
type Slo struct {
	Alerting *Alerting `json:"alerting,omitempty"`
	Description string `json:"description"`
	DestinationDatasource *DestinationDatasource `json:"destinationDatasource,omitempty"`
	Folder *Folder `json:"folder,omitempty"`
	Labels []Label `json:"labels,omitempty"`
	Name string `json:"name"`
	Objectives []Objective `json:"objectives"`
	Query Query `json:"query"`
	ReadOnly *ReadOnly `json:"readOnly,omitempty"`
	Uuid string `json:"uuid"`
}

type _Slo Slo

// NewSlo instantiates a new Slo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlo(description string, name string, objectives []Objective, query Query, uuid string) *Slo {
	this := Slo{}
	this.Description = description
	this.Name = name
	this.Objectives = objectives
	this.Query = query
	this.Uuid = uuid
	return &this
}

// NewSloWithDefaults instantiates a new Slo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloWithDefaults() *Slo {
	this := Slo{}
	return &this
}

// GetAlerting returns the Alerting field value if set, zero value otherwise.
func (o *Slo) GetAlerting() Alerting {
	if o == nil || IsNil(o.Alerting) {
		var ret Alerting
		return ret
	}
	return *o.Alerting
}

// GetAlertingOk returns a tuple with the Alerting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slo) GetAlertingOk() (*Alerting, bool) {
	if o == nil || IsNil(o.Alerting) {
		return nil, false
	}
	return o.Alerting, true
}

// HasAlerting returns a boolean if a field has been set.
func (o *Slo) HasAlerting() bool {
	if o != nil && !IsNil(o.Alerting) {
		return true
	}

	return false
}

// SetAlerting gets a reference to the given Alerting and assigns it to the Alerting field.
func (o *Slo) SetAlerting(v Alerting) {
	o.Alerting = &v
}

// GetDescription returns the Description field value
func (o *Slo) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Slo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Slo) SetDescription(v string) {
	o.Description = v
}

// GetDestinationDatasource returns the DestinationDatasource field value if set, zero value otherwise.
func (o *Slo) GetDestinationDatasource() DestinationDatasource {
	if o == nil || IsNil(o.DestinationDatasource) {
		var ret DestinationDatasource
		return ret
	}
	return *o.DestinationDatasource
}

// GetDestinationDatasourceOk returns a tuple with the DestinationDatasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slo) GetDestinationDatasourceOk() (*DestinationDatasource, bool) {
	if o == nil || IsNil(o.DestinationDatasource) {
		return nil, false
	}
	return o.DestinationDatasource, true
}

// HasDestinationDatasource returns a boolean if a field has been set.
func (o *Slo) HasDestinationDatasource() bool {
	if o != nil && !IsNil(o.DestinationDatasource) {
		return true
	}

	return false
}

// SetDestinationDatasource gets a reference to the given DestinationDatasource and assigns it to the DestinationDatasource field.
func (o *Slo) SetDestinationDatasource(v DestinationDatasource) {
	o.DestinationDatasource = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *Slo) GetFolder() Folder {
	if o == nil || IsNil(o.Folder) {
		var ret Folder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slo) GetFolderOk() (*Folder, bool) {
	if o == nil || IsNil(o.Folder) {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *Slo) HasFolder() bool {
	if o != nil && !IsNil(o.Folder) {
		return true
	}

	return false
}

// SetFolder gets a reference to the given Folder and assigns it to the Folder field.
func (o *Slo) SetFolder(v Folder) {
	o.Folder = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Slo) GetLabels() []Label {
	if o == nil || IsNil(o.Labels) {
		var ret []Label
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slo) GetLabelsOk() ([]Label, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Slo) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *Slo) SetLabels(v []Label) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *Slo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Slo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Slo) SetName(v string) {
	o.Name = v
}

// GetObjectives returns the Objectives field value
func (o *Slo) GetObjectives() []Objective {
	if o == nil {
		var ret []Objective
		return ret
	}

	return o.Objectives
}

// GetObjectivesOk returns a tuple with the Objectives field value
// and a boolean to check if the value has been set.
func (o *Slo) GetObjectivesOk() ([]Objective, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objectives, true
}

// SetObjectives sets field value
func (o *Slo) SetObjectives(v []Objective) {
	o.Objectives = v
}

// GetQuery returns the Query field value
func (o *Slo) GetQuery() Query {
	if o == nil {
		var ret Query
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Slo) GetQueryOk() (*Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Slo) SetQuery(v Query) {
	o.Query = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *Slo) GetReadOnly() ReadOnly {
	if o == nil || IsNil(o.ReadOnly) {
		var ret ReadOnly
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slo) GetReadOnlyOk() (*ReadOnly, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *Slo) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given ReadOnly and assigns it to the ReadOnly field.
func (o *Slo) SetReadOnly(v ReadOnly) {
	o.ReadOnly = &v
}

// GetUuid returns the Uuid field value
func (o *Slo) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Slo) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Slo) SetUuid(v string) {
	o.Uuid = v
}

func (o Slo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Slo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerting) {
		toSerialize["alerting"] = o.Alerting
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.DestinationDatasource) {
		toSerialize["destinationDatasource"] = o.DestinationDatasource
	}
	if !IsNil(o.Folder) {
		toSerialize["folder"] = o.Folder
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	toSerialize["objectives"] = o.Objectives
	toSerialize["query"] = o.Query
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *Slo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
		"objectives",
		"query",
		"uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSlo := _Slo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlo)

	if err != nil {
		return err
	}

	*o = Slo(varSlo)

	return err
}

type NullableSlo struct {
	value *Slo
	isSet bool
}

func (v NullableSlo) Get() *Slo {
	return v.value
}

func (v *NullableSlo) Set(val *Slo) {
	v.value = val
	v.isSet = true
}

func (v NullableSlo) IsSet() bool {
	return v.isSet
}

func (v *NullableSlo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlo(val *Slo) *NullableSlo {
	return &NullableSlo{value: val, isSet: true}
}

func (v NullableSlo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


