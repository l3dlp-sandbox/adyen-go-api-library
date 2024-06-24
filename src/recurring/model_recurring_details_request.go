/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the RecurringDetailsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecurringDetailsRequest{}

// RecurringDetailsRequest struct for RecurringDetailsRequest
type RecurringDetailsRequest struct {
	// The merchant account identifier you want to process the (transaction) request with.
	MerchantAccount string     `json:"merchantAccount"`
	Recurring       *Recurring `json:"recurring,omitempty"`
	// The reference you use to uniquely identify the shopper (e.g. user ID or account ID).
	ShopperReference string `json:"shopperReference"`
}

// NewRecurringDetailsRequest instantiates a new RecurringDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringDetailsRequest(merchantAccount string, shopperReference string) *RecurringDetailsRequest {
	this := RecurringDetailsRequest{}
	this.MerchantAccount = merchantAccount
	this.ShopperReference = shopperReference
	return &this
}

// NewRecurringDetailsRequestWithDefaults instantiates a new RecurringDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringDetailsRequestWithDefaults() *RecurringDetailsRequest {
	this := RecurringDetailsRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *RecurringDetailsRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *RecurringDetailsRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *RecurringDetailsRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *RecurringDetailsRequest) GetRecurring() Recurring {
	if o == nil || common.IsNil(o.Recurring) {
		var ret Recurring
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetailsRequest) GetRecurringOk() (*Recurring, bool) {
	if o == nil || common.IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *RecurringDetailsRequest) HasRecurring() bool {
	if o != nil && !common.IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given Recurring and assigns it to the Recurring field.
func (o *RecurringDetailsRequest) SetRecurring(v Recurring) {
	o.Recurring = &v
}

// GetShopperReference returns the ShopperReference field value
func (o *RecurringDetailsRequest) GetShopperReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value
// and a boolean to check if the value has been set.
func (o *RecurringDetailsRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperReference, true
}

// SetShopperReference sets field value
func (o *RecurringDetailsRequest) SetShopperReference(v string) {
	o.ShopperReference = v
}

func (o RecurringDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurringDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	toSerialize["shopperReference"] = o.ShopperReference
	return toSerialize, nil
}

type NullableRecurringDetailsRequest struct {
	value *RecurringDetailsRequest
	isSet bool
}

func (v NullableRecurringDetailsRequest) Get() *RecurringDetailsRequest {
	return v.value
}

func (v *NullableRecurringDetailsRequest) Set(val *RecurringDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringDetailsRequest(val *RecurringDetailsRequest) *NullableRecurringDetailsRequest {
	return &NullableRecurringDetailsRequest{value: val, isSet: true}
}

func (v NullableRecurringDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
