/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the DeliveryAddress type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeliveryAddress{}

// DeliveryAddress struct for DeliveryAddress
type DeliveryAddress struct {
	// The name of the city.
	City *string `json:"city,omitempty"`
	// The two-character ISO-3166-1 alpha-2 country code. For example, **US**. >If you don't know the country or are not collecting the country from the shopper, provide `country` as `ZZ`.
	Country string `json:"country"`
	// The name of the street. Do not include the number of the building.  For example, if the address is Simon Carmiggeltstraat 6-50, provide **Simon Carmiggeltstraat**.
	Line1 *string `json:"line1,omitempty"`
	// The number of the building.  For example, if the address is Simon Carmiggeltstraat 6-50, provide **6-50**.
	Line2 *string `json:"line2,omitempty"`
	// Additional information about the delivery address.
	Line3 *string `json:"line3,omitempty"`
	// The postal code. Maximum length: * 5 digits for an address in the US. * 10 characters for an address in all other countries.
	PostalCode *string `json:"postalCode,omitempty"`
	// The two-letter ISO 3166-2 state or province code. For example, **CA** in the US or **ON** in Canada. > Required for the US and Canada.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

// NewDeliveryAddress instantiates a new DeliveryAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryAddress(country string) *DeliveryAddress {
	this := DeliveryAddress{}
	this.Country = country
	return &this
}

// NewDeliveryAddressWithDefaults instantiates a new DeliveryAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryAddressWithDefaults() *DeliveryAddress {
	this := DeliveryAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *DeliveryAddress) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *DeliveryAddress) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *DeliveryAddress) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value
func (o *DeliveryAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *DeliveryAddress) SetCountry(v string) {
	o.Country = v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *DeliveryAddress) GetLine1() string {
	if o == nil || common.IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetLine1Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *DeliveryAddress) HasLine1() bool {
	if o != nil && !common.IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *DeliveryAddress) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *DeliveryAddress) GetLine2() string {
	if o == nil || common.IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetLine2Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *DeliveryAddress) HasLine2() bool {
	if o != nil && !common.IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *DeliveryAddress) SetLine2(v string) {
	o.Line2 = &v
}

// GetLine3 returns the Line3 field value if set, zero value otherwise.
func (o *DeliveryAddress) GetLine3() string {
	if o == nil || common.IsNil(o.Line3) {
		var ret string
		return ret
	}
	return *o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetLine3Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line3) {
		return nil, false
	}
	return o.Line3, true
}

// HasLine3 returns a boolean if a field has been set.
func (o *DeliveryAddress) HasLine3() bool {
	if o != nil && !common.IsNil(o.Line3) {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given string and assigns it to the Line3 field.
func (o *DeliveryAddress) SetLine3(v string) {
	o.Line3 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *DeliveryAddress) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *DeliveryAddress) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *DeliveryAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *DeliveryAddress) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAddress) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *DeliveryAddress) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *DeliveryAddress) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

func (o DeliveryAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	toSerialize["country"] = o.Country
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

type NullableDeliveryAddress struct {
	value *DeliveryAddress
	isSet bool
}

func (v NullableDeliveryAddress) Get() *DeliveryAddress {
	return v.value
}

func (v *NullableDeliveryAddress) Set(val *DeliveryAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryAddress(val *DeliveryAddress) *NullableDeliveryAddress {
	return &NullableDeliveryAddress{value: val, isSet: true}
}

func (v NullableDeliveryAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
