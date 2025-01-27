/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the ListExternalTerminalActionsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListExternalTerminalActionsResponse{}

// ListExternalTerminalActionsResponse struct for ListExternalTerminalActionsResponse
type ListExternalTerminalActionsResponse struct {
	// The list of terminal actions.
	Data []ExternalTerminalAction `json:"data,omitempty"`
}

// NewListExternalTerminalActionsResponse instantiates a new ListExternalTerminalActionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListExternalTerminalActionsResponse() *ListExternalTerminalActionsResponse {
	this := ListExternalTerminalActionsResponse{}
	return &this
}

// NewListExternalTerminalActionsResponseWithDefaults instantiates a new ListExternalTerminalActionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListExternalTerminalActionsResponseWithDefaults() *ListExternalTerminalActionsResponse {
	this := ListExternalTerminalActionsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListExternalTerminalActionsResponse) GetData() []ExternalTerminalAction {
	if o == nil || common.IsNil(o.Data) {
		var ret []ExternalTerminalAction
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListExternalTerminalActionsResponse) GetDataOk() ([]ExternalTerminalAction, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListExternalTerminalActionsResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ExternalTerminalAction and assigns it to the Data field.
func (o *ListExternalTerminalActionsResponse) SetData(v []ExternalTerminalAction) {
	o.Data = v
}

func (o ListExternalTerminalActionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListExternalTerminalActionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListExternalTerminalActionsResponse struct {
	value *ListExternalTerminalActionsResponse
	isSet bool
}

func (v NullableListExternalTerminalActionsResponse) Get() *ListExternalTerminalActionsResponse {
	return v.value
}

func (v *NullableListExternalTerminalActionsResponse) Set(val *ListExternalTerminalActionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListExternalTerminalActionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListExternalTerminalActionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListExternalTerminalActionsResponse(val *ListExternalTerminalActionsResponse) *NullableListExternalTerminalActionsResponse {
	return &NullableListExternalTerminalActionsResponse{value: val, isSet: true}
}

func (v NullableListExternalTerminalActionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListExternalTerminalActionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
