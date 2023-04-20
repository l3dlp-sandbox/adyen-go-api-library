/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the ThreeDSRequestorPriorAuthenticationInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDSRequestorPriorAuthenticationInfo{}

// ThreeDSRequestorPriorAuthenticationInfo struct for ThreeDSRequestorPriorAuthenticationInfo
type ThreeDSRequestorPriorAuthenticationInfo struct {
	// Data that documents and supports a specific authentication process. Maximum length: 2048 bytes.
	ThreeDSReqPriorAuthData *string `json:"threeDSReqPriorAuthData,omitempty"`
	// Mechanism used by the Cardholder to previously authenticate to the 3DS Requestor. Allowed values: * **01** — Frictionless authentication occurred by ACS. * **02** — Cardholder challenge occurred by ACS. * **03** — AVS verified. * **04** — Other issuer methods.
	ThreeDSReqPriorAuthMethod *string `json:"threeDSReqPriorAuthMethod,omitempty"`
	// Date and time in UTC of the prior cardholder authentication. Format: YYYYMMDDHHMM
	ThreeDSReqPriorAuthTimestamp *string `json:"threeDSReqPriorAuthTimestamp,omitempty"`
	// This data element provides additional information to the ACS to determine the best approach for handing a request. This data element contains an ACS Transaction ID for a prior authenticated transaction. For example, the first recurring transaction that was authenticated with the cardholder. Length: 30 characters.
	ThreeDSReqPriorRef *string `json:"threeDSReqPriorRef,omitempty"`
}

// NewThreeDSRequestorPriorAuthenticationInfo instantiates a new ThreeDSRequestorPriorAuthenticationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSRequestorPriorAuthenticationInfo() *ThreeDSRequestorPriorAuthenticationInfo {
	this := ThreeDSRequestorPriorAuthenticationInfo{}
	return &this
}

// NewThreeDSRequestorPriorAuthenticationInfoWithDefaults instantiates a new ThreeDSRequestorPriorAuthenticationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSRequestorPriorAuthenticationInfoWithDefaults() *ThreeDSRequestorPriorAuthenticationInfo {
	this := ThreeDSRequestorPriorAuthenticationInfo{}
	return &this
}

// GetThreeDSReqPriorAuthData returns the ThreeDSReqPriorAuthData field value if set, zero value otherwise.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthData() string {
	if o == nil || common.IsNil(o.ThreeDSReqPriorAuthData) {
		var ret string
		return ret
	}
	return *o.ThreeDSReqPriorAuthData
}

// GetThreeDSReqPriorAuthDataOk returns a tuple with the ThreeDSReqPriorAuthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSReqPriorAuthData) {
		return nil, false
	}
	return o.ThreeDSReqPriorAuthData, true
}

// HasThreeDSReqPriorAuthData returns a boolean if a field has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorAuthData() bool {
	if o != nil && !common.IsNil(o.ThreeDSReqPriorAuthData) {
		return true
	}

	return false
}

// SetThreeDSReqPriorAuthData gets a reference to the given string and assigns it to the ThreeDSReqPriorAuthData field.
func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorAuthData(v string) {
	o.ThreeDSReqPriorAuthData = &v
}

// GetThreeDSReqPriorAuthMethod returns the ThreeDSReqPriorAuthMethod field value if set, zero value otherwise.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthMethod() string {
	if o == nil || common.IsNil(o.ThreeDSReqPriorAuthMethod) {
		var ret string
		return ret
	}
	return *o.ThreeDSReqPriorAuthMethod
}

// GetThreeDSReqPriorAuthMethodOk returns a tuple with the ThreeDSReqPriorAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSReqPriorAuthMethod) {
		return nil, false
	}
	return o.ThreeDSReqPriorAuthMethod, true
}

// HasThreeDSReqPriorAuthMethod returns a boolean if a field has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorAuthMethod() bool {
	if o != nil && !common.IsNil(o.ThreeDSReqPriorAuthMethod) {
		return true
	}

	return false
}

// SetThreeDSReqPriorAuthMethod gets a reference to the given string and assigns it to the ThreeDSReqPriorAuthMethod field.
func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorAuthMethod(v string) {
	o.ThreeDSReqPriorAuthMethod = &v
}

// GetThreeDSReqPriorAuthTimestamp returns the ThreeDSReqPriorAuthTimestamp field value if set, zero value otherwise.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthTimestamp() string {
	if o == nil || common.IsNil(o.ThreeDSReqPriorAuthTimestamp) {
		var ret string
		return ret
	}
	return *o.ThreeDSReqPriorAuthTimestamp
}

// GetThreeDSReqPriorAuthTimestampOk returns a tuple with the ThreeDSReqPriorAuthTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthTimestampOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSReqPriorAuthTimestamp) {
		return nil, false
	}
	return o.ThreeDSReqPriorAuthTimestamp, true
}

// HasThreeDSReqPriorAuthTimestamp returns a boolean if a field has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorAuthTimestamp() bool {
	if o != nil && !common.IsNil(o.ThreeDSReqPriorAuthTimestamp) {
		return true
	}

	return false
}

// SetThreeDSReqPriorAuthTimestamp gets a reference to the given string and assigns it to the ThreeDSReqPriorAuthTimestamp field.
func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorAuthTimestamp(v string) {
	o.ThreeDSReqPriorAuthTimestamp = &v
}

// GetThreeDSReqPriorRef returns the ThreeDSReqPriorRef field value if set, zero value otherwise.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorRef() string {
	if o == nil || common.IsNil(o.ThreeDSReqPriorRef) {
		var ret string
		return ret
	}
	return *o.ThreeDSReqPriorRef
}

// GetThreeDSReqPriorRefOk returns a tuple with the ThreeDSReqPriorRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorRefOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSReqPriorRef) {
		return nil, false
	}
	return o.ThreeDSReqPriorRef, true
}

// HasThreeDSReqPriorRef returns a boolean if a field has been set.
func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorRef() bool {
	if o != nil && !common.IsNil(o.ThreeDSReqPriorRef) {
		return true
	}

	return false
}

// SetThreeDSReqPriorRef gets a reference to the given string and assigns it to the ThreeDSReqPriorRef field.
func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorRef(v string) {
	o.ThreeDSReqPriorRef = &v
}

func (o ThreeDSRequestorPriorAuthenticationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSRequestorPriorAuthenticationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ThreeDSReqPriorAuthData) {
		toSerialize["threeDSReqPriorAuthData"] = o.ThreeDSReqPriorAuthData
	}
	if !common.IsNil(o.ThreeDSReqPriorAuthMethod) {
		toSerialize["threeDSReqPriorAuthMethod"] = o.ThreeDSReqPriorAuthMethod
	}
	if !common.IsNil(o.ThreeDSReqPriorAuthTimestamp) {
		toSerialize["threeDSReqPriorAuthTimestamp"] = o.ThreeDSReqPriorAuthTimestamp
	}
	if !common.IsNil(o.ThreeDSReqPriorRef) {
		toSerialize["threeDSReqPriorRef"] = o.ThreeDSReqPriorRef
	}
	return toSerialize, nil
}

type NullableThreeDSRequestorPriorAuthenticationInfo struct {
	value *ThreeDSRequestorPriorAuthenticationInfo
	isSet bool
}

func (v NullableThreeDSRequestorPriorAuthenticationInfo) Get() *ThreeDSRequestorPriorAuthenticationInfo {
	return v.value
}

func (v *NullableThreeDSRequestorPriorAuthenticationInfo) Set(val *ThreeDSRequestorPriorAuthenticationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSRequestorPriorAuthenticationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSRequestorPriorAuthenticationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSRequestorPriorAuthenticationInfo(val *ThreeDSRequestorPriorAuthenticationInfo) *NullableThreeDSRequestorPriorAuthenticationInfo {
	return &NullableThreeDSRequestorPriorAuthenticationInfo{value: val, isSet: true}
}

func (v NullableThreeDSRequestorPriorAuthenticationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSRequestorPriorAuthenticationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ThreeDSRequestorPriorAuthenticationInfo) isValidThreeDSReqPriorAuthMethod() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04"}
	for _, allowed := range allowedEnumValues {
		if o.GetThreeDSReqPriorAuthMethod() == allowed {
			return true
		}
	}
	return false
}
