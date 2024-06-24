/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the JSONPath type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &JSONPath{}

// JSONPath struct for JSONPath
type JSONPath struct {
	Content []string `json:"content,omitempty"`
}

// NewJSONPath instantiates a new JSONPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONPath() *JSONPath {
	this := JSONPath{}
	return &this
}

// NewJSONPathWithDefaults instantiates a new JSONPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONPathWithDefaults() *JSONPath {
	this := JSONPath{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *JSONPath) GetContent() []string {
	if o == nil || common.IsNil(o.Content) {
		var ret []string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONPath) GetContentOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *JSONPath) HasContent() bool {
	if o != nil && !common.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []string and assigns it to the Content field.
func (o *JSONPath) SetContent(v []string) {
	o.Content = v
}

func (o JSONPath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSONPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableJSONPath struct {
	value *JSONPath
	isSet bool
}

func (v NullableJSONPath) Get() *JSONPath {
	return v.value
}

func (v *NullableJSONPath) Set(val *JSONPath) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONPath) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONPath(val *JSONPath) *NullableJSONPath {
	return &NullableJSONPath{value: val, isSet: true}
}

func (v NullableJSONPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
