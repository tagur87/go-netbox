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

// checks if the WritableModuleBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableModuleBayRequest{}

// WritableModuleBayRequest Adds support for custom fields and tags.
type WritableModuleBayRequest struct {
	Device          int32  `json:"device"`
	Name            string `json:"name"`
	InstalledModule int32  `json:"installed_module"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Identifier to reference when renaming installed components
	Position             *string                `json:"position,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableModuleBayRequest WritableModuleBayRequest

// NewWritableModuleBayRequest instantiates a new WritableModuleBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableModuleBayRequest(device int32, name string, installedModule int32) *WritableModuleBayRequest {
	this := WritableModuleBayRequest{}
	this.Device = device
	this.Name = name
	this.InstalledModule = installedModule
	return &this
}

// NewWritableModuleBayRequestWithDefaults instantiates a new WritableModuleBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableModuleBayRequestWithDefaults() *WritableModuleBayRequest {
	this := WritableModuleBayRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *WritableModuleBayRequest) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritableModuleBayRequest) SetDevice(v int32) {
	o.Device = v
}

// GetName returns the Name field value
func (o *WritableModuleBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableModuleBayRequest) SetName(v string) {
	o.Name = v
}

// GetInstalledModule returns the InstalledModule field value
func (o *WritableModuleBayRequest) GetInstalledModule() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InstalledModule
}

// GetInstalledModuleOk returns a tuple with the InstalledModule field value
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetInstalledModuleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstalledModule, true
}

// SetInstalledModule sets field value
func (o *WritableModuleBayRequest) SetInstalledModule(v int32) {
	o.InstalledModule = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableModuleBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableModuleBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableModuleBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *WritableModuleBayRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *WritableModuleBayRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *WritableModuleBayRequest) SetPosition(v string) {
	o.Position = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableModuleBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableModuleBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableModuleBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableModuleBayRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableModuleBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableModuleBayRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableModuleBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableModuleBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableModuleBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableModuleBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableModuleBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["name"] = o.Name
	toSerialize["installed_module"] = o.InstalledModule
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *WritableModuleBayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
		"name",
		"installed_module",
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

	varWritableModuleBayRequest := _WritableModuleBayRequest{}

	err = json.Unmarshal(data, &varWritableModuleBayRequest)

	if err != nil {
		return err
	}

	*o = WritableModuleBayRequest(varWritableModuleBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "installed_module")
		delete(additionalProperties, "label")
		delete(additionalProperties, "position")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableModuleBayRequest struct {
	value *WritableModuleBayRequest
	isSet bool
}

func (v NullableWritableModuleBayRequest) Get() *WritableModuleBayRequest {
	return v.value
}

func (v *NullableWritableModuleBayRequest) Set(val *WritableModuleBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableModuleBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableModuleBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableModuleBayRequest(val *WritableModuleBayRequest) *NullableWritableModuleBayRequest {
	return &NullableWritableModuleBayRequest{value: val, isSet: true}
}

func (v NullableWritableModuleBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableModuleBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
