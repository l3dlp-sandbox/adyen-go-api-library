/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the ResponseAdditionalDataNetworkTokens type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalDataNetworkTokens{}

// ResponseAdditionalDataNetworkTokens struct for ResponseAdditionalDataNetworkTokens
type ResponseAdditionalDataNetworkTokens struct {
	// Indicates whether a network token is available for the specified card.
	NetworkTokenAvailable *string `json:"networkToken.available,omitempty"`
	// The Bank Identification Number of a tokenized card, which is the first six digits of a card number.
	NetworkTokenBin *string `json:"networkToken.bin,omitempty"`
	// The last four digits of a network token.
	NetworkTokenTokenSummary *string `json:"networkToken.tokenSummary,omitempty"`
}

// NewResponseAdditionalDataNetworkTokens instantiates a new ResponseAdditionalDataNetworkTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalDataNetworkTokens() *ResponseAdditionalDataNetworkTokens {
	this := ResponseAdditionalDataNetworkTokens{}
	return &this
}

// NewResponseAdditionalDataNetworkTokensWithDefaults instantiates a new ResponseAdditionalDataNetworkTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalDataNetworkTokensWithDefaults() *ResponseAdditionalDataNetworkTokens {
	this := ResponseAdditionalDataNetworkTokens{}
	return &this
}

// GetNetworkTokenAvailable returns the NetworkTokenAvailable field value if set, zero value otherwise.
func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenAvailable() string {
	if o == nil || common.IsNil(o.NetworkTokenAvailable) {
		var ret string
		return ret
	}
	return *o.NetworkTokenAvailable
}

// GetNetworkTokenAvailableOk returns a tuple with the NetworkTokenAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenAvailableOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkTokenAvailable) {
		return nil, false
	}
	return o.NetworkTokenAvailable, true
}

// HasNetworkTokenAvailable returns a boolean if a field has been set.
func (o *ResponseAdditionalDataNetworkTokens) HasNetworkTokenAvailable() bool {
	if o != nil && !common.IsNil(o.NetworkTokenAvailable) {
		return true
	}

	return false
}

// SetNetworkTokenAvailable gets a reference to the given string and assigns it to the NetworkTokenAvailable field.
func (o *ResponseAdditionalDataNetworkTokens) SetNetworkTokenAvailable(v string) {
	o.NetworkTokenAvailable = &v
}

// GetNetworkTokenBin returns the NetworkTokenBin field value if set, zero value otherwise.
func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenBin() string {
	if o == nil || common.IsNil(o.NetworkTokenBin) {
		var ret string
		return ret
	}
	return *o.NetworkTokenBin
}

// GetNetworkTokenBinOk returns a tuple with the NetworkTokenBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenBinOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkTokenBin) {
		return nil, false
	}
	return o.NetworkTokenBin, true
}

// HasNetworkTokenBin returns a boolean if a field has been set.
func (o *ResponseAdditionalDataNetworkTokens) HasNetworkTokenBin() bool {
	if o != nil && !common.IsNil(o.NetworkTokenBin) {
		return true
	}

	return false
}

// SetNetworkTokenBin gets a reference to the given string and assigns it to the NetworkTokenBin field.
func (o *ResponseAdditionalDataNetworkTokens) SetNetworkTokenBin(v string) {
	o.NetworkTokenBin = &v
}

// GetNetworkTokenTokenSummary returns the NetworkTokenTokenSummary field value if set, zero value otherwise.
func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenTokenSummary() string {
	if o == nil || common.IsNil(o.NetworkTokenTokenSummary) {
		var ret string
		return ret
	}
	return *o.NetworkTokenTokenSummary
}

// GetNetworkTokenTokenSummaryOk returns a tuple with the NetworkTokenTokenSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenTokenSummaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkTokenTokenSummary) {
		return nil, false
	}
	return o.NetworkTokenTokenSummary, true
}

// HasNetworkTokenTokenSummary returns a boolean if a field has been set.
func (o *ResponseAdditionalDataNetworkTokens) HasNetworkTokenTokenSummary() bool {
	if o != nil && !common.IsNil(o.NetworkTokenTokenSummary) {
		return true
	}

	return false
}

// SetNetworkTokenTokenSummary gets a reference to the given string and assigns it to the NetworkTokenTokenSummary field.
func (o *ResponseAdditionalDataNetworkTokens) SetNetworkTokenTokenSummary(v string) {
	o.NetworkTokenTokenSummary = &v
}

func (o ResponseAdditionalDataNetworkTokens) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalDataNetworkTokens) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NetworkTokenAvailable) {
		toSerialize["networkToken.available"] = o.NetworkTokenAvailable
	}
	if !common.IsNil(o.NetworkTokenBin) {
		toSerialize["networkToken.bin"] = o.NetworkTokenBin
	}
	if !common.IsNil(o.NetworkTokenTokenSummary) {
		toSerialize["networkToken.tokenSummary"] = o.NetworkTokenTokenSummary
	}
	return toSerialize, nil
}

type NullableResponseAdditionalDataNetworkTokens struct {
	value *ResponseAdditionalDataNetworkTokens
	isSet bool
}

func (v NullableResponseAdditionalDataNetworkTokens) Get() *ResponseAdditionalDataNetworkTokens {
	return v.value
}

func (v *NullableResponseAdditionalDataNetworkTokens) Set(val *ResponseAdditionalDataNetworkTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalDataNetworkTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalDataNetworkTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalDataNetworkTokens(val *ResponseAdditionalDataNetworkTokens) *NullableResponseAdditionalDataNetworkTokens {
	return &NullableResponseAdditionalDataNetworkTokens{value: val, isSet: true}
}

func (v NullableResponseAdditionalDataNetworkTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalDataNetworkTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
