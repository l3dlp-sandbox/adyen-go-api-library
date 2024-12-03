/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the UpdateSplitConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateSplitConfigurationRequest{}

// UpdateSplitConfigurationRequest struct for UpdateSplitConfigurationRequest
type UpdateSplitConfigurationRequest struct {
	// Your description for the split configuration.
	Description string `json:"description"`
}

// NewUpdateSplitConfigurationRequest instantiates a new UpdateSplitConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSplitConfigurationRequest(description string) *UpdateSplitConfigurationRequest {
	this := UpdateSplitConfigurationRequest{}
	this.Description = description
	return &this
}

// NewUpdateSplitConfigurationRequestWithDefaults instantiates a new UpdateSplitConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSplitConfigurationRequestWithDefaults() *UpdateSplitConfigurationRequest {
	this := UpdateSplitConfigurationRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *UpdateSplitConfigurationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateSplitConfigurationRequest) SetDescription(v string) {
	o.Description = v
}

func (o UpdateSplitConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSplitConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableUpdateSplitConfigurationRequest struct {
	value *UpdateSplitConfigurationRequest
	isSet bool
}

func (v NullableUpdateSplitConfigurationRequest) Get() *UpdateSplitConfigurationRequest {
	return v.value
}

func (v *NullableUpdateSplitConfigurationRequest) Set(val *UpdateSplitConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSplitConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSplitConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSplitConfigurationRequest(val *UpdateSplitConfigurationRequest) *NullableUpdateSplitConfigurationRequest {
	return &NullableUpdateSplitConfigurationRequest{value: val, isSet: true}
}

func (v NullableUpdateSplitConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSplitConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



