/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the AncvDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AncvDetails{}

// AncvDetails struct for AncvDetails
type AncvDetails struct {
	// ANCV account identification (email or account number)
	BeneficiaryId *string `json:"beneficiaryId,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **ancv**
	Type *string `json:"type,omitempty"`
}

// NewAncvDetails instantiates a new AncvDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAncvDetails() *AncvDetails {
	this := AncvDetails{}
	return &this
}

// NewAncvDetailsWithDefaults instantiates a new AncvDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAncvDetailsWithDefaults() *AncvDetails {
	this := AncvDetails{}
	return &this
}

// GetBeneficiaryId returns the BeneficiaryId field value if set, zero value otherwise.
func (o *AncvDetails) GetBeneficiaryId() string {
	if o == nil || common.IsNil(o.BeneficiaryId) {
		var ret string
		return ret
	}
	return *o.BeneficiaryId
}

// GetBeneficiaryIdOk returns a tuple with the BeneficiaryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AncvDetails) GetBeneficiaryIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BeneficiaryId) {
		return nil, false
	}
	return o.BeneficiaryId, true
}

// HasBeneficiaryId returns a boolean if a field has been set.
func (o *AncvDetails) HasBeneficiaryId() bool {
	if o != nil && !common.IsNil(o.BeneficiaryId) {
		return true
	}

	return false
}

// SetBeneficiaryId gets a reference to the given string and assigns it to the BeneficiaryId field.
func (o *AncvDetails) SetBeneficiaryId(v string) {
	o.BeneficiaryId = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *AncvDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AncvDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *AncvDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *AncvDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *AncvDetails) GetRecurringDetailReference() string {
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
func (o *AncvDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *AncvDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *AncvDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *AncvDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AncvDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *AncvDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *AncvDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AncvDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AncvDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AncvDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AncvDetails) SetType(v string) {
	o.Type = &v
}

func (o AncvDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AncvDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BeneficiaryId) {
		toSerialize["beneficiaryId"] = o.BeneficiaryId
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAncvDetails struct {
	value *AncvDetails
	isSet bool
}

func (v NullableAncvDetails) Get() *AncvDetails {
	return v.value
}

func (v *NullableAncvDetails) Set(val *AncvDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAncvDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAncvDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAncvDetails(val *AncvDetails) *NullableAncvDetails {
	return &NullableAncvDetails{value: val, isSet: true}
}

func (v NullableAncvDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAncvDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AncvDetails) isValidType() bool {
	var allowedEnumValues = []string{"ancv"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
