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

// checks if the SameAmountRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SameAmountRestriction{}

// SameAmountRestriction struct for SameAmountRestriction
type SameAmountRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	Value     *bool  `json:"value,omitempty"`
}

// NewSameAmountRestriction instantiates a new SameAmountRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSameAmountRestriction(operation string) *SameAmountRestriction {
	this := SameAmountRestriction{}
	this.Operation = operation
	return &this
}

// NewSameAmountRestrictionWithDefaults instantiates a new SameAmountRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSameAmountRestrictionWithDefaults() *SameAmountRestriction {
	this := SameAmountRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *SameAmountRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *SameAmountRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *SameAmountRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SameAmountRestriction) GetValue() bool {
	if o == nil || common.IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SameAmountRestriction) GetValueOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SameAmountRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *SameAmountRestriction) SetValue(v bool) {
	o.Value = &v
}

func (o SameAmountRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SameAmountRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSameAmountRestriction struct {
	value *SameAmountRestriction
	isSet bool
}

func (v NullableSameAmountRestriction) Get() *SameAmountRestriction {
	return v.value
}

func (v *NullableSameAmountRestriction) Set(val *SameAmountRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableSameAmountRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableSameAmountRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSameAmountRestriction(val *SameAmountRestriction) *NullableSameAmountRestriction {
	return &NullableSameAmountRestriction{value: val, isSet: true}
}

func (v NullableSameAmountRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSameAmountRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
