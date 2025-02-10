/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the StoredPaymentMethodRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredPaymentMethodRequest{}

// StoredPaymentMethodRequest struct for StoredPaymentMethodRequest
type StoredPaymentMethodRequest struct {
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string               `json:"merchantAccount"`
	PaymentMethod   PaymentMethodToStore `json:"paymentMethod"`
	// Defines a recurring payment type. Required when creating a token to store payment details. Allowed values: * `Subscription` – A transaction for a fixed or variable amount, which follows a fixed schedule. * `CardOnFile` – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * `UnscheduledCardOnFile` – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder's balance drops below a certain amount.
	RecurringProcessingModel string `json:"recurringProcessingModel"`
	// The shopper's email address. We recommend that you provide this data, as it is used in velocity fraud checks.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	// The IP address of a shopper.
	ShopperIP *string `json:"shopperIP,omitempty"`
	// A unique identifier for the shopper (for example, user ID or account ID).
	ShopperReference string `json:"shopperReference"`
}

// NewStoredPaymentMethodRequest instantiates a new StoredPaymentMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredPaymentMethodRequest(merchantAccount string, paymentMethod PaymentMethodToStore, recurringProcessingModel string, shopperReference string) *StoredPaymentMethodRequest {
	this := StoredPaymentMethodRequest{}
	this.MerchantAccount = merchantAccount
	this.PaymentMethod = paymentMethod
	this.RecurringProcessingModel = recurringProcessingModel
	this.ShopperReference = shopperReference
	return &this
}

// NewStoredPaymentMethodRequestWithDefaults instantiates a new StoredPaymentMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredPaymentMethodRequestWithDefaults() *StoredPaymentMethodRequest {
	this := StoredPaymentMethodRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StoredPaymentMethodRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StoredPaymentMethodRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *StoredPaymentMethodRequest) GetPaymentMethod() PaymentMethodToStore {
	if o == nil {
		var ret PaymentMethodToStore
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodRequest) GetPaymentMethodOk() (*PaymentMethodToStore, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *StoredPaymentMethodRequest) SetPaymentMethod(v PaymentMethodToStore) {
	o.PaymentMethod = v
}

// GetRecurringProcessingModel returns the RecurringProcessingModel field value
func (o *StoredPaymentMethodRequest) GetRecurringProcessingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringProcessingModel
}

// GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field value
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodRequest) GetRecurringProcessingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurringProcessingModel, true
}

// SetRecurringProcessingModel sets field value
func (o *StoredPaymentMethodRequest) SetRecurringProcessingModel(v string) {
	o.RecurringProcessingModel = v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *StoredPaymentMethodRequest) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodRequest) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *StoredPaymentMethodRequest) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *StoredPaymentMethodRequest) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetShopperIP returns the ShopperIP field value if set, zero value otherwise.
func (o *StoredPaymentMethodRequest) GetShopperIP() string {
	if o == nil || common.IsNil(o.ShopperIP) {
		var ret string
		return ret
	}
	return *o.ShopperIP
}

// GetShopperIPOk returns a tuple with the ShopperIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodRequest) GetShopperIPOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperIP) {
		return nil, false
	}
	return o.ShopperIP, true
}

// HasShopperIP returns a boolean if a field has been set.
func (o *StoredPaymentMethodRequest) HasShopperIP() bool {
	if o != nil && !common.IsNil(o.ShopperIP) {
		return true
	}

	return false
}

// SetShopperIP gets a reference to the given string and assigns it to the ShopperIP field.
func (o *StoredPaymentMethodRequest) SetShopperIP(v string) {
	o.ShopperIP = &v
}

// GetShopperReference returns the ShopperReference field value
func (o *StoredPaymentMethodRequest) GetShopperReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperReference, true
}

// SetShopperReference sets field value
func (o *StoredPaymentMethodRequest) SetShopperReference(v string) {
	o.ShopperReference = v
}

func (o StoredPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredPaymentMethodRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentMethod"] = o.PaymentMethod
	toSerialize["recurringProcessingModel"] = o.RecurringProcessingModel
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	if !common.IsNil(o.ShopperIP) {
		toSerialize["shopperIP"] = o.ShopperIP
	}
	toSerialize["shopperReference"] = o.ShopperReference
	return toSerialize, nil
}

type NullableStoredPaymentMethodRequest struct {
	value *StoredPaymentMethodRequest
	isSet bool
}

func (v NullableStoredPaymentMethodRequest) Get() *StoredPaymentMethodRequest {
	return v.value
}

func (v *NullableStoredPaymentMethodRequest) Set(val *StoredPaymentMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredPaymentMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredPaymentMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredPaymentMethodRequest(val *StoredPaymentMethodRequest) *NullableStoredPaymentMethodRequest {
	return &NullableStoredPaymentMethodRequest{value: val, isSet: true}
}

func (v NullableStoredPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredPaymentMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *StoredPaymentMethodRequest) isValidRecurringProcessingModel() bool {
	var allowedEnumValues = []string{"CardOnFile", "Subscription", "UnscheduledCardOnFile"}
	for _, allowed := range allowedEnumValues {
		if o.GetRecurringProcessingModel() == allowed {
			return true
		}
	}
	return false
}
