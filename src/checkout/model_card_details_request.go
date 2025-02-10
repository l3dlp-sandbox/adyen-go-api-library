/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the CardDetailsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardDetailsRequest{}

// CardDetailsRequest struct for CardDetailsRequest
type CardDetailsRequest struct {
	// A minimum of the first eight digits of the card number. The full card number gives the best result.   You must be [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide) to collect raw card data.
	CardNumber string `json:"cardNumber"`
	// The shopper country.  Format: [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) Example: NL or DE
	CountryCode *string `json:"countryCode,omitempty"`
	// The encrypted card number.
	EncryptedCardNumber *string `json:"encryptedCardNumber,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The card brands you support. This is the [`brands`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods__resParam_paymentMethods-brands) array from your [`/paymentMethods`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods) response.   If not included, our API uses the ones configured for your merchant account and, if provided, the country code.
	SupportedBrands []string `json:"supportedBrands,omitempty"`
}

// NewCardDetailsRequest instantiates a new CardDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardDetailsRequest(cardNumber string, merchantAccount string) *CardDetailsRequest {
	this := CardDetailsRequest{}
	this.CardNumber = cardNumber
	this.MerchantAccount = merchantAccount
	return &this
}

// NewCardDetailsRequestWithDefaults instantiates a new CardDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardDetailsRequestWithDefaults() *CardDetailsRequest {
	this := CardDetailsRequest{}
	return &this
}

// GetCardNumber returns the CardNumber field value
func (o *CardDetailsRequest) GetCardNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value
// and a boolean to check if the value has been set.
func (o *CardDetailsRequest) GetCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardNumber, true
}

// SetCardNumber sets field value
func (o *CardDetailsRequest) SetCardNumber(v string) {
	o.CardNumber = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *CardDetailsRequest) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetailsRequest) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *CardDetailsRequest) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *CardDetailsRequest) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEncryptedCardNumber returns the EncryptedCardNumber field value if set, zero value otherwise.
func (o *CardDetailsRequest) GetEncryptedCardNumber() string {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		var ret string
		return ret
	}
	return *o.EncryptedCardNumber
}

// GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetailsRequest) GetEncryptedCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		return nil, false
	}
	return o.EncryptedCardNumber, true
}

// HasEncryptedCardNumber returns a boolean if a field has been set.
func (o *CardDetailsRequest) HasEncryptedCardNumber() bool {
	if o != nil && !common.IsNil(o.EncryptedCardNumber) {
		return true
	}

	return false
}

// SetEncryptedCardNumber gets a reference to the given string and assigns it to the EncryptedCardNumber field.
func (o *CardDetailsRequest) SetEncryptedCardNumber(v string) {
	o.EncryptedCardNumber = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CardDetailsRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CardDetailsRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CardDetailsRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetSupportedBrands returns the SupportedBrands field value if set, zero value otherwise.
func (o *CardDetailsRequest) GetSupportedBrands() []string {
	if o == nil || common.IsNil(o.SupportedBrands) {
		var ret []string
		return ret
	}
	return o.SupportedBrands
}

// GetSupportedBrandsOk returns a tuple with the SupportedBrands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetailsRequest) GetSupportedBrandsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SupportedBrands) {
		return nil, false
	}
	return o.SupportedBrands, true
}

// HasSupportedBrands returns a boolean if a field has been set.
func (o *CardDetailsRequest) HasSupportedBrands() bool {
	if o != nil && !common.IsNil(o.SupportedBrands) {
		return true
	}

	return false
}

// SetSupportedBrands gets a reference to the given []string and assigns it to the SupportedBrands field.
func (o *CardDetailsRequest) SetSupportedBrands(v []string) {
	o.SupportedBrands = v
}

func (o CardDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cardNumber"] = o.CardNumber
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !common.IsNil(o.EncryptedCardNumber) {
		toSerialize["encryptedCardNumber"] = o.EncryptedCardNumber
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.SupportedBrands) {
		toSerialize["supportedBrands"] = o.SupportedBrands
	}
	return toSerialize, nil
}

type NullableCardDetailsRequest struct {
	value *CardDetailsRequest
	isSet bool
}

func (v NullableCardDetailsRequest) Get() *CardDetailsRequest {
	return v.value
}

func (v *NullableCardDetailsRequest) Set(val *CardDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardDetailsRequest(val *CardDetailsRequest) *NullableCardDetailsRequest {
	return &NullableCardDetailsRequest{value: val, isSet: true}
}

func (v NullableCardDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
