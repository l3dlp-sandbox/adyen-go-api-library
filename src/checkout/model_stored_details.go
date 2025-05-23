/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the StoredDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredDetails{}

// StoredDetails struct for StoredDetails
type StoredDetails struct {
	Bank *BankAccount `json:"bank,omitempty"`
	Card *Card        `json:"card,omitempty"`
	// The email associated with stored payment details.
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// NewStoredDetails instantiates a new StoredDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredDetails() *StoredDetails {
	this := StoredDetails{}
	return &this
}

// NewStoredDetailsWithDefaults instantiates a new StoredDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredDetailsWithDefaults() *StoredDetails {
	this := StoredDetails{}
	return &this
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *StoredDetails) GetBank() BankAccount {
	if o == nil || common.IsNil(o.Bank) {
		var ret BankAccount
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredDetails) GetBankOk() (*BankAccount, bool) {
	if o == nil || common.IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *StoredDetails) HasBank() bool {
	if o != nil && !common.IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given BankAccount and assigns it to the Bank field.
func (o *StoredDetails) SetBank(v BankAccount) {
	o.Bank = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *StoredDetails) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredDetails) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *StoredDetails) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *StoredDetails) SetCard(v Card) {
	o.Card = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *StoredDetails) GetEmailAddress() string {
	if o == nil || common.IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredDetails) GetEmailAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *StoredDetails) HasEmailAddress() bool {
	if o != nil && !common.IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *StoredDetails) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

func (o StoredDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	return toSerialize, nil
}

type NullableStoredDetails struct {
	value *StoredDetails
	isSet bool
}

func (v NullableStoredDetails) Get() *StoredDetails {
	return v.value
}

func (v *NullableStoredDetails) Set(val *StoredDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredDetails(val *StoredDetails) *NullableStoredDetails {
	return &NullableStoredDetails{value: val, isSet: true}
}

func (v NullableStoredDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
