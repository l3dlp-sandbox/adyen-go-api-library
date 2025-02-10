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

// checks if the NegativeBalanceCompensationWarningNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NegativeBalanceCompensationWarningNotificationRequest{}

// NegativeBalanceCompensationWarningNotificationRequest struct for NegativeBalanceCompensationWarningNotificationRequest
type NegativeBalanceCompensationWarningNotificationRequest struct {
	Data NegativeBalanceCompensationWarningNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// When the event was queued.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of webhook.
	Type string `json:"type"`
}

// NewNegativeBalanceCompensationWarningNotificationRequest instantiates a new NegativeBalanceCompensationWarningNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegativeBalanceCompensationWarningNotificationRequest(data NegativeBalanceCompensationWarningNotificationData, environment string, type_ string) *NegativeBalanceCompensationWarningNotificationRequest {
	this := NegativeBalanceCompensationWarningNotificationRequest{}
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewNegativeBalanceCompensationWarningNotificationRequestWithDefaults instantiates a new NegativeBalanceCompensationWarningNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegativeBalanceCompensationWarningNotificationRequestWithDefaults() *NegativeBalanceCompensationWarningNotificationRequest {
	this := NegativeBalanceCompensationWarningNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetData() NegativeBalanceCompensationWarningNotificationData {
	if o == nil {
		var ret NegativeBalanceCompensationWarningNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetDataOk() (*NegativeBalanceCompensationWarningNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NegativeBalanceCompensationWarningNotificationRequest) SetData(v NegativeBalanceCompensationWarningNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *NegativeBalanceCompensationWarningNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetTimestamp() time.Time {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *NegativeBalanceCompensationWarningNotificationRequest) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *NegativeBalanceCompensationWarningNotificationRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NegativeBalanceCompensationWarningNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NegativeBalanceCompensationWarningNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o NegativeBalanceCompensationWarningNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegativeBalanceCompensationWarningNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNegativeBalanceCompensationWarningNotificationRequest struct {
	value *NegativeBalanceCompensationWarningNotificationRequest
	isSet bool
}

func (v NullableNegativeBalanceCompensationWarningNotificationRequest) Get() *NegativeBalanceCompensationWarningNotificationRequest {
	return v.value
}

func (v *NullableNegativeBalanceCompensationWarningNotificationRequest) Set(val *NegativeBalanceCompensationWarningNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNegativeBalanceCompensationWarningNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNegativeBalanceCompensationWarningNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegativeBalanceCompensationWarningNotificationRequest(val *NegativeBalanceCompensationWarningNotificationRequest) *NullableNegativeBalanceCompensationWarningNotificationRequest {
	return &NullableNegativeBalanceCompensationWarningNotificationRequest{value: val, isSet: true}
}

func (v NullableNegativeBalanceCompensationWarningNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegativeBalanceCompensationWarningNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *NegativeBalanceCompensationWarningNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.negativeBalanceCompensationWarning.scheduled"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
