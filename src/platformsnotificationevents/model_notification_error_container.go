/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// NotificationErrorContainer struct for NotificationErrorContainer
type NotificationErrorContainer struct {
	// The Adyen code that is mapped to the error message.
	ErrorCode *string `json:"errorCode,omitempty"`
	// A short explanation of the issue.
	Message *string `json:"message,omitempty"`
}

// NewNotificationErrorContainer instantiates a new NotificationErrorContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationErrorContainer() *NotificationErrorContainer {
	this := NotificationErrorContainer{}
	return &this
}

// NewNotificationErrorContainerWithDefaults instantiates a new NotificationErrorContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationErrorContainerWithDefaults() *NotificationErrorContainer {
	this := NotificationErrorContainer{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *NotificationErrorContainer) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationErrorContainer) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *NotificationErrorContainer) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *NotificationErrorContainer) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationErrorContainer) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationErrorContainer) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationErrorContainer) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotificationErrorContainer) SetMessage(v string) {
	o.Message = &v
}

func (o NotificationErrorContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationErrorContainer struct {
	value *NotificationErrorContainer
	isSet bool
}

func (v NullableNotificationErrorContainer) Get() *NotificationErrorContainer {
	return v.value
}

func (v *NullableNotificationErrorContainer) Set(val *NotificationErrorContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationErrorContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationErrorContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationErrorContainer(val *NotificationErrorContainer) *NullableNotificationErrorContainer {
	return &NullableNotificationErrorContainer{value: val, isSet: true}
}

func (v NullableNotificationErrorContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationErrorContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

