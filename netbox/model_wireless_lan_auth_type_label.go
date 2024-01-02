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

// WirelessLANAuthTypeLabel the model 'WirelessLANAuthTypeLabel'
type WirelessLANAuthTypeLabel string

// List of WirelessLAN_auth_type_label
const (
	WIRELESSLANAUTHTYPELABEL_OPEN              WirelessLANAuthTypeLabel = "Open"
	WIRELESSLANAUTHTYPELABEL_WEP               WirelessLANAuthTypeLabel = "WEP"
	WIRELESSLANAUTHTYPELABEL_WPA_PERSONAL__PSK WirelessLANAuthTypeLabel = "WPA Personal (PSK)"
	WIRELESSLANAUTHTYPELABEL_WPA_ENTERPRISE    WirelessLANAuthTypeLabel = "WPA Enterprise"
)

// All allowed values of WirelessLANAuthTypeLabel enum
var AllowedWirelessLANAuthTypeLabelEnumValues = []WirelessLANAuthTypeLabel{
	"Open",
	"WEP",
	"WPA Personal (PSK)",
	"WPA Enterprise",
}

func (v *WirelessLANAuthTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessLANAuthTypeLabel(value)
	for _, existing := range AllowedWirelessLANAuthTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessLANAuthTypeLabel", value)
}

// NewWirelessLANAuthTypeLabelFromValue returns a pointer to a valid WirelessLANAuthTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessLANAuthTypeLabelFromValue(v string) (*WirelessLANAuthTypeLabel, error) {
	ev := WirelessLANAuthTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessLANAuthTypeLabel: valid values are %v", v, AllowedWirelessLANAuthTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessLANAuthTypeLabel) IsValid() bool {
	for _, existing := range AllowedWirelessLANAuthTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WirelessLAN_auth_type_label value
func (v WirelessLANAuthTypeLabel) Ptr() *WirelessLANAuthTypeLabel {
	return &v
}

type NullableWirelessLANAuthTypeLabel struct {
	value *WirelessLANAuthTypeLabel
	isSet bool
}

func (v NullableWirelessLANAuthTypeLabel) Get() *WirelessLANAuthTypeLabel {
	return v.value
}

func (v *NullableWirelessLANAuthTypeLabel) Set(val *WirelessLANAuthTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLANAuthTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLANAuthTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLANAuthTypeLabel(val *WirelessLANAuthTypeLabel) *NullableWirelessLANAuthTypeLabel {
	return &NullableWirelessLANAuthTypeLabel{value: val, isSet: true}
}

func (v NullableWirelessLANAuthTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLANAuthTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
