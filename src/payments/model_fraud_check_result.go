/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the FraudCheckResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FraudCheckResult{}

// FraudCheckResult struct for FraudCheckResult
type FraudCheckResult struct {
	// The fraud score generated by the risk check.
	AccountScore int32 `json:"accountScore"`
	// The ID of the risk check.
	CheckId int32 `json:"checkId"`
	// The name of the risk check.
	Name string `json:"name"`
}

// NewFraudCheckResult instantiates a new FraudCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFraudCheckResult(accountScore int32, checkId int32, name string) *FraudCheckResult {
	this := FraudCheckResult{}
	this.AccountScore = accountScore
	this.CheckId = checkId
	this.Name = name
	return &this
}

// NewFraudCheckResultWithDefaults instantiates a new FraudCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFraudCheckResultWithDefaults() *FraudCheckResult {
	this := FraudCheckResult{}
	return &this
}

// GetAccountScore returns the AccountScore field value
func (o *FraudCheckResult) GetAccountScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountScore
}

// GetAccountScoreOk returns a tuple with the AccountScore field value
// and a boolean to check if the value has been set.
func (o *FraudCheckResult) GetAccountScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountScore, true
}

// SetAccountScore sets field value
func (o *FraudCheckResult) SetAccountScore(v int32) {
	o.AccountScore = v
}

// GetCheckId returns the CheckId field value
func (o *FraudCheckResult) GetCheckId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value
// and a boolean to check if the value has been set.
func (o *FraudCheckResult) GetCheckIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckId, true
}

// SetCheckId sets field value
func (o *FraudCheckResult) SetCheckId(v int32) {
	o.CheckId = v
}

// GetName returns the Name field value
func (o *FraudCheckResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FraudCheckResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FraudCheckResult) SetName(v string) {
	o.Name = v
}

func (o FraudCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FraudCheckResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountScore"] = o.AccountScore
	toSerialize["checkId"] = o.CheckId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableFraudCheckResult struct {
	value *FraudCheckResult
	isSet bool
}

func (v NullableFraudCheckResult) Get() *FraudCheckResult {
	return v.value
}

func (v *NullableFraudCheckResult) Set(val *FraudCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFraudCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFraudCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFraudCheckResult(val *FraudCheckResult) *NullableFraudCheckResult {
	return &NullableFraudCheckResult{value: val, isSet: true}
}

func (v NullableFraudCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFraudCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
