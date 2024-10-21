/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the GetTermsOfServiceAcceptanceInfosResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTermsOfServiceAcceptanceInfosResponse{}

// GetTermsOfServiceAcceptanceInfosResponse struct for GetTermsOfServiceAcceptanceInfosResponse
type GetTermsOfServiceAcceptanceInfosResponse struct {
	// The Terms of Service acceptance information.
	Data []TermsOfServiceAcceptanceInfo `json:"data,omitempty"`
}

// NewGetTermsOfServiceAcceptanceInfosResponse instantiates a new GetTermsOfServiceAcceptanceInfosResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTermsOfServiceAcceptanceInfosResponse() *GetTermsOfServiceAcceptanceInfosResponse {
	this := GetTermsOfServiceAcceptanceInfosResponse{}
	return &this
}

// NewGetTermsOfServiceAcceptanceInfosResponseWithDefaults instantiates a new GetTermsOfServiceAcceptanceInfosResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTermsOfServiceAcceptanceInfosResponseWithDefaults() *GetTermsOfServiceAcceptanceInfosResponse {
	this := GetTermsOfServiceAcceptanceInfosResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTermsOfServiceAcceptanceInfosResponse) GetData() []TermsOfServiceAcceptanceInfo {
	if o == nil || common.IsNil(o.Data) {
		var ret []TermsOfServiceAcceptanceInfo
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceAcceptanceInfosResponse) GetDataOk() ([]TermsOfServiceAcceptanceInfo, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTermsOfServiceAcceptanceInfosResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TermsOfServiceAcceptanceInfo and assigns it to the Data field.
func (o *GetTermsOfServiceAcceptanceInfosResponse) SetData(v []TermsOfServiceAcceptanceInfo) {
	o.Data = v
}

func (o GetTermsOfServiceAcceptanceInfosResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTermsOfServiceAcceptanceInfosResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetTermsOfServiceAcceptanceInfosResponse struct {
	value *GetTermsOfServiceAcceptanceInfosResponse
	isSet bool
}

func (v NullableGetTermsOfServiceAcceptanceInfosResponse) Get() *GetTermsOfServiceAcceptanceInfosResponse {
	return v.value
}

func (v *NullableGetTermsOfServiceAcceptanceInfosResponse) Set(val *GetTermsOfServiceAcceptanceInfosResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTermsOfServiceAcceptanceInfosResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTermsOfServiceAcceptanceInfosResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTermsOfServiceAcceptanceInfosResponse(val *GetTermsOfServiceAcceptanceInfosResponse) *NullableGetTermsOfServiceAcceptanceInfosResponse {
	return &NullableGetTermsOfServiceAcceptanceInfosResponse{value: val, isSet: true}
}

func (v NullableGetTermsOfServiceAcceptanceInfosResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTermsOfServiceAcceptanceInfosResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
