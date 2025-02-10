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

// checks if the DocumentPage type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DocumentPage{}

// DocumentPage struct for DocumentPage
type DocumentPage struct {
	PageName   *string `json:"pageName,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty"`
	Type       *string `json:"type,omitempty"`
}

// NewDocumentPage instantiates a new DocumentPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentPage() *DocumentPage {
	this := DocumentPage{}
	return &this
}

// NewDocumentPageWithDefaults instantiates a new DocumentPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentPageWithDefaults() *DocumentPage {
	this := DocumentPage{}
	return &this
}

// GetPageName returns the PageName field value if set, zero value otherwise.
func (o *DocumentPage) GetPageName() string {
	if o == nil || common.IsNil(o.PageName) {
		var ret string
		return ret
	}
	return *o.PageName
}

// GetPageNameOk returns a tuple with the PageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPage) GetPageNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PageName) {
		return nil, false
	}
	return o.PageName, true
}

// HasPageName returns a boolean if a field has been set.
func (o *DocumentPage) HasPageName() bool {
	if o != nil && !common.IsNil(o.PageName) {
		return true
	}

	return false
}

// SetPageName gets a reference to the given string and assigns it to the PageName field.
func (o *DocumentPage) SetPageName(v string) {
	o.PageName = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DocumentPage) GetPageNumber() int32 {
	if o == nil || common.IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPage) GetPageNumberOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DocumentPage) HasPageNumber() bool {
	if o != nil && !common.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *DocumentPage) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DocumentPage) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPage) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DocumentPage) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DocumentPage) SetType(v string) {
	o.Type = &v
}

func (o DocumentPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PageName) {
		toSerialize["pageName"] = o.PageName
	}
	if !common.IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDocumentPage struct {
	value *DocumentPage
	isSet bool
}

func (v NullableDocumentPage) Get() *DocumentPage {
	return v.value
}

func (v *NullableDocumentPage) Set(val *DocumentPage) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPage) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPage(val *DocumentPage) *NullableDocumentPage {
	return &NullableDocumentPage{value: val, isSet: true}
}

func (v NullableDocumentPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DocumentPage) isValidType() bool {
	var allowedEnumValues = []string{"BACK", "FRONT", "UNDEFINED"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
