/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the SourceOfFunds type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SourceOfFunds{}

// SourceOfFunds struct for SourceOfFunds
type SourceOfFunds struct {
	// The unique identifier of the business line that will be the source of funds.This must be a business line for a **receivePayments** or **receiveFromPlatformPayments** capability.
	// Deprecated since Legal Entity Management API v3
	// This field will be removed in v4.
	AcquiringBusinessLineId *string `json:"acquiringBusinessLineId,omitempty"`
	// Indicates whether the funds are coming from transactions processed by Adyen. If **false**, a `description` is required.
	AdyenProcessedFunds *bool `json:"adyenProcessedFunds,omitempty"`
	// Text describing the source of funds. For example, for `type` **business**, provide a description of where the business transactions come from, such as payments through bank transfer. Required when `adyenProcessedFunds` is **false**.
	Description *string `json:"description,omitempty"`
	// The type of the source of funds. Possible value: **business**.
	Type *string `json:"type,omitempty"`
}

// NewSourceOfFunds instantiates a new SourceOfFunds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceOfFunds() *SourceOfFunds {
	this := SourceOfFunds{}
	return &this
}

// NewSourceOfFundsWithDefaults instantiates a new SourceOfFunds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceOfFundsWithDefaults() *SourceOfFunds {
	this := SourceOfFunds{}
	return &this
}

// GetAcquiringBusinessLineId returns the AcquiringBusinessLineId field value if set, zero value otherwise.
// Deprecated since Legal Entity Management API v3
// This field will be removed in v4.
func (o *SourceOfFunds) GetAcquiringBusinessLineId() string {
	if o == nil || common.IsNil(o.AcquiringBusinessLineId) {
		var ret string
		return ret
	}
	return *o.AcquiringBusinessLineId
}

// GetAcquiringBusinessLineIdOk returns a tuple with the AcquiringBusinessLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Legal Entity Management API v3
// This field will be removed in v4.
func (o *SourceOfFunds) GetAcquiringBusinessLineIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcquiringBusinessLineId) {
		return nil, false
	}
	return o.AcquiringBusinessLineId, true
}

// HasAcquiringBusinessLineId returns a boolean if a field has been set.
func (o *SourceOfFunds) HasAcquiringBusinessLineId() bool {
	if o != nil && !common.IsNil(o.AcquiringBusinessLineId) {
		return true
	}

	return false
}

// SetAcquiringBusinessLineId gets a reference to the given string and assigns it to the AcquiringBusinessLineId field.
// Deprecated since Legal Entity Management API v3
// This field will be removed in v4.
func (o *SourceOfFunds) SetAcquiringBusinessLineId(v string) {
	o.AcquiringBusinessLineId = &v
}

// GetAdyenProcessedFunds returns the AdyenProcessedFunds field value if set, zero value otherwise.
func (o *SourceOfFunds) GetAdyenProcessedFunds() bool {
	if o == nil || common.IsNil(o.AdyenProcessedFunds) {
		var ret bool
		return ret
	}
	return *o.AdyenProcessedFunds
}

// GetAdyenProcessedFundsOk returns a tuple with the AdyenProcessedFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOfFunds) GetAdyenProcessedFundsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AdyenProcessedFunds) {
		return nil, false
	}
	return o.AdyenProcessedFunds, true
}

// HasAdyenProcessedFunds returns a boolean if a field has been set.
func (o *SourceOfFunds) HasAdyenProcessedFunds() bool {
	if o != nil && !common.IsNil(o.AdyenProcessedFunds) {
		return true
	}

	return false
}

// SetAdyenProcessedFunds gets a reference to the given bool and assigns it to the AdyenProcessedFunds field.
func (o *SourceOfFunds) SetAdyenProcessedFunds(v bool) {
	o.AdyenProcessedFunds = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SourceOfFunds) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOfFunds) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SourceOfFunds) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SourceOfFunds) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceOfFunds) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOfFunds) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceOfFunds) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceOfFunds) SetType(v string) {
	o.Type = &v
}

func (o SourceOfFunds) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceOfFunds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcquiringBusinessLineId) {
		toSerialize["acquiringBusinessLineId"] = o.AcquiringBusinessLineId
	}
	if !common.IsNil(o.AdyenProcessedFunds) {
		toSerialize["adyenProcessedFunds"] = o.AdyenProcessedFunds
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSourceOfFunds struct {
	value *SourceOfFunds
	isSet bool
}

func (v NullableSourceOfFunds) Get() *SourceOfFunds {
	return v.value
}

func (v *NullableSourceOfFunds) Set(val *SourceOfFunds) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceOfFunds) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceOfFunds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceOfFunds(val *SourceOfFunds) *NullableSourceOfFunds {
	return &NullableSourceOfFunds{value: val, isSet: true}
}

func (v NullableSourceOfFunds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceOfFunds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SourceOfFunds) isValidType() bool {
	var allowedEnumValues = []string{"business"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
