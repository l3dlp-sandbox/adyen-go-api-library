/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the MatchingTransactionsRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MatchingTransactionsRestriction{}

// MatchingTransactionsRestriction struct for MatchingTransactionsRestriction
type MatchingTransactionsRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	// The number of transactions.
	Value *int32 `json:"value,omitempty"`
}

// NewMatchingTransactionsRestriction instantiates a new MatchingTransactionsRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchingTransactionsRestriction(operation string) *MatchingTransactionsRestriction {
	this := MatchingTransactionsRestriction{}
	this.Operation = operation
	return &this
}

// NewMatchingTransactionsRestrictionWithDefaults instantiates a new MatchingTransactionsRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchingTransactionsRestrictionWithDefaults() *MatchingTransactionsRestriction {
	this := MatchingTransactionsRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *MatchingTransactionsRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *MatchingTransactionsRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *MatchingTransactionsRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MatchingTransactionsRestriction) GetValue() int32 {
	if o == nil || common.IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchingTransactionsRestriction) GetValueOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MatchingTransactionsRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *MatchingTransactionsRestriction) SetValue(v int32) {
	o.Value = &v
}

func (o MatchingTransactionsRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchingTransactionsRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMatchingTransactionsRestriction struct {
	value *MatchingTransactionsRestriction
	isSet bool
}

func (v NullableMatchingTransactionsRestriction) Get() *MatchingTransactionsRestriction {
	return v.value
}

func (v *NullableMatchingTransactionsRestriction) Set(val *MatchingTransactionsRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchingTransactionsRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchingTransactionsRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchingTransactionsRestriction(val *MatchingTransactionsRestriction) *NullableMatchingTransactionsRestriction {
	return &NullableMatchingTransactionsRestriction{value: val, isSet: true}
}

func (v NullableMatchingTransactionsRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchingTransactionsRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
