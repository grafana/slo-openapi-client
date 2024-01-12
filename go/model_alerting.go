/*
Grafana SLO API

This API CRUDs SLO objects for the Grafana plugin.  Modifying an SLO object will create or update recording and alerting rules in a connected Prometheus instance and create or update dashboards in Grafana.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the Alerting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Alerting{}

// Alerting struct for Alerting
type Alerting struct {
	Annotations []Label `json:"annotations,omitempty"`
	FastBurn *AlertingMetadata `json:"fastBurn,omitempty"`
	Labels interface{} `json:"labels,omitempty"`
	SlowBurn *AlertingMetadata `json:"slowBurn,omitempty"`
}

// NewAlerting instantiates a new Alerting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlerting() *Alerting {
	this := Alerting{}
	return &this
}

// NewAlertingWithDefaults instantiates a new Alerting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingWithDefaults() *Alerting {
	this := Alerting{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *Alerting) GetAnnotations() []Label {
	if o == nil || IsNil(o.Annotations) {
		var ret []Label
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alerting) GetAnnotationsOk() ([]Label, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *Alerting) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []Label and assigns it to the Annotations field.
func (o *Alerting) SetAnnotations(v []Label) {
	o.Annotations = v
}

// GetFastBurn returns the FastBurn field value if set, zero value otherwise.
func (o *Alerting) GetFastBurn() AlertingMetadata {
	if o == nil || IsNil(o.FastBurn) {
		var ret AlertingMetadata
		return ret
	}
	return *o.FastBurn
}

// GetFastBurnOk returns a tuple with the FastBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alerting) GetFastBurnOk() (*AlertingMetadata, bool) {
	if o == nil || IsNil(o.FastBurn) {
		return nil, false
	}
	return o.FastBurn, true
}

// HasFastBurn returns a boolean if a field has been set.
func (o *Alerting) HasFastBurn() bool {
	if o != nil && !IsNil(o.FastBurn) {
		return true
	}

	return false
}

// SetFastBurn gets a reference to the given AlertingMetadata and assigns it to the FastBurn field.
func (o *Alerting) SetFastBurn(v AlertingMetadata) {
	o.FastBurn = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alerting) GetLabels() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alerting) GetLabelsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Alerting) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given interface{} and assigns it to the Labels field.
func (o *Alerting) SetLabels(v interface{}) {
	o.Labels = v
}

// GetSlowBurn returns the SlowBurn field value if set, zero value otherwise.
func (o *Alerting) GetSlowBurn() AlertingMetadata {
	if o == nil || IsNil(o.SlowBurn) {
		var ret AlertingMetadata
		return ret
	}
	return *o.SlowBurn
}

// GetSlowBurnOk returns a tuple with the SlowBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alerting) GetSlowBurnOk() (*AlertingMetadata, bool) {
	if o == nil || IsNil(o.SlowBurn) {
		return nil, false
	}
	return o.SlowBurn, true
}

// HasSlowBurn returns a boolean if a field has been set.
func (o *Alerting) HasSlowBurn() bool {
	if o != nil && !IsNil(o.SlowBurn) {
		return true
	}

	return false
}

// SetSlowBurn gets a reference to the given AlertingMetadata and assigns it to the SlowBurn field.
func (o *Alerting) SetSlowBurn(v AlertingMetadata) {
	o.SlowBurn = &v
}

func (o Alerting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Alerting) ToMap() (map[string]interface{}, error) {
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

type NullableAlerting struct {
	value *Alerting
	isSet bool
}

func (v NullableAlerting) Get() *Alerting {
	return v.value
}

func (v *NullableAlerting) Set(val *Alerting) {
	v.value = val
	v.isSet = true
}

func (v NullableAlerting) IsSet() bool {
	return v.isSet
}

func (v *NullableAlerting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlerting(val *Alerting) *NullableAlerting {
	return &NullableAlerting{value: val, isSet: true}
}

func (v NullableAlerting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlerting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


