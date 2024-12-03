/*
Adyen Stored Value API

API version: 46
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storedvalue

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the StoredValueBalanceMergeRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredValueBalanceMergeRequest{}

// StoredValueBalanceMergeRequest struct for StoredValueBalanceMergeRequest
type StoredValueBalanceMergeRequest struct {
	Amount *Amount `json:"amount,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The collection that contains the type of the payment method and its specific information if available
	PaymentMethod map[string]string `json:"paymentMethod"`
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\"). Maximum length: 80 characters.
	Reference string `json:"reference"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction *string `json:"shopperInteraction,omitempty"`
	ShopperReference *string `json:"shopperReference,omitempty"`
	// The collection that contains the source payment method and its specific information if available. Note that type should not be included since it is inferred from the (target) payment method
	SourcePaymentMethod map[string]string `json:"sourcePaymentMethod"`
	// The physical store, for which this payment is processed.
	Store *string `json:"store,omitempty"`
}

// NewStoredValueBalanceMergeRequest instantiates a new StoredValueBalanceMergeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredValueBalanceMergeRequest(merchantAccount string, paymentMethod map[string]string, reference string, sourcePaymentMethod map[string]string) *StoredValueBalanceMergeRequest {
	this := StoredValueBalanceMergeRequest{}
	this.MerchantAccount = merchantAccount
	this.PaymentMethod = paymentMethod
	this.Reference = reference
	this.SourcePaymentMethod = sourcePaymentMethod
	return &this
}

// NewStoredValueBalanceMergeRequestWithDefaults instantiates a new StoredValueBalanceMergeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredValueBalanceMergeRequestWithDefaults() *StoredValueBalanceMergeRequest {
	this := StoredValueBalanceMergeRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *StoredValueBalanceMergeRequest) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *StoredValueBalanceMergeRequest) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *StoredValueBalanceMergeRequest) SetAmount(v Amount) {
	o.Amount = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StoredValueBalanceMergeRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StoredValueBalanceMergeRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *StoredValueBalanceMergeRequest) GetPaymentMethod() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetPaymentMethodOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *StoredValueBalanceMergeRequest) SetPaymentMethod(v map[string]string) {
	o.PaymentMethod = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
func (o *StoredValueBalanceMergeRequest) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *StoredValueBalanceMergeRequest) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
func (o *StoredValueBalanceMergeRequest) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetReference returns the Reference field value
func (o *StoredValueBalanceMergeRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *StoredValueBalanceMergeRequest) SetReference(v string) {
	o.Reference = v
}

// GetShopperInteraction returns the ShopperInteraction field value if set, zero value otherwise.
func (o *StoredValueBalanceMergeRequest) GetShopperInteraction() string {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		var ret string
		return ret
	}
	return *o.ShopperInteraction
}

// GetShopperInteractionOk returns a tuple with the ShopperInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetShopperInteractionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		return nil, false
	}
	return o.ShopperInteraction, true
}

// HasShopperInteraction returns a boolean if a field has been set.
func (o *StoredValueBalanceMergeRequest) HasShopperInteraction() bool {
	if o != nil && !common.IsNil(o.ShopperInteraction) {
		return true
	}

	return false
}

// SetShopperInteraction gets a reference to the given string and assigns it to the ShopperInteraction field.
func (o *StoredValueBalanceMergeRequest) SetShopperInteraction(v string) {
	o.ShopperInteraction = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *StoredValueBalanceMergeRequest) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *StoredValueBalanceMergeRequest) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *StoredValueBalanceMergeRequest) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetSourcePaymentMethod returns the SourcePaymentMethod field value
func (o *StoredValueBalanceMergeRequest) GetSourcePaymentMethod() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.SourcePaymentMethod
}

// GetSourcePaymentMethodOk returns a tuple with the SourcePaymentMethod field value
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetSourcePaymentMethodOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePaymentMethod, true
}

// SetSourcePaymentMethod sets field value
func (o *StoredValueBalanceMergeRequest) SetSourcePaymentMethod(v map[string]string) {
	o.SourcePaymentMethod = v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *StoredValueBalanceMergeRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueBalanceMergeRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *StoredValueBalanceMergeRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *StoredValueBalanceMergeRequest) SetStore(v string) {
	o.Store = &v
}

func (o StoredValueBalanceMergeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredValueBalanceMergeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentMethod"] = o.PaymentMethod
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.ShopperInteraction) {
		toSerialize["shopperInteraction"] = o.ShopperInteraction
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	toSerialize["sourcePaymentMethod"] = o.SourcePaymentMethod
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	return toSerialize, nil
}

type NullableStoredValueBalanceMergeRequest struct {
	value *StoredValueBalanceMergeRequest
	isSet bool
}

func (v NullableStoredValueBalanceMergeRequest) Get() *StoredValueBalanceMergeRequest {
	return v.value
}

func (v *NullableStoredValueBalanceMergeRequest) Set(val *StoredValueBalanceMergeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredValueBalanceMergeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredValueBalanceMergeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredValueBalanceMergeRequest(val *StoredValueBalanceMergeRequest) *NullableStoredValueBalanceMergeRequest {
	return &NullableStoredValueBalanceMergeRequest{value: val, isSet: true}
}

func (v NullableStoredValueBalanceMergeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredValueBalanceMergeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *StoredValueBalanceMergeRequest) isValidShopperInteraction() bool {
    var allowedEnumValues = []string{ "Ecommerce", "ContAuth", "Moto", "POS" }
    for _, allowed := range allowedEnumValues {
        if o.GetShopperInteraction() == allowed {
            return true
        }
    }
    return false
}

