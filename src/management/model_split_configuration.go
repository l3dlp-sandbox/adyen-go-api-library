/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the SplitConfiguration type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SplitConfiguration{}

// SplitConfiguration struct for SplitConfiguration
type SplitConfiguration struct {
	// Your description for the split configuration.
	Description string `json:"description"`
	// Array of rules that define the split configuration behavior.
	Rules []SplitConfigurationRule `json:"rules"`
	// Unique identifier of the split configuration.
	SplitConfigurationId *string `json:"splitConfigurationId,omitempty"`
	// List of stores to which the split configuration applies.
	Stores []string `json:"stores,omitempty"`
}

// NewSplitConfiguration instantiates a new SplitConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitConfiguration(description string, rules []SplitConfigurationRule) *SplitConfiguration {
	this := SplitConfiguration{}
	this.Description = description
	this.Rules = rules
	return &this
}

// NewSplitConfigurationWithDefaults instantiates a new SplitConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitConfigurationWithDefaults() *SplitConfiguration {
	this := SplitConfiguration{}
	return &this
}

// GetDescription returns the Description field value
func (o *SplitConfiguration) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SplitConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SplitConfiguration) SetDescription(v string) {
	o.Description = v
}

// GetRules returns the Rules field value
func (o *SplitConfiguration) GetRules() []SplitConfigurationRule {
	if o == nil {
		var ret []SplitConfigurationRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *SplitConfiguration) GetRulesOk() ([]SplitConfigurationRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *SplitConfiguration) SetRules(v []SplitConfigurationRule) {
	o.Rules = v
}

// GetSplitConfigurationId returns the SplitConfigurationId field value if set, zero value otherwise.
func (o *SplitConfiguration) GetSplitConfigurationId() string {
	if o == nil || common.IsNil(o.SplitConfigurationId) {
		var ret string
		return ret
	}
	return *o.SplitConfigurationId
}

// GetSplitConfigurationIdOk returns a tuple with the SplitConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfiguration) GetSplitConfigurationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SplitConfigurationId) {
		return nil, false
	}
	return o.SplitConfigurationId, true
}

// HasSplitConfigurationId returns a boolean if a field has been set.
func (o *SplitConfiguration) HasSplitConfigurationId() bool {
	if o != nil && !common.IsNil(o.SplitConfigurationId) {
		return true
	}

	return false
}

// SetSplitConfigurationId gets a reference to the given string and assigns it to the SplitConfigurationId field.
func (o *SplitConfiguration) SetSplitConfigurationId(v string) {
	o.SplitConfigurationId = &v
}

// GetStores returns the Stores field value if set, zero value otherwise.
func (o *SplitConfiguration) GetStores() []string {
	if o == nil || common.IsNil(o.Stores) {
		var ret []string
		return ret
	}
	return o.Stores
}

// GetStoresOk returns a tuple with the Stores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfiguration) GetStoresOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Stores) {
		return nil, false
	}
	return o.Stores, true
}

// HasStores returns a boolean if a field has been set.
func (o *SplitConfiguration) HasStores() bool {
	if o != nil && !common.IsNil(o.Stores) {
		return true
	}

	return false
}

// SetStores gets a reference to the given []string and assigns it to the Stores field.
func (o *SplitConfiguration) SetStores(v []string) {
	o.Stores = v
}

func (o SplitConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["rules"] = o.Rules
	if !common.IsNil(o.SplitConfigurationId) {
		toSerialize["splitConfigurationId"] = o.SplitConfigurationId
	}
	if !common.IsNil(o.Stores) {
		toSerialize["stores"] = o.Stores
	}
	return toSerialize, nil
}

type NullableSplitConfiguration struct {
	value *SplitConfiguration
	isSet bool
}

func (v NullableSplitConfiguration) Get() *SplitConfiguration {
	return v.value
}

func (v *NullableSplitConfiguration) Set(val *SplitConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitConfiguration(val *SplitConfiguration) *NullableSplitConfiguration {
	return &NullableSplitConfiguration{value: val, isSet: true}
}

func (v NullableSplitConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
