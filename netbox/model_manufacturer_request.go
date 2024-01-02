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

// checks if the ManufacturerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManufacturerRequest{}

// ManufacturerRequest Adds support for custom fields and tags.
type ManufacturerRequest struct {
	Name                 string                 `json:"name"`
	Slug                 string                 `json:"slug"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManufacturerRequest ManufacturerRequest

// NewManufacturerRequest instantiates a new ManufacturerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManufacturerRequest(name string, slug string) *ManufacturerRequest {
	this := ManufacturerRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewManufacturerRequestWithDefaults instantiates a new ManufacturerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManufacturerRequestWithDefaults() *ManufacturerRequest {
	this := ManufacturerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ManufacturerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManufacturerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManufacturerRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *ManufacturerRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *ManufacturerRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *ManufacturerRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManufacturerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManufacturerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManufacturerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManufacturerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ManufacturerRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManufacturerRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ManufacturerRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ManufacturerRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ManufacturerRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManufacturerRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ManufacturerRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ManufacturerRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ManufacturerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManufacturerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
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

func (o *ManufacturerRequest) UnmarshalJSON(data []byte) (err error) {
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

	varManufacturerRequest := _ManufacturerRequest{}

	err = json.Unmarshal(data, &varManufacturerRequest)

	if err != nil {
		return err
	}

	*o = ManufacturerRequest(varManufacturerRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManufacturerRequest struct {
	value *ManufacturerRequest
	isSet bool
}

func (v NullableManufacturerRequest) Get() *ManufacturerRequest {
	return v.value
}

func (v *NullableManufacturerRequest) Set(val *ManufacturerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManufacturerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManufacturerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManufacturerRequest(val *ManufacturerRequest) *NullableManufacturerRequest {
	return &NullableManufacturerRequest{value: val, isSet: true}
}

func (v NullableManufacturerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManufacturerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
