/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the TransferRequestReview type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferRequestReview{}

// TransferRequestReview struct for TransferRequestReview
type TransferRequestReview struct {
	// Specifies the number of [approvals](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers/approve) required to process the transfer.
	NumberOfApprovalsRequired *int32 `json:"numberOfApprovalsRequired,omitempty"`
	// Specifies whether you will initiate Strong Customer Authentication (SCA) in thePOST [/transfers/approve](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers/approve) request.  Only applies to transfers made with an Adyen [business account](https://docs.adyen.com/platforms/business-accounts).
	ScaOnApproval *bool `json:"scaOnApproval,omitempty"`
}

// NewTransferRequestReview instantiates a new TransferRequestReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRequestReview() *TransferRequestReview {
	this := TransferRequestReview{}
	return &this
}

// NewTransferRequestReviewWithDefaults instantiates a new TransferRequestReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRequestReviewWithDefaults() *TransferRequestReview {
	this := TransferRequestReview{}
	return &this
}

// GetNumberOfApprovalsRequired returns the NumberOfApprovalsRequired field value if set, zero value otherwise.
func (o *TransferRequestReview) GetNumberOfApprovalsRequired() int32 {
	if o == nil || common.IsNil(o.NumberOfApprovalsRequired) {
		var ret int32
		return ret
	}
	return *o.NumberOfApprovalsRequired
}

// GetNumberOfApprovalsRequiredOk returns a tuple with the NumberOfApprovalsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequestReview) GetNumberOfApprovalsRequiredOk() (*int32, bool) {
	if o == nil || common.IsNil(o.NumberOfApprovalsRequired) {
		return nil, false
	}
	return o.NumberOfApprovalsRequired, true
}

// HasNumberOfApprovalsRequired returns a boolean if a field has been set.
func (o *TransferRequestReview) HasNumberOfApprovalsRequired() bool {
	if o != nil && !common.IsNil(o.NumberOfApprovalsRequired) {
		return true
	}

	return false
}

// SetNumberOfApprovalsRequired gets a reference to the given int32 and assigns it to the NumberOfApprovalsRequired field.
func (o *TransferRequestReview) SetNumberOfApprovalsRequired(v int32) {
	o.NumberOfApprovalsRequired = &v
}

// GetScaOnApproval returns the ScaOnApproval field value if set, zero value otherwise.
func (o *TransferRequestReview) GetScaOnApproval() bool {
	if o == nil || common.IsNil(o.ScaOnApproval) {
		var ret bool
		return ret
	}
	return *o.ScaOnApproval
}

// GetScaOnApprovalOk returns a tuple with the ScaOnApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequestReview) GetScaOnApprovalOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ScaOnApproval) {
		return nil, false
	}
	return o.ScaOnApproval, true
}

// HasScaOnApproval returns a boolean if a field has been set.
func (o *TransferRequestReview) HasScaOnApproval() bool {
	if o != nil && !common.IsNil(o.ScaOnApproval) {
		return true
	}

	return false
}

// SetScaOnApproval gets a reference to the given bool and assigns it to the ScaOnApproval field.
func (o *TransferRequestReview) SetScaOnApproval(v bool) {
	o.ScaOnApproval = &v
}

func (o TransferRequestReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRequestReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NumberOfApprovalsRequired) {
		toSerialize["numberOfApprovalsRequired"] = o.NumberOfApprovalsRequired
	}
	if !common.IsNil(o.ScaOnApproval) {
		toSerialize["scaOnApproval"] = o.ScaOnApproval
	}
	return toSerialize, nil
}

type NullableTransferRequestReview struct {
	value *TransferRequestReview
	isSet bool
}

func (v NullableTransferRequestReview) Get() *TransferRequestReview {
	return v.value
}

func (v *NullableTransferRequestReview) Set(val *TransferRequestReview) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRequestReview) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRequestReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRequestReview(val *TransferRequestReview) *NullableTransferRequestReview {
	return &NullableTransferRequestReview{value: val, isSet: true}
}

func (v NullableTransferRequestReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRequestReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



