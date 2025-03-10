/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the AmountNonZeroDecimalsRequirement type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AmountNonZeroDecimalsRequirement{}

// AmountNonZeroDecimalsRequirement struct for AmountNonZeroDecimalsRequirement
type AmountNonZeroDecimalsRequirement struct {
	// Specifies for which routes the amount in a transfer request must have no non-zero decimal places, so the transfer can only be processed if the amount consists of round numbers.
	Description *string `json:"description,omitempty"`
	// **amountNonZeroDecimalsRequirement**
	Type string `json:"type"`
}

// NewAmountNonZeroDecimalsRequirement instantiates a new AmountNonZeroDecimalsRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmountNonZeroDecimalsRequirement(type_ string) *AmountNonZeroDecimalsRequirement {
	this := AmountNonZeroDecimalsRequirement{}
	this.Type = type_
	return &this
}

// NewAmountNonZeroDecimalsRequirementWithDefaults instantiates a new AmountNonZeroDecimalsRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountNonZeroDecimalsRequirementWithDefaults() *AmountNonZeroDecimalsRequirement {
	this := AmountNonZeroDecimalsRequirement{}
	var type_ string = "amountNonZeroDecimalsRequirement"
	this.Type = type_
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmountNonZeroDecimalsRequirement) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountNonZeroDecimalsRequirement) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmountNonZeroDecimalsRequirement) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmountNonZeroDecimalsRequirement) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *AmountNonZeroDecimalsRequirement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AmountNonZeroDecimalsRequirement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AmountNonZeroDecimalsRequirement) SetType(v string) {
	o.Type = v
}

func (o AmountNonZeroDecimalsRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmountNonZeroDecimalsRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAmountNonZeroDecimalsRequirement struct {
	value *AmountNonZeroDecimalsRequirement
	isSet bool
}

func (v NullableAmountNonZeroDecimalsRequirement) Get() *AmountNonZeroDecimalsRequirement {
	return v.value
}

func (v *NullableAmountNonZeroDecimalsRequirement) Set(val *AmountNonZeroDecimalsRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableAmountNonZeroDecimalsRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableAmountNonZeroDecimalsRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmountNonZeroDecimalsRequirement(val *AmountNonZeroDecimalsRequirement) *NullableAmountNonZeroDecimalsRequirement {
	return &NullableAmountNonZeroDecimalsRequirement{value: val, isSet: true}
}

func (v NullableAmountNonZeroDecimalsRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmountNonZeroDecimalsRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AmountNonZeroDecimalsRequirement) isValidType() bool {
	var allowedEnumValues = []string{"amountNonZeroDecimalsRequirement"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
