/*
Adyen Data Protection API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dataprotection

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the SubjectErasureByPspReferenceRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubjectErasureByPspReferenceRequest{}

// SubjectErasureByPspReferenceRequest struct for SubjectErasureByPspReferenceRequest
type SubjectErasureByPspReferenceRequest struct {
	// Set this to **true** if you want to delete shopper-related data, even if the shopper has an existing recurring transaction. This only deletes the shopper-related data for the specific payment, but does not cancel the existing recurring transaction.
	ForceErasure *bool `json:"forceErasure,omitempty"`
	// Your merchant account
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// The PSP reference of the payment. We will delete all shopper-related data for this payment.
	PspReference *string `json:"pspReference,omitempty"`
}

// NewSubjectErasureByPspReferenceRequest instantiates a new SubjectErasureByPspReferenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectErasureByPspReferenceRequest() *SubjectErasureByPspReferenceRequest {
	this := SubjectErasureByPspReferenceRequest{}
	return &this
}

// NewSubjectErasureByPspReferenceRequestWithDefaults instantiates a new SubjectErasureByPspReferenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectErasureByPspReferenceRequestWithDefaults() *SubjectErasureByPspReferenceRequest {
	this := SubjectErasureByPspReferenceRequest{}
	return &this
}

// GetForceErasure returns the ForceErasure field value if set, zero value otherwise.
func (o *SubjectErasureByPspReferenceRequest) GetForceErasure() bool {
	if o == nil || common.IsNil(o.ForceErasure) {
		var ret bool
		return ret
	}
	return *o.ForceErasure
}

// GetForceErasureOk returns a tuple with the ForceErasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectErasureByPspReferenceRequest) GetForceErasureOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ForceErasure) {
		return nil, false
	}
	return o.ForceErasure, true
}

// HasForceErasure returns a boolean if a field has been set.
func (o *SubjectErasureByPspReferenceRequest) HasForceErasure() bool {
	if o != nil && !common.IsNil(o.ForceErasure) {
		return true
	}

	return false
}

// SetForceErasure gets a reference to the given bool and assigns it to the ForceErasure field.
func (o *SubjectErasureByPspReferenceRequest) SetForceErasure(v bool) {
	o.ForceErasure = &v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *SubjectErasureByPspReferenceRequest) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectErasureByPspReferenceRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *SubjectErasureByPspReferenceRequest) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *SubjectErasureByPspReferenceRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *SubjectErasureByPspReferenceRequest) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectErasureByPspReferenceRequest) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *SubjectErasureByPspReferenceRequest) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *SubjectErasureByPspReferenceRequest) SetPspReference(v string) {
	o.PspReference = &v
}

func (o SubjectErasureByPspReferenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectErasureByPspReferenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ForceErasure) {
		toSerialize["forceErasure"] = o.ForceErasure
	}
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	return toSerialize, nil
}

type NullableSubjectErasureByPspReferenceRequest struct {
	value *SubjectErasureByPspReferenceRequest
	isSet bool
}

func (v NullableSubjectErasureByPspReferenceRequest) Get() *SubjectErasureByPspReferenceRequest {
	return v.value
}

func (v *NullableSubjectErasureByPspReferenceRequest) Set(val *SubjectErasureByPspReferenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectErasureByPspReferenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectErasureByPspReferenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectErasureByPspReferenceRequest(val *SubjectErasureByPspReferenceRequest) *NullableSubjectErasureByPspReferenceRequest {
	return &NullableSubjectErasureByPspReferenceRequest{value: val, isSet: true}
}

func (v NullableSubjectErasureByPspReferenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectErasureByPspReferenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
