/*
Grafana SLO API

This API CRUDs SLO objects for the Grafana plugin.  Modifying an SLO object will create or update recording and alerting rules in a connected Prometheus instance and create or update dashboards in Grafana.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RatioQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatioQuery{}

// RatioQuery struct for RatioQuery
type RatioQuery struct {
	GroupByLabels []string `json:"groupByLabels,omitempty"`
	SuccessMetric MetricDef `json:"successMetric"`
	TotalMetric MetricDef `json:"totalMetric"`
}

type _RatioQuery RatioQuery

// NewRatioQuery instantiates a new RatioQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatioQuery(successMetric MetricDef, totalMetric MetricDef) *RatioQuery {
	this := RatioQuery{}
	this.SuccessMetric = successMetric
	this.TotalMetric = totalMetric
	return &this
}

// NewRatioQueryWithDefaults instantiates a new RatioQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatioQueryWithDefaults() *RatioQuery {
	this := RatioQuery{}
	return &this
}

// GetGroupByLabels returns the GroupByLabels field value if set, zero value otherwise.
func (o *RatioQuery) GetGroupByLabels() []string {
	if o == nil || IsNil(o.GroupByLabels) {
		var ret []string
		return ret
	}
	return o.GroupByLabels
}

// GetGroupByLabelsOk returns a tuple with the GroupByLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatioQuery) GetGroupByLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByLabels) {
		return nil, false
	}
	return o.GroupByLabels, true
}

// HasGroupByLabels returns a boolean if a field has been set.
func (o *RatioQuery) HasGroupByLabels() bool {
	if o != nil && !IsNil(o.GroupByLabels) {
		return true
	}

	return false
}

// SetGroupByLabels gets a reference to the given []string and assigns it to the GroupByLabels field.
func (o *RatioQuery) SetGroupByLabels(v []string) {
	o.GroupByLabels = v
}

// GetSuccessMetric returns the SuccessMetric field value
func (o *RatioQuery) GetSuccessMetric() MetricDef {
	if o == nil {
		var ret MetricDef
		return ret
	}

	return o.SuccessMetric
}

// GetSuccessMetricOk returns a tuple with the SuccessMetric field value
// and a boolean to check if the value has been set.
func (o *RatioQuery) GetSuccessMetricOk() (*MetricDef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessMetric, true
}

// SetSuccessMetric sets field value
func (o *RatioQuery) SetSuccessMetric(v MetricDef) {
	o.SuccessMetric = v
}

// GetTotalMetric returns the TotalMetric field value
func (o *RatioQuery) GetTotalMetric() MetricDef {
	if o == nil {
		var ret MetricDef
		return ret
	}

	return o.TotalMetric
}

// GetTotalMetricOk returns a tuple with the TotalMetric field value
// and a boolean to check if the value has been set.
func (o *RatioQuery) GetTotalMetricOk() (*MetricDef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMetric, true
}

// SetTotalMetric sets field value
func (o *RatioQuery) SetTotalMetric(v MetricDef) {
	o.TotalMetric = v
}

func (o RatioQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatioQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupByLabels) {
		toSerialize["groupByLabels"] = o.GroupByLabels
	}
	toSerialize["successMetric"] = o.SuccessMetric
	toSerialize["totalMetric"] = o.TotalMetric
	return toSerialize, nil
}

func (o *RatioQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"successMetric",
		"totalMetric",
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

	varRatioQuery := _RatioQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRatioQuery)

	if err != nil {
		return err
	}

	*o = RatioQuery(varRatioQuery)

	return err
}

type NullableRatioQuery struct {
	value *RatioQuery
	isSet bool
}

func (v NullableRatioQuery) Get() *RatioQuery {
	return v.value
}

func (v *NullableRatioQuery) Set(val *RatioQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRatioQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRatioQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatioQuery(val *RatioQuery) *NullableRatioQuery {
	return &NullableRatioQuery{value: val, isSet: true}
}

func (v NullableRatioQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatioQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


