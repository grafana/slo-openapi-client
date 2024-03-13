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

// checks if the Folder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Folder{}

// Folder struct for Folder
type Folder struct {
	Uid *string `json:"uid,omitempty"`
}

// NewFolder instantiates a new Folder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolder() *Folder {
	this := Folder{}
	return &this
}

// NewFolderWithDefaults instantiates a new Folder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderWithDefaults() *Folder {
	this := Folder{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *Folder) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *Folder) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *Folder) SetUid(v string) {
	o.Uid = &v
}

func (o Folder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Folder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableFolder struct {
	value *Folder
	isSet bool
}

func (v NullableFolder) Get() *Folder {
	return v.value
}

func (v *NullableFolder) Set(val *Folder) {
	v.value = val
	v.isSet = true
}

func (v NullableFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolder(val *Folder) *NullableFolder {
	return &NullableFolder{value: val, isSet: true}
}

func (v NullableFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


