/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Phone type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Phone{}

// Phone struct for Phone
type Phone struct {
	// Country code. Length: 1–3 characters.
	Cc *string `json:"cc,omitempty"`
	// Subscriber number. Maximum length: 15 characters.
	Subscriber *string `json:"subscriber,omitempty"`
}

// NewPhone instantiates a new Phone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhone() *Phone {
	this := Phone{}
	return &this
}

// NewPhoneWithDefaults instantiates a new Phone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneWithDefaults() *Phone {
	this := Phone{}
	return &this
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *Phone) GetCc() string {
	if o == nil || common.IsNil(o.Cc) {
		var ret string
		return ret
	}
	return *o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetCcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cc) {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *Phone) HasCc() bool {
	if o != nil && !common.IsNil(o.Cc) {
		return true
	}

	return false
}

// SetCc gets a reference to the given string and assigns it to the Cc field.
func (o *Phone) SetCc(v string) {
	o.Cc = &v
}

// GetSubscriber returns the Subscriber field value if set, zero value otherwise.
func (o *Phone) GetSubscriber() string {
	if o == nil || common.IsNil(o.Subscriber) {
		var ret string
		return ret
	}
	return *o.Subscriber
}

// GetSubscriberOk returns a tuple with the Subscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetSubscriberOk() (*string, bool) {
	if o == nil || common.IsNil(o.Subscriber) {
		return nil, false
	}
	return o.Subscriber, true
}

// HasSubscriber returns a boolean if a field has been set.
func (o *Phone) HasSubscriber() bool {
	if o != nil && !common.IsNil(o.Subscriber) {
		return true
	}

	return false
}

// SetSubscriber gets a reference to the given string and assigns it to the Subscriber field.
func (o *Phone) SetSubscriber(v string) {
	o.Subscriber = &v
}

func (o Phone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Phone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Cc) {
		toSerialize["cc"] = o.Cc
	}
	if !common.IsNil(o.Subscriber) {
		toSerialize["subscriber"] = o.Subscriber
	}
	return toSerialize, nil
}

type NullablePhone struct {
	value *Phone
	isSet bool
}

func (v NullablePhone) Get() *Phone {
	return v.value
}

func (v *NullablePhone) Set(val *Phone) {
	v.value = val
	v.isSet = true
}

func (v NullablePhone) IsSet() bool {
	return v.isSet
}

func (v *NullablePhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhone(val *Phone) *NullablePhone {
	return &NullablePhone{value: val, isSet: true}
}

func (v NullablePhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
