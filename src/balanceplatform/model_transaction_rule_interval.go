/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the TransactionRuleInterval type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRuleInterval{}

// TransactionRuleInterval struct for TransactionRuleInterval
type TransactionRuleInterval struct {
	// The day of month, used when the `duration.unit` is **months**. If not provided, by default, this is set to **1**, the first day of the month.
	DayOfMonth *int32 `json:"dayOfMonth,omitempty"`
	// The day of week, used when the `duration.unit` is **weeks**. If not provided, by default, this is set to **monday**.  Possible values: **sunday**, **monday**, **tuesday**, **wednesday**, **thursday**, **friday**.
	DayOfWeek *string   `json:"dayOfWeek,omitempty"`
	Duration  *Duration `json:"duration,omitempty"`
	// The time of day, in **hh:mm:ss** format, used when the `duration.unit` is **hours**. If not provided, by default, this is set to **00:00:00**.
	TimeOfDay *string `json:"timeOfDay,omitempty"`
	// The [time zone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). For example, **Europe/Amsterdam**. By default, this is set to **UTC**.
	TimeZone *string `json:"timeZone,omitempty"`
	// The [type of interval](https://docs.adyen.com/issuing/transaction-rules#time-intervals) during which the rule conditions and limits apply, and how often counters are reset.  Possible values:   * **perTransaction**: conditions are evaluated and the counters are reset for every transaction.  * **daily**: the counters are reset daily at 00:00:00 CET.  * **weekly**: the counters are reset every Monday at 00:00:00 CET.   * **monthly**: the counters reset every first day of the month at 00:00:00 CET.   * **lifetime**: conditions are applied to the lifetime of the payment instrument.  * **rolling**: conditions are applied and the counters are reset based on a `duration`. If the reset date and time are not provided, Adyen applies the default reset time similar to fixed intervals. For example, if the duration is every two weeks, the counter resets every third Monday at 00:00:00 CET.  * **sliding**: conditions are applied and the counters are reset based on the current time and a `duration` that you specify.
	Type string `json:"type"`
}

// NewTransactionRuleInterval instantiates a new TransactionRuleInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleInterval(type_ string) *TransactionRuleInterval {
	this := TransactionRuleInterval{}
	this.Type = type_
	return &this
}

// NewTransactionRuleIntervalWithDefaults instantiates a new TransactionRuleInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleIntervalWithDefaults() *TransactionRuleInterval {
	this := TransactionRuleInterval{}
	return &this
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *TransactionRuleInterval) GetDayOfMonth() int32 {
	if o == nil || common.IsNil(o.DayOfMonth) {
		var ret int32
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInterval) GetDayOfMonthOk() (*int32, bool) {
	if o == nil || common.IsNil(o.DayOfMonth) {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *TransactionRuleInterval) HasDayOfMonth() bool {
	if o != nil && !common.IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int32 and assigns it to the DayOfMonth field.
func (o *TransactionRuleInterval) SetDayOfMonth(v int32) {
	o.DayOfMonth = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *TransactionRuleInterval) GetDayOfWeek() string {
	if o == nil || common.IsNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInterval) GetDayOfWeekOk() (*string, bool) {
	if o == nil || common.IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *TransactionRuleInterval) HasDayOfWeek() bool {
	if o != nil && !common.IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *TransactionRuleInterval) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *TransactionRuleInterval) GetDuration() Duration {
	if o == nil || common.IsNil(o.Duration) {
		var ret Duration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInterval) GetDurationOk() (*Duration, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *TransactionRuleInterval) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given Duration and assigns it to the Duration field.
func (o *TransactionRuleInterval) SetDuration(v Duration) {
	o.Duration = &v
}

// GetTimeOfDay returns the TimeOfDay field value if set, zero value otherwise.
func (o *TransactionRuleInterval) GetTimeOfDay() string {
	if o == nil || common.IsNil(o.TimeOfDay) {
		var ret string
		return ret
	}
	return *o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInterval) GetTimeOfDayOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeOfDay) {
		return nil, false
	}
	return o.TimeOfDay, true
}

// HasTimeOfDay returns a boolean if a field has been set.
func (o *TransactionRuleInterval) HasTimeOfDay() bool {
	if o != nil && !common.IsNil(o.TimeOfDay) {
		return true
	}

	return false
}

// SetTimeOfDay gets a reference to the given string and assigns it to the TimeOfDay field.
func (o *TransactionRuleInterval) SetTimeOfDay(v string) {
	o.TimeOfDay = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *TransactionRuleInterval) GetTimeZone() string {
	if o == nil || common.IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInterval) GetTimeZoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TransactionRuleInterval) HasTimeZone() bool {
	if o != nil && !common.IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *TransactionRuleInterval) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetType returns the Type field value
func (o *TransactionRuleInterval) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionRuleInterval) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionRuleInterval) SetType(v string) {
	o.Type = v
}

func (o TransactionRuleInterval) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DayOfMonth) {
		toSerialize["dayOfMonth"] = o.DayOfMonth
	}
	if !common.IsNil(o.DayOfWeek) {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.TimeOfDay) {
		toSerialize["timeOfDay"] = o.TimeOfDay
	}
	if !common.IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransactionRuleInterval struct {
	value *TransactionRuleInterval
	isSet bool
}

func (v NullableTransactionRuleInterval) Get() *TransactionRuleInterval {
	return v.value
}

func (v *NullableTransactionRuleInterval) Set(val *TransactionRuleInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleInterval(val *TransactionRuleInterval) *NullableTransactionRuleInterval {
	return &NullableTransactionRuleInterval{value: val, isSet: true}
}

func (v NullableTransactionRuleInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransactionRuleInterval) isValidDayOfWeek() bool {
	var allowedEnumValues = []string{"friday", "monday", "saturday", "sunday", "thursday", "tuesday", "wednesday"}
	for _, allowed := range allowedEnumValues {
		if o.GetDayOfWeek() == allowed {
			return true
		}
	}
	return false
}
func (o *TransactionRuleInterval) isValidType() bool {
	var allowedEnumValues = []string{"daily", "lifetime", "monthly", "perTransaction", "rolling", "sliding", "weekly"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
