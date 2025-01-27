/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the EstimationTrackingData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EstimationTrackingData{}

// EstimationTrackingData struct for EstimationTrackingData
type EstimationTrackingData struct {
	// The estimated time the beneficiary should have access to the funds.
	EstimatedArrivalTime time.Time `json:"estimatedArrivalTime"`
	// The type of tracking event.   Possible values:   - **estimation**: the estimated date and time of when the funds will be credited has been determined.
	Type string `json:"type"`
}

// NewEstimationTrackingData instantiates a new EstimationTrackingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimationTrackingData(estimatedArrivalTime time.Time, type_ string) *EstimationTrackingData {
	this := EstimationTrackingData{}
	this.EstimatedArrivalTime = estimatedArrivalTime
	this.Type = type_
	return &this
}

// NewEstimationTrackingDataWithDefaults instantiates a new EstimationTrackingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimationTrackingDataWithDefaults() *EstimationTrackingData {
	this := EstimationTrackingData{}
	var type_ string = "estimation"
	this.Type = type_
	return &this
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value
func (o *EstimationTrackingData) GetEstimatedArrivalTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value
// and a boolean to check if the value has been set.
func (o *EstimationTrackingData) GetEstimatedArrivalTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedArrivalTime, true
}

// SetEstimatedArrivalTime sets field value
func (o *EstimationTrackingData) SetEstimatedArrivalTime(v time.Time) {
	o.EstimatedArrivalTime = v
}

// GetType returns the Type field value
func (o *EstimationTrackingData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EstimationTrackingData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EstimationTrackingData) SetType(v string) {
	o.Type = v
}

func (o EstimationTrackingData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimationTrackingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["estimatedArrivalTime"] = o.EstimatedArrivalTime
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableEstimationTrackingData struct {
	value *EstimationTrackingData
	isSet bool
}

func (v NullableEstimationTrackingData) Get() *EstimationTrackingData {
	return v.value
}

func (v *NullableEstimationTrackingData) Set(val *EstimationTrackingData) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimationTrackingData) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimationTrackingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimationTrackingData(val *EstimationTrackingData) *NullableEstimationTrackingData {
	return &NullableEstimationTrackingData{value: val, isSet: true}
}

func (v NullableEstimationTrackingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimationTrackingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *EstimationTrackingData) isValidType() bool {
	var allowedEnumValues = []string{"estimation"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
