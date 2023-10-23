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

// checks if the OfflineProcessing type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OfflineProcessing{}

// OfflineProcessing struct for OfflineProcessing
type OfflineProcessing struct {
	// The maximum offline transaction amount for chip cards, in the processing currency and specified in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	ChipFloorLimit *int32 `json:"chipFloorLimit,omitempty"`
	// The maximum offline transaction amount for swiped cards, in the specified currency.
	OfflineSwipeLimits []MinorUnitsMonetaryValue `json:"offlineSwipeLimits,omitempty"`
}

// NewOfflineProcessing instantiates a new OfflineProcessing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfflineProcessing() *OfflineProcessing {
	this := OfflineProcessing{}
	return &this
}

// NewOfflineProcessingWithDefaults instantiates a new OfflineProcessing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfflineProcessingWithDefaults() *OfflineProcessing {
	this := OfflineProcessing{}
	return &this
}

// GetChipFloorLimit returns the ChipFloorLimit field value if set, zero value otherwise.
func (o *OfflineProcessing) GetChipFloorLimit() int32 {
	if o == nil || common.IsNil(o.ChipFloorLimit) {
		var ret int32
		return ret
	}
	return *o.ChipFloorLimit
}

// GetChipFloorLimitOk returns a tuple with the ChipFloorLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineProcessing) GetChipFloorLimitOk() (*int32, bool) {
	if o == nil || common.IsNil(o.ChipFloorLimit) {
		return nil, false
	}
	return o.ChipFloorLimit, true
}

// HasChipFloorLimit returns a boolean if a field has been set.
func (o *OfflineProcessing) HasChipFloorLimit() bool {
	if o != nil && !common.IsNil(o.ChipFloorLimit) {
		return true
	}

	return false
}

// SetChipFloorLimit gets a reference to the given int32 and assigns it to the ChipFloorLimit field.
func (o *OfflineProcessing) SetChipFloorLimit(v int32) {
	o.ChipFloorLimit = &v
}

// GetOfflineSwipeLimits returns the OfflineSwipeLimits field value if set, zero value otherwise.
func (o *OfflineProcessing) GetOfflineSwipeLimits() []MinorUnitsMonetaryValue {
	if o == nil || common.IsNil(o.OfflineSwipeLimits) {
		var ret []MinorUnitsMonetaryValue
		return ret
	}
	return o.OfflineSwipeLimits
}

// GetOfflineSwipeLimitsOk returns a tuple with the OfflineSwipeLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineProcessing) GetOfflineSwipeLimitsOk() ([]MinorUnitsMonetaryValue, bool) {
	if o == nil || common.IsNil(o.OfflineSwipeLimits) {
		return nil, false
	}
	return o.OfflineSwipeLimits, true
}

// HasOfflineSwipeLimits returns a boolean if a field has been set.
func (o *OfflineProcessing) HasOfflineSwipeLimits() bool {
	if o != nil && !common.IsNil(o.OfflineSwipeLimits) {
		return true
	}

	return false
}

// SetOfflineSwipeLimits gets a reference to the given []MinorUnitsMonetaryValue and assigns it to the OfflineSwipeLimits field.
func (o *OfflineProcessing) SetOfflineSwipeLimits(v []MinorUnitsMonetaryValue) {
	o.OfflineSwipeLimits = v
}

func (o OfflineProcessing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfflineProcessing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ChipFloorLimit) {
		toSerialize["chipFloorLimit"] = o.ChipFloorLimit
	}
	if !common.IsNil(o.OfflineSwipeLimits) {
		toSerialize["offlineSwipeLimits"] = o.OfflineSwipeLimits
	}
	return toSerialize, nil
}

type NullableOfflineProcessing struct {
	value *OfflineProcessing
	isSet bool
}

func (v NullableOfflineProcessing) Get() *OfflineProcessing {
	return v.value
}

func (v *NullableOfflineProcessing) Set(val *OfflineProcessing) {
	v.value = val
	v.isSet = true
}

func (v NullableOfflineProcessing) IsSet() bool {
	return v.isSet
}

func (v *NullableOfflineProcessing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfflineProcessing(val *OfflineProcessing) *NullableOfflineProcessing {
	return &NullableOfflineProcessing{value: val, isSet: true}
}

func (v NullableOfflineProcessing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfflineProcessing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
