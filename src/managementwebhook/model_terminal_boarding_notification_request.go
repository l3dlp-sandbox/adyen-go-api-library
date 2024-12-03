/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
	"time"
)

// checks if the TerminalBoardingNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalBoardingNotificationRequest{}

// TerminalBoardingNotificationRequest struct for TerminalBoardingNotificationRequest
type TerminalBoardingNotificationRequest struct {
	// Timestamp for when the webhook was created.
	CreatedAt time.Time `json:"createdAt"`
	Data TerminalBoardingData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of notification.
	Type string `json:"type"`
}

// NewTerminalBoardingNotificationRequest instantiates a new TerminalBoardingNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalBoardingNotificationRequest(createdAt time.Time, data TerminalBoardingData, environment string, type_ string) *TerminalBoardingNotificationRequest {
	this := TerminalBoardingNotificationRequest{}
	this.CreatedAt = createdAt
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewTerminalBoardingNotificationRequestWithDefaults instantiates a new TerminalBoardingNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalBoardingNotificationRequestWithDefaults() *TerminalBoardingNotificationRequest {
	this := TerminalBoardingNotificationRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TerminalBoardingNotificationRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TerminalBoardingNotificationRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TerminalBoardingNotificationRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value
func (o *TerminalBoardingNotificationRequest) GetData() TerminalBoardingData {
	if o == nil {
		var ret TerminalBoardingData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TerminalBoardingNotificationRequest) GetDataOk() (*TerminalBoardingData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TerminalBoardingNotificationRequest) SetData(v TerminalBoardingData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *TerminalBoardingNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *TerminalBoardingNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *TerminalBoardingNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *TerminalBoardingNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerminalBoardingNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TerminalBoardingNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o TerminalBoardingNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalBoardingNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTerminalBoardingNotificationRequest struct {
	value *TerminalBoardingNotificationRequest
	isSet bool
}

func (v NullableTerminalBoardingNotificationRequest) Get() *TerminalBoardingNotificationRequest {
	return v.value
}

func (v *NullableTerminalBoardingNotificationRequest) Set(val *TerminalBoardingNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalBoardingNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalBoardingNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalBoardingNotificationRequest(val *TerminalBoardingNotificationRequest) *NullableTerminalBoardingNotificationRequest {
	return &NullableTerminalBoardingNotificationRequest{value: val, isSet: true}
}

func (v NullableTerminalBoardingNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalBoardingNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *TerminalBoardingNotificationRequest) isValidType() bool {
    var allowedEnumValues = []string{ "terminalBoarding.triggered" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

