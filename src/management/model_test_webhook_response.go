/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TestWebhookResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TestWebhookResponse{}

// TestWebhookResponse struct for TestWebhookResponse
type TestWebhookResponse struct {
	// List with test results. Each test webhook we send has a list element with the result.
	Data []TestOutput `json:"data,omitempty"`
}

// NewTestWebhookResponse instantiates a new TestWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestWebhookResponse() *TestWebhookResponse {
	this := TestWebhookResponse{}
	return &this
}

// NewTestWebhookResponseWithDefaults instantiates a new TestWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestWebhookResponseWithDefaults() *TestWebhookResponse {
	this := TestWebhookResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TestWebhookResponse) GetData() []TestOutput {
	if o == nil || common.IsNil(o.Data) {
		var ret []TestOutput
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestWebhookResponse) GetDataOk() ([]TestOutput, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TestWebhookResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TestOutput and assigns it to the Data field.
func (o *TestWebhookResponse) SetData(v []TestOutput) {
	o.Data = v
}

func (o TestWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTestWebhookResponse struct {
	value *TestWebhookResponse
	isSet bool
}

func (v NullableTestWebhookResponse) Get() *TestWebhookResponse {
	return v.value
}

func (v *NullableTestWebhookResponse) Set(val *TestWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTestWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTestWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestWebhookResponse(val *TestWebhookResponse) *NullableTestWebhookResponse {
	return &NullableTestWebhookResponse{value: val, isSet: true}
}

func (v NullableTestWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



