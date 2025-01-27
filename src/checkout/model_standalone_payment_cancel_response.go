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

// checks if the StandalonePaymentCancelResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StandalonePaymentCancelResponse{}

// StandalonePaymentCancelResponse struct for StandalonePaymentCancelResponse
type StandalonePaymentCancelResponse struct {
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// The [`reference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__reqParam_reference) of the payment to cancel.
	PaymentReference string `json:"paymentReference"`
	// Adyen's 16-character reference associated with the cancel request.
	PspReference string `json:"pspReference"`
	// Your reference for the cancel request.
	Reference *string `json:"reference,omitempty"`
	// The status of your request. This will always have the value **received**.
	Status string `json:"status"`
}

// NewStandalonePaymentCancelResponse instantiates a new StandalonePaymentCancelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandalonePaymentCancelResponse(merchantAccount string, paymentReference string, pspReference string, status string) *StandalonePaymentCancelResponse {
	this := StandalonePaymentCancelResponse{}
	this.MerchantAccount = merchantAccount
	this.PaymentReference = paymentReference
	this.PspReference = pspReference
	this.Status = status
	return &this
}

// NewStandalonePaymentCancelResponseWithDefaults instantiates a new StandalonePaymentCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandalonePaymentCancelResponseWithDefaults() *StandalonePaymentCancelResponse {
	this := StandalonePaymentCancelResponse{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StandalonePaymentCancelResponse) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StandalonePaymentCancelResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentReference returns the PaymentReference field value
func (o *StandalonePaymentCancelResponse) GetPaymentReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentReference
}

// GetPaymentReferenceOk returns a tuple with the PaymentReference field value
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelResponse) GetPaymentReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentReference, true
}

// SetPaymentReference sets field value
func (o *StandalonePaymentCancelResponse) SetPaymentReference(v string) {
	o.PaymentReference = v
}

// GetPspReference returns the PspReference field value
func (o *StandalonePaymentCancelResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *StandalonePaymentCancelResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *StandalonePaymentCancelResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *StandalonePaymentCancelResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *StandalonePaymentCancelResponse) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value
func (o *StandalonePaymentCancelResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StandalonePaymentCancelResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StandalonePaymentCancelResponse) SetStatus(v string) {
	o.Status = v
}

func (o StandalonePaymentCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandalonePaymentCancelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentReference"] = o.PaymentReference
	toSerialize["pspReference"] = o.PspReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableStandalonePaymentCancelResponse struct {
	value *StandalonePaymentCancelResponse
	isSet bool
}

func (v NullableStandalonePaymentCancelResponse) Get() *StandalonePaymentCancelResponse {
	return v.value
}

func (v *NullableStandalonePaymentCancelResponse) Set(val *StandalonePaymentCancelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStandalonePaymentCancelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStandalonePaymentCancelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandalonePaymentCancelResponse(val *StandalonePaymentCancelResponse) *NullableStandalonePaymentCancelResponse {
	return &NullableStandalonePaymentCancelResponse{value: val, isSet: true}
}

func (v NullableStandalonePaymentCancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandalonePaymentCancelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *StandalonePaymentCancelResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"received"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
