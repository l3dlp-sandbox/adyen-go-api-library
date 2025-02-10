/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the RecurringDetailWrapper type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecurringDetailWrapper{}

// RecurringDetailWrapper struct for RecurringDetailWrapper
type RecurringDetailWrapper struct {
	RecurringDetail *RecurringDetail `json:"RecurringDetail,omitempty"`
}

// NewRecurringDetailWrapper instantiates a new RecurringDetailWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringDetailWrapper() *RecurringDetailWrapper {
	this := RecurringDetailWrapper{}
	return &this
}

// NewRecurringDetailWrapperWithDefaults instantiates a new RecurringDetailWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringDetailWrapperWithDefaults() *RecurringDetailWrapper {
	this := RecurringDetailWrapper{}
	return &this
}

// GetRecurringDetail returns the RecurringDetail field value if set, zero value otherwise.
func (o *RecurringDetailWrapper) GetRecurringDetail() RecurringDetail {
	if o == nil || common.IsNil(o.RecurringDetail) {
		var ret RecurringDetail
		return ret
	}
	return *o.RecurringDetail
}

// GetRecurringDetailOk returns a tuple with the RecurringDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetailWrapper) GetRecurringDetailOk() (*RecurringDetail, bool) {
	if o == nil || common.IsNil(o.RecurringDetail) {
		return nil, false
	}
	return o.RecurringDetail, true
}

// HasRecurringDetail returns a boolean if a field has been set.
func (o *RecurringDetailWrapper) HasRecurringDetail() bool {
	if o != nil && !common.IsNil(o.RecurringDetail) {
		return true
	}

	return false
}

// SetRecurringDetail gets a reference to the given RecurringDetail and assigns it to the RecurringDetail field.
func (o *RecurringDetailWrapper) SetRecurringDetail(v RecurringDetail) {
	o.RecurringDetail = &v
}

func (o RecurringDetailWrapper) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurringDetailWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RecurringDetail) {
		toSerialize["RecurringDetail"] = o.RecurringDetail
	}
	return toSerialize, nil
}

type NullableRecurringDetailWrapper struct {
	value *RecurringDetailWrapper
	isSet bool
}

func (v NullableRecurringDetailWrapper) Get() *RecurringDetailWrapper {
	return v.value
}

func (v *NullableRecurringDetailWrapper) Set(val *RecurringDetailWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringDetailWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringDetailWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringDetailWrapper(val *RecurringDetailWrapper) *NullableRecurringDetailWrapper {
	return &NullableRecurringDetailWrapper{value: val, isSet: true}
}

func (v NullableRecurringDetailWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringDetailWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
