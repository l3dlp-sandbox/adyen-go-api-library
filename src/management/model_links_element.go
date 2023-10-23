/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the LinksElement type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LinksElement{}

// LinksElement struct for LinksElement
type LinksElement struct {
	Href *string `json:"href,omitempty"`
}

// NewLinksElement instantiates a new LinksElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksElement() *LinksElement {
	this := LinksElement{}
	return &this
}

// NewLinksElementWithDefaults instantiates a new LinksElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksElementWithDefaults() *LinksElement {
	this := LinksElement{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LinksElement) GetHref() string {
	if o == nil || common.IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksElement) GetHrefOk() (*string, bool) {
	if o == nil || common.IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LinksElement) HasHref() bool {
	if o != nil && !common.IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LinksElement) SetHref(v string) {
	o.Href = &v
}

func (o LinksElement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableLinksElement struct {
	value *LinksElement
	isSet bool
}

func (v NullableLinksElement) Get() *LinksElement {
	return v.value
}

func (v *NullableLinksElement) Set(val *LinksElement) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksElement) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksElement(val *LinksElement) *NullableLinksElement {
	return &NullableLinksElement{value: val, isSet: true}
}

func (v NullableLinksElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
