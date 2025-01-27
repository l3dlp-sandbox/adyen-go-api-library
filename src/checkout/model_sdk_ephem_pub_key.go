/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the SDKEphemPubKey type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SDKEphemPubKey{}

// SDKEphemPubKey struct for SDKEphemPubKey
type SDKEphemPubKey struct {
	// The `crv` value as received from the 3D Secure 2 SDK.
	Crv *string `json:"crv,omitempty"`
	// The `kty` value as received from the 3D Secure 2 SDK.
	Kty *string `json:"kty,omitempty"`
	// The `x` value as received from the 3D Secure 2 SDK.
	X *string `json:"x,omitempty"`
	// The `y` value as received from the 3D Secure 2 SDK.
	Y *string `json:"y,omitempty"`
}

// NewSDKEphemPubKey instantiates a new SDKEphemPubKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDKEphemPubKey() *SDKEphemPubKey {
	this := SDKEphemPubKey{}
	return &this
}

// NewSDKEphemPubKeyWithDefaults instantiates a new SDKEphemPubKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDKEphemPubKeyWithDefaults() *SDKEphemPubKey {
	this := SDKEphemPubKey{}
	return &this
}

// GetCrv returns the Crv field value if set, zero value otherwise.
func (o *SDKEphemPubKey) GetCrv() string {
	if o == nil || common.IsNil(o.Crv) {
		var ret string
		return ret
	}
	return *o.Crv
}

// GetCrvOk returns a tuple with the Crv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDKEphemPubKey) GetCrvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Crv) {
		return nil, false
	}
	return o.Crv, true
}

// HasCrv returns a boolean if a field has been set.
func (o *SDKEphemPubKey) HasCrv() bool {
	if o != nil && !common.IsNil(o.Crv) {
		return true
	}

	return false
}

// SetCrv gets a reference to the given string and assigns it to the Crv field.
func (o *SDKEphemPubKey) SetCrv(v string) {
	o.Crv = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *SDKEphemPubKey) GetKty() string {
	if o == nil || common.IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDKEphemPubKey) GetKtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *SDKEphemPubKey) HasKty() bool {
	if o != nil && !common.IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *SDKEphemPubKey) SetKty(v string) {
	o.Kty = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SDKEphemPubKey) GetX() string {
	if o == nil || common.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDKEphemPubKey) GetXOk() (*string, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SDKEphemPubKey) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *SDKEphemPubKey) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SDKEphemPubKey) GetY() string {
	if o == nil || common.IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDKEphemPubKey) GetYOk() (*string, bool) {
	if o == nil || common.IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SDKEphemPubKey) HasY() bool {
	if o != nil && !common.IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *SDKEphemPubKey) SetY(v string) {
	o.Y = &v
}

func (o SDKEphemPubKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SDKEphemPubKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Crv) {
		toSerialize["crv"] = o.Crv
	}
	if !common.IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !common.IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !common.IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableSDKEphemPubKey struct {
	value *SDKEphemPubKey
	isSet bool
}

func (v NullableSDKEphemPubKey) Get() *SDKEphemPubKey {
	return v.value
}

func (v *NullableSDKEphemPubKey) Set(val *SDKEphemPubKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSDKEphemPubKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSDKEphemPubKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDKEphemPubKey(val *SDKEphemPubKey) *NullableSDKEphemPubKey {
	return &NullableSDKEphemPubKey{value: val, isSet: true}
}

func (v NullableSDKEphemPubKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSDKEphemPubKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
