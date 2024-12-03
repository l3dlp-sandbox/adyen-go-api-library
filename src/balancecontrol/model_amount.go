/*
Adyen Balance Control API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balancecontrol

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Amount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Amount{}

// Amount struct for Amount
type Amount struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Currency string `json:"currency"`
	// The amount of the transaction, in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	Value int64 `json:"value"`
}

// NewAmount instantiates a new Amount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmount(currency string, value int64) *Amount {
	this := Amount{}
	this.Currency = currency
	this.Value = value
	return &this
}

// NewAmountWithDefaults instantiates a new Amount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountWithDefaults() *Amount {
	this := Amount{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Amount) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Amount) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Amount) SetCurrency(v string) {
	o.Currency = v
}

// GetValue returns the Value field value
func (o *Amount) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Amount) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Amount) SetValue(v int64) {
	o.Value = v
}

func (o Amount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Amount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableAmount struct {
	value *Amount
	isSet bool
}

func (v NullableAmount) Get() *Amount {
	return v.value
}

func (v *NullableAmount) Set(val *Amount) {
	v.value = val
	v.isSet = true
}

func (v NullableAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmount(val *Amount) *NullableAmount {
	return &NullableAmount{value: val, isSet: true}
}

func (v NullableAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
