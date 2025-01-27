/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the GenerateApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GenerateApiKeyResponse{}

// GenerateApiKeyResponse struct for GenerateApiKeyResponse
type GenerateApiKeyResponse struct {
	// The generated API key.
	ApiKey string `json:"apiKey"`
}

// NewGenerateApiKeyResponse instantiates a new GenerateApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateApiKeyResponse(apiKey string) *GenerateApiKeyResponse {
	this := GenerateApiKeyResponse{}
	this.ApiKey = apiKey
	return &this
}

// NewGenerateApiKeyResponseWithDefaults instantiates a new GenerateApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateApiKeyResponseWithDefaults() *GenerateApiKeyResponse {
	this := GenerateApiKeyResponse{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *GenerateApiKeyResponse) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *GenerateApiKeyResponse) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *GenerateApiKeyResponse) SetApiKey(v string) {
	o.ApiKey = v
}

func (o GenerateApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey
	return toSerialize, nil
}

type NullableGenerateApiKeyResponse struct {
	value *GenerateApiKeyResponse
	isSet bool
}

func (v NullableGenerateApiKeyResponse) Get() *GenerateApiKeyResponse {
	return v.value
}

func (v *NullableGenerateApiKeyResponse) Set(val *GenerateApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateApiKeyResponse(val *GenerateApiKeyResponse) *NullableGenerateApiKeyResponse {
	return &NullableGenerateApiKeyResponse{value: val, isSet: true}
}

func (v NullableGenerateApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
