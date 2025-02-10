/*
Negative balance compensation warning

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package negativebalancewarningwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the NegativeBalanceCompensationWarningNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NegativeBalanceCompensationWarningNotificationData{}

// NegativeBalanceCompensationWarningNotificationData struct for NegativeBalanceCompensationWarningNotificationData
type NegativeBalanceCompensationWarningNotificationData struct {
	AccountHolder *ResourceReference `json:"accountHolder,omitempty"`
	Amount        *Amount            `json:"amount,omitempty"`
	// The unique identifier of the balance platform.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
	// The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// The ID of the resource.
	Id *string `json:"id,omitempty"`
	// The balance account ID of the account that will be used to compensate the balance account whose balance is negative.
	LiableBalanceAccountId *string `json:"liableBalanceAccountId,omitempty"`
	// The date the balance for the account became negative.
	NegativeBalanceSince *time.Time `json:"negativeBalanceSince,omitempty"`
	// The date when a compensation transfer to the account is scheduled to happen.
	ScheduledCompensationAt *time.Time `json:"scheduledCompensationAt,omitempty"`
}

// NewNegativeBalanceCompensationWarningNotificationData instantiates a new NegativeBalanceCompensationWarningNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegativeBalanceCompensationWarningNotificationData() *NegativeBalanceCompensationWarningNotificationData {
	this := NegativeBalanceCompensationWarningNotificationData{}
	return &this
}

// NewNegativeBalanceCompensationWarningNotificationDataWithDefaults instantiates a new NegativeBalanceCompensationWarningNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegativeBalanceCompensationWarningNotificationDataWithDefaults() *NegativeBalanceCompensationWarningNotificationData {
	this := NegativeBalanceCompensationWarningNotificationData{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetAccountHolder() ResourceReference {
	if o == nil || common.IsNil(o.AccountHolder) {
		var ret ResourceReference
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetAccountHolderOk() (*ResourceReference, bool) {
	if o == nil || common.IsNil(o.AccountHolder) {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasAccountHolder() bool {
	if o != nil && !common.IsNil(o.AccountHolder) {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given ResourceReference and assigns it to the AccountHolder field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetAccountHolder(v ResourceReference) {
	o.AccountHolder = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetAmount(v Amount) {
	o.Amount = &v
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetId(v string) {
	o.Id = &v
}

// GetLiableBalanceAccountId returns the LiableBalanceAccountId field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetLiableBalanceAccountId() string {
	if o == nil || common.IsNil(o.LiableBalanceAccountId) {
		var ret string
		return ret
	}
	return *o.LiableBalanceAccountId
}

// GetLiableBalanceAccountIdOk returns a tuple with the LiableBalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetLiableBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiableBalanceAccountId) {
		return nil, false
	}
	return o.LiableBalanceAccountId, true
}

// HasLiableBalanceAccountId returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasLiableBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.LiableBalanceAccountId) {
		return true
	}

	return false
}

// SetLiableBalanceAccountId gets a reference to the given string and assigns it to the LiableBalanceAccountId field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetLiableBalanceAccountId(v string) {
	o.LiableBalanceAccountId = &v
}

// GetNegativeBalanceSince returns the NegativeBalanceSince field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetNegativeBalanceSince() time.Time {
	if o == nil || common.IsNil(o.NegativeBalanceSince) {
		var ret time.Time
		return ret
	}
	return *o.NegativeBalanceSince
}

// GetNegativeBalanceSinceOk returns a tuple with the NegativeBalanceSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetNegativeBalanceSinceOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.NegativeBalanceSince) {
		return nil, false
	}
	return o.NegativeBalanceSince, true
}

// HasNegativeBalanceSince returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasNegativeBalanceSince() bool {
	if o != nil && !common.IsNil(o.NegativeBalanceSince) {
		return true
	}

	return false
}

// SetNegativeBalanceSince gets a reference to the given time.Time and assigns it to the NegativeBalanceSince field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetNegativeBalanceSince(v time.Time) {
	o.NegativeBalanceSince = &v
}

// GetScheduledCompensationAt returns the ScheduledCompensationAt field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationData) GetScheduledCompensationAt() time.Time {
	if o == nil || common.IsNil(o.ScheduledCompensationAt) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledCompensationAt
}

// GetScheduledCompensationAtOk returns a tuple with the ScheduledCompensationAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) GetScheduledCompensationAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ScheduledCompensationAt) {
		return nil, false
	}
	return o.ScheduledCompensationAt, true
}

// HasScheduledCompensationAt returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationData) HasScheduledCompensationAt() bool {
	if o != nil && !common.IsNil(o.ScheduledCompensationAt) {
		return true
	}

	return false
}

// SetScheduledCompensationAt gets a reference to the given time.Time and assigns it to the ScheduledCompensationAt field.
func (o *NegativeBalanceCompensationWarningNotificationData) SetScheduledCompensationAt(v time.Time) {
	o.ScheduledCompensationAt = &v
}

func (o NegativeBalanceCompensationWarningNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegativeBalanceCompensationWarningNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountHolder) {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.LiableBalanceAccountId) {
		toSerialize["liableBalanceAccountId"] = o.LiableBalanceAccountId
	}
	if !common.IsNil(o.NegativeBalanceSince) {
		toSerialize["negativeBalanceSince"] = o.NegativeBalanceSince
	}
	if !common.IsNil(o.ScheduledCompensationAt) {
		toSerialize["scheduledCompensationAt"] = o.ScheduledCompensationAt
	}
	return toSerialize, nil
}

type NullableNegativeBalanceCompensationWarningNotificationData struct {
	value *NegativeBalanceCompensationWarningNotificationData
	isSet bool
}

func (v NullableNegativeBalanceCompensationWarningNotificationData) Get() *NegativeBalanceCompensationWarningNotificationData {
	return v.value
}

func (v *NullableNegativeBalanceCompensationWarningNotificationData) Set(val *NegativeBalanceCompensationWarningNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableNegativeBalanceCompensationWarningNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableNegativeBalanceCompensationWarningNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegativeBalanceCompensationWarningNotificationData(val *NegativeBalanceCompensationWarningNotificationData) *NullableNegativeBalanceCompensationWarningNotificationData {
	return &NullableNegativeBalanceCompensationWarningNotificationData{value: val, isSet: true}
}

func (v NullableNegativeBalanceCompensationWarningNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegativeBalanceCompensationWarningNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
