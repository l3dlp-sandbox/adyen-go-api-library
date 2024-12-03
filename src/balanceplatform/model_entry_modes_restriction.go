/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the EntryModesRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EntryModesRestriction{}

// EntryModesRestriction struct for EntryModesRestriction
type EntryModesRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	// List of point-of-sale entry modes.  Possible values: **barcode**, **chip**, **cof**, **contactless**, **magstripe**, **manual**, **ocr**, **server**.  
	Value []string `json:"value,omitempty"`
}

// NewEntryModesRestriction instantiates a new EntryModesRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryModesRestriction(operation string) *EntryModesRestriction {
	this := EntryModesRestriction{}
	this.Operation = operation
	return &this
}

// NewEntryModesRestrictionWithDefaults instantiates a new EntryModesRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryModesRestrictionWithDefaults() *EntryModesRestriction {
	this := EntryModesRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *EntryModesRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *EntryModesRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *EntryModesRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntryModesRestriction) GetValue() []string {
	if o == nil || common.IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryModesRestriction) GetValueOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntryModesRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *EntryModesRestriction) SetValue(v []string) {
	o.Value = v
}

func (o EntryModesRestriction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntryModesRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEntryModesRestriction struct {
	value *EntryModesRestriction
	isSet bool
}

func (v NullableEntryModesRestriction) Get() *EntryModesRestriction {
	return v.value
}

func (v *NullableEntryModesRestriction) Set(val *EntryModesRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryModesRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryModesRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryModesRestriction(val *EntryModesRestriction) *NullableEntryModesRestriction {
	return &NullableEntryModesRestriction{value: val, isSet: true}
}

func (v NullableEntryModesRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryModesRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



