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

// checks if the NumberAndBicAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NumberAndBicAccountIdentification{}

// NumberAndBicAccountIdentification struct for NumberAndBicAccountIdentification
type NumberAndBicAccountIdentification struct {
	// The bank account number, without separators or whitespace. The length and format depends on the bank or country.
	AccountNumber                string                        `json:"accountNumber"`
	AdditionalBankIdentification *AdditionalBankIdentification `json:"additionalBankIdentification,omitempty"`
	// The bank's 8- or 11-character BIC or SWIFT code.
	Bic string `json:"bic"`
	// **numberAndBic**
	Type string `json:"type"`
}

// NewNumberAndBicAccountIdentification instantiates a new NumberAndBicAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberAndBicAccountIdentification(accountNumber string, bic string, type_ string) *NumberAndBicAccountIdentification {
	this := NumberAndBicAccountIdentification{}
	this.AccountNumber = accountNumber
	this.Bic = bic
	this.Type = type_
	return &this
}

// NewNumberAndBicAccountIdentificationWithDefaults instantiates a new NumberAndBicAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberAndBicAccountIdentificationWithDefaults() *NumberAndBicAccountIdentification {
	this := NumberAndBicAccountIdentification{}
	var type_ string = "numberAndBic"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *NumberAndBicAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *NumberAndBicAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAdditionalBankIdentification returns the AdditionalBankIdentification field value if set, zero value otherwise.
func (o *NumberAndBicAccountIdentification) GetAdditionalBankIdentification() AdditionalBankIdentification {
	if o == nil || common.IsNil(o.AdditionalBankIdentification) {
		var ret AdditionalBankIdentification
		return ret
	}
	return *o.AdditionalBankIdentification
}

// GetAdditionalBankIdentificationOk returns a tuple with the AdditionalBankIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetAdditionalBankIdentificationOk() (*AdditionalBankIdentification, bool) {
	if o == nil || common.IsNil(o.AdditionalBankIdentification) {
		return nil, false
	}
	return o.AdditionalBankIdentification, true
}

// HasAdditionalBankIdentification returns a boolean if a field has been set.
func (o *NumberAndBicAccountIdentification) HasAdditionalBankIdentification() bool {
	if o != nil && !common.IsNil(o.AdditionalBankIdentification) {
		return true
	}

	return false
}

// SetAdditionalBankIdentification gets a reference to the given AdditionalBankIdentification and assigns it to the AdditionalBankIdentification field.
func (o *NumberAndBicAccountIdentification) SetAdditionalBankIdentification(v AdditionalBankIdentification) {
	o.AdditionalBankIdentification = &v
}

// GetBic returns the Bic field value
func (o *NumberAndBicAccountIdentification) GetBic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bic
}

// GetBicOk returns a tuple with the Bic field value
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetBicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bic, true
}

// SetBic sets field value
func (o *NumberAndBicAccountIdentification) SetBic(v string) {
	o.Bic = v
}

// GetType returns the Type field value
func (o *NumberAndBicAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NumberAndBicAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o NumberAndBicAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumberAndBicAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	if !common.IsNil(o.AdditionalBankIdentification) {
		toSerialize["additionalBankIdentification"] = o.AdditionalBankIdentification
	}
	toSerialize["bic"] = o.Bic
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNumberAndBicAccountIdentification struct {
	value *NumberAndBicAccountIdentification
	isSet bool
}

func (v NullableNumberAndBicAccountIdentification) Get() *NumberAndBicAccountIdentification {
	return v.value
}

func (v *NullableNumberAndBicAccountIdentification) Set(val *NumberAndBicAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberAndBicAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberAndBicAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberAndBicAccountIdentification(val *NumberAndBicAccountIdentification) *NullableNumberAndBicAccountIdentification {
	return &NullableNumberAndBicAccountIdentification{value: val, isSet: true}
}

func (v NullableNumberAndBicAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberAndBicAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *NumberAndBicAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"numberAndBic"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
