/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PlatformChargebackLogic type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlatformChargebackLogic{}

// PlatformChargebackLogic struct for PlatformChargebackLogic
type PlatformChargebackLogic struct {
	// The method of handling the chargeback.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**, **deductAccordingToSplitRatio**.
	Behavior *string `json:"behavior,omitempty"`
	// The unique identifier of the balance account to which the chargeback fees are booked. By default, the chargeback fees are booked to your liable balance account.
	CostAllocationAccount *string `json:"costAllocationAccount,omitempty"`
	// The unique identifier of the balance account against which the disputed amount is booked.  Required if `behavior` is **deductFromOneBalanceAccount**.
	TargetAccount *string `json:"targetAccount,omitempty"`
}

// NewPlatformChargebackLogic instantiates a new PlatformChargebackLogic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformChargebackLogic() *PlatformChargebackLogic {
	this := PlatformChargebackLogic{}
	return &this
}

// NewPlatformChargebackLogicWithDefaults instantiates a new PlatformChargebackLogic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformChargebackLogicWithDefaults() *PlatformChargebackLogic {
	this := PlatformChargebackLogic{}
	return &this
}

// GetBehavior returns the Behavior field value if set, zero value otherwise.
func (o *PlatformChargebackLogic) GetBehavior() string {
	if o == nil || common.IsNil(o.Behavior) {
		var ret string
		return ret
	}
	return *o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformChargebackLogic) GetBehaviorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Behavior) {
		return nil, false
	}
	return o.Behavior, true
}

// HasBehavior returns a boolean if a field has been set.
func (o *PlatformChargebackLogic) HasBehavior() bool {
	if o != nil && !common.IsNil(o.Behavior) {
		return true
	}

	return false
}

// SetBehavior gets a reference to the given string and assigns it to the Behavior field.
func (o *PlatformChargebackLogic) SetBehavior(v string) {
	o.Behavior = &v
}

// GetCostAllocationAccount returns the CostAllocationAccount field value if set, zero value otherwise.
func (o *PlatformChargebackLogic) GetCostAllocationAccount() string {
	if o == nil || common.IsNil(o.CostAllocationAccount) {
		var ret string
		return ret
	}
	return *o.CostAllocationAccount
}

// GetCostAllocationAccountOk returns a tuple with the CostAllocationAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformChargebackLogic) GetCostAllocationAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.CostAllocationAccount) {
		return nil, false
	}
	return o.CostAllocationAccount, true
}

// HasCostAllocationAccount returns a boolean if a field has been set.
func (o *PlatformChargebackLogic) HasCostAllocationAccount() bool {
	if o != nil && !common.IsNil(o.CostAllocationAccount) {
		return true
	}

	return false
}

// SetCostAllocationAccount gets a reference to the given string and assigns it to the CostAllocationAccount field.
func (o *PlatformChargebackLogic) SetCostAllocationAccount(v string) {
	o.CostAllocationAccount = &v
}

// GetTargetAccount returns the TargetAccount field value if set, zero value otherwise.
func (o *PlatformChargebackLogic) GetTargetAccount() string {
	if o == nil || common.IsNil(o.TargetAccount) {
		var ret string
		return ret
	}
	return *o.TargetAccount
}

// GetTargetAccountOk returns a tuple with the TargetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformChargebackLogic) GetTargetAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TargetAccount) {
		return nil, false
	}
	return o.TargetAccount, true
}

// HasTargetAccount returns a boolean if a field has been set.
func (o *PlatformChargebackLogic) HasTargetAccount() bool {
	if o != nil && !common.IsNil(o.TargetAccount) {
		return true
	}

	return false
}

// SetTargetAccount gets a reference to the given string and assigns it to the TargetAccount field.
func (o *PlatformChargebackLogic) SetTargetAccount(v string) {
	o.TargetAccount = &v
}

func (o PlatformChargebackLogic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformChargebackLogic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Behavior) {
		toSerialize["behavior"] = o.Behavior
	}
	if !common.IsNil(o.CostAllocationAccount) {
		toSerialize["costAllocationAccount"] = o.CostAllocationAccount
	}
	if !common.IsNil(o.TargetAccount) {
		toSerialize["targetAccount"] = o.TargetAccount
	}
	return toSerialize, nil
}

type NullablePlatformChargebackLogic struct {
	value *PlatformChargebackLogic
	isSet bool
}

func (v NullablePlatformChargebackLogic) Get() *PlatformChargebackLogic {
	return v.value
}

func (v *NullablePlatformChargebackLogic) Set(val *PlatformChargebackLogic) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformChargebackLogic) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformChargebackLogic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformChargebackLogic(val *PlatformChargebackLogic) *NullablePlatformChargebackLogic {
	return &NullablePlatformChargebackLogic{value: val, isSet: true}
}

func (v NullablePlatformChargebackLogic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformChargebackLogic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PlatformChargebackLogic) isValidBehavior() bool {
	var allowedEnumValues = []string{"deductAccordingToSplitRatio", "deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetBehavior() == allowed {
			return true
		}
	}
	return false
}
