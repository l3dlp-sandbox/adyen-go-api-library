/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the BalanceAccountNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceAccountNotificationRequest{}

// BalanceAccountNotificationRequest struct for BalanceAccountNotificationRequest
type BalanceAccountNotificationRequest struct {
	Data BalanceAccountNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// When the event was queued.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of webhook.
	Type string `json:"type"`
}

// NewBalanceAccountNotificationRequest instantiates a new BalanceAccountNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceAccountNotificationRequest(data BalanceAccountNotificationData, environment string, type_ string) *BalanceAccountNotificationRequest {
	this := BalanceAccountNotificationRequest{}
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewBalanceAccountNotificationRequestWithDefaults instantiates a new BalanceAccountNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceAccountNotificationRequestWithDefaults() *BalanceAccountNotificationRequest {
	this := BalanceAccountNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BalanceAccountNotificationRequest) GetData() BalanceAccountNotificationData {
	if o == nil {
		var ret BalanceAccountNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BalanceAccountNotificationRequest) GetDataOk() (*BalanceAccountNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BalanceAccountNotificationRequest) SetData(v BalanceAccountNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *BalanceAccountNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *BalanceAccountNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *BalanceAccountNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BalanceAccountNotificationRequest) GetTimestamp() time.Time {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountNotificationRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BalanceAccountNotificationRequest) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *BalanceAccountNotificationRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value
func (o *BalanceAccountNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BalanceAccountNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BalanceAccountNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o BalanceAccountNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceAccountNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBalanceAccountNotificationRequest struct {
	value *BalanceAccountNotificationRequest
	isSet bool
}

func (v NullableBalanceAccountNotificationRequest) Get() *BalanceAccountNotificationRequest {
	return v.value
}

func (v *NullableBalanceAccountNotificationRequest) Set(val *BalanceAccountNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceAccountNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceAccountNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceAccountNotificationRequest(val *BalanceAccountNotificationRequest) *NullableBalanceAccountNotificationRequest {
	return &NullableBalanceAccountNotificationRequest{value: val, isSet: true}
}

func (v NullableBalanceAccountNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceAccountNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BalanceAccountNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.balanceAccount.updated", "balancePlatform.balanceAccount.created"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
