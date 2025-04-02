/*
Grafana SLO API

This API CRUDs SLO objects for the Grafana plugin.  Modifying an SLO object will create or update recording and alerting rules in a connected Prometheus instance and create or update dashboards in Grafana.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"bytes"
	"encoding/json"
)

// checks if the ReportV1ReportDefinitionSloList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportV1ReportDefinitionSloList{}

// ReportV1ReportDefinitionSloList struct for ReportV1ReportDefinitionSloList
type ReportV1ReportDefinitionSloList struct {
	Slos []ReportV1ReportSlo `json:"slos"`
}

type _ReportV1ReportDefinitionSloList ReportV1ReportDefinitionSloList

// NewReportV1ReportDefinitionSloList instantiates a new ReportV1ReportDefinitionSloList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportV1ReportDefinitionSloList(slos []ReportV1ReportSlo) *ReportV1ReportDefinitionSloList {
	this := ReportV1ReportDefinitionSloList{}
	this.Slos = slos
	return &this
}

// NewReportV1ReportDefinitionSloListWithDefaults instantiates a new ReportV1ReportDefinitionSloList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportV1ReportDefinitionSloListWithDefaults() *ReportV1ReportDefinitionSloList {
	this := ReportV1ReportDefinitionSloList{}
	return &this
}

// GetSlos returns the Slos field value
func (o *ReportV1ReportDefinitionSloList) GetSlos() []ReportV1ReportSlo {
	if o == nil {
		var ret []ReportV1ReportSlo
		return ret
	}

	return o.Slos
}

// GetSlosOk returns a tuple with the Slos field value
// and a boolean to check if the value has been set.
func (o *ReportV1ReportDefinitionSloList) GetSlosOk() ([]ReportV1ReportSlo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slos, true
}

// SetSlos sets field value
func (o *ReportV1ReportDefinitionSloList) SetSlos(v []ReportV1ReportSlo) {
	o.Slos = v
}

func (o ReportV1ReportDefinitionSloList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportV1ReportDefinitionSloList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slos"] = o.Slos
	return toSerialize, nil
}

func (o *ReportV1ReportDefinitionSloList) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varReportV1ReportDefinitionSloList := _ReportV1ReportDefinitionSloList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varReportV1ReportDefinitionSloList)

	if err != nil {
		return err
	}

	*o = ReportV1ReportDefinitionSloList(varReportV1ReportDefinitionSloList)

	return err
}

type NullableReportV1ReportDefinitionSloList struct {
	value *ReportV1ReportDefinitionSloList
	isSet bool
}

func (v NullableReportV1ReportDefinitionSloList) Get() *ReportV1ReportDefinitionSloList {
	return v.value
}

func (v *NullableReportV1ReportDefinitionSloList) Set(val *ReportV1ReportDefinitionSloList) {
	v.value = val
	v.isSet = true
}

func (v NullableReportV1ReportDefinitionSloList) IsSet() bool {
	return v.isSet
}

func (v *NullableReportV1ReportDefinitionSloList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportV1ReportDefinitionSloList(val *ReportV1ReportDefinitionSloList) *NullableReportV1ReportDefinitionSloList {
	return &NullableReportV1ReportDefinitionSloList{value: val, isSet: true}
}

func (v NullableReportV1ReportDefinitionSloList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportV1ReportDefinitionSloList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
