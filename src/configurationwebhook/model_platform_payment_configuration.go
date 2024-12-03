/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the PlatformPaymentConfiguration type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlatformPaymentConfiguration{}

// PlatformPaymentConfiguration struct for PlatformPaymentConfiguration
type PlatformPaymentConfiguration struct {
	// Specifies at what time a [sales day](https://docs.adyen.com/platforms/settle-funds/sales-day-settlement#sales-day) ends for this account.  Possible values: Time in **\"HH:MM\"** format. **HH** ranges from **00** to **07**. **MM** must be **00**.  Default value: **\"00:00\"**.
	SalesDayClosingTime *string `json:"salesDayClosingTime,omitempty"`
	// Specifies after how many business days the funds in a [settlement batch](https://docs.adyen.com/platforms/settle-funds/sales-day-settlement#settlement-batch) are made available in this balance account.  Possible values: **1** to **10**, or **null**. * Setting this value to an integer enables Sales day settlement in this balance account. See how Sales day settlement works in your [marketplace](https://docs.adyen.com/marketplaces/settle-funds/sales-day-settlement) or [platform](https://docs.adyen.com/platforms/settle-funds/sales-day-settlement). * Setting this value to **null** enables Pass-through settlement in this balance account. See how Pass-through settlement works in your [marketplace](https://docs.adyen.com/marketplaces/settle-funds/pass-through-settlement) or [platform](https://docs.adyen.com/platforms/settle-funds/pass-through-settlement).  Default value: **null**.
	SettlementDelayDays *int32 `json:"settlementDelayDays,omitempty"`
}

// NewPlatformPaymentConfiguration instantiates a new PlatformPaymentConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformPaymentConfiguration() *PlatformPaymentConfiguration {
	this := PlatformPaymentConfiguration{}
	return &this
}

// NewPlatformPaymentConfigurationWithDefaults instantiates a new PlatformPaymentConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformPaymentConfigurationWithDefaults() *PlatformPaymentConfiguration {
	this := PlatformPaymentConfiguration{}
	return &this
}

// GetSalesDayClosingTime returns the SalesDayClosingTime field value if set, zero value otherwise.
func (o *PlatformPaymentConfiguration) GetSalesDayClosingTime() string {
	if o == nil || common.IsNil(o.SalesDayClosingTime) {
		var ret string
		return ret
	}
	return *o.SalesDayClosingTime
}

// GetSalesDayClosingTimeOk returns a tuple with the SalesDayClosingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPaymentConfiguration) GetSalesDayClosingTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SalesDayClosingTime) {
		return nil, false
	}
	return o.SalesDayClosingTime, true
}

// HasSalesDayClosingTime returns a boolean if a field has been set.
func (o *PlatformPaymentConfiguration) HasSalesDayClosingTime() bool {
	if o != nil && !common.IsNil(o.SalesDayClosingTime) {
		return true
	}

	return false
}

// SetSalesDayClosingTime gets a reference to the given string and assigns it to the SalesDayClosingTime field.
func (o *PlatformPaymentConfiguration) SetSalesDayClosingTime(v string) {
	o.SalesDayClosingTime = &v
}

// GetSettlementDelayDays returns the SettlementDelayDays field value if set, zero value otherwise.
func (o *PlatformPaymentConfiguration) GetSettlementDelayDays() int32 {
	if o == nil || common.IsNil(o.SettlementDelayDays) {
		var ret int32
		return ret
	}
	return *o.SettlementDelayDays
}

// GetSettlementDelayDaysOk returns a tuple with the SettlementDelayDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPaymentConfiguration) GetSettlementDelayDaysOk() (*int32, bool) {
	if o == nil || common.IsNil(o.SettlementDelayDays) {
		return nil, false
	}
	return o.SettlementDelayDays, true
}

// HasSettlementDelayDays returns a boolean if a field has been set.
func (o *PlatformPaymentConfiguration) HasSettlementDelayDays() bool {
	if o != nil && !common.IsNil(o.SettlementDelayDays) {
		return true
	}

	return false
}

// SetSettlementDelayDays gets a reference to the given int32 and assigns it to the SettlementDelayDays field.
func (o *PlatformPaymentConfiguration) SetSettlementDelayDays(v int32) {
	o.SettlementDelayDays = &v
}

func (o PlatformPaymentConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformPaymentConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SalesDayClosingTime) {
		toSerialize["salesDayClosingTime"] = o.SalesDayClosingTime
	}
	if !common.IsNil(o.SettlementDelayDays) {
		toSerialize["settlementDelayDays"] = o.SettlementDelayDays
	}
	return toSerialize, nil
}

type NullablePlatformPaymentConfiguration struct {
	value *PlatformPaymentConfiguration
	isSet bool
}

func (v NullablePlatformPaymentConfiguration) Get() *PlatformPaymentConfiguration {
	return v.value
}

func (v *NullablePlatformPaymentConfiguration) Set(val *PlatformPaymentConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformPaymentConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformPaymentConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformPaymentConfiguration(val *PlatformPaymentConfiguration) *NullablePlatformPaymentConfiguration {
	return &NullablePlatformPaymentConfiguration{value: val, isSet: true}
}

func (v NullablePlatformPaymentConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformPaymentConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
