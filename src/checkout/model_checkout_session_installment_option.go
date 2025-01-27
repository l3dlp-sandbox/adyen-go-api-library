/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the CheckoutSessionInstallmentOption type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutSessionInstallmentOption{}

// CheckoutSessionInstallmentOption struct for CheckoutSessionInstallmentOption
type CheckoutSessionInstallmentOption struct {
	// Defines the type of installment plan. If not set, defaults to **regular**.  Possible values: * **regular** * **revolving**
	Plans []string `json:"plans,omitempty"`
	// Preselected number of installments offered for this payment method.
	PreselectedValue *int32 `json:"preselectedValue,omitempty"`
	// An array of the number of installments that the shopper can choose from. For example, **[2,3,5]**. This cannot be specified simultaneously with `maxValue`.
	Values []int32 `json:"values,omitempty"`
}

// NewCheckoutSessionInstallmentOption instantiates a new CheckoutSessionInstallmentOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionInstallmentOption() *CheckoutSessionInstallmentOption {
	this := CheckoutSessionInstallmentOption{}
	return &this
}

// NewCheckoutSessionInstallmentOptionWithDefaults instantiates a new CheckoutSessionInstallmentOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionInstallmentOptionWithDefaults() *CheckoutSessionInstallmentOption {
	this := CheckoutSessionInstallmentOption{}
	return &this
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *CheckoutSessionInstallmentOption) GetPlans() []string {
	if o == nil || common.IsNil(o.Plans) {
		var ret []string
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionInstallmentOption) GetPlansOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *CheckoutSessionInstallmentOption) HasPlans() bool {
	if o != nil && !common.IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []string and assigns it to the Plans field.
func (o *CheckoutSessionInstallmentOption) SetPlans(v []string) {
	o.Plans = v
}

// GetPreselectedValue returns the PreselectedValue field value if set, zero value otherwise.
func (o *CheckoutSessionInstallmentOption) GetPreselectedValue() int32 {
	if o == nil || common.IsNil(o.PreselectedValue) {
		var ret int32
		return ret
	}
	return *o.PreselectedValue
}

// GetPreselectedValueOk returns a tuple with the PreselectedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionInstallmentOption) GetPreselectedValueOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PreselectedValue) {
		return nil, false
	}
	return o.PreselectedValue, true
}

// HasPreselectedValue returns a boolean if a field has been set.
func (o *CheckoutSessionInstallmentOption) HasPreselectedValue() bool {
	if o != nil && !common.IsNil(o.PreselectedValue) {
		return true
	}

	return false
}

// SetPreselectedValue gets a reference to the given int32 and assigns it to the PreselectedValue field.
func (o *CheckoutSessionInstallmentOption) SetPreselectedValue(v int32) {
	o.PreselectedValue = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *CheckoutSessionInstallmentOption) GetValues() []int32 {
	if o == nil || common.IsNil(o.Values) {
		var ret []int32
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionInstallmentOption) GetValuesOk() ([]int32, bool) {
	if o == nil || common.IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *CheckoutSessionInstallmentOption) HasValues() bool {
	if o != nil && !common.IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []int32 and assigns it to the Values field.
func (o *CheckoutSessionInstallmentOption) SetValues(v []int32) {
	o.Values = v
}

func (o CheckoutSessionInstallmentOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSessionInstallmentOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	if !common.IsNil(o.PreselectedValue) {
		toSerialize["preselectedValue"] = o.PreselectedValue
	}
	if !common.IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableCheckoutSessionInstallmentOption struct {
	value *CheckoutSessionInstallmentOption
	isSet bool
}

func (v NullableCheckoutSessionInstallmentOption) Get() *CheckoutSessionInstallmentOption {
	return v.value
}

func (v *NullableCheckoutSessionInstallmentOption) Set(val *CheckoutSessionInstallmentOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionInstallmentOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionInstallmentOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionInstallmentOption(val *CheckoutSessionInstallmentOption) *NullableCheckoutSessionInstallmentOption {
	return &NullableCheckoutSessionInstallmentOption{value: val, isSet: true}
}

func (v NullableCheckoutSessionInstallmentOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionInstallmentOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
