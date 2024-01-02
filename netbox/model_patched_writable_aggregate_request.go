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

// checks if the PatchedWritableAggregateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableAggregateRequest{}

// PatchedWritableAggregateRequest Adds support for custom fields and tags.
type PatchedWritableAggregateRequest struct {
	Prefix *string `json:"prefix,omitempty"`
	// Regional Internet Registry responsible for this IP space
	Rir                  *int32                 `json:"rir,omitempty"`
	Tenant               NullableInt32          `json:"tenant,omitempty"`
	DateAdded            NullableString         `json:"date_added,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableAggregateRequest PatchedWritableAggregateRequest

// NewPatchedWritableAggregateRequest instantiates a new PatchedWritableAggregateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableAggregateRequest() *PatchedWritableAggregateRequest {
	this := PatchedWritableAggregateRequest{}
	return &this
}

// NewPatchedWritableAggregateRequestWithDefaults instantiates a new PatchedWritableAggregateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableAggregateRequestWithDefaults() *PatchedWritableAggregateRequest {
	this := PatchedWritableAggregateRequest{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PatchedWritableAggregateRequest) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableAggregateRequest) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PatchedWritableAggregateRequest) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRir returns the Rir field value if set, zero value otherwise.
func (o *PatchedWritableAggregateRequest) GetRir() int32 {
	if o == nil || IsNil(o.Rir) {
		var ret int32
		return ret
	}
	return *o.Rir
}

// GetRirOk returns a tuple with the Rir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableAggregateRequest) GetRirOk() (*int32, bool) {
	if o == nil || IsNil(o.Rir) {
		return nil, false
	}
	return o.Rir, true
}

// HasRir returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasRir() bool {
	if o != nil && !IsNil(o.Rir) {
		return true
	}

	return false
}

// SetRir gets a reference to the given int32 and assigns it to the Rir field.
func (o *PatchedWritableAggregateRequest) SetRir(v int32) {
	o.Rir = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableAggregateRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableAggregateRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritableAggregateRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableAggregateRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableAggregateRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableAggregateRequest) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded.Get()) {
		var ret string
		return ret
	}
	return *o.DateAdded.Get()
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableAggregateRequest) GetDateAddedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateAdded.Get(), o.DateAdded.IsSet()
}

// HasDateAdded returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasDateAdded() bool {
	if o != nil && o.DateAdded.IsSet() {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given NullableString and assigns it to the DateAdded field.
func (o *PatchedWritableAggregateRequest) SetDateAdded(v string) {
	o.DateAdded.Set(&v)
}

// SetDateAddedNil sets the value for DateAdded to be an explicit nil
func (o *PatchedWritableAggregateRequest) SetDateAddedNil() {
	o.DateAdded.Set(nil)
}

// UnsetDateAdded ensures that no value is present for DateAdded, not even an explicit nil
func (o *PatchedWritableAggregateRequest) UnsetDateAdded() {
	o.DateAdded.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableAggregateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableAggregateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableAggregateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableAggregateRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableAggregateRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableAggregateRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableAggregateRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableAggregateRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableAggregateRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableAggregateRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableAggregateRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableAggregateRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableAggregateRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableAggregateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableAggregateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Rir) {
		toSerialize["rir"] = o.Rir
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.DateAdded.IsSet() {
		toSerialize["date_added"] = o.DateAdded.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *PatchedWritableAggregateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableAggregateRequest := _PatchedWritableAggregateRequest{}

	err = json.Unmarshal(data, &varPatchedWritableAggregateRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableAggregateRequest(varPatchedWritableAggregateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "rir")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "date_added")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableAggregateRequest struct {
	value *PatchedWritableAggregateRequest
	isSet bool
}

func (v NullablePatchedWritableAggregateRequest) Get() *PatchedWritableAggregateRequest {
	return v.value
}

func (v *NullablePatchedWritableAggregateRequest) Set(val *PatchedWritableAggregateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableAggregateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableAggregateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableAggregateRequest(val *PatchedWritableAggregateRequest) *NullablePatchedWritableAggregateRequest {
	return &NullablePatchedWritableAggregateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableAggregateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableAggregateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
