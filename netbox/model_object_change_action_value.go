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

// ObjectChangeActionValue * `create` - Created * `update` - Updated * `delete` - Deleted
type ObjectChangeActionValue string

// List of ObjectChange_action_value
const (
	OBJECTCHANGEACTIONVALUE_CREATE ObjectChangeActionValue = "create"
	OBJECTCHANGEACTIONVALUE_UPDATE ObjectChangeActionValue = "update"
	OBJECTCHANGEACTIONVALUE_DELETE ObjectChangeActionValue = "delete"
)

// All allowed values of ObjectChangeActionValue enum
var AllowedObjectChangeActionValueEnumValues = []ObjectChangeActionValue{
	"create",
	"update",
	"delete",
}

func (v *ObjectChangeActionValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObjectChangeActionValue(value)
	for _, existing := range AllowedObjectChangeActionValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObjectChangeActionValue", value)
}

// NewObjectChangeActionValueFromValue returns a pointer to a valid ObjectChangeActionValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObjectChangeActionValueFromValue(v string) (*ObjectChangeActionValue, error) {
	ev := ObjectChangeActionValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObjectChangeActionValue: valid values are %v", v, AllowedObjectChangeActionValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObjectChangeActionValue) IsValid() bool {
	for _, existing := range AllowedObjectChangeActionValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjectChange_action_value value
func (v ObjectChangeActionValue) Ptr() *ObjectChangeActionValue {
	return &v
}

type NullableObjectChangeActionValue struct {
	value *ObjectChangeActionValue
	isSet bool
}

func (v NullableObjectChangeActionValue) Get() *ObjectChangeActionValue {
	return v.value
}

func (v *NullableObjectChangeActionValue) Set(val *ObjectChangeActionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectChangeActionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectChangeActionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectChangeActionValue(val *ObjectChangeActionValue) *NullableObjectChangeActionValue {
	return &NullableObjectChangeActionValue{value: val, isSet: true}
}

func (v NullableObjectChangeActionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectChangeActionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
