/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the UpdateSweepConfigurationV2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateSweepConfigurationV2{}

// UpdateSweepConfigurationV2 struct for UpdateSweepConfigurationV2
type UpdateSweepConfigurationV2 struct {
	// The type of transfer that results from the sweep.  Possible values:   - **bank**: Sweep to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id).  - **internal**: Transfer to another [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  Required when setting `priorities`.
	Category     *string            `json:"category,omitempty"`
	Counterparty *SweepCounterparty `json:"counterparty,omitempty"`
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) in uppercase. For example, **EUR**.  The sweep currency must match any of the [balances currencies](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balanceAccounts/{id}__resParam_balances).
	Currency *string `json:"currency,omitempty"`
	// The message that will be used in the sweep transfer's description body with a maximum length of 140 characters.  If the message is longer after replacing placeholders, the message will be cut off at 140 characters.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the sweep.
	Id *string `json:"id,omitempty"`
	// The list of priorities for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. You can provide multiple priorities. Adyen will try to pay out using the priority you list first. If that's not possible, it moves on to the next option in the order of your provided priorities.  Possible values:  * **regular**: for normal, low-value transactions.  * **fast**: a faster way to transfer funds, but the fees are higher. Recommended for high-priority, low-value transactions.  * **wire**: the fastest way to transfer funds, but this has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: for instant funds transfers in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: for high-value transfers to a recipient in a different country.  * **internal**: for transfers to an Adyen-issued business bank account (by bank account number/IBAN).  Set `category` to **bank**. For more details, see optional priorities setup for [marketplaces](https://docs.adyen.com/marketplaces/payout-to-users/scheduled-payouts#optional-priorities-setup) or [platforms](https://docs.adyen.com/platforms/payout-to-users/scheduled-payouts#optional-priorities-setup).
	Priorities []string `json:"priorities,omitempty"`
	// The reason for disabling the sweep.
	Reason *string `json:"reason,omitempty"`
	// Your reference for the sweep configuration.
	Reference *string `json:"reference,omitempty"`
	// The reference sent to or received from the counterparty. Only alphanumeric characters are allowed.
	ReferenceForBeneficiary *string        `json:"referenceForBeneficiary,omitempty"`
	Schedule                *SweepSchedule `json:"schedule,omitempty"`
	// The status of the sweep. If not provided, by default, this is set to **active**.  Possible values:    * **active**:  the sweep is enabled and funds will be pulled in or pushed out based on the defined configuration.    * **inactive**: the sweep is disabled and cannot be triggered.
	Status        *string `json:"status,omitempty"`
	SweepAmount   *Amount `json:"sweepAmount,omitempty"`
	TargetAmount  *Amount `json:"targetAmount,omitempty"`
	TriggerAmount *Amount `json:"triggerAmount,omitempty"`
	// The direction of sweep, whether pushing out or pulling in funds to the balance account. If not provided, by default, this is set to **push**.  Possible values:   * **push**: _push out funds_ to a destination balance account or transfer instrument.   * **pull**: _pull in funds_ from a source merchant account, transfer instrument, or balance account.
	Type *string `json:"type,omitempty"`
}

// NewUpdateSweepConfigurationV2 instantiates a new UpdateSweepConfigurationV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSweepConfigurationV2() *UpdateSweepConfigurationV2 {
	this := UpdateSweepConfigurationV2{}
	var type_ string = "push"
	this.Type = &type_
	return &this
}

// NewUpdateSweepConfigurationV2WithDefaults instantiates a new UpdateSweepConfigurationV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSweepConfigurationV2WithDefaults() *UpdateSweepConfigurationV2 {
	this := UpdateSweepConfigurationV2{}
	var type_ string = "push"
	this.Type = &type_
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetCategory() string {
	if o == nil || common.IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetCategoryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasCategory() bool {
	if o != nil && !common.IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *UpdateSweepConfigurationV2) SetCategory(v string) {
	o.Category = &v
}

// GetCounterparty returns the Counterparty field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetCounterparty() SweepCounterparty {
	if o == nil || common.IsNil(o.Counterparty) {
		var ret SweepCounterparty
		return ret
	}
	return *o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetCounterpartyOk() (*SweepCounterparty, bool) {
	if o == nil || common.IsNil(o.Counterparty) {
		return nil, false
	}
	return o.Counterparty, true
}

// HasCounterparty returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasCounterparty() bool {
	if o != nil && !common.IsNil(o.Counterparty) {
		return true
	}

	return false
}

// SetCounterparty gets a reference to the given SweepCounterparty and assigns it to the Counterparty field.
func (o *UpdateSweepConfigurationV2) SetCounterparty(v SweepCounterparty) {
	o.Counterparty = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UpdateSweepConfigurationV2) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateSweepConfigurationV2) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateSweepConfigurationV2) SetId(v string) {
	o.Id = &v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetPriorities() []string {
	if o == nil || common.IsNil(o.Priorities) {
		var ret []string
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetPrioritiesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasPriorities() bool {
	if o != nil && !common.IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []string and assigns it to the Priorities field.
func (o *UpdateSweepConfigurationV2) SetPriorities(v []string) {
	o.Priorities = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UpdateSweepConfigurationV2) SetReason(v string) {
	o.Reason = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *UpdateSweepConfigurationV2) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceForBeneficiary returns the ReferenceForBeneficiary field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetReferenceForBeneficiary() string {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		var ret string
		return ret
	}
	return *o.ReferenceForBeneficiary
}

// GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetReferenceForBeneficiaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		return nil, false
	}
	return o.ReferenceForBeneficiary, true
}

// HasReferenceForBeneficiary returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasReferenceForBeneficiary() bool {
	if o != nil && !common.IsNil(o.ReferenceForBeneficiary) {
		return true
	}

	return false
}

// SetReferenceForBeneficiary gets a reference to the given string and assigns it to the ReferenceForBeneficiary field.
func (o *UpdateSweepConfigurationV2) SetReferenceForBeneficiary(v string) {
	o.ReferenceForBeneficiary = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetSchedule() SweepSchedule {
	if o == nil || common.IsNil(o.Schedule) {
		var ret SweepSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetScheduleOk() (*SweepSchedule, bool) {
	if o == nil || common.IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasSchedule() bool {
	if o != nil && !common.IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given SweepSchedule and assigns it to the Schedule field.
func (o *UpdateSweepConfigurationV2) SetSchedule(v SweepSchedule) {
	o.Schedule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateSweepConfigurationV2) SetStatus(v string) {
	o.Status = &v
}

// GetSweepAmount returns the SweepAmount field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetSweepAmount() Amount {
	if o == nil || common.IsNil(o.SweepAmount) {
		var ret Amount
		return ret
	}
	return *o.SweepAmount
}

// GetSweepAmountOk returns a tuple with the SweepAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetSweepAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.SweepAmount) {
		return nil, false
	}
	return o.SweepAmount, true
}

// HasSweepAmount returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasSweepAmount() bool {
	if o != nil && !common.IsNil(o.SweepAmount) {
		return true
	}

	return false
}

// SetSweepAmount gets a reference to the given Amount and assigns it to the SweepAmount field.
func (o *UpdateSweepConfigurationV2) SetSweepAmount(v Amount) {
	o.SweepAmount = &v
}

// GetTargetAmount returns the TargetAmount field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetTargetAmount() Amount {
	if o == nil || common.IsNil(o.TargetAmount) {
		var ret Amount
		return ret
	}
	return *o.TargetAmount
}

// GetTargetAmountOk returns a tuple with the TargetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetTargetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.TargetAmount) {
		return nil, false
	}
	return o.TargetAmount, true
}

// HasTargetAmount returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasTargetAmount() bool {
	if o != nil && !common.IsNil(o.TargetAmount) {
		return true
	}

	return false
}

// SetTargetAmount gets a reference to the given Amount and assigns it to the TargetAmount field.
func (o *UpdateSweepConfigurationV2) SetTargetAmount(v Amount) {
	o.TargetAmount = &v
}

// GetTriggerAmount returns the TriggerAmount field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetTriggerAmount() Amount {
	if o == nil || common.IsNil(o.TriggerAmount) {
		var ret Amount
		return ret
	}
	return *o.TriggerAmount
}

// GetTriggerAmountOk returns a tuple with the TriggerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetTriggerAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.TriggerAmount) {
		return nil, false
	}
	return o.TriggerAmount, true
}

// HasTriggerAmount returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasTriggerAmount() bool {
	if o != nil && !common.IsNil(o.TriggerAmount) {
		return true
	}

	return false
}

// SetTriggerAmount gets a reference to the given Amount and assigns it to the TriggerAmount field.
func (o *UpdateSweepConfigurationV2) SetTriggerAmount(v Amount) {
	o.TriggerAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateSweepConfigurationV2) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSweepConfigurationV2) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateSweepConfigurationV2) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateSweepConfigurationV2) SetType(v string) {
	o.Type = &v
}

func (o UpdateSweepConfigurationV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSweepConfigurationV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !common.IsNil(o.Counterparty) {
		toSerialize["counterparty"] = o.Counterparty
	}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Priorities) {
		toSerialize["priorities"] = o.Priorities
	}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.ReferenceForBeneficiary) {
		toSerialize["referenceForBeneficiary"] = o.ReferenceForBeneficiary
	}
	if !common.IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.SweepAmount) {
		toSerialize["sweepAmount"] = o.SweepAmount
	}
	if !common.IsNil(o.TargetAmount) {
		toSerialize["targetAmount"] = o.TargetAmount
	}
	if !common.IsNil(o.TriggerAmount) {
		toSerialize["triggerAmount"] = o.TriggerAmount
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUpdateSweepConfigurationV2 struct {
	value *UpdateSweepConfigurationV2
	isSet bool
}

func (v NullableUpdateSweepConfigurationV2) Get() *UpdateSweepConfigurationV2 {
	return v.value
}

func (v *NullableUpdateSweepConfigurationV2) Set(val *UpdateSweepConfigurationV2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSweepConfigurationV2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSweepConfigurationV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSweepConfigurationV2(val *UpdateSweepConfigurationV2) *NullableUpdateSweepConfigurationV2 {
	return &NullableUpdateSweepConfigurationV2{value: val, isSet: true}
}

func (v NullableUpdateSweepConfigurationV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSweepConfigurationV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpdateSweepConfigurationV2) isValidCategory() bool {
	var allowedEnumValues = []string{"bank", "internal", "platformPayment"}
	for _, allowed := range allowedEnumValues {
		if o.GetCategory() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSweepConfigurationV2) isValidReason() bool {
	var allowedEnumValues = []string{"amountLimitExceeded", "approved", "balanceAccountTemporarilyBlockedByTransactionRule", "counterpartyAccountBlocked", "counterpartyAccountClosed", "counterpartyAccountNotFound", "counterpartyAddressRequired", "counterpartyBankTimedOut", "counterpartyBankUnavailable", "declined", "declinedByTransactionRule", "error", "notEnoughBalance", "pendingApproval", "refusedByCounterpartyBank", "routeNotFound", "scaFailed", "unknown"}
	for _, allowed := range allowedEnumValues {
		if o.GetReason() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSweepConfigurationV2) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "inactive"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSweepConfigurationV2) isValidType() bool {
	var allowedEnumValues = []string{"pull", "push"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
