/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the PayToDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayToDetails{}

// PayToDetails struct for PayToDetails
type PayToDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
    // Deprecated since Adyen Checkout API v49
    // Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// The shopper's banking details or payId reference, used to complete payment.
	ShopperAccountIdentifier *string `json:"shopperAccountIdentifier,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **payto**
	Type *string `json:"type,omitempty"`
}

// NewPayToDetails instantiates a new PayToDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayToDetails() *PayToDetails {
	this := PayToDetails{}
	var type_ string = "payto"
	this.Type = &type_
	return &this
}

// NewPayToDetailsWithDefaults instantiates a new PayToDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayToDetailsWithDefaults() *PayToDetails {
	this := PayToDetails{}
	var type_ string = "payto"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PayToDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayToDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PayToDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PayToDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PayToDetails) GetRecurringDetailReference() string {
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
func (o *PayToDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *PayToDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PayToDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetShopperAccountIdentifier returns the ShopperAccountIdentifier field value if set, zero value otherwise.
func (o *PayToDetails) GetShopperAccountIdentifier() string {
	if o == nil || common.IsNil(o.ShopperAccountIdentifier) {
		var ret string
		return ret
	}
	return *o.ShopperAccountIdentifier
}

// GetShopperAccountIdentifierOk returns a tuple with the ShopperAccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayToDetails) GetShopperAccountIdentifierOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperAccountIdentifier) {
		return nil, false
	}
	return o.ShopperAccountIdentifier, true
}

// HasShopperAccountIdentifier returns a boolean if a field has been set.
func (o *PayToDetails) HasShopperAccountIdentifier() bool {
	if o != nil && !common.IsNil(o.ShopperAccountIdentifier) {
		return true
	}

	return false
}

// SetShopperAccountIdentifier gets a reference to the given string and assigns it to the ShopperAccountIdentifier field.
func (o *PayToDetails) SetShopperAccountIdentifier(v string) {
	o.ShopperAccountIdentifier = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *PayToDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayToDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *PayToDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *PayToDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PayToDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayToDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PayToDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PayToDetails) SetType(v string) {
	o.Type = &v
}

func (o PayToDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayToDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.ShopperAccountIdentifier) {
		toSerialize["shopperAccountIdentifier"] = o.ShopperAccountIdentifier
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePayToDetails struct {
	value *PayToDetails
	isSet bool
}

func (v NullablePayToDetails) Get() *PayToDetails {
	return v.value
}

func (v *NullablePayToDetails) Set(val *PayToDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePayToDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePayToDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayToDetails(val *PayToDetails) *NullablePayToDetails {
	return &NullablePayToDetails{value: val, isSet: true}
}

func (v NullablePayToDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayToDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *PayToDetails) isValidType() bool {
    var allowedEnumValues = []string{ "payto" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

