/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the Opi type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Opi{}

// Opi struct for Opi
type Opi struct {
	// Indicates if Pay at table is enabled.
	EnablePayAtTable *bool `json:"enablePayAtTable,omitempty"`
	// The store number to use for Pay at Table.
	PayAtTableStoreNumber *string `json:"payAtTableStoreNumber,omitempty"`
	// The URL and port number used for Pay at Table communication.
	PayAtTableURL *string `json:"payAtTableURL,omitempty"`
}

// NewOpi instantiates a new Opi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpi() *Opi {
	this := Opi{}
	return &this
}

// NewOpiWithDefaults instantiates a new Opi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpiWithDefaults() *Opi {
	this := Opi{}
	return &this
}

// GetEnablePayAtTable returns the EnablePayAtTable field value if set, zero value otherwise.
func (o *Opi) GetEnablePayAtTable() bool {
	if o == nil || common.IsNil(o.EnablePayAtTable) {
		var ret bool
		return ret
	}
	return *o.EnablePayAtTable
}

// GetEnablePayAtTableOk returns a tuple with the EnablePayAtTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opi) GetEnablePayAtTableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnablePayAtTable) {
		return nil, false
	}
	return o.EnablePayAtTable, true
}

// HasEnablePayAtTable returns a boolean if a field has been set.
func (o *Opi) HasEnablePayAtTable() bool {
	if o != nil && !common.IsNil(o.EnablePayAtTable) {
		return true
	}

	return false
}

// SetEnablePayAtTable gets a reference to the given bool and assigns it to the EnablePayAtTable field.
func (o *Opi) SetEnablePayAtTable(v bool) {
	o.EnablePayAtTable = &v
}

// GetPayAtTableStoreNumber returns the PayAtTableStoreNumber field value if set, zero value otherwise.
func (o *Opi) GetPayAtTableStoreNumber() string {
	if o == nil || common.IsNil(o.PayAtTableStoreNumber) {
		var ret string
		return ret
	}
	return *o.PayAtTableStoreNumber
}

// GetPayAtTableStoreNumberOk returns a tuple with the PayAtTableStoreNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opi) GetPayAtTableStoreNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayAtTableStoreNumber) {
		return nil, false
	}
	return o.PayAtTableStoreNumber, true
}

// HasPayAtTableStoreNumber returns a boolean if a field has been set.
func (o *Opi) HasPayAtTableStoreNumber() bool {
	if o != nil && !common.IsNil(o.PayAtTableStoreNumber) {
		return true
	}

	return false
}

// SetPayAtTableStoreNumber gets a reference to the given string and assigns it to the PayAtTableStoreNumber field.
func (o *Opi) SetPayAtTableStoreNumber(v string) {
	o.PayAtTableStoreNumber = &v
}

// GetPayAtTableURL returns the PayAtTableURL field value if set, zero value otherwise.
func (o *Opi) GetPayAtTableURL() string {
	if o == nil || common.IsNil(o.PayAtTableURL) {
		var ret string
		return ret
	}
	return *o.PayAtTableURL
}

// GetPayAtTableURLOk returns a tuple with the PayAtTableURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opi) GetPayAtTableURLOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayAtTableURL) {
		return nil, false
	}
	return o.PayAtTableURL, true
}

// HasPayAtTableURL returns a boolean if a field has been set.
func (o *Opi) HasPayAtTableURL() bool {
	if o != nil && !common.IsNil(o.PayAtTableURL) {
		return true
	}

	return false
}

// SetPayAtTableURL gets a reference to the given string and assigns it to the PayAtTableURL field.
func (o *Opi) SetPayAtTableURL(v string) {
	o.PayAtTableURL = &v
}

func (o Opi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Opi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EnablePayAtTable) {
		toSerialize["enablePayAtTable"] = o.EnablePayAtTable
	}
	if !common.IsNil(o.PayAtTableStoreNumber) {
		toSerialize["payAtTableStoreNumber"] = o.PayAtTableStoreNumber
	}
	if !common.IsNil(o.PayAtTableURL) {
		toSerialize["payAtTableURL"] = o.PayAtTableURL
	}
	return toSerialize, nil
}

type NullableOpi struct {
	value *Opi
	isSet bool
}

func (v NullableOpi) Get() *Opi {
	return v.value
}

func (v *NullableOpi) Set(val *Opi) {
	v.value = val
	v.isSet = true
}

func (v NullableOpi) IsSet() bool {
	return v.isSet
}

func (v *NullableOpi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpi(val *Opi) *NullableOpi {
	return &NullableOpi{value: val, isSet: true}
}

func (v NullableOpi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
