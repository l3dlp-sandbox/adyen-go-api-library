/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the UpiIntentDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpiIntentDetails{}

// UpiIntentDetails struct for UpiIntentDetails
type UpiIntentDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// The `shopperNotificationReference` returned in the response when you requested to notify the shopper. Used for recurring payment only.
	ShopperNotificationReference *string `json:"shopperNotificationReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **upi_intent**
	Type string `json:"type"`
}

// NewUpiIntentDetails instantiates a new UpiIntentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpiIntentDetails(type_ string) *UpiIntentDetails {
	this := UpiIntentDetails{}
	this.Type = type_
	return &this
}

// NewUpiIntentDetailsWithDefaults instantiates a new UpiIntentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpiIntentDetailsWithDefaults() *UpiIntentDetails {
	this := UpiIntentDetails{}
	var type_ string = "upi_intent"
	this.Type = type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *UpiIntentDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpiIntentDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *UpiIntentDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *UpiIntentDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated
func (o *UpiIntentDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpiIntentDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *UpiIntentDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated
func (o *UpiIntentDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetShopperNotificationReference returns the ShopperNotificationReference field value if set, zero value otherwise.
func (o *UpiIntentDetails) GetShopperNotificationReference() string {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		var ret string
		return ret
	}
	return *o.ShopperNotificationReference
}

// GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpiIntentDetails) GetShopperNotificationReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		return nil, false
	}
	return o.ShopperNotificationReference, true
}

// HasShopperNotificationReference returns a boolean if a field has been set.
func (o *UpiIntentDetails) HasShopperNotificationReference() bool {
	if o != nil && !common.IsNil(o.ShopperNotificationReference) {
		return true
	}

	return false
}

// SetShopperNotificationReference gets a reference to the given string and assigns it to the ShopperNotificationReference field.
func (o *UpiIntentDetails) SetShopperNotificationReference(v string) {
	o.ShopperNotificationReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *UpiIntentDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpiIntentDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *UpiIntentDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *UpiIntentDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value
func (o *UpiIntentDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpiIntentDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpiIntentDetails) SetType(v string) {
	o.Type = v
}

func (o UpiIntentDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpiIntentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.ShopperNotificationReference) {
		toSerialize["shopperNotificationReference"] = o.ShopperNotificationReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableUpiIntentDetails struct {
	value *UpiIntentDetails
	isSet bool
}

func (v NullableUpiIntentDetails) Get() *UpiIntentDetails {
	return v.value
}

func (v *NullableUpiIntentDetails) Set(val *UpiIntentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpiIntentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpiIntentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpiIntentDetails(val *UpiIntentDetails) *NullableUpiIntentDetails {
	return &NullableUpiIntentDetails{value: val, isSet: true}
}

func (v NullableUpiIntentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpiIntentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpiIntentDetails) isValidType() bool {
	var allowedEnumValues = []string{"upi_intent"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
