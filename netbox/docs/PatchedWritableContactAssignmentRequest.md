# PatchedWritableContactAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**Contact** | Pointer to **int32** |  | [optional] 
**Role** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to [**ContactAssignmentPriorityValue**](ContactAssignmentPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewPatchedWritableContactAssignmentRequest

`func NewPatchedWritableContactAssignmentRequest() *PatchedWritableContactAssignmentRequest`

NewPatchedWritableContactAssignmentRequest instantiates a new PatchedWritableContactAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableContactAssignmentRequestWithDefaults

`func NewPatchedWritableContactAssignmentRequestWithDefaults() *PatchedWritableContactAssignmentRequest`

NewPatchedWritableContactAssignmentRequestWithDefaults instantiates a new PatchedWritableContactAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *PatchedWritableContactAssignmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedWritableContactAssignmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedWritableContactAssignmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedWritableContactAssignmentRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedWritableContactAssignmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedWritableContactAssignmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedWritableContactAssignmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedWritableContactAssignmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetContact

`func (o *PatchedWritableContactAssignmentRequest) GetContact() int32`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatchedWritableContactAssignmentRequest) GetContactOk() (*int32, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatchedWritableContactAssignmentRequest) SetContact(v int32)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatchedWritableContactAssignmentRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableContactAssignmentRequest) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableContactAssignmentRequest) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableContactAssignmentRequest) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableContactAssignmentRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPriority

`func (o *PatchedWritableContactAssignmentRequest) GetPriority() ContactAssignmentPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedWritableContactAssignmentRequest) GetPriorityOk() (*ContactAssignmentPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedWritableContactAssignmentRequest) SetPriority(v ContactAssignmentPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedWritableContactAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableContactAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableContactAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableContactAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableContactAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


