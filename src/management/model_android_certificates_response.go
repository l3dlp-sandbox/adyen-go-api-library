/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the AndroidCertificatesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AndroidCertificatesResponse{}

// AndroidCertificatesResponse struct for AndroidCertificatesResponse
type AndroidCertificatesResponse struct {
	// Uploaded Android certificates for Android payment terminals.
	Data []AndroidCertificate `json:"data,omitempty"`
}

// NewAndroidCertificatesResponse instantiates a new AndroidCertificatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndroidCertificatesResponse() *AndroidCertificatesResponse {
	this := AndroidCertificatesResponse{}
	return &this
}

// NewAndroidCertificatesResponseWithDefaults instantiates a new AndroidCertificatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndroidCertificatesResponseWithDefaults() *AndroidCertificatesResponse {
	this := AndroidCertificatesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AndroidCertificatesResponse) GetData() []AndroidCertificate {
	if o == nil || common.IsNil(o.Data) {
		var ret []AndroidCertificate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidCertificatesResponse) GetDataOk() ([]AndroidCertificate, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AndroidCertificatesResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AndroidCertificate and assigns it to the Data field.
func (o *AndroidCertificatesResponse) SetData(v []AndroidCertificate) {
	o.Data = v
}

func (o AndroidCertificatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AndroidCertificatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAndroidCertificatesResponse struct {
	value *AndroidCertificatesResponse
	isSet bool
}

func (v NullableAndroidCertificatesResponse) Get() *AndroidCertificatesResponse {
	return v.value
}

func (v *NullableAndroidCertificatesResponse) Set(val *AndroidCertificatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAndroidCertificatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAndroidCertificatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndroidCertificatesResponse(val *AndroidCertificatesResponse) *NullableAndroidCertificatesResponse {
	return &NullableAndroidCertificatesResponse{value: val, isSet: true}
}

func (v NullableAndroidCertificatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndroidCertificatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}