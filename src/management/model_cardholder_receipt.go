/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the CardholderReceipt type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardholderReceipt{}

// CardholderReceipt struct for CardholderReceipt
type CardholderReceipt struct {
	// A custom header to show on the shopper receipt for an authorised transaction. Allows one or two comma-separated header lines, and blank lines. For example, `header,header,filler`
	HeaderForAuthorizedReceipt *string `json:"headerForAuthorizedReceipt,omitempty"`
}

// NewCardholderReceipt instantiates a new CardholderReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardholderReceipt() *CardholderReceipt {
	this := CardholderReceipt{}
	return &this
}

// NewCardholderReceiptWithDefaults instantiates a new CardholderReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardholderReceiptWithDefaults() *CardholderReceipt {
	this := CardholderReceipt{}
	return &this
}

// GetHeaderForAuthorizedReceipt returns the HeaderForAuthorizedReceipt field value if set, zero value otherwise.
func (o *CardholderReceipt) GetHeaderForAuthorizedReceipt() string {
	if o == nil || common.IsNil(o.HeaderForAuthorizedReceipt) {
		var ret string
		return ret
	}
	return *o.HeaderForAuthorizedReceipt
}

// GetHeaderForAuthorizedReceiptOk returns a tuple with the HeaderForAuthorizedReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderReceipt) GetHeaderForAuthorizedReceiptOk() (*string, bool) {
	if o == nil || common.IsNil(o.HeaderForAuthorizedReceipt) {
		return nil, false
	}
	return o.HeaderForAuthorizedReceipt, true
}

// HasHeaderForAuthorizedReceipt returns a boolean if a field has been set.
func (o *CardholderReceipt) HasHeaderForAuthorizedReceipt() bool {
	if o != nil && !common.IsNil(o.HeaderForAuthorizedReceipt) {
		return true
	}

	return false
}

// SetHeaderForAuthorizedReceipt gets a reference to the given string and assigns it to the HeaderForAuthorizedReceipt field.
func (o *CardholderReceipt) SetHeaderForAuthorizedReceipt(v string) {
	o.HeaderForAuthorizedReceipt = &v
}

func (o CardholderReceipt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardholderReceipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.HeaderForAuthorizedReceipt) {
		toSerialize["headerForAuthorizedReceipt"] = o.HeaderForAuthorizedReceipt
	}
	return toSerialize, nil
}

type NullableCardholderReceipt struct {
	value *CardholderReceipt
	isSet bool
}

func (v NullableCardholderReceipt) Get() *CardholderReceipt {
	return v.value
}

func (v *NullableCardholderReceipt) Set(val *CardholderReceipt) {
	v.value = val
	v.isSet = true
}

func (v NullableCardholderReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableCardholderReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardholderReceipt(val *CardholderReceipt) *NullableCardholderReceipt {
	return &NullableCardholderReceipt{value: val, isSet: true}
}

func (v NullableCardholderReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardholderReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
