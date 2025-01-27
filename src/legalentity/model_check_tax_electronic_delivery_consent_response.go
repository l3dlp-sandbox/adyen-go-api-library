/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the CheckTaxElectronicDeliveryConsentResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckTaxElectronicDeliveryConsentResponse{}

// CheckTaxElectronicDeliveryConsentResponse struct for CheckTaxElectronicDeliveryConsentResponse
type CheckTaxElectronicDeliveryConsentResponse struct {
	// Consent to electronically deliver tax form US1099-K.
	US1099k *bool `json:"US1099k,omitempty"`
}

// NewCheckTaxElectronicDeliveryConsentResponse instantiates a new CheckTaxElectronicDeliveryConsentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckTaxElectronicDeliveryConsentResponse() *CheckTaxElectronicDeliveryConsentResponse {
	this := CheckTaxElectronicDeliveryConsentResponse{}
	return &this
}

// NewCheckTaxElectronicDeliveryConsentResponseWithDefaults instantiates a new CheckTaxElectronicDeliveryConsentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckTaxElectronicDeliveryConsentResponseWithDefaults() *CheckTaxElectronicDeliveryConsentResponse {
	this := CheckTaxElectronicDeliveryConsentResponse{}
	return &this
}

// GetUS1099k returns the US1099k field value if set, zero value otherwise.
func (o *CheckTaxElectronicDeliveryConsentResponse) GetUS1099k() bool {
	if o == nil || common.IsNil(o.US1099k) {
		var ret bool
		return ret
	}
	return *o.US1099k
}

// GetUS1099kOk returns a tuple with the US1099k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckTaxElectronicDeliveryConsentResponse) GetUS1099kOk() (*bool, bool) {
	if o == nil || common.IsNil(o.US1099k) {
		return nil, false
	}
	return o.US1099k, true
}

// HasUS1099k returns a boolean if a field has been set.
func (o *CheckTaxElectronicDeliveryConsentResponse) HasUS1099k() bool {
	if o != nil && !common.IsNil(o.US1099k) {
		return true
	}

	return false
}

// SetUS1099k gets a reference to the given bool and assigns it to the US1099k field.
func (o *CheckTaxElectronicDeliveryConsentResponse) SetUS1099k(v bool) {
	o.US1099k = &v
}

func (o CheckTaxElectronicDeliveryConsentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckTaxElectronicDeliveryConsentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.US1099k) {
		toSerialize["US1099k"] = o.US1099k
	}
	return toSerialize, nil
}

type NullableCheckTaxElectronicDeliveryConsentResponse struct {
	value *CheckTaxElectronicDeliveryConsentResponse
	isSet bool
}

func (v NullableCheckTaxElectronicDeliveryConsentResponse) Get() *CheckTaxElectronicDeliveryConsentResponse {
	return v.value
}

func (v *NullableCheckTaxElectronicDeliveryConsentResponse) Set(val *CheckTaxElectronicDeliveryConsentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckTaxElectronicDeliveryConsentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckTaxElectronicDeliveryConsentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckTaxElectronicDeliveryConsentResponse(val *CheckTaxElectronicDeliveryConsentResponse) *NullableCheckTaxElectronicDeliveryConsentResponse {
	return &NullableCheckTaxElectronicDeliveryConsentResponse{value: val, isSet: true}
}

func (v NullableCheckTaxElectronicDeliveryConsentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckTaxElectronicDeliveryConsentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
