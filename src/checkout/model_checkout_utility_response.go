/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the CheckoutUtilityResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutUtilityResponse{}

// CheckoutUtilityResponse struct for CheckoutUtilityResponse
type CheckoutUtilityResponse struct {
	// The list of origin keys for all requested domains. For each list item, the key is the domain and the value is the origin key.
	OriginKeys *map[string]string `json:"originKeys,omitempty"`
}

// NewCheckoutUtilityResponse instantiates a new CheckoutUtilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutUtilityResponse() *CheckoutUtilityResponse {
	this := CheckoutUtilityResponse{}
	return &this
}

// NewCheckoutUtilityResponseWithDefaults instantiates a new CheckoutUtilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutUtilityResponseWithDefaults() *CheckoutUtilityResponse {
	this := CheckoutUtilityResponse{}
	return &this
}

// GetOriginKeys returns the OriginKeys field value if set, zero value otherwise.
func (o *CheckoutUtilityResponse) GetOriginKeys() map[string]string {
	if o == nil || common.IsNil(o.OriginKeys) {
		var ret map[string]string
		return ret
	}
	return *o.OriginKeys
}

// GetOriginKeysOk returns a tuple with the OriginKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutUtilityResponse) GetOriginKeysOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.OriginKeys) {
		return nil, false
	}
	return o.OriginKeys, true
}

// HasOriginKeys returns a boolean if a field has been set.
func (o *CheckoutUtilityResponse) HasOriginKeys() bool {
	if o != nil && !common.IsNil(o.OriginKeys) {
		return true
	}

	return false
}

// SetOriginKeys gets a reference to the given map[string]string and assigns it to the OriginKeys field.
func (o *CheckoutUtilityResponse) SetOriginKeys(v map[string]string) {
	o.OriginKeys = &v
}

func (o CheckoutUtilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutUtilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OriginKeys) {
		toSerialize["originKeys"] = o.OriginKeys
	}
	return toSerialize, nil
}

type NullableCheckoutUtilityResponse struct {
	value *CheckoutUtilityResponse
	isSet bool
}

func (v NullableCheckoutUtilityResponse) Get() *CheckoutUtilityResponse {
	return v.value
}

func (v *NullableCheckoutUtilityResponse) Set(val *CheckoutUtilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutUtilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutUtilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutUtilityResponse(val *CheckoutUtilityResponse) *NullableCheckoutUtilityResponse {
	return &NullableCheckoutUtilityResponse{value: val, isSet: true}
}

func (v NullableCheckoutUtilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutUtilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
