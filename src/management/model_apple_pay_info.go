/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the ApplePayInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApplePayInfo{}

// ApplePayInfo struct for ApplePayInfo
type ApplePayInfo struct {
	// The list of merchant domains. Maximum: 99 domains per request.  For more information, see [Apple Pay documentation](https://docs.adyen.com/payment-methods/apple-pay/web-drop-in?tab=adyen-certificate-live_1#going-live).
	Domains []string `json:"domains,omitempty"`
}

// NewApplePayInfo instantiates a new ApplePayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplePayInfo() *ApplePayInfo {
	this := ApplePayInfo{}
	return &this
}

// NewApplePayInfoWithDefaults instantiates a new ApplePayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplePayInfoWithDefaults() *ApplePayInfo {
	this := ApplePayInfo{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *ApplePayInfo) GetDomains() []string {
	if o == nil || common.IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplePayInfo) GetDomainsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *ApplePayInfo) HasDomains() bool {
	if o != nil && !common.IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *ApplePayInfo) SetDomains(v []string) {
	o.Domains = v
}

func (o ApplePayInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplePayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	return toSerialize, nil
}

type NullableApplePayInfo struct {
	value *ApplePayInfo
	isSet bool
}

func (v NullableApplePayInfo) Get() *ApplePayInfo {
	return v.value
}

func (v *NullableApplePayInfo) Set(val *ApplePayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApplePayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApplePayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplePayInfo(val *ApplePayInfo) *NullableApplePayInfo {
	return &NullableApplePayInfo{value: val, isSet: true}
}

func (v NullableApplePayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplePayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
