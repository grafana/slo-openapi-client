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

// checks if the SloV00FailureRatioQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00FailureRatioQuery{}

// SloV00FailureRatioQuery struct for SloV00FailureRatioQuery
type SloV00FailureRatioQuery struct {
	FailureMetric SloV00MetricDef `json:"failureMetric"`
	GroupByLabels []string        `json:"groupByLabels,omitempty"`
	TotalMetric   SloV00MetricDef `json:"totalMetric"`
}

type _SloV00FailureRatioQuery SloV00FailureRatioQuery

// NewSloV00FailureRatioQuery instantiates a new SloV00FailureRatioQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00FailureRatioQuery(failureMetric SloV00MetricDef, totalMetric SloV00MetricDef) *SloV00FailureRatioQuery {
	this := SloV00FailureRatioQuery{}
	this.FailureMetric = failureMetric
	this.TotalMetric = totalMetric
	return &this
}

// NewSloV00FailureRatioQueryWithDefaults instantiates a new SloV00FailureRatioQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00FailureRatioQueryWithDefaults() *SloV00FailureRatioQuery {
	this := SloV00FailureRatioQuery{}
	return &this
}

// GetFailureMetric returns the FailureMetric field value
func (o *SloV00FailureRatioQuery) GetFailureMetric() SloV00MetricDef {
	if o == nil {
		var ret SloV00MetricDef
		return ret
	}

	return o.FailureMetric
}

// GetFailureMetricOk returns a tuple with the FailureMetric field value
// and a boolean to check if the value has been set.
func (o *SloV00FailureRatioQuery) GetFailureMetricOk() (*SloV00MetricDef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureMetric, true
}

// SetFailureMetric sets field value
func (o *SloV00FailureRatioQuery) SetFailureMetric(v SloV00MetricDef) {
	o.FailureMetric = v
}

// GetGroupByLabels returns the GroupByLabels field value if set, zero value otherwise.
func (o *SloV00FailureRatioQuery) GetGroupByLabels() []string {
	if o == nil || IsNil(o.GroupByLabels) {
		var ret []string
		return ret
	}
	return o.GroupByLabels
}

// GetGroupByLabelsOk returns a tuple with the GroupByLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00FailureRatioQuery) GetGroupByLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByLabels) {
		return nil, false
	}
	return o.GroupByLabels, true
}

// HasGroupByLabels returns a boolean if a field has been set.
func (o *SloV00FailureRatioQuery) HasGroupByLabels() bool {
	if o != nil && !IsNil(o.GroupByLabels) {
		return true
	}

	return false
}

// SetGroupByLabels gets a reference to the given []string and assigns it to the GroupByLabels field.
func (o *SloV00FailureRatioQuery) SetGroupByLabels(v []string) {
	o.GroupByLabels = v
}

// GetTotalMetric returns the TotalMetric field value
func (o *SloV00FailureRatioQuery) GetTotalMetric() SloV00MetricDef {
	if o == nil {
		var ret SloV00MetricDef
		return ret
	}

	return o.TotalMetric
}

// GetTotalMetricOk returns a tuple with the TotalMetric field value
// and a boolean to check if the value has been set.
func (o *SloV00FailureRatioQuery) GetTotalMetricOk() (*SloV00MetricDef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMetric, true
}

// SetTotalMetric sets field value
func (o *SloV00FailureRatioQuery) SetTotalMetric(v SloV00MetricDef) {
	o.TotalMetric = v
}

func (o SloV00FailureRatioQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00FailureRatioQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["failureMetric"] = o.FailureMetric
	if !IsNil(o.GroupByLabels) {
		toSerialize["groupByLabels"] = o.GroupByLabels
	}
	toSerialize["totalMetric"] = o.TotalMetric
	return toSerialize, nil
}

func (o *SloV00FailureRatioQuery) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varSloV00FailureRatioQuery := _SloV00FailureRatioQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSloV00FailureRatioQuery)

	if err != nil {
		return err
	}

	*o = SloV00FailureRatioQuery(varSloV00FailureRatioQuery)

	return err
}

type NullableSloV00FailureRatioQuery struct {
	value *SloV00FailureRatioQuery
	isSet bool
}

func (v NullableSloV00FailureRatioQuery) Get() *SloV00FailureRatioQuery {
	return v.value
}

func (v *NullableSloV00FailureRatioQuery) Set(val *SloV00FailureRatioQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00FailureRatioQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00FailureRatioQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00FailureRatioQuery(val *SloV00FailureRatioQuery) *NullableSloV00FailureRatioQuery {
	return &NullableSloV00FailureRatioQuery{value: val, isSet: true}
}

func (v NullableSloV00FailureRatioQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00FailureRatioQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
