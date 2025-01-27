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

// checks if the RatepayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RatepayDetails{}

// RatepayDetails struct for RatepayDetails
type RatepayDetails struct {
	// The address where to send the invoice.
	BillingAddress *string `json:"billingAddress,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The address where the goods should be delivered.
	DeliveryAddress *string `json:"deliveryAddress,omitempty"`
	// Shopper name, date of birth, phone number, and email address.
	PersonalDetails *string `json:"personalDetails,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **ratepay**
	Type string `json:"type"`
}

// NewRatepayDetails instantiates a new RatepayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatepayDetails(type_ string) *RatepayDetails {
	this := RatepayDetails{}
	this.Type = type_
	return &this
}

// NewRatepayDetailsWithDefaults instantiates a new RatepayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatepayDetailsWithDefaults() *RatepayDetails {
	this := RatepayDetails{}
	var type_ string = "ratepay"
	this.Type = type_
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *RatepayDetails) GetBillingAddress() string {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret string
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatepayDetails) GetBillingAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *RatepayDetails) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given string and assigns it to the BillingAddress field.
func (o *RatepayDetails) SetBillingAddress(v string) {
	o.BillingAddress = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *RatepayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatepayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *RatepayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *RatepayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise.
func (o *RatepayDetails) GetDeliveryAddress() string {
	if o == nil || common.IsNil(o.DeliveryAddress) {
		var ret string
		return ret
	}
	return *o.DeliveryAddress
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatepayDetails) GetDeliveryAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliveryAddress) {
		return nil, false
	}
	return o.DeliveryAddress, true
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *RatepayDetails) HasDeliveryAddress() bool {
	if o != nil && !common.IsNil(o.DeliveryAddress) {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given string and assigns it to the DeliveryAddress field.
func (o *RatepayDetails) SetDeliveryAddress(v string) {
	o.DeliveryAddress = &v
}

// GetPersonalDetails returns the PersonalDetails field value if set, zero value otherwise.
func (o *RatepayDetails) GetPersonalDetails() string {
	if o == nil || common.IsNil(o.PersonalDetails) {
		var ret string
		return ret
	}
	return *o.PersonalDetails
}

// GetPersonalDetailsOk returns a tuple with the PersonalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatepayDetails) GetPersonalDetailsOk() (*string, bool) {
	if o == nil || common.IsNil(o.PersonalDetails) {
		return nil, false
	}
	return o.PersonalDetails, true
}

// HasPersonalDetails returns a boolean if a field has been set.
func (o *RatepayDetails) HasPersonalDetails() bool {
	if o != nil && !common.IsNil(o.PersonalDetails) {
		return true
	}

	return false
}

// SetPersonalDetails gets a reference to the given string and assigns it to the PersonalDetails field.
func (o *RatepayDetails) SetPersonalDetails(v string) {
	o.PersonalDetails = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *RatepayDetails) GetRecurringDetailReference() string {
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
func (o *RatepayDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *RatepayDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *RatepayDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *RatepayDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatepayDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *RatepayDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *RatepayDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value
func (o *RatepayDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RatepayDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RatepayDetails) SetType(v string) {
	o.Type = v
}

func (o RatepayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatepayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.DeliveryAddress) {
		toSerialize["deliveryAddress"] = o.DeliveryAddress
	}
	if !common.IsNil(o.PersonalDetails) {
		toSerialize["personalDetails"] = o.PersonalDetails
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableRatepayDetails struct {
	value *RatepayDetails
	isSet bool
}

func (v NullableRatepayDetails) Get() *RatepayDetails {
	return v.value
}

func (v *NullableRatepayDetails) Set(val *RatepayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRatepayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRatepayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatepayDetails(val *RatepayDetails) *NullableRatepayDetails {
	return &NullableRatepayDetails{value: val, isSet: true}
}

func (v NullableRatepayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatepayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *RatepayDetails) isValidType() bool {
	var allowedEnumValues = []string{"ratepay", "ratepay_directdebit"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
