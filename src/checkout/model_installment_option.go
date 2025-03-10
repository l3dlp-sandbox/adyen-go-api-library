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

// checks if the InstallmentOption type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InstallmentOption{}

// InstallmentOption struct for InstallmentOption
type InstallmentOption struct {
	// The maximum number of installments offered for this payment method.
	MaxValue *int32 `json:"maxValue,omitempty"`
	// Defines the type of installment plan. If not set, defaults to **regular**.  Possible values: * **regular** * **revolving**
	Plans []string `json:"plans,omitempty"`
	// Preselected number of installments offered for this payment method.
	PreselectedValue *int32 `json:"preselectedValue,omitempty"`
	// An array of the number of installments that the shopper can choose from. For example, **[2,3,5]**. This cannot be specified simultaneously with `maxValue`.
	Values []int32 `json:"values,omitempty"`
}

// NewInstallmentOption instantiates a new InstallmentOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallmentOption() *InstallmentOption {
	this := InstallmentOption{}
	return &this
}

// NewInstallmentOptionWithDefaults instantiates a new InstallmentOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallmentOptionWithDefaults() *InstallmentOption {
	this := InstallmentOption{}
	return &this
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *InstallmentOption) GetMaxValue() int32 {
	if o == nil || common.IsNil(o.MaxValue) {
		var ret int32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallmentOption) GetMaxValueOk() (*int32, bool) {
	if o == nil || common.IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *InstallmentOption) HasMaxValue() bool {
	if o != nil && !common.IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int32 and assigns it to the MaxValue field.
func (o *InstallmentOption) SetMaxValue(v int32) {
	o.MaxValue = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *InstallmentOption) GetPlans() []string {
	if o == nil || common.IsNil(o.Plans) {
		var ret []string
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallmentOption) GetPlansOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *InstallmentOption) HasPlans() bool {
	if o != nil && !common.IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []string and assigns it to the Plans field.
func (o *InstallmentOption) SetPlans(v []string) {
	o.Plans = v
}

// GetPreselectedValue returns the PreselectedValue field value if set, zero value otherwise.
func (o *InstallmentOption) GetPreselectedValue() int32 {
	if o == nil || common.IsNil(o.PreselectedValue) {
		var ret int32
		return ret
	}
	return *o.PreselectedValue
}

// GetPreselectedValueOk returns a tuple with the PreselectedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallmentOption) GetPreselectedValueOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PreselectedValue) {
		return nil, false
	}
	return o.PreselectedValue, true
}

// HasPreselectedValue returns a boolean if a field has been set.
func (o *InstallmentOption) HasPreselectedValue() bool {
	if o != nil && !common.IsNil(o.PreselectedValue) {
		return true
	}

	return false
}

// SetPreselectedValue gets a reference to the given int32 and assigns it to the PreselectedValue field.
func (o *InstallmentOption) SetPreselectedValue(v int32) {
	o.PreselectedValue = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *InstallmentOption) GetValues() []int32 {
	if o == nil || common.IsNil(o.Values) {
		var ret []int32
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallmentOption) GetValuesOk() ([]int32, bool) {
	if o == nil || common.IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InstallmentOption) HasValues() bool {
	if o != nil && !common.IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []int32 and assigns it to the Values field.
func (o *InstallmentOption) SetValues(v []int32) {
	o.Values = v
}

func (o InstallmentOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallmentOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MaxValue) {
		toSerialize["maxValue"] = o.MaxValue
	}
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

type NullableInstallmentOption struct {
	value *InstallmentOption
	isSet bool
}

func (v NullableInstallmentOption) Get() *InstallmentOption {
	return v.value
}

func (v *NullableInstallmentOption) Set(val *InstallmentOption) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallmentOption) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallmentOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallmentOption(val *InstallmentOption) *NullableInstallmentOption {
	return &NullableInstallmentOption{value: val, isSet: true}
}

func (v NullableInstallmentOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallmentOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
