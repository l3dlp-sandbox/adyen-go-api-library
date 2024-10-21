/*
Configuration webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the BankAccountDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankAccountDetails{}

// BankAccountDetails struct for BankAccountDetails
type BankAccountDetails struct {
	// The bank account number, without separators or whitespace.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The bank account type.  Possible values: **checking** or **savings**. Defaults to **checking**.
	AccountType *string `json:"accountType,omitempty"`
	// The bank account branch number, without separators or whitespace
	BranchNumber *string `json:"branchNumber,omitempty"`
	// Business accounts with a `formFactor` value of **physical** are business accounts issued under the central bank of that country. The default value is **physical** for NL, US, and UK business accounts.   Adyen creates a local IBAN for business accounts when the `formFactor` value is set to **virtual**. The local IBANs that are supported are for DE and FR, which reference a physical NL account, with funds being routed through the central bank of NL.
	FormFactor *string `json:"formFactor,omitempty"`
	// The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard.
	Iban *string `json:"iban,omitempty"`
	// The [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace.
	RoutingNumber *string `json:"routingNumber,omitempty"`
	// The [sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or whitespace.
	SortCode *string `json:"sortCode,omitempty"`
	// **iban** or **usLocal** or **ukLocal**
	Type string `json:"type"`
}

// NewBankAccountDetails instantiates a new BankAccountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountDetails(type_ string) *BankAccountDetails {
	this := BankAccountDetails{}
	var accountType string = "checking"
	this.AccountType = &accountType
	var formFactor string = "physical"
	this.FormFactor = &formFactor
	this.Type = type_
	return &this
}

// NewBankAccountDetailsWithDefaults instantiates a new BankAccountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountDetailsWithDefaults() *BankAccountDetails {
	this := BankAccountDetails{}
	var accountType string = "checking"
	this.AccountType = &accountType
	var formFactor string = "physical"
	this.FormFactor = &formFactor
	var type_ string = "iban"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *BankAccountDetails) GetAccountNumber() string {
	if o == nil || common.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *BankAccountDetails) HasAccountNumber() bool {
	if o != nil && !common.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *BankAccountDetails) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *BankAccountDetails) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *BankAccountDetails) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *BankAccountDetails) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBranchNumber returns the BranchNumber field value if set, zero value otherwise.
func (o *BankAccountDetails) GetBranchNumber() string {
	if o == nil || common.IsNil(o.BranchNumber) {
		var ret string
		return ret
	}
	return *o.BranchNumber
}

// GetBranchNumberOk returns a tuple with the BranchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetBranchNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.BranchNumber) {
		return nil, false
	}
	return o.BranchNumber, true
}

// HasBranchNumber returns a boolean if a field has been set.
func (o *BankAccountDetails) HasBranchNumber() bool {
	if o != nil && !common.IsNil(o.BranchNumber) {
		return true
	}

	return false
}

// SetBranchNumber gets a reference to the given string and assigns it to the BranchNumber field.
func (o *BankAccountDetails) SetBranchNumber(v string) {
	o.BranchNumber = &v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise.
func (o *BankAccountDetails) GetFormFactor() string {
	if o == nil || common.IsNil(o.FormFactor) {
		var ret string
		return ret
	}
	return *o.FormFactor
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetFormFactorOk() (*string, bool) {
	if o == nil || common.IsNil(o.FormFactor) {
		return nil, false
	}
	return o.FormFactor, true
}

// HasFormFactor returns a boolean if a field has been set.
func (o *BankAccountDetails) HasFormFactor() bool {
	if o != nil && !common.IsNil(o.FormFactor) {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given string and assigns it to the FormFactor field.
func (o *BankAccountDetails) SetFormFactor(v string) {
	o.FormFactor = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *BankAccountDetails) GetIban() string {
	if o == nil || common.IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetIbanOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *BankAccountDetails) HasIban() bool {
	if o != nil && !common.IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *BankAccountDetails) SetIban(v string) {
	o.Iban = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *BankAccountDetails) GetRoutingNumber() string {
	if o == nil || common.IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetRoutingNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *BankAccountDetails) HasRoutingNumber() bool {
	if o != nil && !common.IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *BankAccountDetails) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetSortCode returns the SortCode field value if set, zero value otherwise.
func (o *BankAccountDetails) GetSortCode() string {
	if o == nil || common.IsNil(o.SortCode) {
		var ret string
		return ret
	}
	return *o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetSortCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SortCode) {
		return nil, false
	}
	return o.SortCode, true
}

// HasSortCode returns a boolean if a field has been set.
func (o *BankAccountDetails) HasSortCode() bool {
	if o != nil && !common.IsNil(o.SortCode) {
		return true
	}

	return false
}

// SetSortCode gets a reference to the given string and assigns it to the SortCode field.
func (o *BankAccountDetails) SetSortCode(v string) {
	o.SortCode = &v
}

// GetType returns the Type field value
func (o *BankAccountDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BankAccountDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BankAccountDetails) SetType(v string) {
	o.Type = v
}

func (o BankAccountDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !common.IsNil(o.BranchNumber) {
		toSerialize["branchNumber"] = o.BranchNumber
	}
	if !common.IsNil(o.FormFactor) {
		toSerialize["formFactor"] = o.FormFactor
	}
	if !common.IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !common.IsNil(o.RoutingNumber) {
		toSerialize["routingNumber"] = o.RoutingNumber
	}
	if !common.IsNil(o.SortCode) {
		toSerialize["sortCode"] = o.SortCode
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBankAccountDetails struct {
	value *BankAccountDetails
	isSet bool
}

func (v NullableBankAccountDetails) Get() *BankAccountDetails {
	return v.value
}

func (v *NullableBankAccountDetails) Set(val *BankAccountDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountDetails(val *BankAccountDetails) *NullableBankAccountDetails {
	return &NullableBankAccountDetails{value: val, isSet: true}
}

func (v NullableBankAccountDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
