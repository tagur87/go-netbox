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

// checks if the NestedCableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedCableRequest{}

// NestedCableRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedCableRequest struct {
	Label                *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedCableRequest NestedCableRequest

// NewNestedCableRequest instantiates a new NestedCableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCableRequest() *NestedCableRequest {
	this := NestedCableRequest{}
	return &this
}

// NewNestedCableRequestWithDefaults instantiates a new NestedCableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedCableRequestWithDefaults() *NestedCableRequest {
	this := NestedCableRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *NestedCableRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCableRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *NestedCableRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *NestedCableRequest) SetLabel(v string) {
	o.Label = &v
}

func (o NestedCableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedCableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedCableRequest) UnmarshalJSON(data []byte) (err error) {
	varNestedCableRequest := _NestedCableRequest{}

	err = json.Unmarshal(data, &varNestedCableRequest)

	if err != nil {
		return err
	}

	*o = NestedCableRequest(varNestedCableRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedCableRequest struct {
	value *NestedCableRequest
	isSet bool
}

func (v NullableNestedCableRequest) Get() *NestedCableRequest {
	return v.value
}

func (v *NullableNestedCableRequest) Set(val *NestedCableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCableRequest(val *NestedCableRequest) *NullableNestedCableRequest {
	return &NullableNestedCableRequest{value: val, isSet: true}
}

func (v NullableNestedCableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
