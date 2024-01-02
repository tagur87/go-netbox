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

// checks if the PowerFeedSupply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerFeedSupply{}

// PowerFeedSupply struct for PowerFeedSupply
type PowerFeedSupply struct {
	Value                *PatchedWritablePowerFeedRequestSupply `json:"value,omitempty"`
	Label                *PowerFeedSupplyLabel                  `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerFeedSupply PowerFeedSupply

// NewPowerFeedSupply instantiates a new PowerFeedSupply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFeedSupply() *PowerFeedSupply {
	this := PowerFeedSupply{}
	return &this
}

// NewPowerFeedSupplyWithDefaults instantiates a new PowerFeedSupply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFeedSupplyWithDefaults() *PowerFeedSupply {
	this := PowerFeedSupply{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PowerFeedSupply) GetValue() PatchedWritablePowerFeedRequestSupply {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritablePowerFeedRequestSupply
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedSupply) GetValueOk() (*PatchedWritablePowerFeedRequestSupply, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PowerFeedSupply) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritablePowerFeedRequestSupply and assigns it to the Value field.
func (o *PowerFeedSupply) SetValue(v PatchedWritablePowerFeedRequestSupply) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerFeedSupply) GetLabel() PowerFeedSupplyLabel {
	if o == nil || IsNil(o.Label) {
		var ret PowerFeedSupplyLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedSupply) GetLabelOk() (*PowerFeedSupplyLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerFeedSupply) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given PowerFeedSupplyLabel and assigns it to the Label field.
func (o *PowerFeedSupply) SetLabel(v PowerFeedSupplyLabel) {
	o.Label = &v
}

func (o PowerFeedSupply) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerFeedSupply) ToMap() (map[string]interface{}, error) {
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

func (o *PowerFeedSupply) UnmarshalJSON(data []byte) (err error) {
	varPowerFeedSupply := _PowerFeedSupply{}

	err = json.Unmarshal(data, &varPowerFeedSupply)

	if err != nil {
		return err
	}

	*o = PowerFeedSupply(varPowerFeedSupply)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerFeedSupply struct {
	value *PowerFeedSupply
	isSet bool
}

func (v NullablePowerFeedSupply) Get() *PowerFeedSupply {
	return v.value
}

func (v *NullablePowerFeedSupply) Set(val *PowerFeedSupply) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedSupply) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedSupply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedSupply(val *PowerFeedSupply) *NullablePowerFeedSupply {
	return &NullablePowerFeedSupply{value: val, isSet: true}
}

func (v NullablePowerFeedSupply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedSupply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
