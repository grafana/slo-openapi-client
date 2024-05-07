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

// checks if the ApiSLOListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSLOListResponse{}

// ApiSLOListResponse struct for ApiSLOListResponse
type ApiSLOListResponse struct {
	Slos []SloV00Slo `json:"slos"`
}

type _ApiSLOListResponse ApiSLOListResponse

// NewApiSLOListResponse instantiates a new ApiSLOListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSLOListResponse(slos []SloV00Slo) *ApiSLOListResponse {
	this := ApiSLOListResponse{}
	this.Slos = slos
	return &this
}

// NewApiSLOListResponseWithDefaults instantiates a new ApiSLOListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSLOListResponseWithDefaults() *ApiSLOListResponse {
	this := ApiSLOListResponse{}
	return &this
}

// GetSlos returns the Slos field value
func (o *ApiSLOListResponse) GetSlos() []SloV00Slo {
	if o == nil {
		var ret []SloV00Slo
		return ret
	}

	return o.Slos
}

// GetSlosOk returns a tuple with the Slos field value
// and a boolean to check if the value has been set.
func (o *ApiSLOListResponse) GetSlosOk() ([]SloV00Slo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slos, true
}

// SetSlos sets field value
func (o *ApiSLOListResponse) SetSlos(v []SloV00Slo) {
	o.Slos = v
}

func (o ApiSLOListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSLOListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slos"] = o.Slos
	return toSerialize, nil
}

func (o *ApiSLOListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"slos",
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

	varApiSLOListResponse := _ApiSLOListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiSLOListResponse)

	if err != nil {
		return err
	}

	*o = ApiSLOListResponse(varApiSLOListResponse)

	return err
}

type NullableApiSLOListResponse struct {
	value *ApiSLOListResponse
	isSet bool
}

func (v NullableApiSLOListResponse) Get() *ApiSLOListResponse {
	return v.value
}

func (v *NullableApiSLOListResponse) Set(val *ApiSLOListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSLOListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSLOListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSLOListResponse(val *ApiSLOListResponse) *NullableApiSLOListResponse {
	return &NullableApiSLOListResponse{value: val, isSet: true}
}

func (v NullableApiSLOListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSLOListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

