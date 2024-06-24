/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the AdditionalDataWallets type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataWallets{}

// AdditionalDataWallets struct for AdditionalDataWallets
type AdditionalDataWallets struct {
	// The Android Pay token retrieved from the SDK.
	AndroidpayToken *string `json:"androidpay.token,omitempty"`
	// The Mastercard Masterpass Transaction ID retrieved from the SDK.
	MasterpassTransactionId *string `json:"masterpass.transactionId,omitempty"`
	// The Apple Pay token retrieved from the SDK.
	PaymentToken *string `json:"payment.token,omitempty"`
	// The Google Pay token retrieved from the SDK.
	PaywithgoogleToken *string `json:"paywithgoogle.token,omitempty"`
	// The Samsung Pay token retrieved from the SDK.
	SamsungpayToken *string `json:"samsungpay.token,omitempty"`
	// The Visa Checkout Call ID retrieved from the SDK.
	VisacheckoutCallId *string `json:"visacheckout.callId,omitempty"`
}

// NewAdditionalDataWallets instantiates a new AdditionalDataWallets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataWallets() *AdditionalDataWallets {
	this := AdditionalDataWallets{}
	return &this
}

// NewAdditionalDataWalletsWithDefaults instantiates a new AdditionalDataWallets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataWalletsWithDefaults() *AdditionalDataWallets {
	this := AdditionalDataWallets{}
	return &this
}

// GetAndroidpayToken returns the AndroidpayToken field value if set, zero value otherwise.
func (o *AdditionalDataWallets) GetAndroidpayToken() string {
	if o == nil || common.IsNil(o.AndroidpayToken) {
		var ret string
		return ret
	}
	return *o.AndroidpayToken
}

// GetAndroidpayTokenOk returns a tuple with the AndroidpayToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataWallets) GetAndroidpayTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.AndroidpayToken) {
		return nil, false
	}
	return o.AndroidpayToken, true
}

// HasAndroidpayToken returns a boolean if a field has been set.
func (o *AdditionalDataWallets) HasAndroidpayToken() bool {
	if o != nil && !common.IsNil(o.AndroidpayToken) {
		return true
	}

	return false
}

// SetAndroidpayToken gets a reference to the given string and assigns it to the AndroidpayToken field.
func (o *AdditionalDataWallets) SetAndroidpayToken(v string) {
	o.AndroidpayToken = &v
}

// GetMasterpassTransactionId returns the MasterpassTransactionId field value if set, zero value otherwise.
func (o *AdditionalDataWallets) GetMasterpassTransactionId() string {
	if o == nil || common.IsNil(o.MasterpassTransactionId) {
		var ret string
		return ret
	}
	return *o.MasterpassTransactionId
}

// GetMasterpassTransactionIdOk returns a tuple with the MasterpassTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataWallets) GetMasterpassTransactionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MasterpassTransactionId) {
		return nil, false
	}
	return o.MasterpassTransactionId, true
}

// HasMasterpassTransactionId returns a boolean if a field has been set.
func (o *AdditionalDataWallets) HasMasterpassTransactionId() bool {
	if o != nil && !common.IsNil(o.MasterpassTransactionId) {
		return true
	}

	return false
}

// SetMasterpassTransactionId gets a reference to the given string and assigns it to the MasterpassTransactionId field.
func (o *AdditionalDataWallets) SetMasterpassTransactionId(v string) {
	o.MasterpassTransactionId = &v
}

// GetPaymentToken returns the PaymentToken field value if set, zero value otherwise.
func (o *AdditionalDataWallets) GetPaymentToken() string {
	if o == nil || common.IsNil(o.PaymentToken) {
		var ret string
		return ret
	}
	return *o.PaymentToken
}

// GetPaymentTokenOk returns a tuple with the PaymentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataWallets) GetPaymentTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentToken) {
		return nil, false
	}
	return o.PaymentToken, true
}

// HasPaymentToken returns a boolean if a field has been set.
func (o *AdditionalDataWallets) HasPaymentToken() bool {
	if o != nil && !common.IsNil(o.PaymentToken) {
		return true
	}

	return false
}

// SetPaymentToken gets a reference to the given string and assigns it to the PaymentToken field.
func (o *AdditionalDataWallets) SetPaymentToken(v string) {
	o.PaymentToken = &v
}

// GetPaywithgoogleToken returns the PaywithgoogleToken field value if set, zero value otherwise.
func (o *AdditionalDataWallets) GetPaywithgoogleToken() string {
	if o == nil || common.IsNil(o.PaywithgoogleToken) {
		var ret string
		return ret
	}
	return *o.PaywithgoogleToken
}

// GetPaywithgoogleTokenOk returns a tuple with the PaywithgoogleToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataWallets) GetPaywithgoogleTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaywithgoogleToken) {
		return nil, false
	}
	return o.PaywithgoogleToken, true
}

// HasPaywithgoogleToken returns a boolean if a field has been set.
func (o *AdditionalDataWallets) HasPaywithgoogleToken() bool {
	if o != nil && !common.IsNil(o.PaywithgoogleToken) {
		return true
	}

	return false
}

// SetPaywithgoogleToken gets a reference to the given string and assigns it to the PaywithgoogleToken field.
func (o *AdditionalDataWallets) SetPaywithgoogleToken(v string) {
	o.PaywithgoogleToken = &v
}

// GetSamsungpayToken returns the SamsungpayToken field value if set, zero value otherwise.
func (o *AdditionalDataWallets) GetSamsungpayToken() string {
	if o == nil || common.IsNil(o.SamsungpayToken) {
		var ret string
		return ret
	}
	return *o.SamsungpayToken
}

// GetSamsungpayTokenOk returns a tuple with the SamsungpayToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataWallets) GetSamsungpayTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.SamsungpayToken) {
		return nil, false
	}
	return o.SamsungpayToken, true
}

// HasSamsungpayToken returns a boolean if a field has been set.
func (o *AdditionalDataWallets) HasSamsungpayToken() bool {
	if o != nil && !common.IsNil(o.SamsungpayToken) {
		return true
	}

	return false
}

// SetSamsungpayToken gets a reference to the given string and assigns it to the SamsungpayToken field.
func (o *AdditionalDataWallets) SetSamsungpayToken(v string) {
	o.SamsungpayToken = &v
}

// GetVisacheckoutCallId returns the VisacheckoutCallId field value if set, zero value otherwise.
func (o *AdditionalDataWallets) GetVisacheckoutCallId() string {
	if o == nil || common.IsNil(o.VisacheckoutCallId) {
		var ret string
		return ret
	}
	return *o.VisacheckoutCallId
}

// GetVisacheckoutCallIdOk returns a tuple with the VisacheckoutCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataWallets) GetVisacheckoutCallIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.VisacheckoutCallId) {
		return nil, false
	}
	return o.VisacheckoutCallId, true
}

// HasVisacheckoutCallId returns a boolean if a field has been set.
func (o *AdditionalDataWallets) HasVisacheckoutCallId() bool {
	if o != nil && !common.IsNil(o.VisacheckoutCallId) {
		return true
	}

	return false
}

// SetVisacheckoutCallId gets a reference to the given string and assigns it to the VisacheckoutCallId field.
func (o *AdditionalDataWallets) SetVisacheckoutCallId(v string) {
	o.VisacheckoutCallId = &v
}

func (o AdditionalDataWallets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataWallets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AndroidpayToken) {
		toSerialize["androidpay.token"] = o.AndroidpayToken
	}
	if !common.IsNil(o.MasterpassTransactionId) {
		toSerialize["masterpass.transactionId"] = o.MasterpassTransactionId
	}
	if !common.IsNil(o.PaymentToken) {
		toSerialize["payment.token"] = o.PaymentToken
	}
	if !common.IsNil(o.PaywithgoogleToken) {
		toSerialize["paywithgoogle.token"] = o.PaywithgoogleToken
	}
	if !common.IsNil(o.SamsungpayToken) {
		toSerialize["samsungpay.token"] = o.SamsungpayToken
	}
	if !common.IsNil(o.VisacheckoutCallId) {
		toSerialize["visacheckout.callId"] = o.VisacheckoutCallId
	}
	return toSerialize, nil
}

type NullableAdditionalDataWallets struct {
	value *AdditionalDataWallets
	isSet bool
}

func (v NullableAdditionalDataWallets) Get() *AdditionalDataWallets {
	return v.value
}

func (v *NullableAdditionalDataWallets) Set(val *AdditionalDataWallets) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataWallets) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataWallets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataWallets(val *AdditionalDataWallets) *NullableAdditionalDataWallets {
	return &NullableAdditionalDataWallets{value: val, isSet: true}
}

func (v NullableAdditionalDataWallets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataWallets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
