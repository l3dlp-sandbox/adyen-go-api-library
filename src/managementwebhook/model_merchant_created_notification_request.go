/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the MerchantCreatedNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MerchantCreatedNotificationRequest{}

// MerchantCreatedNotificationRequest struct for MerchantCreatedNotificationRequest
type MerchantCreatedNotificationRequest struct {
	// Timestamp for when the webhook was created.
	CreatedAt time.Time                     `json:"createdAt"`
	Data      AccountCreateNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of notification.
	Type string `json:"type"`
}

// NewMerchantCreatedNotificationRequest instantiates a new MerchantCreatedNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreatedNotificationRequest(createdAt time.Time, data AccountCreateNotificationData, environment string, type_ string) *MerchantCreatedNotificationRequest {
	this := MerchantCreatedNotificationRequest{}
	this.CreatedAt = createdAt
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewMerchantCreatedNotificationRequestWithDefaults instantiates a new MerchantCreatedNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreatedNotificationRequestWithDefaults() *MerchantCreatedNotificationRequest {
	this := MerchantCreatedNotificationRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *MerchantCreatedNotificationRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MerchantCreatedNotificationRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MerchantCreatedNotificationRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value
func (o *MerchantCreatedNotificationRequest) GetData() AccountCreateNotificationData {
	if o == nil {
		var ret AccountCreateNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MerchantCreatedNotificationRequest) GetDataOk() (*AccountCreateNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MerchantCreatedNotificationRequest) SetData(v AccountCreateNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *MerchantCreatedNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *MerchantCreatedNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *MerchantCreatedNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *MerchantCreatedNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MerchantCreatedNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MerchantCreatedNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o MerchantCreatedNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCreatedNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMerchantCreatedNotificationRequest struct {
	value *MerchantCreatedNotificationRequest
	isSet bool
}

func (v NullableMerchantCreatedNotificationRequest) Get() *MerchantCreatedNotificationRequest {
	return v.value
}

func (v *NullableMerchantCreatedNotificationRequest) Set(val *MerchantCreatedNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreatedNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreatedNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreatedNotificationRequest(val *MerchantCreatedNotificationRequest) *NullableMerchantCreatedNotificationRequest {
	return &NullableMerchantCreatedNotificationRequest{value: val, isSet: true}
}

func (v NullableMerchantCreatedNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreatedNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *MerchantCreatedNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"merchant.created"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
