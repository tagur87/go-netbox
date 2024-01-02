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

// PatchedWritablePowerFeedRequestType * `primary` - Primary * `redundant` - Redundant
type PatchedWritablePowerFeedRequestType string

// List of PatchedWritablePowerFeedRequest_type
const (
	PATCHEDWRITABLEPOWERFEEDREQUESTTYPE_PRIMARY   PatchedWritablePowerFeedRequestType = "primary"
	PATCHEDWRITABLEPOWERFEEDREQUESTTYPE_REDUNDANT PatchedWritablePowerFeedRequestType = "redundant"
)

// All allowed values of PatchedWritablePowerFeedRequestType enum
var AllowedPatchedWritablePowerFeedRequestTypeEnumValues = []PatchedWritablePowerFeedRequestType{
	"primary",
	"redundant",
}

func (v *PatchedWritablePowerFeedRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePowerFeedRequestType(value)
	for _, existing := range AllowedPatchedWritablePowerFeedRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePowerFeedRequestType", value)
}

// NewPatchedWritablePowerFeedRequestTypeFromValue returns a pointer to a valid PatchedWritablePowerFeedRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePowerFeedRequestTypeFromValue(v string) (*PatchedWritablePowerFeedRequestType, error) {
	ev := PatchedWritablePowerFeedRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePowerFeedRequestType: valid values are %v", v, AllowedPatchedWritablePowerFeedRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePowerFeedRequestType) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePowerFeedRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePowerFeedRequest_type value
func (v PatchedWritablePowerFeedRequestType) Ptr() *PatchedWritablePowerFeedRequestType {
	return &v
}

type NullablePatchedWritablePowerFeedRequestType struct {
	value *PatchedWritablePowerFeedRequestType
	isSet bool
}

func (v NullablePatchedWritablePowerFeedRequestType) Get() *PatchedWritablePowerFeedRequestType {
	return v.value
}

func (v *NullablePatchedWritablePowerFeedRequestType) Set(val *PatchedWritablePowerFeedRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerFeedRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerFeedRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerFeedRequestType(val *PatchedWritablePowerFeedRequestType) *NullablePatchedWritablePowerFeedRequestType {
	return &NullablePatchedWritablePowerFeedRequestType{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerFeedRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerFeedRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
