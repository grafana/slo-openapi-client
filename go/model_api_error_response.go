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

// checks if the ApiErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorResponse{}

// ApiErrorResponse struct for ApiErrorResponse
type ApiErrorResponse struct {
	Code int64 `json:"code"`
	Error string `json:"error"`
}

type _ApiErrorResponse ApiErrorResponse

// NewApiErrorResponse instantiates a new ApiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorResponse(code int64, error_ string) *ApiErrorResponse {
	this := ApiErrorResponse{}
	this.Code = code
	this.Error = error_
	return &this
}

// NewApiErrorResponseWithDefaults instantiates a new ApiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorResponseWithDefaults() *ApiErrorResponse {
	this := ApiErrorResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *ApiErrorResponse) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiErrorResponse) SetCode(v int64) {
	o.Code = v
}

// GetError returns the Error field value
func (o *ApiErrorResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ApiErrorResponse) SetError(v string) {
	o.Error = v
}

func (o ApiErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *ApiErrorResponse) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	varApiErrorResponse := _ApiErrorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varApiErrorResponse)

	if err != nil {
		return err
	}

	*o = ApiErrorResponse(varApiErrorResponse)

	return err
}

type NullableApiErrorResponse struct {
	value *ApiErrorResponse
	isSet bool
}

func (v NullableApiErrorResponse) Get() *ApiErrorResponse {
	return v.value
}

func (v *NullableApiErrorResponse) Set(val *ApiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorResponse(val *ApiErrorResponse) *NullableApiErrorResponse {
	return &NullableApiErrorResponse{value: val, isSet: true}
}

func (v NullableApiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


