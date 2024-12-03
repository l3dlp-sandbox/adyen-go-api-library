/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TimeOfDayRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TimeOfDayRestriction{}

// TimeOfDayRestriction struct for TimeOfDayRestriction
type TimeOfDayRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string     `json:"operation"`
	Value     *TimeOfDay `json:"value,omitempty"`
}

// NewTimeOfDayRestriction instantiates a new TimeOfDayRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeOfDayRestriction(operation string) *TimeOfDayRestriction {
	this := TimeOfDayRestriction{}
	this.Operation = operation
	return &this
}

// NewTimeOfDayRestrictionWithDefaults instantiates a new TimeOfDayRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeOfDayRestrictionWithDefaults() *TimeOfDayRestriction {
	this := TimeOfDayRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *TimeOfDayRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *TimeOfDayRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *TimeOfDayRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TimeOfDayRestriction) GetValue() TimeOfDay {
	if o == nil || common.IsNil(o.Value) {
		var ret TimeOfDay
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeOfDayRestriction) GetValueOk() (*TimeOfDay, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TimeOfDayRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given TimeOfDay and assigns it to the Value field.
func (o *TimeOfDayRestriction) SetValue(v TimeOfDay) {
	o.Value = &v
}

func (o TimeOfDayRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeOfDayRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTimeOfDayRestriction struct {
	value *TimeOfDayRestriction
	isSet bool
}

func (v NullableTimeOfDayRestriction) Get() *TimeOfDayRestriction {
	return v.value
}

func (v *NullableTimeOfDayRestriction) Set(val *TimeOfDayRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeOfDayRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeOfDayRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeOfDayRestriction(val *TimeOfDayRestriction) *NullableTimeOfDayRestriction {
	return &NullableTimeOfDayRestriction{value: val, isSet: true}
}

func (v NullableTimeOfDayRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeOfDayRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
