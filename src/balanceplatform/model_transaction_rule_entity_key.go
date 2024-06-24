/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the TransactionRuleEntityKey type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRuleEntityKey{}

// TransactionRuleEntityKey struct for TransactionRuleEntityKey
type TransactionRuleEntityKey struct {
	// The unique identifier of the resource.
	EntityReference *string `json:"entityReference,omitempty"`
	// The type of resource.  Possible values: **balancePlatform**, **paymentInstrumentGroup**, **accountHolder**, **balanceAccount**, or **paymentInstrument**.
	EntityType *string `json:"entityType,omitempty"`
}

// NewTransactionRuleEntityKey instantiates a new TransactionRuleEntityKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleEntityKey() *TransactionRuleEntityKey {
	this := TransactionRuleEntityKey{}
	return &this
}

// NewTransactionRuleEntityKeyWithDefaults instantiates a new TransactionRuleEntityKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleEntityKeyWithDefaults() *TransactionRuleEntityKey {
	this := TransactionRuleEntityKey{}
	return &this
}

// GetEntityReference returns the EntityReference field value if set, zero value otherwise.
func (o *TransactionRuleEntityKey) GetEntityReference() string {
	if o == nil || common.IsNil(o.EntityReference) {
		var ret string
		return ret
	}
	return *o.EntityReference
}

// GetEntityReferenceOk returns a tuple with the EntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleEntityKey) GetEntityReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntityReference) {
		return nil, false
	}
	return o.EntityReference, true
}

// HasEntityReference returns a boolean if a field has been set.
func (o *TransactionRuleEntityKey) HasEntityReference() bool {
	if o != nil && !common.IsNil(o.EntityReference) {
		return true
	}

	return false
}

// SetEntityReference gets a reference to the given string and assigns it to the EntityReference field.
func (o *TransactionRuleEntityKey) SetEntityReference(v string) {
	o.EntityReference = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *TransactionRuleEntityKey) GetEntityType() string {
	if o == nil || common.IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleEntityKey) GetEntityTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *TransactionRuleEntityKey) HasEntityType() bool {
	if o != nil && !common.IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *TransactionRuleEntityKey) SetEntityType(v string) {
	o.EntityType = &v
}

func (o TransactionRuleEntityKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleEntityKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EntityReference) {
		toSerialize["entityReference"] = o.EntityReference
	}
	if !common.IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	return toSerialize, nil
}

type NullableTransactionRuleEntityKey struct {
	value *TransactionRuleEntityKey
	isSet bool
}

func (v NullableTransactionRuleEntityKey) Get() *TransactionRuleEntityKey {
	return v.value
}

func (v *NullableTransactionRuleEntityKey) Set(val *TransactionRuleEntityKey) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleEntityKey) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleEntityKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleEntityKey(val *TransactionRuleEntityKey) *NullableTransactionRuleEntityKey {
	return &NullableTransactionRuleEntityKey{value: val, isSet: true}
}

func (v NullableTransactionRuleEntityKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleEntityKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
