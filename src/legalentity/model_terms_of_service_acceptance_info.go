/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the TermsOfServiceAcceptanceInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TermsOfServiceAcceptanceInfo{}

// TermsOfServiceAcceptanceInfo struct for TermsOfServiceAcceptanceInfo
type TermsOfServiceAcceptanceInfo struct {
	// The unique identifier of the user that accepted the Terms of Service.
	AcceptedBy *string `json:"acceptedBy,omitempty"`
	// The unique identifier of the legal entity for which the Terms of Service are accepted.
	AcceptedFor *string `json:"acceptedFor,omitempty"`
	// The date when the Terms of Service were accepted.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// An Adyen-generated reference for the accepted Terms of Service.
	Id *string `json:"id,omitempty"`
	// The type of Terms of Service.
	Type *string `json:"type,omitempty"`
}

// NewTermsOfServiceAcceptanceInfo instantiates a new TermsOfServiceAcceptanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTermsOfServiceAcceptanceInfo() *TermsOfServiceAcceptanceInfo {
	this := TermsOfServiceAcceptanceInfo{}
	return &this
}

// NewTermsOfServiceAcceptanceInfoWithDefaults instantiates a new TermsOfServiceAcceptanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermsOfServiceAcceptanceInfoWithDefaults() *TermsOfServiceAcceptanceInfo {
	this := TermsOfServiceAcceptanceInfo{}
	return &this
}

// GetAcceptedBy returns the AcceptedBy field value if set, zero value otherwise.
func (o *TermsOfServiceAcceptanceInfo) GetAcceptedBy() string {
	if o == nil || common.IsNil(o.AcceptedBy) {
		var ret string
		return ret
	}
	return *o.AcceptedBy
}

// GetAcceptedByOk returns a tuple with the AcceptedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsOfServiceAcceptanceInfo) GetAcceptedByOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcceptedBy) {
		return nil, false
	}
	return o.AcceptedBy, true
}

// HasAcceptedBy returns a boolean if a field has been set.
func (o *TermsOfServiceAcceptanceInfo) HasAcceptedBy() bool {
	if o != nil && !common.IsNil(o.AcceptedBy) {
		return true
	}

	return false
}

// SetAcceptedBy gets a reference to the given string and assigns it to the AcceptedBy field.
func (o *TermsOfServiceAcceptanceInfo) SetAcceptedBy(v string) {
	o.AcceptedBy = &v
}

// GetAcceptedFor returns the AcceptedFor field value if set, zero value otherwise.
func (o *TermsOfServiceAcceptanceInfo) GetAcceptedFor() string {
	if o == nil || common.IsNil(o.AcceptedFor) {
		var ret string
		return ret
	}
	return *o.AcceptedFor
}

// GetAcceptedForOk returns a tuple with the AcceptedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsOfServiceAcceptanceInfo) GetAcceptedForOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcceptedFor) {
		return nil, false
	}
	return o.AcceptedFor, true
}

// HasAcceptedFor returns a boolean if a field has been set.
func (o *TermsOfServiceAcceptanceInfo) HasAcceptedFor() bool {
	if o != nil && !common.IsNil(o.AcceptedFor) {
		return true
	}

	return false
}

// SetAcceptedFor gets a reference to the given string and assigns it to the AcceptedFor field.
func (o *TermsOfServiceAcceptanceInfo) SetAcceptedFor(v string) {
	o.AcceptedFor = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TermsOfServiceAcceptanceInfo) GetCreatedAt() time.Time {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsOfServiceAcceptanceInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TermsOfServiceAcceptanceInfo) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TermsOfServiceAcceptanceInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TermsOfServiceAcceptanceInfo) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsOfServiceAcceptanceInfo) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TermsOfServiceAcceptanceInfo) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TermsOfServiceAcceptanceInfo) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TermsOfServiceAcceptanceInfo) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsOfServiceAcceptanceInfo) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TermsOfServiceAcceptanceInfo) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TermsOfServiceAcceptanceInfo) SetType(v string) {
	o.Type = &v
}

func (o TermsOfServiceAcceptanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TermsOfServiceAcceptanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcceptedBy) {
		toSerialize["acceptedBy"] = o.AcceptedBy
	}
	if !common.IsNil(o.AcceptedFor) {
		toSerialize["acceptedFor"] = o.AcceptedFor
	}
	if !common.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTermsOfServiceAcceptanceInfo struct {
	value *TermsOfServiceAcceptanceInfo
	isSet bool
}

func (v NullableTermsOfServiceAcceptanceInfo) Get() *TermsOfServiceAcceptanceInfo {
	return v.value
}

func (v *NullableTermsOfServiceAcceptanceInfo) Set(val *TermsOfServiceAcceptanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTermsOfServiceAcceptanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTermsOfServiceAcceptanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermsOfServiceAcceptanceInfo(val *TermsOfServiceAcceptanceInfo) *NullableTermsOfServiceAcceptanceInfo {
	return &NullableTermsOfServiceAcceptanceInfo{value: val, isSet: true}
}

func (v NullableTermsOfServiceAcceptanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermsOfServiceAcceptanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TermsOfServiceAcceptanceInfo) isValidType() bool {
	var allowedEnumValues = []string{"adyenAccount", "adyenCapital", "adyenCard", "adyenForPlatformsAdvanced", "adyenForPlatformsManage", "adyenFranchisee", "adyenIssuing"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}