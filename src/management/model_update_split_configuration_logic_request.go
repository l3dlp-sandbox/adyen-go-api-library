/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the UpdateSplitConfigurationLogicRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateSplitConfigurationLogicRequest{}

// UpdateSplitConfigurationLogicRequest struct for UpdateSplitConfigurationLogicRequest
type UpdateSplitConfigurationLogicRequest struct {
	// Deducts the acquiring fees (the aggregated amount of interchange and scheme fee) from the specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AcquiringFees        *string               `json:"acquiringFees,omitempty"`
	AdditionalCommission *AdditionalCommission `json:"additionalCommission,omitempty"`
	// Deducts the transaction fee due to Adyen under [blended rates](https://www.adyen.com/knowledge-hub/guides/payments-training-guide/get-the-best-from-your-card-processing) from the specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AdyenCommission *string `json:"adyenCommission,omitempty"`
	// Deducts the fees due to Adyen (markup or commission) from the specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AdyenFees *string `json:"adyenFees,omitempty"`
	// Deducts the transaction fee due to Adyen under [Interchange ++ pricing](https://www.adyen.com/what-is-interchange) from the specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	AdyenMarkup *string `json:"adyenMarkup,omitempty"`
	// Specifies how and from which balance account(s) to deduct the chargeback amount.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**, **deductAccordingToSplitRatio**.
	Chargeback *string `json:"chargeback,omitempty"`
	// Deducts the chargeback costs from the specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**
	ChargebackCostAllocation *string    `json:"chargebackCostAllocation,omitempty"`
	Commission               Commission `json:"commission"`
	// Deducts the interchange fee from specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	Interchange *string `json:"interchange,omitempty"`
	// Deducts all transaction fees incurred by the payment from the specified balance account. The transaction fees include the acquiring fees (interchange and scheme fee), and the fees due to Adyen (markup or commission). You can book any and all these fees to different balance account by specifying other transaction fee parameters in your split configuration profile:  - [`adyenCommission`](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/(merchantId)/splitConfigurations#request-rules-splitLogic-adyenCommission): The transaction fee due to Adyen under [blended rates](https://www.adyen.com/knowledge-hub/interchange-fees-explained#interchange-vs-blended). - [`adyenMarkup`](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/(merchantId)/splitConfigurations#request-rules-splitLogic-adyenMarkup): The transaction fee due to Adyen under [Interchange ++ pricing](https://www.adyen.com/knowledge-hub/interchange-fees-explained#interchange-vs-blended). - [`schemeFee`](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/(merchantId)/splitConfigurations#request-rules-splitLogic-schemeFee): The fee paid to the card scheme for using their network. - [`interchange`](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/(merchantId)/splitConfigurations#request-rules-splitLogic-interchange): The fee paid to the issuer for each payment transaction made with the card network. - [`adyenFees`](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/(merchantId)/splitConfigurations#request-rules-splitLogic-adyenFees): The aggregated amount of Adyen's commission and markup. - [`acquiringFees`](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/(merchantId)/splitConfigurations#request-rules-splitLogic-acquiringFees): The aggregated amount of the interchange and scheme fees.  If you don't include at least one transaction fee type in the `splitLogic` object, Adyen updates the payment request with the `paymentFee` parameter, booking all transaction fees to your platform's liable balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	PaymentFee *string `json:"paymentFee,omitempty"`
	// Specifies how and from which balance account(s) to deduct the refund amount.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**, **deductAccordingToSplitRatio**
	Refund *string `json:"refund,omitempty"`
	// Deducts the refund costs from the specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**
	RefundCostAllocation *string `json:"refundCostAllocation,omitempty"`
	// Books the amount left over after currency conversion to the specified balance account.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**.
	Remainder *string `json:"remainder,omitempty"`
	// Deducts the scheme fee from the specified balance account.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	SchemeFee *string `json:"schemeFee,omitempty"`
	// Unique identifier of the collection of split instructions that are applied when the rule conditions are met.
	SplitLogicId *string `json:"splitLogicId,omitempty"`
	// Books the surcharge amount to the specified balance account.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**
	Surcharge *string `json:"surcharge,omitempty"`
	// Books the tips (gratuity) to the specified balance account.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**.
	Tip *string `json:"tip,omitempty"`
}

// NewUpdateSplitConfigurationLogicRequest instantiates a new UpdateSplitConfigurationLogicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSplitConfigurationLogicRequest(commission Commission) *UpdateSplitConfigurationLogicRequest {
	this := UpdateSplitConfigurationLogicRequest{}
	this.Commission = commission
	return &this
}

// NewUpdateSplitConfigurationLogicRequestWithDefaults instantiates a new UpdateSplitConfigurationLogicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSplitConfigurationLogicRequestWithDefaults() *UpdateSplitConfigurationLogicRequest {
	this := UpdateSplitConfigurationLogicRequest{}
	return &this
}

// GetAcquiringFees returns the AcquiringFees field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetAcquiringFees() string {
	if o == nil || common.IsNil(o.AcquiringFees) {
		var ret string
		return ret
	}
	return *o.AcquiringFees
}

// GetAcquiringFeesOk returns a tuple with the AcquiringFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetAcquiringFeesOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcquiringFees) {
		return nil, false
	}
	return o.AcquiringFees, true
}

// HasAcquiringFees returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasAcquiringFees() bool {
	if o != nil && !common.IsNil(o.AcquiringFees) {
		return true
	}

	return false
}

// SetAcquiringFees gets a reference to the given string and assigns it to the AcquiringFees field.
func (o *UpdateSplitConfigurationLogicRequest) SetAcquiringFees(v string) {
	o.AcquiringFees = &v
}

// GetAdditionalCommission returns the AdditionalCommission field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetAdditionalCommission() AdditionalCommission {
	if o == nil || common.IsNil(o.AdditionalCommission) {
		var ret AdditionalCommission
		return ret
	}
	return *o.AdditionalCommission
}

// GetAdditionalCommissionOk returns a tuple with the AdditionalCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetAdditionalCommissionOk() (*AdditionalCommission, bool) {
	if o == nil || common.IsNil(o.AdditionalCommission) {
		return nil, false
	}
	return o.AdditionalCommission, true
}

// HasAdditionalCommission returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasAdditionalCommission() bool {
	if o != nil && !common.IsNil(o.AdditionalCommission) {
		return true
	}

	return false
}

// SetAdditionalCommission gets a reference to the given AdditionalCommission and assigns it to the AdditionalCommission field.
func (o *UpdateSplitConfigurationLogicRequest) SetAdditionalCommission(v AdditionalCommission) {
	o.AdditionalCommission = &v
}

// GetAdyenCommission returns the AdyenCommission field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetAdyenCommission() string {
	if o == nil || common.IsNil(o.AdyenCommission) {
		var ret string
		return ret
	}
	return *o.AdyenCommission
}

// GetAdyenCommissionOk returns a tuple with the AdyenCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetAdyenCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdyenCommission) {
		return nil, false
	}
	return o.AdyenCommission, true
}

// HasAdyenCommission returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasAdyenCommission() bool {
	if o != nil && !common.IsNil(o.AdyenCommission) {
		return true
	}

	return false
}

// SetAdyenCommission gets a reference to the given string and assigns it to the AdyenCommission field.
func (o *UpdateSplitConfigurationLogicRequest) SetAdyenCommission(v string) {
	o.AdyenCommission = &v
}

// GetAdyenFees returns the AdyenFees field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetAdyenFees() string {
	if o == nil || common.IsNil(o.AdyenFees) {
		var ret string
		return ret
	}
	return *o.AdyenFees
}

// GetAdyenFeesOk returns a tuple with the AdyenFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetAdyenFeesOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdyenFees) {
		return nil, false
	}
	return o.AdyenFees, true
}

// HasAdyenFees returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasAdyenFees() bool {
	if o != nil && !common.IsNil(o.AdyenFees) {
		return true
	}

	return false
}

// SetAdyenFees gets a reference to the given string and assigns it to the AdyenFees field.
func (o *UpdateSplitConfigurationLogicRequest) SetAdyenFees(v string) {
	o.AdyenFees = &v
}

// GetAdyenMarkup returns the AdyenMarkup field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetAdyenMarkup() string {
	if o == nil || common.IsNil(o.AdyenMarkup) {
		var ret string
		return ret
	}
	return *o.AdyenMarkup
}

// GetAdyenMarkupOk returns a tuple with the AdyenMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetAdyenMarkupOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdyenMarkup) {
		return nil, false
	}
	return o.AdyenMarkup, true
}

// HasAdyenMarkup returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasAdyenMarkup() bool {
	if o != nil && !common.IsNil(o.AdyenMarkup) {
		return true
	}

	return false
}

// SetAdyenMarkup gets a reference to the given string and assigns it to the AdyenMarkup field.
func (o *UpdateSplitConfigurationLogicRequest) SetAdyenMarkup(v string) {
	o.AdyenMarkup = &v
}

// GetChargeback returns the Chargeback field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetChargeback() string {
	if o == nil || common.IsNil(o.Chargeback) {
		var ret string
		return ret
	}
	return *o.Chargeback
}

// GetChargebackOk returns a tuple with the Chargeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetChargebackOk() (*string, bool) {
	if o == nil || common.IsNil(o.Chargeback) {
		return nil, false
	}
	return o.Chargeback, true
}

// HasChargeback returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasChargeback() bool {
	if o != nil && !common.IsNil(o.Chargeback) {
		return true
	}

	return false
}

// SetChargeback gets a reference to the given string and assigns it to the Chargeback field.
func (o *UpdateSplitConfigurationLogicRequest) SetChargeback(v string) {
	o.Chargeback = &v
}

// GetChargebackCostAllocation returns the ChargebackCostAllocation field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetChargebackCostAllocation() string {
	if o == nil || common.IsNil(o.ChargebackCostAllocation) {
		var ret string
		return ret
	}
	return *o.ChargebackCostAllocation
}

// GetChargebackCostAllocationOk returns a tuple with the ChargebackCostAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetChargebackCostAllocationOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChargebackCostAllocation) {
		return nil, false
	}
	return o.ChargebackCostAllocation, true
}

// HasChargebackCostAllocation returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasChargebackCostAllocation() bool {
	if o != nil && !common.IsNil(o.ChargebackCostAllocation) {
		return true
	}

	return false
}

// SetChargebackCostAllocation gets a reference to the given string and assigns it to the ChargebackCostAllocation field.
func (o *UpdateSplitConfigurationLogicRequest) SetChargebackCostAllocation(v string) {
	o.ChargebackCostAllocation = &v
}

// GetCommission returns the Commission field value
func (o *UpdateSplitConfigurationLogicRequest) GetCommission() Commission {
	if o == nil {
		var ret Commission
		return ret
	}

	return o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetCommissionOk() (*Commission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commission, true
}

// SetCommission sets field value
func (o *UpdateSplitConfigurationLogicRequest) SetCommission(v Commission) {
	o.Commission = v
}

// GetInterchange returns the Interchange field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetInterchange() string {
	if o == nil || common.IsNil(o.Interchange) {
		var ret string
		return ret
	}
	return *o.Interchange
}

// GetInterchangeOk returns a tuple with the Interchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetInterchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interchange) {
		return nil, false
	}
	return o.Interchange, true
}

// HasInterchange returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasInterchange() bool {
	if o != nil && !common.IsNil(o.Interchange) {
		return true
	}

	return false
}

// SetInterchange gets a reference to the given string and assigns it to the Interchange field.
func (o *UpdateSplitConfigurationLogicRequest) SetInterchange(v string) {
	o.Interchange = &v
}

// GetPaymentFee returns the PaymentFee field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetPaymentFee() string {
	if o == nil || common.IsNil(o.PaymentFee) {
		var ret string
		return ret
	}
	return *o.PaymentFee
}

// GetPaymentFeeOk returns a tuple with the PaymentFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetPaymentFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentFee) {
		return nil, false
	}
	return o.PaymentFee, true
}

// HasPaymentFee returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasPaymentFee() bool {
	if o != nil && !common.IsNil(o.PaymentFee) {
		return true
	}

	return false
}

// SetPaymentFee gets a reference to the given string and assigns it to the PaymentFee field.
func (o *UpdateSplitConfigurationLogicRequest) SetPaymentFee(v string) {
	o.PaymentFee = &v
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetRefund() string {
	if o == nil || common.IsNil(o.Refund) {
		var ret string
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetRefundOk() (*string, bool) {
	if o == nil || common.IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasRefund() bool {
	if o != nil && !common.IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given string and assigns it to the Refund field.
func (o *UpdateSplitConfigurationLogicRequest) SetRefund(v string) {
	o.Refund = &v
}

// GetRefundCostAllocation returns the RefundCostAllocation field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetRefundCostAllocation() string {
	if o == nil || common.IsNil(o.RefundCostAllocation) {
		var ret string
		return ret
	}
	return *o.RefundCostAllocation
}

// GetRefundCostAllocationOk returns a tuple with the RefundCostAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetRefundCostAllocationOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefundCostAllocation) {
		return nil, false
	}
	return o.RefundCostAllocation, true
}

// HasRefundCostAllocation returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasRefundCostAllocation() bool {
	if o != nil && !common.IsNil(o.RefundCostAllocation) {
		return true
	}

	return false
}

// SetRefundCostAllocation gets a reference to the given string and assigns it to the RefundCostAllocation field.
func (o *UpdateSplitConfigurationLogicRequest) SetRefundCostAllocation(v string) {
	o.RefundCostAllocation = &v
}

// GetRemainder returns the Remainder field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetRemainder() string {
	if o == nil || common.IsNil(o.Remainder) {
		var ret string
		return ret
	}
	return *o.Remainder
}

// GetRemainderOk returns a tuple with the Remainder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetRemainderOk() (*string, bool) {
	if o == nil || common.IsNil(o.Remainder) {
		return nil, false
	}
	return o.Remainder, true
}

// HasRemainder returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasRemainder() bool {
	if o != nil && !common.IsNil(o.Remainder) {
		return true
	}

	return false
}

// SetRemainder gets a reference to the given string and assigns it to the Remainder field.
func (o *UpdateSplitConfigurationLogicRequest) SetRemainder(v string) {
	o.Remainder = &v
}

// GetSchemeFee returns the SchemeFee field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetSchemeFee() string {
	if o == nil || common.IsNil(o.SchemeFee) {
		var ret string
		return ret
	}
	return *o.SchemeFee
}

// GetSchemeFeeOk returns a tuple with the SchemeFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetSchemeFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SchemeFee) {
		return nil, false
	}
	return o.SchemeFee, true
}

// HasSchemeFee returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasSchemeFee() bool {
	if o != nil && !common.IsNil(o.SchemeFee) {
		return true
	}

	return false
}

// SetSchemeFee gets a reference to the given string and assigns it to the SchemeFee field.
func (o *UpdateSplitConfigurationLogicRequest) SetSchemeFee(v string) {
	o.SchemeFee = &v
}

// GetSplitLogicId returns the SplitLogicId field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetSplitLogicId() string {
	if o == nil || common.IsNil(o.SplitLogicId) {
		var ret string
		return ret
	}
	return *o.SplitLogicId
}

// GetSplitLogicIdOk returns a tuple with the SplitLogicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetSplitLogicIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SplitLogicId) {
		return nil, false
	}
	return o.SplitLogicId, true
}

// HasSplitLogicId returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasSplitLogicId() bool {
	if o != nil && !common.IsNil(o.SplitLogicId) {
		return true
	}

	return false
}

// SetSplitLogicId gets a reference to the given string and assigns it to the SplitLogicId field.
func (o *UpdateSplitConfigurationLogicRequest) SetSplitLogicId(v string) {
	o.SplitLogicId = &v
}

// GetSurcharge returns the Surcharge field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetSurcharge() string {
	if o == nil || common.IsNil(o.Surcharge) {
		var ret string
		return ret
	}
	return *o.Surcharge
}

// GetSurchargeOk returns a tuple with the Surcharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetSurchargeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Surcharge) {
		return nil, false
	}
	return o.Surcharge, true
}

// HasSurcharge returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasSurcharge() bool {
	if o != nil && !common.IsNil(o.Surcharge) {
		return true
	}

	return false
}

// SetSurcharge gets a reference to the given string and assigns it to the Surcharge field.
func (o *UpdateSplitConfigurationLogicRequest) SetSurcharge(v string) {
	o.Surcharge = &v
}

// GetTip returns the Tip field value if set, zero value otherwise.
func (o *UpdateSplitConfigurationLogicRequest) GetTip() string {
	if o == nil || common.IsNil(o.Tip) {
		var ret string
		return ret
	}
	return *o.Tip
}

// GetTipOk returns a tuple with the Tip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSplitConfigurationLogicRequest) GetTipOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tip) {
		return nil, false
	}
	return o.Tip, true
}

// HasTip returns a boolean if a field has been set.
func (o *UpdateSplitConfigurationLogicRequest) HasTip() bool {
	if o != nil && !common.IsNil(o.Tip) {
		return true
	}

	return false
}

// SetTip gets a reference to the given string and assigns it to the Tip field.
func (o *UpdateSplitConfigurationLogicRequest) SetTip(v string) {
	o.Tip = &v
}

func (o UpdateSplitConfigurationLogicRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSplitConfigurationLogicRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcquiringFees) {
		toSerialize["acquiringFees"] = o.AcquiringFees
	}
	if !common.IsNil(o.AdditionalCommission) {
		toSerialize["additionalCommission"] = o.AdditionalCommission
	}
	if !common.IsNil(o.AdyenCommission) {
		toSerialize["adyenCommission"] = o.AdyenCommission
	}
	if !common.IsNil(o.AdyenFees) {
		toSerialize["adyenFees"] = o.AdyenFees
	}
	if !common.IsNil(o.AdyenMarkup) {
		toSerialize["adyenMarkup"] = o.AdyenMarkup
	}
	if !common.IsNil(o.Chargeback) {
		toSerialize["chargeback"] = o.Chargeback
	}
	if !common.IsNil(o.ChargebackCostAllocation) {
		toSerialize["chargebackCostAllocation"] = o.ChargebackCostAllocation
	}
	toSerialize["commission"] = o.Commission
	if !common.IsNil(o.Interchange) {
		toSerialize["interchange"] = o.Interchange
	}
	if !common.IsNil(o.PaymentFee) {
		toSerialize["paymentFee"] = o.PaymentFee
	}
	if !common.IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	if !common.IsNil(o.RefundCostAllocation) {
		toSerialize["refundCostAllocation"] = o.RefundCostAllocation
	}
	if !common.IsNil(o.Remainder) {
		toSerialize["remainder"] = o.Remainder
	}
	if !common.IsNil(o.SchemeFee) {
		toSerialize["schemeFee"] = o.SchemeFee
	}
	if !common.IsNil(o.SplitLogicId) {
		toSerialize["splitLogicId"] = o.SplitLogicId
	}
	if !common.IsNil(o.Surcharge) {
		toSerialize["surcharge"] = o.Surcharge
	}
	if !common.IsNil(o.Tip) {
		toSerialize["tip"] = o.Tip
	}
	return toSerialize, nil
}

type NullableUpdateSplitConfigurationLogicRequest struct {
	value *UpdateSplitConfigurationLogicRequest
	isSet bool
}

func (v NullableUpdateSplitConfigurationLogicRequest) Get() *UpdateSplitConfigurationLogicRequest {
	return v.value
}

func (v *NullableUpdateSplitConfigurationLogicRequest) Set(val *UpdateSplitConfigurationLogicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSplitConfigurationLogicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSplitConfigurationLogicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSplitConfigurationLogicRequest(val *UpdateSplitConfigurationLogicRequest) *NullableUpdateSplitConfigurationLogicRequest {
	return &NullableUpdateSplitConfigurationLogicRequest{value: val, isSet: true}
}

func (v NullableUpdateSplitConfigurationLogicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSplitConfigurationLogicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpdateSplitConfigurationLogicRequest) isValidAcquiringFees() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAcquiringFees() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidAdyenCommission() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAdyenCommission() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidAdyenFees() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAdyenFees() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidAdyenMarkup() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetAdyenMarkup() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidChargeback() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount", "deductAccordingToSplitRatio"}
	for _, allowed := range allowedEnumValues {
		if o.GetChargeback() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidChargebackCostAllocation() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetChargebackCostAllocation() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidInterchange() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetInterchange() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidPaymentFee() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetPaymentFee() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidRefund() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount", "deductAccordingToSplitRatio"}
	for _, allowed := range allowedEnumValues {
		if o.GetRefund() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidRefundCostAllocation() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetRefundCostAllocation() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidRemainder() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetRemainder() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidSchemeFee() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetSchemeFee() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidSurcharge() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetSurcharge() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateSplitConfigurationLogicRequest) isValidTip() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetTip() == allowed {
			return true
		}
	}
	return false
}
