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

// checks if the SloV00Objective type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00Objective{}

// SloV00Objective struct for SloV00Objective
type SloV00Objective struct {
	Value float64 `json:"value"`
	Window string `json:"window"`
}

type _SloV00Objective SloV00Objective

// NewSloV00Objective instantiates a new SloV00Objective object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00Objective(value float64, window string) *SloV00Objective {
	this := SloV00Objective{}
	this.Value = value
	this.Window = window
	return &this
}

// NewSloV00ObjectiveWithDefaults instantiates a new SloV00Objective object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00ObjectiveWithDefaults() *SloV00Objective {
	this := SloV00Objective{}
	return &this
}

// GetValue returns the Value field value
func (o *SloV00Objective) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SloV00Objective) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SloV00Objective) SetValue(v float64) {
	o.Value = v
}

// GetWindow returns the Window field value
func (o *SloV00Objective) GetWindow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *SloV00Objective) GetWindowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *SloV00Objective) SetWindow(v string) {
	o.Window = v
}

func (o SloV00Objective) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00Objective) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *SloV00Objective) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"window",
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

	varSloV00Objective := _SloV00Objective{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSloV00Objective)

	if err != nil {
		return err
	}

	*o = SloV00Objective(varSloV00Objective)

	return err
}

type NullableSloV00Objective struct {
	value *SloV00Objective
	isSet bool
}

func (v NullableSloV00Objective) Get() *SloV00Objective {
	return v.value
}

func (v *NullableSloV00Objective) Set(val *SloV00Objective) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00Objective) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00Objective) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00Objective(val *SloV00Objective) *NullableSloV00Objective {
	return &NullableSloV00Objective{value: val, isSet: true}
}

func (v NullableSloV00Objective) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00Objective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


