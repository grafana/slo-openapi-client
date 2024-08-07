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

// checks if the SloV00DashboardRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00DashboardRef{}

// SloV00DashboardRef struct for SloV00DashboardRef
type SloV00DashboardRef struct {
	UID string `json:"UID"`
}

type _SloV00DashboardRef SloV00DashboardRef

// NewSloV00DashboardRef instantiates a new SloV00DashboardRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00DashboardRef(uID string) *SloV00DashboardRef {
	this := SloV00DashboardRef{}
	this.UID = uID
	return &this
}

// NewSloV00DashboardRefWithDefaults instantiates a new SloV00DashboardRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00DashboardRefWithDefaults() *SloV00DashboardRef {
	this := SloV00DashboardRef{}
	return &this
}

// GetUID returns the UID field value
func (o *SloV00DashboardRef) GetUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UID
}

// GetUIDOk returns a tuple with the UID field value
// and a boolean to check if the value has been set.
func (o *SloV00DashboardRef) GetUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UID, true
}

// SetUID sets field value
func (o *SloV00DashboardRef) SetUID(v string) {
	o.UID = v
}

func (o SloV00DashboardRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00DashboardRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["UID"] = o.UID
	return toSerialize, nil
}

func (o *SloV00DashboardRef) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varSloV00DashboardRef := _SloV00DashboardRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSloV00DashboardRef)

	if err != nil {
		return err
	}

	*o = SloV00DashboardRef(varSloV00DashboardRef)

	return err
}

type NullableSloV00DashboardRef struct {
	value *SloV00DashboardRef
	isSet bool
}

func (v NullableSloV00DashboardRef) Get() *SloV00DashboardRef {
	return v.value
}

func (v *NullableSloV00DashboardRef) Set(val *SloV00DashboardRef) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00DashboardRef) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00DashboardRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00DashboardRef(val *SloV00DashboardRef) *NullableSloV00DashboardRef {
	return &NullableSloV00DashboardRef{value: val, isSet: true}
}

func (v NullableSloV00DashboardRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00DashboardRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
