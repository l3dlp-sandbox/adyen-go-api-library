/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the DotpayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DotpayDetails{}

// DotpayDetails struct for DotpayDetails
type DotpayDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The Dotpay issuer value of the shopper's selected bank. Set this to an **id** of a Dotpay issuer to preselect it.
	Issuer string `json:"issuer"`
	// **dotpay**
	Type *string `json:"type,omitempty"`
}

// NewDotpayDetails instantiates a new DotpayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDotpayDetails(issuer string) *DotpayDetails {
	this := DotpayDetails{}
	this.Issuer = issuer
	var type_ string = "dotpay"
	this.Type = &type_
	return &this
}

// NewDotpayDetailsWithDefaults instantiates a new DotpayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDotpayDetailsWithDefaults() *DotpayDetails {
	this := DotpayDetails{}
	var type_ string = "dotpay"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *DotpayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DotpayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *DotpayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *DotpayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetIssuer returns the Issuer field value
func (o *DotpayDetails) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *DotpayDetails) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *DotpayDetails) SetIssuer(v string) {
	o.Issuer = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DotpayDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DotpayDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DotpayDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DotpayDetails) SetType(v string) {
	o.Type = &v
}

func (o DotpayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DotpayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	toSerialize["issuer"] = o.Issuer
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDotpayDetails struct {
	value *DotpayDetails
	isSet bool
}

func (v NullableDotpayDetails) Get() *DotpayDetails {
	return v.value
}

func (v *NullableDotpayDetails) Set(val *DotpayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDotpayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDotpayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDotpayDetails(val *DotpayDetails) *NullableDotpayDetails {
	return &NullableDotpayDetails{value: val, isSet: true}
}

func (v NullableDotpayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDotpayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DotpayDetails) isValidType() bool {
	var allowedEnumValues = []string{"dotpay"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
