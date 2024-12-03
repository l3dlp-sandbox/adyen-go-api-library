/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the BalanceAccountNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceAccountNotificationData{}

// BalanceAccountNotificationData struct for BalanceAccountNotificationData
type BalanceAccountNotificationData struct {
	BalanceAccount *BalanceAccount `json:"balanceAccount,omitempty"`
	// The unique identifier of the balance platform.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
}

// NewBalanceAccountNotificationData instantiates a new BalanceAccountNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceAccountNotificationData() *BalanceAccountNotificationData {
	this := BalanceAccountNotificationData{}
	return &this
}

// NewBalanceAccountNotificationDataWithDefaults instantiates a new BalanceAccountNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceAccountNotificationDataWithDefaults() *BalanceAccountNotificationData {
	this := BalanceAccountNotificationData{}
	return &this
}

// GetBalanceAccount returns the BalanceAccount field value if set, zero value otherwise.
func (o *BalanceAccountNotificationData) GetBalanceAccount() BalanceAccount {
	if o == nil || common.IsNil(o.BalanceAccount) {
		var ret BalanceAccount
		return ret
	}
	return *o.BalanceAccount
}

// GetBalanceAccountOk returns a tuple with the BalanceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountNotificationData) GetBalanceAccountOk() (*BalanceAccount, bool) {
	if o == nil || common.IsNil(o.BalanceAccount) {
		return nil, false
	}
	return o.BalanceAccount, true
}

// HasBalanceAccount returns a boolean if a field has been set.
func (o *BalanceAccountNotificationData) HasBalanceAccount() bool {
	if o != nil && !common.IsNil(o.BalanceAccount) {
		return true
	}

	return false
}

// SetBalanceAccount gets a reference to the given BalanceAccount and assigns it to the BalanceAccount field.
func (o *BalanceAccountNotificationData) SetBalanceAccount(v BalanceAccount) {
	o.BalanceAccount = &v
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *BalanceAccountNotificationData) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountNotificationData) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *BalanceAccountNotificationData) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *BalanceAccountNotificationData) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

func (o BalanceAccountNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceAccountNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccount) {
		toSerialize["balanceAccount"] = o.BalanceAccount
	}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	return toSerialize, nil
}

type NullableBalanceAccountNotificationData struct {
	value *BalanceAccountNotificationData
	isSet bool
}

func (v NullableBalanceAccountNotificationData) Get() *BalanceAccountNotificationData {
	return v.value
}

func (v *NullableBalanceAccountNotificationData) Set(val *BalanceAccountNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceAccountNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceAccountNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceAccountNotificationData(val *BalanceAccountNotificationData) *NullableBalanceAccountNotificationData {
	return &NullableBalanceAccountNotificationData{value: val, isSet: true}
}

func (v NullableBalanceAccountNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceAccountNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



