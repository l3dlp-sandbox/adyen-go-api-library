/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Passcodes type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Passcodes{}

// Passcodes struct for Passcodes
type Passcodes struct {
	// The passcode for the Admin menu and the Settings menu.
	AdminMenuPin *string `json:"adminMenuPin,omitempty"`
	// The passcode for referenced and unreferenced refunds on standalone terminals.
	RefundPin *string `json:"refundPin,omitempty"`
	// The passcode to unlock the terminal screen after a timeout.
	ScreenLockPin *string `json:"screenLockPin,omitempty"`
	// The passcode for the Transactions menu.
	TxMenuPin *string `json:"txMenuPin,omitempty"`
}

// NewPasscodes instantiates a new Passcodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasscodes() *Passcodes {
	this := Passcodes{}
	return &this
}

// NewPasscodesWithDefaults instantiates a new Passcodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasscodesWithDefaults() *Passcodes {
	this := Passcodes{}
	return &this
}

// GetAdminMenuPin returns the AdminMenuPin field value if set, zero value otherwise.
func (o *Passcodes) GetAdminMenuPin() string {
	if o == nil || common.IsNil(o.AdminMenuPin) {
		var ret string
		return ret
	}
	return *o.AdminMenuPin
}

// GetAdminMenuPinOk returns a tuple with the AdminMenuPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Passcodes) GetAdminMenuPinOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdminMenuPin) {
		return nil, false
	}
	return o.AdminMenuPin, true
}

// HasAdminMenuPin returns a boolean if a field has been set.
func (o *Passcodes) HasAdminMenuPin() bool {
	if o != nil && !common.IsNil(o.AdminMenuPin) {
		return true
	}

	return false
}

// SetAdminMenuPin gets a reference to the given string and assigns it to the AdminMenuPin field.
func (o *Passcodes) SetAdminMenuPin(v string) {
	o.AdminMenuPin = &v
}

// GetRefundPin returns the RefundPin field value if set, zero value otherwise.
func (o *Passcodes) GetRefundPin() string {
	if o == nil || common.IsNil(o.RefundPin) {
		var ret string
		return ret
	}
	return *o.RefundPin
}

// GetRefundPinOk returns a tuple with the RefundPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Passcodes) GetRefundPinOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefundPin) {
		return nil, false
	}
	return o.RefundPin, true
}

// HasRefundPin returns a boolean if a field has been set.
func (o *Passcodes) HasRefundPin() bool {
	if o != nil && !common.IsNil(o.RefundPin) {
		return true
	}

	return false
}

// SetRefundPin gets a reference to the given string and assigns it to the RefundPin field.
func (o *Passcodes) SetRefundPin(v string) {
	o.RefundPin = &v
}

// GetScreenLockPin returns the ScreenLockPin field value if set, zero value otherwise.
func (o *Passcodes) GetScreenLockPin() string {
	if o == nil || common.IsNil(o.ScreenLockPin) {
		var ret string
		return ret
	}
	return *o.ScreenLockPin
}

// GetScreenLockPinOk returns a tuple with the ScreenLockPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Passcodes) GetScreenLockPinOk() (*string, bool) {
	if o == nil || common.IsNil(o.ScreenLockPin) {
		return nil, false
	}
	return o.ScreenLockPin, true
}

// HasScreenLockPin returns a boolean if a field has been set.
func (o *Passcodes) HasScreenLockPin() bool {
	if o != nil && !common.IsNil(o.ScreenLockPin) {
		return true
	}

	return false
}

// SetScreenLockPin gets a reference to the given string and assigns it to the ScreenLockPin field.
func (o *Passcodes) SetScreenLockPin(v string) {
	o.ScreenLockPin = &v
}

// GetTxMenuPin returns the TxMenuPin field value if set, zero value otherwise.
func (o *Passcodes) GetTxMenuPin() string {
	if o == nil || common.IsNil(o.TxMenuPin) {
		var ret string
		return ret
	}
	return *o.TxMenuPin
}

// GetTxMenuPinOk returns a tuple with the TxMenuPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Passcodes) GetTxMenuPinOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxMenuPin) {
		return nil, false
	}
	return o.TxMenuPin, true
}

// HasTxMenuPin returns a boolean if a field has been set.
func (o *Passcodes) HasTxMenuPin() bool {
	if o != nil && !common.IsNil(o.TxMenuPin) {
		return true
	}

	return false
}

// SetTxMenuPin gets a reference to the given string and assigns it to the TxMenuPin field.
func (o *Passcodes) SetTxMenuPin(v string) {
	o.TxMenuPin = &v
}

func (o Passcodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Passcodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdminMenuPin) {
		toSerialize["adminMenuPin"] = o.AdminMenuPin
	}
	if !common.IsNil(o.RefundPin) {
		toSerialize["refundPin"] = o.RefundPin
	}
	if !common.IsNil(o.ScreenLockPin) {
		toSerialize["screenLockPin"] = o.ScreenLockPin
	}
	if !common.IsNil(o.TxMenuPin) {
		toSerialize["txMenuPin"] = o.TxMenuPin
	}
	return toSerialize, nil
}

type NullablePasscodes struct {
	value *Passcodes
	isSet bool
}

func (v NullablePasscodes) Get() *Passcodes {
	return v.value
}

func (v *NullablePasscodes) Set(val *Passcodes) {
	v.value = val
	v.isSet = true
}

func (v NullablePasscodes) IsSet() bool {
	return v.isSet
}

func (v *NullablePasscodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasscodes(val *Passcodes) *NullablePasscodes {
	return &NullablePasscodes{value: val, isSet: true}
}

func (v NullablePasscodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasscodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



