/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the UpiCollectDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpiCollectDetails{}

// UpiCollectDetails struct for UpiCollectDetails
type UpiCollectDetails struct {
	// The sequence number for the debit. For example, send **2** if this is the second debit for the subscription. The sequence number is included in the notification sent to the shopper.
	BillingSequenceNumber string `json:"billingSequenceNumber"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// The `shopperNotificationReference` returned in the response when you requested to notify the shopper. Used for recurring payment only.
	ShopperNotificationReference *string `json:"shopperNotificationReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **upi_collect**
	Type string `json:"type"`
	// The virtual payment address for UPI.
	VirtualPaymentAddress *string `json:"virtualPaymentAddress,omitempty"`
}

// NewUpiCollectDetails instantiates a new UpiCollectDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpiCollectDetails(billingSequenceNumber string, type_ string) *UpiCollectDetails {
	this := UpiCollectDetails{}
	this.BillingSequenceNumber = billingSequenceNumber
	this.Type = type_
	return &this
}

// NewUpiCollectDetailsWithDefaults instantiates a new UpiCollectDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpiCollectDetailsWithDefaults() *UpiCollectDetails {
	this := UpiCollectDetails{}
	var type_ string = "upi_collect"
	this.Type = type_
	return &this
}

// GetBillingSequenceNumber returns the BillingSequenceNumber field value
func (o *UpiCollectDetails) GetBillingSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingSequenceNumber
}

// GetBillingSequenceNumberOk returns a tuple with the BillingSequenceNumber field value
// and a boolean to check if the value has been set.
func (o *UpiCollectDetails) GetBillingSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingSequenceNumber, true
}

// SetBillingSequenceNumber sets field value
func (o *UpiCollectDetails) SetBillingSequenceNumber(v string) {
	o.BillingSequenceNumber = v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *UpiCollectDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpiCollectDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *UpiCollectDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *UpiCollectDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *UpiCollectDetails) GetRecurringDetailReference() string {
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
func (o *UpiCollectDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *UpiCollectDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *UpiCollectDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetShopperNotificationReference returns the ShopperNotificationReference field value if set, zero value otherwise.
func (o *UpiCollectDetails) GetShopperNotificationReference() string {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		var ret string
		return ret
	}
	return *o.ShopperNotificationReference
}

// GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpiCollectDetails) GetShopperNotificationReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		return nil, false
	}
	return o.ShopperNotificationReference, true
}

// HasShopperNotificationReference returns a boolean if a field has been set.
func (o *UpiCollectDetails) HasShopperNotificationReference() bool {
	if o != nil && !common.IsNil(o.ShopperNotificationReference) {
		return true
	}

	return false
}

// SetShopperNotificationReference gets a reference to the given string and assigns it to the ShopperNotificationReference field.
func (o *UpiCollectDetails) SetShopperNotificationReference(v string) {
	o.ShopperNotificationReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *UpiCollectDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpiCollectDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *UpiCollectDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *UpiCollectDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value
func (o *UpiCollectDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpiCollectDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpiCollectDetails) SetType(v string) {
	o.Type = v
}

// GetVirtualPaymentAddress returns the VirtualPaymentAddress field value if set, zero value otherwise.
func (o *UpiCollectDetails) GetVirtualPaymentAddress() string {
	if o == nil || common.IsNil(o.VirtualPaymentAddress) {
		var ret string
		return ret
	}
	return *o.VirtualPaymentAddress
}

// GetVirtualPaymentAddressOk returns a tuple with the VirtualPaymentAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpiCollectDetails) GetVirtualPaymentAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.VirtualPaymentAddress) {
		return nil, false
	}
	return o.VirtualPaymentAddress, true
}

// HasVirtualPaymentAddress returns a boolean if a field has been set.
func (o *UpiCollectDetails) HasVirtualPaymentAddress() bool {
	if o != nil && !common.IsNil(o.VirtualPaymentAddress) {
		return true
	}

	return false
}

// SetVirtualPaymentAddress gets a reference to the given string and assigns it to the VirtualPaymentAddress field.
func (o *UpiCollectDetails) SetVirtualPaymentAddress(v string) {
	o.VirtualPaymentAddress = &v
}

func (o UpiCollectDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpiCollectDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billingSequenceNumber"] = o.BillingSequenceNumber
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

type NullableUpiCollectDetails struct {
	value *UpiCollectDetails
	isSet bool
}

func (v NullableUpiCollectDetails) Get() *UpiCollectDetails {
	return v.value
}

func (v *NullableUpiCollectDetails) Set(val *UpiCollectDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpiCollectDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpiCollectDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpiCollectDetails(val *UpiCollectDetails) *NullableUpiCollectDetails {
	return &NullableUpiCollectDetails{value: val, isSet: true}
}

func (v NullableUpiCollectDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpiCollectDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpiCollectDetails) isValidType() bool {
	var allowedEnumValues = []string{"upi_collect"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
