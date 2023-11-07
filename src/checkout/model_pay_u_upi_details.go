/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PayUUpiDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayUUpiDetails{}

// PayUUpiDetails struct for PayUUpiDetails
type PayUUpiDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// The `shopperNotificationReference` returned in the response when you requested to notify the shopper. Used for recurring payment only.
	ShopperNotificationReference *string `json:"shopperNotificationReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **payu_IN_upi**
	Type string `json:"type"`
	// The virtual payment address for UPI.
	VirtualPaymentAddress *string `json:"virtualPaymentAddress,omitempty"`
}

// NewPayUUpiDetails instantiates a new PayUUpiDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayUUpiDetails(type_ string) *PayUUpiDetails {
	this := PayUUpiDetails{}
	this.Type = type_
	return &this
}

// NewPayUUpiDetailsWithDefaults instantiates a new PayUUpiDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayUUpiDetailsWithDefaults() *PayUUpiDetails {
	this := PayUUpiDetails{}
	var type_ string = "payu_IN_upi"
	this.Type = type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PayUUpiDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayUUpiDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PayUUpiDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PayUUpiDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated
func (o *PayUUpiDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PayUUpiDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *PayUUpiDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated
func (o *PayUUpiDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetShopperNotificationReference returns the ShopperNotificationReference field value if set, zero value otherwise.
func (o *PayUUpiDetails) GetShopperNotificationReference() string {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		var ret string
		return ret
	}
	return *o.ShopperNotificationReference
}

// GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayUUpiDetails) GetShopperNotificationReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		return nil, false
	}
	return o.ShopperNotificationReference, true
}

// HasShopperNotificationReference returns a boolean if a field has been set.
func (o *PayUUpiDetails) HasShopperNotificationReference() bool {
	if o != nil && !common.IsNil(o.ShopperNotificationReference) {
		return true
	}

	return false
}

// SetShopperNotificationReference gets a reference to the given string and assigns it to the ShopperNotificationReference field.
func (o *PayUUpiDetails) SetShopperNotificationReference(v string) {
	o.ShopperNotificationReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *PayUUpiDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayUUpiDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *PayUUpiDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *PayUUpiDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value
func (o *PayUUpiDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PayUUpiDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PayUUpiDetails) SetType(v string) {
	o.Type = v
}

// GetVirtualPaymentAddress returns the VirtualPaymentAddress field value if set, zero value otherwise.
func (o *PayUUpiDetails) GetVirtualPaymentAddress() string {
	if o == nil || common.IsNil(o.VirtualPaymentAddress) {
		var ret string
		return ret
	}
	return *o.VirtualPaymentAddress
}

// GetVirtualPaymentAddressOk returns a tuple with the VirtualPaymentAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayUUpiDetails) GetVirtualPaymentAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.VirtualPaymentAddress) {
		return nil, false
	}
	return o.VirtualPaymentAddress, true
}

// HasVirtualPaymentAddress returns a boolean if a field has been set.
func (o *PayUUpiDetails) HasVirtualPaymentAddress() bool {
	if o != nil && !common.IsNil(o.VirtualPaymentAddress) {
		return true
	}

	return false
}

// SetVirtualPaymentAddress gets a reference to the given string and assigns it to the VirtualPaymentAddress field.
func (o *PayUUpiDetails) SetVirtualPaymentAddress(v string) {
	o.VirtualPaymentAddress = &v
}

func (o PayUUpiDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayUUpiDetails) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.VirtualPaymentAddress) {
		toSerialize["virtualPaymentAddress"] = o.VirtualPaymentAddress
	}
	return toSerialize, nil
}

type NullablePayUUpiDetails struct {
	value *PayUUpiDetails
	isSet bool
}

func (v NullablePayUUpiDetails) Get() *PayUUpiDetails {
	return v.value
}

func (v *NullablePayUUpiDetails) Set(val *PayUUpiDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePayUUpiDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePayUUpiDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayUUpiDetails(val *PayUUpiDetails) *NullablePayUUpiDetails {
	return &NullablePayUUpiDetails{value: val, isSet: true}
}

func (v NullablePayUUpiDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayUUpiDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PayUUpiDetails) isValidType() bool {
	var allowedEnumValues = []string{"payu_IN_upi"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
