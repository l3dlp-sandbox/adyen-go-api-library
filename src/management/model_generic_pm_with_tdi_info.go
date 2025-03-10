/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the GenericPmWithTdiInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GenericPmWithTdiInfo{}

// GenericPmWithTdiInfo struct for GenericPmWithTdiInfo
type GenericPmWithTdiInfo struct {
	TransactionDescription *TransactionDescriptionInfo `json:"transactionDescription,omitempty"`
}

// NewGenericPmWithTdiInfo instantiates a new GenericPmWithTdiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericPmWithTdiInfo() *GenericPmWithTdiInfo {
	this := GenericPmWithTdiInfo{}
	return &this
}

// NewGenericPmWithTdiInfoWithDefaults instantiates a new GenericPmWithTdiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericPmWithTdiInfoWithDefaults() *GenericPmWithTdiInfo {
	this := GenericPmWithTdiInfo{}
	return &this
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *GenericPmWithTdiInfo) GetTransactionDescription() TransactionDescriptionInfo {
	if o == nil || common.IsNil(o.TransactionDescription) {
		var ret TransactionDescriptionInfo
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericPmWithTdiInfo) GetTransactionDescriptionOk() (*TransactionDescriptionInfo, bool) {
	if o == nil || common.IsNil(o.TransactionDescription) {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *GenericPmWithTdiInfo) HasTransactionDescription() bool {
	if o != nil && !common.IsNil(o.TransactionDescription) {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given TransactionDescriptionInfo and assigns it to the TransactionDescription field.
func (o *GenericPmWithTdiInfo) SetTransactionDescription(v TransactionDescriptionInfo) {
	o.TransactionDescription = &v
}

func (o GenericPmWithTdiInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericPmWithTdiInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransactionDescription) {
		toSerialize["transactionDescription"] = o.TransactionDescription
	}
	return toSerialize, nil
}

type NullableGenericPmWithTdiInfo struct {
	value *GenericPmWithTdiInfo
	isSet bool
}

func (v NullableGenericPmWithTdiInfo) Get() *GenericPmWithTdiInfo {
	return v.value
}

func (v *NullableGenericPmWithTdiInfo) Set(val *GenericPmWithTdiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericPmWithTdiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericPmWithTdiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericPmWithTdiInfo(val *GenericPmWithTdiInfo) *NullableGenericPmWithTdiInfo {
	return &NullableGenericPmWithTdiInfo{value: val, isSet: true}
}

func (v NullableGenericPmWithTdiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericPmWithTdiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
