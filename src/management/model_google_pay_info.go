/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the GooglePayInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GooglePayInfo{}

// GooglePayInfo struct for GooglePayInfo
type GooglePayInfo struct {
	// Google Pay [Merchant ID](https://support.google.com/paymentscenter/answer/7163092?hl=en). Character length and limitations: 16 alphanumeric characters or 20 numeric characters.
	MerchantId string `json:"merchantId"`
	// Indicates whether the Google Pay Merchant ID is used for several merchant accounts. Default value: **false**.
	ReuseMerchantId *bool `json:"reuseMerchantId,omitempty"`
}

// NewGooglePayInfo instantiates a new GooglePayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGooglePayInfo(merchantId string) *GooglePayInfo {
	this := GooglePayInfo{}
	this.MerchantId = merchantId
	return &this
}

// NewGooglePayInfoWithDefaults instantiates a new GooglePayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGooglePayInfoWithDefaults() *GooglePayInfo {
	this := GooglePayInfo{}
	return &this
}

// GetMerchantId returns the MerchantId field value
func (o *GooglePayInfo) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *GooglePayInfo) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *GooglePayInfo) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetReuseMerchantId returns the ReuseMerchantId field value if set, zero value otherwise.
func (o *GooglePayInfo) GetReuseMerchantId() bool {
	if o == nil || common.IsNil(o.ReuseMerchantId) {
		var ret bool
		return ret
	}
	return *o.ReuseMerchantId
}

// GetReuseMerchantIdOk returns a tuple with the ReuseMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglePayInfo) GetReuseMerchantIdOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReuseMerchantId) {
		return nil, false
	}
	return o.ReuseMerchantId, true
}

// HasReuseMerchantId returns a boolean if a field has been set.
func (o *GooglePayInfo) HasReuseMerchantId() bool {
	if o != nil && !common.IsNil(o.ReuseMerchantId) {
		return true
	}

	return false
}

// SetReuseMerchantId gets a reference to the given bool and assigns it to the ReuseMerchantId field.
func (o *GooglePayInfo) SetReuseMerchantId(v bool) {
	o.ReuseMerchantId = &v
}

func (o GooglePayInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GooglePayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantId"] = o.MerchantId
	if !common.IsNil(o.ReuseMerchantId) {
		toSerialize["reuseMerchantId"] = o.ReuseMerchantId
	}
	return toSerialize, nil
}

type NullableGooglePayInfo struct {
	value *GooglePayInfo
	isSet bool
}

func (v NullableGooglePayInfo) Get() *GooglePayInfo {
	return v.value
}

func (v *NullableGooglePayInfo) Set(val *GooglePayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGooglePayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGooglePayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGooglePayInfo(val *GooglePayInfo) *NullableGooglePayInfo {
	return &NullableGooglePayInfo{value: val, isSet: true}
}

func (v NullableGooglePayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGooglePayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
