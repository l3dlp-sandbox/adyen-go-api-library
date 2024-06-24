/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the ThresholdRepayment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThresholdRepayment{}

// ThresholdRepayment struct for ThresholdRepayment
type ThresholdRepayment struct {
	Amount Amount `json:"amount"`
}

// NewThresholdRepayment instantiates a new ThresholdRepayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdRepayment(amount Amount) *ThresholdRepayment {
	this := ThresholdRepayment{}
	this.Amount = amount
	return &this
}

// NewThresholdRepaymentWithDefaults instantiates a new ThresholdRepayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdRepaymentWithDefaults() *ThresholdRepayment {
	this := ThresholdRepayment{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ThresholdRepayment) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ThresholdRepayment) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ThresholdRepayment) SetAmount(v Amount) {
	o.Amount = v
}

func (o ThresholdRepayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdRepayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableThresholdRepayment struct {
	value *ThresholdRepayment
	isSet bool
}

func (v NullableThresholdRepayment) Get() *ThresholdRepayment {
	return v.value
}

func (v *NullableThresholdRepayment) Set(val *ThresholdRepayment) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdRepayment) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdRepayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdRepayment(val *ThresholdRepayment) *NullableThresholdRepayment {
	return &NullableThresholdRepayment{value: val, isSet: true}
}

func (v NullableThresholdRepayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdRepayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
