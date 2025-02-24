/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the RepaymentTerm type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RepaymentTerm{}

// RepaymentTerm struct for RepaymentTerm
type RepaymentTerm struct {
	// The estimated term for repaying the grant, in days.
	EstimatedDays int32 `json:"estimatedDays"`
	// The maximum term for repaying the grant, in days. Only applies when `contractType` is **loan**.
	MaximumDays *int32 `json:"maximumDays,omitempty"`
}

// NewRepaymentTerm instantiates a new RepaymentTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepaymentTerm(estimatedDays int32) *RepaymentTerm {
	this := RepaymentTerm{}
	this.EstimatedDays = estimatedDays
	return &this
}

// NewRepaymentTermWithDefaults instantiates a new RepaymentTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepaymentTermWithDefaults() *RepaymentTerm {
	this := RepaymentTerm{}
	return &this
}

// GetEstimatedDays returns the EstimatedDays field value
func (o *RepaymentTerm) GetEstimatedDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedDays
}

// GetEstimatedDaysOk returns a tuple with the EstimatedDays field value
// and a boolean to check if the value has been set.
func (o *RepaymentTerm) GetEstimatedDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedDays, true
}

// SetEstimatedDays sets field value
func (o *RepaymentTerm) SetEstimatedDays(v int32) {
	o.EstimatedDays = v
}

// GetMaximumDays returns the MaximumDays field value if set, zero value otherwise.
func (o *RepaymentTerm) GetMaximumDays() int32 {
	if o == nil || common.IsNil(o.MaximumDays) {
		var ret int32
		return ret
	}
	return *o.MaximumDays
}

// GetMaximumDaysOk returns a tuple with the MaximumDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepaymentTerm) GetMaximumDaysOk() (*int32, bool) {
	if o == nil || common.IsNil(o.MaximumDays) {
		return nil, false
	}
	return o.MaximumDays, true
}

// HasMaximumDays returns a boolean if a field has been set.
func (o *RepaymentTerm) HasMaximumDays() bool {
	if o != nil && !common.IsNil(o.MaximumDays) {
		return true
	}

	return false
}

// SetMaximumDays gets a reference to the given int32 and assigns it to the MaximumDays field.
func (o *RepaymentTerm) SetMaximumDays(v int32) {
	o.MaximumDays = &v
}

func (o RepaymentTerm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepaymentTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["estimatedDays"] = o.EstimatedDays
	if !common.IsNil(o.MaximumDays) {
		toSerialize["maximumDays"] = o.MaximumDays
	}
	return toSerialize, nil
}

type NullableRepaymentTerm struct {
	value *RepaymentTerm
	isSet bool
}

func (v NullableRepaymentTerm) Get() *RepaymentTerm {
	return v.value
}

func (v *NullableRepaymentTerm) Set(val *RepaymentTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableRepaymentTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableRepaymentTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepaymentTerm(val *RepaymentTerm) *NullableRepaymentTerm {
	return &NullableRepaymentTerm{value: val, isSet: true}
}

func (v NullableRepaymentTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepaymentTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
