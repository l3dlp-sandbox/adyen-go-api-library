/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the TransactionRuleReference type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRuleReference{}

// TransactionRuleReference struct for TransactionRuleReference
type TransactionRuleReference struct {
	// The description of the resource.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the resource.
	Id *string `json:"id,omitempty"`
	// The outcome type of the rule.
	OutcomeType *string `json:"outcomeType,omitempty"`
	// The reference for the resource.
	Reference *string `json:"reference,omitempty"`
	// The score of the rule in case it's a scoreBased rule.
	Score *int32 `json:"score,omitempty"`
}

// NewTransactionRuleReference instantiates a new TransactionRuleReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleReference() *TransactionRuleReference {
	this := TransactionRuleReference{}
	return &this
}

// NewTransactionRuleReferenceWithDefaults instantiates a new TransactionRuleReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleReferenceWithDefaults() *TransactionRuleReference {
	this := TransactionRuleReference{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionRuleReference) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleReference) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionRuleReference) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionRuleReference) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionRuleReference) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleReference) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionRuleReference) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionRuleReference) SetId(v string) {
	o.Id = &v
}

// GetOutcomeType returns the OutcomeType field value if set, zero value otherwise.
func (o *TransactionRuleReference) GetOutcomeType() string {
	if o == nil || common.IsNil(o.OutcomeType) {
		var ret string
		return ret
	}
	return *o.OutcomeType
}

// GetOutcomeTypeOk returns a tuple with the OutcomeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleReference) GetOutcomeTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OutcomeType) {
		return nil, false
	}
	return o.OutcomeType, true
}

// HasOutcomeType returns a boolean if a field has been set.
func (o *TransactionRuleReference) HasOutcomeType() bool {
	if o != nil && !common.IsNil(o.OutcomeType) {
		return true
	}

	return false
}

// SetOutcomeType gets a reference to the given string and assigns it to the OutcomeType field.
func (o *TransactionRuleReference) SetOutcomeType(v string) {
	o.OutcomeType = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TransactionRuleReference) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleReference) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TransactionRuleReference) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TransactionRuleReference) SetReference(v string) {
	o.Reference = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *TransactionRuleReference) GetScore() int32 {
	if o == nil || common.IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleReference) GetScoreOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *TransactionRuleReference) HasScore() bool {
	if o != nil && !common.IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *TransactionRuleReference) SetScore(v int32) {
	o.Score = &v
}

func (o TransactionRuleReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.OutcomeType) {
		toSerialize["outcomeType"] = o.OutcomeType
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return toSerialize, nil
}

type NullableTransactionRuleReference struct {
	value *TransactionRuleReference
	isSet bool
}

func (v NullableTransactionRuleReference) Get() *TransactionRuleReference {
	return v.value
}

func (v *NullableTransactionRuleReference) Set(val *TransactionRuleReference) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleReference) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleReference(val *TransactionRuleReference) *NullableTransactionRuleReference {
	return &NullableTransactionRuleReference{value: val, isSet: true}
}

func (v NullableTransactionRuleReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
