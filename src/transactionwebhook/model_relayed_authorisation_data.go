/*
Transaction webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactionwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the RelayedAuthorisationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RelayedAuthorisationData{}

// RelayedAuthorisationData struct for RelayedAuthorisationData
type RelayedAuthorisationData struct {
	// Contains key-value pairs of your references and descriptions, for example, `customId`:`your-own-custom-field-12345`.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Your reference for the relayed authorisation data.
	Reference *string `json:"reference,omitempty"`
}

// NewRelayedAuthorisationData instantiates a new RelayedAuthorisationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelayedAuthorisationData() *RelayedAuthorisationData {
	this := RelayedAuthorisationData{}
	return &this
}

// NewRelayedAuthorisationDataWithDefaults instantiates a new RelayedAuthorisationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelayedAuthorisationDataWithDefaults() *RelayedAuthorisationData {
	this := RelayedAuthorisationData{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RelayedAuthorisationData) GetMetadata() map[string]string {
	if o == nil || common.IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayedAuthorisationData) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RelayedAuthorisationData) HasMetadata() bool {
	if o != nil && !common.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *RelayedAuthorisationData) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *RelayedAuthorisationData) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayedAuthorisationData) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *RelayedAuthorisationData) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *RelayedAuthorisationData) SetReference(v string) {
	o.Reference = &v
}

func (o RelayedAuthorisationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelayedAuthorisationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableRelayedAuthorisationData struct {
	value *RelayedAuthorisationData
	isSet bool
}

func (v NullableRelayedAuthorisationData) Get() *RelayedAuthorisationData {
	return v.value
}

func (v *NullableRelayedAuthorisationData) Set(val *RelayedAuthorisationData) {
	v.value = val
	v.isSet = true
}

func (v NullableRelayedAuthorisationData) IsSet() bool {
	return v.isSet
}

func (v *NullableRelayedAuthorisationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelayedAuthorisationData(val *RelayedAuthorisationData) *NullableRelayedAuthorisationData {
	return &NullableRelayedAuthorisationData{value: val, isSet: true}
}

func (v NullableRelayedAuthorisationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelayedAuthorisationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



