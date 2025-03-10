/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the EBankingFinlandDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EBankingFinlandDetails{}

// EBankingFinlandDetails struct for EBankingFinlandDetails
type EBankingFinlandDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The Ebanking Finland issuer value of the shopper's selected bank.
	Issuer *string `json:"issuer,omitempty"`
	// **ebanking_FI**
	Type string `json:"type"`
}

// NewEBankingFinlandDetails instantiates a new EBankingFinlandDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEBankingFinlandDetails(type_ string) *EBankingFinlandDetails {
	this := EBankingFinlandDetails{}
	this.Type = type_
	return &this
}

// NewEBankingFinlandDetailsWithDefaults instantiates a new EBankingFinlandDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEBankingFinlandDetailsWithDefaults() *EBankingFinlandDetails {
	this := EBankingFinlandDetails{}
	var type_ string = "ebanking_FI"
	this.Type = type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *EBankingFinlandDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBankingFinlandDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *EBankingFinlandDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *EBankingFinlandDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *EBankingFinlandDetails) GetIssuer() string {
	if o == nil || common.IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBankingFinlandDetails) GetIssuerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *EBankingFinlandDetails) HasIssuer() bool {
	if o != nil && !common.IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *EBankingFinlandDetails) SetIssuer(v string) {
	o.Issuer = &v
}

// GetType returns the Type field value
func (o *EBankingFinlandDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EBankingFinlandDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EBankingFinlandDetails) SetType(v string) {
	o.Type = v
}

func (o EBankingFinlandDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EBankingFinlandDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableEBankingFinlandDetails struct {
	value *EBankingFinlandDetails
	isSet bool
}

func (v NullableEBankingFinlandDetails) Get() *EBankingFinlandDetails {
	return v.value
}

func (v *NullableEBankingFinlandDetails) Set(val *EBankingFinlandDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEBankingFinlandDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEBankingFinlandDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBankingFinlandDetails(val *EBankingFinlandDetails) *NullableEBankingFinlandDetails {
	return &NullableEBankingFinlandDetails{value: val, isSet: true}
}

func (v NullableEBankingFinlandDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBankingFinlandDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *EBankingFinlandDetails) isValidType() bool {
	var allowedEnumValues = []string{"ebanking_FI"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
