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

// checks if the DashboardRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardRef{}

// DashboardRef struct for DashboardRef
type DashboardRef struct {
	UID string `json:"UID"`
}

type _DashboardRef DashboardRef

// NewDashboardRef instantiates a new DashboardRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardRef(uID string) *DashboardRef {
	this := DashboardRef{}
	this.UID = uID
	return &this
}

// NewDashboardRefWithDefaults instantiates a new DashboardRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardRefWithDefaults() *DashboardRef {
	this := DashboardRef{}
	return &this
}

// GetUID returns the UID field value
func (o *DashboardRef) GetUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UID
}

// GetUIDOk returns a tuple with the UID field value
// and a boolean to check if the value has been set.
func (o *DashboardRef) GetUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UID, true
}

// SetUID sets field value
func (o *DashboardRef) SetUID(v string) {
	o.UID = v
}

func (o DashboardRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["UID"] = o.UID
	return toSerialize, nil
}

func (o *DashboardRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"UID",
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

	varDashboardRef := _DashboardRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDashboardRef)

	if err != nil {
		return err
	}

	*o = DashboardRef(varDashboardRef)

	return err
}

type NullableDashboardRef struct {
	value *DashboardRef
	isSet bool
}

func (v NullableDashboardRef) Get() *DashboardRef {
	return v.value
}

func (v *NullableDashboardRef) Set(val *DashboardRef) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardRef) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardRef(val *DashboardRef) *NullableDashboardRef {
	return &NullableDashboardRef{value: val, isSet: true}
}

func (v NullableDashboardRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


