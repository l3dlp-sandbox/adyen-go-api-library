/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the SupportedCardTypes type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SupportedCardTypes{}

// SupportedCardTypes struct for SupportedCardTypes
type SupportedCardTypes struct {
	// Set to **true** to accept credit cards.
	Credit *bool `json:"credit,omitempty"`
	// Set to **true** to accept debit cards.
	Debit *bool `json:"debit,omitempty"`
	// Set to **true** to accept cards that allow deferred debit.
	DeferredDebit *bool `json:"deferredDebit,omitempty"`
	// Set to **true** to accept prepaid cards.
	Prepaid *bool `json:"prepaid,omitempty"`
	// Set to **true** to accept card types for which the terminal can't determine the funding source while offline.
	Unknown *bool `json:"unknown,omitempty"`
}

// NewSupportedCardTypes instantiates a new SupportedCardTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedCardTypes() *SupportedCardTypes {
	this := SupportedCardTypes{}
	return &this
}

// NewSupportedCardTypesWithDefaults instantiates a new SupportedCardTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedCardTypesWithDefaults() *SupportedCardTypes {
	this := SupportedCardTypes{}
	return &this
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *SupportedCardTypes) GetCredit() bool {
	if o == nil || common.IsNil(o.Credit) {
		var ret bool
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedCardTypes) GetCreditOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *SupportedCardTypes) HasCredit() bool {
	if o != nil && !common.IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given bool and assigns it to the Credit field.
func (o *SupportedCardTypes) SetCredit(v bool) {
	o.Credit = &v
}

// GetDebit returns the Debit field value if set, zero value otherwise.
func (o *SupportedCardTypes) GetDebit() bool {
	if o == nil || common.IsNil(o.Debit) {
		var ret bool
		return ret
	}
	return *o.Debit
}

// GetDebitOk returns a tuple with the Debit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedCardTypes) GetDebitOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Debit) {
		return nil, false
	}
	return o.Debit, true
}

// HasDebit returns a boolean if a field has been set.
func (o *SupportedCardTypes) HasDebit() bool {
	if o != nil && !common.IsNil(o.Debit) {
		return true
	}

	return false
}

// SetDebit gets a reference to the given bool and assigns it to the Debit field.
func (o *SupportedCardTypes) SetDebit(v bool) {
	o.Debit = &v
}

// GetDeferredDebit returns the DeferredDebit field value if set, zero value otherwise.
func (o *SupportedCardTypes) GetDeferredDebit() bool {
	if o == nil || common.IsNil(o.DeferredDebit) {
		var ret bool
		return ret
	}
	return *o.DeferredDebit
}

// GetDeferredDebitOk returns a tuple with the DeferredDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedCardTypes) GetDeferredDebitOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DeferredDebit) {
		return nil, false
	}
	return o.DeferredDebit, true
}

// HasDeferredDebit returns a boolean if a field has been set.
func (o *SupportedCardTypes) HasDeferredDebit() bool {
	if o != nil && !common.IsNil(o.DeferredDebit) {
		return true
	}

	return false
}

// SetDeferredDebit gets a reference to the given bool and assigns it to the DeferredDebit field.
func (o *SupportedCardTypes) SetDeferredDebit(v bool) {
	o.DeferredDebit = &v
}

// GetPrepaid returns the Prepaid field value if set, zero value otherwise.
func (o *SupportedCardTypes) GetPrepaid() bool {
	if o == nil || common.IsNil(o.Prepaid) {
		var ret bool
		return ret
	}
	return *o.Prepaid
}

// GetPrepaidOk returns a tuple with the Prepaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedCardTypes) GetPrepaidOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Prepaid) {
		return nil, false
	}
	return o.Prepaid, true
}

// HasPrepaid returns a boolean if a field has been set.
func (o *SupportedCardTypes) HasPrepaid() bool {
	if o != nil && !common.IsNil(o.Prepaid) {
		return true
	}

	return false
}

// SetPrepaid gets a reference to the given bool and assigns it to the Prepaid field.
func (o *SupportedCardTypes) SetPrepaid(v bool) {
	o.Prepaid = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *SupportedCardTypes) GetUnknown() bool {
	if o == nil || common.IsNil(o.Unknown) {
		var ret bool
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedCardTypes) GetUnknownOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Unknown) {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *SupportedCardTypes) HasUnknown() bool {
	if o != nil && !common.IsNil(o.Unknown) {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given bool and assigns it to the Unknown field.
func (o *SupportedCardTypes) SetUnknown(v bool) {
	o.Unknown = &v
}

func (o SupportedCardTypes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedCardTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !common.IsNil(o.Debit) {
		toSerialize["debit"] = o.Debit
	}
	if !common.IsNil(o.DeferredDebit) {
		toSerialize["deferredDebit"] = o.DeferredDebit
	}
	if !common.IsNil(o.Prepaid) {
		toSerialize["prepaid"] = o.Prepaid
	}
	if !common.IsNil(o.Unknown) {
		toSerialize["unknown"] = o.Unknown
	}
	return toSerialize, nil
}

type NullableSupportedCardTypes struct {
	value *SupportedCardTypes
	isSet bool
}

func (v NullableSupportedCardTypes) Get() *SupportedCardTypes {
	return v.value
}

func (v *NullableSupportedCardTypes) Set(val *SupportedCardTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedCardTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedCardTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedCardTypes(val *SupportedCardTypes) *NullableSupportedCardTypes {
	return &NullableSupportedCardTypes{value: val, isSet: true}
}

func (v NullableSupportedCardTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedCardTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
