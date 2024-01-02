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

// checks if the NestedVLAN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVLAN{}

// NestedVLAN Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVLAN struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// Numeric VLAN ID (1-4094)
	Vid                  int32  `json:"vid"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedVLAN NestedVLAN

// NewNestedVLAN instantiates a new NestedVLAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVLAN(id int32, url string, display string, vid int32, name string) *NestedVLAN {
	this := NestedVLAN{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Vid = vid
	this.Name = name
	return &this
}

// NewNestedVLANWithDefaults instantiates a new NestedVLAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVLANWithDefaults() *NestedVLAN {
	this := NestedVLAN{}
	return &this
}

// GetId returns the Id field value
func (o *NestedVLAN) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedVLAN) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedVLAN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedVLAN) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedVLAN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedVLAN) SetDisplay(v string) {
	o.Display = v
}

// GetVid returns the Vid field value
func (o *NestedVLAN) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *NestedVLAN) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *NestedVLAN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVLAN) SetName(v string) {
	o.Name = v
}

func (o NestedVLAN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVLAN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["vid"] = o.Vid
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedVLAN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"vid",
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

	varNestedVLAN := _NestedVLAN{}

	err = json.Unmarshal(data, &varNestedVLAN)

	if err != nil {
		return err
	}

	*o = NestedVLAN(varNestedVLAN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedVLAN struct {
	value *NestedVLAN
	isSet bool
}

func (v NullableNestedVLAN) Get() *NestedVLAN {
	return v.value
}

func (v *NullableNestedVLAN) Set(val *NestedVLAN) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVLAN) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVLAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVLAN(val *NestedVLAN) *NullableNestedVLAN {
	return &NullableNestedVLAN{value: val, isSet: true}
}

func (v NullableNestedVLAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVLAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
