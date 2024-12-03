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

// checks if the Key type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Key{}

// Key struct for Key
type Key struct {
	// The unique identifier of the shared key.
	Identifier *string `json:"identifier,omitempty"`
	// The secure passphrase to protect the shared key.
	Passphrase *string `json:"passphrase,omitempty"`
	// The version number of the shared key.
	Version *int32 `json:"version,omitempty"`
}

// NewKey instantiates a new Key object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKey() *Key {
	this := Key{}
	return &this
}

// NewKeyWithDefaults instantiates a new Key object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyWithDefaults() *Key {
	this := Key{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Key) GetIdentifier() string {
	if o == nil || common.IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetIdentifierOk() (*string, bool) {
	if o == nil || common.IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Key) HasIdentifier() bool {
	if o != nil && !common.IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Key) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *Key) GetPassphrase() string {
	if o == nil || common.IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetPassphraseOk() (*string, bool) {
	if o == nil || common.IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *Key) HasPassphrase() bool {
	if o != nil && !common.IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *Key) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Key) GetVersion() int32 {
	if o == nil || common.IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetVersionOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Key) HasVersion() bool {
	if o != nil && !common.IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Key) SetVersion(v int32) {
	o.Version = &v
}

func (o Key) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Key) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !common.IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if !common.IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableKey struct {
	value *Key
	isSet bool
}

func (v NullableKey) Get() *Key {
	return v.value
}

func (v *NullableKey) Set(val *Key) {
	v.value = val
	v.isSet = true
}

func (v NullableKey) IsSet() bool {
	return v.isSet
}

func (v *NullableKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKey(val *Key) *NullableKey {
	return &NullableKey{value: val, isSet: true}
}

func (v NullableKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



