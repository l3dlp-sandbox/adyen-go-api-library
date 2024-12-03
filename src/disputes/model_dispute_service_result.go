/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the DisputeServiceResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DisputeServiceResult{}

// DisputeServiceResult struct for DisputeServiceResult
type DisputeServiceResult struct {
	// The general error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Indicates whether the request succeeded.
	Success bool `json:"success"`
}

// NewDisputeServiceResult instantiates a new DisputeServiceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisputeServiceResult(success bool) *DisputeServiceResult {
	this := DisputeServiceResult{}
	this.Success = success
	return &this
}

// NewDisputeServiceResultWithDefaults instantiates a new DisputeServiceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisputeServiceResultWithDefaults() *DisputeServiceResult {
	this := DisputeServiceResult{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *DisputeServiceResult) GetErrorMessage() string {
	if o == nil || common.IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeServiceResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *DisputeServiceResult) HasErrorMessage() bool {
	if o != nil && !common.IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *DisputeServiceResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetSuccess returns the Success field value
func (o *DisputeServiceResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *DisputeServiceResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *DisputeServiceResult) SetSuccess(v bool) {
	o.Success = v
}

func (o DisputeServiceResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisputeServiceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableDisputeServiceResult struct {
	value *DisputeServiceResult
	isSet bool
}

func (v NullableDisputeServiceResult) Get() *DisputeServiceResult {
	return v.value
}

func (v *NullableDisputeServiceResult) Set(val *DisputeServiceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDisputeServiceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDisputeServiceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisputeServiceResult(val *DisputeServiceResult) *NullableDisputeServiceResult {
	return &NullableDisputeServiceResult{value: val, isSet: true}
}

func (v NullableDisputeServiceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisputeServiceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
