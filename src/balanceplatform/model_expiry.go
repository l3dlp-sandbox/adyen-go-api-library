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

// checks if the Expiry type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Expiry{}

// Expiry struct for Expiry
type Expiry struct {
	// The month in which the card will expire.
	Month *string `json:"month,omitempty"`
	// The year in which the card will expire.
	Year *string `json:"year,omitempty"`
}

// NewExpiry instantiates a new Expiry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiry() *Expiry {
	this := Expiry{}
	return &this
}

// NewExpiryWithDefaults instantiates a new Expiry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiryWithDefaults() *Expiry {
	this := Expiry{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *Expiry) GetMonth() string {
	if o == nil || common.IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expiry) GetMonthOk() (*string, bool) {
	if o == nil || common.IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *Expiry) HasMonth() bool {
	if o != nil && !common.IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *Expiry) SetMonth(v string) {
	o.Month = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *Expiry) GetYear() string {
	if o == nil || common.IsNil(o.Year) {
		var ret string
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expiry) GetYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *Expiry) HasYear() bool {
	if o != nil && !common.IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given string and assigns it to the Year field.
func (o *Expiry) SetYear(v string) {
	o.Year = &v
}

func (o Expiry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Expiry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !common.IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	return toSerialize, nil
}

type NullableExpiry struct {
	value *Expiry
	isSet bool
}

func (v NullableExpiry) Get() *Expiry {
	return v.value
}

func (v *NullableExpiry) Set(val *Expiry) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiry(val *Expiry) *NullableExpiry {
	return &NullableExpiry{value: val, isSet: true}
}

func (v NullableExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



