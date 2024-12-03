/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Nexo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Nexo{}

// Nexo struct for Nexo
type Nexo struct {
	DisplayUrls   *NotificationUrl `json:"displayUrls,omitempty"`
	EncryptionKey *Key             `json:"encryptionKey,omitempty"`
	EventUrls     *EventUrl        `json:"eventUrls,omitempty"`
	// One or more URLs to send event messages to when using Terminal API.
	// Deprecated since Management API v1
	// Use `eventUrls` instead.
	NexoEventUrls []string      `json:"nexoEventUrls,omitempty"`
	Notification  *Notification `json:"notification,omitempty"`
}

// NewNexo instantiates a new Nexo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNexo() *Nexo {
	this := Nexo{}
	return &this
}

// NewNexoWithDefaults instantiates a new Nexo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNexoWithDefaults() *Nexo {
	this := Nexo{}
	return &this
}

// GetDisplayUrls returns the DisplayUrls field value if set, zero value otherwise.
func (o *Nexo) GetDisplayUrls() NotificationUrl {
	if o == nil || common.IsNil(o.DisplayUrls) {
		var ret NotificationUrl
		return ret
	}
	return *o.DisplayUrls
}

// GetDisplayUrlsOk returns a tuple with the DisplayUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nexo) GetDisplayUrlsOk() (*NotificationUrl, bool) {
	if o == nil || common.IsNil(o.DisplayUrls) {
		return nil, false
	}
	return o.DisplayUrls, true
}

// HasDisplayUrls returns a boolean if a field has been set.
func (o *Nexo) HasDisplayUrls() bool {
	if o != nil && !common.IsNil(o.DisplayUrls) {
		return true
	}

	return false
}

// SetDisplayUrls gets a reference to the given NotificationUrl and assigns it to the DisplayUrls field.
func (o *Nexo) SetDisplayUrls(v NotificationUrl) {
	o.DisplayUrls = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *Nexo) GetEncryptionKey() Key {
	if o == nil || common.IsNil(o.EncryptionKey) {
		var ret Key
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nexo) GetEncryptionKeyOk() (*Key, bool) {
	if o == nil || common.IsNil(o.EncryptionKey) {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *Nexo) HasEncryptionKey() bool {
	if o != nil && !common.IsNil(o.EncryptionKey) {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given Key and assigns it to the EncryptionKey field.
func (o *Nexo) SetEncryptionKey(v Key) {
	o.EncryptionKey = &v
}

// GetEventUrls returns the EventUrls field value if set, zero value otherwise.
func (o *Nexo) GetEventUrls() EventUrl {
	if o == nil || common.IsNil(o.EventUrls) {
		var ret EventUrl
		return ret
	}
	return *o.EventUrls
}

// GetEventUrlsOk returns a tuple with the EventUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nexo) GetEventUrlsOk() (*EventUrl, bool) {
	if o == nil || common.IsNil(o.EventUrls) {
		return nil, false
	}
	return o.EventUrls, true
}

// HasEventUrls returns a boolean if a field has been set.
func (o *Nexo) HasEventUrls() bool {
	if o != nil && !common.IsNil(o.EventUrls) {
		return true
	}

	return false
}

// SetEventUrls gets a reference to the given EventUrl and assigns it to the EventUrls field.
func (o *Nexo) SetEventUrls(v EventUrl) {
	o.EventUrls = &v
}

// GetNexoEventUrls returns the NexoEventUrls field value if set, zero value otherwise.
// Deprecated since Management API v1
// Use `eventUrls` instead.
func (o *Nexo) GetNexoEventUrls() []string {
	if o == nil || common.IsNil(o.NexoEventUrls) {
		var ret []string
		return ret
	}
	return o.NexoEventUrls
}

// GetNexoEventUrlsOk returns a tuple with the NexoEventUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Management API v1
// Use `eventUrls` instead.
func (o *Nexo) GetNexoEventUrlsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.NexoEventUrls) {
		return nil, false
	}
	return o.NexoEventUrls, true
}

// HasNexoEventUrls returns a boolean if a field has been set.
func (o *Nexo) HasNexoEventUrls() bool {
	if o != nil && !common.IsNil(o.NexoEventUrls) {
		return true
	}

	return false
}

// SetNexoEventUrls gets a reference to the given []string and assigns it to the NexoEventUrls field.
// Deprecated since Management API v1
// Use `eventUrls` instead.
func (o *Nexo) SetNexoEventUrls(v []string) {
	o.NexoEventUrls = v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *Nexo) GetNotification() Notification {
	if o == nil || common.IsNil(o.Notification) {
		var ret Notification
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nexo) GetNotificationOk() (*Notification, bool) {
	if o == nil || common.IsNil(o.Notification) {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *Nexo) HasNotification() bool {
	if o != nil && !common.IsNil(o.Notification) {
		return true
	}

	return false
}

// SetNotification gets a reference to the given Notification and assigns it to the Notification field.
func (o *Nexo) SetNotification(v Notification) {
	o.Notification = &v
}

func (o Nexo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Nexo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DisplayUrls) {
		toSerialize["displayUrls"] = o.DisplayUrls
	}
	if !common.IsNil(o.EncryptionKey) {
		toSerialize["encryptionKey"] = o.EncryptionKey
	}
	if !common.IsNil(o.EventUrls) {
		toSerialize["eventUrls"] = o.EventUrls
	}
	if !common.IsNil(o.NexoEventUrls) {
		toSerialize["nexoEventUrls"] = o.NexoEventUrls
	}
	if !common.IsNil(o.Notification) {
		toSerialize["notification"] = o.Notification
	}
	return toSerialize, nil
}

type NullableNexo struct {
	value *Nexo
	isSet bool
}

func (v NullableNexo) Get() *Nexo {
	return v.value
}

func (v *NullableNexo) Set(val *Nexo) {
	v.value = val
	v.isSet = true
}

func (v NullableNexo) IsSet() bool {
	return v.isSet
}

func (v *NullableNexo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNexo(val *Nexo) *NullableNexo {
	return &NullableNexo{value: val, isSet: true}
}

func (v NullableNexo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNexo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
