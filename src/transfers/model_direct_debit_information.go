/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the DirectDebitInformation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DirectDebitInformation{}

// DirectDebitInformation struct for DirectDebitInformation
type DirectDebitInformation struct {
	// The date when the direct debit mandate was accepted by your user, in [ISO-8601](https://www.w3.org/TR/NOTE-datetime) format.
	DateOfSignature *time.Time `json:"dateOfSignature,omitempty"`
	// The date when the funds are deducted from your user's balance account.
	DueDate *time.Time `json:"dueDate,omitempty"`
	// Your unique identifier for the direct debit mandate.
	MandateId *string `json:"mandateId,omitempty"`
	// Identifies the direct debit transfer's type. Possible values: **OneOff**, **First**, **Recurring**, **Final**.
	SequenceType *string `json:"sequenceType,omitempty"`
}

// NewDirectDebitInformation instantiates a new DirectDebitInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitInformation() *DirectDebitInformation {
	this := DirectDebitInformation{}
	return &this
}

// NewDirectDebitInformationWithDefaults instantiates a new DirectDebitInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitInformationWithDefaults() *DirectDebitInformation {
	this := DirectDebitInformation{}
	return &this
}

// GetDateOfSignature returns the DateOfSignature field value if set, zero value otherwise.
func (o *DirectDebitInformation) GetDateOfSignature() time.Time {
	if o == nil || common.IsNil(o.DateOfSignature) {
		var ret time.Time
		return ret
	}
	return *o.DateOfSignature
}

// GetDateOfSignatureOk returns a tuple with the DateOfSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitInformation) GetDateOfSignatureOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.DateOfSignature) {
		return nil, false
	}
	return o.DateOfSignature, true
}

// HasDateOfSignature returns a boolean if a field has been set.
func (o *DirectDebitInformation) HasDateOfSignature() bool {
	if o != nil && !common.IsNil(o.DateOfSignature) {
		return true
	}

	return false
}

// SetDateOfSignature gets a reference to the given time.Time and assigns it to the DateOfSignature field.
func (o *DirectDebitInformation) SetDateOfSignature(v time.Time) {
	o.DateOfSignature = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *DirectDebitInformation) GetDueDate() time.Time {
	if o == nil || common.IsNil(o.DueDate) {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitInformation) GetDueDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *DirectDebitInformation) HasDueDate() bool {
	if o != nil && !common.IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *DirectDebitInformation) SetDueDate(v time.Time) {
	o.DueDate = &v
}

// GetMandateId returns the MandateId field value if set, zero value otherwise.
func (o *DirectDebitInformation) GetMandateId() string {
	if o == nil || common.IsNil(o.MandateId) {
		var ret string
		return ret
	}
	return *o.MandateId
}

// GetMandateIdOk returns a tuple with the MandateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitInformation) GetMandateIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MandateId) {
		return nil, false
	}
	return o.MandateId, true
}

// HasMandateId returns a boolean if a field has been set.
func (o *DirectDebitInformation) HasMandateId() bool {
	if o != nil && !common.IsNil(o.MandateId) {
		return true
	}

	return false
}

// SetMandateId gets a reference to the given string and assigns it to the MandateId field.
func (o *DirectDebitInformation) SetMandateId(v string) {
	o.MandateId = &v
}

// GetSequenceType returns the SequenceType field value if set, zero value otherwise.
func (o *DirectDebitInformation) GetSequenceType() string {
	if o == nil || common.IsNil(o.SequenceType) {
		var ret string
		return ret
	}
	return *o.SequenceType
}

// GetSequenceTypeOk returns a tuple with the SequenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitInformation) GetSequenceTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SequenceType) {
		return nil, false
	}
	return o.SequenceType, true
}

// HasSequenceType returns a boolean if a field has been set.
func (o *DirectDebitInformation) HasSequenceType() bool {
	if o != nil && !common.IsNil(o.SequenceType) {
		return true
	}

	return false
}

// SetSequenceType gets a reference to the given string and assigns it to the SequenceType field.
func (o *DirectDebitInformation) SetSequenceType(v string) {
	o.SequenceType = &v
}

func (o DirectDebitInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DateOfSignature) {
		toSerialize["dateOfSignature"] = o.DateOfSignature
	}
	if !common.IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !common.IsNil(o.MandateId) {
		toSerialize["mandateId"] = o.MandateId
	}
	if !common.IsNil(o.SequenceType) {
		toSerialize["sequenceType"] = o.SequenceType
	}
	return toSerialize, nil
}

type NullableDirectDebitInformation struct {
	value *DirectDebitInformation
	isSet bool
}

func (v NullableDirectDebitInformation) Get() *DirectDebitInformation {
	return v.value
}

func (v *NullableDirectDebitInformation) Set(val *DirectDebitInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitInformation(val *DirectDebitInformation) *NullableDirectDebitInformation {
	return &NullableDirectDebitInformation{value: val, isSet: true}
}

func (v NullableDirectDebitInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
