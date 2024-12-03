/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the AdditionalDataRatepay type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataRatepay{}

// AdditionalDataRatepay struct for AdditionalDataRatepay
type AdditionalDataRatepay struct {
	// Amount the customer has to pay each month.
	RatepayInstallmentAmount *string `json:"ratepay.installmentAmount,omitempty"`
	// Interest rate of this installment.
	RatepayInterestRate *string `json:"ratepay.interestRate,omitempty"`
	// Amount of the last installment.
	RatepayLastInstallmentAmount *string `json:"ratepay.lastInstallmentAmount,omitempty"`
	// Calendar day of the first payment.
	RatepayPaymentFirstday *string `json:"ratepay.paymentFirstday,omitempty"`
	// Date the merchant delivered the goods to the customer.
	RatepaydataDeliveryDate *string `json:"ratepaydata.deliveryDate,omitempty"`
	// Date by which the customer must settle the payment.
	RatepaydataDueDate *string `json:"ratepaydata.dueDate,omitempty"`
	// Invoice date, defined by the merchant. If not included, the invoice date is set to the delivery date.
	RatepaydataInvoiceDate *string `json:"ratepaydata.invoiceDate,omitempty"`
	// Identification name or number for the invoice, defined by the merchant.
	RatepaydataInvoiceId *string `json:"ratepaydata.invoiceId,omitempty"`
}

// NewAdditionalDataRatepay instantiates a new AdditionalDataRatepay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataRatepay() *AdditionalDataRatepay {
	this := AdditionalDataRatepay{}
	return &this
}

// NewAdditionalDataRatepayWithDefaults instantiates a new AdditionalDataRatepay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataRatepayWithDefaults() *AdditionalDataRatepay {
	this := AdditionalDataRatepay{}
	return &this
}

// GetRatepayInstallmentAmount returns the RatepayInstallmentAmount field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepayInstallmentAmount() string {
	if o == nil || common.IsNil(o.RatepayInstallmentAmount) {
		var ret string
		return ret
	}
	return *o.RatepayInstallmentAmount
}

// GetRatepayInstallmentAmountOk returns a tuple with the RatepayInstallmentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepayInstallmentAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepayInstallmentAmount) {
		return nil, false
	}
	return o.RatepayInstallmentAmount, true
}

// HasRatepayInstallmentAmount returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepayInstallmentAmount() bool {
	if o != nil && !common.IsNil(o.RatepayInstallmentAmount) {
		return true
	}

	return false
}

// SetRatepayInstallmentAmount gets a reference to the given string and assigns it to the RatepayInstallmentAmount field.
func (o *AdditionalDataRatepay) SetRatepayInstallmentAmount(v string) {
	o.RatepayInstallmentAmount = &v
}

// GetRatepayInterestRate returns the RatepayInterestRate field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepayInterestRate() string {
	if o == nil || common.IsNil(o.RatepayInterestRate) {
		var ret string
		return ret
	}
	return *o.RatepayInterestRate
}

// GetRatepayInterestRateOk returns a tuple with the RatepayInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepayInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepayInterestRate) {
		return nil, false
	}
	return o.RatepayInterestRate, true
}

// HasRatepayInterestRate returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepayInterestRate() bool {
	if o != nil && !common.IsNil(o.RatepayInterestRate) {
		return true
	}

	return false
}

// SetRatepayInterestRate gets a reference to the given string and assigns it to the RatepayInterestRate field.
func (o *AdditionalDataRatepay) SetRatepayInterestRate(v string) {
	o.RatepayInterestRate = &v
}

// GetRatepayLastInstallmentAmount returns the RatepayLastInstallmentAmount field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepayLastInstallmentAmount() string {
	if o == nil || common.IsNil(o.RatepayLastInstallmentAmount) {
		var ret string
		return ret
	}
	return *o.RatepayLastInstallmentAmount
}

// GetRatepayLastInstallmentAmountOk returns a tuple with the RatepayLastInstallmentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepayLastInstallmentAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepayLastInstallmentAmount) {
		return nil, false
	}
	return o.RatepayLastInstallmentAmount, true
}

// HasRatepayLastInstallmentAmount returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepayLastInstallmentAmount() bool {
	if o != nil && !common.IsNil(o.RatepayLastInstallmentAmount) {
		return true
	}

	return false
}

// SetRatepayLastInstallmentAmount gets a reference to the given string and assigns it to the RatepayLastInstallmentAmount field.
func (o *AdditionalDataRatepay) SetRatepayLastInstallmentAmount(v string) {
	o.RatepayLastInstallmentAmount = &v
}

// GetRatepayPaymentFirstday returns the RatepayPaymentFirstday field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepayPaymentFirstday() string {
	if o == nil || common.IsNil(o.RatepayPaymentFirstday) {
		var ret string
		return ret
	}
	return *o.RatepayPaymentFirstday
}

// GetRatepayPaymentFirstdayOk returns a tuple with the RatepayPaymentFirstday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepayPaymentFirstdayOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepayPaymentFirstday) {
		return nil, false
	}
	return o.RatepayPaymentFirstday, true
}

// HasRatepayPaymentFirstday returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepayPaymentFirstday() bool {
	if o != nil && !common.IsNil(o.RatepayPaymentFirstday) {
		return true
	}

	return false
}

// SetRatepayPaymentFirstday gets a reference to the given string and assigns it to the RatepayPaymentFirstday field.
func (o *AdditionalDataRatepay) SetRatepayPaymentFirstday(v string) {
	o.RatepayPaymentFirstday = &v
}

// GetRatepaydataDeliveryDate returns the RatepaydataDeliveryDate field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepaydataDeliveryDate() string {
	if o == nil || common.IsNil(o.RatepaydataDeliveryDate) {
		var ret string
		return ret
	}
	return *o.RatepaydataDeliveryDate
}

// GetRatepaydataDeliveryDateOk returns a tuple with the RatepaydataDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepaydataDeliveryDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepaydataDeliveryDate) {
		return nil, false
	}
	return o.RatepaydataDeliveryDate, true
}

// HasRatepaydataDeliveryDate returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepaydataDeliveryDate() bool {
	if o != nil && !common.IsNil(o.RatepaydataDeliveryDate) {
		return true
	}

	return false
}

// SetRatepaydataDeliveryDate gets a reference to the given string and assigns it to the RatepaydataDeliveryDate field.
func (o *AdditionalDataRatepay) SetRatepaydataDeliveryDate(v string) {
	o.RatepaydataDeliveryDate = &v
}

// GetRatepaydataDueDate returns the RatepaydataDueDate field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepaydataDueDate() string {
	if o == nil || common.IsNil(o.RatepaydataDueDate) {
		var ret string
		return ret
	}
	return *o.RatepaydataDueDate
}

// GetRatepaydataDueDateOk returns a tuple with the RatepaydataDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepaydataDueDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepaydataDueDate) {
		return nil, false
	}
	return o.RatepaydataDueDate, true
}

// HasRatepaydataDueDate returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepaydataDueDate() bool {
	if o != nil && !common.IsNil(o.RatepaydataDueDate) {
		return true
	}

	return false
}

// SetRatepaydataDueDate gets a reference to the given string and assigns it to the RatepaydataDueDate field.
func (o *AdditionalDataRatepay) SetRatepaydataDueDate(v string) {
	o.RatepaydataDueDate = &v
}

// GetRatepaydataInvoiceDate returns the RatepaydataInvoiceDate field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepaydataInvoiceDate() string {
	if o == nil || common.IsNil(o.RatepaydataInvoiceDate) {
		var ret string
		return ret
	}
	return *o.RatepaydataInvoiceDate
}

// GetRatepaydataInvoiceDateOk returns a tuple with the RatepaydataInvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepaydataInvoiceDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepaydataInvoiceDate) {
		return nil, false
	}
	return o.RatepaydataInvoiceDate, true
}

// HasRatepaydataInvoiceDate returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepaydataInvoiceDate() bool {
	if o != nil && !common.IsNil(o.RatepaydataInvoiceDate) {
		return true
	}

	return false
}

// SetRatepaydataInvoiceDate gets a reference to the given string and assigns it to the RatepaydataInvoiceDate field.
func (o *AdditionalDataRatepay) SetRatepaydataInvoiceDate(v string) {
	o.RatepaydataInvoiceDate = &v
}

// GetRatepaydataInvoiceId returns the RatepaydataInvoiceId field value if set, zero value otherwise.
func (o *AdditionalDataRatepay) GetRatepaydataInvoiceId() string {
	if o == nil || common.IsNil(o.RatepaydataInvoiceId) {
		var ret string
		return ret
	}
	return *o.RatepaydataInvoiceId
}

// GetRatepaydataInvoiceIdOk returns a tuple with the RatepaydataInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRatepay) GetRatepaydataInvoiceIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RatepaydataInvoiceId) {
		return nil, false
	}
	return o.RatepaydataInvoiceId, true
}

// HasRatepaydataInvoiceId returns a boolean if a field has been set.
func (o *AdditionalDataRatepay) HasRatepaydataInvoiceId() bool {
	if o != nil && !common.IsNil(o.RatepaydataInvoiceId) {
		return true
	}

	return false
}

// SetRatepaydataInvoiceId gets a reference to the given string and assigns it to the RatepaydataInvoiceId field.
func (o *AdditionalDataRatepay) SetRatepaydataInvoiceId(v string) {
	o.RatepaydataInvoiceId = &v
}

func (o AdditionalDataRatepay) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataRatepay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RatepayInstallmentAmount) {
		toSerialize["ratepay.installmentAmount"] = o.RatepayInstallmentAmount
	}
	if !common.IsNil(o.RatepayInterestRate) {
		toSerialize["ratepay.interestRate"] = o.RatepayInterestRate
	}
	if !common.IsNil(o.RatepayLastInstallmentAmount) {
		toSerialize["ratepay.lastInstallmentAmount"] = o.RatepayLastInstallmentAmount
	}
	if !common.IsNil(o.RatepayPaymentFirstday) {
		toSerialize["ratepay.paymentFirstday"] = o.RatepayPaymentFirstday
	}
	if !common.IsNil(o.RatepaydataDeliveryDate) {
		toSerialize["ratepaydata.deliveryDate"] = o.RatepaydataDeliveryDate
	}
	if !common.IsNil(o.RatepaydataDueDate) {
		toSerialize["ratepaydata.dueDate"] = o.RatepaydataDueDate
	}
	if !common.IsNil(o.RatepaydataInvoiceDate) {
		toSerialize["ratepaydata.invoiceDate"] = o.RatepaydataInvoiceDate
	}
	if !common.IsNil(o.RatepaydataInvoiceId) {
		toSerialize["ratepaydata.invoiceId"] = o.RatepaydataInvoiceId
	}
	return toSerialize, nil
}

type NullableAdditionalDataRatepay struct {
	value *AdditionalDataRatepay
	isSet bool
}

func (v NullableAdditionalDataRatepay) Get() *AdditionalDataRatepay {
	return v.value
}

func (v *NullableAdditionalDataRatepay) Set(val *AdditionalDataRatepay) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataRatepay) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataRatepay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataRatepay(val *AdditionalDataRatepay) *NullableAdditionalDataRatepay {
	return &NullableAdditionalDataRatepay{value: val, isSet: true}
}

func (v NullableAdditionalDataRatepay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataRatepay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



