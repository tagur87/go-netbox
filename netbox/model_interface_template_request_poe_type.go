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

// InterfaceTemplateRequestPoeType * `type1-ieee802.3af` - 802.3af (Type 1) * `type2-ieee802.3at` - 802.3at (Type 2) * `type3-ieee802.3bt` - 802.3bt (Type 3) * `type4-ieee802.3bt` - 802.3bt (Type 4) * `passive-24v-2pair` - Passive 24V (2-pair) * `passive-24v-4pair` - Passive 24V (4-pair) * `passive-48v-2pair` - Passive 48V (2-pair) * `passive-48v-4pair` - Passive 48V (4-pair)
type InterfaceTemplateRequestPoeType string

// List of InterfaceTemplateRequest_poe_type
const (
	INTERFACETEMPLATEREQUESTPOETYPE_TYPE1_IEEE802_3AF InterfaceTemplateRequestPoeType = "type1-ieee802.3af"
	INTERFACETEMPLATEREQUESTPOETYPE_TYPE2_IEEE802_3AT InterfaceTemplateRequestPoeType = "type2-ieee802.3at"
	INTERFACETEMPLATEREQUESTPOETYPE_TYPE3_IEEE802_3BT InterfaceTemplateRequestPoeType = "type3-ieee802.3bt"
	INTERFACETEMPLATEREQUESTPOETYPE_TYPE4_IEEE802_3BT InterfaceTemplateRequestPoeType = "type4-ieee802.3bt"
	INTERFACETEMPLATEREQUESTPOETYPE_PASSIVE_24V_2PAIR InterfaceTemplateRequestPoeType = "passive-24v-2pair"
	INTERFACETEMPLATEREQUESTPOETYPE_PASSIVE_24V_4PAIR InterfaceTemplateRequestPoeType = "passive-24v-4pair"
	INTERFACETEMPLATEREQUESTPOETYPE_PASSIVE_48V_2PAIR InterfaceTemplateRequestPoeType = "passive-48v-2pair"
	INTERFACETEMPLATEREQUESTPOETYPE_PASSIVE_48V_4PAIR InterfaceTemplateRequestPoeType = "passive-48v-4pair"
	INTERFACETEMPLATEREQUESTPOETYPE_EMPTY             InterfaceTemplateRequestPoeType = ""
	INTERFACETEMPLATEREQUESTPOETYPE_NULL              InterfaceTemplateRequestPoeType = "null"
)

// All allowed values of InterfaceTemplateRequestPoeType enum
var AllowedInterfaceTemplateRequestPoeTypeEnumValues = []InterfaceTemplateRequestPoeType{
	"type1-ieee802.3af",
	"type2-ieee802.3at",
	"type3-ieee802.3bt",
	"type4-ieee802.3bt",
	"passive-24v-2pair",
	"passive-24v-4pair",
	"passive-48v-2pair",
	"passive-48v-4pair",
	"",
	"null",
}

func (v *InterfaceTemplateRequestPoeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceTemplateRequestPoeType(value)
	for _, existing := range AllowedInterfaceTemplateRequestPoeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceTemplateRequestPoeType", value)
}

// NewInterfaceTemplateRequestPoeTypeFromValue returns a pointer to a valid InterfaceTemplateRequestPoeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceTemplateRequestPoeTypeFromValue(v string) (*InterfaceTemplateRequestPoeType, error) {
	ev := InterfaceTemplateRequestPoeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceTemplateRequestPoeType: valid values are %v", v, AllowedInterfaceTemplateRequestPoeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceTemplateRequestPoeType) IsValid() bool {
	for _, existing := range AllowedInterfaceTemplateRequestPoeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceTemplateRequest_poe_type value
func (v InterfaceTemplateRequestPoeType) Ptr() *InterfaceTemplateRequestPoeType {
	return &v
}

type NullableInterfaceTemplateRequestPoeType struct {
	value *InterfaceTemplateRequestPoeType
	isSet bool
}

func (v NullableInterfaceTemplateRequestPoeType) Get() *InterfaceTemplateRequestPoeType {
	return v.value
}

func (v *NullableInterfaceTemplateRequestPoeType) Set(val *InterfaceTemplateRequestPoeType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplateRequestPoeType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplateRequestPoeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplateRequestPoeType(val *InterfaceTemplateRequestPoeType) *NullableInterfaceTemplateRequestPoeType {
	return &NullableInterfaceTemplateRequestPoeType{value: val, isSet: true}
}

func (v NullableInterfaceTemplateRequestPoeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplateRequestPoeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
