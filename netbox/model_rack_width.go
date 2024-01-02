/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.2 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the RackWidth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackWidth{}

// RackWidth struct for RackWidth
type RackWidth struct {
	Value                *RackWidthValue `json:"value,omitempty"`
	Label                *RackWidthLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackWidth RackWidth

// NewRackWidth instantiates a new RackWidth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackWidth() *RackWidth {
	this := RackWidth{}
	return &this
}

// NewRackWidthWithDefaults instantiates a new RackWidth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackWidthWithDefaults() *RackWidth {
	this := RackWidth{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RackWidth) GetValue() RackWidthValue {
	if o == nil || IsNil(o.Value) {
		var ret RackWidthValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackWidth) GetValueOk() (*RackWidthValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RackWidth) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given RackWidthValue and assigns it to the Value field.
func (o *RackWidth) SetValue(v RackWidthValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RackWidth) GetLabel() RackWidthLabel {
	if o == nil || IsNil(o.Label) {
		var ret RackWidthLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackWidth) GetLabelOk() (*RackWidthLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RackWidth) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given RackWidthLabel and assigns it to the Label field.
func (o *RackWidth) SetLabel(v RackWidthLabel) {
	o.Label = &v
}

func (o RackWidth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackWidth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RackWidth) UnmarshalJSON(data []byte) (err error) {
	varRackWidth := _RackWidth{}

	err = json.Unmarshal(data, &varRackWidth)

	if err != nil {
		return err
	}

	*o = RackWidth(varRackWidth)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackWidth struct {
	value *RackWidth
	isSet bool
}

func (v NullableRackWidth) Get() *RackWidth {
	return v.value
}

func (v *NullableRackWidth) Set(val *RackWidth) {
	v.value = val
	v.isSet = true
}

func (v NullableRackWidth) IsSet() bool {
	return v.isSet
}

func (v *NullableRackWidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackWidth(val *RackWidth) *NullableRackWidth {
	return &NullableRackWidth{value: val, isSet: true}
}

func (v NullableRackWidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackWidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
