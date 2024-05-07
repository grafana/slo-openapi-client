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

// checks if the SloV00Alerting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00Alerting{}

// SloV00Alerting struct for SloV00Alerting
type SloV00Alerting struct {
	Annotations []SloV00Label `json:"annotations,omitempty"`
	FastBurn *SloV00AlertingMetadata `json:"fastBurn,omitempty"`
	Labels interface{} `json:"labels,omitempty"`
	SlowBurn *SloV00AlertingMetadata `json:"slowBurn,omitempty"`
}

// NewSloV00Alerting instantiates a new SloV00Alerting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00Alerting() *SloV00Alerting {
	this := SloV00Alerting{}
	return &this
}

// NewSloV00AlertingWithDefaults instantiates a new SloV00Alerting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00AlertingWithDefaults() *SloV00Alerting {
	this := SloV00Alerting{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *SloV00Alerting) GetAnnotations() []SloV00Label {
	if o == nil || IsNil(o.Annotations) {
		var ret []SloV00Label
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Alerting) GetAnnotationsOk() ([]SloV00Label, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *SloV00Alerting) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []SloV00Label and assigns it to the Annotations field.
func (o *SloV00Alerting) SetAnnotations(v []SloV00Label) {
	o.Annotations = v
}

// GetFastBurn returns the FastBurn field value if set, zero value otherwise.
func (o *SloV00Alerting) GetFastBurn() SloV00AlertingMetadata {
	if o == nil || IsNil(o.FastBurn) {
		var ret SloV00AlertingMetadata
		return ret
	}
	return *o.FastBurn
}

// GetFastBurnOk returns a tuple with the FastBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Alerting) GetFastBurnOk() (*SloV00AlertingMetadata, bool) {
	if o == nil || IsNil(o.FastBurn) {
		return nil, false
	}
	return o.FastBurn, true
}

// HasFastBurn returns a boolean if a field has been set.
func (o *SloV00Alerting) HasFastBurn() bool {
	if o != nil && !IsNil(o.FastBurn) {
		return true
	}

	return false
}

// SetFastBurn gets a reference to the given SloV00AlertingMetadata and assigns it to the FastBurn field.
func (o *SloV00Alerting) SetFastBurn(v SloV00AlertingMetadata) {
	o.FastBurn = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SloV00Alerting) GetLabels() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SloV00Alerting) GetLabelsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SloV00Alerting) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given interface{} and assigns it to the Labels field.
func (o *SloV00Alerting) SetLabels(v interface{}) {
	o.Labels = v
}

// GetSlowBurn returns the SlowBurn field value if set, zero value otherwise.
func (o *SloV00Alerting) GetSlowBurn() SloV00AlertingMetadata {
	if o == nil || IsNil(o.SlowBurn) {
		var ret SloV00AlertingMetadata
		return ret
	}
	return *o.SlowBurn
}

// GetSlowBurnOk returns a tuple with the SlowBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Alerting) GetSlowBurnOk() (*SloV00AlertingMetadata, bool) {
	if o == nil || IsNil(o.SlowBurn) {
		return nil, false
	}
	return o.SlowBurn, true
}

// HasSlowBurn returns a boolean if a field has been set.
func (o *SloV00Alerting) HasSlowBurn() bool {
	if o != nil && !IsNil(o.SlowBurn) {
		return true
	}

	return false
}

// SetSlowBurn gets a reference to the given SloV00AlertingMetadata and assigns it to the SlowBurn field.
func (o *SloV00Alerting) SetSlowBurn(v SloV00AlertingMetadata) {
	o.SlowBurn = &v
}

func (o SloV00Alerting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00Alerting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.FastBurn) {
		toSerialize["fastBurn"] = o.FastBurn
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.SlowBurn) {
		toSerialize["slowBurn"] = o.SlowBurn
	}
	return toSerialize, nil
}

type NullableSloV00Alerting struct {
	value *SloV00Alerting
	isSet bool
}

func (v NullableSloV00Alerting) Get() *SloV00Alerting {
	return v.value
}

func (v *NullableSloV00Alerting) Set(val *SloV00Alerting) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00Alerting) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00Alerting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00Alerting(val *SloV00Alerting) *NullableSloV00Alerting {
	return &NullableSloV00Alerting{value: val, isSet: true}
}

func (v NullableSloV00Alerting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00Alerting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


