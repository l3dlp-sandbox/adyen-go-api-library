/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
	"time"
)

// checks if the TransferData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferData{}

// TransferData struct for TransferData
type TransferData struct {
	AccountHolder *ResourceReference `json:"accountHolder,omitempty"`
	Amount Amount `json:"amount"`
	BalanceAccount *ResourceReference `json:"balanceAccount,omitempty"`
	// The unique identifier of the balance platform.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
	// The list of the latest balance statuses in the transfer.
	Balances []BalanceMutation `json:"balances,omitempty"`
	// The category of the transfer.  Possible values:   - **bank**: a transfer involving a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.  - **card**: a transfer involving a third-party card.  - **internal**: a transfer between [balance accounts](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  - **issuedCard**: a transfer initiated by an Adyen-issued card.  - **platformPayment**: funds movements related to payments that are acquired for your users.
	Category string `json:"category"`
	CategoryData *TransferDataCategoryData `json:"categoryData,omitempty"`
	Counterparty *TransferNotificationCounterParty `json:"counterparty,omitempty"`
	// The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Your description for the transfer. It is used by most banks as the transfer description. We recommend sending a maximum of 140 characters, otherwise the description may be truncated.  Supported characters: **[a-z] [A-Z] [0-9] / - ?** **: ( ) . , ' + Space**  Supported characters for **regular** and **fast** transfers to a US counterparty: **[a-z] [A-Z] [0-9] & $ % # @** **~ = + - _ ' \" ! ?**
	Description *string `json:"description,omitempty"`
	DirectDebitInformation *DirectDebitInformation `json:"directDebitInformation,omitempty"`
	// The direction of the transfer.  Possible values: **incoming**, **outgoing**.
	Direction *string `json:"direction,omitempty"`
	// The list of events leading up to the current status of the transfer.
	Events []TransferEvent `json:"events,omitempty"`
	// The ID of the resource.
	Id *string `json:"id,omitempty"`
	PaymentInstrument *PaymentInstrument `json:"paymentInstrument,omitempty"`
	// Additional information about the status of the transfer.
	Reason *string `json:"reason,omitempty"`
	// Your reference for the transfer, used internally within your platform. If you don't provide this in the request, Adyen generates a unique reference.
	Reference *string `json:"reference,omitempty"`
	//  A reference that is sent to the recipient. This reference is also sent in all webhooks related to the transfer, so you can use it to track statuses for both the source and recipient of funds.   Supported characters: **a-z**, **A-Z**, **0-9**.The maximum length depends on the `category`.   - **internal**: 80 characters  - **bank**: 35 characters when transferring to an IBAN, 15 characters for others.
	ReferenceForBeneficiary *string `json:"referenceForBeneficiary,omitempty"`
	Review *TransferReview `json:"review,omitempty"`
	// The sequence number of the transfer webhook. The numbers start from 1 and increase with each new webhook for a specific transfer.  The sequence number can help you restore the correct sequence of events even if they arrive out of order.
	SequenceNumber *int32 `json:"sequenceNumber,omitempty"`
	// The result of the transfer.   For example, **authorised**, **refused**, or **error**.
	Status string `json:"status"`
	Tracking *TransferDataTracking `json:"tracking,omitempty"`
	TransactionRulesResult *TransactionRulesResult `json:"transactionRulesResult,omitempty"`
	// The type of transfer or transaction. For example, **refund**, **payment**, **internalTransfer**, **bankTransfer**.
	Type *string `json:"type,omitempty"`
}

// NewTransferData instantiates a new TransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferData(amount Amount, category string, status string) *TransferData {
	this := TransferData{}
	this.Amount = amount
	this.Category = category
	this.Status = status
	return &this
}

// NewTransferDataWithDefaults instantiates a new TransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDataWithDefaults() *TransferData {
	this := TransferData{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *TransferData) GetAccountHolder() ResourceReference {
	if o == nil || common.IsNil(o.AccountHolder) {
		var ret ResourceReference
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetAccountHolderOk() (*ResourceReference, bool) {
	if o == nil || common.IsNil(o.AccountHolder) {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *TransferData) HasAccountHolder() bool {
	if o != nil && !common.IsNil(o.AccountHolder) {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given ResourceReference and assigns it to the AccountHolder field.
func (o *TransferData) SetAccountHolder(v ResourceReference) {
	o.AccountHolder = &v
}

// GetAmount returns the Amount field value
func (o *TransferData) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferData) SetAmount(v Amount) {
	o.Amount = v
}

// GetBalanceAccount returns the BalanceAccount field value if set, zero value otherwise.
func (o *TransferData) GetBalanceAccount() ResourceReference {
	if o == nil || common.IsNil(o.BalanceAccount) {
		var ret ResourceReference
		return ret
	}
	return *o.BalanceAccount
}

// GetBalanceAccountOk returns a tuple with the BalanceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetBalanceAccountOk() (*ResourceReference, bool) {
	if o == nil || common.IsNil(o.BalanceAccount) {
		return nil, false
	}
	return o.BalanceAccount, true
}

// HasBalanceAccount returns a boolean if a field has been set.
func (o *TransferData) HasBalanceAccount() bool {
	if o != nil && !common.IsNil(o.BalanceAccount) {
		return true
	}

	return false
}

// SetBalanceAccount gets a reference to the given ResourceReference and assigns it to the BalanceAccount field.
func (o *TransferData) SetBalanceAccount(v ResourceReference) {
	o.BalanceAccount = &v
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *TransferData) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *TransferData) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *TransferData) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *TransferData) GetBalances() []BalanceMutation {
	if o == nil || common.IsNil(o.Balances) {
		var ret []BalanceMutation
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetBalancesOk() ([]BalanceMutation, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *TransferData) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []BalanceMutation and assigns it to the Balances field.
func (o *TransferData) SetBalances(v []BalanceMutation) {
	o.Balances = v
}

// GetCategory returns the Category field value
func (o *TransferData) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *TransferData) SetCategory(v string) {
	o.Category = v
}

// GetCategoryData returns the CategoryData field value if set, zero value otherwise.
func (o *TransferData) GetCategoryData() TransferDataCategoryData {
	if o == nil || common.IsNil(o.CategoryData) {
		var ret TransferDataCategoryData
		return ret
	}
	return *o.CategoryData
}

// GetCategoryDataOk returns a tuple with the CategoryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetCategoryDataOk() (*TransferDataCategoryData, bool) {
	if o == nil || common.IsNil(o.CategoryData) {
		return nil, false
	}
	return o.CategoryData, true
}

// HasCategoryData returns a boolean if a field has been set.
func (o *TransferData) HasCategoryData() bool {
	if o != nil && !common.IsNil(o.CategoryData) {
		return true
	}

	return false
}

// SetCategoryData gets a reference to the given TransferDataCategoryData and assigns it to the CategoryData field.
func (o *TransferData) SetCategoryData(v TransferDataCategoryData) {
	o.CategoryData = &v
}

// GetCounterparty returns the Counterparty field value if set, zero value otherwise.
func (o *TransferData) GetCounterparty() TransferNotificationCounterParty {
	if o == nil || common.IsNil(o.Counterparty) {
		var ret TransferNotificationCounterParty
		return ret
	}
	return *o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetCounterpartyOk() (*TransferNotificationCounterParty, bool) {
	if o == nil || common.IsNil(o.Counterparty) {
		return nil, false
	}
	return o.Counterparty, true
}

// HasCounterparty returns a boolean if a field has been set.
func (o *TransferData) HasCounterparty() bool {
	if o != nil && !common.IsNil(o.Counterparty) {
		return true
	}

	return false
}

// SetCounterparty gets a reference to the given TransferNotificationCounterParty and assigns it to the Counterparty field.
func (o *TransferData) SetCounterparty(v TransferNotificationCounterParty) {
	o.Counterparty = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *TransferData) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *TransferData) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *TransferData) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransferData) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransferData) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransferData) SetDescription(v string) {
	o.Description = &v
}

// GetDirectDebitInformation returns the DirectDebitInformation field value if set, zero value otherwise.
func (o *TransferData) GetDirectDebitInformation() DirectDebitInformation {
	if o == nil || common.IsNil(o.DirectDebitInformation) {
		var ret DirectDebitInformation
		return ret
	}
	return *o.DirectDebitInformation
}

// GetDirectDebitInformationOk returns a tuple with the DirectDebitInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetDirectDebitInformationOk() (*DirectDebitInformation, bool) {
	if o == nil || common.IsNil(o.DirectDebitInformation) {
		return nil, false
	}
	return o.DirectDebitInformation, true
}

// HasDirectDebitInformation returns a boolean if a field has been set.
func (o *TransferData) HasDirectDebitInformation() bool {
	if o != nil && !common.IsNil(o.DirectDebitInformation) {
		return true
	}

	return false
}

// SetDirectDebitInformation gets a reference to the given DirectDebitInformation and assigns it to the DirectDebitInformation field.
func (o *TransferData) SetDirectDebitInformation(v DirectDebitInformation) {
	o.DirectDebitInformation = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *TransferData) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *TransferData) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *TransferData) SetDirection(v string) {
	o.Direction = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *TransferData) GetEvents() []TransferEvent {
	if o == nil || common.IsNil(o.Events) {
		var ret []TransferEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetEventsOk() ([]TransferEvent, bool) {
	if o == nil || common.IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *TransferData) HasEvents() bool {
	if o != nil && !common.IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []TransferEvent and assigns it to the Events field.
func (o *TransferData) SetEvents(v []TransferEvent) {
	o.Events = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransferData) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransferData) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransferData) SetId(v string) {
	o.Id = &v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise.
func (o *TransferData) GetPaymentInstrument() PaymentInstrument {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		var ret PaymentInstrument
		return ret
	}
	return *o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetPaymentInstrumentOk() (*PaymentInstrument, bool) {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		return nil, false
	}
	return o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *TransferData) HasPaymentInstrument() bool {
	if o != nil && !common.IsNil(o.PaymentInstrument) {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given PaymentInstrument and assigns it to the PaymentInstrument field.
func (o *TransferData) SetPaymentInstrument(v PaymentInstrument) {
	o.PaymentInstrument = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TransferData) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TransferData) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TransferData) SetReason(v string) {
	o.Reason = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TransferData) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TransferData) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TransferData) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceForBeneficiary returns the ReferenceForBeneficiary field value if set, zero value otherwise.
func (o *TransferData) GetReferenceForBeneficiary() string {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		var ret string
		return ret
	}
	return *o.ReferenceForBeneficiary
}

// GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetReferenceForBeneficiaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		return nil, false
	}
	return o.ReferenceForBeneficiary, true
}

// HasReferenceForBeneficiary returns a boolean if a field has been set.
func (o *TransferData) HasReferenceForBeneficiary() bool {
	if o != nil && !common.IsNil(o.ReferenceForBeneficiary) {
		return true
	}

	return false
}

// SetReferenceForBeneficiary gets a reference to the given string and assigns it to the ReferenceForBeneficiary field.
func (o *TransferData) SetReferenceForBeneficiary(v string) {
	o.ReferenceForBeneficiary = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *TransferData) GetReview() TransferReview {
	if o == nil || common.IsNil(o.Review) {
		var ret TransferReview
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetReviewOk() (*TransferReview, bool) {
	if o == nil || common.IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *TransferData) HasReview() bool {
	if o != nil && !common.IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given TransferReview and assigns it to the Review field.
func (o *TransferData) SetReview(v TransferReview) {
	o.Review = &v
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *TransferData) GetSequenceNumber() int32 {
	if o == nil || common.IsNil(o.SequenceNumber) {
		var ret int32
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetSequenceNumberOk() (*int32, bool) {
	if o == nil || common.IsNil(o.SequenceNumber) {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *TransferData) HasSequenceNumber() bool {
	if o != nil && !common.IsNil(o.SequenceNumber) {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int32 and assigns it to the SequenceNumber field.
func (o *TransferData) SetSequenceNumber(v int32) {
	o.SequenceNumber = &v
}

// GetStatus returns the Status field value
func (o *TransferData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferData) SetStatus(v string) {
	o.Status = v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *TransferData) GetTracking() TransferDataTracking {
	if o == nil || common.IsNil(o.Tracking) {
		var ret TransferDataTracking
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetTrackingOk() (*TransferDataTracking, bool) {
	if o == nil || common.IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *TransferData) HasTracking() bool {
	if o != nil && !common.IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given TransferDataTracking and assigns it to the Tracking field.
func (o *TransferData) SetTracking(v TransferDataTracking) {
	o.Tracking = &v
}

// GetTransactionRulesResult returns the TransactionRulesResult field value if set, zero value otherwise.
func (o *TransferData) GetTransactionRulesResult() TransactionRulesResult {
	if o == nil || common.IsNil(o.TransactionRulesResult) {
		var ret TransactionRulesResult
		return ret
	}
	return *o.TransactionRulesResult
}

// GetTransactionRulesResultOk returns a tuple with the TransactionRulesResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetTransactionRulesResultOk() (*TransactionRulesResult, bool) {
	if o == nil || common.IsNil(o.TransactionRulesResult) {
		return nil, false
	}
	return o.TransactionRulesResult, true
}

// HasTransactionRulesResult returns a boolean if a field has been set.
func (o *TransferData) HasTransactionRulesResult() bool {
	if o != nil && !common.IsNil(o.TransactionRulesResult) {
		return true
	}

	return false
}

// SetTransactionRulesResult gets a reference to the given TransactionRulesResult and assigns it to the TransactionRulesResult field.
func (o *TransferData) SetTransactionRulesResult(v TransactionRulesResult) {
	o.TransactionRulesResult = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferData) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferData) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransferData) SetType(v string) {
	o.Type = &v
}

func (o TransferData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountHolder) {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.BalanceAccount) {
		toSerialize["balanceAccount"] = o.BalanceAccount
	}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	toSerialize["category"] = o.Category
	if !common.IsNil(o.CategoryData) {
		toSerialize["categoryData"] = o.CategoryData
	}
	if !common.IsNil(o.Counterparty) {
		toSerialize["counterparty"] = o.Counterparty
	}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.DirectDebitInformation) {
		toSerialize["directDebitInformation"] = o.DirectDebitInformation
	}
	if !common.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !common.IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.PaymentInstrument) {
		toSerialize["paymentInstrument"] = o.PaymentInstrument
	}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.ReferenceForBeneficiary) {
		toSerialize["referenceForBeneficiary"] = o.ReferenceForBeneficiary
	}
	if !common.IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	if !common.IsNil(o.SequenceNumber) {
		toSerialize["sequenceNumber"] = o.SequenceNumber
	}
	toSerialize["status"] = o.Status
	if !common.IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !common.IsNil(o.TransactionRulesResult) {
		toSerialize["transactionRulesResult"] = o.TransactionRulesResult
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTransferData struct {
	value *TransferData
	isSet bool
}

func (v NullableTransferData) Get() *TransferData {
	return v.value
}

func (v *NullableTransferData) Set(val *TransferData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferData(val *TransferData) *NullableTransferData {
	return &NullableTransferData{value: val, isSet: true}
}

func (v NullableTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *TransferData) isValidCategory() bool {
    var allowedEnumValues = []string{ "bank", "card", "internal", "issuedCard", "platformPayment" }
    for _, allowed := range allowedEnumValues {
        if o.GetCategory() == allowed {
            return true
        }
    }
    return false
}
func (o *TransferData) isValidDirection() bool {
    var allowedEnumValues = []string{ "incoming", "outgoing" }
    for _, allowed := range allowedEnumValues {
        if o.GetDirection() == allowed {
            return true
        }
    }
    return false
}
func (o *TransferData) isValidReason() bool {
    var allowedEnumValues = []string{ "accountHierarchyNotActive", "amountLimitExceeded", "approved", "balanceAccountTemporarilyBlockedByTransactionRule", "counterpartyAccountBlocked", "counterpartyAccountClosed", "counterpartyAccountNotFound", "counterpartyAddressRequired", "counterpartyBankTimedOut", "counterpartyBankUnavailable", "declined", "declinedByTransactionRule", "directDebitNotSupported", "error", "notEnoughBalance", "pendingApproval", "pendingExecution", "refusedByCounterpartyBank", "refusedByCustomer", "routeNotFound", "scaFailed", "transferInstrumentDoesNotExist", "unknown" }
    for _, allowed := range allowedEnumValues {
        if o.GetReason() == allowed {
            return true
        }
    }
    return false
}
func (o *TransferData) isValidStatus() bool {
    var allowedEnumValues = []string{ "approvalPending", "atmWithdrawal", "atmWithdrawalReversalPending", "atmWithdrawalReversed", "authAdjustmentAuthorised", "authAdjustmentError", "authAdjustmentRefused", "authorised", "bankTransfer", "bankTransferPending", "booked", "bookingPending", "cancelled", "capturePending", "captureReversalPending", "captureReversed", "captured", "capturedExternally", "chargeback", "chargebackExternally", "chargebackPending", "chargebackReversalPending", "chargebackReversed", "credited", "depositCorrection", "depositCorrectionPending", "dispute", "disputeClosed", "disputeExpired", "disputeNeedsReview", "error", "expired", "failed", "fee", "feePending", "internalTransfer", "internalTransferPending", "invoiceDeduction", "invoiceDeductionPending", "manualCorrectionPending", "manuallyCorrected", "matchedStatement", "matchedStatementPending", "merchantPayin", "merchantPayinPending", "merchantPayinReversed", "merchantPayinReversedPending", "miscCost", "miscCostPending", "paymentCost", "paymentCostPending", "pendingApproval", "pendingExecution", "received", "refundPending", "refundReversalPending", "refundReversed", "refunded", "refundedExternally", "refused", "rejected", "reserveAdjustment", "reserveAdjustmentPending", "returned", "secondChargeback", "secondChargebackPending", "undefined" }
    for _, allowed := range allowedEnumValues {
        if o.GetStatus() == allowed {
            return true
        }
    }
    return false
}
func (o *TransferData) isValidType() bool {
    var allowedEnumValues = []string{ "payment", "capture", "captureReversal", "refund", "refundReversal", "chargeback", "chargebackCorrection", "chargebackReversal", "chargebackReversalCorrection", "secondChargeback", "secondChargebackCorrection", "atmWithdrawal", "atmWithdrawalReversal", "internalTransfer", "internalDirectDebit", "manualCorrection", "invoiceDeduction", "depositCorrection", "reserveAdjustment", "bankTransfer", "bankDirectDebit", "cardTransfer", "miscCost", "paymentCost", "fee", "leftover", "grant", "capitalFundsCollection", "cashOutInstruction", "cashoutFee", "cashoutRepayment", "cashoutFunding", "repayment", "installment", "installmentReversal", "balanceAdjustment", "balanceRollover", "balanceMigration" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

