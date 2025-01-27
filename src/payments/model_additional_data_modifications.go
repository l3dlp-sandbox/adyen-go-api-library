/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the AdditionalDataModifications type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataModifications{}

// AdditionalDataModifications struct for AdditionalDataModifications
type AdditionalDataModifications struct {
	// This is the installment option selected by the shopper. It is required only if specified by the user.
	InstallmentPaymentDataSelectedInstallmentOption *string `json:"installmentPaymentData.selectedInstallmentOption,omitempty"`
}

// NewAdditionalDataModifications instantiates a new AdditionalDataModifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataModifications() *AdditionalDataModifications {
	this := AdditionalDataModifications{}
	return &this
}

// NewAdditionalDataModificationsWithDefaults instantiates a new AdditionalDataModifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataModificationsWithDefaults() *AdditionalDataModifications {
	this := AdditionalDataModifications{}
	return &this
}

// GetInstallmentPaymentDataSelectedInstallmentOption returns the InstallmentPaymentDataSelectedInstallmentOption field value if set, zero value otherwise.
func (o *AdditionalDataModifications) GetInstallmentPaymentDataSelectedInstallmentOption() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataSelectedInstallmentOption) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataSelectedInstallmentOption
}

// GetInstallmentPaymentDataSelectedInstallmentOptionOk returns a tuple with the InstallmentPaymentDataSelectedInstallmentOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataModifications) GetInstallmentPaymentDataSelectedInstallmentOptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataSelectedInstallmentOption) {
		return nil, false
	}
	return o.InstallmentPaymentDataSelectedInstallmentOption, true
}

// HasInstallmentPaymentDataSelectedInstallmentOption returns a boolean if a field has been set.
func (o *AdditionalDataModifications) HasInstallmentPaymentDataSelectedInstallmentOption() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataSelectedInstallmentOption) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataSelectedInstallmentOption gets a reference to the given string and assigns it to the InstallmentPaymentDataSelectedInstallmentOption field.
func (o *AdditionalDataModifications) SetInstallmentPaymentDataSelectedInstallmentOption(v string) {
	o.InstallmentPaymentDataSelectedInstallmentOption = &v
}

func (o AdditionalDataModifications) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataModifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.InstallmentPaymentDataSelectedInstallmentOption) {
		toSerialize["installmentPaymentData.selectedInstallmentOption"] = o.InstallmentPaymentDataSelectedInstallmentOption
	}
	return toSerialize, nil
}

type NullableAdditionalDataModifications struct {
	value *AdditionalDataModifications
	isSet bool
}

func (v NullableAdditionalDataModifications) Get() *AdditionalDataModifications {
	return v.value
}

func (v *NullableAdditionalDataModifications) Set(val *AdditionalDataModifications) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataModifications) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataModifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataModifications(val *AdditionalDataModifications) *NullableAdditionalDataModifications {
	return &NullableAdditionalDataModifications{value: val, isSet: true}
}

func (v NullableAdditionalDataModifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataModifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
