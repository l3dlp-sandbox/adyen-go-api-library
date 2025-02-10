/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the TerminalBoardingNotificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalBoardingNotificationResponse{}

// TerminalBoardingNotificationResponse struct for TerminalBoardingNotificationResponse
type TerminalBoardingNotificationResponse struct {
	// Respond with any **2xx** HTTP status code to [accept the webhook](https://docs.adyen.com/development-resources/webhooks#accept-notifications).
	NotificationResponse *string `json:"notificationResponse,omitempty"`
}

// NewTerminalBoardingNotificationResponse instantiates a new TerminalBoardingNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalBoardingNotificationResponse() *TerminalBoardingNotificationResponse {
	this := TerminalBoardingNotificationResponse{}
	return &this
}

// NewTerminalBoardingNotificationResponseWithDefaults instantiates a new TerminalBoardingNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalBoardingNotificationResponseWithDefaults() *TerminalBoardingNotificationResponse {
	this := TerminalBoardingNotificationResponse{}
	return &this
}

// GetNotificationResponse returns the NotificationResponse field value if set, zero value otherwise.
func (o *TerminalBoardingNotificationResponse) GetNotificationResponse() string {
	if o == nil || common.IsNil(o.NotificationResponse) {
		var ret string
		return ret
	}
	return *o.NotificationResponse
}

// GetNotificationResponseOk returns a tuple with the NotificationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalBoardingNotificationResponse) GetNotificationResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotificationResponse) {
		return nil, false
	}
	return o.NotificationResponse, true
}

// HasNotificationResponse returns a boolean if a field has been set.
func (o *TerminalBoardingNotificationResponse) HasNotificationResponse() bool {
	if o != nil && !common.IsNil(o.NotificationResponse) {
		return true
	}

	return false
}

// SetNotificationResponse gets a reference to the given string and assigns it to the NotificationResponse field.
func (o *TerminalBoardingNotificationResponse) SetNotificationResponse(v string) {
	o.NotificationResponse = &v
}

func (o TerminalBoardingNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalBoardingNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NotificationResponse) {
		toSerialize["notificationResponse"] = o.NotificationResponse
	}
	return toSerialize, nil
}

type NullableTerminalBoardingNotificationResponse struct {
	value *TerminalBoardingNotificationResponse
	isSet bool
}

func (v NullableTerminalBoardingNotificationResponse) Get() *TerminalBoardingNotificationResponse {
	return v.value
}

func (v *NullableTerminalBoardingNotificationResponse) Set(val *TerminalBoardingNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalBoardingNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalBoardingNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalBoardingNotificationResponse(val *TerminalBoardingNotificationResponse) *NullableTerminalBoardingNotificationResponse {
	return &NullableTerminalBoardingNotificationResponse{value: val, isSet: true}
}

func (v NullableTerminalBoardingNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalBoardingNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
