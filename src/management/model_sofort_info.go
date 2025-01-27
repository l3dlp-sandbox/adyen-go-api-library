/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the SofortInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SofortInfo{}

// SofortInfo struct for SofortInfo
type SofortInfo struct {
	// Sofort currency code. For example, **EUR**.
	CurrencyCode string `json:"currencyCode"`
	// Sofort logo. Format: Base64-encoded string.
	Logo string `json:"logo"`
}

// NewSofortInfo instantiates a new SofortInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSofortInfo(currencyCode string, logo string) *SofortInfo {
	this := SofortInfo{}
	this.CurrencyCode = currencyCode
	this.Logo = logo
	return &this
}

// NewSofortInfoWithDefaults instantiates a new SofortInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSofortInfoWithDefaults() *SofortInfo {
	this := SofortInfo{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *SofortInfo) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *SofortInfo) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *SofortInfo) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetLogo returns the Logo field value
func (o *SofortInfo) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *SofortInfo) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *SofortInfo) SetLogo(v string) {
	o.Logo = v
}

func (o SofortInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SofortInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currencyCode"] = o.CurrencyCode
	toSerialize["logo"] = o.Logo
	return toSerialize, nil
}

type NullableSofortInfo struct {
	value *SofortInfo
	isSet bool
}

func (v NullableSofortInfo) Get() *SofortInfo {
	return v.value
}

func (v *NullableSofortInfo) Set(val *SofortInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSofortInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSofortInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSofortInfo(val *SofortInfo) *NullableSofortInfo {
	return &NullableSofortInfo{value: val, isSet: true}
}

func (v NullableSofortInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSofortInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
