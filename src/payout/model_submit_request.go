/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the SubmitRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubmitRequest{}

// SubmitRequest struct for SubmitRequest
type SubmitRequest struct {
	// This field contains additional data, which may be required for a particular request.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Amount         Amount             `json:"amount"`
	// The date of birth. Format: ISO-8601; example: YYYY-MM-DD  For Paysafecard it must be the same as used when registering the Paysafecard account.  > This field is mandatory for natural persons.  > This field is required to update the existing `dateOfBirth` that is associated with this recurring contract.
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	// The type of the entity the payout is processed for.  Allowed values: * NaturalPerson * Company > This field is required to update the existing `entityType` that is associated with this recurring contract.
	EntityType *string `json:"entityType,omitempty"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset *int32 `json:"fraudOffset,omitempty"`
	// The merchant account identifier you want to process the transaction request with.
	MerchantAccount string `json:"merchantAccount"`
	// The shopper's nationality.  A valid value is an ISO 2-character country code (e.g. 'NL').  > This field is required to update the existing nationality that is associated with this recurring contract.
	Nationality *string   `json:"nationality,omitempty"`
	Recurring   Recurring `json:"recurring"`
	// The merchant reference for this payout. This reference will be used in all communication to the merchant about the status of the payout. Although it is a good idea to make sure it is unique, this is not a requirement.
	Reference string `json:"reference"`
	// This is the `recurringDetailReference` you want to use for this payout.  You can use the value LATEST to select the most recently used recurring detail.
	SelectedRecurringDetailReference string `json:"selectedRecurringDetailReference"`
	// The shopper's email address.
	ShopperEmail string `json:"shopperEmail"`
	ShopperName  *Name  `json:"shopperName,omitempty"`
	// The shopper's reference for the payout transaction.
	ShopperReference string `json:"shopperReference"`
	// The description of this payout. This description is shown on the bank statement of the shopper (if this is supported by the chosen payment method).
	ShopperStatement *string `json:"shopperStatement,omitempty"`
	// The shopper's social security number.
	SocialSecurityNumber *string `json:"socialSecurityNumber,omitempty"`
}

// NewSubmitRequest instantiates a new SubmitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitRequest(amount Amount, merchantAccount string, recurring Recurring, reference string, selectedRecurringDetailReference string, shopperEmail string, shopperReference string) *SubmitRequest {
	this := SubmitRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.Recurring = recurring
	this.Reference = reference
	this.SelectedRecurringDetailReference = selectedRecurringDetailReference
	this.ShopperEmail = shopperEmail
	this.ShopperReference = shopperReference
	return &this
}

// NewSubmitRequestWithDefaults instantiates a new SubmitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitRequestWithDefaults() *SubmitRequest {
	this := SubmitRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *SubmitRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *SubmitRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *SubmitRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAmount returns the Amount field value
func (o *SubmitRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *SubmitRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *SubmitRequest) GetDateOfBirth() string {
	if o == nil || common.IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetDateOfBirthOk() (*string, bool) {
	if o == nil || common.IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *SubmitRequest) HasDateOfBirth() bool {
	if o != nil && !common.IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *SubmitRequest) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *SubmitRequest) GetEntityType() string {
	if o == nil || common.IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *SubmitRequest) HasEntityType() bool {
	if o != nil && !common.IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *SubmitRequest) SetEntityType(v string) {
	o.EntityType = &v
}

// GetFraudOffset returns the FraudOffset field value if set, zero value otherwise.
func (o *SubmitRequest) GetFraudOffset() int32 {
	if o == nil || common.IsNil(o.FraudOffset) {
		var ret int32
		return ret
	}
	return *o.FraudOffset
}

// GetFraudOffsetOk returns a tuple with the FraudOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetFraudOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FraudOffset) {
		return nil, false
	}
	return o.FraudOffset, true
}

// HasFraudOffset returns a boolean if a field has been set.
func (o *SubmitRequest) HasFraudOffset() bool {
	if o != nil && !common.IsNil(o.FraudOffset) {
		return true
	}

	return false
}

// SetFraudOffset gets a reference to the given int32 and assigns it to the FraudOffset field.
func (o *SubmitRequest) SetFraudOffset(v int32) {
	o.FraudOffset = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *SubmitRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *SubmitRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *SubmitRequest) GetNationality() string {
	if o == nil || common.IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetNationalityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *SubmitRequest) HasNationality() bool {
	if o != nil && !common.IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *SubmitRequest) SetNationality(v string) {
	o.Nationality = &v
}

// GetRecurring returns the Recurring field value
func (o *SubmitRequest) GetRecurring() Recurring {
	if o == nil {
		var ret Recurring
		return ret
	}

	return o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetRecurringOk() (*Recurring, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurring, true
}

// SetRecurring sets field value
func (o *SubmitRequest) SetRecurring(v Recurring) {
	o.Recurring = v
}

// GetReference returns the Reference field value
func (o *SubmitRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *SubmitRequest) SetReference(v string) {
	o.Reference = v
}

// GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field value
func (o *SubmitRequest) GetSelectedRecurringDetailReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SelectedRecurringDetailReference
}

// GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field value
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedRecurringDetailReference, true
}

// SetSelectedRecurringDetailReference sets field value
func (o *SubmitRequest) SetSelectedRecurringDetailReference(v string) {
	o.SelectedRecurringDetailReference = v
}

// GetShopperEmail returns the ShopperEmail field value
func (o *SubmitRequest) GetShopperEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetShopperEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperEmail, true
}

// SetShopperEmail sets field value
func (o *SubmitRequest) SetShopperEmail(v string) {
	o.ShopperEmail = v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *SubmitRequest) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *SubmitRequest) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *SubmitRequest) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetShopperReference returns the ShopperReference field value
func (o *SubmitRequest) GetShopperReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperReference, true
}

// SetShopperReference sets field value
func (o *SubmitRequest) SetShopperReference(v string) {
	o.ShopperReference = v
}

// GetShopperStatement returns the ShopperStatement field value if set, zero value otherwise.
func (o *SubmitRequest) GetShopperStatement() string {
	if o == nil || common.IsNil(o.ShopperStatement) {
		var ret string
		return ret
	}
	return *o.ShopperStatement
}

// GetShopperStatementOk returns a tuple with the ShopperStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetShopperStatementOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperStatement) {
		return nil, false
	}
	return o.ShopperStatement, true
}

// HasShopperStatement returns a boolean if a field has been set.
func (o *SubmitRequest) HasShopperStatement() bool {
	if o != nil && !common.IsNil(o.ShopperStatement) {
		return true
	}

	return false
}

// SetShopperStatement gets a reference to the given string and assigns it to the ShopperStatement field.
func (o *SubmitRequest) SetShopperStatement(v string) {
	o.ShopperStatement = &v
}

// GetSocialSecurityNumber returns the SocialSecurityNumber field value if set, zero value otherwise.
func (o *SubmitRequest) GetSocialSecurityNumber() string {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		var ret string
		return ret
	}
	return *o.SocialSecurityNumber
}

// GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRequest) GetSocialSecurityNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		return nil, false
	}
	return o.SocialSecurityNumber, true
}

// HasSocialSecurityNumber returns a boolean if a field has been set.
func (o *SubmitRequest) HasSocialSecurityNumber() bool {
	if o != nil && !common.IsNil(o.SocialSecurityNumber) {
		return true
	}

	return false
}

// SetSocialSecurityNumber gets a reference to the given string and assigns it to the SocialSecurityNumber field.
func (o *SubmitRequest) SetSocialSecurityNumber(v string) {
	o.SocialSecurityNumber = &v
}

func (o SubmitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !common.IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !common.IsNil(o.FraudOffset) {
		toSerialize["fraudOffset"] = o.FraudOffset
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	toSerialize["recurring"] = o.Recurring
	toSerialize["reference"] = o.Reference
	toSerialize["selectedRecurringDetailReference"] = o.SelectedRecurringDetailReference
	toSerialize["shopperEmail"] = o.ShopperEmail
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	toSerialize["shopperReference"] = o.ShopperReference
	if !common.IsNil(o.ShopperStatement) {
		toSerialize["shopperStatement"] = o.ShopperStatement
	}
	if !common.IsNil(o.SocialSecurityNumber) {
		toSerialize["socialSecurityNumber"] = o.SocialSecurityNumber
	}
	return toSerialize, nil
}

type NullableSubmitRequest struct {
	value *SubmitRequest
	isSet bool
}

func (v NullableSubmitRequest) Get() *SubmitRequest {
	return v.value
}

func (v *NullableSubmitRequest) Set(val *SubmitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitRequest(val *SubmitRequest) *NullableSubmitRequest {
	return &NullableSubmitRequest{value: val, isSet: true}
}

func (v NullableSubmitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SubmitRequest) isValidEntityType() bool {
	var allowedEnumValues = []string{"NaturalPerson", "Company"}
	for _, allowed := range allowedEnumValues {
		if o.GetEntityType() == allowed {
			return true
		}
	}
	return false
}
