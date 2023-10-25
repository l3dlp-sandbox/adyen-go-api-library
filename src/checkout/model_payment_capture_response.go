/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PaymentCaptureResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentCaptureResponse{}

// PaymentCaptureResponse struct for PaymentCaptureResponse
type PaymentCaptureResponse struct {
	Amount Amount `json:"amount"`
	// Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). > This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Atome, Clearpay, Klarna, Ratepay, Walley, and Zip.
	LineItems []LineItem `json:"lineItems,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to capture.
	PaymentPspReference     string                   `json:"paymentPspReference"`
	PlatformChargebackLogic *PlatformChargebackLogic `json:"platformChargebackLogic,omitempty"`
	// Adyen's 16-character reference associated with the capture request.
	PspReference string `json:"pspReference"`
	// Your reference for the capture request.
	Reference *string `json:"reference,omitempty"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information).
	Splits []Split `json:"splits,omitempty"`
	// The status of your request. This will always have the value **received**.
	Status string `json:"status"`
	// List of sub-merchants.
	SubMerchants []SubMerchantInfo `json:"subMerchants,omitempty"`
}

// NewPaymentCaptureResponse instantiates a new PaymentCaptureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCaptureResponse(amount Amount, merchantAccount string, paymentPspReference string, pspReference string, status string) *PaymentCaptureResponse {
	this := PaymentCaptureResponse{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.PaymentPspReference = paymentPspReference
	this.PspReference = pspReference
	this.Status = status
	return &this
}

// NewPaymentCaptureResponseWithDefaults instantiates a new PaymentCaptureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCaptureResponseWithDefaults() *PaymentCaptureResponse {
	this := PaymentCaptureResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentCaptureResponse) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentCaptureResponse) SetAmount(v Amount) {
	o.Amount = v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentCaptureResponse) GetLineItems() []LineItem {
	if o == nil || common.IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || common.IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentCaptureResponse) HasLineItems() bool {
	if o != nil && !common.IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *PaymentCaptureResponse) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentCaptureResponse) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentCaptureResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentPspReference returns the PaymentPspReference field value
func (o *PaymentCaptureResponse) GetPaymentPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentPspReference
}

// GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetPaymentPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentPspReference, true
}

// SetPaymentPspReference sets field value
func (o *PaymentCaptureResponse) SetPaymentPspReference(v string) {
	o.PaymentPspReference = v
}

// GetPlatformChargebackLogic returns the PlatformChargebackLogic field value if set, zero value otherwise.
func (o *PaymentCaptureResponse) GetPlatformChargebackLogic() PlatformChargebackLogic {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		var ret PlatformChargebackLogic
		return ret
	}
	return *o.PlatformChargebackLogic
}

// GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool) {
	if o == nil || common.IsNil(o.PlatformChargebackLogic) {
		return nil, false
	}
	return o.PlatformChargebackLogic, true
}

// HasPlatformChargebackLogic returns a boolean if a field has been set.
func (o *PaymentCaptureResponse) HasPlatformChargebackLogic() bool {
	if o != nil && !common.IsNil(o.PlatformChargebackLogic) {
		return true
	}

	return false
}

// SetPlatformChargebackLogic gets a reference to the given PlatformChargebackLogic and assigns it to the PlatformChargebackLogic field.
func (o *PaymentCaptureResponse) SetPlatformChargebackLogic(v PlatformChargebackLogic) {
	o.PlatformChargebackLogic = &v
}

// GetPspReference returns the PspReference field value
func (o *PaymentCaptureResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *PaymentCaptureResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentCaptureResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentCaptureResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentCaptureResponse) SetReference(v string) {
	o.Reference = &v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *PaymentCaptureResponse) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *PaymentCaptureResponse) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *PaymentCaptureResponse) SetSplits(v []Split) {
	o.Splits = v
}

// GetStatus returns the Status field value
func (o *PaymentCaptureResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentCaptureResponse) SetStatus(v string) {
	o.Status = v
}

// GetSubMerchants returns the SubMerchants field value if set, zero value otherwise.
func (o *PaymentCaptureResponse) GetSubMerchants() []SubMerchantInfo {
	if o == nil || common.IsNil(o.SubMerchants) {
		var ret []SubMerchantInfo
		return ret
	}
	return o.SubMerchants
}

// GetSubMerchantsOk returns a tuple with the SubMerchants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCaptureResponse) GetSubMerchantsOk() ([]SubMerchantInfo, bool) {
	if o == nil || common.IsNil(o.SubMerchants) {
		return nil, false
	}
	return o.SubMerchants, true
}

// HasSubMerchants returns a boolean if a field has been set.
func (o *PaymentCaptureResponse) HasSubMerchants() bool {
	if o != nil && !common.IsNil(o.SubMerchants) {
		return true
	}

	return false
}

// SetSubMerchants gets a reference to the given []SubMerchantInfo and assigns it to the SubMerchants field.
func (o *PaymentCaptureResponse) SetSubMerchants(v []SubMerchantInfo) {
	o.SubMerchants = v
}

func (o PaymentCaptureResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCaptureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentPspReference"] = o.PaymentPspReference
	if !common.IsNil(o.PlatformChargebackLogic) {
		toSerialize["platformChargebackLogic"] = o.PlatformChargebackLogic
	}
	toSerialize["pspReference"] = o.PspReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	toSerialize["status"] = o.Status
	if !common.IsNil(o.SubMerchants) {
		toSerialize["subMerchants"] = o.SubMerchants
	}
	return toSerialize, nil
}

type NullablePaymentCaptureResponse struct {
	value *PaymentCaptureResponse
	isSet bool
}

func (v NullablePaymentCaptureResponse) Get() *PaymentCaptureResponse {
	return v.value
}

func (v *NullablePaymentCaptureResponse) Set(val *PaymentCaptureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCaptureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCaptureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCaptureResponse(val *PaymentCaptureResponse) *NullablePaymentCaptureResponse {
	return &NullablePaymentCaptureResponse{value: val, isSet: true}
}

func (v NullablePaymentCaptureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCaptureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentCaptureResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"received"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
