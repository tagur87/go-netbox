/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.2 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableDeviceRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableDeviceRoleRequest{}

// PatchedWritableDeviceRoleRequest Adds support for custom fields and tags.
type PatchedWritableDeviceRoleRequest struct {
	Name  *string `json:"name,omitempty"`
	Slug  *string `json:"slug,omitempty"`
	Color *string `json:"color,omitempty"`
	// Virtual machines may be assigned to this role
	VmRole               *bool                  `json:"vm_role,omitempty"`
	ConfigTemplate       NullableInt32          `json:"config_template,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableDeviceRoleRequest PatchedWritableDeviceRoleRequest

// NewPatchedWritableDeviceRoleRequest instantiates a new PatchedWritableDeviceRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableDeviceRoleRequest() *PatchedWritableDeviceRoleRequest {
	this := PatchedWritableDeviceRoleRequest{}
	return &this
}

// NewPatchedWritableDeviceRoleRequestWithDefaults instantiates a new PatchedWritableDeviceRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableDeviceRoleRequestWithDefaults() *PatchedWritableDeviceRoleRequest {
	this := PatchedWritableDeviceRoleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableDeviceRoleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceRoleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableDeviceRoleRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedWritableDeviceRoleRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceRoleRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedWritableDeviceRoleRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PatchedWritableDeviceRoleRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceRoleRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PatchedWritableDeviceRoleRequest) SetColor(v string) {
	o.Color = &v
}

// GetVmRole returns the VmRole field value if set, zero value otherwise.
func (o *PatchedWritableDeviceRoleRequest) GetVmRole() bool {
	if o == nil || IsNil(o.VmRole) {
		var ret bool
		return ret
	}
	return *o.VmRole
}

// GetVmRoleOk returns a tuple with the VmRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceRoleRequest) GetVmRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.VmRole) {
		return nil, false
	}
	return o.VmRole, true
}

// HasVmRole returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasVmRole() bool {
	if o != nil && !IsNil(o.VmRole) {
		return true
	}

	return false
}

// SetVmRole gets a reference to the given bool and assigns it to the VmRole field.
func (o *PatchedWritableDeviceRoleRequest) SetVmRole(v bool) {
	o.VmRole = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableDeviceRoleRequest) GetConfigTemplate() int32 {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret int32
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableDeviceRoleRequest) GetConfigTemplateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableInt32 and assigns it to the ConfigTemplate field.
func (o *PatchedWritableDeviceRoleRequest) SetConfigTemplate(v int32) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *PatchedWritableDeviceRoleRequest) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *PatchedWritableDeviceRoleRequest) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableDeviceRoleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableDeviceRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableDeviceRoleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceRoleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableDeviceRoleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableDeviceRoleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceRoleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableDeviceRoleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableDeviceRoleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableDeviceRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableDeviceRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.VmRole) {
		toSerialize["vm_role"] = o.VmRole
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
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

func (o *PatchedWritableDeviceRoleRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableDeviceRoleRequest := _PatchedWritableDeviceRoleRequest{}

	err = json.Unmarshal(data, &varPatchedWritableDeviceRoleRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableDeviceRoleRequest(varPatchedWritableDeviceRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		delete(additionalProperties, "vm_role")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableDeviceRoleRequest struct {
	value *PatchedWritableDeviceRoleRequest
	isSet bool
}

func (v NullablePatchedWritableDeviceRoleRequest) Get() *PatchedWritableDeviceRoleRequest {
	return v.value
}

func (v *NullablePatchedWritableDeviceRoleRequest) Set(val *PatchedWritableDeviceRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableDeviceRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableDeviceRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableDeviceRoleRequest(val *PatchedWritableDeviceRoleRequest) *NullablePatchedWritableDeviceRoleRequest {
	return &NullablePatchedWritableDeviceRoleRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableDeviceRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableDeviceRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}