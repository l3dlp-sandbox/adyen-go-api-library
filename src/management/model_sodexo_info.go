/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the SodexoInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SodexoInfo{}

// SodexoInfo struct for SodexoInfo
type SodexoInfo struct {
	// Sodexo merchantContactPhone
	MerchantContactPhone string `json:"merchantContactPhone"`
}

// NewSodexoInfo instantiates a new SodexoInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodexoInfo(merchantContactPhone string) *SodexoInfo {
	this := SodexoInfo{}
	this.MerchantContactPhone = merchantContactPhone
	return &this
}

// NewSodexoInfoWithDefaults instantiates a new SodexoInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodexoInfoWithDefaults() *SodexoInfo {
	this := SodexoInfo{}
	return &this
}

// GetMerchantContactPhone returns the MerchantContactPhone field value
func (o *SodexoInfo) GetMerchantContactPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantContactPhone
}

// GetMerchantContactPhoneOk returns a tuple with the MerchantContactPhone field value
// and a boolean to check if the value has been set.
func (o *SodexoInfo) GetMerchantContactPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantContactPhone, true
}

// SetMerchantContactPhone sets field value
func (o *SodexoInfo) SetMerchantContactPhone(v string) {
	o.MerchantContactPhone = v
}

func (o SodexoInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodexoInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantContactPhone"] = o.MerchantContactPhone
	return toSerialize, nil
}

type NullableSodexoInfo struct {
	value *SodexoInfo
	isSet bool
}

func (v NullableSodexoInfo) Get() *SodexoInfo {
	return v.value
}

func (v *NullableSodexoInfo) Set(val *SodexoInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSodexoInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSodexoInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodexoInfo(val *SodexoInfo) *NullableSodexoInfo {
	return &NullableSodexoInfo{value: val, isSet: true}
}

func (v NullableSodexoInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodexoInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
