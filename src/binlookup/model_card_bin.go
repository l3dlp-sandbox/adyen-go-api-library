/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the CardBin type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardBin{}

// CardBin struct for CardBin
type CardBin struct {
	// The first 6 digit of the card number. Enable this field via merchant account settings.
	Bin *string `json:"bin,omitempty"`
	// If true, it indicates a commercial card. Enable this field via merchant account settings.
	Commercial *bool `json:"commercial,omitempty"`
	// The card funding source. Valid values are: * CHARGE * CREDIT * DEBIT * DEFERRED_DEBIT * PREPAID * PREPAID_RELOADABLE * PREPAID_NONRELOADABLE > Enable this field via merchant account settings.
	FundingSource *string `json:"fundingSource,omitempty"`
	// Indicates availability of funds.  Visa: * \"I\" (fast funds are supported) * \"N\" (otherwise)  Mastercard: * \"I\" (product type is Prepaid or Debit, or issuing country is in CEE/HGEM list) * \"N\" (otherwise) > Returned when you verify a card BIN or estimate costs, and only if `payoutEligible` is different from \"N\" or \"U\".
	FundsAvailability *string `json:"fundsAvailability,omitempty"`
	// The first 8 digit of the card number. Enable this field via merchant account settings.
	IssuerBin *string `json:"issuerBin,omitempty"`
	// The issuing bank of the card.
	IssuingBank *string `json:"issuingBank,omitempty"`
	// The country where the card was issued from.
	IssuingCountry *string `json:"issuingCountry,omitempty"`
	// The currency of the card.
	IssuingCurrency *string `json:"issuingCurrency,omitempty"`
	// The payment method associated with the card (e.g. visa, mc, or amex).
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// Indicates whether a payout is eligible or not for this card.  Visa: * \"Y\" * \"N\"  Mastercard: * \"Y\" (domestic and cross-border) * \"D\" (only domestic) * \"N\" (no MoneySend) * \"U\" (unknown) > Returned when you verify a card BIN or estimate costs, and only if `payoutEligible` is different from \"N\" or \"U\".
	PayoutEligible *string `json:"payoutEligible,omitempty"`
	// The last four digits of the card number.
	Summary *string `json:"summary,omitempty"`
}

// NewCardBin instantiates a new CardBin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardBin() *CardBin {
	this := CardBin{}
	return &this
}

// NewCardBinWithDefaults instantiates a new CardBin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardBinWithDefaults() *CardBin {
	this := CardBin{}
	return &this
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *CardBin) GetBin() string {
	if o == nil || common.IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetBinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *CardBin) HasBin() bool {
	if o != nil && !common.IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *CardBin) SetBin(v string) {
	o.Bin = &v
}

// GetCommercial returns the Commercial field value if set, zero value otherwise.
func (o *CardBin) GetCommercial() bool {
	if o == nil || common.IsNil(o.Commercial) {
		var ret bool
		return ret
	}
	return *o.Commercial
}

// GetCommercialOk returns a tuple with the Commercial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetCommercialOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Commercial) {
		return nil, false
	}
	return o.Commercial, true
}

// HasCommercial returns a boolean if a field has been set.
func (o *CardBin) HasCommercial() bool {
	if o != nil && !common.IsNil(o.Commercial) {
		return true
	}

	return false
}

// SetCommercial gets a reference to the given bool and assigns it to the Commercial field.
func (o *CardBin) SetCommercial(v bool) {
	o.Commercial = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *CardBin) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *CardBin) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *CardBin) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetFundsAvailability returns the FundsAvailability field value if set, zero value otherwise.
func (o *CardBin) GetFundsAvailability() string {
	if o == nil || common.IsNil(o.FundsAvailability) {
		var ret string
		return ret
	}
	return *o.FundsAvailability
}

// GetFundsAvailabilityOk returns a tuple with the FundsAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetFundsAvailabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundsAvailability) {
		return nil, false
	}
	return o.FundsAvailability, true
}

// HasFundsAvailability returns a boolean if a field has been set.
func (o *CardBin) HasFundsAvailability() bool {
	if o != nil && !common.IsNil(o.FundsAvailability) {
		return true
	}

	return false
}

// SetFundsAvailability gets a reference to the given string and assigns it to the FundsAvailability field.
func (o *CardBin) SetFundsAvailability(v string) {
	o.FundsAvailability = &v
}

// GetIssuerBin returns the IssuerBin field value if set, zero value otherwise.
func (o *CardBin) GetIssuerBin() string {
	if o == nil || common.IsNil(o.IssuerBin) {
		var ret string
		return ret
	}
	return *o.IssuerBin
}

// GetIssuerBinOk returns a tuple with the IssuerBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetIssuerBinOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerBin) {
		return nil, false
	}
	return o.IssuerBin, true
}

// HasIssuerBin returns a boolean if a field has been set.
func (o *CardBin) HasIssuerBin() bool {
	if o != nil && !common.IsNil(o.IssuerBin) {
		return true
	}

	return false
}

// SetIssuerBin gets a reference to the given string and assigns it to the IssuerBin field.
func (o *CardBin) SetIssuerBin(v string) {
	o.IssuerBin = &v
}

// GetIssuingBank returns the IssuingBank field value if set, zero value otherwise.
func (o *CardBin) GetIssuingBank() string {
	if o == nil || common.IsNil(o.IssuingBank) {
		var ret string
		return ret
	}
	return *o.IssuingBank
}

// GetIssuingBankOk returns a tuple with the IssuingBank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetIssuingBankOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuingBank) {
		return nil, false
	}
	return o.IssuingBank, true
}

// HasIssuingBank returns a boolean if a field has been set.
func (o *CardBin) HasIssuingBank() bool {
	if o != nil && !common.IsNil(o.IssuingBank) {
		return true
	}

	return false
}

// SetIssuingBank gets a reference to the given string and assigns it to the IssuingBank field.
func (o *CardBin) SetIssuingBank(v string) {
	o.IssuingBank = &v
}

// GetIssuingCountry returns the IssuingCountry field value if set, zero value otherwise.
func (o *CardBin) GetIssuingCountry() string {
	if o == nil || common.IsNil(o.IssuingCountry) {
		var ret string
		return ret
	}
	return *o.IssuingCountry
}

// GetIssuingCountryOk returns a tuple with the IssuingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetIssuingCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuingCountry) {
		return nil, false
	}
	return o.IssuingCountry, true
}

// HasIssuingCountry returns a boolean if a field has been set.
func (o *CardBin) HasIssuingCountry() bool {
	if o != nil && !common.IsNil(o.IssuingCountry) {
		return true
	}

	return false
}

// SetIssuingCountry gets a reference to the given string and assigns it to the IssuingCountry field.
func (o *CardBin) SetIssuingCountry(v string) {
	o.IssuingCountry = &v
}

// GetIssuingCurrency returns the IssuingCurrency field value if set, zero value otherwise.
func (o *CardBin) GetIssuingCurrency() string {
	if o == nil || common.IsNil(o.IssuingCurrency) {
		var ret string
		return ret
	}
	return *o.IssuingCurrency
}

// GetIssuingCurrencyOk returns a tuple with the IssuingCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetIssuingCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuingCurrency) {
		return nil, false
	}
	return o.IssuingCurrency, true
}

// HasIssuingCurrency returns a boolean if a field has been set.
func (o *CardBin) HasIssuingCurrency() bool {
	if o != nil && !common.IsNil(o.IssuingCurrency) {
		return true
	}

	return false
}

// SetIssuingCurrency gets a reference to the given string and assigns it to the IssuingCurrency field.
func (o *CardBin) SetIssuingCurrency(v string) {
	o.IssuingCurrency = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *CardBin) GetPaymentMethod() string {
	if o == nil || common.IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetPaymentMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CardBin) HasPaymentMethod() bool {
	if o != nil && !common.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *CardBin) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPayoutEligible returns the PayoutEligible field value if set, zero value otherwise.
func (o *CardBin) GetPayoutEligible() string {
	if o == nil || common.IsNil(o.PayoutEligible) {
		var ret string
		return ret
	}
	return *o.PayoutEligible
}

// GetPayoutEligibleOk returns a tuple with the PayoutEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetPayoutEligibleOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayoutEligible) {
		return nil, false
	}
	return o.PayoutEligible, true
}

// HasPayoutEligible returns a boolean if a field has been set.
func (o *CardBin) HasPayoutEligible() bool {
	if o != nil && !common.IsNil(o.PayoutEligible) {
		return true
	}

	return false
}

// SetPayoutEligible gets a reference to the given string and assigns it to the PayoutEligible field.
func (o *CardBin) SetPayoutEligible(v string) {
	o.PayoutEligible = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *CardBin) GetSummary() string {
	if o == nil || common.IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBin) GetSummaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *CardBin) HasSummary() bool {
	if o != nil && !common.IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *CardBin) SetSummary(v string) {
	o.Summary = &v
}

func (o CardBin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardBin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	if !common.IsNil(o.Commercial) {
		toSerialize["commercial"] = o.Commercial
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	if !common.IsNil(o.FundsAvailability) {
		toSerialize["fundsAvailability"] = o.FundsAvailability
	}
	if !common.IsNil(o.IssuerBin) {
		toSerialize["issuerBin"] = o.IssuerBin
	}
	if !common.IsNil(o.IssuingBank) {
		toSerialize["issuingBank"] = o.IssuingBank
	}
	if !common.IsNil(o.IssuingCountry) {
		toSerialize["issuingCountry"] = o.IssuingCountry
	}
	if !common.IsNil(o.IssuingCurrency) {
		toSerialize["issuingCurrency"] = o.IssuingCurrency
	}
	if !common.IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !common.IsNil(o.PayoutEligible) {
		toSerialize["payoutEligible"] = o.PayoutEligible
	}
	if !common.IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	return toSerialize, nil
}

type NullableCardBin struct {
	value *CardBin
	isSet bool
}

func (v NullableCardBin) Get() *CardBin {
	return v.value
}

func (v *NullableCardBin) Set(val *CardBin) {
	v.value = val
	v.isSet = true
}

func (v NullableCardBin) IsSet() bool {
	return v.isSet
}

func (v *NullableCardBin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardBin(val *CardBin) *NullableCardBin {
	return &NullableCardBin{value: val, isSet: true}
}

func (v NullableCardBin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardBin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
