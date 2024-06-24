/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the PaypalUpdateOrderRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaypalUpdateOrderRequest{}

// PaypalUpdateOrderRequest struct for PaypalUpdateOrderRequest
type PaypalUpdateOrderRequest struct {
	Amount *Amount `json:"amount,omitempty"`
	// The list of new delivery methods and the cost of each.
	DeliveryMethods []DeliveryMethod `json:"deliveryMethods,omitempty"`
	// The `paymentData` from the client side. This value changes every time you make a `/paypal/updateOrder` request.
	PaymentData *string `json:"paymentData,omitempty"`
	// The original `pspReference` from the `/payments` response.
	PspReference *string `json:"pspReference,omitempty"`
	// The original `sessionId` from the `/sessions` response.
	SessionId *string   `json:"sessionId,omitempty"`
	TaxTotal  *TaxTotal `json:"taxTotal,omitempty"`
}

// NewPaypalUpdateOrderRequest instantiates a new PaypalUpdateOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalUpdateOrderRequest() *PaypalUpdateOrderRequest {
	this := PaypalUpdateOrderRequest{}
	return &this
}

// NewPaypalUpdateOrderRequestWithDefaults instantiates a new PaypalUpdateOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalUpdateOrderRequestWithDefaults() *PaypalUpdateOrderRequest {
	this := PaypalUpdateOrderRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaypalUpdateOrderRequest) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderRequest) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaypalUpdateOrderRequest) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *PaypalUpdateOrderRequest) SetAmount(v Amount) {
	o.Amount = &v
}

// GetDeliveryMethods returns the DeliveryMethods field value if set, zero value otherwise.
func (o *PaypalUpdateOrderRequest) GetDeliveryMethods() []DeliveryMethod {
	if o == nil || common.IsNil(o.DeliveryMethods) {
		var ret []DeliveryMethod
		return ret
	}
	return o.DeliveryMethods
}

// GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderRequest) GetDeliveryMethodsOk() ([]DeliveryMethod, bool) {
	if o == nil || common.IsNil(o.DeliveryMethods) {
		return nil, false
	}
	return o.DeliveryMethods, true
}

// HasDeliveryMethods returns a boolean if a field has been set.
func (o *PaypalUpdateOrderRequest) HasDeliveryMethods() bool {
	if o != nil && !common.IsNil(o.DeliveryMethods) {
		return true
	}

	return false
}

// SetDeliveryMethods gets a reference to the given []DeliveryMethod and assigns it to the DeliveryMethods field.
func (o *PaypalUpdateOrderRequest) SetDeliveryMethods(v []DeliveryMethod) {
	o.DeliveryMethods = v
}

// GetPaymentData returns the PaymentData field value if set, zero value otherwise.
func (o *PaypalUpdateOrderRequest) GetPaymentData() string {
	if o == nil || common.IsNil(o.PaymentData) {
		var ret string
		return ret
	}
	return *o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderRequest) GetPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentData) {
		return nil, false
	}
	return o.PaymentData, true
}

// HasPaymentData returns a boolean if a field has been set.
func (o *PaypalUpdateOrderRequest) HasPaymentData() bool {
	if o != nil && !common.IsNil(o.PaymentData) {
		return true
	}

	return false
}

// SetPaymentData gets a reference to the given string and assigns it to the PaymentData field.
func (o *PaypalUpdateOrderRequest) SetPaymentData(v string) {
	o.PaymentData = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *PaypalUpdateOrderRequest) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderRequest) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *PaypalUpdateOrderRequest) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *PaypalUpdateOrderRequest) SetPspReference(v string) {
	o.PspReference = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *PaypalUpdateOrderRequest) GetSessionId() string {
	if o == nil || common.IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderRequest) GetSessionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *PaypalUpdateOrderRequest) HasSessionId() bool {
	if o != nil && !common.IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *PaypalUpdateOrderRequest) SetSessionId(v string) {
	o.SessionId = &v
}

// GetTaxTotal returns the TaxTotal field value if set, zero value otherwise.
func (o *PaypalUpdateOrderRequest) GetTaxTotal() TaxTotal {
	if o == nil || common.IsNil(o.TaxTotal) {
		var ret TaxTotal
		return ret
	}
	return *o.TaxTotal
}

// GetTaxTotalOk returns a tuple with the TaxTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderRequest) GetTaxTotalOk() (*TaxTotal, bool) {
	if o == nil || common.IsNil(o.TaxTotal) {
		return nil, false
	}
	return o.TaxTotal, true
}

// HasTaxTotal returns a boolean if a field has been set.
func (o *PaypalUpdateOrderRequest) HasTaxTotal() bool {
	if o != nil && !common.IsNil(o.TaxTotal) {
		return true
	}

	return false
}

// SetTaxTotal gets a reference to the given TaxTotal and assigns it to the TaxTotal field.
func (o *PaypalUpdateOrderRequest) SetTaxTotal(v TaxTotal) {
	o.TaxTotal = &v
}

func (o PaypalUpdateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaypalUpdateOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.DeliveryMethods) {
		toSerialize["deliveryMethods"] = o.DeliveryMethods
	}
	if !common.IsNil(o.PaymentData) {
		toSerialize["paymentData"] = o.PaymentData
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	if !common.IsNil(o.TaxTotal) {
		toSerialize["taxTotal"] = o.TaxTotal
	}
	return toSerialize, nil
}

type NullablePaypalUpdateOrderRequest struct {
	value *PaypalUpdateOrderRequest
	isSet bool
}

func (v NullablePaypalUpdateOrderRequest) Get() *PaypalUpdateOrderRequest {
	return v.value
}

func (v *NullablePaypalUpdateOrderRequest) Set(val *PaypalUpdateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalUpdateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalUpdateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalUpdateOrderRequest(val *PaypalUpdateOrderRequest) *NullablePaypalUpdateOrderRequest {
	return &NullablePaypalUpdateOrderRequest{value: val, isSet: true}
}

func (v NullablePaypalUpdateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalUpdateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
