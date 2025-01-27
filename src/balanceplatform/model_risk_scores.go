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

// checks if the RiskScores type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RiskScores{}

// RiskScores struct for RiskScores
type RiskScores struct {
	// Transaction risk score provided by Mastercard. Values provided by Mastercard range between 0 (lowest risk) to 998 (highest risk).
	Mastercard *int32 `json:"mastercard,omitempty"`
	// Transaction risk score provided by Visa. Values provided by Visa range between 01 (lowest risk) to 99 (highest risk).
	Visa *int32 `json:"visa,omitempty"`
}

// NewRiskScores instantiates a new RiskScores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskScores() *RiskScores {
	this := RiskScores{}
	return &this
}

// NewRiskScoresWithDefaults instantiates a new RiskScores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskScoresWithDefaults() *RiskScores {
	this := RiskScores{}
	return &this
}

// GetMastercard returns the Mastercard field value if set, zero value otherwise.
func (o *RiskScores) GetMastercard() int32 {
	if o == nil || common.IsNil(o.Mastercard) {
		var ret int32
		return ret
	}
	return *o.Mastercard
}

// GetMastercardOk returns a tuple with the Mastercard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskScores) GetMastercardOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Mastercard) {
		return nil, false
	}
	return o.Mastercard, true
}

// HasMastercard returns a boolean if a field has been set.
func (o *RiskScores) HasMastercard() bool {
	if o != nil && !common.IsNil(o.Mastercard) {
		return true
	}

	return false
}

// SetMastercard gets a reference to the given int32 and assigns it to the Mastercard field.
func (o *RiskScores) SetMastercard(v int32) {
	o.Mastercard = &v
}

// GetVisa returns the Visa field value if set, zero value otherwise.
func (o *RiskScores) GetVisa() int32 {
	if o == nil || common.IsNil(o.Visa) {
		var ret int32
		return ret
	}
	return *o.Visa
}

// GetVisaOk returns a tuple with the Visa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskScores) GetVisaOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Visa) {
		return nil, false
	}
	return o.Visa, true
}

// HasVisa returns a boolean if a field has been set.
func (o *RiskScores) HasVisa() bool {
	if o != nil && !common.IsNil(o.Visa) {
		return true
	}

	return false
}

// SetVisa gets a reference to the given int32 and assigns it to the Visa field.
func (o *RiskScores) SetVisa(v int32) {
	o.Visa = &v
}

func (o RiskScores) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskScores) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Mastercard) {
		toSerialize["mastercard"] = o.Mastercard
	}
	if !common.IsNil(o.Visa) {
		toSerialize["visa"] = o.Visa
	}
	return toSerialize, nil
}

type NullableRiskScores struct {
	value *RiskScores
	isSet bool
}

func (v NullableRiskScores) Get() *RiskScores {
	return v.value
}

func (v *NullableRiskScores) Set(val *RiskScores) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskScores) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskScores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskScores(val *RiskScores) *NullableRiskScores {
	return &NullableRiskScores{value: val, isSet: true}
}

func (v NullableRiskScores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskScores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
