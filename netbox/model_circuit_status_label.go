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

// CircuitStatusLabel the model 'CircuitStatusLabel'
type CircuitStatusLabel string

// List of Circuit_status_label
const (
	CIRCUITSTATUSLABEL_PLANNED        CircuitStatusLabel = "Planned"
	CIRCUITSTATUSLABEL_PROVISIONING   CircuitStatusLabel = "Provisioning"
	CIRCUITSTATUSLABEL_ACTIVE         CircuitStatusLabel = "Active"
	CIRCUITSTATUSLABEL_OFFLINE        CircuitStatusLabel = "Offline"
	CIRCUITSTATUSLABEL_DEPROVISIONING CircuitStatusLabel = "Deprovisioning"
	CIRCUITSTATUSLABEL_DECOMMISSIONED CircuitStatusLabel = "Decommissioned"
)

// All allowed values of CircuitStatusLabel enum
var AllowedCircuitStatusLabelEnumValues = []CircuitStatusLabel{
	"Planned",
	"Provisioning",
	"Active",
	"Offline",
	"Deprovisioning",
	"Decommissioned",
}

func (v *CircuitStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CircuitStatusLabel(value)
	for _, existing := range AllowedCircuitStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CircuitStatusLabel", value)
}

// NewCircuitStatusLabelFromValue returns a pointer to a valid CircuitStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCircuitStatusLabelFromValue(v string) (*CircuitStatusLabel, error) {
	ev := CircuitStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CircuitStatusLabel: valid values are %v", v, AllowedCircuitStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CircuitStatusLabel) IsValid() bool {
	for _, existing := range AllowedCircuitStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Circuit_status_label value
func (v CircuitStatusLabel) Ptr() *CircuitStatusLabel {
	return &v
}

type NullableCircuitStatusLabel struct {
	value *CircuitStatusLabel
	isSet bool
}

func (v NullableCircuitStatusLabel) Get() *CircuitStatusLabel {
	return v.value
}

func (v *NullableCircuitStatusLabel) Set(val *CircuitStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitStatusLabel(val *CircuitStatusLabel) *NullableCircuitStatusLabel {
	return &NullableCircuitStatusLabel{value: val, isSet: true}
}

func (v NullableCircuitStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
