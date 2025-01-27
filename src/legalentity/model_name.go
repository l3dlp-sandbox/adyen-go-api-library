/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the Name type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Name{}

// Name struct for Name
type Name struct {
	// The individual's first name. Must not be blank.
	FirstName string `json:"firstName"`
	// The infix in the individual's name, if any.
	Infix *string `json:"infix,omitempty"`
	// The individual's last name. Must not be blank.
	LastName string `json:"lastName"`
}

// NewName instantiates a new Name object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewName(firstName string, lastName string) *Name {
	this := Name{}
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewNameWithDefaults instantiates a new Name object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameWithDefaults() *Name {
	this := Name{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *Name) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *Name) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *Name) SetFirstName(v string) {
	o.FirstName = v
}

// GetInfix returns the Infix field value if set, zero value otherwise.
func (o *Name) GetInfix() string {
	if o == nil || common.IsNil(o.Infix) {
		var ret string
		return ret
	}
	return *o.Infix
}

// GetInfixOk returns a tuple with the Infix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetInfixOk() (*string, bool) {
	if o == nil || common.IsNil(o.Infix) {
		return nil, false
	}
	return o.Infix, true
}

// HasInfix returns a boolean if a field has been set.
func (o *Name) HasInfix() bool {
	if o != nil && !common.IsNil(o.Infix) {
		return true
	}

	return false
}

// SetInfix gets a reference to the given string and assigns it to the Infix field.
func (o *Name) SetInfix(v string) {
	o.Infix = &v
}

// GetLastName returns the LastName field value
func (o *Name) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *Name) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *Name) SetLastName(v string) {
	o.LastName = v
}

func (o Name) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Name) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstName"] = o.FirstName
	if !common.IsNil(o.Infix) {
		toSerialize["infix"] = o.Infix
	}
	toSerialize["lastName"] = o.LastName
	return toSerialize, nil
}

type NullableName struct {
	value *Name
	isSet bool
}

func (v NullableName) Get() *Name {
	return v.value
}

func (v *NullableName) Set(val *Name) {
	v.value = val
	v.isSet = true
}

func (v NullableName) IsSet() bool {
	return v.isSet
}

func (v *NullableName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableName(val *Name) *NullableName {
	return &NullableName{value: val, isSet: true}
}

func (v NullableName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
