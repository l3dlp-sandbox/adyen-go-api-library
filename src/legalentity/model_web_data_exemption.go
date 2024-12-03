/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the WebDataExemption type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WebDataExemption{}

// WebDataExemption struct for WebDataExemption
type WebDataExemption struct {
	// The reason why the web data was not provided. Possible value: **noOnlinePresence**.
	Reason *string `json:"reason,omitempty"`
}

// NewWebDataExemption instantiates a new WebDataExemption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebDataExemption() *WebDataExemption {
	this := WebDataExemption{}
	return &this
}

// NewWebDataExemptionWithDefaults instantiates a new WebDataExemption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebDataExemptionWithDefaults() *WebDataExemption {
	this := WebDataExemption{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *WebDataExemption) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebDataExemption) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *WebDataExemption) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *WebDataExemption) SetReason(v string) {
	o.Reason = &v
}

func (o WebDataExemption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebDataExemption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableWebDataExemption struct {
	value *WebDataExemption
	isSet bool
}

func (v NullableWebDataExemption) Get() *WebDataExemption {
	return v.value
}

func (v *NullableWebDataExemption) Set(val *WebDataExemption) {
	v.value = val
	v.isSet = true
}

func (v NullableWebDataExemption) IsSet() bool {
	return v.isSet
}

func (v *NullableWebDataExemption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebDataExemption(val *WebDataExemption) *NullableWebDataExemption {
	return &NullableWebDataExemption{value: val, isSet: true}
}

func (v NullableWebDataExemption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebDataExemption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *WebDataExemption) isValidReason() bool {
    var allowedEnumValues = []string{ "noOnlinePresence", "notCollectedDuringOnboarding" }
    for _, allowed := range allowedEnumValues {
        if o.GetReason() == allowed {
            return true
        }
    }
    return false
}

