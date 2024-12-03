/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the PhoneNumber type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PhoneNumber{}

// PhoneNumber struct for PhoneNumber
type PhoneNumber struct {
	// The two-character ISO-3166-1 alpha-2 country code of the phone number. For example, **US** or **NL**.
	PhoneCountryCode *string `json:"phoneCountryCode,omitempty"`
	// The phone number. The inclusion of the phone number country code is not necessary.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The type of the phone number. Possible values: **Landline**, **Mobile**, **SIP**, **Fax**.
	PhoneType *string `json:"phoneType,omitempty"`
}

// NewPhoneNumber instantiates a new PhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumber() *PhoneNumber {
	this := PhoneNumber{}
	return &this
}

// NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumberWithDefaults() *PhoneNumber {
	this := PhoneNumber{}
	return &this
}

// GetPhoneCountryCode returns the PhoneCountryCode field value if set, zero value otherwise.
func (o *PhoneNumber) GetPhoneCountryCode() string {
	if o == nil || common.IsNil(o.PhoneCountryCode) {
		var ret string
		return ret
	}
	return *o.PhoneCountryCode
}

// GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetPhoneCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PhoneCountryCode) {
		return nil, false
	}
	return o.PhoneCountryCode, true
}

// HasPhoneCountryCode returns a boolean if a field has been set.
func (o *PhoneNumber) HasPhoneCountryCode() bool {
	if o != nil && !common.IsNil(o.PhoneCountryCode) {
		return true
	}

	return false
}

// SetPhoneCountryCode gets a reference to the given string and assigns it to the PhoneCountryCode field.
func (o *PhoneNumber) SetPhoneCountryCode(v string) {
	o.PhoneCountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PhoneNumber) GetPhoneNumber() string {
	if o == nil || common.IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetPhoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PhoneNumber) HasPhoneNumber() bool {
	if o != nil && !common.IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PhoneNumber) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneType returns the PhoneType field value if set, zero value otherwise.
func (o *PhoneNumber) GetPhoneType() string {
	if o == nil || common.IsNil(o.PhoneType) {
		var ret string
		return ret
	}
	return *o.PhoneType
}

// GetPhoneTypeOk returns a tuple with the PhoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetPhoneTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PhoneType) {
		return nil, false
	}
	return o.PhoneType, true
}

// HasPhoneType returns a boolean if a field has been set.
func (o *PhoneNumber) HasPhoneType() bool {
	if o != nil && !common.IsNil(o.PhoneType) {
		return true
	}

	return false
}

// SetPhoneType gets a reference to the given string and assigns it to the PhoneType field.
func (o *PhoneNumber) SetPhoneType(v string) {
	o.PhoneType = &v
}

func (o PhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PhoneCountryCode) {
		toSerialize["phoneCountryCode"] = o.PhoneCountryCode
	}
	if !common.IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !common.IsNil(o.PhoneType) {
		toSerialize["phoneType"] = o.PhoneType
	}
	return toSerialize, nil
}

type NullablePhoneNumber struct {
	value *PhoneNumber
	isSet bool
}

func (v NullablePhoneNumber) Get() *PhoneNumber {
	return v.value
}

func (v *NullablePhoneNumber) Set(val *PhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumber(val *PhoneNumber) *NullablePhoneNumber {
	return &NullablePhoneNumber{value: val, isSet: true}
}

func (v NullablePhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *PhoneNumber) isValidPhoneType() bool {
    var allowedEnumValues = []string{ "Fax", "Landline", "Mobile", "SIP" }
    for _, allowed := range allowedEnumValues {
        if o.GetPhoneType() == allowed {
            return true
        }
    }
    return false
}

