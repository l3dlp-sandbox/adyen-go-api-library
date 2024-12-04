/*
Transaction webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactionwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TransferNotificationValidationFact type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferNotificationValidationFact{}

// TransferNotificationValidationFact struct for TransferNotificationValidationFact
type TransferNotificationValidationFact struct {
	// The evaluation result of the validation fact.
	Result *string `json:"result,omitempty"`
	// The type of the validation fact.
	Type *string `json:"type,omitempty"`
}

// NewTransferNotificationValidationFact instantiates a new TransferNotificationValidationFact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferNotificationValidationFact() *TransferNotificationValidationFact {
	this := TransferNotificationValidationFact{}
	return &this
}

// NewTransferNotificationValidationFactWithDefaults instantiates a new TransferNotificationValidationFact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferNotificationValidationFactWithDefaults() *TransferNotificationValidationFact {
	this := TransferNotificationValidationFact{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TransferNotificationValidationFact) GetResult() string {
	if o == nil || common.IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationValidationFact) GetResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TransferNotificationValidationFact) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *TransferNotificationValidationFact) SetResult(v string) {
	o.Result = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferNotificationValidationFact) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationValidationFact) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferNotificationValidationFact) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransferNotificationValidationFact) SetType(v string) {
	o.Type = &v
}

func (o TransferNotificationValidationFact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferNotificationValidationFact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTransferNotificationValidationFact struct {
	value *TransferNotificationValidationFact
	isSet bool
}

func (v NullableTransferNotificationValidationFact) Get() *TransferNotificationValidationFact {
	return v.value
}

func (v *NullableTransferNotificationValidationFact) Set(val *TransferNotificationValidationFact) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferNotificationValidationFact) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferNotificationValidationFact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferNotificationValidationFact(val *TransferNotificationValidationFact) *NullableTransferNotificationValidationFact {
	return &NullableTransferNotificationValidationFact{value: val, isSet: true}
}

func (v NullableTransferNotificationValidationFact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferNotificationValidationFact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}