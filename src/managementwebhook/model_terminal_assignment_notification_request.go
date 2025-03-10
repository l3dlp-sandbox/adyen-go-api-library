/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the TerminalAssignmentNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalAssignmentNotificationRequest{}

// TerminalAssignmentNotificationRequest struct for TerminalAssignmentNotificationRequest
type TerminalAssignmentNotificationRequest struct {
	// The unique identifier of the merchant/company account to which the terminal is assigned.
	AssignedToAccount string `json:"assignedToAccount"`
	// The unique identifier of the store to which the terminal is assigned.
	AssignedToStore *string `json:"assignedToStore,omitempty"`
	// The date and time when an event has been completed.
	EventDate string `json:"eventDate"`
	// The PSP reference of the request from which the notification originates.
	PspReference string `json:"pspReference"`
	// The unique identifier of the terminal.
	UniqueTerminalId string `json:"uniqueTerminalId"`
}

// NewTerminalAssignmentNotificationRequest instantiates a new TerminalAssignmentNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalAssignmentNotificationRequest(assignedToAccount string, eventDate string, pspReference string, uniqueTerminalId string) *TerminalAssignmentNotificationRequest {
	this := TerminalAssignmentNotificationRequest{}
	this.AssignedToAccount = assignedToAccount
	this.EventDate = eventDate
	this.PspReference = pspReference
	this.UniqueTerminalId = uniqueTerminalId
	return &this
}

// NewTerminalAssignmentNotificationRequestWithDefaults instantiates a new TerminalAssignmentNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalAssignmentNotificationRequestWithDefaults() *TerminalAssignmentNotificationRequest {
	this := TerminalAssignmentNotificationRequest{}
	return &this
}

// GetAssignedToAccount returns the AssignedToAccount field value
func (o *TerminalAssignmentNotificationRequest) GetAssignedToAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedToAccount
}

// GetAssignedToAccountOk returns a tuple with the AssignedToAccount field value
// and a boolean to check if the value has been set.
func (o *TerminalAssignmentNotificationRequest) GetAssignedToAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedToAccount, true
}

// SetAssignedToAccount sets field value
func (o *TerminalAssignmentNotificationRequest) SetAssignedToAccount(v string) {
	o.AssignedToAccount = v
}

// GetAssignedToStore returns the AssignedToStore field value if set, zero value otherwise.
func (o *TerminalAssignmentNotificationRequest) GetAssignedToStore() string {
	if o == nil || common.IsNil(o.AssignedToStore) {
		var ret string
		return ret
	}
	return *o.AssignedToStore
}

// GetAssignedToStoreOk returns a tuple with the AssignedToStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAssignmentNotificationRequest) GetAssignedToStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssignedToStore) {
		return nil, false
	}
	return o.AssignedToStore, true
}

// HasAssignedToStore returns a boolean if a field has been set.
func (o *TerminalAssignmentNotificationRequest) HasAssignedToStore() bool {
	if o != nil && !common.IsNil(o.AssignedToStore) {
		return true
	}

	return false
}

// SetAssignedToStore gets a reference to the given string and assigns it to the AssignedToStore field.
func (o *TerminalAssignmentNotificationRequest) SetAssignedToStore(v string) {
	o.AssignedToStore = &v
}

// GetEventDate returns the EventDate field value
func (o *TerminalAssignmentNotificationRequest) GetEventDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value
// and a boolean to check if the value has been set.
func (o *TerminalAssignmentNotificationRequest) GetEventDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDate, true
}

// SetEventDate sets field value
func (o *TerminalAssignmentNotificationRequest) SetEventDate(v string) {
	o.EventDate = v
}

// GetPspReference returns the PspReference field value
func (o *TerminalAssignmentNotificationRequest) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *TerminalAssignmentNotificationRequest) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *TerminalAssignmentNotificationRequest) SetPspReference(v string) {
	o.PspReference = v
}

// GetUniqueTerminalId returns the UniqueTerminalId field value
func (o *TerminalAssignmentNotificationRequest) GetUniqueTerminalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueTerminalId
}

// GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field value
// and a boolean to check if the value has been set.
func (o *TerminalAssignmentNotificationRequest) GetUniqueTerminalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueTerminalId, true
}

// SetUniqueTerminalId sets field value
func (o *TerminalAssignmentNotificationRequest) SetUniqueTerminalId(v string) {
	o.UniqueTerminalId = v
}

func (o TerminalAssignmentNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalAssignmentNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignedToAccount"] = o.AssignedToAccount
	if !common.IsNil(o.AssignedToStore) {
		toSerialize["assignedToStore"] = o.AssignedToStore
	}
	toSerialize["eventDate"] = o.EventDate
	toSerialize["pspReference"] = o.PspReference
	toSerialize["uniqueTerminalId"] = o.UniqueTerminalId
	return toSerialize, nil
}

type NullableTerminalAssignmentNotificationRequest struct {
	value *TerminalAssignmentNotificationRequest
	isSet bool
}

func (v NullableTerminalAssignmentNotificationRequest) Get() *TerminalAssignmentNotificationRequest {
	return v.value
}

func (v *NullableTerminalAssignmentNotificationRequest) Set(val *TerminalAssignmentNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalAssignmentNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalAssignmentNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalAssignmentNotificationRequest(val *TerminalAssignmentNotificationRequest) *NullableTerminalAssignmentNotificationRequest {
	return &NullableTerminalAssignmentNotificationRequest{value: val, isSet: true}
}

func (v NullableTerminalAssignmentNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalAssignmentNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
