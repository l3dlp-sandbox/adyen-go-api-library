/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the TransferReview type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferReview{}

// TransferReview struct for TransferReview
type TransferReview struct {
	// Shows the number of [approvals](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers/approve) required to process the transfer.
	NumberOfApprovalsRequired *int32 `json:"numberOfApprovalsRequired,omitempty"`
}

// NewTransferReview instantiates a new TransferReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferReview() *TransferReview {
	this := TransferReview{}
	return &this
}

// NewTransferReviewWithDefaults instantiates a new TransferReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferReviewWithDefaults() *TransferReview {
	this := TransferReview{}
	return &this
}

// GetNumberOfApprovalsRequired returns the NumberOfApprovalsRequired field value if set, zero value otherwise.
func (o *TransferReview) GetNumberOfApprovalsRequired() int32 {
	if o == nil || common.IsNil(o.NumberOfApprovalsRequired) {
		var ret int32
		return ret
	}
	return *o.NumberOfApprovalsRequired
}

// GetNumberOfApprovalsRequiredOk returns a tuple with the NumberOfApprovalsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferReview) GetNumberOfApprovalsRequiredOk() (*int32, bool) {
	if o == nil || common.IsNil(o.NumberOfApprovalsRequired) {
		return nil, false
	}
	return o.NumberOfApprovalsRequired, true
}

// HasNumberOfApprovalsRequired returns a boolean if a field has been set.
func (o *TransferReview) HasNumberOfApprovalsRequired() bool {
	if o != nil && !common.IsNil(o.NumberOfApprovalsRequired) {
		return true
	}

	return false
}

// SetNumberOfApprovalsRequired gets a reference to the given int32 and assigns it to the NumberOfApprovalsRequired field.
func (o *TransferReview) SetNumberOfApprovalsRequired(v int32) {
	o.NumberOfApprovalsRequired = &v
}

func (o TransferReview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NumberOfApprovalsRequired) {
		toSerialize["numberOfApprovalsRequired"] = o.NumberOfApprovalsRequired
	}
	return toSerialize, nil
}

type NullableTransferReview struct {
	value *TransferReview
	isSet bool
}

func (v NullableTransferReview) Get() *TransferReview {
	return v.value
}

func (v *NullableTransferReview) Set(val *TransferReview) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferReview) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferReview(val *TransferReview) *NullableTransferReview {
	return &NullableTransferReview{value: val, isSet: true}
}

func (v NullableTransferReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
