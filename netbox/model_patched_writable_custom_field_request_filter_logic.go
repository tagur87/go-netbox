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

// PatchedWritableCustomFieldRequestFilterLogic Loose matches any instance of a given string; exact matches the entire field.  * `disabled` - Disabled * `loose` - Loose * `exact` - Exact
type PatchedWritableCustomFieldRequestFilterLogic string

// List of PatchedWritableCustomFieldRequest_filter_logic
const (
	PATCHEDWRITABLECUSTOMFIELDREQUESTFILTERLOGIC_DISABLED PatchedWritableCustomFieldRequestFilterLogic = "disabled"
	PATCHEDWRITABLECUSTOMFIELDREQUESTFILTERLOGIC_LOOSE    PatchedWritableCustomFieldRequestFilterLogic = "loose"
	PATCHEDWRITABLECUSTOMFIELDREQUESTFILTERLOGIC_EXACT    PatchedWritableCustomFieldRequestFilterLogic = "exact"
)

// All allowed values of PatchedWritableCustomFieldRequestFilterLogic enum
var AllowedPatchedWritableCustomFieldRequestFilterLogicEnumValues = []PatchedWritableCustomFieldRequestFilterLogic{
	"disabled",
	"loose",
	"exact",
}

func (v *PatchedWritableCustomFieldRequestFilterLogic) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableCustomFieldRequestFilterLogic(value)
	for _, existing := range AllowedPatchedWritableCustomFieldRequestFilterLogicEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableCustomFieldRequestFilterLogic", value)
}

// NewPatchedWritableCustomFieldRequestFilterLogicFromValue returns a pointer to a valid PatchedWritableCustomFieldRequestFilterLogic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableCustomFieldRequestFilterLogicFromValue(v string) (*PatchedWritableCustomFieldRequestFilterLogic, error) {
	ev := PatchedWritableCustomFieldRequestFilterLogic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableCustomFieldRequestFilterLogic: valid values are %v", v, AllowedPatchedWritableCustomFieldRequestFilterLogicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableCustomFieldRequestFilterLogic) IsValid() bool {
	for _, existing := range AllowedPatchedWritableCustomFieldRequestFilterLogicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableCustomFieldRequest_filter_logic value
func (v PatchedWritableCustomFieldRequestFilterLogic) Ptr() *PatchedWritableCustomFieldRequestFilterLogic {
	return &v
}

type NullablePatchedWritableCustomFieldRequestFilterLogic struct {
	value *PatchedWritableCustomFieldRequestFilterLogic
	isSet bool
}

func (v NullablePatchedWritableCustomFieldRequestFilterLogic) Get() *PatchedWritableCustomFieldRequestFilterLogic {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldRequestFilterLogic) Set(val *PatchedWritableCustomFieldRequestFilterLogic) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldRequestFilterLogic) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldRequestFilterLogic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldRequestFilterLogic(val *PatchedWritableCustomFieldRequestFilterLogic) *NullablePatchedWritableCustomFieldRequestFilterLogic {
	return &NullablePatchedWritableCustomFieldRequestFilterLogic{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldRequestFilterLogic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldRequestFilterLogic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
