/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the InternationalTransactionRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InternationalTransactionRestriction{}

// InternationalTransactionRestriction struct for InternationalTransactionRestriction
type InternationalTransactionRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	// Boolean indicating whether transaction is an international transaction.  Possible values:  - **true**: The transaction is an international transaction.  - **false**: The transaction is a domestic transaction.
	Value *bool `json:"value,omitempty"`
}

// NewInternationalTransactionRestriction instantiates a new InternationalTransactionRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternationalTransactionRestriction(operation string) *InternationalTransactionRestriction {
	this := InternationalTransactionRestriction{}
	this.Operation = operation
	return &this
}

// NewInternationalTransactionRestrictionWithDefaults instantiates a new InternationalTransactionRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternationalTransactionRestrictionWithDefaults() *InternationalTransactionRestriction {
	this := InternationalTransactionRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *InternationalTransactionRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *InternationalTransactionRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *InternationalTransactionRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InternationalTransactionRestriction) GetValue() bool {
	if o == nil || common.IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalTransactionRestriction) GetValueOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InternationalTransactionRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *InternationalTransactionRestriction) SetValue(v bool) {
	o.Value = &v
}

func (o InternationalTransactionRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternationalTransactionRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableInternationalTransactionRestriction struct {
	value *InternationalTransactionRestriction
	isSet bool
}

func (v NullableInternationalTransactionRestriction) Get() *InternationalTransactionRestriction {
	return v.value
}

func (v *NullableInternationalTransactionRestriction) Set(val *InternationalTransactionRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableInternationalTransactionRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableInternationalTransactionRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternationalTransactionRestriction(val *InternationalTransactionRestriction) *NullableInternationalTransactionRestriction {
	return &NullableInternationalTransactionRestriction{value: val, isSet: true}
}

func (v NullableInternationalTransactionRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternationalTransactionRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
