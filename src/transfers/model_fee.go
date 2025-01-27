/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the Fee type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Fee{}

// Fee struct for Fee
type Fee struct {
	Amount Amount `json:"amount"`
}

// NewFee instantiates a new Fee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFee(amount Amount) *Fee {
	this := Fee{}
	this.Amount = amount
	return &this
}

// NewFeeWithDefaults instantiates a new Fee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeWithDefaults() *Fee {
	this := Fee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Fee) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Fee) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Fee) SetAmount(v Amount) {
	o.Amount = v
}

func (o Fee) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Fee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableFee struct {
	value *Fee
	isSet bool
}

func (v NullableFee) Get() *Fee {
	return v.value
}

func (v *NullableFee) Set(val *Fee) {
	v.value = val
	v.isSet = true
}

func (v NullableFee) IsSet() bool {
	return v.isSet
}

func (v *NullableFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFee(val *Fee) *NullableFee {
	return &NullableFee{value: val, isSet: true}
}

func (v NullableFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
