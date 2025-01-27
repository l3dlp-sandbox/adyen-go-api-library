/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the PaymentMethodCreatedNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodCreatedNotificationRequest{}

// PaymentMethodCreatedNotificationRequest struct for PaymentMethodCreatedNotificationRequest
type PaymentMethodCreatedNotificationRequest struct {
	// Timestamp for when the webhook was created.
	CreatedAt time.Time                  `json:"createdAt"`
	Data      MidServiceNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of notification.
	Type string `json:"type"`
}

// NewPaymentMethodCreatedNotificationRequest instantiates a new PaymentMethodCreatedNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCreatedNotificationRequest(createdAt time.Time, data MidServiceNotificationData, environment string, type_ string) *PaymentMethodCreatedNotificationRequest {
	this := PaymentMethodCreatedNotificationRequest{}
	this.CreatedAt = createdAt
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewPaymentMethodCreatedNotificationRequestWithDefaults instantiates a new PaymentMethodCreatedNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCreatedNotificationRequestWithDefaults() *PaymentMethodCreatedNotificationRequest {
	this := PaymentMethodCreatedNotificationRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentMethodCreatedNotificationRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreatedNotificationRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentMethodCreatedNotificationRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value
func (o *PaymentMethodCreatedNotificationRequest) GetData() MidServiceNotificationData {
	if o == nil {
		var ret MidServiceNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreatedNotificationRequest) GetDataOk() (*MidServiceNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaymentMethodCreatedNotificationRequest) SetData(v MidServiceNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *PaymentMethodCreatedNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreatedNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *PaymentMethodCreatedNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *PaymentMethodCreatedNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreatedNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodCreatedNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o PaymentMethodCreatedNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCreatedNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePaymentMethodCreatedNotificationRequest struct {
	value *PaymentMethodCreatedNotificationRequest
	isSet bool
}

func (v NullablePaymentMethodCreatedNotificationRequest) Get() *PaymentMethodCreatedNotificationRequest {
	return v.value
}

func (v *NullablePaymentMethodCreatedNotificationRequest) Set(val *PaymentMethodCreatedNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCreatedNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCreatedNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCreatedNotificationRequest(val *PaymentMethodCreatedNotificationRequest) *NullablePaymentMethodCreatedNotificationRequest {
	return &NullablePaymentMethodCreatedNotificationRequest{value: val, isSet: true}
}

func (v NullablePaymentMethodCreatedNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCreatedNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentMethodCreatedNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"paymentMethod.created"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
