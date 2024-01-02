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

// checks if the WritablePlatformRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritablePlatformRequest{}

// WritablePlatformRequest Adds support for custom fields and tags.
type WritablePlatformRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	// Optionally limit this platform to devices of a certain manufacturer
	Manufacturer         NullableInt32          `json:"manufacturer,omitempty"`
	ConfigTemplate       NullableInt32          `json:"config_template,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritablePlatformRequest WritablePlatformRequest

// NewWritablePlatformRequest instantiates a new WritablePlatformRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritablePlatformRequest(name string, slug string) *WritablePlatformRequest {
	this := WritablePlatformRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewWritablePlatformRequestWithDefaults instantiates a new WritablePlatformRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritablePlatformRequestWithDefaults() *WritablePlatformRequest {
	this := WritablePlatformRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritablePlatformRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritablePlatformRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritablePlatformRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WritablePlatformRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritablePlatformRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritablePlatformRequest) SetSlug(v string) {
	o.Slug = v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePlatformRequest) GetManufacturer() int32 {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret int32
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePlatformRequest) GetManufacturerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *WritablePlatformRequest) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableInt32 and assigns it to the Manufacturer field.
func (o *WritablePlatformRequest) SetManufacturer(v int32) {
	o.Manufacturer.Set(&v)
}

// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *WritablePlatformRequest) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *WritablePlatformRequest) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePlatformRequest) GetConfigTemplate() int32 {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret int32
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePlatformRequest) GetConfigTemplateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *WritablePlatformRequest) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableInt32 and assigns it to the ConfigTemplate field.
func (o *WritablePlatformRequest) SetConfigTemplate(v int32) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *WritablePlatformRequest) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *WritablePlatformRequest) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritablePlatformRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatformRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritablePlatformRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritablePlatformRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritablePlatformRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatformRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritablePlatformRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritablePlatformRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritablePlatformRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatformRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritablePlatformRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritablePlatformRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritablePlatformRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritablePlatformRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
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

func (o *WritablePlatformRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varWritablePlatformRequest := _WritablePlatformRequest{}

	err = json.Unmarshal(data, &varWritablePlatformRequest)

	if err != nil {
		return err
	}

	*o = WritablePlatformRequest(varWritablePlatformRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritablePlatformRequest struct {
	value *WritablePlatformRequest
	isSet bool
}

func (v NullableWritablePlatformRequest) Get() *WritablePlatformRequest {
	return v.value
}

func (v *NullableWritablePlatformRequest) Set(val *WritablePlatformRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritablePlatformRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritablePlatformRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritablePlatformRequest(val *WritablePlatformRequest) *NullableWritablePlatformRequest {
	return &NullableWritablePlatformRequest{value: val, isSet: true}
}

func (v NullableWritablePlatformRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritablePlatformRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
