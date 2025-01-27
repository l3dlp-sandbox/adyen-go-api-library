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

// checks if the AccountHolderNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountHolderNotificationRequest{}

// AccountHolderNotificationRequest struct for AccountHolderNotificationRequest
type AccountHolderNotificationRequest struct {
	Data AccountHolderNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// When the event was queued.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of webhook.
	Type string `json:"type"`
}

// NewAccountHolderNotificationRequest instantiates a new AccountHolderNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderNotificationRequest(data AccountHolderNotificationData, environment string, type_ string) *AccountHolderNotificationRequest {
	this := AccountHolderNotificationRequest{}
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewAccountHolderNotificationRequestWithDefaults instantiates a new AccountHolderNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderNotificationRequestWithDefaults() *AccountHolderNotificationRequest {
	this := AccountHolderNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AccountHolderNotificationRequest) GetData() AccountHolderNotificationData {
	if o == nil {
		var ret AccountHolderNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountHolderNotificationRequest) GetDataOk() (*AccountHolderNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AccountHolderNotificationRequest) SetData(v AccountHolderNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *AccountHolderNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *AccountHolderNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *AccountHolderNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AccountHolderNotificationRequest) GetTimestamp() time.Time {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderNotificationRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AccountHolderNotificationRequest) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *AccountHolderNotificationRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value
func (o *AccountHolderNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountHolderNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountHolderNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o AccountHolderNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountHolderNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAccountHolderNotificationRequest struct {
	value *AccountHolderNotificationRequest
	isSet bool
}

func (v NullableAccountHolderNotificationRequest) Get() *AccountHolderNotificationRequest {
	return v.value
}

func (v *NullableAccountHolderNotificationRequest) Set(val *AccountHolderNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderNotificationRequest(val *AccountHolderNotificationRequest) *NullableAccountHolderNotificationRequest {
	return &NullableAccountHolderNotificationRequest{value: val, isSet: true}
}

func (v NullableAccountHolderNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AccountHolderNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.accountHolder.updated", "balancePlatform.accountHolder.created"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
