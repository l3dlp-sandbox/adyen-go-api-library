/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the WeChatPayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WeChatPayDetails{}

// WeChatPayDetails struct for WeChatPayDetails
type WeChatPayDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// **wechatpay**
	Type *string `json:"type,omitempty"`
}

// NewWeChatPayDetails instantiates a new WeChatPayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeChatPayDetails() *WeChatPayDetails {
	this := WeChatPayDetails{}
	var type_ string = "wechatpay"
	this.Type = &type_
	return &this
}

// NewWeChatPayDetailsWithDefaults instantiates a new WeChatPayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeChatPayDetailsWithDefaults() *WeChatPayDetails {
	this := WeChatPayDetails{}
	var type_ string = "wechatpay"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *WeChatPayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeChatPayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *WeChatPayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *WeChatPayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WeChatPayDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeChatPayDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WeChatPayDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WeChatPayDetails) SetType(v string) {
	o.Type = &v
}

func (o WeChatPayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeChatPayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableWeChatPayDetails struct {
	value *WeChatPayDetails
	isSet bool
}

func (v NullableWeChatPayDetails) Get() *WeChatPayDetails {
	return v.value
}

func (v *NullableWeChatPayDetails) Set(val *WeChatPayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableWeChatPayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableWeChatPayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeChatPayDetails(val *WeChatPayDetails) *NullableWeChatPayDetails {
	return &NullableWeChatPayDetails{value: val, isSet: true}
}

func (v NullableWeChatPayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeChatPayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *WeChatPayDetails) isValidType() bool {
	var allowedEnumValues = []string{"wechatpay", "wechatpay_pos"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
