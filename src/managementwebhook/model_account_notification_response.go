/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the AccountNotificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountNotificationResponse{}

// AccountNotificationResponse struct for AccountNotificationResponse
type AccountNotificationResponse struct {
	// Respond with any **2xx** HTTP status code to [accept the webhook](https://docs.adyen.com/development-resources/webhooks#accept-notifications).
	NotificationResponse *string `json:"notificationResponse,omitempty"`
}

// NewAccountNotificationResponse instantiates a new AccountNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountNotificationResponse() *AccountNotificationResponse {
	this := AccountNotificationResponse{}
	return &this
}

// NewAccountNotificationResponseWithDefaults instantiates a new AccountNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountNotificationResponseWithDefaults() *AccountNotificationResponse {
	this := AccountNotificationResponse{}
	return &this
}

// GetNotificationResponse returns the NotificationResponse field value if set, zero value otherwise.
func (o *AccountNotificationResponse) GetNotificationResponse() string {
	if o == nil || common.IsNil(o.NotificationResponse) {
		var ret string
		return ret
	}
	return *o.NotificationResponse
}

// GetNotificationResponseOk returns a tuple with the NotificationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNotificationResponse) GetNotificationResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotificationResponse) {
		return nil, false
	}
	return o.NotificationResponse, true
}

// HasNotificationResponse returns a boolean if a field has been set.
func (o *AccountNotificationResponse) HasNotificationResponse() bool {
	if o != nil && !common.IsNil(o.NotificationResponse) {
		return true
	}

	return false
}

// SetNotificationResponse gets a reference to the given string and assigns it to the NotificationResponse field.
func (o *AccountNotificationResponse) SetNotificationResponse(v string) {
	o.NotificationResponse = &v
}

func (o AccountNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NotificationResponse) {
		toSerialize["notificationResponse"] = o.NotificationResponse
	}
	return toSerialize, nil
}

type NullableAccountNotificationResponse struct {
	value *AccountNotificationResponse
	isSet bool
}

func (v NullableAccountNotificationResponse) Get() *AccountNotificationResponse {
	return v.value
}

func (v *NullableAccountNotificationResponse) Set(val *AccountNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountNotificationResponse(val *AccountNotificationResponse) *NullableAccountNotificationResponse {
	return &NullableAccountNotificationResponse{value: val, isSet: true}
}

func (v NullableAccountNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
