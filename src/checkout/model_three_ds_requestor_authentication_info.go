/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the ThreeDSRequestorAuthenticationInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDSRequestorAuthenticationInfo{}

// ThreeDSRequestorAuthenticationInfo struct for ThreeDSRequestorAuthenticationInfo
type ThreeDSRequestorAuthenticationInfo struct {
	// Data that documents and supports a specific authentication process. Maximum length: 2048 bytes.
	ThreeDSReqAuthData *string `json:"threeDSReqAuthData,omitempty"`
	// Mechanism used by the Cardholder to authenticate to the 3DS Requestor. Allowed values: * **01** — No 3DS Requestor authentication occurred (for example, cardholder “logged in” as guest). * **02** — Login to the cardholder account at the 3DS Requestor system using 3DS Requestor’s own credentials. * **03** — Login to the cardholder account at the 3DS Requestor system using federated ID. * **04** — Login to the cardholder account at the 3DS Requestor system using issuer credentials. * **05** — Login to the cardholder account at the 3DS Requestor system using third-party authentication. * **06** — Login to the cardholder account at the 3DS Requestor system using FIDO Authenticator.
	ThreeDSReqAuthMethod *string `json:"threeDSReqAuthMethod,omitempty"`
	// Date and time in UTC of the cardholder authentication. Format: YYYYMMDDHHMM
	ThreeDSReqAuthTimestamp *string `json:"threeDSReqAuthTimestamp,omitempty"`
}

// NewThreeDSRequestorAuthenticationInfo instantiates a new ThreeDSRequestorAuthenticationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSRequestorAuthenticationInfo() *ThreeDSRequestorAuthenticationInfo {
	this := ThreeDSRequestorAuthenticationInfo{}
	return &this
}

// NewThreeDSRequestorAuthenticationInfoWithDefaults instantiates a new ThreeDSRequestorAuthenticationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSRequestorAuthenticationInfoWithDefaults() *ThreeDSRequestorAuthenticationInfo {
	this := ThreeDSRequestorAuthenticationInfo{}
	return &this
}

// GetThreeDSReqAuthData returns the ThreeDSReqAuthData field value if set, zero value otherwise.
func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthData() string {
	if o == nil || common.IsNil(o.ThreeDSReqAuthData) {
		var ret string
		return ret
	}
	return *o.ThreeDSReqAuthData
}

// GetThreeDSReqAuthDataOk returns a tuple with the ThreeDSReqAuthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSReqAuthData) {
		return nil, false
	}
	return o.ThreeDSReqAuthData, true
}

// HasThreeDSReqAuthData returns a boolean if a field has been set.
func (o *ThreeDSRequestorAuthenticationInfo) HasThreeDSReqAuthData() bool {
	if o != nil && !common.IsNil(o.ThreeDSReqAuthData) {
		return true
	}

	return false
}

// SetThreeDSReqAuthData gets a reference to the given string and assigns it to the ThreeDSReqAuthData field.
func (o *ThreeDSRequestorAuthenticationInfo) SetThreeDSReqAuthData(v string) {
	o.ThreeDSReqAuthData = &v
}

// GetThreeDSReqAuthMethod returns the ThreeDSReqAuthMethod field value if set, zero value otherwise.
func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthMethod() string {
	if o == nil || common.IsNil(o.ThreeDSReqAuthMethod) {
		var ret string
		return ret
	}
	return *o.ThreeDSReqAuthMethod
}

// GetThreeDSReqAuthMethodOk returns a tuple with the ThreeDSReqAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSReqAuthMethod) {
		return nil, false
	}
	return o.ThreeDSReqAuthMethod, true
}

// HasThreeDSReqAuthMethod returns a boolean if a field has been set.
func (o *ThreeDSRequestorAuthenticationInfo) HasThreeDSReqAuthMethod() bool {
	if o != nil && !common.IsNil(o.ThreeDSReqAuthMethod) {
		return true
	}

	return false
}

// SetThreeDSReqAuthMethod gets a reference to the given string and assigns it to the ThreeDSReqAuthMethod field.
func (o *ThreeDSRequestorAuthenticationInfo) SetThreeDSReqAuthMethod(v string) {
	o.ThreeDSReqAuthMethod = &v
}

// GetThreeDSReqAuthTimestamp returns the ThreeDSReqAuthTimestamp field value if set, zero value otherwise.
func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthTimestamp() string {
	if o == nil || common.IsNil(o.ThreeDSReqAuthTimestamp) {
		var ret string
		return ret
	}
	return *o.ThreeDSReqAuthTimestamp
}

// GetThreeDSReqAuthTimestampOk returns a tuple with the ThreeDSReqAuthTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthTimestampOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSReqAuthTimestamp) {
		return nil, false
	}
	return o.ThreeDSReqAuthTimestamp, true
}

// HasThreeDSReqAuthTimestamp returns a boolean if a field has been set.
func (o *ThreeDSRequestorAuthenticationInfo) HasThreeDSReqAuthTimestamp() bool {
	if o != nil && !common.IsNil(o.ThreeDSReqAuthTimestamp) {
		return true
	}

	return false
}

// SetThreeDSReqAuthTimestamp gets a reference to the given string and assigns it to the ThreeDSReqAuthTimestamp field.
func (o *ThreeDSRequestorAuthenticationInfo) SetThreeDSReqAuthTimestamp(v string) {
	o.ThreeDSReqAuthTimestamp = &v
}

func (o ThreeDSRequestorAuthenticationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSRequestorAuthenticationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ThreeDSReqAuthData) {
		toSerialize["threeDSReqAuthData"] = o.ThreeDSReqAuthData
	}
	if !common.IsNil(o.ThreeDSReqAuthMethod) {
		toSerialize["threeDSReqAuthMethod"] = o.ThreeDSReqAuthMethod
	}
	if !common.IsNil(o.ThreeDSReqAuthTimestamp) {
		toSerialize["threeDSReqAuthTimestamp"] = o.ThreeDSReqAuthTimestamp
	}
	return toSerialize, nil
}

type NullableThreeDSRequestorAuthenticationInfo struct {
	value *ThreeDSRequestorAuthenticationInfo
	isSet bool
}

func (v NullableThreeDSRequestorAuthenticationInfo) Get() *ThreeDSRequestorAuthenticationInfo {
	return v.value
}

func (v *NullableThreeDSRequestorAuthenticationInfo) Set(val *ThreeDSRequestorAuthenticationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSRequestorAuthenticationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSRequestorAuthenticationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSRequestorAuthenticationInfo(val *ThreeDSRequestorAuthenticationInfo) *NullableThreeDSRequestorAuthenticationInfo {
	return &NullableThreeDSRequestorAuthenticationInfo{value: val, isSet: true}
}

func (v NullableThreeDSRequestorAuthenticationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSRequestorAuthenticationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ThreeDSRequestorAuthenticationInfo) isValidThreeDSReqAuthMethod() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05", "06"}
	for _, allowed := range allowedEnumValues {
		if o.GetThreeDSReqAuthMethod() == allowed {
			return true
		}
	}
	return false
}
