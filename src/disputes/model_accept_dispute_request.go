/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the AcceptDisputeRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptDisputeRequest{}

// AcceptDisputeRequest struct for AcceptDisputeRequest
type AcceptDisputeRequest struct {
	// The PSP reference assigned to the dispute.
	DisputePspReference string `json:"disputePspReference"`
	// The merchant account identifier, for which you want to process the dispute transaction.
	MerchantAccountCode string `json:"merchantAccountCode"`
}

// NewAcceptDisputeRequest instantiates a new AcceptDisputeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptDisputeRequest(disputePspReference string, merchantAccountCode string) *AcceptDisputeRequest {
	this := AcceptDisputeRequest{}
	this.DisputePspReference = disputePspReference
	this.MerchantAccountCode = merchantAccountCode
	return &this
}

// NewAcceptDisputeRequestWithDefaults instantiates a new AcceptDisputeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptDisputeRequestWithDefaults() *AcceptDisputeRequest {
	this := AcceptDisputeRequest{}
	return &this
}

// GetDisputePspReference returns the DisputePspReference field value
func (o *AcceptDisputeRequest) GetDisputePspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisputePspReference
}

// GetDisputePspReferenceOk returns a tuple with the DisputePspReference field value
// and a boolean to check if the value has been set.
func (o *AcceptDisputeRequest) GetDisputePspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputePspReference, true
}

// SetDisputePspReference sets field value
func (o *AcceptDisputeRequest) SetDisputePspReference(v string) {
	o.DisputePspReference = v
}

// GetMerchantAccountCode returns the MerchantAccountCode field value
func (o *AcceptDisputeRequest) GetMerchantAccountCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountCode
}

// GetMerchantAccountCodeOk returns a tuple with the MerchantAccountCode field value
// and a boolean to check if the value has been set.
func (o *AcceptDisputeRequest) GetMerchantAccountCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountCode, true
}

// SetMerchantAccountCode sets field value
func (o *AcceptDisputeRequest) SetMerchantAccountCode(v string) {
	o.MerchantAccountCode = v
}

func (o AcceptDisputeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptDisputeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disputePspReference"] = o.DisputePspReference
	toSerialize["merchantAccountCode"] = o.MerchantAccountCode
	return toSerialize, nil
}

type NullableAcceptDisputeRequest struct {
	value *AcceptDisputeRequest
	isSet bool
}

func (v NullableAcceptDisputeRequest) Get() *AcceptDisputeRequest {
	return v.value
}

func (v *NullableAcceptDisputeRequest) Set(val *AcceptDisputeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptDisputeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptDisputeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptDisputeRequest(val *AcceptDisputeRequest) *NullableAcceptDisputeRequest {
	return &NullableAcceptDisputeRequest{value: val, isSet: true}
}

func (v NullableAcceptDisputeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptDisputeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
