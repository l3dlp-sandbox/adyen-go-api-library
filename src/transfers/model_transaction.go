/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Transaction{}

// Transaction struct for Transaction
type Transaction struct {
	AccountHolder  ResourceReference `json:"accountHolder"`
	Amount         Amount            `json:"amount"`
	BalanceAccount ResourceReference `json:"balanceAccount"`
	// The unique identifier of the balance platform.
	BalancePlatform string `json:"balancePlatform"`
	// The date the transaction was booked into the balance account.
	BookingDate time.Time `json:"bookingDate"`
	// The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// The `description` from the `/transfers` request.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the transaction.
	Id                string             `json:"id"`
	PaymentInstrument *PaymentInstrument `json:"paymentInstrument,omitempty"`
	// The reference sent to or received from the counterparty.  * For outgoing funds, this is the [`referenceForBeneficiary`](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__resParam_referenceForBeneficiary) from the  [`/transfers`](https://docs.adyen.com/api-explorer/#/transfers/latest/post/transfers__reqParam_referenceForBeneficiary) request.   * For incoming funds, this is the reference from the sender.
	ReferenceForBeneficiary *string `json:"referenceForBeneficiary,omitempty"`
	// The status of the transaction.   Possible values:  * **pending**: The transaction is still pending.  * **booked**: The transaction has been booked to the balance account.
	Status   string        `json:"status"`
	Transfer *TransferView `json:"transfer,omitempty"`
	// The date the transfer amount becomes available in the balance account.
	ValueDate time.Time `json:"valueDate"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(accountHolder ResourceReference, amount Amount, balanceAccount ResourceReference, balancePlatform string, bookingDate time.Time, id string, status string, valueDate time.Time) *Transaction {
	this := Transaction{}
	this.AccountHolder = accountHolder
	this.Amount = amount
	this.BalanceAccount = balanceAccount
	this.BalancePlatform = balancePlatform
	this.BookingDate = bookingDate
	this.Id = id
	this.Status = status
	this.ValueDate = valueDate
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value
func (o *Transaction) GetAccountHolder() ResourceReference {
	if o == nil {
		var ret ResourceReference
		return ret
	}

	return o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAccountHolderOk() (*ResourceReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolder, true
}

// SetAccountHolder sets field value
func (o *Transaction) SetAccountHolder(v ResourceReference) {
	o.AccountHolder = v
}

// GetAmount returns the Amount field value
func (o *Transaction) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Transaction) SetAmount(v Amount) {
	o.Amount = v
}

// GetBalanceAccount returns the BalanceAccount field value
func (o *Transaction) GetBalanceAccount() ResourceReference {
	if o == nil {
		var ret ResourceReference
		return ret
	}

	return o.BalanceAccount
}

// GetBalanceAccountOk returns a tuple with the BalanceAccount field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetBalanceAccountOk() (*ResourceReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceAccount, true
}

// SetBalanceAccount sets field value
func (o *Transaction) SetBalanceAccount(v ResourceReference) {
	o.BalanceAccount = v
}

// GetBalancePlatform returns the BalancePlatform field value
func (o *Transaction) GetBalancePlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetBalancePlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalancePlatform, true
}

// SetBalancePlatform sets field value
func (o *Transaction) SetBalancePlatform(v string) {
	o.BalancePlatform = v
}

// GetBookingDate returns the BookingDate field value
func (o *Transaction) GetBookingDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BookingDate
}

// GetBookingDateOk returns a tuple with the BookingDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetBookingDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookingDate, true
}

// SetBookingDate sets field value
func (o *Transaction) SetBookingDate(v time.Time) {
	o.BookingDate = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Transaction) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Transaction) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Transaction) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Transaction) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Transaction) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Transaction) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *Transaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v string) {
	o.Id = v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise.
func (o *Transaction) GetPaymentInstrument() PaymentInstrument {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		var ret PaymentInstrument
		return ret
	}
	return *o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPaymentInstrumentOk() (*PaymentInstrument, bool) {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		return nil, false
	}
	return o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *Transaction) HasPaymentInstrument() bool {
	if o != nil && !common.IsNil(o.PaymentInstrument) {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given PaymentInstrument and assigns it to the PaymentInstrument field.
func (o *Transaction) SetPaymentInstrument(v PaymentInstrument) {
	o.PaymentInstrument = &v
}

// GetReferenceForBeneficiary returns the ReferenceForBeneficiary field value if set, zero value otherwise.
func (o *Transaction) GetReferenceForBeneficiary() string {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		var ret string
		return ret
	}
	return *o.ReferenceForBeneficiary
}

// GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetReferenceForBeneficiaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		return nil, false
	}
	return o.ReferenceForBeneficiary, true
}

// HasReferenceForBeneficiary returns a boolean if a field has been set.
func (o *Transaction) HasReferenceForBeneficiary() bool {
	if o != nil && !common.IsNil(o.ReferenceForBeneficiary) {
		return true
	}

	return false
}

// SetReferenceForBeneficiary gets a reference to the given string and assigns it to the ReferenceForBeneficiary field.
func (o *Transaction) SetReferenceForBeneficiary(v string) {
	o.ReferenceForBeneficiary = &v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v string) {
	o.Status = v
}

// GetTransfer returns the Transfer field value if set, zero value otherwise.
func (o *Transaction) GetTransfer() TransferView {
	if o == nil || common.IsNil(o.Transfer) {
		var ret TransferView
		return ret
	}
	return *o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransferOk() (*TransferView, bool) {
	if o == nil || common.IsNil(o.Transfer) {
		return nil, false
	}
	return o.Transfer, true
}

// HasTransfer returns a boolean if a field has been set.
func (o *Transaction) HasTransfer() bool {
	if o != nil && !common.IsNil(o.Transfer) {
		return true
	}

	return false
}

// SetTransfer gets a reference to the given TransferView and assigns it to the Transfer field.
func (o *Transaction) SetTransfer(v TransferView) {
	o.Transfer = &v
}

// GetValueDate returns the ValueDate field value
func (o *Transaction) GetValueDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetValueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueDate, true
}

// SetValueDate sets field value
func (o *Transaction) SetValueDate(v time.Time) {
	o.ValueDate = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountHolder"] = o.AccountHolder
	toSerialize["amount"] = o.Amount
	toSerialize["balanceAccount"] = o.BalanceAccount
	toSerialize["balancePlatform"] = o.BalancePlatform
	toSerialize["bookingDate"] = o.BookingDate
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !common.IsNil(o.PaymentInstrument) {
		toSerialize["paymentInstrument"] = o.PaymentInstrument
	}
	if !common.IsNil(o.ReferenceForBeneficiary) {
		toSerialize["referenceForBeneficiary"] = o.ReferenceForBeneficiary
	}
	toSerialize["status"] = o.Status
	if !common.IsNil(o.Transfer) {
		toSerialize["transfer"] = o.Transfer
	}
	toSerialize["valueDate"] = o.ValueDate
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Transaction) isValidStatus() bool {
	var allowedEnumValues = []string{"booked", "pending"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
