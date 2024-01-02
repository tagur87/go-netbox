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

// checks if the NestedCircuitType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedCircuitType{}

// NestedCircuitType Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedCircuitType struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	Slug                 string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedCircuitType NestedCircuitType

// NewNestedCircuitType instantiates a new NestedCircuitType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCircuitType(id int32, url string, display string, name string, slug string) *NestedCircuitType {
	this := NestedCircuitType{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedCircuitTypeWithDefaults instantiates a new NestedCircuitType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedCircuitTypeWithDefaults() *NestedCircuitType {
	this := NestedCircuitType{}
	return &this
}

// GetId returns the Id field value
func (o *NestedCircuitType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedCircuitType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedCircuitType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedCircuitType) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedCircuitType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedCircuitType) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedCircuitType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedCircuitType) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedCircuitType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedCircuitType) SetSlug(v string) {
	o.Slug = v
}

func (o NestedCircuitType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedCircuitType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedCircuitType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
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

	varNestedCircuitType := _NestedCircuitType{}

	err = json.Unmarshal(data, &varNestedCircuitType)

	if err != nil {
		return err
	}

	*o = NestedCircuitType(varNestedCircuitType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedCircuitType struct {
	value *NestedCircuitType
	isSet bool
}

func (v NullableNestedCircuitType) Get() *NestedCircuitType {
	return v.value
}

func (v *NullableNestedCircuitType) Set(val *NestedCircuitType) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCircuitType) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCircuitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCircuitType(val *NestedCircuitType) *NullableNestedCircuitType {
	return &NullableNestedCircuitType{value: val, isSet: true}
}

func (v NullableNestedCircuitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCircuitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
