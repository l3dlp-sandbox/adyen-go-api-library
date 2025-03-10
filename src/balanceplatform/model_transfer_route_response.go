/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the TransferRouteResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferRouteResponse{}

// TransferRouteResponse struct for TransferRouteResponse
type TransferRouteResponse struct {
	// List of available priorities for a transfer, along with requirements. Use this information to initiate a transfer.
	TransferRoutes []TransferRoute `json:"transferRoutes,omitempty"`
}

// NewTransferRouteResponse instantiates a new TransferRouteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRouteResponse() *TransferRouteResponse {
	this := TransferRouteResponse{}
	return &this
}

// NewTransferRouteResponseWithDefaults instantiates a new TransferRouteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRouteResponseWithDefaults() *TransferRouteResponse {
	this := TransferRouteResponse{}
	return &this
}

// GetTransferRoutes returns the TransferRoutes field value if set, zero value otherwise.
func (o *TransferRouteResponse) GetTransferRoutes() []TransferRoute {
	if o == nil || common.IsNil(o.TransferRoutes) {
		var ret []TransferRoute
		return ret
	}
	return o.TransferRoutes
}

// GetTransferRoutesOk returns a tuple with the TransferRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRouteResponse) GetTransferRoutesOk() ([]TransferRoute, bool) {
	if o == nil || common.IsNil(o.TransferRoutes) {
		return nil, false
	}
	return o.TransferRoutes, true
}

// HasTransferRoutes returns a boolean if a field has been set.
func (o *TransferRouteResponse) HasTransferRoutes() bool {
	if o != nil && !common.IsNil(o.TransferRoutes) {
		return true
	}

	return false
}

// SetTransferRoutes gets a reference to the given []TransferRoute and assigns it to the TransferRoutes field.
func (o *TransferRouteResponse) SetTransferRoutes(v []TransferRoute) {
	o.TransferRoutes = v
}

func (o TransferRouteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRouteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransferRoutes) {
		toSerialize["transferRoutes"] = o.TransferRoutes
	}
	return toSerialize, nil
}

type NullableTransferRouteResponse struct {
	value *TransferRouteResponse
	isSet bool
}

func (v NullableTransferRouteResponse) Get() *TransferRouteResponse {
	return v.value
}

func (v *NullableTransferRouteResponse) Set(val *TransferRouteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRouteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRouteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRouteResponse(val *TransferRouteResponse) *NullableTransferRouteResponse {
	return &NullableTransferRouteResponse{value: val, isSet: true}
}

func (v NullableTransferRouteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRouteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
