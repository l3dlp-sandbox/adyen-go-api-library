/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the ListNetworkTokensResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListNetworkTokensResponse{}

// ListNetworkTokensResponse struct for ListNetworkTokensResponse
type ListNetworkTokensResponse struct {
	// List of network tokens.
	NetworkTokens []NetworkToken `json:"networkTokens,omitempty"`
}

// NewListNetworkTokensResponse instantiates a new ListNetworkTokensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworkTokensResponse() *ListNetworkTokensResponse {
	this := ListNetworkTokensResponse{}
	return &this
}

// NewListNetworkTokensResponseWithDefaults instantiates a new ListNetworkTokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworkTokensResponseWithDefaults() *ListNetworkTokensResponse {
	this := ListNetworkTokensResponse{}
	return &this
}

// GetNetworkTokens returns the NetworkTokens field value if set, zero value otherwise.
func (o *ListNetworkTokensResponse) GetNetworkTokens() []NetworkToken {
	if o == nil || common.IsNil(o.NetworkTokens) {
		var ret []NetworkToken
		return ret
	}
	return o.NetworkTokens
}

// GetNetworkTokensOk returns a tuple with the NetworkTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkTokensResponse) GetNetworkTokensOk() ([]NetworkToken, bool) {
	if o == nil || common.IsNil(o.NetworkTokens) {
		return nil, false
	}
	return o.NetworkTokens, true
}

// HasNetworkTokens returns a boolean if a field has been set.
func (o *ListNetworkTokensResponse) HasNetworkTokens() bool {
	if o != nil && !common.IsNil(o.NetworkTokens) {
		return true
	}

	return false
}

// SetNetworkTokens gets a reference to the given []NetworkToken and assigns it to the NetworkTokens field.
func (o *ListNetworkTokensResponse) SetNetworkTokens(v []NetworkToken) {
	o.NetworkTokens = v
}

func (o ListNetworkTokensResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNetworkTokensResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NetworkTokens) {
		toSerialize["networkTokens"] = o.NetworkTokens
	}
	return toSerialize, nil
}

type NullableListNetworkTokensResponse struct {
	value *ListNetworkTokensResponse
	isSet bool
}

func (v NullableListNetworkTokensResponse) Get() *ListNetworkTokensResponse {
	return v.value
}

func (v *NullableListNetworkTokensResponse) Set(val *ListNetworkTokensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListNetworkTokensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListNetworkTokensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNetworkTokensResponse(val *ListNetworkTokensResponse) *NullableListNetworkTokensResponse {
	return &NullableListNetworkTokensResponse{value: val, isSet: true}
}

func (v NullableListNetworkTokensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNetworkTokensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
