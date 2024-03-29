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

// checks if the DestinationDatasource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationDatasource{}

// DestinationDatasource struct for DestinationDatasource
type DestinationDatasource struct {
	Type *string `json:"type,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewDestinationDatasource instantiates a new DestinationDatasource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationDatasource() *DestinationDatasource {
	this := DestinationDatasource{}
	return &this
}

// NewDestinationDatasourceWithDefaults instantiates a new DestinationDatasource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationDatasourceWithDefaults() *DestinationDatasource {
	this := DestinationDatasource{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DestinationDatasource) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDatasource) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DestinationDatasource) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DestinationDatasource) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *DestinationDatasource) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDatasource) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *DestinationDatasource) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *DestinationDatasource) SetUid(v string) {
	o.Uid = &v
}

func (o DestinationDatasource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationDatasource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableDestinationDatasource struct {
	value *DestinationDatasource
	isSet bool
}

func (v NullableDestinationDatasource) Get() *DestinationDatasource {
	return v.value
}

func (v *NullableDestinationDatasource) Set(val *DestinationDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationDatasource(val *DestinationDatasource) *NullableDestinationDatasource {
	return &NullableDestinationDatasource{value: val, isSet: true}
}

func (v NullableDestinationDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


