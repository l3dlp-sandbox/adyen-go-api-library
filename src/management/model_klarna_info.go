/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the KlarnaInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlarnaInfo{}

// KlarnaInfo struct for KlarnaInfo
type KlarnaInfo struct {
	// Indicates the status of [Automatic capture](https://docs.adyen.com/online-payments/capture#automatic-capture). Default value: **false**.
	AutoCapture *bool `json:"autoCapture,omitempty"`
	// The email address for disputes.
	DisputeEmail string `json:"disputeEmail"`
	// The region of operation. For example, **NA**, **EU**, **CH**, **AU**.
	Region string `json:"region"`
	// The email address of merchant support.
	SupportEmail string `json:"supportEmail"`
}

// NewKlarnaInfo instantiates a new KlarnaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaInfo(disputeEmail string, region string, supportEmail string) *KlarnaInfo {
	this := KlarnaInfo{}
	this.DisputeEmail = disputeEmail
	this.Region = region
	this.SupportEmail = supportEmail
	return &this
}

// NewKlarnaInfoWithDefaults instantiates a new KlarnaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaInfoWithDefaults() *KlarnaInfo {
	this := KlarnaInfo{}
	return &this
}

// GetAutoCapture returns the AutoCapture field value if set, zero value otherwise.
func (o *KlarnaInfo) GetAutoCapture() bool {
	if o == nil || common.IsNil(o.AutoCapture) {
		var ret bool
		return ret
	}
	return *o.AutoCapture
}

// GetAutoCaptureOk returns a tuple with the AutoCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaInfo) GetAutoCaptureOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoCapture) {
		return nil, false
	}
	return o.AutoCapture, true
}

// HasAutoCapture returns a boolean if a field has been set.
func (o *KlarnaInfo) HasAutoCapture() bool {
	if o != nil && !common.IsNil(o.AutoCapture) {
		return true
	}

	return false
}

// SetAutoCapture gets a reference to the given bool and assigns it to the AutoCapture field.
func (o *KlarnaInfo) SetAutoCapture(v bool) {
	o.AutoCapture = &v
}

// GetDisputeEmail returns the DisputeEmail field value
func (o *KlarnaInfo) GetDisputeEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisputeEmail
}

// GetDisputeEmailOk returns a tuple with the DisputeEmail field value
// and a boolean to check if the value has been set.
func (o *KlarnaInfo) GetDisputeEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputeEmail, true
}

// SetDisputeEmail sets field value
func (o *KlarnaInfo) SetDisputeEmail(v string) {
	o.DisputeEmail = v
}

// GetRegion returns the Region field value
func (o *KlarnaInfo) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *KlarnaInfo) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *KlarnaInfo) SetRegion(v string) {
	o.Region = v
}

// GetSupportEmail returns the SupportEmail field value
func (o *KlarnaInfo) GetSupportEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
func (o *KlarnaInfo) GetSupportEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmail, true
}

// SetSupportEmail sets field value
func (o *KlarnaInfo) SetSupportEmail(v string) {
	o.SupportEmail = v
}

func (o KlarnaInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AutoCapture) {
		toSerialize["autoCapture"] = o.AutoCapture
	}
	toSerialize["disputeEmail"] = o.DisputeEmail
	toSerialize["region"] = o.Region
	toSerialize["supportEmail"] = o.SupportEmail
	return toSerialize, nil
}

type NullableKlarnaInfo struct {
	value *KlarnaInfo
	isSet bool
}

func (v NullableKlarnaInfo) Get() *KlarnaInfo {
	return v.value
}

func (v *NullableKlarnaInfo) Set(val *KlarnaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaInfo(val *KlarnaInfo) *NullableKlarnaInfo {
	return &NullableKlarnaInfo{value: val, isSet: true}
}

func (v NullableKlarnaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *KlarnaInfo) isValidRegion() bool {
	var allowedEnumValues = []string{"NA", "EU", "CH", "AU"}
	for _, allowed := range allowedEnumValues {
		if o.GetRegion() == allowed {
			return true
		}
	}
	return false
}
