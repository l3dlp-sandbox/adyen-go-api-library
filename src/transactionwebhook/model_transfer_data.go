/*
Transaction webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactionwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the TransferData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferData{}

// TransferData struct for TransferData
type TransferData struct {
	// The ID of the resource.
	Id *string `json:"id,omitempty"`
	// The [`reference`](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__reqParam_reference) from the `/transfers` request. If you haven't provided any, Adyen generates a unique reference.
	Reference string `json:"reference"`
}

// NewTransferData instantiates a new TransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferData(reference string) *TransferData {
	this := TransferData{}
	this.Reference = reference
	return &this
}

// NewTransferDataWithDefaults instantiates a new TransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDataWithDefaults() *TransferData {
	this := TransferData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransferData) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransferData) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransferData) SetId(v string) {
	o.Id = &v
}

// GetReference returns the Reference field value
func (o *TransferData) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *TransferData) SetReference(v string) {
	o.Reference = v
}

func (o TransferData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["reference"] = o.Reference
	return toSerialize, nil
}

type NullableTransferData struct {
	value *TransferData
	isSet bool
}

func (v NullableTransferData) Get() *TransferData {
	return v.value
}

func (v *NullableTransferData) Set(val *TransferData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferData(val *TransferData) *NullableTransferData {
	return &NullableTransferData{value: val, isSet: true}
}

func (v NullableTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
