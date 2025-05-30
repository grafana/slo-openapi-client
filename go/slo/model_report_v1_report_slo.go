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

// checks if the ReportV1ReportSlo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportV1ReportSlo{}

// ReportV1ReportSlo struct for ReportV1ReportSlo
type ReportV1ReportSlo struct {
	SloUuid string `json:"sloUuid"`
}

type _ReportV1ReportSlo ReportV1ReportSlo

// NewReportV1ReportSlo instantiates a new ReportV1ReportSlo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportV1ReportSlo(sloUuid string) *ReportV1ReportSlo {
	this := ReportV1ReportSlo{}
	this.SloUuid = sloUuid
	return &this
}

// NewReportV1ReportSloWithDefaults instantiates a new ReportV1ReportSlo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportV1ReportSloWithDefaults() *ReportV1ReportSlo {
	this := ReportV1ReportSlo{}
	return &this
}

// GetSloUuid returns the SloUuid field value
func (o *ReportV1ReportSlo) GetSloUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SloUuid
}

// GetSloUuidOk returns a tuple with the SloUuid field value
// and a boolean to check if the value has been set.
func (o *ReportV1ReportSlo) GetSloUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SloUuid, true
}

// SetSloUuid sets field value
func (o *ReportV1ReportSlo) SetSloUuid(v string) {
	o.SloUuid = v
}

func (o ReportV1ReportSlo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportV1ReportSlo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sloUuid"] = o.SloUuid
	return toSerialize, nil
}

func (o *ReportV1ReportSlo) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varReportV1ReportSlo := _ReportV1ReportSlo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varReportV1ReportSlo)

	if err != nil {
		return err
	}

	*o = ReportV1ReportSlo(varReportV1ReportSlo)

	return err
}

type NullableReportV1ReportSlo struct {
	value *ReportV1ReportSlo
	isSet bool
}

func (v NullableReportV1ReportSlo) Get() *ReportV1ReportSlo {
	return v.value
}

func (v *NullableReportV1ReportSlo) Set(val *ReportV1ReportSlo) {
	v.value = val
	v.isSet = true
}

func (v NullableReportV1ReportSlo) IsSet() bool {
	return v.isSet
}

func (v *NullableReportV1ReportSlo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportV1ReportSlo(val *ReportV1ReportSlo) *NullableReportV1ReportSlo {
	return &NullableReportV1ReportSlo{value: val, isSet: true}
}

func (v NullableReportV1ReportSlo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportV1ReportSlo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
