/*
Authentication webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acswebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the RelayedAuthenticationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RelayedAuthenticationResponse{}

// RelayedAuthenticationResponse struct for RelayedAuthenticationResponse
type RelayedAuthenticationResponse struct {
	AuthenticationDecision AuthenticationDecision `json:"authenticationDecision"`
}

// NewRelayedAuthenticationResponse instantiates a new RelayedAuthenticationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelayedAuthenticationResponse(authenticationDecision AuthenticationDecision) *RelayedAuthenticationResponse {
	this := RelayedAuthenticationResponse{}
	this.AuthenticationDecision = authenticationDecision
	return &this
}

// NewRelayedAuthenticationResponseWithDefaults instantiates a new RelayedAuthenticationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelayedAuthenticationResponseWithDefaults() *RelayedAuthenticationResponse {
	this := RelayedAuthenticationResponse{}
	return &this
}

// GetAuthenticationDecision returns the AuthenticationDecision field value
func (o *RelayedAuthenticationResponse) GetAuthenticationDecision() AuthenticationDecision {
	if o == nil {
		var ret AuthenticationDecision
		return ret
	}

	return o.AuthenticationDecision
}

// GetAuthenticationDecisionOk returns a tuple with the AuthenticationDecision field value
// and a boolean to check if the value has been set.
func (o *RelayedAuthenticationResponse) GetAuthenticationDecisionOk() (*AuthenticationDecision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationDecision, true
}

// SetAuthenticationDecision sets field value
func (o *RelayedAuthenticationResponse) SetAuthenticationDecision(v AuthenticationDecision) {
	o.AuthenticationDecision = v
}

func (o RelayedAuthenticationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelayedAuthenticationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticationDecision"] = o.AuthenticationDecision
	return toSerialize, nil
}

type NullableRelayedAuthenticationResponse struct {
	value *RelayedAuthenticationResponse
	isSet bool
}

func (v NullableRelayedAuthenticationResponse) Get() *RelayedAuthenticationResponse {
	return v.value
}

func (v *NullableRelayedAuthenticationResponse) Set(val *RelayedAuthenticationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRelayedAuthenticationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRelayedAuthenticationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelayedAuthenticationResponse(val *RelayedAuthenticationResponse) *NullableRelayedAuthenticationResponse {
	return &NullableRelayedAuthenticationResponse{value: val, isSet: true}
}

func (v NullableRelayedAuthenticationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelayedAuthenticationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
