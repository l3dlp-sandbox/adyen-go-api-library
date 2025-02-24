/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the CounterpartyInfoV3 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CounterpartyInfoV3{}

// CounterpartyInfoV3 struct for CounterpartyInfoV3
type CounterpartyInfoV3 struct {
	// The unique identifier of the counterparty [balance account](https://docs.adyen.com/api-explorer/balanceplatform/latest/post/balanceAccounts#responses-200-id).
	BalanceAccountId *string        `json:"balanceAccountId,omitempty"`
	BankAccount      *BankAccountV3 `json:"bankAccount,omitempty"`
	Card             *Card          `json:"card,omitempty"`
	// The unique identifier of the counterparty [transfer instrument](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments#responses-200-id).
	TransferInstrumentId *string `json:"transferInstrumentId,omitempty"`
}

// NewCounterpartyInfoV3 instantiates a new CounterpartyInfoV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCounterpartyInfoV3() *CounterpartyInfoV3 {
	this := CounterpartyInfoV3{}
	return &this
}

// NewCounterpartyInfoV3WithDefaults instantiates a new CounterpartyInfoV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCounterpartyInfoV3WithDefaults() *CounterpartyInfoV3 {
	this := CounterpartyInfoV3{}
	return &this
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *CounterpartyInfoV3) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CounterpartyInfoV3) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *CounterpartyInfoV3) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *CounterpartyInfoV3) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *CounterpartyInfoV3) GetBankAccount() BankAccountV3 {
	if o == nil || common.IsNil(o.BankAccount) {
		var ret BankAccountV3
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CounterpartyInfoV3) GetBankAccountOk() (*BankAccountV3, bool) {
	if o == nil || common.IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *CounterpartyInfoV3) HasBankAccount() bool {
	if o != nil && !common.IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given BankAccountV3 and assigns it to the BankAccount field.
func (o *CounterpartyInfoV3) SetBankAccount(v BankAccountV3) {
	o.BankAccount = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *CounterpartyInfoV3) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CounterpartyInfoV3) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *CounterpartyInfoV3) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *CounterpartyInfoV3) SetCard(v Card) {
	o.Card = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value if set, zero value otherwise.
func (o *CounterpartyInfoV3) GetTransferInstrumentId() string {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		var ret string
		return ret
	}
	return *o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CounterpartyInfoV3) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		return nil, false
	}
	return o.TransferInstrumentId, true
}

// HasTransferInstrumentId returns a boolean if a field has been set.
func (o *CounterpartyInfoV3) HasTransferInstrumentId() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentId) {
		return true
	}

	return false
}

// SetTransferInstrumentId gets a reference to the given string and assigns it to the TransferInstrumentId field.
func (o *CounterpartyInfoV3) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = &v
}

func (o CounterpartyInfoV3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CounterpartyInfoV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	if !common.IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.TransferInstrumentId) {
		toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	}
	return toSerialize, nil
}

type NullableCounterpartyInfoV3 struct {
	value *CounterpartyInfoV3
	isSet bool
}

func (v NullableCounterpartyInfoV3) Get() *CounterpartyInfoV3 {
	return v.value
}

func (v *NullableCounterpartyInfoV3) Set(val *CounterpartyInfoV3) {
	v.value = val
	v.isSet = true
}

func (v NullableCounterpartyInfoV3) IsSet() bool {
	return v.isSet
}

func (v *NullableCounterpartyInfoV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCounterpartyInfoV3(val *CounterpartyInfoV3) *NullableCounterpartyInfoV3 {
	return &NullableCounterpartyInfoV3{value: val, isSet: true}
}

func (v NullableCounterpartyInfoV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCounterpartyInfoV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
