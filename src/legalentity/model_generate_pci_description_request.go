/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the GeneratePciDescriptionRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GeneratePciDescriptionRequest{}

// GeneratePciDescriptionRequest struct for GeneratePciDescriptionRequest
type GeneratePciDescriptionRequest struct {
	// An array of additional sales channels to generate PCI questionnaires. Include the relevant sales channels if you need your user to sign PCI questionnaires. Not required if you [create stores](https://docs.adyen.com/platforms) and [add payment methods](https://docs.adyen.com/adyen-for-platforms-model) before you generate the questionnaires.  Possible values: *  **eCommerce** *  **pos** *  **ecomMoto** *  **posMoto**
	AdditionalSalesChannels []string `json:"additionalSalesChannels,omitempty"`
	// Sets the language of the PCI questionnaire. Its value is a two-character [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code, for example, **en**.
	Language *string `json:"language,omitempty"`
}

// NewGeneratePciDescriptionRequest instantiates a new GeneratePciDescriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePciDescriptionRequest() *GeneratePciDescriptionRequest {
	this := GeneratePciDescriptionRequest{}
	return &this
}

// NewGeneratePciDescriptionRequestWithDefaults instantiates a new GeneratePciDescriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePciDescriptionRequestWithDefaults() *GeneratePciDescriptionRequest {
	this := GeneratePciDescriptionRequest{}
	return &this
}

// GetAdditionalSalesChannels returns the AdditionalSalesChannels field value if set, zero value otherwise.
func (o *GeneratePciDescriptionRequest) GetAdditionalSalesChannels() []string {
	if o == nil || common.IsNil(o.AdditionalSalesChannels) {
		var ret []string
		return ret
	}
	return o.AdditionalSalesChannels
}

// GetAdditionalSalesChannelsOk returns a tuple with the AdditionalSalesChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePciDescriptionRequest) GetAdditionalSalesChannelsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AdditionalSalesChannels) {
		return nil, false
	}
	return o.AdditionalSalesChannels, true
}

// HasAdditionalSalesChannels returns a boolean if a field has been set.
func (o *GeneratePciDescriptionRequest) HasAdditionalSalesChannels() bool {
	if o != nil && !common.IsNil(o.AdditionalSalesChannels) {
		return true
	}

	return false
}

// SetAdditionalSalesChannels gets a reference to the given []string and assigns it to the AdditionalSalesChannels field.
func (o *GeneratePciDescriptionRequest) SetAdditionalSalesChannels(v []string) {
	o.AdditionalSalesChannels = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GeneratePciDescriptionRequest) GetLanguage() string {
	if o == nil || common.IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePciDescriptionRequest) GetLanguageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GeneratePciDescriptionRequest) HasLanguage() bool {
	if o != nil && !common.IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GeneratePciDescriptionRequest) SetLanguage(v string) {
	o.Language = &v
}

func (o GeneratePciDescriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneratePciDescriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalSalesChannels) {
		toSerialize["additionalSalesChannels"] = o.AdditionalSalesChannels
	}
	if !common.IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableGeneratePciDescriptionRequest struct {
	value *GeneratePciDescriptionRequest
	isSet bool
}

func (v NullableGeneratePciDescriptionRequest) Get() *GeneratePciDescriptionRequest {
	return v.value
}

func (v *NullableGeneratePciDescriptionRequest) Set(val *GeneratePciDescriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePciDescriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePciDescriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePciDescriptionRequest(val *GeneratePciDescriptionRequest) *NullableGeneratePciDescriptionRequest {
	return &NullableGeneratePciDescriptionRequest{value: val, isSet: true}
}

func (v NullableGeneratePciDescriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratePciDescriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
