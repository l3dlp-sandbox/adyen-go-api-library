/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the BrandVariantsRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BrandVariantsRestriction{}

// BrandVariantsRestriction struct for BrandVariantsRestriction
type BrandVariantsRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	// List of card brand variants.  Possible values:   - **mc**, **mccredit**, **mccommercialcredit_b2b**, **mcdebit**, **mcbusinessdebit**, **mcbusinessworlddebit**, **mcprepaid**, **mcmaestro**   - **visa**, **visacredit**, **visadebit**, **visaprepaid**.  You can specify a rule for a generic variant. For example, to create a rule for all Mastercard payment instruments, use **mc**. The rule is applied to all payment instruments under **mc**, such as **mcbusinessdebit** and **mcdebit**.  
	Value []string `json:"value,omitempty"`
}

// NewBrandVariantsRestriction instantiates a new BrandVariantsRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandVariantsRestriction(operation string) *BrandVariantsRestriction {
	this := BrandVariantsRestriction{}
	this.Operation = operation
	return &this
}

// NewBrandVariantsRestrictionWithDefaults instantiates a new BrandVariantsRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandVariantsRestrictionWithDefaults() *BrandVariantsRestriction {
	this := BrandVariantsRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *BrandVariantsRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *BrandVariantsRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *BrandVariantsRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BrandVariantsRestriction) GetValue() []string {
	if o == nil || common.IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandVariantsRestriction) GetValueOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BrandVariantsRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *BrandVariantsRestriction) SetValue(v []string) {
	o.Value = v
}

func (o BrandVariantsRestriction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandVariantsRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableBrandVariantsRestriction struct {
	value *BrandVariantsRestriction
	isSet bool
}

func (v NullableBrandVariantsRestriction) Get() *BrandVariantsRestriction {
	return v.value
}

func (v *NullableBrandVariantsRestriction) Set(val *BrandVariantsRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandVariantsRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandVariantsRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandVariantsRestriction(val *BrandVariantsRestriction) *NullableBrandVariantsRestriction {
	return &NullableBrandVariantsRestriction{value: val, isSet: true}
}

func (v NullableBrandVariantsRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandVariantsRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



