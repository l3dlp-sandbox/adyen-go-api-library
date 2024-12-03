/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TransferInstrumentInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferInstrumentInfo{}

// TransferInstrumentInfo struct for TransferInstrumentInfo
type TransferInstrumentInfo struct {
	BankAccount BankAccountInfo `json:"bankAccount"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) that owns the transfer instrument.
	LegalEntityId string `json:"legalEntityId"`
	// The type of transfer instrument.  Possible value: **bankAccount**.
	Type string `json:"type"`
}

// NewTransferInstrumentInfo instantiates a new TransferInstrumentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferInstrumentInfo(bankAccount BankAccountInfo, legalEntityId string, type_ string) *TransferInstrumentInfo {
	this := TransferInstrumentInfo{}
	this.BankAccount = bankAccount
	this.LegalEntityId = legalEntityId
	this.Type = type_
	return &this
}

// NewTransferInstrumentInfoWithDefaults instantiates a new TransferInstrumentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferInstrumentInfoWithDefaults() *TransferInstrumentInfo {
	this := TransferInstrumentInfo{}
	return &this
}

// GetBankAccount returns the BankAccount field value
func (o *TransferInstrumentInfo) GetBankAccount() BankAccountInfo {
	if o == nil {
		var ret BankAccountInfo
		return ret
	}

	return o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
func (o *TransferInstrumentInfo) GetBankAccountOk() (*BankAccountInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccount, true
}

// SetBankAccount sets field value
func (o *TransferInstrumentInfo) SetBankAccount(v BankAccountInfo) {
	o.BankAccount = v
}

// GetLegalEntityId returns the LegalEntityId field value
func (o *TransferInstrumentInfo) GetLegalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value
// and a boolean to check if the value has been set.
func (o *TransferInstrumentInfo) GetLegalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalEntityId, true
}

// SetLegalEntityId sets field value
func (o *TransferInstrumentInfo) SetLegalEntityId(v string) {
	o.LegalEntityId = v
}

// GetType returns the Type field value
func (o *TransferInstrumentInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferInstrumentInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferInstrumentInfo) SetType(v string) {
	o.Type = v
}

func (o TransferInstrumentInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferInstrumentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bankAccount"] = o.BankAccount
	toSerialize["legalEntityId"] = o.LegalEntityId
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransferInstrumentInfo struct {
	value *TransferInstrumentInfo
	isSet bool
}

func (v NullableTransferInstrumentInfo) Get() *TransferInstrumentInfo {
	return v.value
}

func (v *NullableTransferInstrumentInfo) Set(val *TransferInstrumentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferInstrumentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferInstrumentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferInstrumentInfo(val *TransferInstrumentInfo) *NullableTransferInstrumentInfo {
	return &NullableTransferInstrumentInfo{value: val, isSet: true}
}

func (v NullableTransferInstrumentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferInstrumentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *TransferInstrumentInfo) isValidType() bool {
    var allowedEnumValues = []string{ "bankAccount", "recurringDetail" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

