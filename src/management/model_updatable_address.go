/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the UpdatableAddress type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdatableAddress{}

// UpdatableAddress struct for UpdatableAddress
type UpdatableAddress struct {
	// The name of the city.
	City *string `json:"city,omitempty"`
	// The street address.
	Line1 *string `json:"line1,omitempty"`
	// Second address line.
	Line2 *string `json:"line2,omitempty"`
	// Third address line.
	Line3 *string `json:"line3,omitempty"`
	// The postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The state or province code as defined in [ISO 3166-2](https://www.iso.org/standard/72483.html). For example, **ON** for Ontario, Canada.  Required for the following countries:  - Australia - Brazil - Canada - India - Mexico - New Zealand - United States
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

// NewUpdatableAddress instantiates a new UpdatableAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatableAddress() *UpdatableAddress {
	this := UpdatableAddress{}
	return &this
}

// NewUpdatableAddressWithDefaults instantiates a new UpdatableAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatableAddressWithDefaults() *UpdatableAddress {
	this := UpdatableAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UpdatableAddress) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatableAddress) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UpdatableAddress) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UpdatableAddress) SetCity(v string) {
	o.City = &v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *UpdatableAddress) GetLine1() string {
	if o == nil || common.IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatableAddress) GetLine1Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *UpdatableAddress) HasLine1() bool {
	if o != nil && !common.IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *UpdatableAddress) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *UpdatableAddress) GetLine2() string {
	if o == nil || common.IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatableAddress) GetLine2Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *UpdatableAddress) HasLine2() bool {
	if o != nil && !common.IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *UpdatableAddress) SetLine2(v string) {
	o.Line2 = &v
}

// GetLine3 returns the Line3 field value if set, zero value otherwise.
func (o *UpdatableAddress) GetLine3() string {
	if o == nil || common.IsNil(o.Line3) {
		var ret string
		return ret
	}
	return *o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatableAddress) GetLine3Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line3) {
		return nil, false
	}
	return o.Line3, true
}

// HasLine3 returns a boolean if a field has been set.
func (o *UpdatableAddress) HasLine3() bool {
	if o != nil && !common.IsNil(o.Line3) {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given string and assigns it to the Line3 field.
func (o *UpdatableAddress) SetLine3(v string) {
	o.Line3 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *UpdatableAddress) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatableAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *UpdatableAddress) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *UpdatableAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *UpdatableAddress) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatableAddress) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *UpdatableAddress) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *UpdatableAddress) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

func (o UpdatableAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatableAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !common.IsNil(o.Line1) {
		toSerialize["line1"] = o.Line1
	}
	if !common.IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	if !common.IsNil(o.Line3) {
		toSerialize["line3"] = o.Line3
	}
	if !common.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !common.IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	return toSerialize, nil
}

type NullableUpdatableAddress struct {
	value *UpdatableAddress
	isSet bool
}

func (v NullableUpdatableAddress) Get() *UpdatableAddress {
	return v.value
}

func (v *NullableUpdatableAddress) Set(val *UpdatableAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatableAddress(val *UpdatableAddress) *NullableUpdatableAddress {
	return &NullableUpdatableAddress{value: val, isSet: true}
}

func (v NullableUpdatableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
