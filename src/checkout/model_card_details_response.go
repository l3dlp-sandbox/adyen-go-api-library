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

// checks if the CardDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardDetailsResponse{}

// CardDetailsResponse struct for CardDetailsResponse
type CardDetailsResponse struct {
	// The list of brands identified for the card.
	Brands []CardBrandDetails `json:"brands,omitempty"`
	// The funding source of the card, for example **DEBIT**, **CREDIT**, or **PREPAID**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// Indicates if this is a commercial card or a consumer card. If **true**, it is a commercial card. If **false**, it is a consumer card.
	IsCardCommercial *bool `json:"isCardCommercial,omitempty"`
	// The two-letter country code  of the country where the card was issued.
	IssuingCountryCode *string `json:"issuingCountryCode,omitempty"`
}

// NewCardDetailsResponse instantiates a new CardDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardDetailsResponse() *CardDetailsResponse {
	this := CardDetailsResponse{}
	return &this
}

// NewCardDetailsResponseWithDefaults instantiates a new CardDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardDetailsResponseWithDefaults() *CardDetailsResponse {
	this := CardDetailsResponse{}
	return &this
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *CardDetailsResponse) GetBrands() []CardBrandDetails {
	if o == nil || common.IsNil(o.Brands) {
		var ret []CardBrandDetails
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetailsResponse) GetBrandsOk() ([]CardBrandDetails, bool) {
	if o == nil || common.IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *CardDetailsResponse) HasBrands() bool {
	if o != nil && !common.IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []CardBrandDetails and assigns it to the Brands field.
func (o *CardDetailsResponse) SetBrands(v []CardBrandDetails) {
	o.Brands = v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *CardDetailsResponse) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetailsResponse) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *CardDetailsResponse) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *CardDetailsResponse) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetIsCardCommercial returns the IsCardCommercial field value if set, zero value otherwise.
func (o *CardDetailsResponse) GetIsCardCommercial() bool {
	if o == nil || common.IsNil(o.IsCardCommercial) {
		var ret bool
		return ret
	}
	return *o.IsCardCommercial
}

// GetIsCardCommercialOk returns a tuple with the IsCardCommercial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetailsResponse) GetIsCardCommercialOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsCardCommercial) {
		return nil, false
	}
	return o.IsCardCommercial, true
}

// HasIsCardCommercial returns a boolean if a field has been set.
func (o *CardDetailsResponse) HasIsCardCommercial() bool {
	if o != nil && !common.IsNil(o.IsCardCommercial) {
		return true
	}

	return false
}

// SetIsCardCommercial gets a reference to the given bool and assigns it to the IsCardCommercial field.
func (o *CardDetailsResponse) SetIsCardCommercial(v bool) {
	o.IsCardCommercial = &v
}

// GetIssuingCountryCode returns the IssuingCountryCode field value if set, zero value otherwise.
func (o *CardDetailsResponse) GetIssuingCountryCode() string {
	if o == nil || common.IsNil(o.IssuingCountryCode) {
		var ret string
		return ret
	}
	return *o.IssuingCountryCode
}

// GetIssuingCountryCodeOk returns a tuple with the IssuingCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetailsResponse) GetIssuingCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuingCountryCode) {
		return nil, false
	}
	return o.IssuingCountryCode, true
}

// HasIssuingCountryCode returns a boolean if a field has been set.
func (o *CardDetailsResponse) HasIssuingCountryCode() bool {
	if o != nil && !common.IsNil(o.IssuingCountryCode) {
		return true
	}

	return false
}

// SetIssuingCountryCode gets a reference to the given string and assigns it to the IssuingCountryCode field.
func (o *CardDetailsResponse) SetIssuingCountryCode(v string) {
	o.IssuingCountryCode = &v
}

func (o CardDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Brands) {
		toSerialize["brands"] = o.Brands
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	if !common.IsNil(o.IsCardCommercial) {
		toSerialize["isCardCommercial"] = o.IsCardCommercial
	}
	if !common.IsNil(o.IssuingCountryCode) {
		toSerialize["issuingCountryCode"] = o.IssuingCountryCode
	}
	return toSerialize, nil
}

type NullableCardDetailsResponse struct {
	value *CardDetailsResponse
	isSet bool
}

func (v NullableCardDetailsResponse) Get() *CardDetailsResponse {
	return v.value
}

func (v *NullableCardDetailsResponse) Set(val *CardDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardDetailsResponse(val *CardDetailsResponse) *NullableCardDetailsResponse {
	return &NullableCardDetailsResponse{value: val, isSet: true}
}

func (v NullableCardDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
