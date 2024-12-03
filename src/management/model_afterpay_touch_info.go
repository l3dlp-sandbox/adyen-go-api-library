/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the AfterpayTouchInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AfterpayTouchInfo{}

// AfterpayTouchInfo struct for AfterpayTouchInfo
type AfterpayTouchInfo struct {
	// Support Url
	SupportUrl string `json:"supportUrl"`
}

// NewAfterpayTouchInfo instantiates a new AfterpayTouchInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfterpayTouchInfo(supportUrl string) *AfterpayTouchInfo {
	this := AfterpayTouchInfo{}
	this.SupportUrl = supportUrl
	return &this
}

// NewAfterpayTouchInfoWithDefaults instantiates a new AfterpayTouchInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfterpayTouchInfoWithDefaults() *AfterpayTouchInfo {
	this := AfterpayTouchInfo{}
	return &this
}

// GetSupportUrl returns the SupportUrl field value
func (o *AfterpayTouchInfo) GetSupportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportUrl
}

// GetSupportUrlOk returns a tuple with the SupportUrl field value
// and a boolean to check if the value has been set.
func (o *AfterpayTouchInfo) GetSupportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportUrl, true
}

// SetSupportUrl sets field value
func (o *AfterpayTouchInfo) SetSupportUrl(v string) {
	o.SupportUrl = v
}

func (o AfterpayTouchInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfterpayTouchInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supportUrl"] = o.SupportUrl
	return toSerialize, nil
}

type NullableAfterpayTouchInfo struct {
	value *AfterpayTouchInfo
	isSet bool
}

func (v NullableAfterpayTouchInfo) Get() *AfterpayTouchInfo {
	return v.value
}

func (v *NullableAfterpayTouchInfo) Set(val *AfterpayTouchInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAfterpayTouchInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAfterpayTouchInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfterpayTouchInfo(val *AfterpayTouchInfo) *NullableAfterpayTouchInfo {
	return &NullableAfterpayTouchInfo{value: val, isSet: true}
}

func (v NullableAfterpayTouchInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfterpayTouchInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



