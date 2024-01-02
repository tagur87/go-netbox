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

// CustomFieldTypeValue * `text` - Text * `longtext` - Text (long) * `integer` - Integer * `decimal` - Decimal * `boolean` - Boolean (true/false) * `date` - Date * `datetime` - Date & time * `url` - URL * `json` - JSON * `select` - Selection * `multiselect` - Multiple selection * `object` - Object * `multiobject` - Multiple objects
type CustomFieldTypeValue string

// List of CustomField_type_value
const (
	CUSTOMFIELDTYPEVALUE_TEXT        CustomFieldTypeValue = "text"
	CUSTOMFIELDTYPEVALUE_LONGTEXT    CustomFieldTypeValue = "longtext"
	CUSTOMFIELDTYPEVALUE_INTEGER     CustomFieldTypeValue = "integer"
	CUSTOMFIELDTYPEVALUE_DECIMAL     CustomFieldTypeValue = "decimal"
	CUSTOMFIELDTYPEVALUE_BOOLEAN     CustomFieldTypeValue = "boolean"
	CUSTOMFIELDTYPEVALUE_DATE        CustomFieldTypeValue = "date"
	CUSTOMFIELDTYPEVALUE_DATETIME    CustomFieldTypeValue = "datetime"
	CUSTOMFIELDTYPEVALUE_URL         CustomFieldTypeValue = "url"
	CUSTOMFIELDTYPEVALUE_JSON        CustomFieldTypeValue = "json"
	CUSTOMFIELDTYPEVALUE_SELECT      CustomFieldTypeValue = "select"
	CUSTOMFIELDTYPEVALUE_MULTISELECT CustomFieldTypeValue = "multiselect"
	CUSTOMFIELDTYPEVALUE_OBJECT      CustomFieldTypeValue = "object"
	CUSTOMFIELDTYPEVALUE_MULTIOBJECT CustomFieldTypeValue = "multiobject"
)

// All allowed values of CustomFieldTypeValue enum
var AllowedCustomFieldTypeValueEnumValues = []CustomFieldTypeValue{
	"text",
	"longtext",
	"integer",
	"decimal",
	"boolean",
	"date",
	"datetime",
	"url",
	"json",
	"select",
	"multiselect",
	"object",
	"multiobject",
}

func (v *CustomFieldTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldTypeValue(value)
	for _, existing := range AllowedCustomFieldTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldTypeValue", value)
}

// NewCustomFieldTypeValueFromValue returns a pointer to a valid CustomFieldTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldTypeValueFromValue(v string) (*CustomFieldTypeValue, error) {
	ev := CustomFieldTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldTypeValue: valid values are %v", v, AllowedCustomFieldTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldTypeValue) IsValid() bool {
	for _, existing := range AllowedCustomFieldTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomField_type_value value
func (v CustomFieldTypeValue) Ptr() *CustomFieldTypeValue {
	return &v
}

type NullableCustomFieldTypeValue struct {
	value *CustomFieldTypeValue
	isSet bool
}

func (v NullableCustomFieldTypeValue) Get() *CustomFieldTypeValue {
	return v.value
}

func (v *NullableCustomFieldTypeValue) Set(val *CustomFieldTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldTypeValue(val *CustomFieldTypeValue) *NullableCustomFieldTypeValue {
	return &NullableCustomFieldTypeValue{value: val, isSet: true}
}

func (v NullableCustomFieldTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
