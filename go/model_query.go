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

// checks if the Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Query{}

// Query struct for Query
type Query struct {
	Freeform *FreeformQuery `json:"freeform,omitempty"`
	Histogram *HistogramQuery `json:"histogram,omitempty"`
	Ratio *RatioQuery `json:"ratio,omitempty"`
	Threshold *ThresholdQuery `json:"threshold,omitempty"`
	Type string `json:"type"`
}

type _Query Query

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery(type_ string) *Query {
	this := Query{}
	this.Type = type_
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetFreeform returns the Freeform field value if set, zero value otherwise.
func (o *Query) GetFreeform() FreeformQuery {
	if o == nil || IsNil(o.Freeform) {
		var ret FreeformQuery
		return ret
	}
	return *o.Freeform
}

// GetFreeformOk returns a tuple with the Freeform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetFreeformOk() (*FreeformQuery, bool) {
	if o == nil || IsNil(o.Freeform) {
		return nil, false
	}
	return o.Freeform, true
}

// HasFreeform returns a boolean if a field has been set.
func (o *Query) HasFreeform() bool {
	if o != nil && !IsNil(o.Freeform) {
		return true
	}

	return false
}

// SetFreeform gets a reference to the given FreeformQuery and assigns it to the Freeform field.
func (o *Query) SetFreeform(v FreeformQuery) {
	o.Freeform = &v
}

// GetHistogram returns the Histogram field value if set, zero value otherwise.
func (o *Query) GetHistogram() HistogramQuery {
	if o == nil || IsNil(o.Histogram) {
		var ret HistogramQuery
		return ret
	}
	return *o.Histogram
}

// GetHistogramOk returns a tuple with the Histogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetHistogramOk() (*HistogramQuery, bool) {
	if o == nil || IsNil(o.Histogram) {
		return nil, false
	}
	return o.Histogram, true
}

// HasHistogram returns a boolean if a field has been set.
func (o *Query) HasHistogram() bool {
	if o != nil && !IsNil(o.Histogram) {
		return true
	}

	return false
}

// SetHistogram gets a reference to the given HistogramQuery and assigns it to the Histogram field.
func (o *Query) SetHistogram(v HistogramQuery) {
	o.Histogram = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *Query) GetRatio() RatioQuery {
	if o == nil || IsNil(o.Ratio) {
		var ret RatioQuery
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetRatioOk() (*RatioQuery, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *Query) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given RatioQuery and assigns it to the Ratio field.
func (o *Query) SetRatio(v RatioQuery) {
	o.Ratio = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Query) GetThreshold() ThresholdQuery {
	if o == nil || IsNil(o.Threshold) {
		var ret ThresholdQuery
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetThresholdOk() (*ThresholdQuery, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Query) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given ThresholdQuery and assigns it to the Threshold field.
func (o *Query) SetThreshold(v ThresholdQuery) {
	o.Threshold = &v
}

// GetType returns the Type field value
func (o *Query) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Query) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Query) SetType(v string) {
	o.Type = v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Freeform) {
		toSerialize["freeform"] = o.Freeform
	}
	if !IsNil(o.Histogram) {
		toSerialize["histogram"] = o.Histogram
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

func (o *Query) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varQuery := _Query{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuery)

	if err != nil {
		return err
	}

	*o = Query(varQuery)

	return err
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


