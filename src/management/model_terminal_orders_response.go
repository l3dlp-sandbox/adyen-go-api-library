/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the TerminalOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalOrdersResponse{}

// TerminalOrdersResponse struct for TerminalOrdersResponse
type TerminalOrdersResponse struct {
	// List of orders for payment terminal packages and parts.
	Data []TerminalOrder `json:"data,omitempty"`
}

// NewTerminalOrdersResponse instantiates a new TerminalOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalOrdersResponse() *TerminalOrdersResponse {
	this := TerminalOrdersResponse{}
	return &this
}

// NewTerminalOrdersResponseWithDefaults instantiates a new TerminalOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalOrdersResponseWithDefaults() *TerminalOrdersResponse {
	this := TerminalOrdersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TerminalOrdersResponse) GetData() []TerminalOrder {
	if o == nil || common.IsNil(o.Data) {
		var ret []TerminalOrder
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrdersResponse) GetDataOk() ([]TerminalOrder, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TerminalOrdersResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TerminalOrder and assigns it to the Data field.
func (o *TerminalOrdersResponse) SetData(v []TerminalOrder) {
	o.Data = v
}

func (o TerminalOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTerminalOrdersResponse struct {
	value *TerminalOrdersResponse
	isSet bool
}

func (v NullableTerminalOrdersResponse) Get() *TerminalOrdersResponse {
	return v.value
}

func (v *NullableTerminalOrdersResponse) Set(val *TerminalOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalOrdersResponse(val *TerminalOrdersResponse) *NullableTerminalOrdersResponse {
	return &NullableTerminalOrdersResponse{value: val, isSet: true}
}

func (v NullableTerminalOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
