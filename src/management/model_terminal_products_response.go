/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the TerminalProductsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalProductsResponse{}

// TerminalProductsResponse struct for TerminalProductsResponse
type TerminalProductsResponse struct {
	// Terminal products that can be ordered.
	Data []TerminalProduct `json:"data,omitempty"`
}

// NewTerminalProductsResponse instantiates a new TerminalProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalProductsResponse() *TerminalProductsResponse {
	this := TerminalProductsResponse{}
	return &this
}

// NewTerminalProductsResponseWithDefaults instantiates a new TerminalProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalProductsResponseWithDefaults() *TerminalProductsResponse {
	this := TerminalProductsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TerminalProductsResponse) GetData() []TerminalProduct {
	if o == nil || common.IsNil(o.Data) {
		var ret []TerminalProduct
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProductsResponse) GetDataOk() ([]TerminalProduct, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TerminalProductsResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TerminalProduct and assigns it to the Data field.
func (o *TerminalProductsResponse) SetData(v []TerminalProduct) {
	o.Data = v
}

func (o TerminalProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTerminalProductsResponse struct {
	value *TerminalProductsResponse
	isSet bool
}

func (v NullableTerminalProductsResponse) Get() *TerminalProductsResponse {
	return v.value
}

func (v *NullableTerminalProductsResponse) Set(val *TerminalProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalProductsResponse(val *TerminalProductsResponse) *NullableTerminalProductsResponse {
	return &NullableTerminalProductsResponse{value: val, isSet: true}
}

func (v NullableTerminalProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
