/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the Phone type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Phone{}

// Phone struct for Phone
type Phone struct {
	// The full phone number provided as a single string.  For example, **\"0031 6 11 22 33 44\"**, **\"+316/1122-3344\"**,    or **\"(0031) 611223344\"**.
	Number string `json:"number"`
	// Type of phone number. Possible values:  **Landline**, **Mobile**.
	Type string `json:"type"`
}

// NewPhone instantiates a new Phone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhone(number string, type_ string) *Phone {
	this := Phone{}
	this.Number = number
	this.Type = type_
	return &this
}

// NewPhoneWithDefaults instantiates a new Phone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneWithDefaults() *Phone {
	this := Phone{}
	return &this
}

// GetNumber returns the Number field value
func (o *Phone) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Phone) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Phone) SetNumber(v string) {
	o.Number = v
}

// GetType returns the Type field value
func (o *Phone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Phone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Phone) SetType(v string) {
	o.Type = v
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
	toSerialize["number"] = o.Number
	toSerialize["type"] = o.Type
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

func (o *Phone) isValidType() bool {
	var allowedEnumValues = []string{"landline", "mobile"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
