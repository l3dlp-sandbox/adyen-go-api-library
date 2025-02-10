/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the GetAcceptedTermsOfServiceDocumentResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAcceptedTermsOfServiceDocumentResponse{}

// GetAcceptedTermsOfServiceDocumentResponse struct for GetAcceptedTermsOfServiceDocumentResponse
type GetAcceptedTermsOfServiceDocumentResponse struct {
	// The accepted Terms of Service document in the requested format represented as a Base64-encoded bytes array.
	Document *string `json:"document,omitempty"`
	// The unique identifier of the legal entity.
	Id *string `json:"id,omitempty"`
	// An Adyen-generated reference for the accepted Terms of Service.
	TermsOfServiceAcceptanceReference *string `json:"termsOfServiceAcceptanceReference,omitempty"`
	// The format of the Terms of Service document.
	TermsOfServiceDocumentFormat *string `json:"termsOfServiceDocumentFormat,omitempty"`
}

// NewGetAcceptedTermsOfServiceDocumentResponse instantiates a new GetAcceptedTermsOfServiceDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAcceptedTermsOfServiceDocumentResponse() *GetAcceptedTermsOfServiceDocumentResponse {
	this := GetAcceptedTermsOfServiceDocumentResponse{}
	return &this
}

// NewGetAcceptedTermsOfServiceDocumentResponseWithDefaults instantiates a new GetAcceptedTermsOfServiceDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAcceptedTermsOfServiceDocumentResponseWithDefaults() *GetAcceptedTermsOfServiceDocumentResponse {
	this := GetAcceptedTermsOfServiceDocumentResponse{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetDocument() string {
	if o == nil || common.IsNil(o.Document) {
		var ret string
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetDocumentOk() (*string, bool) {
	if o == nil || common.IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) HasDocument() bool {
	if o != nil && !common.IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given string and assigns it to the Document field.
func (o *GetAcceptedTermsOfServiceDocumentResponse) SetDocument(v string) {
	o.Document = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetAcceptedTermsOfServiceDocumentResponse) SetId(v string) {
	o.Id = &v
}

// GetTermsOfServiceAcceptanceReference returns the TermsOfServiceAcceptanceReference field value if set, zero value otherwise.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetTermsOfServiceAcceptanceReference() string {
	if o == nil || common.IsNil(o.TermsOfServiceAcceptanceReference) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceAcceptanceReference
}

// GetTermsOfServiceAcceptanceReferenceOk returns a tuple with the TermsOfServiceAcceptanceReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetTermsOfServiceAcceptanceReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TermsOfServiceAcceptanceReference) {
		return nil, false
	}
	return o.TermsOfServiceAcceptanceReference, true
}

// HasTermsOfServiceAcceptanceReference returns a boolean if a field has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) HasTermsOfServiceAcceptanceReference() bool {
	if o != nil && !common.IsNil(o.TermsOfServiceAcceptanceReference) {
		return true
	}

	return false
}

// SetTermsOfServiceAcceptanceReference gets a reference to the given string and assigns it to the TermsOfServiceAcceptanceReference field.
func (o *GetAcceptedTermsOfServiceDocumentResponse) SetTermsOfServiceAcceptanceReference(v string) {
	o.TermsOfServiceAcceptanceReference = &v
}

// GetTermsOfServiceDocumentFormat returns the TermsOfServiceDocumentFormat field value if set, zero value otherwise.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetTermsOfServiceDocumentFormat() string {
	if o == nil || common.IsNil(o.TermsOfServiceDocumentFormat) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceDocumentFormat
}

// GetTermsOfServiceDocumentFormatOk returns a tuple with the TermsOfServiceDocumentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) GetTermsOfServiceDocumentFormatOk() (*string, bool) {
	if o == nil || common.IsNil(o.TermsOfServiceDocumentFormat) {
		return nil, false
	}
	return o.TermsOfServiceDocumentFormat, true
}

// HasTermsOfServiceDocumentFormat returns a boolean if a field has been set.
func (o *GetAcceptedTermsOfServiceDocumentResponse) HasTermsOfServiceDocumentFormat() bool {
	if o != nil && !common.IsNil(o.TermsOfServiceDocumentFormat) {
		return true
	}

	return false
}

// SetTermsOfServiceDocumentFormat gets a reference to the given string and assigns it to the TermsOfServiceDocumentFormat field.
func (o *GetAcceptedTermsOfServiceDocumentResponse) SetTermsOfServiceDocumentFormat(v string) {
	o.TermsOfServiceDocumentFormat = &v
}

func (o GetAcceptedTermsOfServiceDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAcceptedTermsOfServiceDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TermsOfServiceAcceptanceReference) {
		toSerialize["termsOfServiceAcceptanceReference"] = o.TermsOfServiceAcceptanceReference
	}
	if !common.IsNil(o.TermsOfServiceDocumentFormat) {
		toSerialize["termsOfServiceDocumentFormat"] = o.TermsOfServiceDocumentFormat
	}
	return toSerialize, nil
}

type NullableGetAcceptedTermsOfServiceDocumentResponse struct {
	value *GetAcceptedTermsOfServiceDocumentResponse
	isSet bool
}

func (v NullableGetAcceptedTermsOfServiceDocumentResponse) Get() *GetAcceptedTermsOfServiceDocumentResponse {
	return v.value
}

func (v *NullableGetAcceptedTermsOfServiceDocumentResponse) Set(val *GetAcceptedTermsOfServiceDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAcceptedTermsOfServiceDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAcceptedTermsOfServiceDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAcceptedTermsOfServiceDocumentResponse(val *GetAcceptedTermsOfServiceDocumentResponse) *NullableGetAcceptedTermsOfServiceDocumentResponse {
	return &NullableGetAcceptedTermsOfServiceDocumentResponse{value: val, isSet: true}
}

func (v NullableGetAcceptedTermsOfServiceDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAcceptedTermsOfServiceDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *GetAcceptedTermsOfServiceDocumentResponse) isValidTermsOfServiceDocumentFormat() bool {
	var allowedEnumValues = []string{"JSON", "PDF", "TXT"}
	for _, allowed := range allowedEnumValues {
		if o.GetTermsOfServiceDocumentFormat() == allowed {
			return true
		}
	}
	return false
}
