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

// checks if the PatchedWritableVirtualDeviceContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableVirtualDeviceContextRequest{}

// PatchedWritableVirtualDeviceContextRequest Adds support for custom fields and tags.
type PatchedWritableVirtualDeviceContextRequest struct {
	Name   *string       `json:"name,omitempty"`
	Device NullableInt32 `json:"device,omitempty"`
	// Numeric identifier unique to the parent device
	Identifier           NullableInt32                                     `json:"identifier,omitempty"`
	Tenant               NullableInt32                                     `json:"tenant,omitempty"`
	PrimaryIp4           NullableInt32                                     `json:"primary_ip4,omitempty"`
	PrimaryIp6           NullableInt32                                     `json:"primary_ip6,omitempty"`
	Status               *PatchedWritableVirtualDeviceContextRequestStatus `json:"status,omitempty"`
	Description          *string                                           `json:"description,omitempty"`
	Comments             *string                                           `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                                `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                            `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableVirtualDeviceContextRequest PatchedWritableVirtualDeviceContextRequest

// NewPatchedWritableVirtualDeviceContextRequest instantiates a new PatchedWritableVirtualDeviceContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableVirtualDeviceContextRequest() *PatchedWritableVirtualDeviceContextRequest {
	this := PatchedWritableVirtualDeviceContextRequest{}
	return &this
}

// NewPatchedWritableVirtualDeviceContextRequestWithDefaults instantiates a new PatchedWritableVirtualDeviceContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableVirtualDeviceContextRequestWithDefaults() *PatchedWritableVirtualDeviceContextRequest {
	this := PatchedWritableVirtualDeviceContextRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableVirtualDeviceContextRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetName(v string) {
	o.Name = &v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVirtualDeviceContextRequest) GetDevice() int32 {
	if o == nil || IsNil(o.Device.Get()) {
		var ret int32
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVirtualDeviceContextRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableInt32 and assigns it to the Device field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetDevice(v int32) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVirtualDeviceContextRequest) GetIdentifier() int32 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int32
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVirtualDeviceContextRequest) GetIdentifierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt32 and assigns it to the Identifier field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetIdentifier(v int32) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVirtualDeviceContextRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVirtualDeviceContextRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVirtualDeviceContextRequest) GetPrimaryIp4() int32 {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret int32
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVirtualDeviceContextRequest) GetPrimaryIp4Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableInt32 and assigns it to the PrimaryIp4 field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetPrimaryIp4(v int32) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVirtualDeviceContextRequest) GetPrimaryIp6() int32 {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret int32
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVirtualDeviceContextRequest) GetPrimaryIp6Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableInt32 and assigns it to the PrimaryIp6 field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetPrimaryIp6(v int32) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *PatchedWritableVirtualDeviceContextRequest) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableVirtualDeviceContextRequest) GetStatus() PatchedWritableVirtualDeviceContextRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableVirtualDeviceContextRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) GetStatusOk() (*PatchedWritableVirtualDeviceContextRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableVirtualDeviceContextRequestStatus and assigns it to the Status field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetStatus(v PatchedWritableVirtualDeviceContextRequestStatus) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableVirtualDeviceContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableVirtualDeviceContextRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableVirtualDeviceContextRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableVirtualDeviceContextRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableVirtualDeviceContextRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableVirtualDeviceContextRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableVirtualDeviceContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableVirtualDeviceContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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

func (o *PatchedWritableVirtualDeviceContextRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableVirtualDeviceContextRequest := _PatchedWritableVirtualDeviceContextRequest{}

	err = json.Unmarshal(data, &varPatchedWritableVirtualDeviceContextRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableVirtualDeviceContextRequest(varPatchedWritableVirtualDeviceContextRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "device")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "primary_ip4")
		delete(additionalProperties, "primary_ip6")
		delete(additionalProperties, "status")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableVirtualDeviceContextRequest struct {
	value *PatchedWritableVirtualDeviceContextRequest
	isSet bool
}

func (v NullablePatchedWritableVirtualDeviceContextRequest) Get() *PatchedWritableVirtualDeviceContextRequest {
	return v.value
}

func (v *NullablePatchedWritableVirtualDeviceContextRequest) Set(val *PatchedWritableVirtualDeviceContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableVirtualDeviceContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableVirtualDeviceContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableVirtualDeviceContextRequest(val *PatchedWritableVirtualDeviceContextRequest) *NullablePatchedWritableVirtualDeviceContextRequest {
	return &NullablePatchedWritableVirtualDeviceContextRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableVirtualDeviceContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableVirtualDeviceContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
