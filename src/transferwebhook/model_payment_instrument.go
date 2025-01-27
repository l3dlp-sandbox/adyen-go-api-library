/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the PaymentInstrument type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentInstrument{}

// PaymentInstrument struct for PaymentInstrument
type PaymentInstrument struct {
	// The description of the resource.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the resource.
	Id *string `json:"id,omitempty"`
	// The reference for the resource.
	Reference *string `json:"reference,omitempty"`
	// The type of wallet that the network token is associated with.
	TokenType *string `json:"tokenType,omitempty"`
}

// NewPaymentInstrument instantiates a new PaymentInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInstrument() *PaymentInstrument {
	this := PaymentInstrument{}
	return &this
}

// NewPaymentInstrumentWithDefaults instantiates a new PaymentInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInstrumentWithDefaults() *PaymentInstrument {
	this := PaymentInstrument{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentInstrument) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrument) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentInstrument) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentInstrument) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentInstrument) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrument) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentInstrument) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentInstrument) SetId(v string) {
	o.Id = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentInstrument) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrument) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentInstrument) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentInstrument) SetReference(v string) {
	o.Reference = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *PaymentInstrument) GetTokenType() string {
	if o == nil || common.IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrument) GetTokenTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *PaymentInstrument) HasTokenType() bool {
	if o != nil && !common.IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *PaymentInstrument) SetTokenType(v string) {
	o.TokenType = &v
}

func (o PaymentInstrument) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInstrument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.TokenType) {
		toSerialize["tokenType"] = o.TokenType
	}
	return toSerialize, nil
}

type NullablePaymentInstrument struct {
	value *PaymentInstrument
	isSet bool
}

func (v NullablePaymentInstrument) Get() *PaymentInstrument {
	return v.value
}

func (v *NullablePaymentInstrument) Set(val *PaymentInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrument(val *PaymentInstrument) *NullablePaymentInstrument {
	return &NullablePaymentInstrument{value: val, isSet: true}
}

func (v NullablePaymentInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
