/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the StoreSplitConfiguration type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreSplitConfiguration{}

// StoreSplitConfiguration struct for StoreSplitConfiguration
type StoreSplitConfiguration struct {
	// The [unique identifier of the balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balanceAccounts/{id}__queryParam_id) to which the split amount must be booked, depending on the defined [split logic](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/_merchantId_/splitConfigurations#request-rules-splitLogic).
	BalanceAccountId *string `json:"balanceAccountId,omitempty"`
	// The unique identifier of the [split configuration profile](https://docs.adyen.com/platforms/automatic-split-configuration/create-split-configuration/).
	SplitConfigurationId *string `json:"splitConfigurationId,omitempty"`
}

// NewStoreSplitConfiguration instantiates a new StoreSplitConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreSplitConfiguration() *StoreSplitConfiguration {
	this := StoreSplitConfiguration{}
	return &this
}

// NewStoreSplitConfigurationWithDefaults instantiates a new StoreSplitConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreSplitConfigurationWithDefaults() *StoreSplitConfiguration {
	this := StoreSplitConfiguration{}
	return &this
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *StoreSplitConfiguration) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreSplitConfiguration) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *StoreSplitConfiguration) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *StoreSplitConfiguration) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetSplitConfigurationId returns the SplitConfigurationId field value if set, zero value otherwise.
func (o *StoreSplitConfiguration) GetSplitConfigurationId() string {
	if o == nil || common.IsNil(o.SplitConfigurationId) {
		var ret string
		return ret
	}
	return *o.SplitConfigurationId
}

// GetSplitConfigurationIdOk returns a tuple with the SplitConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreSplitConfiguration) GetSplitConfigurationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SplitConfigurationId) {
		return nil, false
	}
	return o.SplitConfigurationId, true
}

// HasSplitConfigurationId returns a boolean if a field has been set.
func (o *StoreSplitConfiguration) HasSplitConfigurationId() bool {
	if o != nil && !common.IsNil(o.SplitConfigurationId) {
		return true
	}

	return false
}

// SetSplitConfigurationId gets a reference to the given string and assigns it to the SplitConfigurationId field.
func (o *StoreSplitConfiguration) SetSplitConfigurationId(v string) {
	o.SplitConfigurationId = &v
}

func (o StoreSplitConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreSplitConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	if !common.IsNil(o.SplitConfigurationId) {
		toSerialize["splitConfigurationId"] = o.SplitConfigurationId
	}
	return toSerialize, nil
}

type NullableStoreSplitConfiguration struct {
	value *StoreSplitConfiguration
	isSet bool
}

func (v NullableStoreSplitConfiguration) Get() *StoreSplitConfiguration {
	return v.value
}

func (v *NullableStoreSplitConfiguration) Set(val *StoreSplitConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreSplitConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreSplitConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreSplitConfiguration(val *StoreSplitConfiguration) *NullableStoreSplitConfiguration {
	return &NullableStoreSplitConfiguration{value: val, isSet: true}
}

func (v NullableStoreSplitConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreSplitConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
