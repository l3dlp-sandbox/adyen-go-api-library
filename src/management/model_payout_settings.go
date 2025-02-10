/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the PayoutSettings type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayoutSettings{}

// PayoutSettings struct for PayoutSettings
type PayoutSettings struct {
	// Indicates if payouts to the bank account are allowed. This value is set automatically based on the status of the verification process. The value is:  * **true** if `verificationStatus` is **valid**. * **false** for all other values.
	Allowed *bool `json:"allowed,omitempty"`
	// Indicates if payouts to this bank account are enabled. Default: **true**.  To receive payouts into this bank account, both `enabled` and `allowed` must be **true**.
	Enabled *bool `json:"enabled,omitempty"`
	// The date when Adyen starts paying out to this bank account.  Format: [ISO 8601](https://www.w3.org/TR/NOTE-datetime), for example, **2019-11-23T12:25:28Z** or **2020-05-27T20:25:28+08:00**.  If not specified, the `enabled` field indicates if payouts are enabled for this bank account.  If a date is specified and:  * `enabled`: **true**, payouts are enabled starting the specified date. * `enabled`: **false**, payouts are disabled until the specified date. On the specified date, `enabled` changes to **true** and this field is reset to **null**.
	EnabledFromDate *string `json:"enabledFromDate,omitempty"`
	// The unique identifier of the payout setting.
	Id string `json:"id"`
	// Determines how long it takes for the funds to reach the bank account. Adyen pays out based on the [payout frequency](https://docs.adyen.com/account/getting-paid#payout-frequency). Depending on the currencies and banks involved in transferring the money, it may take up to three days for the payout funds to arrive in the bank account.   Possible values: * **first**: same day. * **urgent**: the next day. * **normal**: between 1 and 3 days.
	Priority *string `json:"priority,omitempty"`
	// The unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments) that contains the details of the bank account.
	TransferInstrumentId string `json:"transferInstrumentId"`
	// The status of the verification process for the bank account.  Possible values: * **valid**: the verification was successful. * **pending**: the verification is in progress. * **invalid**: the information provided is not complete. * **rejected**:  there are reasons to refuse working with this entity.
	VerificationStatus *string `json:"verificationStatus,omitempty"`
}

// NewPayoutSettings instantiates a new PayoutSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutSettings(id string, transferInstrumentId string) *PayoutSettings {
	this := PayoutSettings{}
	this.Id = id
	this.TransferInstrumentId = transferInstrumentId
	return &this
}

// NewPayoutSettingsWithDefaults instantiates a new PayoutSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutSettingsWithDefaults() *PayoutSettings {
	this := PayoutSettings{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *PayoutSettings) GetAllowed() bool {
	if o == nil || common.IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettings) GetAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *PayoutSettings) HasAllowed() bool {
	if o != nil && !common.IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *PayoutSettings) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PayoutSettings) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PayoutSettings) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PayoutSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnabledFromDate returns the EnabledFromDate field value if set, zero value otherwise.
func (o *PayoutSettings) GetEnabledFromDate() string {
	if o == nil || common.IsNil(o.EnabledFromDate) {
		var ret string
		return ret
	}
	return *o.EnabledFromDate
}

// GetEnabledFromDateOk returns a tuple with the EnabledFromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettings) GetEnabledFromDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnabledFromDate) {
		return nil, false
	}
	return o.EnabledFromDate, true
}

// HasEnabledFromDate returns a boolean if a field has been set.
func (o *PayoutSettings) HasEnabledFromDate() bool {
	if o != nil && !common.IsNil(o.EnabledFromDate) {
		return true
	}

	return false
}

// SetEnabledFromDate gets a reference to the given string and assigns it to the EnabledFromDate field.
func (o *PayoutSettings) SetEnabledFromDate(v string) {
	o.EnabledFromDate = &v
}

// GetId returns the Id field value
func (o *PayoutSettings) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PayoutSettings) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PayoutSettings) SetId(v string) {
	o.Id = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PayoutSettings) GetPriority() string {
	if o == nil || common.IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettings) GetPriorityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PayoutSettings) HasPriority() bool {
	if o != nil && !common.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *PayoutSettings) SetPriority(v string) {
	o.Priority = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value
func (o *PayoutSettings) GetTransferInstrumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value
// and a boolean to check if the value has been set.
func (o *PayoutSettings) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferInstrumentId, true
}

// SetTransferInstrumentId sets field value
func (o *PayoutSettings) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *PayoutSettings) GetVerificationStatus() string {
	if o == nil || common.IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettings) GetVerificationStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *PayoutSettings) HasVerificationStatus() bool {
	if o != nil && !common.IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *PayoutSettings) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o PayoutSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.EnabledFromDate) {
		toSerialize["enabledFromDate"] = o.EnabledFromDate
	}
	toSerialize["id"] = o.Id
	if !common.IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	if !common.IsNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullablePayoutSettings struct {
	value *PayoutSettings
	isSet bool
}

func (v NullablePayoutSettings) Get() *PayoutSettings {
	return v.value
}

func (v *NullablePayoutSettings) Set(val *PayoutSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutSettings(val *PayoutSettings) *NullablePayoutSettings {
	return &NullablePayoutSettings{value: val, isSet: true}
}

func (v NullablePayoutSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PayoutSettings) isValidPriority() bool {
	var allowedEnumValues = []string{"first", "normal", "urgent"}
	for _, allowed := range allowedEnumValues {
		if o.GetPriority() == allowed {
			return true
		}
	}
	return false
}
func (o *PayoutSettings) isValidVerificationStatus() bool {
	var allowedEnumValues = []string{"invalid", "pending", "rejected", "valid"}
	for _, allowed := range allowedEnumValues {
		if o.GetVerificationStatus() == allowed {
			return true
		}
	}
	return false
}
