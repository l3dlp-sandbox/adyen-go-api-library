/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the CommonField type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CommonField{}

// CommonField struct for CommonField
type CommonField struct {
	// Name of the field. For example, Name of External Platform.
	Name *string `json:"name,omitempty"`
	// Version of the field. For example, Version of External Platform.
	Version *string `json:"version,omitempty"`
}

// NewCommonField instantiates a new CommonField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonField() *CommonField {
	this := CommonField{}
	return &this
}

// NewCommonFieldWithDefaults instantiates a new CommonField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonFieldWithDefaults() *CommonField {
	this := CommonField{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommonField) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonField) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommonField) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommonField) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CommonField) GetVersion() string {
	if o == nil || common.IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonField) GetVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CommonField) HasVersion() bool {
	if o != nil && !common.IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CommonField) SetVersion(v string) {
	o.Version = &v
}

func (o CommonField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCommonField struct {
	value *CommonField
	isSet bool
}

func (v NullableCommonField) Get() *CommonField {
	return v.value
}

func (v *NullableCommonField) Set(val *CommonField) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonField) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonField(val *CommonField) *NullableCommonField {
	return &NullableCommonField{value: val, isSet: true}
}

func (v NullableCommonField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
