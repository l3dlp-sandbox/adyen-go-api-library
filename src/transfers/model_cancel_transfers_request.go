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

// checks if the CancelTransfersRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelTransfersRequest{}

// CancelTransfersRequest struct for CancelTransfersRequest
type CancelTransfersRequest struct {
	// Contains the unique identifiers of the transfers that you want to cancel.
	TransferIds []string `json:"transferIds,omitempty"`
}

// NewCancelTransfersRequest instantiates a new CancelTransfersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelTransfersRequest() *CancelTransfersRequest {
	this := CancelTransfersRequest{}
	return &this
}

// NewCancelTransfersRequestWithDefaults instantiates a new CancelTransfersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelTransfersRequestWithDefaults() *CancelTransfersRequest {
	this := CancelTransfersRequest{}
	return &this
}

// GetTransferIds returns the TransferIds field value if set, zero value otherwise.
func (o *CancelTransfersRequest) GetTransferIds() []string {
	if o == nil || common.IsNil(o.TransferIds) {
		var ret []string
		return ret
	}
	return o.TransferIds
}

// GetTransferIdsOk returns a tuple with the TransferIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelTransfersRequest) GetTransferIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.TransferIds) {
		return nil, false
	}
	return o.TransferIds, true
}

// HasTransferIds returns a boolean if a field has been set.
func (o *CancelTransfersRequest) HasTransferIds() bool {
	if o != nil && !common.IsNil(o.TransferIds) {
		return true
	}

	return false
}

// SetTransferIds gets a reference to the given []string and assigns it to the TransferIds field.
func (o *CancelTransfersRequest) SetTransferIds(v []string) {
	o.TransferIds = v
}

func (o CancelTransfersRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelTransfersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransferIds) {
		toSerialize["transferIds"] = o.TransferIds
	}
	return toSerialize, nil
}

type NullableCancelTransfersRequest struct {
	value *CancelTransfersRequest
	isSet bool
}

func (v NullableCancelTransfersRequest) Get() *CancelTransfersRequest {
	return v.value
}

func (v *NullableCancelTransfersRequest) Set(val *CancelTransfersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelTransfersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelTransfersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelTransfersRequest(val *CancelTransfersRequest) *NullableCancelTransfersRequest {
	return &NullableCancelTransfersRequest{value: val, isSet: true}
}

func (v NullableCancelTransfersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelTransfersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
