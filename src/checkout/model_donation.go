/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the Donation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Donation{}

// Donation struct for Donation
type Donation struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes/).
	Currency string `json:"currency"`
	// The [type of donation](https://docs.adyen.com/online-payments/donations/#donation-types).  Possible values: * **roundup**: a donation where the original transaction amount is rounded up as a donation. * **fixedAmounts**: a donation where you show fixed donations amounts that the shopper can select from.
	DonationType string `json:"donationType"`
	// The maximum amount a transaction can be rounded up to make a donation. This field is only present when `donationType` is **roundup**.
	MaxRoundupAmount *int64 `json:"maxRoundupAmount,omitempty"`
	// The fixed donation amounts in [minor units](https://docs.adyen.com/development-resources/currency-codes//#minor-units). This field is only present when `donationType` is **fixedAmounts**.
	Values []int64 `json:"values,omitempty"`
}

// NewDonation instantiates a new Donation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDonation(currency string, donationType string) *Donation {
	this := Donation{}
	this.Currency = currency
	this.DonationType = donationType
	return &this
}

// NewDonationWithDefaults instantiates a new Donation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDonationWithDefaults() *Donation {
	this := Donation{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Donation) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Donation) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Donation) SetCurrency(v string) {
	o.Currency = v
}

// GetDonationType returns the DonationType field value
func (o *Donation) GetDonationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DonationType
}

// GetDonationTypeOk returns a tuple with the DonationType field value
// and a boolean to check if the value has been set.
func (o *Donation) GetDonationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DonationType, true
}

// SetDonationType sets field value
func (o *Donation) SetDonationType(v string) {
	o.DonationType = v
}

// GetMaxRoundupAmount returns the MaxRoundupAmount field value if set, zero value otherwise.
func (o *Donation) GetMaxRoundupAmount() int64 {
	if o == nil || common.IsNil(o.MaxRoundupAmount) {
		var ret int64
		return ret
	}
	return *o.MaxRoundupAmount
}

// GetMaxRoundupAmountOk returns a tuple with the MaxRoundupAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Donation) GetMaxRoundupAmountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxRoundupAmount) {
		return nil, false
	}
	return o.MaxRoundupAmount, true
}

// HasMaxRoundupAmount returns a boolean if a field has been set.
func (o *Donation) HasMaxRoundupAmount() bool {
	if o != nil && !common.IsNil(o.MaxRoundupAmount) {
		return true
	}

	return false
}

// SetMaxRoundupAmount gets a reference to the given int64 and assigns it to the MaxRoundupAmount field.
func (o *Donation) SetMaxRoundupAmount(v int64) {
	o.MaxRoundupAmount = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Donation) GetValues() []int64 {
	if o == nil || common.IsNil(o.Values) {
		var ret []int64
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Donation) GetValuesOk() ([]int64, bool) {
	if o == nil || common.IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Donation) HasValues() bool {
	if o != nil && !common.IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []int64 and assigns it to the Values field.
func (o *Donation) SetValues(v []int64) {
	o.Values = v
}

func (o Donation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Donation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["donationType"] = o.DonationType
	if !common.IsNil(o.MaxRoundupAmount) {
		toSerialize["maxRoundupAmount"] = o.MaxRoundupAmount
	}
	if !common.IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableDonation struct {
	value *Donation
	isSet bool
}

func (v NullableDonation) Get() *Donation {
	return v.value
}

func (v *NullableDonation) Set(val *Donation) {
	v.value = val
	v.isSet = true
}

func (v NullableDonation) IsSet() bool {
	return v.isSet
}

func (v *NullableDonation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonation(val *Donation) *NullableDonation {
	return &NullableDonation{value: val, isSet: true}
}

func (v NullableDonation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
