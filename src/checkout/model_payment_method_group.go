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

// checks if the PaymentMethodGroup type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodGroup{}

// PaymentMethodGroup struct for PaymentMethodGroup
type PaymentMethodGroup struct {
	// The name of the group.
	Name *string `json:"name,omitempty"`
	// Echo data to be used if the payment method is displayed as part of this group.
	PaymentMethodData *string `json:"paymentMethodData,omitempty"`
	// The unique code of the group.
	Type *string `json:"type,omitempty"`
}

// NewPaymentMethodGroup instantiates a new PaymentMethodGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodGroup() *PaymentMethodGroup {
	this := PaymentMethodGroup{}
	return &this
}

// NewPaymentMethodGroupWithDefaults instantiates a new PaymentMethodGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodGroupWithDefaults() *PaymentMethodGroup {
	this := PaymentMethodGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethodGroup) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodGroup) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethodGroup) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethodGroup) SetName(v string) {
	o.Name = &v
}

// GetPaymentMethodData returns the PaymentMethodData field value if set, zero value otherwise.
func (o *PaymentMethodGroup) GetPaymentMethodData() string {
	if o == nil || common.IsNil(o.PaymentMethodData) {
		var ret string
		return ret
	}
	return *o.PaymentMethodData
}

// GetPaymentMethodDataOk returns a tuple with the PaymentMethodData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodGroup) GetPaymentMethodDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodData) {
		return nil, false
	}
	return o.PaymentMethodData, true
}

// HasPaymentMethodData returns a boolean if a field has been set.
func (o *PaymentMethodGroup) HasPaymentMethodData() bool {
	if o != nil && !common.IsNil(o.PaymentMethodData) {
		return true
	}

	return false
}

// SetPaymentMethodData gets a reference to the given string and assigns it to the PaymentMethodData field.
func (o *PaymentMethodGroup) SetPaymentMethodData(v string) {
	o.PaymentMethodData = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodGroup) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodGroup) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodGroup) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodGroup) SetType(v string) {
	o.Type = &v
}

func (o PaymentMethodGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.PaymentMethodData) {
		toSerialize["paymentMethodData"] = o.PaymentMethodData
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePaymentMethodGroup struct {
	value *PaymentMethodGroup
	isSet bool
}

func (v NullablePaymentMethodGroup) Get() *PaymentMethodGroup {
	return v.value
}

func (v *NullablePaymentMethodGroup) Set(val *PaymentMethodGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodGroup(val *PaymentMethodGroup) *NullablePaymentMethodGroup {
	return &NullablePaymentMethodGroup{value: val, isSet: true}
}

func (v NullablePaymentMethodGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
