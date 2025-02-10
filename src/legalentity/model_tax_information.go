/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the TaxInformation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TaxInformation{}

// TaxInformation struct for TaxInformation
type TaxInformation struct {
	// The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code.
	Country *string `json:"country,omitempty"`
	// The tax ID number (TIN) of the organization or individual.
	Number *string `json:"number,omitempty"`
	// The TIN type depending on the country where it was issued. Only provide if the country has multiple tax IDs: Singapore, Sweden, the UK, or the US. For example, provide **SSN**, **EIN**, or **ITIN** for the US.
	Type *string `json:"type,omitempty"`
}

// NewTaxInformation instantiates a new TaxInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxInformation() *TaxInformation {
	this := TaxInformation{}
	return &this
}

// NewTaxInformationWithDefaults instantiates a new TaxInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxInformationWithDefaults() *TaxInformation {
	this := TaxInformation{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TaxInformation) GetCountry() string {
	if o == nil || common.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxInformation) GetCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TaxInformation) HasCountry() bool {
	if o != nil && !common.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TaxInformation) SetCountry(v string) {
	o.Country = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *TaxInformation) GetNumber() string {
	if o == nil || common.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxInformation) GetNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *TaxInformation) HasNumber() bool {
	if o != nil && !common.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *TaxInformation) SetNumber(v string) {
	o.Number = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaxInformation) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxInformation) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaxInformation) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaxInformation) SetType(v string) {
	o.Type = &v
}

func (o TaxInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !common.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTaxInformation struct {
	value *TaxInformation
	isSet bool
}

func (v NullableTaxInformation) Get() *TaxInformation {
	return v.value
}

func (v *NullableTaxInformation) Set(val *TaxInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxInformation(val *TaxInformation) *NullableTaxInformation {
	return &NullableTaxInformation{value: val, isSet: true}
}

func (v NullableTaxInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
