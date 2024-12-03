/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the ModifyRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModifyRequest{}

// ModifyRequest struct for ModifyRequest
type ModifyRequest struct {
	// This field contains additional data, which may be required for a particular payout request.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The PSP reference received in the `/submitThirdParty` response.
	OriginalReference string `json:"originalReference"`
}

// NewModifyRequest instantiates a new ModifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyRequest(merchantAccount string, originalReference string) *ModifyRequest {
	this := ModifyRequest{}
	this.MerchantAccount = merchantAccount
	this.OriginalReference = originalReference
	return &this
}

// NewModifyRequestWithDefaults instantiates a new ModifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyRequestWithDefaults() *ModifyRequest {
	this := ModifyRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ModifyRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ModifyRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ModifyRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *ModifyRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *ModifyRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *ModifyRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetOriginalReference returns the OriginalReference field value
func (o *ModifyRequest) GetOriginalReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalReference
}

// GetOriginalReferenceOk returns a tuple with the OriginalReference field value
// and a boolean to check if the value has been set.
func (o *ModifyRequest) GetOriginalReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalReference, true
}

// SetOriginalReference sets field value
func (o *ModifyRequest) SetOriginalReference(v string) {
	o.OriginalReference = v
}

func (o ModifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["originalReference"] = o.OriginalReference
	return toSerialize, nil
}

type NullableModifyRequest struct {
	value *ModifyRequest
	isSet bool
}

func (v NullableModifyRequest) Get() *ModifyRequest {
	return v.value
}

func (v *NullableModifyRequest) Set(val *ModifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyRequest(val *ModifyRequest) *NullableModifyRequest {
	return &NullableModifyRequest{value: val, isSet: true}
}

func (v NullableModifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



