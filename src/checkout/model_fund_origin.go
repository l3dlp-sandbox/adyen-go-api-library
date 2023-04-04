/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the FundOrigin type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundOrigin{}

// FundOrigin struct for FundOrigin
type FundOrigin struct {
	BillingAddress *Address `json:"billingAddress,omitempty"`
	ShopperName    *Name    `json:"shopperName,omitempty"`
}

// NewFundOrigin instantiates a new FundOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundOrigin() *FundOrigin {
	this := FundOrigin{}
	return &this
}

// NewFundOriginWithDefaults instantiates a new FundOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundOriginWithDefaults() *FundOrigin {
	this := FundOrigin{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *FundOrigin) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOrigin) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *FundOrigin) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *FundOrigin) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *FundOrigin) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOrigin) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *FundOrigin) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *FundOrigin) SetShopperName(v Name) {
	o.ShopperName = &v
}

func (o FundOrigin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	return toSerialize, nil
}

type NullableFundOrigin struct {
	value *FundOrigin
	isSet bool
}

func (v NullableFundOrigin) Get() *FundOrigin {
	return v.value
}

func (v *NullableFundOrigin) Set(val *FundOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableFundOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableFundOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundOrigin(val *FundOrigin) *NullableFundOrigin {
	return &NullableFundOrigin{value: val, isSet: true}
}

func (v NullableFundOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
