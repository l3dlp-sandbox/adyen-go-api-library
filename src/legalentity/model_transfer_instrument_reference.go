/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the TransferInstrumentReference type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferInstrumentReference{}

// TransferInstrumentReference struct for TransferInstrumentReference
type TransferInstrumentReference struct {
	// The masked IBAN or bank account number.
	AccountIdentifier string `json:"accountIdentifier"`
	// The unique identifier of the resource.
	Id string `json:"id"`
	// Four last digits of the bank account number. If the transfer instrument is created using [instant bank account verification](https://docs.adyen.com/release-notes/platforms-and-financial-products#releaseNote=2023-05-08-hosted-onboarding), and it is a virtual bank account, these digits may be different from the last four digits of the masked account number.
	RealLastFour *string `json:"realLastFour,omitempty"`
	// Identifies if the bank account was created through [instant bank verification](https://docs.adyen.com/release-notes/platforms-and-financial-products#releaseNote=2023-05-08-hosted-onboarding).
	TrustedSource *bool `json:"trustedSource,omitempty"`
}

// NewTransferInstrumentReference instantiates a new TransferInstrumentReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferInstrumentReference(accountIdentifier string, id string) *TransferInstrumentReference {
	this := TransferInstrumentReference{}
	this.AccountIdentifier = accountIdentifier
	this.Id = id
	return &this
}

// NewTransferInstrumentReferenceWithDefaults instantiates a new TransferInstrumentReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferInstrumentReferenceWithDefaults() *TransferInstrumentReference {
	this := TransferInstrumentReference{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value
func (o *TransferInstrumentReference) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *TransferInstrumentReference) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value
func (o *TransferInstrumentReference) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetId returns the Id field value
func (o *TransferInstrumentReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferInstrumentReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferInstrumentReference) SetId(v string) {
	o.Id = v
}

// GetRealLastFour returns the RealLastFour field value if set, zero value otherwise.
func (o *TransferInstrumentReference) GetRealLastFour() string {
	if o == nil || common.IsNil(o.RealLastFour) {
		var ret string
		return ret
	}
	return *o.RealLastFour
}

// GetRealLastFourOk returns a tuple with the RealLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInstrumentReference) GetRealLastFourOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealLastFour) {
		return nil, false
	}
	return o.RealLastFour, true
}

// HasRealLastFour returns a boolean if a field has been set.
func (o *TransferInstrumentReference) HasRealLastFour() bool {
	if o != nil && !common.IsNil(o.RealLastFour) {
		return true
	}

	return false
}

// SetRealLastFour gets a reference to the given string and assigns it to the RealLastFour field.
func (o *TransferInstrumentReference) SetRealLastFour(v string) {
	o.RealLastFour = &v
}

// GetTrustedSource returns the TrustedSource field value if set, zero value otherwise.
func (o *TransferInstrumentReference) GetTrustedSource() bool {
	if o == nil || common.IsNil(o.TrustedSource) {
		var ret bool
		return ret
	}
	return *o.TrustedSource
}

// GetTrustedSourceOk returns a tuple with the TrustedSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInstrumentReference) GetTrustedSourceOk() (*bool, bool) {
	if o == nil || common.IsNil(o.TrustedSource) {
		return nil, false
	}
	return o.TrustedSource, true
}

// HasTrustedSource returns a boolean if a field has been set.
func (o *TransferInstrumentReference) HasTrustedSource() bool {
	if o != nil && !common.IsNil(o.TrustedSource) {
		return true
	}

	return false
}

// SetTrustedSource gets a reference to the given bool and assigns it to the TrustedSource field.
func (o *TransferInstrumentReference) SetTrustedSource(v bool) {
	o.TrustedSource = &v
}

func (o TransferInstrumentReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferInstrumentReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountIdentifier"] = o.AccountIdentifier
	toSerialize["id"] = o.Id
	if !common.IsNil(o.RealLastFour) {
		toSerialize["realLastFour"] = o.RealLastFour
	}
	if !common.IsNil(o.TrustedSource) {
		toSerialize["trustedSource"] = o.TrustedSource
	}
	return toSerialize, nil
}

type NullableTransferInstrumentReference struct {
	value *TransferInstrumentReference
	isSet bool
}

func (v NullableTransferInstrumentReference) Get() *TransferInstrumentReference {
	return v.value
}

func (v *NullableTransferInstrumentReference) Set(val *TransferInstrumentReference) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferInstrumentReference) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferInstrumentReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferInstrumentReference(val *TransferInstrumentReference) *NullableTransferInstrumentReference {
	return &NullableTransferInstrumentReference{value: val, isSet: true}
}

func (v NullableTransferInstrumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferInstrumentReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
