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

// checks if the NestedProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedProviderRequest{}

// NestedProviderRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedProviderRequest struct {
	// Full name of the provider
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedProviderRequest NestedProviderRequest

// NewNestedProviderRequest instantiates a new NestedProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedProviderRequest(name string, slug string) *NestedProviderRequest {
	this := NestedProviderRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedProviderRequestWithDefaults instantiates a new NestedProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedProviderRequestWithDefaults() *NestedProviderRequest {
	this := NestedProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedProviderRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedProviderRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedProviderRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedProviderRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedProviderRequest) UnmarshalJSON(data []byte) (err error) {
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

	varNestedProviderRequest := _NestedProviderRequest{}

	err = json.Unmarshal(data, &varNestedProviderRequest)

	if err != nil {
		return err
	}

	*o = NestedProviderRequest(varNestedProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedProviderRequest struct {
	value *NestedProviderRequest
	isSet bool
}

func (v NullableNestedProviderRequest) Get() *NestedProviderRequest {
	return v.value
}

func (v *NullableNestedProviderRequest) Set(val *NestedProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedProviderRequest(val *NestedProviderRequest) *NullableNestedProviderRequest {
	return &NullableNestedProviderRequest{value: val, isSet: true}
}

func (v NullableNestedProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
