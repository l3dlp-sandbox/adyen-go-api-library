/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the CheckoutNativeRedirectAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutNativeRedirectAction{}

// CheckoutNativeRedirectAction struct for CheckoutNativeRedirectAction
type CheckoutNativeRedirectAction struct {
	// When the redirect URL must be accessed via POST, use this data to post to the redirect URL.
	Data *map[string]string `json:"data,omitempty"`
	// Specifies the HTTP method, for example GET or POST.
	Method *string `json:"method,omitempty"`
	// Native SDK's redirect data containing the direct issuer link and state data that must be submitted to the /v1/nativeRedirect/redirectResult.
	NativeRedirectData *string `json:"nativeRedirectData,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// **nativeRedirect**
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutNativeRedirectAction instantiates a new CheckoutNativeRedirectAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutNativeRedirectAction(type_ string) *CheckoutNativeRedirectAction {
	this := CheckoutNativeRedirectAction{}
	this.Type = type_
	return &this
}

// NewCheckoutNativeRedirectActionWithDefaults instantiates a new CheckoutNativeRedirectAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutNativeRedirectActionWithDefaults() *CheckoutNativeRedirectAction {
	this := CheckoutNativeRedirectAction{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckoutNativeRedirectAction) GetData() map[string]string {
	if o == nil || common.IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNativeRedirectAction) GetDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckoutNativeRedirectAction) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *CheckoutNativeRedirectAction) SetData(v map[string]string) {
	o.Data = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CheckoutNativeRedirectAction) GetMethod() string {
	if o == nil || common.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNativeRedirectAction) GetMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CheckoutNativeRedirectAction) HasMethod() bool {
	if o != nil && !common.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CheckoutNativeRedirectAction) SetMethod(v string) {
	o.Method = &v
}

// GetNativeRedirectData returns the NativeRedirectData field value if set, zero value otherwise.
func (o *CheckoutNativeRedirectAction) GetNativeRedirectData() string {
	if o == nil || common.IsNil(o.NativeRedirectData) {
		var ret string
		return ret
	}
	return *o.NativeRedirectData
}

// GetNativeRedirectDataOk returns a tuple with the NativeRedirectData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNativeRedirectAction) GetNativeRedirectDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.NativeRedirectData) {
		return nil, false
	}
	return o.NativeRedirectData, true
}

// HasNativeRedirectData returns a boolean if a field has been set.
func (o *CheckoutNativeRedirectAction) HasNativeRedirectData() bool {
	if o != nil && !common.IsNil(o.NativeRedirectData) {
		return true
	}

	return false
}

// SetNativeRedirectData gets a reference to the given string and assigns it to the NativeRedirectData field.
func (o *CheckoutNativeRedirectAction) SetNativeRedirectData(v string) {
	o.NativeRedirectData = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutNativeRedirectAction) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNativeRedirectAction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutNativeRedirectAction) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutNativeRedirectAction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetType returns the Type field value
func (o *CheckoutNativeRedirectAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutNativeRedirectAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutNativeRedirectAction) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutNativeRedirectAction) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutNativeRedirectAction) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutNativeRedirectAction) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutNativeRedirectAction) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutNativeRedirectAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutNativeRedirectAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !common.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !common.IsNil(o.NativeRedirectData) {
		toSerialize["nativeRedirectData"] = o.NativeRedirectData
	}
	if !common.IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutNativeRedirectAction struct {
	value *CheckoutNativeRedirectAction
	isSet bool
}

func (v NullableCheckoutNativeRedirectAction) Get() *CheckoutNativeRedirectAction {
	return v.value
}

func (v *NullableCheckoutNativeRedirectAction) Set(val *CheckoutNativeRedirectAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutNativeRedirectAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutNativeRedirectAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutNativeRedirectAction(val *CheckoutNativeRedirectAction) *NullableCheckoutNativeRedirectAction {
	return &NullableCheckoutNativeRedirectAction{value: val, isSet: true}
}

func (v NullableCheckoutNativeRedirectAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutNativeRedirectAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutNativeRedirectAction) isValidType() bool {
	var allowedEnumValues = []string{"nativeRedirect"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
