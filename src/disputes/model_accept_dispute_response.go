/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the AcceptDisputeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcceptDisputeResponse{}

// AcceptDisputeResponse struct for AcceptDisputeResponse
type AcceptDisputeResponse struct {
	DisputeServiceResult DisputeServiceResult `json:"disputeServiceResult"`
}

// NewAcceptDisputeResponse instantiates a new AcceptDisputeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptDisputeResponse(disputeServiceResult DisputeServiceResult) *AcceptDisputeResponse {
	this := AcceptDisputeResponse{}
	this.DisputeServiceResult = disputeServiceResult
	return &this
}

// NewAcceptDisputeResponseWithDefaults instantiates a new AcceptDisputeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptDisputeResponseWithDefaults() *AcceptDisputeResponse {
	this := AcceptDisputeResponse{}
	return &this
}

// GetDisputeServiceResult returns the DisputeServiceResult field value
func (o *AcceptDisputeResponse) GetDisputeServiceResult() DisputeServiceResult {
	if o == nil {
		var ret DisputeServiceResult
		return ret
	}

	return o.DisputeServiceResult
}

// GetDisputeServiceResultOk returns a tuple with the DisputeServiceResult field value
// and a boolean to check if the value has been set.
func (o *AcceptDisputeResponse) GetDisputeServiceResultOk() (*DisputeServiceResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputeServiceResult, true
}

// SetDisputeServiceResult sets field value
func (o *AcceptDisputeResponse) SetDisputeServiceResult(v DisputeServiceResult) {
	o.DisputeServiceResult = v
}

func (o AcceptDisputeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptDisputeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disputeServiceResult"] = o.DisputeServiceResult
	return toSerialize, nil
}

type NullableAcceptDisputeResponse struct {
	value *AcceptDisputeResponse
	isSet bool
}

func (v NullableAcceptDisputeResponse) Get() *AcceptDisputeResponse {
	return v.value
}

func (v *NullableAcceptDisputeResponse) Set(val *AcceptDisputeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptDisputeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptDisputeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptDisputeResponse(val *AcceptDisputeResponse) *NullableAcceptDisputeResponse {
	return &NullableAcceptDisputeResponse{value: val, isSet: true}
}

func (v NullableAcceptDisputeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptDisputeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



