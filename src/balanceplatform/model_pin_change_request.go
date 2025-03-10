/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the PinChangeRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PinChangeRequest{}

// PinChangeRequest struct for PinChangeRequest
type PinChangeRequest struct {
	// The symmetric session key that you encrypted with the [public key](https://docs.adyen.com/api-explorer/balanceplatform/2/get/publicKey) that you received from Adyen.
	EncryptedKey string `json:"encryptedKey"`
	// The encrypted [PIN block](https://www.pcisecuritystandards.org/glossary/pin-block).
	EncryptedPinBlock string `json:"encryptedPinBlock"`
	// The unique identifier of the payment instrument, which is the card for which you are managing the PIN.
	PaymentInstrumentId string `json:"paymentInstrumentId"`
	// The 16-digit token that you used to generate the `encryptedPinBlock`.
	Token string `json:"token"`
}

// NewPinChangeRequest instantiates a new PinChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinChangeRequest(encryptedKey string, encryptedPinBlock string, paymentInstrumentId string, token string) *PinChangeRequest {
	this := PinChangeRequest{}
	this.EncryptedKey = encryptedKey
	this.EncryptedPinBlock = encryptedPinBlock
	this.PaymentInstrumentId = paymentInstrumentId
	this.Token = token
	return &this
}

// NewPinChangeRequestWithDefaults instantiates a new PinChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinChangeRequestWithDefaults() *PinChangeRequest {
	this := PinChangeRequest{}
	return &this
}

// GetEncryptedKey returns the EncryptedKey field value
func (o *PinChangeRequest) GetEncryptedKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptedKey
}

// GetEncryptedKeyOk returns a tuple with the EncryptedKey field value
// and a boolean to check if the value has been set.
func (o *PinChangeRequest) GetEncryptedKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedKey, true
}

// SetEncryptedKey sets field value
func (o *PinChangeRequest) SetEncryptedKey(v string) {
	o.EncryptedKey = v
}

// GetEncryptedPinBlock returns the EncryptedPinBlock field value
func (o *PinChangeRequest) GetEncryptedPinBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptedPinBlock
}

// GetEncryptedPinBlockOk returns a tuple with the EncryptedPinBlock field value
// and a boolean to check if the value has been set.
func (o *PinChangeRequest) GetEncryptedPinBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedPinBlock, true
}

// SetEncryptedPinBlock sets field value
func (o *PinChangeRequest) SetEncryptedPinBlock(v string) {
	o.EncryptedPinBlock = v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value
func (o *PinChangeRequest) GetPaymentInstrumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value
// and a boolean to check if the value has been set.
func (o *PinChangeRequest) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentInstrumentId, true
}

// SetPaymentInstrumentId sets field value
func (o *PinChangeRequest) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = v
}

// GetToken returns the Token field value
func (o *PinChangeRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PinChangeRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PinChangeRequest) SetToken(v string) {
	o.Token = v
}

func (o PinChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinChangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encryptedKey"] = o.EncryptedKey
	toSerialize["encryptedPinBlock"] = o.EncryptedPinBlock
	toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullablePinChangeRequest struct {
	value *PinChangeRequest
	isSet bool
}

func (v NullablePinChangeRequest) Get() *PinChangeRequest {
	return v.value
}

func (v *NullablePinChangeRequest) Set(val *PinChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePinChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePinChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinChangeRequest(val *PinChangeRequest) *NullablePinChangeRequest {
	return &NullablePinChangeRequest{value: val, isSet: true}
}

func (v NullablePinChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
