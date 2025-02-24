/*
Adyen Stored Value API

API version: 46
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storedvalue

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the StoredValueVoidRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredValueVoidRequest{}

// StoredValueVoidRequest struct for StoredValueVoidRequest
type StoredValueVoidRequest struct {
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The original pspReference of the payment to modify.
	OriginalReference string `json:"originalReference"`
	// Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// The physical store, for which this payment is processed.
	Store *string `json:"store,omitempty"`
	// The reference of the tender.
	TenderReference *string `json:"tenderReference,omitempty"`
	// The unique ID of a POS terminal.
	UniqueTerminalId *string `json:"uniqueTerminalId,omitempty"`
}

// NewStoredValueVoidRequest instantiates a new StoredValueVoidRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredValueVoidRequest(merchantAccount string, originalReference string) *StoredValueVoidRequest {
	this := StoredValueVoidRequest{}
	this.MerchantAccount = merchantAccount
	this.OriginalReference = originalReference
	return &this
}

// NewStoredValueVoidRequestWithDefaults instantiates a new StoredValueVoidRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredValueVoidRequestWithDefaults() *StoredValueVoidRequest {
	this := StoredValueVoidRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StoredValueVoidRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StoredValueVoidRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StoredValueVoidRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetOriginalReference returns the OriginalReference field value
func (o *StoredValueVoidRequest) GetOriginalReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalReference
}

// GetOriginalReferenceOk returns a tuple with the OriginalReference field value
// and a boolean to check if the value has been set.
func (o *StoredValueVoidRequest) GetOriginalReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalReference, true
}

// SetOriginalReference sets field value
func (o *StoredValueVoidRequest) SetOriginalReference(v string) {
	o.OriginalReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *StoredValueVoidRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueVoidRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *StoredValueVoidRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *StoredValueVoidRequest) SetReference(v string) {
	o.Reference = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *StoredValueVoidRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueVoidRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *StoredValueVoidRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *StoredValueVoidRequest) SetStore(v string) {
	o.Store = &v
}

// GetTenderReference returns the TenderReference field value if set, zero value otherwise.
func (o *StoredValueVoidRequest) GetTenderReference() string {
	if o == nil || common.IsNil(o.TenderReference) {
		var ret string
		return ret
	}
	return *o.TenderReference
}

// GetTenderReferenceOk returns a tuple with the TenderReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueVoidRequest) GetTenderReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TenderReference) {
		return nil, false
	}
	return o.TenderReference, true
}

// HasTenderReference returns a boolean if a field has been set.
func (o *StoredValueVoidRequest) HasTenderReference() bool {
	if o != nil && !common.IsNil(o.TenderReference) {
		return true
	}

	return false
}

// SetTenderReference gets a reference to the given string and assigns it to the TenderReference field.
func (o *StoredValueVoidRequest) SetTenderReference(v string) {
	o.TenderReference = &v
}

// GetUniqueTerminalId returns the UniqueTerminalId field value if set, zero value otherwise.
func (o *StoredValueVoidRequest) GetUniqueTerminalId() string {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		var ret string
		return ret
	}
	return *o.UniqueTerminalId
}

// GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueVoidRequest) GetUniqueTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.UniqueTerminalId) {
		return nil, false
	}
	return o.UniqueTerminalId, true
}

// HasUniqueTerminalId returns a boolean if a field has been set.
func (o *StoredValueVoidRequest) HasUniqueTerminalId() bool {
	if o != nil && !common.IsNil(o.UniqueTerminalId) {
		return true
	}

	return false
}

// SetUniqueTerminalId gets a reference to the given string and assigns it to the UniqueTerminalId field.
func (o *StoredValueVoidRequest) SetUniqueTerminalId(v string) {
	o.UniqueTerminalId = &v
}

func (o StoredValueVoidRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredValueVoidRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["originalReference"] = o.OriginalReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	if !common.IsNil(o.TenderReference) {
		toSerialize["tenderReference"] = o.TenderReference
	}
	if !common.IsNil(o.UniqueTerminalId) {
		toSerialize["uniqueTerminalId"] = o.UniqueTerminalId
	}
	return toSerialize, nil
}

type NullableStoredValueVoidRequest struct {
	value *StoredValueVoidRequest
	isSet bool
}

func (v NullableStoredValueVoidRequest) Get() *StoredValueVoidRequest {
	return v.value
}

func (v *NullableStoredValueVoidRequest) Set(val *StoredValueVoidRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredValueVoidRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredValueVoidRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredValueVoidRequest(val *StoredValueVoidRequest) *NullableStoredValueVoidRequest {
	return &NullableStoredValueVoidRequest{value: val, isSet: true}
}

func (v NullableStoredValueVoidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredValueVoidRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
