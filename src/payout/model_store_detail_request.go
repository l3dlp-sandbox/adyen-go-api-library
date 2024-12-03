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

// checks if the StoreDetailRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreDetailRequest{}

// StoreDetailRequest struct for StoreDetailRequest
type StoreDetailRequest struct {
	// This field contains additional data, which may be required for a particular request.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Bank *BankAccount `json:"bank,omitempty"`
	BillingAddress *Address `json:"billingAddress,omitempty"`
	Card *Card `json:"card,omitempty"`
	// The date of birth. Format: [ISO-8601](https://www.w3.org/TR/NOTE-datetime); example: YYYY-MM-DD For Paysafecard it must be the same as used when registering the Paysafecard account. > This field is mandatory for natural persons.
	DateOfBirth string `json:"dateOfBirth"`
	// The type of the entity the payout is processed for.
	EntityType string `json:"entityType"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset *int32 `json:"fraudOffset,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The shopper's nationality.  A valid value is an ISO 2-character country code (e.g. 'NL').
	Nationality string `json:"nationality"`
	Recurring Recurring `json:"recurring"`
	// The name of the brand to make a payout to.  For Paysafecard it must be set to `paysafecard`.
	SelectedBrand *string `json:"selectedBrand,omitempty"`
	// The shopper's email address.
	ShopperEmail string `json:"shopperEmail"`
	ShopperName *Name `json:"shopperName,omitempty"`
	// The shopper's reference for the payment transaction.
	ShopperReference string `json:"shopperReference"`
	// The shopper's social security number.
	SocialSecurityNumber *string `json:"socialSecurityNumber,omitempty"`
	// The shopper's phone number.
	TelephoneNumber *string `json:"telephoneNumber,omitempty"`
}

// NewStoreDetailRequest instantiates a new StoreDetailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDetailRequest(dateOfBirth string, entityType string, merchantAccount string, nationality string, recurring Recurring, shopperEmail string, shopperReference string) *StoreDetailRequest {
	this := StoreDetailRequest{}
	this.DateOfBirth = dateOfBirth
	this.EntityType = entityType
	this.MerchantAccount = merchantAccount
	this.Nationality = nationality
	this.Recurring = recurring
	this.ShopperEmail = shopperEmail
	this.ShopperReference = shopperReference
	return &this
}

// NewStoreDetailRequestWithDefaults instantiates a new StoreDetailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDetailRequestWithDefaults() *StoreDetailRequest {
	this := StoreDetailRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *StoreDetailRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetBank() BankAccount {
	if o == nil || common.IsNil(o.Bank) {
		var ret BankAccount
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetBankOk() (*BankAccount, bool) {
	if o == nil || common.IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasBank() bool {
	if o != nil && !common.IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given BankAccount and assigns it to the Bank field.
func (o *StoreDetailRequest) SetBank(v BankAccount) {
	o.Bank = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *StoreDetailRequest) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *StoreDetailRequest) SetCard(v Card) {
	o.Card = &v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *StoreDetailRequest) GetDateOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetDateOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *StoreDetailRequest) SetDateOfBirth(v string) {
	o.DateOfBirth = v
}

// GetEntityType returns the EntityType field value
func (o *StoreDetailRequest) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *StoreDetailRequest) SetEntityType(v string) {
	o.EntityType = v
}

// GetFraudOffset returns the FraudOffset field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetFraudOffset() int32 {
	if o == nil || common.IsNil(o.FraudOffset) {
		var ret int32
		return ret
	}
	return *o.FraudOffset
}

// GetFraudOffsetOk returns a tuple with the FraudOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetFraudOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FraudOffset) {
		return nil, false
	}
	return o.FraudOffset, true
}

// HasFraudOffset returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasFraudOffset() bool {
	if o != nil && !common.IsNil(o.FraudOffset) {
		return true
	}

	return false
}

// SetFraudOffset gets a reference to the given int32 and assigns it to the FraudOffset field.
func (o *StoreDetailRequest) SetFraudOffset(v int32) {
	o.FraudOffset = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StoreDetailRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StoreDetailRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetNationality returns the Nationality field value
func (o *StoreDetailRequest) GetNationality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetNationalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nationality, true
}

// SetNationality sets field value
func (o *StoreDetailRequest) SetNationality(v string) {
	o.Nationality = v
}

// GetRecurring returns the Recurring field value
func (o *StoreDetailRequest) GetRecurring() Recurring {
	if o == nil {
		var ret Recurring
		return ret
	}

	return o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetRecurringOk() (*Recurring, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurring, true
}

// SetRecurring sets field value
func (o *StoreDetailRequest) SetRecurring(v Recurring) {
	o.Recurring = v
}

// GetSelectedBrand returns the SelectedBrand field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetSelectedBrand() string {
	if o == nil || common.IsNil(o.SelectedBrand) {
		var ret string
		return ret
	}
	return *o.SelectedBrand
}

// GetSelectedBrandOk returns a tuple with the SelectedBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetSelectedBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelectedBrand) {
		return nil, false
	}
	return o.SelectedBrand, true
}

// HasSelectedBrand returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasSelectedBrand() bool {
	if o != nil && !common.IsNil(o.SelectedBrand) {
		return true
	}

	return false
}

// SetSelectedBrand gets a reference to the given string and assigns it to the SelectedBrand field.
func (o *StoreDetailRequest) SetSelectedBrand(v string) {
	o.SelectedBrand = &v
}

// GetShopperEmail returns the ShopperEmail field value
func (o *StoreDetailRequest) GetShopperEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetShopperEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperEmail, true
}

// SetShopperEmail sets field value
func (o *StoreDetailRequest) SetShopperEmail(v string) {
	o.ShopperEmail = v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *StoreDetailRequest) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetShopperReference returns the ShopperReference field value
func (o *StoreDetailRequest) GetShopperReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperReference, true
}

// SetShopperReference sets field value
func (o *StoreDetailRequest) SetShopperReference(v string) {
	o.ShopperReference = v
}

// GetSocialSecurityNumber returns the SocialSecurityNumber field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetSocialSecurityNumber() string {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		var ret string
		return ret
	}
	return *o.SocialSecurityNumber
}

// GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetSocialSecurityNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		return nil, false
	}
	return o.SocialSecurityNumber, true
}

// HasSocialSecurityNumber returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasSocialSecurityNumber() bool {
	if o != nil && !common.IsNil(o.SocialSecurityNumber) {
		return true
	}

	return false
}

// SetSocialSecurityNumber gets a reference to the given string and assigns it to the SocialSecurityNumber field.
func (o *StoreDetailRequest) SetSocialSecurityNumber(v string) {
	o.SocialSecurityNumber = &v
}

// GetTelephoneNumber returns the TelephoneNumber field value if set, zero value otherwise.
func (o *StoreDetailRequest) GetTelephoneNumber() string {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		var ret string
		return ret
	}
	return *o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailRequest) GetTelephoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		return nil, false
	}
	return o.TelephoneNumber, true
}

// HasTelephoneNumber returns a boolean if a field has been set.
func (o *StoreDetailRequest) HasTelephoneNumber() bool {
	if o != nil && !common.IsNil(o.TelephoneNumber) {
		return true
	}

	return false
}

// SetTelephoneNumber gets a reference to the given string and assigns it to the TelephoneNumber field.
func (o *StoreDetailRequest) SetTelephoneNumber(v string) {
	o.TelephoneNumber = &v
}

func (o StoreDetailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDetailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	toSerialize["dateOfBirth"] = o.DateOfBirth
	toSerialize["entityType"] = o.EntityType
	if !common.IsNil(o.FraudOffset) {
		toSerialize["fraudOffset"] = o.FraudOffset
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["nationality"] = o.Nationality
	toSerialize["recurring"] = o.Recurring
	if !common.IsNil(o.SelectedBrand) {
		toSerialize["selectedBrand"] = o.SelectedBrand
	}
	toSerialize["shopperEmail"] = o.ShopperEmail
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	toSerialize["shopperReference"] = o.ShopperReference
	if !common.IsNil(o.SocialSecurityNumber) {
		toSerialize["socialSecurityNumber"] = o.SocialSecurityNumber
	}
	if !common.IsNil(o.TelephoneNumber) {
		toSerialize["telephoneNumber"] = o.TelephoneNumber
	}
	return toSerialize, nil
}

type NullableStoreDetailRequest struct {
	value *StoreDetailRequest
	isSet bool
}

func (v NullableStoreDetailRequest) Get() *StoreDetailRequest {
	return v.value
}

func (v *NullableStoreDetailRequest) Set(val *StoreDetailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDetailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDetailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDetailRequest(val *StoreDetailRequest) *NullableStoreDetailRequest {
	return &NullableStoreDetailRequest{value: val, isSet: true}
}

func (v NullableStoreDetailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDetailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *StoreDetailRequest) isValidEntityType() bool {
    var allowedEnumValues = []string{ "NaturalPerson", "Company" }
    for _, allowed := range allowedEnumValues {
        if o.GetEntityType() == allowed {
            return true
        }
    }
    return false
}

