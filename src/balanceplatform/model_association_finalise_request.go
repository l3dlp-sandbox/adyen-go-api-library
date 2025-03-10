/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the AssociationFinaliseRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssociationFinaliseRequest{}

// AssociationFinaliseRequest struct for AssociationFinaliseRequest
type AssociationFinaliseRequest struct {
	// The list of unique identifiers of the resources that you are associating with the SCA device.  Maximum: 5 strings.
	Ids                          []string                               `json:"ids"`
	StrongCustomerAuthentication AssociationDelegatedAuthenticationData `json:"strongCustomerAuthentication"`
	// The type of resource that you are associating with the SCA device.  Possible value: **PaymentInstrument**
	Type string `json:"type"`
}

// NewAssociationFinaliseRequest instantiates a new AssociationFinaliseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationFinaliseRequest(ids []string, strongCustomerAuthentication AssociationDelegatedAuthenticationData, type_ string) *AssociationFinaliseRequest {
	this := AssociationFinaliseRequest{}
	this.Ids = ids
	this.StrongCustomerAuthentication = strongCustomerAuthentication
	this.Type = type_
	return &this
}

// NewAssociationFinaliseRequestWithDefaults instantiates a new AssociationFinaliseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationFinaliseRequestWithDefaults() *AssociationFinaliseRequest {
	this := AssociationFinaliseRequest{}
	return &this
}

// GetIds returns the Ids field value
func (o *AssociationFinaliseRequest) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *AssociationFinaliseRequest) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *AssociationFinaliseRequest) SetIds(v []string) {
	o.Ids = v
}

// GetStrongCustomerAuthentication returns the StrongCustomerAuthentication field value
func (o *AssociationFinaliseRequest) GetStrongCustomerAuthentication() AssociationDelegatedAuthenticationData {
	if o == nil {
		var ret AssociationDelegatedAuthenticationData
		return ret
	}

	return o.StrongCustomerAuthentication
}

// GetStrongCustomerAuthenticationOk returns a tuple with the StrongCustomerAuthentication field value
// and a boolean to check if the value has been set.
func (o *AssociationFinaliseRequest) GetStrongCustomerAuthenticationOk() (*AssociationDelegatedAuthenticationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrongCustomerAuthentication, true
}

// SetStrongCustomerAuthentication sets field value
func (o *AssociationFinaliseRequest) SetStrongCustomerAuthentication(v AssociationDelegatedAuthenticationData) {
	o.StrongCustomerAuthentication = v
}

// GetType returns the Type field value
func (o *AssociationFinaliseRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssociationFinaliseRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssociationFinaliseRequest) SetType(v string) {
	o.Type = v
}

func (o AssociationFinaliseRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssociationFinaliseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	toSerialize["strongCustomerAuthentication"] = o.StrongCustomerAuthentication
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAssociationFinaliseRequest struct {
	value *AssociationFinaliseRequest
	isSet bool
}

func (v NullableAssociationFinaliseRequest) Get() *AssociationFinaliseRequest {
	return v.value
}

func (v *NullableAssociationFinaliseRequest) Set(val *AssociationFinaliseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationFinaliseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationFinaliseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationFinaliseRequest(val *AssociationFinaliseRequest) *NullableAssociationFinaliseRequest {
	return &NullableAssociationFinaliseRequest{value: val, isSet: true}
}

func (v NullableAssociationFinaliseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationFinaliseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AssociationFinaliseRequest) isValidType() bool {
	var allowedEnumValues = []string{"PaymentInstrument"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
