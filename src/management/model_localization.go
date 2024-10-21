/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the Localization type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Localization{}

// Localization struct for Localization
type Localization struct {
	// Language of the terminal.
	Language *string `json:"language,omitempty"`
	// Secondary language of the terminal.
	SecondaryLanguage *string `json:"secondaryLanguage,omitempty"`
	// The time zone of the terminal.
	Timezone *string `json:"timezone,omitempty"`
}

// NewLocalization instantiates a new Localization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalization() *Localization {
	this := Localization{}
	return &this
}

// NewLocalizationWithDefaults instantiates a new Localization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalizationWithDefaults() *Localization {
	this := Localization{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Localization) GetLanguage() string {
	if o == nil || common.IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Localization) GetLanguageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Localization) HasLanguage() bool {
	if o != nil && !common.IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Localization) SetLanguage(v string) {
	o.Language = &v
}

// GetSecondaryLanguage returns the SecondaryLanguage field value if set, zero value otherwise.
func (o *Localization) GetSecondaryLanguage() string {
	if o == nil || common.IsNil(o.SecondaryLanguage) {
		var ret string
		return ret
	}
	return *o.SecondaryLanguage
}

// GetSecondaryLanguageOk returns a tuple with the SecondaryLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Localization) GetSecondaryLanguageOk() (*string, bool) {
	if o == nil || common.IsNil(o.SecondaryLanguage) {
		return nil, false
	}
	return o.SecondaryLanguage, true
}

// HasSecondaryLanguage returns a boolean if a field has been set.
func (o *Localization) HasSecondaryLanguage() bool {
	if o != nil && !common.IsNil(o.SecondaryLanguage) {
		return true
	}

	return false
}

// SetSecondaryLanguage gets a reference to the given string and assigns it to the SecondaryLanguage field.
func (o *Localization) SetSecondaryLanguage(v string) {
	o.SecondaryLanguage = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Localization) GetTimezone() string {
	if o == nil || common.IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Localization) GetTimezoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Localization) HasTimezone() bool {
	if o != nil && !common.IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Localization) SetTimezone(v string) {
	o.Timezone = &v
}

func (o Localization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Localization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !common.IsNil(o.SecondaryLanguage) {
		toSerialize["secondaryLanguage"] = o.SecondaryLanguage
	}
	if !common.IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

type NullableLocalization struct {
	value *Localization
	isSet bool
}

func (v NullableLocalization) Get() *Localization {
	return v.value
}

func (v *NullableLocalization) Set(val *Localization) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalization(val *Localization) *NullableLocalization {
	return &NullableLocalization{value: val, isSet: true}
}

func (v NullableLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
