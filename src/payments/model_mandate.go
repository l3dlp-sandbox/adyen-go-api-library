/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Mandate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Mandate{}

// Mandate struct for Mandate
type Mandate struct {
	// The billing amount (in minor units) of the recurring transactions.
	Amount string `json:"amount"`
	// The limitation rule of the billing amount.  Possible values:  * **max**: The transaction amount can not exceed the `amount`.   * **exact**: The transaction amount should be the same as the `amount`.  
	AmountRule *string `json:"amountRule,omitempty"`
	// The rule to specify the period, within which the recurring debit can happen, relative to the mandate recurring date.  Possible values:   * **on**: On a specific date.   * **before**:  Before and on a specific date.   * **after**: On and after a specific date.  
	BillingAttemptsRule *string `json:"billingAttemptsRule,omitempty"`
	// The number of the day, on which the recurring debit can happen. Should be within the same calendar month as the mandate recurring date.  Possible values: 1-31 based on the `frequency`.
	BillingDay *string `json:"billingDay,omitempty"`
	// The number of transactions that can be performed within the given frequency.
	Count *string `json:"count,omitempty"`
	// End date of the billing plan, in YYYY-MM-DD format.
	EndsAt string `json:"endsAt"`
	// The frequency with which a shopper should be charged.  Possible values: **daily**, **weekly**, **biWeekly**, **monthly**, **quarterly**, **halfYearly**, **yearly**.
	Frequency string `json:"frequency"`
	// The message shown by UPI to the shopper on the approval screen.
	Remarks *string `json:"remarks,omitempty"`
	// Start date of the billing plan, in YYYY-MM-DD format. By default, the transaction date.
	StartsAt *string `json:"startsAt,omitempty"`
}

// NewMandate instantiates a new Mandate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMandate(amount string, endsAt string, frequency string) *Mandate {
	this := Mandate{}
	this.Amount = amount
	this.EndsAt = endsAt
	this.Frequency = frequency
	return &this
}

// NewMandateWithDefaults instantiates a new Mandate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMandateWithDefaults() *Mandate {
	this := Mandate{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Mandate) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Mandate) SetAmount(v string) {
	o.Amount = v
}

// GetAmountRule returns the AmountRule field value if set, zero value otherwise.
func (o *Mandate) GetAmountRule() string {
	if o == nil || common.IsNil(o.AmountRule) {
		var ret string
		return ret
	}
	return *o.AmountRule
}

// GetAmountRuleOk returns a tuple with the AmountRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mandate) GetAmountRuleOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountRule) {
		return nil, false
	}
	return o.AmountRule, true
}

// HasAmountRule returns a boolean if a field has been set.
func (o *Mandate) HasAmountRule() bool {
	if o != nil && !common.IsNil(o.AmountRule) {
		return true
	}

	return false
}

// SetAmountRule gets a reference to the given string and assigns it to the AmountRule field.
func (o *Mandate) SetAmountRule(v string) {
	o.AmountRule = &v
}

// GetBillingAttemptsRule returns the BillingAttemptsRule field value if set, zero value otherwise.
func (o *Mandate) GetBillingAttemptsRule() string {
	if o == nil || common.IsNil(o.BillingAttemptsRule) {
		var ret string
		return ret
	}
	return *o.BillingAttemptsRule
}

// GetBillingAttemptsRuleOk returns a tuple with the BillingAttemptsRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mandate) GetBillingAttemptsRuleOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAttemptsRule) {
		return nil, false
	}
	return o.BillingAttemptsRule, true
}

// HasBillingAttemptsRule returns a boolean if a field has been set.
func (o *Mandate) HasBillingAttemptsRule() bool {
	if o != nil && !common.IsNil(o.BillingAttemptsRule) {
		return true
	}

	return false
}

// SetBillingAttemptsRule gets a reference to the given string and assigns it to the BillingAttemptsRule field.
func (o *Mandate) SetBillingAttemptsRule(v string) {
	o.BillingAttemptsRule = &v
}

// GetBillingDay returns the BillingDay field value if set, zero value otherwise.
func (o *Mandate) GetBillingDay() string {
	if o == nil || common.IsNil(o.BillingDay) {
		var ret string
		return ret
	}
	return *o.BillingDay
}

// GetBillingDayOk returns a tuple with the BillingDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mandate) GetBillingDayOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingDay) {
		return nil, false
	}
	return o.BillingDay, true
}

// HasBillingDay returns a boolean if a field has been set.
func (o *Mandate) HasBillingDay() bool {
	if o != nil && !common.IsNil(o.BillingDay) {
		return true
	}

	return false
}

// SetBillingDay gets a reference to the given string and assigns it to the BillingDay field.
func (o *Mandate) SetBillingDay(v string) {
	o.BillingDay = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Mandate) GetCount() string {
	if o == nil || common.IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mandate) GetCountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Mandate) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *Mandate) SetCount(v string) {
	o.Count = &v
}

// GetEndsAt returns the EndsAt field value
func (o *Mandate) GetEndsAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetEndsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndsAt, true
}

// SetEndsAt sets field value
func (o *Mandate) SetEndsAt(v string) {
	o.EndsAt = v
}

// GetFrequency returns the Frequency field value
func (o *Mandate) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *Mandate) SetFrequency(v string) {
	o.Frequency = v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *Mandate) GetRemarks() string {
	if o == nil || common.IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mandate) GetRemarksOk() (*string, bool) {
	if o == nil || common.IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *Mandate) HasRemarks() bool {
	if o != nil && !common.IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *Mandate) SetRemarks(v string) {
	o.Remarks = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *Mandate) GetStartsAt() string {
	if o == nil || common.IsNil(o.StartsAt) {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mandate) GetStartsAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *Mandate) HasStartsAt() bool {
	if o != nil && !common.IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *Mandate) SetStartsAt(v string) {
	o.StartsAt = &v
}

func (o Mandate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mandate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.AmountRule) {
		toSerialize["amountRule"] = o.AmountRule
	}
	if !common.IsNil(o.BillingAttemptsRule) {
		toSerialize["billingAttemptsRule"] = o.BillingAttemptsRule
	}
	if !common.IsNil(o.BillingDay) {
		toSerialize["billingDay"] = o.BillingDay
	}
	if !common.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["endsAt"] = o.EndsAt
	toSerialize["frequency"] = o.Frequency
	if !common.IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !common.IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	return toSerialize, nil
}

type NullableMandate struct {
	value *Mandate
	isSet bool
}

func (v NullableMandate) Get() *Mandate {
	return v.value
}

func (v *NullableMandate) Set(val *Mandate) {
	v.value = val
	v.isSet = true
}

func (v NullableMandate) IsSet() bool {
	return v.isSet
}

func (v *NullableMandate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMandate(val *Mandate) *NullableMandate {
	return &NullableMandate{value: val, isSet: true}
}

func (v NullableMandate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMandate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *Mandate) isValidAmountRule() bool {
    var allowedEnumValues = []string{ "max", "exact" }
    for _, allowed := range allowedEnumValues {
        if o.GetAmountRule() == allowed {
            return true
        }
    }
    return false
}
func (o *Mandate) isValidBillingAttemptsRule() bool {
    var allowedEnumValues = []string{ "on", "before", "after" }
    for _, allowed := range allowedEnumValues {
        if o.GetBillingAttemptsRule() == allowed {
            return true
        }
    }
    return false
}
func (o *Mandate) isValidFrequency() bool {
    var allowedEnumValues = []string{ "adhoc", "daily", "weekly", "biWeekly", "monthly", "quarterly", "halfYearly", "yearly" }
    for _, allowed := range allowedEnumValues {
        if o.GetFrequency() == allowed {
            return true
        }
    }
    return false
}

