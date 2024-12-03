/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the UpdateNetworkTokenRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateNetworkTokenRequest{}

// UpdateNetworkTokenRequest struct for UpdateNetworkTokenRequest
type UpdateNetworkTokenRequest struct {
	// The new status of the network token. Possible values: **active**, **suspended**, **closed**. The **closed** status is final and cannot be changed.
	Status *string `json:"status,omitempty"`
}

// NewUpdateNetworkTokenRequest instantiates a new UpdateNetworkTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkTokenRequest() *UpdateNetworkTokenRequest {
	this := UpdateNetworkTokenRequest{}
	return &this
}

// NewUpdateNetworkTokenRequestWithDefaults instantiates a new UpdateNetworkTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkTokenRequestWithDefaults() *UpdateNetworkTokenRequest {
	this := UpdateNetworkTokenRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateNetworkTokenRequest) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkTokenRequest) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateNetworkTokenRequest) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateNetworkTokenRequest) SetStatus(v string) {
	o.Status = &v
}

func (o UpdateNetworkTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUpdateNetworkTokenRequest struct {
	value *UpdateNetworkTokenRequest
	isSet bool
}

func (v NullableUpdateNetworkTokenRequest) Get() *UpdateNetworkTokenRequest {
	return v.value
}

func (v *NullableUpdateNetworkTokenRequest) Set(val *UpdateNetworkTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkTokenRequest(val *UpdateNetworkTokenRequest) *NullableUpdateNetworkTokenRequest {
	return &NullableUpdateNetworkTokenRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *UpdateNetworkTokenRequest) isValidStatus() bool {
    var allowedEnumValues = []string{ "active", "suspended", "closed" }
    for _, allowed := range allowedEnumValues {
        if o.GetStatus() == allowed {
            return true
        }
    }
    return false
}

