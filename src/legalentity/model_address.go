/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Address{}

// Address struct for Address
type Address struct {
	// The name of the city. Required if `stateOrProvince` is provided.  If you specify the city, you must also send `postalCode` and `street`.
	City *string `json:"city,omitempty"`
	// The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code.
	Country string `json:"country"`
	// Postal code. Required if `stateOrProvince` and/or `city` is provided.
	PostalCode *string `json:"postalCode,omitempty"`
	// The two-letter ISO 3166-2 state or province code. For example, **CA** in the US.  If you specify the state or province, you must also send `city`, `postalCode`, and `street`.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// The name of the street, and the house or building number. Required if `stateOrProvince` and/or `city` is provided.
	Street *string `json:"street,omitempty"`
	// The apartment, unit, or suite number.
	Street2 *string `json:"street2,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(country string) *Address {
	this := Address{}
	this.Country = country
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Address) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Address) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Address) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value
func (o *Address) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Address) SetCountry(v string) {
	o.Country = v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Address) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Address) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Address) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *Address) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *Address) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *Address) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *Address) GetStreet() string {
	if o == nil || common.IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *Address) HasStreet() bool {
	if o != nil && !common.IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *Address) SetStreet(v string) {
	o.Street = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *Address) GetStreet2() string {
	if o == nil || common.IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreet2Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *Address) HasStreet2() bool {
	if o != nil && !common.IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *Address) SetStreet2(v string) {
	o.Street2 = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	toSerialize["country"] = o.Country
	if !common.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !common.IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	if !common.IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !common.IsNil(o.Street2) {
		toSerialize["street2"] = o.Street2
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
