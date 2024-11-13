/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the UpdatePaymentInstrument type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdatePaymentInstrument{}

// UpdatePaymentInstrument struct for UpdatePaymentInstrument
type UpdatePaymentInstrument struct {
	// Contains optional, additional business account details. Returned when you create a payment instrument with `type` **bankAccount**.
	// Deprecated since Configuration API v2
	// Please use `bankAccount` object instead
	AdditionalBankAccountIdentifications []PaymentInstrumentAdditionalBankAccountIdentificationsInner `json:"additionalBankAccountIdentifications,omitempty"`
	// The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/balanceAccounts__resParam_id) associated with the payment instrument.
	BalanceAccountId string              `json:"balanceAccountId"`
	BankAccount      *BankAccountDetails `json:"bankAccount,omitempty"`
	Card             *Card               `json:"card,omitempty"`
	// Your description for the payment instrument, maximum 300 characters.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the payment instrument.
	Id string `json:"id"`
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**.
	IssuingCountryCode string `json:"issuingCountryCode"`
	// The unique identifier of the [payment instrument group](https://docs.adyen.com/api-explorer/#/balanceplatform/v1/post/paymentInstrumentGroups__resParam_id) to which the payment instrument belongs.
	PaymentInstrumentGroupId *string `json:"paymentInstrumentGroupId,omitempty"`
	// Your reference for the payment instrument, maximum 150 characters.
	Reference *string `json:"reference,omitempty"`
	// The status of the payment instrument. If a status is not specified when creating a payment instrument, it is set to **active** by default. However, there can be exceptions for cards based on the `card.formFactor` and the `issuingCountryCode`. For example, when issuing physical cards in the US, the default status is **inactive**.  Possible values:    * **active**:  The payment instrument is active and can be used to make payments.    * **inactive**: The payment instrument is inactive and cannot be used to make payments.    * **suspended**: The payment instrument is suspended, either because it was stolen or lost.    * **closed**: The payment instrument is permanently closed. This action cannot be undone.
	Status *string `json:"status,omitempty"`
	// Comment for the status of the payment instrument.  Required if `statusReason` is **other**.
	StatusComment *string `json:"statusComment,omitempty"`
	// The reason for the status of the payment instrument.  Possible values: **accountClosure**, **damaged**, **endOfLife**, **expired**, **lost**, **stolen**, **suspectedFraud**, **transactionRule**, **other**. If the reason is **other**, you must also send the `statusComment` parameter describing the status change.
	StatusReason *string `json:"statusReason,omitempty"`
	// The type of payment instrument.  Possible values: **card**, **bankAccount**.
	Type string `json:"type"`
}

// NewUpdatePaymentInstrument instantiates a new UpdatePaymentInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentInstrument(balanceAccountId string, id string, issuingCountryCode string, type_ string) *UpdatePaymentInstrument {
	this := UpdatePaymentInstrument{}
	this.BalanceAccountId = balanceAccountId
	this.Id = id
	this.IssuingCountryCode = issuingCountryCode
	this.Type = type_
	return &this
}

// NewUpdatePaymentInstrumentWithDefaults instantiates a new UpdatePaymentInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentInstrumentWithDefaults() *UpdatePaymentInstrument {
	this := UpdatePaymentInstrument{}
	return &this
}

// GetAdditionalBankAccountIdentifications returns the AdditionalBankAccountIdentifications field value if set, zero value otherwise.
// Deprecated since Configuration API v2
// Please use `bankAccount` object instead
func (o *UpdatePaymentInstrument) GetAdditionalBankAccountIdentifications() []PaymentInstrumentAdditionalBankAccountIdentificationsInner {
	if o == nil || common.IsNil(o.AdditionalBankAccountIdentifications) {
		var ret []PaymentInstrumentAdditionalBankAccountIdentificationsInner
		return ret
	}
	return o.AdditionalBankAccountIdentifications
}

// GetAdditionalBankAccountIdentificationsOk returns a tuple with the AdditionalBankAccountIdentifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Configuration API v2
// Please use `bankAccount` object instead
func (o *UpdatePaymentInstrument) GetAdditionalBankAccountIdentificationsOk() ([]PaymentInstrumentAdditionalBankAccountIdentificationsInner, bool) {
	if o == nil || common.IsNil(o.AdditionalBankAccountIdentifications) {
		return nil, false
	}
	return o.AdditionalBankAccountIdentifications, true
}

// HasAdditionalBankAccountIdentifications returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasAdditionalBankAccountIdentifications() bool {
	if o != nil && !common.IsNil(o.AdditionalBankAccountIdentifications) {
		return true
	}

	return false
}

// SetAdditionalBankAccountIdentifications gets a reference to the given []PaymentInstrumentAdditionalBankAccountIdentificationsInner and assigns it to the AdditionalBankAccountIdentifications field.
// Deprecated since Configuration API v2
// Please use `bankAccount` object instead
func (o *UpdatePaymentInstrument) SetAdditionalBankAccountIdentifications(v []PaymentInstrumentAdditionalBankAccountIdentificationsInner) {
	o.AdditionalBankAccountIdentifications = v
}

// GetBalanceAccountId returns the BalanceAccountId field value
func (o *UpdatePaymentInstrument) GetBalanceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceAccountId, true
}

// SetBalanceAccountId sets field value
func (o *UpdatePaymentInstrument) SetBalanceAccountId(v string) {
	o.BalanceAccountId = v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetBankAccount() BankAccountDetails {
	if o == nil || common.IsNil(o.BankAccount) {
		var ret BankAccountDetails
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetBankAccountOk() (*BankAccountDetails, bool) {
	if o == nil || common.IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasBankAccount() bool {
	if o != nil && !common.IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given BankAccountDetails and assigns it to the BankAccount field.
func (o *UpdatePaymentInstrument) SetBankAccount(v BankAccountDetails) {
	o.BankAccount = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *UpdatePaymentInstrument) SetCard(v Card) {
	o.Card = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdatePaymentInstrument) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *UpdatePaymentInstrument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdatePaymentInstrument) SetId(v string) {
	o.Id = v
}

// GetIssuingCountryCode returns the IssuingCountryCode field value
func (o *UpdatePaymentInstrument) GetIssuingCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuingCountryCode
}

// GetIssuingCountryCodeOk returns a tuple with the IssuingCountryCode field value
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetIssuingCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuingCountryCode, true
}

// SetIssuingCountryCode sets field value
func (o *UpdatePaymentInstrument) SetIssuingCountryCode(v string) {
	o.IssuingCountryCode = v
}

// GetPaymentInstrumentGroupId returns the PaymentInstrumentGroupId field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetPaymentInstrumentGroupId() string {
	if o == nil || common.IsNil(o.PaymentInstrumentGroupId) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentGroupId
}

// GetPaymentInstrumentGroupIdOk returns a tuple with the PaymentInstrumentGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetPaymentInstrumentGroupIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentInstrumentGroupId) {
		return nil, false
	}
	return o.PaymentInstrumentGroupId, true
}

// HasPaymentInstrumentGroupId returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasPaymentInstrumentGroupId() bool {
	if o != nil && !common.IsNil(o.PaymentInstrumentGroupId) {
		return true
	}

	return false
}

// SetPaymentInstrumentGroupId gets a reference to the given string and assigns it to the PaymentInstrumentGroupId field.
func (o *UpdatePaymentInstrument) SetPaymentInstrumentGroupId(v string) {
	o.PaymentInstrumentGroupId = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *UpdatePaymentInstrument) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdatePaymentInstrument) SetStatus(v string) {
	o.Status = &v
}

// GetStatusComment returns the StatusComment field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetStatusComment() string {
	if o == nil || common.IsNil(o.StatusComment) {
		var ret string
		return ret
	}
	return *o.StatusComment
}

// GetStatusCommentOk returns a tuple with the StatusComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetStatusCommentOk() (*string, bool) {
	if o == nil || common.IsNil(o.StatusComment) {
		return nil, false
	}
	return o.StatusComment, true
}

// HasStatusComment returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasStatusComment() bool {
	if o != nil && !common.IsNil(o.StatusComment) {
		return true
	}

	return false
}

// SetStatusComment gets a reference to the given string and assigns it to the StatusComment field.
func (o *UpdatePaymentInstrument) SetStatusComment(v string) {
	o.StatusComment = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *UpdatePaymentInstrument) GetStatusReason() string {
	if o == nil || common.IsNil(o.StatusReason) {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetStatusReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.StatusReason) {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *UpdatePaymentInstrument) HasStatusReason() bool {
	if o != nil && !common.IsNil(o.StatusReason) {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *UpdatePaymentInstrument) SetStatusReason(v string) {
	o.StatusReason = &v
}

// GetType returns the Type field value
func (o *UpdatePaymentInstrument) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdatePaymentInstrument) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdatePaymentInstrument) SetType(v string) {
	o.Type = v
}

func (o UpdatePaymentInstrument) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentInstrument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalBankAccountIdentifications) {
		toSerialize["additionalBankAccountIdentifications"] = o.AdditionalBankAccountIdentifications
	}
	toSerialize["balanceAccountId"] = o.BalanceAccountId
	if !common.IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["issuingCountryCode"] = o.IssuingCountryCode
	if !common.IsNil(o.PaymentInstrumentGroupId) {
		toSerialize["paymentInstrumentGroupId"] = o.PaymentInstrumentGroupId
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StatusComment) {
		toSerialize["statusComment"] = o.StatusComment
	}
	if !common.IsNil(o.StatusReason) {
		toSerialize["statusReason"] = o.StatusReason
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableUpdatePaymentInstrument struct {
	value *UpdatePaymentInstrument
	isSet bool
}

func (v NullableUpdatePaymentInstrument) Get() *UpdatePaymentInstrument {
	return v.value
}

func (v *NullableUpdatePaymentInstrument) Set(val *UpdatePaymentInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentInstrument(val *UpdatePaymentInstrument) *NullableUpdatePaymentInstrument {
	return &NullableUpdatePaymentInstrument{value: val, isSet: true}
}

func (v NullableUpdatePaymentInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpdatePaymentInstrument) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "closed", "inactive", "suspended"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdatePaymentInstrument) isValidStatusReason() bool {
	var allowedEnumValues = []string{"accountClosure", "damaged", "endOfLife", "expired", "lost", "other", "stolen", "suspectedFraud", "transactionRule"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatusReason() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdatePaymentInstrument) isValidType() bool {
	var allowedEnumValues = []string{"bankAccount", "card"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
