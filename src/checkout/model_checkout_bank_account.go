/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the CheckoutBankAccount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutBankAccount{}

// CheckoutBankAccount struct for CheckoutBankAccount
type CheckoutBankAccount struct {
	// The type of the bank account.
	AccountType *string `json:"accountType,omitempty"`
	// The bank account number (without separators).
	BankAccountNumber *string `json:"bankAccountNumber,omitempty"`
	// The bank city.
	BankCity *string `json:"bankCity,omitempty"`
	// The location id of the bank. The field value is `nil` in most cases.
	BankLocationId *string `json:"bankLocationId,omitempty"`
	// The name of the bank.
	BankName *string `json:"bankName,omitempty"`
	// The [Business Identifier Code](https://en.wikipedia.org/wiki/ISO_9362) (BIC) is the SWIFT address assigned to a bank. The field value is `nil` in most cases.
	Bic *string `json:"bic,omitempty"`
	// Country code where the bank is located.  A valid value is an ISO two-character country code (e.g. 'NL').
	CountryCode *string `json:"countryCode,omitempty"`
	// The [International Bank Account Number](https://en.wikipedia.org/wiki/International_Bank_Account_Number) (IBAN).
	Iban *string `json:"iban,omitempty"`
	// The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don't accept 'ø'. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. > If provided details don't match the required format, the response returns the error message: 203 'Invalid bank account holder name'.
	OwnerName *string `json:"ownerName,omitempty"`
	// The bank account holder's tax ID.
	TaxId *string `json:"taxId,omitempty"`
}

// NewCheckoutBankAccount instantiates a new CheckoutBankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutBankAccount() *CheckoutBankAccount {
	this := CheckoutBankAccount{}
	return &this
}

// NewCheckoutBankAccountWithDefaults instantiates a new CheckoutBankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutBankAccountWithDefaults() *CheckoutBankAccount {
	this := CheckoutBankAccount{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *CheckoutBankAccount) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetBankAccountNumber() string {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasBankAccountNumber() bool {
	if o != nil && !common.IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *CheckoutBankAccount) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetBankCity returns the BankCity field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetBankCity() string {
	if o == nil || common.IsNil(o.BankCity) {
		var ret string
		return ret
	}
	return *o.BankCity
}

// GetBankCityOk returns a tuple with the BankCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetBankCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankCity) {
		return nil, false
	}
	return o.BankCity, true
}

// HasBankCity returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasBankCity() bool {
	if o != nil && !common.IsNil(o.BankCity) {
		return true
	}

	return false
}

// SetBankCity gets a reference to the given string and assigns it to the BankCity field.
func (o *CheckoutBankAccount) SetBankCity(v string) {
	o.BankCity = &v
}

// GetBankLocationId returns the BankLocationId field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetBankLocationId() string {
	if o == nil || common.IsNil(o.BankLocationId) {
		var ret string
		return ret
	}
	return *o.BankLocationId
}

// GetBankLocationIdOk returns a tuple with the BankLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetBankLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankLocationId) {
		return nil, false
	}
	return o.BankLocationId, true
}

// HasBankLocationId returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasBankLocationId() bool {
	if o != nil && !common.IsNil(o.BankLocationId) {
		return true
	}

	return false
}

// SetBankLocationId gets a reference to the given string and assigns it to the BankLocationId field.
func (o *CheckoutBankAccount) SetBankLocationId(v string) {
	o.BankLocationId = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetBankName() string {
	if o == nil || common.IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetBankNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasBankName() bool {
	if o != nil && !common.IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *CheckoutBankAccount) SetBankName(v string) {
	o.BankName = &v
}

// GetBic returns the Bic field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetBic() string {
	if o == nil || common.IsNil(o.Bic) {
		var ret string
		return ret
	}
	return *o.Bic
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetBicOk() (*string, bool) {
	if o == nil || common.IsNil(o.Bic) {
		return nil, false
	}
	return o.Bic, true
}

// HasBic returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasBic() bool {
	if o != nil && !common.IsNil(o.Bic) {
		return true
	}

	return false
}

// SetBic gets a reference to the given string and assigns it to the Bic field.
func (o *CheckoutBankAccount) SetBic(v string) {
	o.Bic = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *CheckoutBankAccount) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetIban() string {
	if o == nil || common.IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetIbanOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasIban() bool {
	if o != nil && !common.IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *CheckoutBankAccount) SetIban(v string) {
	o.Iban = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetOwnerName() string {
	if o == nil || common.IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetOwnerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasOwnerName() bool {
	if o != nil && !common.IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *CheckoutBankAccount) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *CheckoutBankAccount) GetTaxId() string {
	if o == nil || common.IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankAccount) GetTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *CheckoutBankAccount) HasTaxId() bool {
	if o != nil && !common.IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *CheckoutBankAccount) SetTaxId(v string) {
	o.TaxId = &v
}

func (o CheckoutBankAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutBankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !common.IsNil(o.BankAccountNumber) {
		toSerialize["bankAccountNumber"] = o.BankAccountNumber
	}
	if !common.IsNil(o.BankCity) {
		toSerialize["bankCity"] = o.BankCity
	}
	if !common.IsNil(o.BankLocationId) {
		toSerialize["bankLocationId"] = o.BankLocationId
	}
	if !common.IsNil(o.BankName) {
		toSerialize["bankName"] = o.BankName
	}
	if !common.IsNil(o.Bic) {
		toSerialize["bic"] = o.Bic
	}
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !common.IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !common.IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !common.IsNil(o.TaxId) {
		toSerialize["taxId"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableCheckoutBankAccount struct {
	value *CheckoutBankAccount
	isSet bool
}

func (v NullableCheckoutBankAccount) Get() *CheckoutBankAccount {
	return v.value
}

func (v *NullableCheckoutBankAccount) Set(val *CheckoutBankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutBankAccount(val *CheckoutBankAccount) *NullableCheckoutBankAccount {
	return &NullableCheckoutBankAccount{value: val, isSet: true}
}

func (v NullableCheckoutBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutBankAccount) isValidAccountType() bool {
	var allowedEnumValues = []string{"balance", "checking", "deposit", "general", "other", "payment", "savings"}
	for _, allowed := range allowedEnumValues {
		if o.GetAccountType() == allowed {
			return true
		}
	}
	return false
}
