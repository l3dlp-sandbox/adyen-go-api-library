/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the Url type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Url{}

// Url struct for Url
type Url struct {
	// Indicates if the message sent to this URL should be encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
	// The password for authentication of the notifications.
	Password *string `json:"password,omitempty"`
	// The URL in the format: http(s)://domain.com.
	Url *string `json:"url,omitempty"`
	// The username for authentication of the notifications.
	Username *string `json:"username,omitempty"`
}

// NewUrl instantiates a new Url object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrl() *Url {
	this := Url{}
	return &this
}

// NewUrlWithDefaults instantiates a new Url object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlWithDefaults() *Url {
	this := Url{}
	return &this
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *Url) GetEncrypted() bool {
	if o == nil || common.IsNil(o.Encrypted) {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Url) GetEncryptedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *Url) HasEncrypted() bool {
	if o != nil && !common.IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *Url) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Url) GetPassword() string {
	if o == nil || common.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Url) GetPasswordOk() (*string, bool) {
	if o == nil || common.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Url) HasPassword() bool {
	if o != nil && !common.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Url) SetPassword(v string) {
	o.Password = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Url) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Url) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Url) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Url) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Url) GetUsername() string {
	if o == nil || common.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Url) GetUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Url) HasUsername() bool {
	if o != nil && !common.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Url) SetUsername(v string) {
	o.Username = &v
}

func (o Url) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Url) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	if !common.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !common.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableUrl struct {
	value *Url
	isSet bool
}

func (v NullableUrl) Get() *Url {
	return v.value
}

func (v *NullableUrl) Set(val *Url) {
	v.value = val
	v.isSet = true
}

func (v NullableUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrl(val *Url) *NullableUrl {
	return &NullableUrl{value: val, isSet: true}
}

func (v NullableUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
