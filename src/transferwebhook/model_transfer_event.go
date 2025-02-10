/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the TransferEvent type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferEvent{}

// TransferEvent struct for TransferEvent
type TransferEvent struct {
	Amount *Amount `json:"amount,omitempty"`
	// The amount adjustments in this transfer.
	AmountAdjustments []AmountAdjustment `json:"amountAdjustments,omitempty"`
	// Scheme unique arn identifier useful for tracing captures, chargebacks, refunds, etc.
	Arn *string `json:"arn,omitempty"`
	// The date when the transfer request was sent.
	BookingDate *time.Time `json:"bookingDate,omitempty"`
	// The estimated time when the beneficiary should have access to the funds.
	EstimatedArrivalTime *time.Time `json:"estimatedArrivalTime,omitempty"`
	// A list of event data.
	EventsData     []TransferEventEventsDataInner `json:"eventsData,omitempty"`
	ExternalReason *ExternalReason                `json:"externalReason,omitempty"`
	// The unique identifier of the transfer event.
	Id           *string       `json:"id,omitempty"`
	Modification *Modification `json:"modification,omitempty"`
	// The list of balance mutations per event.
	Mutations      []BalanceMutation `json:"mutations,omitempty"`
	OriginalAmount *Amount           `json:"originalAmount,omitempty"`
	// The reason for the transfer status.
	Reason *string `json:"reason,omitempty"`
	// The status of the transfer event.
	Status       *string                    `json:"status,omitempty"`
	TrackingData *TransferEventTrackingData `json:"trackingData,omitempty"`
	// The id of the transaction that is related to this accounting event. Only sent for events of type **accounting** where the balance changes.
	TransactionId *string `json:"transactionId,omitempty"`
	// The type of the transfer event. Possible values: **accounting**, **tracking**.
	Type *string `json:"type,omitempty"`
	// The date when the tracking status was updated.
	UpdateDate *time.Time `json:"updateDate,omitempty"`
	// The date when the funds are expected to be deducted from or credited to the balance account. This date can be in either the past or future.
	ValueDate *time.Time `json:"valueDate,omitempty"`
}

// NewTransferEvent instantiates a new TransferEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEvent() *TransferEvent {
	this := TransferEvent{}
	return &this
}

// NewTransferEventWithDefaults instantiates a new TransferEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventWithDefaults() *TransferEvent {
	this := TransferEvent{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransferEvent) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransferEvent) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *TransferEvent) SetAmount(v Amount) {
	o.Amount = &v
}

// GetAmountAdjustments returns the AmountAdjustments field value if set, zero value otherwise.
func (o *TransferEvent) GetAmountAdjustments() []AmountAdjustment {
	if o == nil || common.IsNil(o.AmountAdjustments) {
		var ret []AmountAdjustment
		return ret
	}
	return o.AmountAdjustments
}

// GetAmountAdjustmentsOk returns a tuple with the AmountAdjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetAmountAdjustmentsOk() ([]AmountAdjustment, bool) {
	if o == nil || common.IsNil(o.AmountAdjustments) {
		return nil, false
	}
	return o.AmountAdjustments, true
}

// HasAmountAdjustments returns a boolean if a field has been set.
func (o *TransferEvent) HasAmountAdjustments() bool {
	if o != nil && !common.IsNil(o.AmountAdjustments) {
		return true
	}

	return false
}

// SetAmountAdjustments gets a reference to the given []AmountAdjustment and assigns it to the AmountAdjustments field.
func (o *TransferEvent) SetAmountAdjustments(v []AmountAdjustment) {
	o.AmountAdjustments = v
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *TransferEvent) GetArn() string {
	if o == nil || common.IsNil(o.Arn) {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetArnOk() (*string, bool) {
	if o == nil || common.IsNil(o.Arn) {
		return nil, false
	}
	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *TransferEvent) HasArn() bool {
	if o != nil && !common.IsNil(o.Arn) {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *TransferEvent) SetArn(v string) {
	o.Arn = &v
}

// GetBookingDate returns the BookingDate field value if set, zero value otherwise.
func (o *TransferEvent) GetBookingDate() time.Time {
	if o == nil || common.IsNil(o.BookingDate) {
		var ret time.Time
		return ret
	}
	return *o.BookingDate
}

// GetBookingDateOk returns a tuple with the BookingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetBookingDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.BookingDate) {
		return nil, false
	}
	return o.BookingDate, true
}

// HasBookingDate returns a boolean if a field has been set.
func (o *TransferEvent) HasBookingDate() bool {
	if o != nil && !common.IsNil(o.BookingDate) {
		return true
	}

	return false
}

// SetBookingDate gets a reference to the given time.Time and assigns it to the BookingDate field.
func (o *TransferEvent) SetBookingDate(v time.Time) {
	o.BookingDate = &v
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise.
func (o *TransferEvent) GetEstimatedArrivalTime() time.Time {
	if o == nil || common.IsNil(o.EstimatedArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetEstimatedArrivalTimeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.EstimatedArrivalTime) {
		return nil, false
	}
	return o.EstimatedArrivalTime, true
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *TransferEvent) HasEstimatedArrivalTime() bool {
	if o != nil && !common.IsNil(o.EstimatedArrivalTime) {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given time.Time and assigns it to the EstimatedArrivalTime field.
func (o *TransferEvent) SetEstimatedArrivalTime(v time.Time) {
	o.EstimatedArrivalTime = &v
}

// GetEventsData returns the EventsData field value if set, zero value otherwise.
func (o *TransferEvent) GetEventsData() []TransferEventEventsDataInner {
	if o == nil || common.IsNil(o.EventsData) {
		var ret []TransferEventEventsDataInner
		return ret
	}
	return o.EventsData
}

// GetEventsDataOk returns a tuple with the EventsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetEventsDataOk() ([]TransferEventEventsDataInner, bool) {
	if o == nil || common.IsNil(o.EventsData) {
		return nil, false
	}
	return o.EventsData, true
}

// HasEventsData returns a boolean if a field has been set.
func (o *TransferEvent) HasEventsData() bool {
	if o != nil && !common.IsNil(o.EventsData) {
		return true
	}

	return false
}

// SetEventsData gets a reference to the given []TransferEventEventsDataInner and assigns it to the EventsData field.
func (o *TransferEvent) SetEventsData(v []TransferEventEventsDataInner) {
	o.EventsData = v
}

// GetExternalReason returns the ExternalReason field value if set, zero value otherwise.
func (o *TransferEvent) GetExternalReason() ExternalReason {
	if o == nil || common.IsNil(o.ExternalReason) {
		var ret ExternalReason
		return ret
	}
	return *o.ExternalReason
}

// GetExternalReasonOk returns a tuple with the ExternalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetExternalReasonOk() (*ExternalReason, bool) {
	if o == nil || common.IsNil(o.ExternalReason) {
		return nil, false
	}
	return o.ExternalReason, true
}

// HasExternalReason returns a boolean if a field has been set.
func (o *TransferEvent) HasExternalReason() bool {
	if o != nil && !common.IsNil(o.ExternalReason) {
		return true
	}

	return false
}

// SetExternalReason gets a reference to the given ExternalReason and assigns it to the ExternalReason field.
func (o *TransferEvent) SetExternalReason(v ExternalReason) {
	o.ExternalReason = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransferEvent) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransferEvent) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransferEvent) SetId(v string) {
	o.Id = &v
}

// GetModification returns the Modification field value if set, zero value otherwise.
func (o *TransferEvent) GetModification() Modification {
	if o == nil || common.IsNil(o.Modification) {
		var ret Modification
		return ret
	}
	return *o.Modification
}

// GetModificationOk returns a tuple with the Modification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetModificationOk() (*Modification, bool) {
	if o == nil || common.IsNil(o.Modification) {
		return nil, false
	}
	return o.Modification, true
}

// HasModification returns a boolean if a field has been set.
func (o *TransferEvent) HasModification() bool {
	if o != nil && !common.IsNil(o.Modification) {
		return true
	}

	return false
}

// SetModification gets a reference to the given Modification and assigns it to the Modification field.
func (o *TransferEvent) SetModification(v Modification) {
	o.Modification = &v
}

// GetMutations returns the Mutations field value if set, zero value otherwise.
func (o *TransferEvent) GetMutations() []BalanceMutation {
	if o == nil || common.IsNil(o.Mutations) {
		var ret []BalanceMutation
		return ret
	}
	return o.Mutations
}

// GetMutationsOk returns a tuple with the Mutations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetMutationsOk() ([]BalanceMutation, bool) {
	if o == nil || common.IsNil(o.Mutations) {
		return nil, false
	}
	return o.Mutations, true
}

// HasMutations returns a boolean if a field has been set.
func (o *TransferEvent) HasMutations() bool {
	if o != nil && !common.IsNil(o.Mutations) {
		return true
	}

	return false
}

// SetMutations gets a reference to the given []BalanceMutation and assigns it to the Mutations field.
func (o *TransferEvent) SetMutations(v []BalanceMutation) {
	o.Mutations = v
}

// GetOriginalAmount returns the OriginalAmount field value if set, zero value otherwise.
func (o *TransferEvent) GetOriginalAmount() Amount {
	if o == nil || common.IsNil(o.OriginalAmount) {
		var ret Amount
		return ret
	}
	return *o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetOriginalAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.OriginalAmount) {
		return nil, false
	}
	return o.OriginalAmount, true
}

// HasOriginalAmount returns a boolean if a field has been set.
func (o *TransferEvent) HasOriginalAmount() bool {
	if o != nil && !common.IsNil(o.OriginalAmount) {
		return true
	}

	return false
}

// SetOriginalAmount gets a reference to the given Amount and assigns it to the OriginalAmount field.
func (o *TransferEvent) SetOriginalAmount(v Amount) {
	o.OriginalAmount = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TransferEvent) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TransferEvent) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TransferEvent) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransferEvent) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransferEvent) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransferEvent) SetStatus(v string) {
	o.Status = &v
}

// GetTrackingData returns the TrackingData field value if set, zero value otherwise.
func (o *TransferEvent) GetTrackingData() TransferEventTrackingData {
	if o == nil || common.IsNil(o.TrackingData) {
		var ret TransferEventTrackingData
		return ret
	}
	return *o.TrackingData
}

// GetTrackingDataOk returns a tuple with the TrackingData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTrackingDataOk() (*TransferEventTrackingData, bool) {
	if o == nil || common.IsNil(o.TrackingData) {
		return nil, false
	}
	return o.TrackingData, true
}

// HasTrackingData returns a boolean if a field has been set.
func (o *TransferEvent) HasTrackingData() bool {
	if o != nil && !common.IsNil(o.TrackingData) {
		return true
	}

	return false
}

// SetTrackingData gets a reference to the given TransferEventTrackingData and assigns it to the TrackingData field.
func (o *TransferEvent) SetTrackingData(v TransferEventTrackingData) {
	o.TrackingData = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TransferEvent) GetTransactionId() string {
	if o == nil || common.IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTransactionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TransferEvent) HasTransactionId() bool {
	if o != nil && !common.IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *TransferEvent) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferEvent) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferEvent) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransferEvent) SetType(v string) {
	o.Type = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *TransferEvent) GetUpdateDate() time.Time {
	if o == nil || common.IsNil(o.UpdateDate) {
		var ret time.Time
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetUpdateDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.UpdateDate) {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *TransferEvent) HasUpdateDate() bool {
	if o != nil && !common.IsNil(o.UpdateDate) {
		return true
	}

	return false
}

// SetUpdateDate gets a reference to the given time.Time and assigns it to the UpdateDate field.
func (o *TransferEvent) SetUpdateDate(v time.Time) {
	o.UpdateDate = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *TransferEvent) GetValueDate() time.Time {
	if o == nil || common.IsNil(o.ValueDate) {
		var ret time.Time
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetValueDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ValueDate) {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *TransferEvent) HasValueDate() bool {
	if o != nil && !common.IsNil(o.ValueDate) {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given time.Time and assigns it to the ValueDate field.
func (o *TransferEvent) SetValueDate(v time.Time) {
	o.ValueDate = &v
}

func (o TransferEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.AmountAdjustments) {
		toSerialize["amountAdjustments"] = o.AmountAdjustments
	}
	if !common.IsNil(o.Arn) {
		toSerialize["arn"] = o.Arn
	}
	if !common.IsNil(o.BookingDate) {
		toSerialize["bookingDate"] = o.BookingDate
	}
	if !common.IsNil(o.EstimatedArrivalTime) {
		toSerialize["estimatedArrivalTime"] = o.EstimatedArrivalTime
	}
	if !common.IsNil(o.EventsData) {
		toSerialize["eventsData"] = o.EventsData
	}
	if !common.IsNil(o.ExternalReason) {
		toSerialize["externalReason"] = o.ExternalReason
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Modification) {
		toSerialize["modification"] = o.Modification
	}
	if !common.IsNil(o.Mutations) {
		toSerialize["mutations"] = o.Mutations
	}
	if !common.IsNil(o.OriginalAmount) {
		toSerialize["originalAmount"] = o.OriginalAmount
	}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TrackingData) {
		toSerialize["trackingData"] = o.TrackingData
	}
	if !common.IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UpdateDate) {
		toSerialize["updateDate"] = o.UpdateDate
	}
	if !common.IsNil(o.ValueDate) {
		toSerialize["valueDate"] = o.ValueDate
	}
	return toSerialize, nil
}

type NullableTransferEvent struct {
	value *TransferEvent
	isSet bool
}

func (v NullableTransferEvent) Get() *TransferEvent {
	return v.value
}

func (v *NullableTransferEvent) Set(val *TransferEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEvent(val *TransferEvent) *NullableTransferEvent {
	return &NullableTransferEvent{value: val, isSet: true}
}

func (v NullableTransferEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferEvent) isValidReason() bool {
	var allowedEnumValues = []string{"accountHierarchyNotActive", "amountLimitExceeded", "approved", "balanceAccountTemporarilyBlockedByTransactionRule", "counterpartyAccountBlocked", "counterpartyAccountClosed", "counterpartyAccountNotFound", "counterpartyAddressRequired", "counterpartyBankTimedOut", "counterpartyBankUnavailable", "declined", "declinedByTransactionRule", "directDebitNotSupported", "error", "notEnoughBalance", "pendingApproval", "pendingExecution", "refusedByCounterpartyBank", "refusedByCustomer", "routeNotFound", "scaFailed", "transferInstrumentDoesNotExist", "unknown"}
	for _, allowed := range allowedEnumValues {
		if o.GetReason() == allowed {
			return true
		}
	}
	return false
}
func (o *TransferEvent) isValidStatus() bool {
	var allowedEnumValues = []string{"approvalPending", "atmWithdrawal", "atmWithdrawalReversalPending", "atmWithdrawalReversed", "authAdjustmentAuthorised", "authAdjustmentError", "authAdjustmentRefused", "authorised", "bankTransfer", "bankTransferPending", "booked", "bookingPending", "cancelled", "capturePending", "captureReversalPending", "captureReversed", "captured", "capturedExternally", "chargeback", "chargebackExternally", "chargebackPending", "chargebackReversalPending", "chargebackReversed", "credited", "depositCorrection", "depositCorrectionPending", "dispute", "disputeClosed", "disputeExpired", "disputeNeedsReview", "error", "expired", "failed", "fee", "feePending", "internalTransfer", "internalTransferPending", "invoiceDeduction", "invoiceDeductionPending", "manualCorrectionPending", "manuallyCorrected", "matchedStatement", "matchedStatementPending", "merchantPayin", "merchantPayinPending", "merchantPayinReversed", "merchantPayinReversedPending", "miscCost", "miscCostPending", "paymentCost", "paymentCostPending", "pendingApproval", "pendingExecution", "received", "refundPending", "refundReversalPending", "refundReversed", "refunded", "refundedExternally", "refused", "rejected", "reserveAdjustment", "reserveAdjustmentPending", "returned", "secondChargeback", "secondChargebackPending", "undefined"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *TransferEvent) isValidType() bool {
	var allowedEnumValues = []string{"accounting", "tracking"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
