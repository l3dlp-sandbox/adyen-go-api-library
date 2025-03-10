/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the Card type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Card{}

// Card struct for Card
type Card struct {
	CardHolder         PartyIdentification `json:"cardHolder"`
	CardIdentification CardIdentification  `json:"cardIdentification"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard(cardHolder PartyIdentification, cardIdentification CardIdentification) *Card {
	this := Card{}
	this.CardHolder = cardHolder
	this.CardIdentification = cardIdentification
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetCardHolder returns the CardHolder field value
func (o *Card) GetCardHolder() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.CardHolder
}

// GetCardHolderOk returns a tuple with the CardHolder field value
// and a boolean to check if the value has been set.
func (o *Card) GetCardHolderOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardHolder, true
}

// SetCardHolder sets field value
func (o *Card) SetCardHolder(v PartyIdentification) {
	o.CardHolder = v
}

// GetCardIdentification returns the CardIdentification field value
func (o *Card) GetCardIdentification() CardIdentification {
	if o == nil {
		var ret CardIdentification
		return ret
	}

	return o.CardIdentification
}

// GetCardIdentificationOk returns a tuple with the CardIdentification field value
// and a boolean to check if the value has been set.
func (o *Card) GetCardIdentificationOk() (*CardIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardIdentification, true
}

// SetCardIdentification sets field value
func (o *Card) SetCardIdentification(v CardIdentification) {
	o.CardIdentification = v
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Card) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cardHolder"] = o.CardHolder
	toSerialize["cardIdentification"] = o.CardIdentification
	return toSerialize, nil
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
