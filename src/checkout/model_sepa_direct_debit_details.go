/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the SepaDirectDebitDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SepaDirectDebitDetails{}

// SepaDirectDebitDetails struct for SepaDirectDebitDetails
type SepaDirectDebitDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The International Bank Account Number (IBAN).
	Iban string `json:"iban"`
	// The name of the bank account holder.
	OwnerName string `json:"ownerName"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// The unique identifier of your user's verified transfer instrument, which you can use to top up their balance accounts.
	TransferInstrumentId *string `json:"transferInstrumentId,omitempty"`
	// **sepadirectdebit**
	Type *string `json:"type,omitempty"`
}

// NewSepaDirectDebitDetails instantiates a new SepaDirectDebitDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSepaDirectDebitDetails(iban string, ownerName string) *SepaDirectDebitDetails {
	this := SepaDirectDebitDetails{}
	this.Iban = iban
	this.OwnerName = ownerName
	var type_ string = "sepadirectdebit"
	this.Type = &type_
	return &this
}

// NewSepaDirectDebitDetailsWithDefaults instantiates a new SepaDirectDebitDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSepaDirectDebitDetailsWithDefaults() *SepaDirectDebitDetails {
	this := SepaDirectDebitDetails{}
	var type_ string = "sepadirectdebit"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *SepaDirectDebitDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SepaDirectDebitDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *SepaDirectDebitDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *SepaDirectDebitDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetIban returns the Iban field value
func (o *SepaDirectDebitDetails) GetIban() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iban
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
func (o *SepaDirectDebitDetails) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iban, true
}

// SetIban sets field value
func (o *SepaDirectDebitDetails) SetIban(v string) {
	o.Iban = v
}

// GetOwnerName returns the OwnerName field value
func (o *SepaDirectDebitDetails) GetOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value
// and a boolean to check if the value has been set.
func (o *SepaDirectDebitDetails) GetOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerName, true
}

// SetOwnerName sets field value
func (o *SepaDirectDebitDetails) SetOwnerName(v string) {
	o.OwnerName = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *SepaDirectDebitDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *SepaDirectDebitDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *SepaDirectDebitDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *SepaDirectDebitDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *SepaDirectDebitDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SepaDirectDebitDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *SepaDirectDebitDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *SepaDirectDebitDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value if set, zero value otherwise.
func (o *SepaDirectDebitDetails) GetTransferInstrumentId() string {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		var ret string
		return ret
	}
	return *o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SepaDirectDebitDetails) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		return nil, false
	}
	return o.TransferInstrumentId, true
}

// HasTransferInstrumentId returns a boolean if a field has been set.
func (o *SepaDirectDebitDetails) HasTransferInstrumentId() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentId) {
		return true
	}

	return false
}

// SetTransferInstrumentId gets a reference to the given string and assigns it to the TransferInstrumentId field.
func (o *SepaDirectDebitDetails) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SepaDirectDebitDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SepaDirectDebitDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SepaDirectDebitDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SepaDirectDebitDetails) SetType(v string) {
	o.Type = &v
}

func (o SepaDirectDebitDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SepaDirectDebitDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	toSerialize["iban"] = o.Iban
	toSerialize["ownerName"] = o.OwnerName
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.TransferInstrumentId) {
		toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSepaDirectDebitDetails struct {
	value *SepaDirectDebitDetails
	isSet bool
}

func (v NullableSepaDirectDebitDetails) Get() *SepaDirectDebitDetails {
	return v.value
}

func (v *NullableSepaDirectDebitDetails) Set(val *SepaDirectDebitDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSepaDirectDebitDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSepaDirectDebitDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSepaDirectDebitDetails(val *SepaDirectDebitDetails) *NullableSepaDirectDebitDetails {
	return &NullableSepaDirectDebitDetails{value: val, isSet: true}
}

func (v NullableSepaDirectDebitDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSepaDirectDebitDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SepaDirectDebitDetails) isValidType() bool {
	var allowedEnumValues = []string{"sepadirectdebit", "sepadirectdebit_amazonpay"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
