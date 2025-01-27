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

// checks if the CheckoutOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutOrderResponse{}

// CheckoutOrderResponse struct for CheckoutOrderResponse
type CheckoutOrderResponse struct {
	Amount *Amount `json:"amount,omitempty"`
	// The expiry date for the order.
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// The encrypted order data.
	OrderData *string `json:"orderData,omitempty"`
	// The `pspReference` that belongs to the order.
	PspReference string `json:"pspReference"`
	// The merchant reference for the order.
	Reference       *string `json:"reference,omitempty"`
	RemainingAmount *Amount `json:"remainingAmount,omitempty"`
}

// NewCheckoutOrderResponse instantiates a new CheckoutOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutOrderResponse(pspReference string) *CheckoutOrderResponse {
	this := CheckoutOrderResponse{}
	this.PspReference = pspReference
	return &this
}

// NewCheckoutOrderResponseWithDefaults instantiates a new CheckoutOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutOrderResponseWithDefaults() *CheckoutOrderResponse {
	this := CheckoutOrderResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CheckoutOrderResponse) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutOrderResponse) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CheckoutOrderResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *CheckoutOrderResponse) SetAmount(v Amount) {
	o.Amount = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CheckoutOrderResponse) GetExpiresAt() string {
	if o == nil || common.IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutOrderResponse) GetExpiresAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CheckoutOrderResponse) HasExpiresAt() bool {
	if o != nil && !common.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CheckoutOrderResponse) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetOrderData returns the OrderData field value if set, zero value otherwise.
func (o *CheckoutOrderResponse) GetOrderData() string {
	if o == nil || common.IsNil(o.OrderData) {
		var ret string
		return ret
	}
	return *o.OrderData
}

// GetOrderDataOk returns a tuple with the OrderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutOrderResponse) GetOrderDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderData) {
		return nil, false
	}
	return o.OrderData, true
}

// HasOrderData returns a boolean if a field has been set.
func (o *CheckoutOrderResponse) HasOrderData() bool {
	if o != nil && !common.IsNil(o.OrderData) {
		return true
	}

	return false
}

// SetOrderData gets a reference to the given string and assigns it to the OrderData field.
func (o *CheckoutOrderResponse) SetOrderData(v string) {
	o.OrderData = &v
}

// GetPspReference returns the PspReference field value
func (o *CheckoutOrderResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *CheckoutOrderResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *CheckoutOrderResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CheckoutOrderResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutOrderResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CheckoutOrderResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CheckoutOrderResponse) SetReference(v string) {
	o.Reference = &v
}

// GetRemainingAmount returns the RemainingAmount field value if set, zero value otherwise.
func (o *CheckoutOrderResponse) GetRemainingAmount() Amount {
	if o == nil || common.IsNil(o.RemainingAmount) {
		var ret Amount
		return ret
	}
	return *o.RemainingAmount
}

// GetRemainingAmountOk returns a tuple with the RemainingAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutOrderResponse) GetRemainingAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.RemainingAmount) {
		return nil, false
	}
	return o.RemainingAmount, true
}

// HasRemainingAmount returns a boolean if a field has been set.
func (o *CheckoutOrderResponse) HasRemainingAmount() bool {
	if o != nil && !common.IsNil(o.RemainingAmount) {
		return true
	}

	return false
}

// SetRemainingAmount gets a reference to the given Amount and assigns it to the RemainingAmount field.
func (o *CheckoutOrderResponse) SetRemainingAmount(v Amount) {
	o.RemainingAmount = &v
}

func (o CheckoutOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !common.IsNil(o.OrderData) {
		toSerialize["orderData"] = o.OrderData
	}
	toSerialize["pspReference"] = o.PspReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.RemainingAmount) {
		toSerialize["remainingAmount"] = o.RemainingAmount
	}
	return toSerialize, nil
}

type NullableCheckoutOrderResponse struct {
	value *CheckoutOrderResponse
	isSet bool
}

func (v NullableCheckoutOrderResponse) Get() *CheckoutOrderResponse {
	return v.value
}

func (v *NullableCheckoutOrderResponse) Set(val *CheckoutOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutOrderResponse(val *CheckoutOrderResponse) *NullableCheckoutOrderResponse {
	return &NullableCheckoutOrderResponse{value: val, isSet: true}
}

func (v NullableCheckoutOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
