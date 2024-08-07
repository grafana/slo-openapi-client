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

// checks if the SloV00MetricDef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00MetricDef{}

// SloV00MetricDef struct for SloV00MetricDef
type SloV00MetricDef struct {
	PrometheusMetric string  `json:"prometheusMetric"`
	Type             *string `json:"type,omitempty"`
}

type _SloV00MetricDef SloV00MetricDef

// NewSloV00MetricDef instantiates a new SloV00MetricDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00MetricDef(prometheusMetric string) *SloV00MetricDef {
	this := SloV00MetricDef{}
	this.PrometheusMetric = prometheusMetric
	return &this
}

// NewSloV00MetricDefWithDefaults instantiates a new SloV00MetricDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00MetricDefWithDefaults() *SloV00MetricDef {
	this := SloV00MetricDef{}
	return &this
}

// GetPrometheusMetric returns the PrometheusMetric field value
func (o *SloV00MetricDef) GetPrometheusMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrometheusMetric
}

// GetPrometheusMetricOk returns a tuple with the PrometheusMetric field value
// and a boolean to check if the value has been set.
func (o *SloV00MetricDef) GetPrometheusMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrometheusMetric, true
}

// SetPrometheusMetric sets field value
func (o *SloV00MetricDef) SetPrometheusMetric(v string) {
	o.PrometheusMetric = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SloV00MetricDef) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00MetricDef) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SloV00MetricDef) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SloV00MetricDef) SetType(v string) {
	o.Type = &v
}

func (o SloV00MetricDef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00MetricDef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prometheusMetric"] = o.PrometheusMetric
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *SloV00MetricDef) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varSloV00MetricDef := _SloV00MetricDef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSloV00MetricDef)

	if err != nil {
		return err
	}

	*o = SloV00MetricDef(varSloV00MetricDef)

	return err
}

type NullableSloV00MetricDef struct {
	value *SloV00MetricDef
	isSet bool
}

func (v NullableSloV00MetricDef) Get() *SloV00MetricDef {
	return v.value
}

func (v *NullableSloV00MetricDef) Set(val *SloV00MetricDef) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00MetricDef) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00MetricDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00MetricDef(val *SloV00MetricDef) *NullableSloV00MetricDef {
	return &NullableSloV00MetricDef{value: val, isSet: true}
}

func (v NullableSloV00MetricDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00MetricDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
