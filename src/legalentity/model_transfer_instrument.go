/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the TransferInstrument type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferInstrument{}

// TransferInstrument struct for TransferInstrument
type TransferInstrument struct {
	BankAccount BankAccountInfo `json:"bankAccount"`
	// List of capabilities for this transfer instrument.
	Capabilities *map[string]SupportingEntityCapability `json:"capabilities,omitempty"`
	// List of documents uploaded for the transfer instrument.
	DocumentDetails []DocumentReference `json:"documentDetails,omitempty"`
	// The unique identifier of the transfer instrument.
	Id string `json:"id"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) that owns the transfer instrument.
	LegalEntityId string `json:"legalEntityId"`
	// The verification errors related to capabilities for this transfer instrument.
	Problems []CapabilityProblem `json:"problems,omitempty"`
	// The type of transfer instrument.  Possible value: **bankAccount**.
	Type string `json:"type"`
}

// NewTransferInstrument instantiates a new TransferInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferInstrument(bankAccount BankAccountInfo, id string, legalEntityId string, type_ string) *TransferInstrument {
	this := TransferInstrument{}
	this.BankAccount = bankAccount
	this.Id = id
	this.LegalEntityId = legalEntityId
	this.Type = type_
	return &this
}

// NewTransferInstrumentWithDefaults instantiates a new TransferInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferInstrumentWithDefaults() *TransferInstrument {
	this := TransferInstrument{}
	return &this
}

// GetBankAccount returns the BankAccount field value
func (o *TransferInstrument) GetBankAccount() BankAccountInfo {
	if o == nil {
		var ret BankAccountInfo
		return ret
	}

	return o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
func (o *TransferInstrument) GetBankAccountOk() (*BankAccountInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccount, true
}

// SetBankAccount sets field value
func (o *TransferInstrument) SetBankAccount(v BankAccountInfo) {
	o.BankAccount = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *TransferInstrument) GetCapabilities() map[string]SupportingEntityCapability {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret map[string]SupportingEntityCapability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInstrument) GetCapabilitiesOk() (*map[string]SupportingEntityCapability, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *TransferInstrument) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]SupportingEntityCapability and assigns it to the Capabilities field.
func (o *TransferInstrument) SetCapabilities(v map[string]SupportingEntityCapability) {
	o.Capabilities = &v
}

// GetDocumentDetails returns the DocumentDetails field value if set, zero value otherwise.
func (o *TransferInstrument) GetDocumentDetails() []DocumentReference {
	if o == nil || common.IsNil(o.DocumentDetails) {
		var ret []DocumentReference
		return ret
	}
	return o.DocumentDetails
}

// GetDocumentDetailsOk returns a tuple with the DocumentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInstrument) GetDocumentDetailsOk() ([]DocumentReference, bool) {
	if o == nil || common.IsNil(o.DocumentDetails) {
		return nil, false
	}
	return o.DocumentDetails, true
}

// HasDocumentDetails returns a boolean if a field has been set.
func (o *TransferInstrument) HasDocumentDetails() bool {
	if o != nil && !common.IsNil(o.DocumentDetails) {
		return true
	}

	return false
}

// SetDocumentDetails gets a reference to the given []DocumentReference and assigns it to the DocumentDetails field.
func (o *TransferInstrument) SetDocumentDetails(v []DocumentReference) {
	o.DocumentDetails = v
}

// GetId returns the Id field value
func (o *TransferInstrument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferInstrument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferInstrument) SetId(v string) {
	o.Id = v
}

// GetLegalEntityId returns the LegalEntityId field value
func (o *TransferInstrument) GetLegalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value
// and a boolean to check if the value has been set.
func (o *TransferInstrument) GetLegalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalEntityId, true
}

// SetLegalEntityId sets field value
func (o *TransferInstrument) SetLegalEntityId(v string) {
	o.LegalEntityId = v
}

// GetProblems returns the Problems field value if set, zero value otherwise.
func (o *TransferInstrument) GetProblems() []CapabilityProblem {
	if o == nil || common.IsNil(o.Problems) {
		var ret []CapabilityProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInstrument) GetProblemsOk() ([]CapabilityProblem, bool) {
	if o == nil || common.IsNil(o.Problems) {
		return nil, false
	}
	return o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *TransferInstrument) HasProblems() bool {
	if o != nil && !common.IsNil(o.Problems) {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []CapabilityProblem and assigns it to the Problems field.
func (o *TransferInstrument) SetProblems(v []CapabilityProblem) {
	o.Problems = v
}

// GetType returns the Type field value
func (o *TransferInstrument) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferInstrument) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferInstrument) SetType(v string) {
	o.Type = v
}

func (o TransferInstrument) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferInstrument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bankAccount"] = o.BankAccount
	if !common.IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !common.IsNil(o.DocumentDetails) {
		toSerialize["documentDetails"] = o.DocumentDetails
	}
	toSerialize["id"] = o.Id
	toSerialize["legalEntityId"] = o.LegalEntityId
	if !common.IsNil(o.Problems) {
		toSerialize["problems"] = o.Problems
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransferInstrument struct {
	value *TransferInstrument
	isSet bool
}

func (v NullableTransferInstrument) Get() *TransferInstrument {
	return v.value
}

func (v *NullableTransferInstrument) Set(val *TransferInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferInstrument(val *TransferInstrument) *NullableTransferInstrument {
	return &NullableTransferInstrument{value: val, isSet: true}
}

func (v NullableTransferInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferInstrument) isValidType() bool {
	var allowedEnumValues = []string{"bankAccount", "recurringDetail"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
