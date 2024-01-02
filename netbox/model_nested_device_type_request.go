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

// checks if the NestedDeviceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedDeviceTypeRequest{}

// NestedDeviceTypeRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedDeviceTypeRequest struct {
	Model                string `json:"model"`
	Slug                 string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedDeviceTypeRequest NestedDeviceTypeRequest

// NewNestedDeviceTypeRequest instantiates a new NestedDeviceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDeviceTypeRequest(model string, slug string) *NestedDeviceTypeRequest {
	this := NestedDeviceTypeRequest{}
	this.Model = model
	this.Slug = slug
	return &this
}

// NewNestedDeviceTypeRequestWithDefaults instantiates a new NestedDeviceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDeviceTypeRequestWithDefaults() *NestedDeviceTypeRequest {
	this := NestedDeviceTypeRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *NestedDeviceTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *NestedDeviceTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *NestedDeviceTypeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceTypeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedDeviceTypeRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedDeviceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedDeviceTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
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

	varNestedDeviceTypeRequest := _NestedDeviceTypeRequest{}

	err = json.Unmarshal(data, &varNestedDeviceTypeRequest)

	if err != nil {
		return err
	}

	*o = NestedDeviceTypeRequest(varNestedDeviceTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedDeviceTypeRequest struct {
	value *NestedDeviceTypeRequest
	isSet bool
}

func (v NullableNestedDeviceTypeRequest) Get() *NestedDeviceTypeRequest {
	return v.value
}

func (v *NullableNestedDeviceTypeRequest) Set(val *NestedDeviceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDeviceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDeviceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDeviceTypeRequest(val *NestedDeviceTypeRequest) *NullableNestedDeviceTypeRequest {
	return &NullableNestedDeviceTypeRequest{value: val, isSet: true}
}

func (v NullableNestedDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDeviceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
