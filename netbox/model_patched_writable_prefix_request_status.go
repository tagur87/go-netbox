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

// PatchedWritablePrefixRequestStatus Operational status of this prefix  * `container` - Container * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated
type PatchedWritablePrefixRequestStatus string

// List of PatchedWritablePrefixRequest_status
const (
	PATCHEDWRITABLEPREFIXREQUESTSTATUS_CONTAINER  PatchedWritablePrefixRequestStatus = "container"
	PATCHEDWRITABLEPREFIXREQUESTSTATUS_ACTIVE     PatchedWritablePrefixRequestStatus = "active"
	PATCHEDWRITABLEPREFIXREQUESTSTATUS_RESERVED   PatchedWritablePrefixRequestStatus = "reserved"
	PATCHEDWRITABLEPREFIXREQUESTSTATUS_DEPRECATED PatchedWritablePrefixRequestStatus = "deprecated"
)

// All allowed values of PatchedWritablePrefixRequestStatus enum
var AllowedPatchedWritablePrefixRequestStatusEnumValues = []PatchedWritablePrefixRequestStatus{
	"container",
	"active",
	"reserved",
	"deprecated",
}

func (v *PatchedWritablePrefixRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePrefixRequestStatus(value)
	for _, existing := range AllowedPatchedWritablePrefixRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePrefixRequestStatus", value)
}

// NewPatchedWritablePrefixRequestStatusFromValue returns a pointer to a valid PatchedWritablePrefixRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePrefixRequestStatusFromValue(v string) (*PatchedWritablePrefixRequestStatus, error) {
	ev := PatchedWritablePrefixRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePrefixRequestStatus: valid values are %v", v, AllowedPatchedWritablePrefixRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePrefixRequestStatus) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePrefixRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePrefixRequest_status value
func (v PatchedWritablePrefixRequestStatus) Ptr() *PatchedWritablePrefixRequestStatus {
	return &v
}

type NullablePatchedWritablePrefixRequestStatus struct {
	value *PatchedWritablePrefixRequestStatus
	isSet bool
}

func (v NullablePatchedWritablePrefixRequestStatus) Get() *PatchedWritablePrefixRequestStatus {
	return v.value
}

func (v *NullablePatchedWritablePrefixRequestStatus) Set(val *PatchedWritablePrefixRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePrefixRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePrefixRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePrefixRequestStatus(val *PatchedWritablePrefixRequestStatus) *NullablePatchedWritablePrefixRequestStatus {
	return &NullablePatchedWritablePrefixRequestStatus{value: val, isSet: true}
}

func (v NullablePatchedWritablePrefixRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePrefixRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
