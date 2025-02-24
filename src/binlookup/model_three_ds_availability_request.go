/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the ThreeDSAvailabilityRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDSAvailabilityRequest{}

// ThreeDSAvailabilityRequest struct for ThreeDSAvailabilityRequest
type ThreeDSAvailabilityRequest struct {
	// This field contains additional data, which may be required for a particular request.  The `additionalData` object consists of entries, each of which includes the key and value.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// List of brands.
	Brands []string `json:"brands,omitempty"`
	// Card number or BIN.
	CardNumber *string `json:"cardNumber,omitempty"`
	// The merchant account identifier.
	MerchantAccount string `json:"merchantAccount"`
	// A recurring detail reference corresponding to a card.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// The shopper's reference to uniquely identify this shopper (e.g. user ID or account ID).
	ShopperReference *string `json:"shopperReference,omitempty"`
}

// NewThreeDSAvailabilityRequest instantiates a new ThreeDSAvailabilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSAvailabilityRequest(merchantAccount string) *ThreeDSAvailabilityRequest {
	this := ThreeDSAvailabilityRequest{}
	this.MerchantAccount = merchantAccount
	return &this
}

// NewThreeDSAvailabilityRequestWithDefaults instantiates a new ThreeDSAvailabilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSAvailabilityRequestWithDefaults() *ThreeDSAvailabilityRequest {
	this := ThreeDSAvailabilityRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ThreeDSAvailabilityRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityRequest) GetBrands() []string {
	if o == nil || common.IsNil(o.Brands) {
		var ret []string
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityRequest) GetBrandsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityRequest) HasBrands() bool {
	if o != nil && !common.IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []string and assigns it to the Brands field.
func (o *ThreeDSAvailabilityRequest) SetBrands(v []string) {
	o.Brands = v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityRequest) GetCardNumber() string {
	if o == nil || common.IsNil(o.CardNumber) {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityRequest) GetCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardNumber) {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityRequest) HasCardNumber() bool {
	if o != nil && !common.IsNil(o.CardNumber) {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *ThreeDSAvailabilityRequest) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *ThreeDSAvailabilityRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *ThreeDSAvailabilityRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityRequest) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityRequest) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityRequest) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
func (o *ThreeDSAvailabilityRequest) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityRequest) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityRequest) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *ThreeDSAvailabilityRequest) SetShopperReference(v string) {
	o.ShopperReference = &v
}

func (o ThreeDSAvailabilityRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSAvailabilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.Brands) {
		toSerialize["brands"] = o.Brands
	}
	if !common.IsNil(o.CardNumber) {
		toSerialize["cardNumber"] = o.CardNumber
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	return toSerialize, nil
}

type NullableThreeDSAvailabilityRequest struct {
	value *ThreeDSAvailabilityRequest
	isSet bool
}

func (v NullableThreeDSAvailabilityRequest) Get() *ThreeDSAvailabilityRequest {
	return v.value
}

func (v *NullableThreeDSAvailabilityRequest) Set(val *ThreeDSAvailabilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSAvailabilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSAvailabilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSAvailabilityRequest(val *ThreeDSAvailabilityRequest) *NullableThreeDSAvailabilityRequest {
	return &NullableThreeDSAvailabilityRequest{value: val, isSet: true}
}

func (v NullableThreeDSAvailabilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSAvailabilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
