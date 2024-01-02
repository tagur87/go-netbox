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

// checks if the WritablePowerOutletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritablePowerOutletRequest{}

// WritablePowerOutletRequest Adds support for custom fields and tags.
type WritablePowerOutletRequest struct {
	Device int32         `json:"device"`
	Module NullableInt32 `json:"module,omitempty"`
	Name   string        `json:"name"`
	// Physical label
	Label       *string                                   `json:"label,omitempty"`
	Type        *PatchedWritablePowerOutletRequestType    `json:"type,omitempty"`
	PowerPort   NullableInt32                             `json:"power_port,omitempty"`
	FeedLeg     *PatchedWritablePowerOutletRequestFeedLeg `json:"feed_leg,omitempty"`
	Description *string                                   `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritablePowerOutletRequest WritablePowerOutletRequest

// NewWritablePowerOutletRequest instantiates a new WritablePowerOutletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritablePowerOutletRequest(device int32, name string) *WritablePowerOutletRequest {
	this := WritablePowerOutletRequest{}
	this.Device = device
	this.Name = name
	return &this
}

// NewWritablePowerOutletRequestWithDefaults instantiates a new WritablePowerOutletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritablePowerOutletRequestWithDefaults() *WritablePowerOutletRequest {
	this := WritablePowerOutletRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *WritablePowerOutletRequest) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritablePowerOutletRequest) SetDevice(v int32) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePowerOutletRequest) GetModule() int32 {
	if o == nil || IsNil(o.Module.Get()) {
		var ret int32
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerOutletRequest) GetModuleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableInt32 and assigns it to the Module field.
func (o *WritablePowerOutletRequest) SetModule(v int32) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *WritablePowerOutletRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *WritablePowerOutletRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *WritablePowerOutletRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritablePowerOutletRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritablePowerOutletRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritablePowerOutletRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritablePowerOutletRequest) GetType() PatchedWritablePowerOutletRequestType {
	if o == nil || IsNil(o.Type) {
		var ret PatchedWritablePowerOutletRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetTypeOk() (*PatchedWritablePowerOutletRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritablePowerOutletRequestType and assigns it to the Type field.
func (o *WritablePowerOutletRequest) SetType(v PatchedWritablePowerOutletRequestType) {
	o.Type = &v
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePowerOutletRequest) GetPowerPort() int32 {
	if o == nil || IsNil(o.PowerPort.Get()) {
		var ret int32
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerOutletRequest) GetPowerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableInt32 and assigns it to the PowerPort field.
func (o *WritablePowerOutletRequest) SetPowerPort(v int32) {
	o.PowerPort.Set(&v)
}

// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *WritablePowerOutletRequest) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *WritablePowerOutletRequest) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise.
func (o *WritablePowerOutletRequest) GetFeedLeg() PatchedWritablePowerOutletRequestFeedLeg {
	if o == nil || IsNil(o.FeedLeg) {
		var ret PatchedWritablePowerOutletRequestFeedLeg
		return ret
	}
	return *o.FeedLeg
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetFeedLegOk() (*PatchedWritablePowerOutletRequestFeedLeg, bool) {
	if o == nil || IsNil(o.FeedLeg) {
		return nil, false
	}
	return o.FeedLeg, true
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasFeedLeg() bool {
	if o != nil && !IsNil(o.FeedLeg) {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given PatchedWritablePowerOutletRequestFeedLeg and assigns it to the FeedLeg field.
func (o *WritablePowerOutletRequest) SetFeedLeg(v PatchedWritablePowerOutletRequestFeedLeg) {
	o.FeedLeg = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritablePowerOutletRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritablePowerOutletRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *WritablePowerOutletRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *WritablePowerOutletRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritablePowerOutletRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritablePowerOutletRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritablePowerOutletRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutletRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritablePowerOutletRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritablePowerOutletRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritablePowerOutletRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritablePowerOutletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
	}
	if !IsNil(o.FeedLeg) {
		toSerialize["feed_leg"] = o.FeedLeg
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritablePowerOutletRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
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

	varWritablePowerOutletRequest := _WritablePowerOutletRequest{}

	err = json.Unmarshal(data, &varWritablePowerOutletRequest)

	if err != nil {
		return err
	}

	*o = WritablePowerOutletRequest(varWritablePowerOutletRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "power_port")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritablePowerOutletRequest struct {
	value *WritablePowerOutletRequest
	isSet bool
}

func (v NullableWritablePowerOutletRequest) Get() *WritablePowerOutletRequest {
	return v.value
}

func (v *NullableWritablePowerOutletRequest) Set(val *WritablePowerOutletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritablePowerOutletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritablePowerOutletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritablePowerOutletRequest(val *WritablePowerOutletRequest) *NullableWritablePowerOutletRequest {
	return &NullableWritablePowerOutletRequest{value: val, isSet: true}
}

func (v NullableWritablePowerOutletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritablePowerOutletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
