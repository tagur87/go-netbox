/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.2 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the NestedCustomFieldChoiceSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedCustomFieldChoiceSetRequest{}

// NestedCustomFieldChoiceSetRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedCustomFieldChoiceSetRequest struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedCustomFieldChoiceSetRequest NestedCustomFieldChoiceSetRequest

// NewNestedCustomFieldChoiceSetRequest instantiates a new NestedCustomFieldChoiceSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCustomFieldChoiceSetRequest(name string) *NestedCustomFieldChoiceSetRequest {
	this := NestedCustomFieldChoiceSetRequest{}
	this.Name = name
	return &this
}

// NewNestedCustomFieldChoiceSetRequestWithDefaults instantiates a new NestedCustomFieldChoiceSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedCustomFieldChoiceSetRequestWithDefaults() *NestedCustomFieldChoiceSetRequest {
	this := NestedCustomFieldChoiceSetRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedCustomFieldChoiceSetRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedCustomFieldChoiceSetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedCustomFieldChoiceSetRequest) SetName(v string) {
	o.Name = v
}

func (o NestedCustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedCustomFieldChoiceSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedCustomFieldChoiceSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNestedCustomFieldChoiceSetRequest := _NestedCustomFieldChoiceSetRequest{}

	err = json.Unmarshal(data, &varNestedCustomFieldChoiceSetRequest)

	if err != nil {
		return err
	}

	*o = NestedCustomFieldChoiceSetRequest(varNestedCustomFieldChoiceSetRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedCustomFieldChoiceSetRequest struct {
	value *NestedCustomFieldChoiceSetRequest
	isSet bool
}

func (v NullableNestedCustomFieldChoiceSetRequest) Get() *NestedCustomFieldChoiceSetRequest {
	return v.value
}

func (v *NullableNestedCustomFieldChoiceSetRequest) Set(val *NestedCustomFieldChoiceSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCustomFieldChoiceSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCustomFieldChoiceSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCustomFieldChoiceSetRequest(val *NestedCustomFieldChoiceSetRequest) *NullableNestedCustomFieldChoiceSetRequest {
	return &NullableNestedCustomFieldChoiceSetRequest{value: val, isSet: true}
}

func (v NullableNestedCustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCustomFieldChoiceSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
