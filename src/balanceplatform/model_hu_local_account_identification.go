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

// checks if the HULocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HULocalAccountIdentification{}

// HULocalAccountIdentification struct for HULocalAccountIdentification
type HULocalAccountIdentification struct {
	// The 24-digit bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// **huLocal**
	Type string `json:"type"`
}

// NewHULocalAccountIdentification instantiates a new HULocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHULocalAccountIdentification(accountNumber string, type_ string) *HULocalAccountIdentification {
	this := HULocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.Type = type_
	return &this
}

// NewHULocalAccountIdentificationWithDefaults instantiates a new HULocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHULocalAccountIdentificationWithDefaults() *HULocalAccountIdentification {
	this := HULocalAccountIdentification{}
	var type_ string = "huLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *HULocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *HULocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *HULocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetType returns the Type field value
func (o *HULocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HULocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HULocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o HULocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HULocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableHULocalAccountIdentification struct {
	value *HULocalAccountIdentification
	isSet bool
}

func (v NullableHULocalAccountIdentification) Get() *HULocalAccountIdentification {
	return v.value
}

func (v *NullableHULocalAccountIdentification) Set(val *HULocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableHULocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableHULocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHULocalAccountIdentification(val *HULocalAccountIdentification) *NullableHULocalAccountIdentification {
	return &NullableHULocalAccountIdentification{value: val, isSet: true}
}

func (v NullableHULocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHULocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *HULocalAccountIdentification) isValidType() bool {
    var allowedEnumValues = []string{ "huLocal" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

