/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the ResponseAdditionalDataDomesticError type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalDataDomesticError{}

// ResponseAdditionalDataDomesticError struct for ResponseAdditionalDataDomesticError
type ResponseAdditionalDataDomesticError struct {
	// The reason the transaction was declined, given by the local issuer.  Currently available for merchants in Japan.
	DomesticRefusalReasonRaw *string `json:"domesticRefusalReasonRaw,omitempty"`
	// The action the shopper should take, in a local language.  Currently available in Japanese, for merchants in Japan.
	DomesticShopperAdvice *string `json:"domesticShopperAdvice,omitempty"`
}

// NewResponseAdditionalDataDomesticError instantiates a new ResponseAdditionalDataDomesticError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalDataDomesticError() *ResponseAdditionalDataDomesticError {
	this := ResponseAdditionalDataDomesticError{}
	return &this
}

// NewResponseAdditionalDataDomesticErrorWithDefaults instantiates a new ResponseAdditionalDataDomesticError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalDataDomesticErrorWithDefaults() *ResponseAdditionalDataDomesticError {
	this := ResponseAdditionalDataDomesticError{}
	return &this
}

// GetDomesticRefusalReasonRaw returns the DomesticRefusalReasonRaw field value if set, zero value otherwise.
func (o *ResponseAdditionalDataDomesticError) GetDomesticRefusalReasonRaw() string {
	if o == nil || common.IsNil(o.DomesticRefusalReasonRaw) {
		var ret string
		return ret
	}
	return *o.DomesticRefusalReasonRaw
}

// GetDomesticRefusalReasonRawOk returns a tuple with the DomesticRefusalReasonRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataDomesticError) GetDomesticRefusalReasonRawOk() (*string, bool) {
	if o == nil || common.IsNil(o.DomesticRefusalReasonRaw) {
		return nil, false
	}
	return o.DomesticRefusalReasonRaw, true
}

// HasDomesticRefusalReasonRaw returns a boolean if a field has been set.
func (o *ResponseAdditionalDataDomesticError) HasDomesticRefusalReasonRaw() bool {
	if o != nil && !common.IsNil(o.DomesticRefusalReasonRaw) {
		return true
	}

	return false
}

// SetDomesticRefusalReasonRaw gets a reference to the given string and assigns it to the DomesticRefusalReasonRaw field.
func (o *ResponseAdditionalDataDomesticError) SetDomesticRefusalReasonRaw(v string) {
	o.DomesticRefusalReasonRaw = &v
}

// GetDomesticShopperAdvice returns the DomesticShopperAdvice field value if set, zero value otherwise.
func (o *ResponseAdditionalDataDomesticError) GetDomesticShopperAdvice() string {
	if o == nil || common.IsNil(o.DomesticShopperAdvice) {
		var ret string
		return ret
	}
	return *o.DomesticShopperAdvice
}

// GetDomesticShopperAdviceOk returns a tuple with the DomesticShopperAdvice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataDomesticError) GetDomesticShopperAdviceOk() (*string, bool) {
	if o == nil || common.IsNil(o.DomesticShopperAdvice) {
		return nil, false
	}
	return o.DomesticShopperAdvice, true
}

// HasDomesticShopperAdvice returns a boolean if a field has been set.
func (o *ResponseAdditionalDataDomesticError) HasDomesticShopperAdvice() bool {
	if o != nil && !common.IsNil(o.DomesticShopperAdvice) {
		return true
	}

	return false
}

// SetDomesticShopperAdvice gets a reference to the given string and assigns it to the DomesticShopperAdvice field.
func (o *ResponseAdditionalDataDomesticError) SetDomesticShopperAdvice(v string) {
	o.DomesticShopperAdvice = &v
}

func (o ResponseAdditionalDataDomesticError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalDataDomesticError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DomesticRefusalReasonRaw) {
		toSerialize["domesticRefusalReasonRaw"] = o.DomesticRefusalReasonRaw
	}
	if !common.IsNil(o.DomesticShopperAdvice) {
		toSerialize["domesticShopperAdvice"] = o.DomesticShopperAdvice
	}
	return toSerialize, nil
}

type NullableResponseAdditionalDataDomesticError struct {
	value *ResponseAdditionalDataDomesticError
	isSet bool
}

func (v NullableResponseAdditionalDataDomesticError) Get() *ResponseAdditionalDataDomesticError {
	return v.value
}

func (v *NullableResponseAdditionalDataDomesticError) Set(val *ResponseAdditionalDataDomesticError) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalDataDomesticError) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalDataDomesticError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalDataDomesticError(val *ResponseAdditionalDataDomesticError) *NullableResponseAdditionalDataDomesticError {
	return &NullableResponseAdditionalDataDomesticError{value: val, isSet: true}
}

func (v NullableResponseAdditionalDataDomesticError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalDataDomesticError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
