/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the NOLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NOLocalAccountIdentification{}

// NOLocalAccountIdentification struct for NOLocalAccountIdentification
type NOLocalAccountIdentification struct {
	// The 11-digit bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// **noLocal**
	Type string `json:"type"`
}

// NewNOLocalAccountIdentification instantiates a new NOLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNOLocalAccountIdentification(accountNumber string, type_ string) *NOLocalAccountIdentification {
	this := NOLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.Type = type_
	return &this
}

// NewNOLocalAccountIdentificationWithDefaults instantiates a new NOLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNOLocalAccountIdentificationWithDefaults() *NOLocalAccountIdentification {
	this := NOLocalAccountIdentification{}
	var type_ string = "noLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *NOLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *NOLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *NOLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetType returns the Type field value
func (o *NOLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NOLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NOLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o NOLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NOLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNOLocalAccountIdentification struct {
	value *NOLocalAccountIdentification
	isSet bool
}

func (v NullableNOLocalAccountIdentification) Get() *NOLocalAccountIdentification {
	return v.value
}

func (v *NullableNOLocalAccountIdentification) Set(val *NOLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableNOLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableNOLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNOLocalAccountIdentification(val *NOLocalAccountIdentification) *NullableNOLocalAccountIdentification {
	return &NullableNOLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableNOLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNOLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *NOLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"noLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
