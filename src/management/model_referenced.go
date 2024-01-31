/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the Referenced type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Referenced{}

// Referenced struct for Referenced
type Referenced struct {
	// Indicates whether referenced refunds are enabled on the standalone terminal.
	EnableStandaloneRefunds *bool `json:"enableStandaloneRefunds,omitempty"`
}

// NewReferenced instantiates a new Referenced object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenced() *Referenced {
	this := Referenced{}
	return &this
}

// NewReferencedWithDefaults instantiates a new Referenced object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencedWithDefaults() *Referenced {
	this := Referenced{}
	return &this
}

// GetEnableStandaloneRefunds returns the EnableStandaloneRefunds field value if set, zero value otherwise.
func (o *Referenced) GetEnableStandaloneRefunds() bool {
	if o == nil || common.IsNil(o.EnableStandaloneRefunds) {
		var ret bool
		return ret
	}
	return *o.EnableStandaloneRefunds
}

// GetEnableStandaloneRefundsOk returns a tuple with the EnableStandaloneRefunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referenced) GetEnableStandaloneRefundsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableStandaloneRefunds) {
		return nil, false
	}
	return o.EnableStandaloneRefunds, true
}

// HasEnableStandaloneRefunds returns a boolean if a field has been set.
func (o *Referenced) HasEnableStandaloneRefunds() bool {
	if o != nil && !common.IsNil(o.EnableStandaloneRefunds) {
		return true
	}

	return false
}

// SetEnableStandaloneRefunds gets a reference to the given bool and assigns it to the EnableStandaloneRefunds field.
func (o *Referenced) SetEnableStandaloneRefunds(v bool) {
	o.EnableStandaloneRefunds = &v
}

func (o Referenced) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Referenced) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EnableStandaloneRefunds) {
		toSerialize["enableStandaloneRefunds"] = o.EnableStandaloneRefunds
	}
	return toSerialize, nil
}

type NullableReferenced struct {
	value *Referenced
	isSet bool
}

func (v NullableReferenced) Get() *Referenced {
	return v.value
}

func (v *NullableReferenced) Set(val *Referenced) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenced) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenced) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenced(val *Referenced) *NullableReferenced {
	return &NullableReferenced{value: val, isSet: true}
}

func (v NullableReferenced) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenced) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}