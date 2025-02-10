/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the WeChatPayMiniProgramDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WeChatPayMiniProgramDetails{}

// WeChatPayMiniProgramDetails struct for WeChatPayMiniProgramDetails
type WeChatPayMiniProgramDetails struct {
	AppId *string `json:"appId,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	Openid            *string `json:"openid,omitempty"`
	// **wechatpayMiniProgram**
	Type *string `json:"type,omitempty"`
}

// NewWeChatPayMiniProgramDetails instantiates a new WeChatPayMiniProgramDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeChatPayMiniProgramDetails() *WeChatPayMiniProgramDetails {
	this := WeChatPayMiniProgramDetails{}
	var type_ string = "wechatpayMiniProgram"
	this.Type = &type_
	return &this
}

// NewWeChatPayMiniProgramDetailsWithDefaults instantiates a new WeChatPayMiniProgramDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeChatPayMiniProgramDetailsWithDefaults() *WeChatPayMiniProgramDetails {
	this := WeChatPayMiniProgramDetails{}
	var type_ string = "wechatpayMiniProgram"
	this.Type = &type_
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *WeChatPayMiniProgramDetails) GetAppId() string {
	if o == nil || common.IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeChatPayMiniProgramDetails) GetAppIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *WeChatPayMiniProgramDetails) HasAppId() bool {
	if o != nil && !common.IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *WeChatPayMiniProgramDetails) SetAppId(v string) {
	o.AppId = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *WeChatPayMiniProgramDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeChatPayMiniProgramDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *WeChatPayMiniProgramDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *WeChatPayMiniProgramDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetOpenid returns the Openid field value if set, zero value otherwise.
func (o *WeChatPayMiniProgramDetails) GetOpenid() string {
	if o == nil || common.IsNil(o.Openid) {
		var ret string
		return ret
	}
	return *o.Openid
}

// GetOpenidOk returns a tuple with the Openid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeChatPayMiniProgramDetails) GetOpenidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Openid) {
		return nil, false
	}
	return o.Openid, true
}

// HasOpenid returns a boolean if a field has been set.
func (o *WeChatPayMiniProgramDetails) HasOpenid() bool {
	if o != nil && !common.IsNil(o.Openid) {
		return true
	}

	return false
}

// SetOpenid gets a reference to the given string and assigns it to the Openid field.
func (o *WeChatPayMiniProgramDetails) SetOpenid(v string) {
	o.Openid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WeChatPayMiniProgramDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeChatPayMiniProgramDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WeChatPayMiniProgramDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WeChatPayMiniProgramDetails) SetType(v string) {
	o.Type = &v
}

func (o WeChatPayMiniProgramDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeChatPayMiniProgramDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.Openid) {
		toSerialize["openid"] = o.Openid
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableWeChatPayMiniProgramDetails struct {
	value *WeChatPayMiniProgramDetails
	isSet bool
}

func (v NullableWeChatPayMiniProgramDetails) Get() *WeChatPayMiniProgramDetails {
	return v.value
}

func (v *NullableWeChatPayMiniProgramDetails) Set(val *WeChatPayMiniProgramDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableWeChatPayMiniProgramDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableWeChatPayMiniProgramDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeChatPayMiniProgramDetails(val *WeChatPayMiniProgramDetails) *NullableWeChatPayMiniProgramDetails {
	return &NullableWeChatPayMiniProgramDetails{value: val, isSet: true}
}

func (v NullableWeChatPayMiniProgramDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeChatPayMiniProgramDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *WeChatPayMiniProgramDetails) isValidType() bool {
	var allowedEnumValues = []string{"wechatpayMiniProgram"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
