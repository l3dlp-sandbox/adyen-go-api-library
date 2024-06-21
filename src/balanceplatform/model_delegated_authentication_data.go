/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the DelegatedAuthenticationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DelegatedAuthenticationData{}

// DelegatedAuthenticationData struct for DelegatedAuthenticationData
type DelegatedAuthenticationData struct {
	// A base64-encoded block with the data required to register the SCA device. You obtain this information by using our authentication SDK.
	SdkOutput string `json:"sdkOutput"`
}

// NewDelegatedAuthenticationData instantiates a new DelegatedAuthenticationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegatedAuthenticationData(sdkOutput string) *DelegatedAuthenticationData {
	this := DelegatedAuthenticationData{}
	this.SdkOutput = sdkOutput
	return &this
}

// NewDelegatedAuthenticationDataWithDefaults instantiates a new DelegatedAuthenticationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegatedAuthenticationDataWithDefaults() *DelegatedAuthenticationData {
	this := DelegatedAuthenticationData{}
	return &this
}

// GetSdkOutput returns the SdkOutput field value
func (o *DelegatedAuthenticationData) GetSdkOutput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkOutput
}

// GetSdkOutputOk returns a tuple with the SdkOutput field value
// and a boolean to check if the value has been set.
func (o *DelegatedAuthenticationData) GetSdkOutputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkOutput, true
}

// SetSdkOutput sets field value
func (o *DelegatedAuthenticationData) SetSdkOutput(v string) {
	o.SdkOutput = v
}

func (o DelegatedAuthenticationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegatedAuthenticationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sdkOutput"] = o.SdkOutput
	return toSerialize, nil
}

type NullableDelegatedAuthenticationData struct {
	value *DelegatedAuthenticationData
	isSet bool
}

func (v NullableDelegatedAuthenticationData) Get() *DelegatedAuthenticationData {
	return v.value
}

func (v *NullableDelegatedAuthenticationData) Set(val *DelegatedAuthenticationData) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegatedAuthenticationData) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegatedAuthenticationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegatedAuthenticationData(val *DelegatedAuthenticationData) *NullableDelegatedAuthenticationData {
	return &NullableDelegatedAuthenticationData{value: val, isSet: true}
}

func (v NullableDelegatedAuthenticationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegatedAuthenticationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}