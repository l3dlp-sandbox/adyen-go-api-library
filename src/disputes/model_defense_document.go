/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the DefenseDocument type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DefenseDocument{}

// DefenseDocument struct for DefenseDocument
type DefenseDocument struct {
	// The content of the defense document.
	Content string `json:"content"`
	// The content type of the defense document.
	ContentType string `json:"contentType"`
	// The document type code of the defense document.
	DefenseDocumentTypeCode string `json:"defenseDocumentTypeCode"`
}

// NewDefenseDocument instantiates a new DefenseDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefenseDocument(content string, contentType string, defenseDocumentTypeCode string) *DefenseDocument {
	this := DefenseDocument{}
	this.Content = content
	this.ContentType = contentType
	this.DefenseDocumentTypeCode = defenseDocumentTypeCode
	return &this
}

// NewDefenseDocumentWithDefaults instantiates a new DefenseDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefenseDocumentWithDefaults() *DefenseDocument {
	this := DefenseDocument{}
	return &this
}

// GetContent returns the Content field value
func (o *DefenseDocument) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *DefenseDocument) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *DefenseDocument) SetContent(v string) {
	o.Content = v
}

// GetContentType returns the ContentType field value
func (o *DefenseDocument) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *DefenseDocument) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *DefenseDocument) SetContentType(v string) {
	o.ContentType = v
}

// GetDefenseDocumentTypeCode returns the DefenseDocumentTypeCode field value
func (o *DefenseDocument) GetDefenseDocumentTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefenseDocumentTypeCode
}

// GetDefenseDocumentTypeCodeOk returns a tuple with the DefenseDocumentTypeCode field value
// and a boolean to check if the value has been set.
func (o *DefenseDocument) GetDefenseDocumentTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenseDocumentTypeCode, true
}

// SetDefenseDocumentTypeCode sets field value
func (o *DefenseDocument) SetDefenseDocumentTypeCode(v string) {
	o.DefenseDocumentTypeCode = v
}

func (o DefenseDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefenseDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["contentType"] = o.ContentType
	toSerialize["defenseDocumentTypeCode"] = o.DefenseDocumentTypeCode
	return toSerialize, nil
}

type NullableDefenseDocument struct {
	value *DefenseDocument
	isSet bool
}

func (v NullableDefenseDocument) Get() *DefenseDocument {
	return v.value
}

func (v *NullableDefenseDocument) Set(val *DefenseDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableDefenseDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDefenseDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefenseDocument(val *DefenseDocument) *NullableDefenseDocument {
	return &NullableDefenseDocument{value: val, isSet: true}
}

func (v NullableDefenseDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefenseDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



