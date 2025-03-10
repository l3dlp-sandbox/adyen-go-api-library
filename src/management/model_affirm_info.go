/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the AffirmInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AffirmInfo{}

// AffirmInfo struct for AffirmInfo
type AffirmInfo struct {
	// Merchant support email
	SupportEmail string `json:"supportEmail"`
}

// NewAffirmInfo instantiates a new AffirmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffirmInfo(supportEmail string) *AffirmInfo {
	this := AffirmInfo{}
	this.SupportEmail = supportEmail
	return &this
}

// NewAffirmInfoWithDefaults instantiates a new AffirmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffirmInfoWithDefaults() *AffirmInfo {
	this := AffirmInfo{}
	return &this
}

// GetSupportEmail returns the SupportEmail field value
func (o *AffirmInfo) GetSupportEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
func (o *AffirmInfo) GetSupportEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmail, true
}

// SetSupportEmail sets field value
func (o *AffirmInfo) SetSupportEmail(v string) {
	o.SupportEmail = v
}

func (o AffirmInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AffirmInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supportEmail"] = o.SupportEmail
	return toSerialize, nil
}

type NullableAffirmInfo struct {
	value *AffirmInfo
	isSet bool
}

func (v NullableAffirmInfo) Get() *AffirmInfo {
	return v.value
}

func (v *NullableAffirmInfo) Set(val *AffirmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAffirmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAffirmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffirmInfo(val *AffirmInfo) *NullableAffirmInfo {
	return &NullableAffirmInfo{value: val, isSet: true}
}

func (v NullableAffirmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffirmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
