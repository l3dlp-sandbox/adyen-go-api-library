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

// checks if the TokenRequestorsRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenRequestorsRestriction{}

// TokenRequestorsRestriction struct for TokenRequestorsRestriction
type TokenRequestorsRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string   `json:"operation"`
	Value     []string `json:"value,omitempty"`
}

// NewTokenRequestorsRestriction instantiates a new TokenRequestorsRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequestorsRestriction(operation string) *TokenRequestorsRestriction {
	this := TokenRequestorsRestriction{}
	this.Operation = operation
	return &this
}

// NewTokenRequestorsRestrictionWithDefaults instantiates a new TokenRequestorsRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestorsRestrictionWithDefaults() *TokenRequestorsRestriction {
	this := TokenRequestorsRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *TokenRequestorsRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *TokenRequestorsRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *TokenRequestorsRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TokenRequestorsRestriction) GetValue() []string {
	if o == nil || common.IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequestorsRestriction) GetValueOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TokenRequestorsRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *TokenRequestorsRestriction) SetValue(v []string) {
	o.Value = v
}

func (o TokenRequestorsRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRequestorsRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTokenRequestorsRestriction struct {
	value *TokenRequestorsRestriction
	isSet bool
}

func (v NullableTokenRequestorsRestriction) Get() *TokenRequestorsRestriction {
	return v.value
}

func (v *NullableTokenRequestorsRestriction) Set(val *TokenRequestorsRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequestorsRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequestorsRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequestorsRestriction(val *TokenRequestorsRestriction) *NullableTokenRequestorsRestriction {
	return &NullableTokenRequestorsRestriction{value: val, isSet: true}
}

func (v NullableTokenRequestorsRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequestorsRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
