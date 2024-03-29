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

// checks if the SLOCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SLOCreateResponse{}

// SLOCreateResponse struct for SLOCreateResponse
type SLOCreateResponse struct {
	Message string `json:"message"`
	Uuid string `json:"uuid"`
}

type _SLOCreateResponse SLOCreateResponse

// NewSLOCreateResponse instantiates a new SLOCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOCreateResponse(message string, uuid string) *SLOCreateResponse {
	this := SLOCreateResponse{}
	this.Message = message
	this.Uuid = uuid
	return &this
}

// NewSLOCreateResponseWithDefaults instantiates a new SLOCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOCreateResponseWithDefaults() *SLOCreateResponse {
	this := SLOCreateResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *SLOCreateResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SLOCreateResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SLOCreateResponse) SetMessage(v string) {
	o.Message = v
}

// GetUuid returns the Uuid field value
func (o *SLOCreateResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SLOCreateResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SLOCreateResponse) SetUuid(v string) {
	o.Uuid = v
}

func (o SLOCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SLOCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *SLOCreateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"uuid",
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

	varSLOCreateResponse := _SLOCreateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSLOCreateResponse)

	if err != nil {
		return err
	}

	*o = SLOCreateResponse(varSLOCreateResponse)

	return err
}

type NullableSLOCreateResponse struct {
	value *SLOCreateResponse
	isSet bool
}

func (v NullableSLOCreateResponse) Get() *SLOCreateResponse {
	return v.value
}

func (v *NullableSLOCreateResponse) Set(val *SLOCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOCreateResponse(val *SLOCreateResponse) *NullableSLOCreateResponse {
	return &NullableSLOCreateResponse{value: val, isSet: true}
}

func (v NullableSLOCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


