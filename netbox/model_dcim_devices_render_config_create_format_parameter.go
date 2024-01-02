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

// DcimDevicesRenderConfigCreateFormatParameter the model 'DcimDevicesRenderConfigCreateFormatParameter'
type DcimDevicesRenderConfigCreateFormatParameter string

// List of dcim_devices_render_config_create_format_parameter
const (
	DCIMDEVICESRENDERCONFIGCREATEFORMATPARAMETER_JSON DcimDevicesRenderConfigCreateFormatParameter = "json"
	DCIMDEVICESRENDERCONFIGCREATEFORMATPARAMETER_TXT  DcimDevicesRenderConfigCreateFormatParameter = "txt"
)

// All allowed values of DcimDevicesRenderConfigCreateFormatParameter enum
var AllowedDcimDevicesRenderConfigCreateFormatParameterEnumValues = []DcimDevicesRenderConfigCreateFormatParameter{
	"json",
	"txt",
}

func (v *DcimDevicesRenderConfigCreateFormatParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimDevicesRenderConfigCreateFormatParameter(value)
	for _, existing := range AllowedDcimDevicesRenderConfigCreateFormatParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimDevicesRenderConfigCreateFormatParameter", value)
}

// NewDcimDevicesRenderConfigCreateFormatParameterFromValue returns a pointer to a valid DcimDevicesRenderConfigCreateFormatParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimDevicesRenderConfigCreateFormatParameterFromValue(v string) (*DcimDevicesRenderConfigCreateFormatParameter, error) {
	ev := DcimDevicesRenderConfigCreateFormatParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimDevicesRenderConfigCreateFormatParameter: valid values are %v", v, AllowedDcimDevicesRenderConfigCreateFormatParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimDevicesRenderConfigCreateFormatParameter) IsValid() bool {
	for _, existing := range AllowedDcimDevicesRenderConfigCreateFormatParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_devices_render_config_create_format_parameter value
func (v DcimDevicesRenderConfigCreateFormatParameter) Ptr() *DcimDevicesRenderConfigCreateFormatParameter {
	return &v
}

type NullableDcimDevicesRenderConfigCreateFormatParameter struct {
	value *DcimDevicesRenderConfigCreateFormatParameter
	isSet bool
}

func (v NullableDcimDevicesRenderConfigCreateFormatParameter) Get() *DcimDevicesRenderConfigCreateFormatParameter {
	return v.value
}

func (v *NullableDcimDevicesRenderConfigCreateFormatParameter) Set(val *DcimDevicesRenderConfigCreateFormatParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimDevicesRenderConfigCreateFormatParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimDevicesRenderConfigCreateFormatParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimDevicesRenderConfigCreateFormatParameter(val *DcimDevicesRenderConfigCreateFormatParameter) *NullableDcimDevicesRenderConfigCreateFormatParameter {
	return &NullableDcimDevicesRenderConfigCreateFormatParameter{value: val, isSet: true}
}

func (v NullableDcimDevicesRenderConfigCreateFormatParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimDevicesRenderConfigCreateFormatParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
