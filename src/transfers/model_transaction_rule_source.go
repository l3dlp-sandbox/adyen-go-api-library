/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the TransactionRuleSource type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRuleSource{}

// TransactionRuleSource struct for TransactionRuleSource
type TransactionRuleSource struct {
	// ID of the resource, when applicable.
	Id *string `json:"id,omitempty"`
	// Indicates the type of resource for which the transaction rule is defined.  Possible values:   * **PaymentInstrumentGroup**  * **PaymentInstrument**  * **BalancePlatform**  * **EntityUsageConfiguration**  * **PlatformRule**: The transaction rule is a platform-wide rule imposed by Adyen.
	Type *string `json:"type,omitempty"`
}

// NewTransactionRuleSource instantiates a new TransactionRuleSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleSource() *TransactionRuleSource {
	this := TransactionRuleSource{}
	return &this
}

// NewTransactionRuleSourceWithDefaults instantiates a new TransactionRuleSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleSourceWithDefaults() *TransactionRuleSource {
	this := TransactionRuleSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionRuleSource) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleSource) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionRuleSource) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionRuleSource) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransactionRuleSource) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleSource) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransactionRuleSource) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransactionRuleSource) SetType(v string) {
	o.Type = &v
}

func (o TransactionRuleSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTransactionRuleSource struct {
	value *TransactionRuleSource
	isSet bool
}

func (v NullableTransactionRuleSource) Get() *TransactionRuleSource {
	return v.value
}

func (v *NullableTransactionRuleSource) Set(val *TransactionRuleSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleSource(val *TransactionRuleSource) *NullableTransactionRuleSource {
	return &NullableTransactionRuleSource{value: val, isSet: true}
}

func (v NullableTransactionRuleSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
