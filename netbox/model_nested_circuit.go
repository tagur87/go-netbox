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

// checks if the NestedCircuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedCircuit{}

// NestedCircuit Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedCircuit struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// Unique circuit ID
	Cid                  string `json:"cid"`
	AdditionalProperties map[string]interface{}
}

type _NestedCircuit NestedCircuit

// NewNestedCircuit instantiates a new NestedCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCircuit(id int32, url string, display string, cid string) *NestedCircuit {
	this := NestedCircuit{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Cid = cid
	return &this
}

// NewNestedCircuitWithDefaults instantiates a new NestedCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedCircuitWithDefaults() *NestedCircuit {
	this := NestedCircuit{}
	return &this
}

// GetId returns the Id field value
func (o *NestedCircuit) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedCircuit) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedCircuit) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedCircuit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedCircuit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedCircuit) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedCircuit) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedCircuit) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedCircuit) SetDisplay(v string) {
	o.Display = v
}

// GetCid returns the Cid field value
func (o *NestedCircuit) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *NestedCircuit) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *NestedCircuit) SetCid(v string) {
	o.Cid = v
}

func (o NestedCircuit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedCircuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["cid"] = o.Cid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedCircuit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"cid",
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

	varNestedCircuit := _NestedCircuit{}

	err = json.Unmarshal(data, &varNestedCircuit)

	if err != nil {
		return err
	}

	*o = NestedCircuit(varNestedCircuit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "cid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedCircuit struct {
	value *NestedCircuit
	isSet bool
}

func (v NullableNestedCircuit) Get() *NestedCircuit {
	return v.value
}

func (v *NullableNestedCircuit) Set(val *NestedCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCircuit(val *NestedCircuit) *NullableNestedCircuit {
	return &NullableNestedCircuit{value: val, isSet: true}
}

func (v NullableNestedCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}