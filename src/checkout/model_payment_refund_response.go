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

// checks if the PaymentRefundResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentRefundResponse{}

// PaymentRefundResponse struct for PaymentRefundResponse
type PaymentRefundResponse struct {
	Amount Amount `json:"amount"`
	// Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). > This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Atome, Clearpay, Klarna, Ratepay, Walley, and Zip.
	LineItems []LineItem `json:"lineItems,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// Your reason for the refund request.
	MerchantRefundReason common.NullableString `json:"merchantRefundReason,omitempty"`
	// The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to refund.
	PaymentPspReference string `json:"paymentPspReference"`
	// Adyen's 16-character reference associated with the refund request.
	PspReference string `json:"pspReference"`
	// Your reference for the refund request.
	Reference *string `json:"reference,omitempty"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For more information, see how to process payments for [marketplaces](https://docs.adyen.com/marketplaces/split-payments) or [platforms](https://docs.adyen.com/platforms/online-payments/split-payments/).
	Splits []Split `json:"splits,omitempty"`
	// The status of your request. This will always have the value **received**.
	Status string `json:"status"`
	// The online store or [physical store](https://docs.adyen.com/point-of-sale/design-your-integration/determine-account-structure/#create-stores) that is processing the refund. This must be the same as the store name configured in your Customer Area.  Otherwise, you get an error and the refund fails.
	Store *string `json:"store,omitempty"`
}

// NewPaymentRefundResponse instantiates a new PaymentRefundResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRefundResponse(amount Amount, merchantAccount string, paymentPspReference string, pspReference string, status string) *PaymentRefundResponse {
	this := PaymentRefundResponse{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.PaymentPspReference = paymentPspReference
	this.PspReference = pspReference
	this.Status = status
	return &this
}

// NewPaymentRefundResponseWithDefaults instantiates a new PaymentRefundResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRefundResponseWithDefaults() *PaymentRefundResponse {
	this := PaymentRefundResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentRefundResponse) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentRefundResponse) SetAmount(v Amount) {
	o.Amount = v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentRefundResponse) GetLineItems() []LineItem {
	if o == nil || common.IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || common.IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentRefundResponse) HasLineItems() bool {
	if o != nil && !common.IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *PaymentRefundResponse) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentRefundResponse) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentRefundResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetMerchantRefundReason returns the MerchantRefundReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRefundResponse) GetMerchantRefundReason() string {
	if o == nil || common.IsNil(o.MerchantRefundReason.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantRefundReason.Get()
}

// GetMerchantRefundReasonOk returns a tuple with the MerchantRefundReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRefundResponse) GetMerchantRefundReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantRefundReason.Get(), o.MerchantRefundReason.IsSet()
}

// HasMerchantRefundReason returns a boolean if a field has been set.
func (o *PaymentRefundResponse) HasMerchantRefundReason() bool {
	if o != nil && o.MerchantRefundReason.IsSet() {
		return true
	}

	return false
}

// SetMerchantRefundReason gets a reference to the given NullableString and assigns it to the MerchantRefundReason field.
func (o *PaymentRefundResponse) SetMerchantRefundReason(v string) {
	o.MerchantRefundReason.Set(&v)
}

// SetMerchantRefundReasonNil sets the value for MerchantRefundReason to be an explicit nil
func (o *PaymentRefundResponse) SetMerchantRefundReasonNil() {
	o.MerchantRefundReason.Set(nil)
}

// UnsetMerchantRefundReason ensures that no value is present for MerchantRefundReason, not even an explicit nil
func (o *PaymentRefundResponse) UnsetMerchantRefundReason() {
	o.MerchantRefundReason.Unset()
}

// GetPaymentPspReference returns the PaymentPspReference field value
func (o *PaymentRefundResponse) GetPaymentPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentPspReference
}

// GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetPaymentPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentPspReference, true
}

// SetPaymentPspReference sets field value
func (o *PaymentRefundResponse) SetPaymentPspReference(v string) {
	o.PaymentPspReference = v
}

// GetPspReference returns the PspReference field value
func (o *PaymentRefundResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *PaymentRefundResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentRefundResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentRefundResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentRefundResponse) SetReference(v string) {
	o.Reference = &v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *PaymentRefundResponse) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *PaymentRefundResponse) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *PaymentRefundResponse) SetSplits(v []Split) {
	o.Splits = v
}

// GetStatus returns the Status field value
func (o *PaymentRefundResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentRefundResponse) SetStatus(v string) {
	o.Status = v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *PaymentRefundResponse) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundResponse) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *PaymentRefundResponse) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *PaymentRefundResponse) SetStore(v string) {
	o.Store = &v
}

func (o PaymentRefundResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRefundResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if o.MerchantRefundReason.IsSet() {
		toSerialize["merchantRefundReason"] = o.MerchantRefundReason.Get()
	}
	toSerialize["paymentPspReference"] = o.PaymentPspReference
	toSerialize["pspReference"] = o.PspReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	toSerialize["status"] = o.Status
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	return toSerialize, nil
}

type NullablePaymentRefundResponse struct {
	value *PaymentRefundResponse
	isSet bool
}

func (v NullablePaymentRefundResponse) Get() *PaymentRefundResponse {
	return v.value
}

func (v *NullablePaymentRefundResponse) Set(val *PaymentRefundResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRefundResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRefundResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRefundResponse(val *PaymentRefundResponse) *NullablePaymentRefundResponse {
	return &NullablePaymentRefundResponse{value: val, isSet: true}
}

func (v NullablePaymentRefundResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRefundResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentRefundResponse) isValidMerchantRefundReason() bool {
	var allowedEnumValues = []string{"FRAUD", "CUSTOMER REQUEST", "RETURN", "DUPLICATE", "OTHER"}
	for _, allowed := range allowedEnumValues {
		if o.GetMerchantRefundReason() == allowed {
			return true
		}
	}
	return false
}
func (o *PaymentRefundResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"received"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
