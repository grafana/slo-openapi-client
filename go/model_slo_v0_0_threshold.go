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

// checks if the SloV00Threshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00Threshold{}

// SloV00Threshold struct for SloV00Threshold
type SloV00Threshold struct {
	Operator string `json:"operator"`
	Value float64 `json:"value"`
}

type _SloV00Threshold SloV00Threshold

// NewSloV00Threshold instantiates a new SloV00Threshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00Threshold(operator string, value float64) *SloV00Threshold {
	this := SloV00Threshold{}
	this.Operator = operator
	this.Value = value
	return &this
}

// NewSloV00ThresholdWithDefaults instantiates a new SloV00Threshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00ThresholdWithDefaults() *SloV00Threshold {
	this := SloV00Threshold{}
	return &this
}

// GetOperator returns the Operator field value
func (o *SloV00Threshold) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SloV00Threshold) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SloV00Threshold) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value
func (o *SloV00Threshold) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SloV00Threshold) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SloV00Threshold) SetValue(v float64) {
	o.Value = v
}

func (o SloV00Threshold) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00Threshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *SloV00Threshold) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	varSloV00Threshold := _SloV00Threshold{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSloV00Threshold)

	if err != nil {
		return err
	}

	*o = SloV00Threshold(varSloV00Threshold)

	return err
}

type NullableSloV00Threshold struct {
	value *SloV00Threshold
	isSet bool
}

func (v NullableSloV00Threshold) Get() *SloV00Threshold {
	return v.value
}

func (v *NullableSloV00Threshold) Set(val *SloV00Threshold) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00Threshold) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00Threshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00Threshold(val *SloV00Threshold) *NullableSloV00Threshold {
	return &NullableSloV00Threshold{value: val, isSet: true}
}

func (v NullableSloV00Threshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00Threshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


