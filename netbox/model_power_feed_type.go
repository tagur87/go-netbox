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

// checks if the PowerFeedType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerFeedType{}

// PowerFeedType struct for PowerFeedType
type PowerFeedType struct {
	Value                *PatchedWritablePowerFeedRequestType `json:"value,omitempty"`
	Label                *PowerFeedTypeLabel                  `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerFeedType PowerFeedType

// NewPowerFeedType instantiates a new PowerFeedType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFeedType() *PowerFeedType {
	this := PowerFeedType{}
	return &this
}

// NewPowerFeedTypeWithDefaults instantiates a new PowerFeedType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFeedTypeWithDefaults() *PowerFeedType {
	this := PowerFeedType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PowerFeedType) GetValue() PatchedWritablePowerFeedRequestType {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritablePowerFeedRequestType
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedType) GetValueOk() (*PatchedWritablePowerFeedRequestType, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PowerFeedType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritablePowerFeedRequestType and assigns it to the Value field.
func (o *PowerFeedType) SetValue(v PatchedWritablePowerFeedRequestType) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerFeedType) GetLabel() PowerFeedTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret PowerFeedTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedType) GetLabelOk() (*PowerFeedTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerFeedType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given PowerFeedTypeLabel and assigns it to the Label field.
func (o *PowerFeedType) SetLabel(v PowerFeedTypeLabel) {
	o.Label = &v
}

func (o PowerFeedType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerFeedType) ToMap() (map[string]interface{}, error) {
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

func (o *PowerFeedType) UnmarshalJSON(data []byte) (err error) {
	varPowerFeedType := _PowerFeedType{}

	err = json.Unmarshal(data, &varPowerFeedType)

	if err != nil {
		return err
	}

	*o = PowerFeedType(varPowerFeedType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerFeedType struct {
	value *PowerFeedType
	isSet bool
}

func (v NullablePowerFeedType) Get() *PowerFeedType {
	return v.value
}

func (v *NullablePowerFeedType) Set(val *PowerFeedType) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedType) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedType(val *PowerFeedType) *NullablePowerFeedType {
	return &NullablePowerFeedType{value: val, isSet: true}
}

func (v NullablePowerFeedType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
