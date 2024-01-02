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

// InterfacePoeTypeLabel the model 'InterfacePoeTypeLabel'
type InterfacePoeTypeLabel string

// List of Interface_poe_type_label
const (
	INTERFACEPOETYPELABEL__802_3AF__TYPE_1     InterfacePoeTypeLabel = "802.3af (Type 1)"
	INTERFACEPOETYPELABEL__802_3AT__TYPE_2     InterfacePoeTypeLabel = "802.3at (Type 2)"
	INTERFACEPOETYPELABEL__802_3BT__TYPE_3     InterfacePoeTypeLabel = "802.3bt (Type 3)"
	INTERFACEPOETYPELABEL__802_3BT__TYPE_4     InterfacePoeTypeLabel = "802.3bt (Type 4)"
	INTERFACEPOETYPELABEL_PASSIVE_24_V__2_PAIR InterfacePoeTypeLabel = "Passive 24V (2-pair)"
	INTERFACEPOETYPELABEL_PASSIVE_24_V__4_PAIR InterfacePoeTypeLabel = "Passive 24V (4-pair)"
	INTERFACEPOETYPELABEL_PASSIVE_48_V__2_PAIR InterfacePoeTypeLabel = "Passive 48V (2-pair)"
	INTERFACEPOETYPELABEL_PASSIVE_48_V__4_PAIR InterfacePoeTypeLabel = "Passive 48V (4-pair)"
)

// All allowed values of InterfacePoeTypeLabel enum
var AllowedInterfacePoeTypeLabelEnumValues = []InterfacePoeTypeLabel{
	"802.3af (Type 1)",
	"802.3at (Type 2)",
	"802.3bt (Type 3)",
	"802.3bt (Type 4)",
	"Passive 24V (2-pair)",
	"Passive 24V (4-pair)",
	"Passive 48V (2-pair)",
	"Passive 48V (4-pair)",
}

func (v *InterfacePoeTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfacePoeTypeLabel(value)
	for _, existing := range AllowedInterfacePoeTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfacePoeTypeLabel", value)
}

// NewInterfacePoeTypeLabelFromValue returns a pointer to a valid InterfacePoeTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfacePoeTypeLabelFromValue(v string) (*InterfacePoeTypeLabel, error) {
	ev := InterfacePoeTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfacePoeTypeLabel: valid values are %v", v, AllowedInterfacePoeTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfacePoeTypeLabel) IsValid() bool {
	for _, existing := range AllowedInterfacePoeTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_poe_type_label value
func (v InterfacePoeTypeLabel) Ptr() *InterfacePoeTypeLabel {
	return &v
}

type NullableInterfacePoeTypeLabel struct {
	value *InterfacePoeTypeLabel
	isSet bool
}

func (v NullableInterfacePoeTypeLabel) Get() *InterfacePoeTypeLabel {
	return v.value
}

func (v *NullableInterfacePoeTypeLabel) Set(val *InterfacePoeTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfacePoeTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfacePoeTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfacePoeTypeLabel(val *InterfacePoeTypeLabel) *NullableInterfacePoeTypeLabel {
	return &NullableInterfacePoeTypeLabel{value: val, isSet: true}
}

func (v NullableInterfacePoeTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfacePoeTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}