/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the SplitConfigurationRule type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SplitConfigurationRule{}

// SplitConfigurationRule struct for SplitConfigurationRule
type SplitConfigurationRule struct {
	// The currency condition that defines whether the split logic applies. Its value must be a three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	Currency string `json:"currency"`
	// The funding source condition of the payment method (only for cards).  Possible values: **credit**, **debit**, or **ANY**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// The payment method condition that defines whether the split logic applies.  Possible values: * [Payment method variant](https://docs.adyen.com/development-resources/paymentmethodvariant): Apply the split logic for a specific payment method. * **ANY**: Apply the split logic for all available payment methods.
	PaymentMethod string `json:"paymentMethod"`
	// The unique identifier of the split configuration rule.
	RuleId *string `json:"ruleId,omitempty"`
	// The sales channel condition that defines whether the split logic applies.  Possible values: * **Ecommerce**: Online transactions where the cardholder is present. * **ContAuth**: Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). * **Moto**: Mail-order and telephone-order transactions where the customer is in contact with the merchant via email or telephone. * **POS**: Point-of-sale transactions where the customer is physically present to make a payment using a secure payment terminal. * **ANY**: All sales channels.
	ShopperInteraction string                  `json:"shopperInteraction"`
	SplitLogic         SplitConfigurationLogic `json:"splitLogic"`
}

// NewSplitConfigurationRule instantiates a new SplitConfigurationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitConfigurationRule(currency string, paymentMethod string, shopperInteraction string, splitLogic SplitConfigurationLogic) *SplitConfigurationRule {
	this := SplitConfigurationRule{}
	this.Currency = currency
	this.PaymentMethod = paymentMethod
	this.ShopperInteraction = shopperInteraction
	this.SplitLogic = splitLogic
	return &this
}

// NewSplitConfigurationRuleWithDefaults instantiates a new SplitConfigurationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitConfigurationRuleWithDefaults() *SplitConfigurationRule {
	this := SplitConfigurationRule{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *SplitConfigurationRule) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SplitConfigurationRule) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SplitConfigurationRule) SetCurrency(v string) {
	o.Currency = v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *SplitConfigurationRule) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationRule) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *SplitConfigurationRule) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *SplitConfigurationRule) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *SplitConfigurationRule) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *SplitConfigurationRule) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *SplitConfigurationRule) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *SplitConfigurationRule) GetRuleId() string {
	if o == nil || common.IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationRule) GetRuleIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *SplitConfigurationRule) HasRuleId() bool {
	if o != nil && !common.IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *SplitConfigurationRule) SetRuleId(v string) {
	o.RuleId = &v
}

// GetShopperInteraction returns the ShopperInteraction field value
func (o *SplitConfigurationRule) GetShopperInteraction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperInteraction
}

// GetShopperInteractionOk returns a tuple with the ShopperInteraction field value
// and a boolean to check if the value has been set.
func (o *SplitConfigurationRule) GetShopperInteractionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperInteraction, true
}

// SetShopperInteraction sets field value
func (o *SplitConfigurationRule) SetShopperInteraction(v string) {
	o.ShopperInteraction = v
}

// GetSplitLogic returns the SplitLogic field value
func (o *SplitConfigurationRule) GetSplitLogic() SplitConfigurationLogic {
	if o == nil {
		var ret SplitConfigurationLogic
		return ret
	}

	return o.SplitLogic
}

// GetSplitLogicOk returns a tuple with the SplitLogic field value
// and a boolean to check if the value has been set.
func (o *SplitConfigurationRule) GetSplitLogicOk() (*SplitConfigurationLogic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitLogic, true
}

// SetSplitLogic sets field value
func (o *SplitConfigurationRule) SetSplitLogic(v SplitConfigurationLogic) {
	o.SplitLogic = v
}

func (o SplitConfigurationRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitConfigurationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	toSerialize["paymentMethod"] = o.PaymentMethod
	if !common.IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	toSerialize["shopperInteraction"] = o.ShopperInteraction
	toSerialize["splitLogic"] = o.SplitLogic
	return toSerialize, nil
}

type NullableSplitConfigurationRule struct {
	value *SplitConfigurationRule
	isSet bool
}

func (v NullableSplitConfigurationRule) Get() *SplitConfigurationRule {
	return v.value
}

func (v *NullableSplitConfigurationRule) Set(val *SplitConfigurationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitConfigurationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitConfigurationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitConfigurationRule(val *SplitConfigurationRule) *NullableSplitConfigurationRule {
	return &NullableSplitConfigurationRule{value: val, isSet: true}
}

func (v NullableSplitConfigurationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitConfigurationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SplitConfigurationRule) isValidFundingSource() bool {
	var allowedEnumValues = []string{"credit", "debit", "ANY"}
	for _, allowed := range allowedEnumValues {
		if o.GetFundingSource() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationRule) isValidShopperInteraction() bool {
	var allowedEnumValues = []string{"Ecommerce", "ContAuth", "Moto", "POS", "ANY"}
	for _, allowed := range allowedEnumValues {
		if o.GetShopperInteraction() == allowed {
			return true
		}
	}
	return false
}
