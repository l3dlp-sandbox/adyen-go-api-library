/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the ResponseAdditionalDataBillingAddress type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalDataBillingAddress{}

// ResponseAdditionalDataBillingAddress struct for ResponseAdditionalDataBillingAddress
type ResponseAdditionalDataBillingAddress struct {
	// The billing address city passed in the payment request.
	BillingAddressCity *string `json:"billingAddress.city,omitempty"`
	// The billing address country passed in the payment request.  Example: NL
	BillingAddressCountry *string `json:"billingAddress.country,omitempty"`
	// The billing address house number or name passed in the payment request.
	BillingAddressHouseNumberOrName *string `json:"billingAddress.houseNumberOrName,omitempty"`
	// The billing address postal code passed in the payment request.  Example: 1011 DJ
	BillingAddressPostalCode *string `json:"billingAddress.postalCode,omitempty"`
	// The billing address state or province passed in the payment request.  Example: NH
	BillingAddressStateOrProvince *string `json:"billingAddress.stateOrProvince,omitempty"`
	// The billing address street passed in the payment request.
	BillingAddressStreet *string `json:"billingAddress.street,omitempty"`
}

// NewResponseAdditionalDataBillingAddress instantiates a new ResponseAdditionalDataBillingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalDataBillingAddress() *ResponseAdditionalDataBillingAddress {
	this := ResponseAdditionalDataBillingAddress{}
	return &this
}

// NewResponseAdditionalDataBillingAddressWithDefaults instantiates a new ResponseAdditionalDataBillingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalDataBillingAddressWithDefaults() *ResponseAdditionalDataBillingAddress {
	this := ResponseAdditionalDataBillingAddress{}
	return &this
}

// GetBillingAddressCity returns the BillingAddressCity field value if set, zero value otherwise.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressCity() string {
	if o == nil || common.IsNil(o.BillingAddressCity) {
		var ret string
		return ret
	}
	return *o.BillingAddressCity
}

// GetBillingAddressCityOk returns a tuple with the BillingAddressCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddressCity) {
		return nil, false
	}
	return o.BillingAddressCity, true
}

// HasBillingAddressCity returns a boolean if a field has been set.
func (o *ResponseAdditionalDataBillingAddress) HasBillingAddressCity() bool {
	if o != nil && !common.IsNil(o.BillingAddressCity) {
		return true
	}

	return false
}

// SetBillingAddressCity gets a reference to the given string and assigns it to the BillingAddressCity field.
func (o *ResponseAdditionalDataBillingAddress) SetBillingAddressCity(v string) {
	o.BillingAddressCity = &v
}

// GetBillingAddressCountry returns the BillingAddressCountry field value if set, zero value otherwise.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressCountry() string {
	if o == nil || common.IsNil(o.BillingAddressCountry) {
		var ret string
		return ret
	}
	return *o.BillingAddressCountry
}

// GetBillingAddressCountryOk returns a tuple with the BillingAddressCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddressCountry) {
		return nil, false
	}
	return o.BillingAddressCountry, true
}

// HasBillingAddressCountry returns a boolean if a field has been set.
func (o *ResponseAdditionalDataBillingAddress) HasBillingAddressCountry() bool {
	if o != nil && !common.IsNil(o.BillingAddressCountry) {
		return true
	}

	return false
}

// SetBillingAddressCountry gets a reference to the given string and assigns it to the BillingAddressCountry field.
func (o *ResponseAdditionalDataBillingAddress) SetBillingAddressCountry(v string) {
	o.BillingAddressCountry = &v
}

// GetBillingAddressHouseNumberOrName returns the BillingAddressHouseNumberOrName field value if set, zero value otherwise.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressHouseNumberOrName() string {
	if o == nil || common.IsNil(o.BillingAddressHouseNumberOrName) {
		var ret string
		return ret
	}
	return *o.BillingAddressHouseNumberOrName
}

// GetBillingAddressHouseNumberOrNameOk returns a tuple with the BillingAddressHouseNumberOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressHouseNumberOrNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddressHouseNumberOrName) {
		return nil, false
	}
	return o.BillingAddressHouseNumberOrName, true
}

// HasBillingAddressHouseNumberOrName returns a boolean if a field has been set.
func (o *ResponseAdditionalDataBillingAddress) HasBillingAddressHouseNumberOrName() bool {
	if o != nil && !common.IsNil(o.BillingAddressHouseNumberOrName) {
		return true
	}

	return false
}

// SetBillingAddressHouseNumberOrName gets a reference to the given string and assigns it to the BillingAddressHouseNumberOrName field.
func (o *ResponseAdditionalDataBillingAddress) SetBillingAddressHouseNumberOrName(v string) {
	o.BillingAddressHouseNumberOrName = &v
}

// GetBillingAddressPostalCode returns the BillingAddressPostalCode field value if set, zero value otherwise.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressPostalCode() string {
	if o == nil || common.IsNil(o.BillingAddressPostalCode) {
		var ret string
		return ret
	}
	return *o.BillingAddressPostalCode
}

// GetBillingAddressPostalCodeOk returns a tuple with the BillingAddressPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddressPostalCode) {
		return nil, false
	}
	return o.BillingAddressPostalCode, true
}

// HasBillingAddressPostalCode returns a boolean if a field has been set.
func (o *ResponseAdditionalDataBillingAddress) HasBillingAddressPostalCode() bool {
	if o != nil && !common.IsNil(o.BillingAddressPostalCode) {
		return true
	}

	return false
}

// SetBillingAddressPostalCode gets a reference to the given string and assigns it to the BillingAddressPostalCode field.
func (o *ResponseAdditionalDataBillingAddress) SetBillingAddressPostalCode(v string) {
	o.BillingAddressPostalCode = &v
}

// GetBillingAddressStateOrProvince returns the BillingAddressStateOrProvince field value if set, zero value otherwise.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressStateOrProvince() string {
	if o == nil || common.IsNil(o.BillingAddressStateOrProvince) {
		var ret string
		return ret
	}
	return *o.BillingAddressStateOrProvince
}

// GetBillingAddressStateOrProvinceOk returns a tuple with the BillingAddressStateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddressStateOrProvince) {
		return nil, false
	}
	return o.BillingAddressStateOrProvince, true
}

// HasBillingAddressStateOrProvince returns a boolean if a field has been set.
func (o *ResponseAdditionalDataBillingAddress) HasBillingAddressStateOrProvince() bool {
	if o != nil && !common.IsNil(o.BillingAddressStateOrProvince) {
		return true
	}

	return false
}

// SetBillingAddressStateOrProvince gets a reference to the given string and assigns it to the BillingAddressStateOrProvince field.
func (o *ResponseAdditionalDataBillingAddress) SetBillingAddressStateOrProvince(v string) {
	o.BillingAddressStateOrProvince = &v
}

// GetBillingAddressStreet returns the BillingAddressStreet field value if set, zero value otherwise.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressStreet() string {
	if o == nil || common.IsNil(o.BillingAddressStreet) {
		var ret string
		return ret
	}
	return *o.BillingAddressStreet
}

// GetBillingAddressStreetOk returns a tuple with the BillingAddressStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataBillingAddress) GetBillingAddressStreetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddressStreet) {
		return nil, false
	}
	return o.BillingAddressStreet, true
}

// HasBillingAddressStreet returns a boolean if a field has been set.
func (o *ResponseAdditionalDataBillingAddress) HasBillingAddressStreet() bool {
	if o != nil && !common.IsNil(o.BillingAddressStreet) {
		return true
	}

	return false
}

// SetBillingAddressStreet gets a reference to the given string and assigns it to the BillingAddressStreet field.
func (o *ResponseAdditionalDataBillingAddress) SetBillingAddressStreet(v string) {
	o.BillingAddressStreet = &v
}

func (o ResponseAdditionalDataBillingAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalDataBillingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BillingAddressCity) {
		toSerialize["billingAddress.city"] = o.BillingAddressCity
	}
	if !common.IsNil(o.BillingAddressCountry) {
		toSerialize["billingAddress.country"] = o.BillingAddressCountry
	}
	if !common.IsNil(o.BillingAddressHouseNumberOrName) {
		toSerialize["billingAddress.houseNumberOrName"] = o.BillingAddressHouseNumberOrName
	}
	if !common.IsNil(o.BillingAddressPostalCode) {
		toSerialize["billingAddress.postalCode"] = o.BillingAddressPostalCode
	}
	if !common.IsNil(o.BillingAddressStateOrProvince) {
		toSerialize["billingAddress.stateOrProvince"] = o.BillingAddressStateOrProvince
	}
	if !common.IsNil(o.BillingAddressStreet) {
		toSerialize["billingAddress.street"] = o.BillingAddressStreet
	}
	return toSerialize, nil
}

type NullableResponseAdditionalDataBillingAddress struct {
	value *ResponseAdditionalDataBillingAddress
	isSet bool
}

func (v NullableResponseAdditionalDataBillingAddress) Get() *ResponseAdditionalDataBillingAddress {
	return v.value
}

func (v *NullableResponseAdditionalDataBillingAddress) Set(val *ResponseAdditionalDataBillingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalDataBillingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalDataBillingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalDataBillingAddress(val *ResponseAdditionalDataBillingAddress) *NullableResponseAdditionalDataBillingAddress {
	return &NullableResponseAdditionalDataBillingAddress{value: val, isSet: true}
}

func (v NullableResponseAdditionalDataBillingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalDataBillingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
