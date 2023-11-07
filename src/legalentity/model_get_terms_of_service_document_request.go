/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the GetTermsOfServiceDocumentRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTermsOfServiceDocumentRequest{}

// GetTermsOfServiceDocumentRequest struct for GetTermsOfServiceDocumentRequest
type GetTermsOfServiceDocumentRequest struct {
	// The language to be used for the Terms of Service document, specified by the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. Possible value: **en** for English.
	Language string `json:"language"`
	// The type of Terms of Service.  Possible values: *  **adyenForPlatformsManage** *  **adyenIssuing** *  **adyenForPlatformsAdvanced** *  **adyenCapital** *  **adyenAccount** *  **adyenCard** *  **adyenFranchisee**
	Type string `json:"type"`
}

// NewGetTermsOfServiceDocumentRequest instantiates a new GetTermsOfServiceDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTermsOfServiceDocumentRequest(language string, type_ string) *GetTermsOfServiceDocumentRequest {
	this := GetTermsOfServiceDocumentRequest{}
	this.Language = language
	this.Type = type_
	return &this
}

// NewGetTermsOfServiceDocumentRequestWithDefaults instantiates a new GetTermsOfServiceDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTermsOfServiceDocumentRequestWithDefaults() *GetTermsOfServiceDocumentRequest {
	this := GetTermsOfServiceDocumentRequest{}
	return &this
}

// GetLanguage returns the Language field value
func (o *GetTermsOfServiceDocumentRequest) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentRequest) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *GetTermsOfServiceDocumentRequest) SetLanguage(v string) {
	o.Language = v
}

// GetType returns the Type field value
func (o *GetTermsOfServiceDocumentRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTermsOfServiceDocumentRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTermsOfServiceDocumentRequest) SetType(v string) {
	o.Type = v
}

func (o GetTermsOfServiceDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTermsOfServiceDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableGetTermsOfServiceDocumentRequest struct {
	value *GetTermsOfServiceDocumentRequest
	isSet bool
}

func (v NullableGetTermsOfServiceDocumentRequest) Get() *GetTermsOfServiceDocumentRequest {
	return v.value
}

func (v *NullableGetTermsOfServiceDocumentRequest) Set(val *GetTermsOfServiceDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTermsOfServiceDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTermsOfServiceDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTermsOfServiceDocumentRequest(val *GetTermsOfServiceDocumentRequest) *NullableGetTermsOfServiceDocumentRequest {
	return &NullableGetTermsOfServiceDocumentRequest{value: val, isSet: true}
}

func (v NullableGetTermsOfServiceDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTermsOfServiceDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *GetTermsOfServiceDocumentRequest) isValidType() bool {
	var allowedEnumValues = []string{"adyenAccount", "adyenCapital", "adyenCard", "adyenForPlatformsAdvanced", "adyenForPlatformsManage", "adyenFranchisee", "adyenIssuing"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
