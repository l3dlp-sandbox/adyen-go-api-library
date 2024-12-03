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

// checks if the TimeOfDay type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TimeOfDay{}

// TimeOfDay struct for TimeOfDay
type TimeOfDay struct {
	// The end time in a time-only ISO-8601 extended offset format. For example: **08:00:00+02:00**, **22:30:00-03:00**.  
	EndTime *string `json:"endTime,omitempty"`
	// The start time in a time-only ISO-8601 extended offset format. For example: **08:00:00+02:00**, **22:30:00-03:00**.  
	StartTime *string `json:"startTime,omitempty"`
}

// NewTimeOfDay instantiates a new TimeOfDay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeOfDay() *TimeOfDay {
	this := TimeOfDay{}
	return &this
}

// NewTimeOfDayWithDefaults instantiates a new TimeOfDay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeOfDayWithDefaults() *TimeOfDay {
	this := TimeOfDay{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeOfDay) GetEndTime() string {
	if o == nil || common.IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeOfDay) GetEndTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeOfDay) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *TimeOfDay) SetEndTime(v string) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeOfDay) GetStartTime() string {
	if o == nil || common.IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeOfDay) GetStartTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeOfDay) HasStartTime() bool {
	if o != nil && !common.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *TimeOfDay) SetStartTime(v string) {
	o.StartTime = &v
}

func (o TimeOfDay) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeOfDay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !common.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableTimeOfDay struct {
	value *TimeOfDay
	isSet bool
}

func (v NullableTimeOfDay) Get() *TimeOfDay {
	return v.value
}

func (v *NullableTimeOfDay) Set(val *TimeOfDay) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeOfDay) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeOfDay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeOfDay(val *TimeOfDay) *NullableTimeOfDay {
	return &NullableTimeOfDay{value: val, isSet: true}
}

func (v NullableTimeOfDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeOfDay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



