/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the ReturnTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReturnTransferResponse{}

// ReturnTransferResponse struct for ReturnTransferResponse
type ReturnTransferResponse struct {
	// The unique identifier of the return.
	Id *string `json:"id,omitempty"`
	// Your internal reference for the return.
	Reference *string `json:"reference,omitempty"`
	// The resulting status of the return.  Possible values: **Authorised**, **Declined**.
	Status *string `json:"status,omitempty"`
	// The unique identifier of the original transfer.
	TransferId *string `json:"transferId,omitempty"`
}

// NewReturnTransferResponse instantiates a new ReturnTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnTransferResponse() *ReturnTransferResponse {
	this := ReturnTransferResponse{}
	return &this
}

// NewReturnTransferResponseWithDefaults instantiates a new ReturnTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnTransferResponseWithDefaults() *ReturnTransferResponse {
	this := ReturnTransferResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReturnTransferResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnTransferResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReturnTransferResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReturnTransferResponse) SetId(v string) {
	o.Id = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ReturnTransferResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnTransferResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ReturnTransferResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ReturnTransferResponse) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReturnTransferResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnTransferResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReturnTransferResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReturnTransferResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *ReturnTransferResponse) GetTransferId() string {
	if o == nil || common.IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnTransferResponse) GetTransferIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *ReturnTransferResponse) HasTransferId() bool {
	if o != nil && !common.IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *ReturnTransferResponse) SetTransferId(v string) {
	o.TransferId = &v
}

func (o ReturnTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TransferId) {
		toSerialize["transferId"] = o.TransferId
	}
	return toSerialize, nil
}

type NullableReturnTransferResponse struct {
	value *ReturnTransferResponse
	isSet bool
}

func (v NullableReturnTransferResponse) Get() *ReturnTransferResponse {
	return v.value
}

func (v *NullableReturnTransferResponse) Set(val *ReturnTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnTransferResponse(val *ReturnTransferResponse) *NullableReturnTransferResponse {
	return &NullableReturnTransferResponse{value: val, isSet: true}
}

func (v NullableReturnTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ReturnTransferResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"Authorised", "Declined"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
