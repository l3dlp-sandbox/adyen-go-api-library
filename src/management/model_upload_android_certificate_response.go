/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the UploadAndroidCertificateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UploadAndroidCertificateResponse{}

// UploadAndroidCertificateResponse struct for UploadAndroidCertificateResponse
type UploadAndroidCertificateResponse struct {
	// The unique identifier of the uploaded Android certificate.
	Id *string `json:"id,omitempty"`
}

// NewUploadAndroidCertificateResponse instantiates a new UploadAndroidCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadAndroidCertificateResponse() *UploadAndroidCertificateResponse {
	this := UploadAndroidCertificateResponse{}
	return &this
}

// NewUploadAndroidCertificateResponseWithDefaults instantiates a new UploadAndroidCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadAndroidCertificateResponseWithDefaults() *UploadAndroidCertificateResponse {
	this := UploadAndroidCertificateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UploadAndroidCertificateResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadAndroidCertificateResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UploadAndroidCertificateResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UploadAndroidCertificateResponse) SetId(v string) {
	o.Id = &v
}

func (o UploadAndroidCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadAndroidCertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUploadAndroidCertificateResponse struct {
	value *UploadAndroidCertificateResponse
	isSet bool
}

func (v NullableUploadAndroidCertificateResponse) Get() *UploadAndroidCertificateResponse {
	return v.value
}

func (v *NullableUploadAndroidCertificateResponse) Set(val *UploadAndroidCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadAndroidCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadAndroidCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadAndroidCertificateResponse(val *UploadAndroidCertificateResponse) *NullableUploadAndroidCertificateResponse {
	return &NullableUploadAndroidCertificateResponse{value: val, isSet: true}
}

func (v NullableUploadAndroidCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadAndroidCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



