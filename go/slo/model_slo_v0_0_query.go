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

// checks if the SloV00Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloV00Query{}

// SloV00Query struct for SloV00Query
type SloV00Query struct {
	FailureRatio     *SloV00FailureRatioQuery     `json:"failureRatio,omitempty"`
	FailureThreshold *SloV00FailureThresholdQuery `json:"failureThreshold,omitempty"`
	Freeform         *SloV00FreeformQuery         `json:"freeform,omitempty"`
	GrafanaQueries   *SloV00GrafanaQueries        `json:"grafanaQueries,omitempty"`
	Ratio            *SloV00RatioQuery            `json:"ratio,omitempty"`
	Threshold        *SloV00ThresholdQuery        `json:"threshold,omitempty"`
	Type             string                       `json:"type"`
}

type _SloV00Query SloV00Query

// NewSloV00Query instantiates a new SloV00Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloV00Query(type_ string) *SloV00Query {
	this := SloV00Query{}
	this.Type = type_
	return &this
}

// NewSloV00QueryWithDefaults instantiates a new SloV00Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloV00QueryWithDefaults() *SloV00Query {
	this := SloV00Query{}
	return &this
}

// GetFailureRatio returns the FailureRatio field value if set, zero value otherwise.
func (o *SloV00Query) GetFailureRatio() SloV00FailureRatioQuery {
	if o == nil || IsNil(o.FailureRatio) {
		var ret SloV00FailureRatioQuery
		return ret
	}
	return *o.FailureRatio
}

// GetFailureRatioOk returns a tuple with the FailureRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Query) GetFailureRatioOk() (*SloV00FailureRatioQuery, bool) {
	if o == nil || IsNil(o.FailureRatio) {
		return nil, false
	}
	return o.FailureRatio, true
}

// HasFailureRatio returns a boolean if a field has been set.
func (o *SloV00Query) HasFailureRatio() bool {
	if o != nil && !IsNil(o.FailureRatio) {
		return true
	}

	return false
}

// SetFailureRatio gets a reference to the given SloV00FailureRatioQuery and assigns it to the FailureRatio field.
func (o *SloV00Query) SetFailureRatio(v SloV00FailureRatioQuery) {
	o.FailureRatio = &v
}

// GetFailureThreshold returns the FailureThreshold field value if set, zero value otherwise.
func (o *SloV00Query) GetFailureThreshold() SloV00FailureThresholdQuery {
	if o == nil || IsNil(o.FailureThreshold) {
		var ret SloV00FailureThresholdQuery
		return ret
	}
	return *o.FailureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Query) GetFailureThresholdOk() (*SloV00FailureThresholdQuery, bool) {
	if o == nil || IsNil(o.FailureThreshold) {
		return nil, false
	}
	return o.FailureThreshold, true
}

// HasFailureThreshold returns a boolean if a field has been set.
func (o *SloV00Query) HasFailureThreshold() bool {
	if o != nil && !IsNil(o.FailureThreshold) {
		return true
	}

	return false
}

// SetFailureThreshold gets a reference to the given SloV00FailureThresholdQuery and assigns it to the FailureThreshold field.
func (o *SloV00Query) SetFailureThreshold(v SloV00FailureThresholdQuery) {
	o.FailureThreshold = &v
}

// GetFreeform returns the Freeform field value if set, zero value otherwise.
func (o *SloV00Query) GetFreeform() SloV00FreeformQuery {
	if o == nil || IsNil(o.Freeform) {
		var ret SloV00FreeformQuery
		return ret
	}
	return *o.Freeform
}

// GetFreeformOk returns a tuple with the Freeform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Query) GetFreeformOk() (*SloV00FreeformQuery, bool) {
	if o == nil || IsNil(o.Freeform) {
		return nil, false
	}
	return o.Freeform, true
}

// HasFreeform returns a boolean if a field has been set.
func (o *SloV00Query) HasFreeform() bool {
	if o != nil && !IsNil(o.Freeform) {
		return true
	}

	return false
}

// SetFreeform gets a reference to the given SloV00FreeformQuery and assigns it to the Freeform field.
func (o *SloV00Query) SetFreeform(v SloV00FreeformQuery) {
	o.Freeform = &v
}

// GetGrafanaQueries returns the GrafanaQueries field value if set, zero value otherwise.
func (o *SloV00Query) GetGrafanaQueries() SloV00GrafanaQueries {
	if o == nil || IsNil(o.GrafanaQueries) {
		var ret SloV00GrafanaQueries
		return ret
	}
	return *o.GrafanaQueries
}

// GetGrafanaQueriesOk returns a tuple with the GrafanaQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Query) GetGrafanaQueriesOk() (*SloV00GrafanaQueries, bool) {
	if o == nil || IsNil(o.GrafanaQueries) {
		return nil, false
	}
	return o.GrafanaQueries, true
}

// HasGrafanaQueries returns a boolean if a field has been set.
func (o *SloV00Query) HasGrafanaQueries() bool {
	if o != nil && !IsNil(o.GrafanaQueries) {
		return true
	}

	return false
}

// SetGrafanaQueries gets a reference to the given SloV00GrafanaQueries and assigns it to the GrafanaQueries field.
func (o *SloV00Query) SetGrafanaQueries(v SloV00GrafanaQueries) {
	o.GrafanaQueries = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *SloV00Query) GetRatio() SloV00RatioQuery {
	if o == nil || IsNil(o.Ratio) {
		var ret SloV00RatioQuery
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Query) GetRatioOk() (*SloV00RatioQuery, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *SloV00Query) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given SloV00RatioQuery and assigns it to the Ratio field.
func (o *SloV00Query) SetRatio(v SloV00RatioQuery) {
	o.Ratio = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *SloV00Query) GetThreshold() SloV00ThresholdQuery {
	if o == nil || IsNil(o.Threshold) {
		var ret SloV00ThresholdQuery
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloV00Query) GetThresholdOk() (*SloV00ThresholdQuery, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *SloV00Query) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given SloV00ThresholdQuery and assigns it to the Threshold field.
func (o *SloV00Query) SetThreshold(v SloV00ThresholdQuery) {
	o.Threshold = &v
}

// GetType returns the Type field value
func (o *SloV00Query) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SloV00Query) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SloV00Query) SetType(v string) {
	o.Type = v
}

func (o SloV00Query) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloV00Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailureRatio) {
		toSerialize["failureRatio"] = o.FailureRatio
	}
	if !IsNil(o.FailureThreshold) {
		toSerialize["failureThreshold"] = o.FailureThreshold
	}
	if !IsNil(o.Freeform) {
		toSerialize["freeform"] = o.Freeform
	}
	if !IsNil(o.GrafanaQueries) {
		toSerialize["grafanaQueries"] = o.GrafanaQueries
	}
	if !IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *SloV00Query) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varSloV00Query := _SloV00Query{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSloV00Query)

	if err != nil {
		return err
	}

	*o = SloV00Query(varSloV00Query)

	return err
}

type NullableSloV00Query struct {
	value *SloV00Query
	isSet bool
}

func (v NullableSloV00Query) Get() *SloV00Query {
	return v.value
}

func (v *NullableSloV00Query) Set(val *SloV00Query) {
	v.value = val
	v.isSet = true
}

func (v NullableSloV00Query) IsSet() bool {
	return v.isSet
}

func (v *NullableSloV00Query) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloV00Query(val *SloV00Query) *NullableSloV00Query {
	return &NullableSloV00Query{value: val, isSet: true}
}

func (v NullableSloV00Query) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloV00Query) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
