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

// checks if the PowerOutletTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerOutletTemplateRequest{}

// PowerOutletTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PowerOutletTemplateRequest struct {
	DeviceType NullableNestedDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableNestedModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string                                `json:"label,omitempty"`
	Type                 NullablePowerOutletRequestType         `json:"type,omitempty"`
	PowerPort            NullableNestedPowerPortTemplateRequest `json:"power_port,omitempty"`
	FeedLeg              NullablePowerOutletRequestFeedLeg      `json:"feed_leg,omitempty"`
	Description          *string                                `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerOutletTemplateRequest PowerOutletTemplateRequest

// NewPowerOutletTemplateRequest instantiates a new PowerOutletTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerOutletTemplateRequest(name string) *PowerOutletTemplateRequest {
	this := PowerOutletTemplateRequest{}
	this.Name = name
	return &this
}

// NewPowerOutletTemplateRequestWithDefaults instantiates a new PowerOutletTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerOutletTemplateRequestWithDefaults() *PowerOutletTemplateRequest {
	this := PowerOutletTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletTemplateRequest) GetDeviceType() NestedDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret NestedDeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletTemplateRequest) GetDeviceTypeOk() (*NestedDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PowerOutletTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableNestedDeviceTypeRequest and assigns it to the DeviceType field.
func (o *PowerOutletTemplateRequest) SetDeviceType(v NestedDeviceTypeRequest) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *PowerOutletTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *PowerOutletTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletTemplateRequest) GetModuleType() NestedModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret NestedModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletTemplateRequest) GetModuleTypeOk() (*NestedModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *PowerOutletTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableNestedModuleTypeRequest and assigns it to the ModuleType field.
func (o *PowerOutletTemplateRequest) SetModuleType(v NestedModuleTypeRequest) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *PowerOutletTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *PowerOutletTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *PowerOutletTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerOutletTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerOutletTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerOutletTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerOutletTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletTemplateRequest) GetType() PowerOutletRequestType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret PowerOutletRequestType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletTemplateRequest) GetTypeOk() (*PowerOutletRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PowerOutletTemplateRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullablePowerOutletRequestType and assigns it to the Type field.
func (o *PowerOutletTemplateRequest) SetType(v PowerOutletRequestType) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *PowerOutletTemplateRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PowerOutletTemplateRequest) UnsetType() {
	o.Type.Unset()
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletTemplateRequest) GetPowerPort() NestedPowerPortTemplateRequest {
	if o == nil || IsNil(o.PowerPort.Get()) {
		var ret NestedPowerPortTemplateRequest
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletTemplateRequest) GetPowerPortOk() (*NestedPowerPortTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *PowerOutletTemplateRequest) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableNestedPowerPortTemplateRequest and assigns it to the PowerPort field.
func (o *PowerOutletTemplateRequest) SetPowerPort(v NestedPowerPortTemplateRequest) {
	o.PowerPort.Set(&v)
}

// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *PowerOutletTemplateRequest) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *PowerOutletTemplateRequest) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletTemplateRequest) GetFeedLeg() PowerOutletRequestFeedLeg {
	if o == nil || IsNil(o.FeedLeg.Get()) {
		var ret PowerOutletRequestFeedLeg
		return ret
	}
	return *o.FeedLeg.Get()
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletTemplateRequest) GetFeedLegOk() (*PowerOutletRequestFeedLeg, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeedLeg.Get(), o.FeedLeg.IsSet()
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *PowerOutletTemplateRequest) HasFeedLeg() bool {
	if o != nil && o.FeedLeg.IsSet() {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given NullablePowerOutletRequestFeedLeg and assigns it to the FeedLeg field.
func (o *PowerOutletTemplateRequest) SetFeedLeg(v PowerOutletRequestFeedLeg) {
	o.FeedLeg.Set(&v)
}

// SetFeedLegNil sets the value for FeedLeg to be an explicit nil
func (o *PowerOutletTemplateRequest) SetFeedLegNil() {
	o.FeedLeg.Set(nil)
}

// UnsetFeedLeg ensures that no value is present for FeedLeg, not even an explicit nil
func (o *PowerOutletTemplateRequest) UnsetFeedLeg() {
	o.FeedLeg.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerOutletTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerOutletTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerOutletTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PowerOutletTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerOutletTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
	}
	if o.FeedLeg.IsSet() {
		toSerialize["feed_leg"] = o.FeedLeg.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PowerOutletTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPowerOutletTemplateRequest := _PowerOutletTemplateRequest{}

	err = json.Unmarshal(data, &varPowerOutletTemplateRequest)

	if err != nil {
		return err
	}

	*o = PowerOutletTemplateRequest(varPowerOutletTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "power_port")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerOutletTemplateRequest struct {
	value *PowerOutletTemplateRequest
	isSet bool
}

func (v NullablePowerOutletTemplateRequest) Get() *PowerOutletTemplateRequest {
	return v.value
}

func (v *NullablePowerOutletTemplateRequest) Set(val *PowerOutletTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletTemplateRequest(val *PowerOutletTemplateRequest) *NullablePowerOutletTemplateRequest {
	return &NullablePowerOutletTemplateRequest{value: val, isSet: true}
}

func (v NullablePowerOutletTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
