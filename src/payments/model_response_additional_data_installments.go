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

// checks if the ResponseAdditionalDataInstallments type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalDataInstallments{}

// ResponseAdditionalDataInstallments struct for ResponseAdditionalDataInstallments
type ResponseAdditionalDataInstallments struct {
	// Type of installment. The value of `installmentType` should be **IssuerFinanced**.
	InstallmentPaymentDataInstallmentType *string `json:"installmentPaymentData.installmentType,omitempty"`
	// Annual interest rate.
	InstallmentPaymentDataOptionItemNrAnnualPercentageRate *string `json:"installmentPaymentData.option[itemNr].annualPercentageRate,omitempty"`
	// First Installment Amount in minor units.
	InstallmentPaymentDataOptionItemNrFirstInstallmentAmount *string `json:"installmentPaymentData.option[itemNr].firstInstallmentAmount,omitempty"`
	// Installment fee amount in minor units.
	InstallmentPaymentDataOptionItemNrInstallmentFee *string `json:"installmentPaymentData.option[itemNr].installmentFee,omitempty"`
	// Interest rate for the installment period.
	InstallmentPaymentDataOptionItemNrInterestRate *string `json:"installmentPaymentData.option[itemNr].interestRate,omitempty"`
	// Maximum number of installments possible for this payment.
	InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments *string `json:"installmentPaymentData.option[itemNr].maximumNumberOfInstallments,omitempty"`
	// Minimum number of installments possible for this payment.
	InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments *string `json:"installmentPaymentData.option[itemNr].minimumNumberOfInstallments,omitempty"`
	// Total number of installments possible for this payment.
	InstallmentPaymentDataOptionItemNrNumberOfInstallments *string `json:"installmentPaymentData.option[itemNr].numberOfInstallments,omitempty"`
	// Subsequent Installment Amount in minor units.
	InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount *string `json:"installmentPaymentData.option[itemNr].subsequentInstallmentAmount,omitempty"`
	// Total amount in minor units.
	InstallmentPaymentDataOptionItemNrTotalAmountDue *string `json:"installmentPaymentData.option[itemNr].totalAmountDue,omitempty"`
	// Possible values: * PayInInstallmentsOnly * PayInFullOnly * PayInFullOrInstallments
	InstallmentPaymentDataPaymentOptions *string `json:"installmentPaymentData.paymentOptions,omitempty"`
	// The number of installments that the payment amount should be charged with.  Example: 5 > Only relevant for card payments in countries that support installments.
	InstallmentsValue *string `json:"installments.value,omitempty"`
}

// NewResponseAdditionalDataInstallments instantiates a new ResponseAdditionalDataInstallments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalDataInstallments() *ResponseAdditionalDataInstallments {
	this := ResponseAdditionalDataInstallments{}
	return &this
}

// NewResponseAdditionalDataInstallmentsWithDefaults instantiates a new ResponseAdditionalDataInstallments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalDataInstallmentsWithDefaults() *ResponseAdditionalDataInstallments {
	this := ResponseAdditionalDataInstallments{}
	return &this
}

// GetInstallmentPaymentDataInstallmentType returns the InstallmentPaymentDataInstallmentType field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataInstallmentType() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataInstallmentType) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataInstallmentType
}

// GetInstallmentPaymentDataInstallmentTypeOk returns a tuple with the InstallmentPaymentDataInstallmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataInstallmentTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataInstallmentType) {
		return nil, false
	}
	return o.InstallmentPaymentDataInstallmentType, true
}

// HasInstallmentPaymentDataInstallmentType returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataInstallmentType() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataInstallmentType) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataInstallmentType gets a reference to the given string and assigns it to the InstallmentPaymentDataInstallmentType field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataInstallmentType(v string) {
	o.InstallmentPaymentDataInstallmentType = &v
}

// GetInstallmentPaymentDataOptionItemNrAnnualPercentageRate returns the InstallmentPaymentDataOptionItemNrAnnualPercentageRate field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate
}

// GetInstallmentPaymentDataOptionItemNrAnnualPercentageRateOk returns a tuple with the InstallmentPaymentDataOptionItemNrAnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate, true
}

// HasInstallmentPaymentDataOptionItemNrAnnualPercentageRate returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrAnnualPercentageRate gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrAnnualPercentageRate field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrAnnualPercentageRate(v string) {
	o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate = &v
}

// GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount returns the InstallmentPaymentDataOptionItemNrFirstInstallmentAmount field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount
}

// GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmountOk returns a tuple with the InstallmentPaymentDataOptionItemNrFirstInstallmentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount, true
}

// HasInstallmentPaymentDataOptionItemNrFirstInstallmentAmount returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrFirstInstallmentAmount() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrFirstInstallmentAmount field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount(v string) {
	o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount = &v
}

// GetInstallmentPaymentDataOptionItemNrInstallmentFee returns the InstallmentPaymentDataOptionItemNrInstallmentFee field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInstallmentFee() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrInstallmentFee) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrInstallmentFee
}

// GetInstallmentPaymentDataOptionItemNrInstallmentFeeOk returns a tuple with the InstallmentPaymentDataOptionItemNrInstallmentFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInstallmentFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrInstallmentFee) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrInstallmentFee, true
}

// HasInstallmentPaymentDataOptionItemNrInstallmentFee returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrInstallmentFee() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrInstallmentFee) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrInstallmentFee gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrInstallmentFee field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrInstallmentFee(v string) {
	o.InstallmentPaymentDataOptionItemNrInstallmentFee = &v
}

// GetInstallmentPaymentDataOptionItemNrInterestRate returns the InstallmentPaymentDataOptionItemNrInterestRate field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInterestRate() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrInterestRate) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrInterestRate
}

// GetInstallmentPaymentDataOptionItemNrInterestRateOk returns a tuple with the InstallmentPaymentDataOptionItemNrInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrInterestRate) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrInterestRate, true
}

// HasInstallmentPaymentDataOptionItemNrInterestRate returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrInterestRate() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrInterestRate) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrInterestRate gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrInterestRate field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrInterestRate(v string) {
	o.InstallmentPaymentDataOptionItemNrInterestRate = &v
}

// GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments returns the InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments
}

// GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallmentsOk returns a tuple with the InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallmentsOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments, true
}

// HasInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments(v string) {
	o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments = &v
}

// GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments returns the InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments
}

// GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallmentsOk returns a tuple with the InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallmentsOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments, true
}

// HasInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments(v string) {
	o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments = &v
}

// GetInstallmentPaymentDataOptionItemNrNumberOfInstallments returns the InstallmentPaymentDataOptionItemNrNumberOfInstallments field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrNumberOfInstallments() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrNumberOfInstallments) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrNumberOfInstallments
}

// GetInstallmentPaymentDataOptionItemNrNumberOfInstallmentsOk returns a tuple with the InstallmentPaymentDataOptionItemNrNumberOfInstallments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrNumberOfInstallmentsOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrNumberOfInstallments) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrNumberOfInstallments, true
}

// HasInstallmentPaymentDataOptionItemNrNumberOfInstallments returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrNumberOfInstallments() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrNumberOfInstallments) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrNumberOfInstallments gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrNumberOfInstallments field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrNumberOfInstallments(v string) {
	o.InstallmentPaymentDataOptionItemNrNumberOfInstallments = &v
}

// GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount returns the InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount
}

// GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmountOk returns a tuple with the InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount, true
}

// HasInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount(v string) {
	o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount = &v
}

// GetInstallmentPaymentDataOptionItemNrTotalAmountDue returns the InstallmentPaymentDataOptionItemNrTotalAmountDue field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrTotalAmountDue() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrTotalAmountDue) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataOptionItemNrTotalAmountDue
}

// GetInstallmentPaymentDataOptionItemNrTotalAmountDueOk returns a tuple with the InstallmentPaymentDataOptionItemNrTotalAmountDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrTotalAmountDueOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataOptionItemNrTotalAmountDue) {
		return nil, false
	}
	return o.InstallmentPaymentDataOptionItemNrTotalAmountDue, true
}

// HasInstallmentPaymentDataOptionItemNrTotalAmountDue returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrTotalAmountDue() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataOptionItemNrTotalAmountDue) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataOptionItemNrTotalAmountDue gets a reference to the given string and assigns it to the InstallmentPaymentDataOptionItemNrTotalAmountDue field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrTotalAmountDue(v string) {
	o.InstallmentPaymentDataOptionItemNrTotalAmountDue = &v
}

// GetInstallmentPaymentDataPaymentOptions returns the InstallmentPaymentDataPaymentOptions field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataPaymentOptions() string {
	if o == nil || common.IsNil(o.InstallmentPaymentDataPaymentOptions) {
		var ret string
		return ret
	}
	return *o.InstallmentPaymentDataPaymentOptions
}

// GetInstallmentPaymentDataPaymentOptionsOk returns a tuple with the InstallmentPaymentDataPaymentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataPaymentOptionsOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentPaymentDataPaymentOptions) {
		return nil, false
	}
	return o.InstallmentPaymentDataPaymentOptions, true
}

// HasInstallmentPaymentDataPaymentOptions returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataPaymentOptions() bool {
	if o != nil && !common.IsNil(o.InstallmentPaymentDataPaymentOptions) {
		return true
	}

	return false
}

// SetInstallmentPaymentDataPaymentOptions gets a reference to the given string and assigns it to the InstallmentPaymentDataPaymentOptions field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataPaymentOptions(v string) {
	o.InstallmentPaymentDataPaymentOptions = &v
}

// GetInstallmentsValue returns the InstallmentsValue field value if set, zero value otherwise.
func (o *ResponseAdditionalDataInstallments) GetInstallmentsValue() string {
	if o == nil || common.IsNil(o.InstallmentsValue) {
		var ret string
		return ret
	}
	return *o.InstallmentsValue
}

// GetInstallmentsValueOk returns a tuple with the InstallmentsValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataInstallments) GetInstallmentsValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallmentsValue) {
		return nil, false
	}
	return o.InstallmentsValue, true
}

// HasInstallmentsValue returns a boolean if a field has been set.
func (o *ResponseAdditionalDataInstallments) HasInstallmentsValue() bool {
	if o != nil && !common.IsNil(o.InstallmentsValue) {
		return true
	}

	return false
}

// SetInstallmentsValue gets a reference to the given string and assigns it to the InstallmentsValue field.
func (o *ResponseAdditionalDataInstallments) SetInstallmentsValue(v string) {
	o.InstallmentsValue = &v
}

func (o ResponseAdditionalDataInstallments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalDataInstallments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.InstallmentPaymentDataInstallmentType) {
		toSerialize["installmentPaymentData.installmentType"] = o.InstallmentPaymentDataInstallmentType
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate) {
		toSerialize["installmentPaymentData.option[itemNr].annualPercentageRate"] = o.InstallmentPaymentDataOptionItemNrAnnualPercentageRate
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount) {
		toSerialize["installmentPaymentData.option[itemNr].firstInstallmentAmount"] = o.InstallmentPaymentDataOptionItemNrFirstInstallmentAmount
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrInstallmentFee) {
		toSerialize["installmentPaymentData.option[itemNr].installmentFee"] = o.InstallmentPaymentDataOptionItemNrInstallmentFee
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrInterestRate) {
		toSerialize["installmentPaymentData.option[itemNr].interestRate"] = o.InstallmentPaymentDataOptionItemNrInterestRate
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments) {
		toSerialize["installmentPaymentData.option[itemNr].maximumNumberOfInstallments"] = o.InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments) {
		toSerialize["installmentPaymentData.option[itemNr].minimumNumberOfInstallments"] = o.InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrNumberOfInstallments) {
		toSerialize["installmentPaymentData.option[itemNr].numberOfInstallments"] = o.InstallmentPaymentDataOptionItemNrNumberOfInstallments
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount) {
		toSerialize["installmentPaymentData.option[itemNr].subsequentInstallmentAmount"] = o.InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount
	}
	if !common.IsNil(o.InstallmentPaymentDataOptionItemNrTotalAmountDue) {
		toSerialize["installmentPaymentData.option[itemNr].totalAmountDue"] = o.InstallmentPaymentDataOptionItemNrTotalAmountDue
	}
	if !common.IsNil(o.InstallmentPaymentDataPaymentOptions) {
		toSerialize["installmentPaymentData.paymentOptions"] = o.InstallmentPaymentDataPaymentOptions
	}
	if !common.IsNil(o.InstallmentsValue) {
		toSerialize["installments.value"] = o.InstallmentsValue
	}
	return toSerialize, nil
}

type NullableResponseAdditionalDataInstallments struct {
	value *ResponseAdditionalDataInstallments
	isSet bool
}

func (v NullableResponseAdditionalDataInstallments) Get() *ResponseAdditionalDataInstallments {
	return v.value
}

func (v *NullableResponseAdditionalDataInstallments) Set(val *ResponseAdditionalDataInstallments) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalDataInstallments) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalDataInstallments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalDataInstallments(val *ResponseAdditionalDataInstallments) *NullableResponseAdditionalDataInstallments {
	return &NullableResponseAdditionalDataInstallments{value: val, isSet: true}
}

func (v NullableResponseAdditionalDataInstallments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalDataInstallments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



