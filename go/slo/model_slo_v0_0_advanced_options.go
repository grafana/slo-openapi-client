/*
Grafana SLO API

This API CRUDs SLO objects for the Grafana plugin.  Modifying an SLO object will create or update recording and alerting rules in a connected Prometheus instance and create or update dashboards in Grafana.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the SloV00AdvancedOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00AdvancedOptions{}

// SloV00AdvancedOptions struct for SloV00AdvancedOptions
type SloV00AdvancedOptions struct {
	MinFailures *int64 `json:"minFailures,omitempty"`
}

// NewSloV00AdvancedOptions instantiates a new SloV00AdvancedOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00AdvancedOptions() *SloV00AdvancedOptions {
	this := SloV00AdvancedOptions{}
	return &this
}

// NewSloV00AdvancedOptionsWithDefaults instantiates a new SloV00AdvancedOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00AdvancedOptionsWithDefaults() *SloV00AdvancedOptions {
	this := SloV00AdvancedOptions{}
	return &this
}

// GetMinFailures returns the MinFailures field value if set, zero value otherwise.
func (o *SloV00AdvancedOptions) GetMinFailures() int64 {
	if o == nil || IsNil(o.MinFailures) {
		var ret int64
		return ret
	}
	return *o.MinFailures
}

// GetMinFailuresOk returns a tuple with the MinFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00AdvancedOptions) GetMinFailuresOk() (*int64, bool) {
	if o == nil || IsNil(o.MinFailures) {
		return nil, false
	}
	return o.MinFailures, true
}

// HasMinFailures returns a boolean if a field has been set.
func (o *SloV00AdvancedOptions) HasMinFailures() bool {
	if o != nil && !IsNil(o.MinFailures) {
		return true
	}

	return false
}

// SetMinFailures gets a reference to the given int64 and assigns it to the MinFailures field.
func (o *SloV00AdvancedOptions) SetMinFailures(v int64) {
	o.MinFailures = &v
}

func (o SloV00AdvancedOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00AdvancedOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinFailures) {
		toSerialize["minFailures"] = o.MinFailures
	}
	return toSerialize, nil
}

type NullableSloV00AdvancedOptions struct {
	value *SloV00AdvancedOptions
	isSet bool
}

func (v NullableSloV00AdvancedOptions) Get() *SloV00AdvancedOptions {
	return v.value
}

func (v *NullableSloV00AdvancedOptions) Set(val *SloV00AdvancedOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00AdvancedOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00AdvancedOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00AdvancedOptions(val *SloV00AdvancedOptions) *NullableSloV00AdvancedOptions {
	return &NullableSloV00AdvancedOptions{value: val, isSet: true}
}

func (v NullableSloV00AdvancedOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00AdvancedOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
