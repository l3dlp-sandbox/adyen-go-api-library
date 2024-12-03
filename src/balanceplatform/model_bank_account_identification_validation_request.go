/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the BankAccountIdentificationValidationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankAccountIdentificationValidationRequest{}

// BankAccountIdentificationValidationRequest struct for BankAccountIdentificationValidationRequest
type BankAccountIdentificationValidationRequest struct {
	AccountIdentification BankAccountIdentificationValidationRequestAccountIdentification `json:"accountIdentification"`
}

// NewBankAccountIdentificationValidationRequest instantiates a new BankAccountIdentificationValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountIdentificationValidationRequest(accountIdentification BankAccountIdentificationValidationRequestAccountIdentification) *BankAccountIdentificationValidationRequest {
	this := BankAccountIdentificationValidationRequest{}
	this.AccountIdentification = accountIdentification
	return &this
}

// NewBankAccountIdentificationValidationRequestWithDefaults instantiates a new BankAccountIdentificationValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountIdentificationValidationRequestWithDefaults() *BankAccountIdentificationValidationRequest {
	this := BankAccountIdentificationValidationRequest{}
	return &this
}

// GetAccountIdentification returns the AccountIdentification field value
func (o *BankAccountIdentificationValidationRequest) GetAccountIdentification() BankAccountIdentificationValidationRequestAccountIdentification {
	if o == nil {
		var ret BankAccountIdentificationValidationRequestAccountIdentification
		return ret
	}

	return o.AccountIdentification
}

// GetAccountIdentificationOk returns a tuple with the AccountIdentification field value
// and a boolean to check if the value has been set.
func (o *BankAccountIdentificationValidationRequest) GetAccountIdentificationOk() (*BankAccountIdentificationValidationRequestAccountIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentification, true
}

// SetAccountIdentification sets field value
func (o *BankAccountIdentificationValidationRequest) SetAccountIdentification(v BankAccountIdentificationValidationRequestAccountIdentification) {
	o.AccountIdentification = v
}

func (o BankAccountIdentificationValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountIdentificationValidationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountIdentification"] = o.AccountIdentification
	return toSerialize, nil
}

type NullableBankAccountIdentificationValidationRequest struct {
	value *BankAccountIdentificationValidationRequest
	isSet bool
}

func (v NullableBankAccountIdentificationValidationRequest) Get() *BankAccountIdentificationValidationRequest {
	return v.value
}

func (v *NullableBankAccountIdentificationValidationRequest) Set(val *BankAccountIdentificationValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountIdentificationValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountIdentificationValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountIdentificationValidationRequest(val *BankAccountIdentificationValidationRequest) *NullableBankAccountIdentificationValidationRequest {
	return &NullableBankAccountIdentificationValidationRequest{value: val, isSet: true}
}

func (v NullableBankAccountIdentificationValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountIdentificationValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



