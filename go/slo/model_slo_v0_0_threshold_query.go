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

// checks if the SloV00ThresholdQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00ThresholdQuery{}

// SloV00ThresholdQuery struct for SloV00ThresholdQuery
type SloV00ThresholdQuery struct {
	GroupByLabels       []string        `json:"groupByLabels,omitempty"`
	Threshold           SloV00Threshold `json:"threshold"`
	ThresholdExpression string          `json:"thresholdExpression"`
}

type _SloV00ThresholdQuery SloV00ThresholdQuery

// NewSloV00ThresholdQuery instantiates a new SloV00ThresholdQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00ThresholdQuery(threshold SloV00Threshold, thresholdExpression string) *SloV00ThresholdQuery {
	this := SloV00ThresholdQuery{}
	this.Threshold = threshold
	this.ThresholdExpression = thresholdExpression
	return &this
}

// NewSloV00ThresholdQueryWithDefaults instantiates a new SloV00ThresholdQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00ThresholdQueryWithDefaults() *SloV00ThresholdQuery {
	this := SloV00ThresholdQuery{}
	return &this
}

// GetGroupByLabels returns the GroupByLabels field value if set, zero value otherwise.
func (o *SloV00ThresholdQuery) GetGroupByLabels() []string {
	if o == nil || IsNil(o.GroupByLabels) {
		var ret []string
		return ret
	}
	return o.GroupByLabels
}

// GetGroupByLabelsOk returns a tuple with the GroupByLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00ThresholdQuery) GetGroupByLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByLabels) {
		return nil, false
	}
	return o.GroupByLabels, true
}

// HasGroupByLabels returns a boolean if a field has been set.
func (o *SloV00ThresholdQuery) HasGroupByLabels() bool {
	if o != nil && !IsNil(o.GroupByLabels) {
		return true
	}

	return false
}

// SetGroupByLabels gets a reference to the given []string and assigns it to the GroupByLabels field.
func (o *SloV00ThresholdQuery) SetGroupByLabels(v []string) {
	o.GroupByLabels = v
}

// GetThreshold returns the Threshold field value
func (o *SloV00ThresholdQuery) GetThreshold() SloV00Threshold {
	if o == nil {
		var ret SloV00Threshold
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *SloV00ThresholdQuery) GetThresholdOk() (*SloV00Threshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *SloV00ThresholdQuery) SetThreshold(v SloV00Threshold) {
	o.Threshold = v
}

// GetThresholdExpression returns the ThresholdExpression field value
func (o *SloV00ThresholdQuery) GetThresholdExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThresholdExpression
}

// GetThresholdExpressionOk returns a tuple with the ThresholdExpression field value
// and a boolean to check if the value has been set.
func (o *SloV00ThresholdQuery) GetThresholdExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdExpression, true
}

// SetThresholdExpression sets field value
func (o *SloV00ThresholdQuery) SetThresholdExpression(v string) {
	o.ThresholdExpression = v
}

func (o SloV00ThresholdQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00ThresholdQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupByLabels) {
		toSerialize["groupByLabels"] = o.GroupByLabels
	}
	toSerialize["threshold"] = o.Threshold
	toSerialize["thresholdExpression"] = o.ThresholdExpression
	return toSerialize, nil
}

func (o *SloV00ThresholdQuery) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varSloV00ThresholdQuery := _SloV00ThresholdQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSloV00ThresholdQuery)

	if err != nil {
		return err
	}

	*o = SloV00ThresholdQuery(varSloV00ThresholdQuery)

	return err
}

type NullableSloV00ThresholdQuery struct {
	value *SloV00ThresholdQuery
	isSet bool
}

func (v NullableSloV00ThresholdQuery) Get() *SloV00ThresholdQuery {
	return v.value
}

func (v *NullableSloV00ThresholdQuery) Set(val *SloV00ThresholdQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00ThresholdQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00ThresholdQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00ThresholdQuery(val *SloV00ThresholdQuery) *NullableSloV00ThresholdQuery {
	return &NullableSloV00ThresholdQuery{value: val, isSet: true}
}

func (v NullableSloV00ThresholdQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00ThresholdQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
