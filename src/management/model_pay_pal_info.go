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

// checks if the PayPalInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayPalInfo{}

// PayPalInfo struct for PayPalInfo
type PayPalInfo struct {
	// Indicates if direct (immediate) capture for PayPal is enabled. If set to **true**, this setting overrides the [capture](https://docs.adyen.com/online-payments/capture) settings of your merchant account. Default value: **true**.
	DirectCapture *bool `json:"directCapture,omitempty"`
	// PayPal Merchant ID. Character length and limitations: 13 single-byte alphanumeric characters.
	PayerId string `json:"payerId"`
	// Your business email address.
	Subject string `json:"subject"`
}

// NewPayPalInfo instantiates a new PayPalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayPalInfo(payerId string, subject string) *PayPalInfo {
	this := PayPalInfo{}
	this.PayerId = payerId
	this.Subject = subject
	return &this
}

// NewPayPalInfoWithDefaults instantiates a new PayPalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayPalInfoWithDefaults() *PayPalInfo {
	this := PayPalInfo{}
	return &this
}

// GetDirectCapture returns the DirectCapture field value if set, zero value otherwise.
func (o *PayPalInfo) GetDirectCapture() bool {
	if o == nil || common.IsNil(o.DirectCapture) {
		var ret bool
		return ret
	}
	return *o.DirectCapture
}

// GetDirectCaptureOk returns a tuple with the DirectCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalInfo) GetDirectCaptureOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DirectCapture) {
		return nil, false
	}
	return o.DirectCapture, true
}

// HasDirectCapture returns a boolean if a field has been set.
func (o *PayPalInfo) HasDirectCapture() bool {
	if o != nil && !common.IsNil(o.DirectCapture) {
		return true
	}

	return false
}

// SetDirectCapture gets a reference to the given bool and assigns it to the DirectCapture field.
func (o *PayPalInfo) SetDirectCapture(v bool) {
	o.DirectCapture = &v
}

// GetPayerId returns the PayerId field value
func (o *PayPalInfo) GetPayerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayerId
}

// GetPayerIdOk returns a tuple with the PayerId field value
// and a boolean to check if the value has been set.
func (o *PayPalInfo) GetPayerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerId, true
}

// SetPayerId sets field value
func (o *PayPalInfo) SetPayerId(v string) {
	o.PayerId = v
}

// GetSubject returns the Subject field value
func (o *PayPalInfo) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *PayPalInfo) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *PayPalInfo) SetSubject(v string) {
	o.Subject = v
}

func (o PayPalInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayPalInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DirectCapture) {
		toSerialize["directCapture"] = o.DirectCapture
	}
	toSerialize["payerId"] = o.PayerId
	toSerialize["subject"] = o.Subject
	return toSerialize, nil
}

type NullablePayPalInfo struct {
	value *PayPalInfo
	isSet bool
}

func (v NullablePayPalInfo) Get() *PayPalInfo {
	return v.value
}

func (v *NullablePayPalInfo) Set(val *PayPalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePayPalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePayPalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayPalInfo(val *PayPalInfo) *NullablePayPalInfo {
	return &NullablePayPalInfo{value: val, isSet: true}
}

func (v NullablePayPalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayPalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
