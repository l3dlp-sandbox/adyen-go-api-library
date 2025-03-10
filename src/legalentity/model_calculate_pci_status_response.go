/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the CalculatePciStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CalculatePciStatusResponse{}

// CalculatePciStatusResponse struct for CalculatePciStatusResponse
type CalculatePciStatusResponse struct {
	// Indicates if the user is required to sign PCI questionnaires. If **false**, they do not need to sign any questionnaires.
	SigningRequired *bool `json:"signingRequired,omitempty"`
}

// NewCalculatePciStatusResponse instantiates a new CalculatePciStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalculatePciStatusResponse() *CalculatePciStatusResponse {
	this := CalculatePciStatusResponse{}
	return &this
}

// NewCalculatePciStatusResponseWithDefaults instantiates a new CalculatePciStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculatePciStatusResponseWithDefaults() *CalculatePciStatusResponse {
	this := CalculatePciStatusResponse{}
	return &this
}

// GetSigningRequired returns the SigningRequired field value if set, zero value otherwise.
func (o *CalculatePciStatusResponse) GetSigningRequired() bool {
	if o == nil || common.IsNil(o.SigningRequired) {
		var ret bool
		return ret
	}
	return *o.SigningRequired
}

// GetSigningRequiredOk returns a tuple with the SigningRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculatePciStatusResponse) GetSigningRequiredOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SigningRequired) {
		return nil, false
	}
	return o.SigningRequired, true
}

// HasSigningRequired returns a boolean if a field has been set.
func (o *CalculatePciStatusResponse) HasSigningRequired() bool {
	if o != nil && !common.IsNil(o.SigningRequired) {
		return true
	}

	return false
}

// SetSigningRequired gets a reference to the given bool and assigns it to the SigningRequired field.
func (o *CalculatePciStatusResponse) SetSigningRequired(v bool) {
	o.SigningRequired = &v
}

func (o CalculatePciStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalculatePciStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SigningRequired) {
		toSerialize["signingRequired"] = o.SigningRequired
	}
	return toSerialize, nil
}

type NullableCalculatePciStatusResponse struct {
	value *CalculatePciStatusResponse
	isSet bool
}

func (v NullableCalculatePciStatusResponse) Get() *CalculatePciStatusResponse {
	return v.value
}

func (v *NullableCalculatePciStatusResponse) Set(val *CalculatePciStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculatePciStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculatePciStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculatePciStatusResponse(val *CalculatePciStatusResponse) *NullableCalculatePciStatusResponse {
	return &NullableCalculatePciStatusResponse{value: val, isSet: true}
}

func (v NullableCalculatePciStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculatePciStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
