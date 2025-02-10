/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the EftDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EftDetails{}

// EftDetails struct for EftDetails
type EftDetails struct {
	// The bank account number (without separators).
	BankAccountNumber *string `json:"bankAccountNumber,omitempty"`
	// The financial institution code.
	BankCode *string `json:"bankCode,omitempty"`
	// The bank routing number of the account.
	BankLocationId *string `json:"bankLocationId,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don't accept 'ø'. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. > If provided details don't match the required format, the response returns the error message: 203 'Invalid bank account holder name'.
	OwnerName *string `json:"ownerName,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **eft**
	Type *string `json:"type,omitempty"`
}

// NewEftDetails instantiates a new EftDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEftDetails() *EftDetails {
	this := EftDetails{}
	var type_ string = "eft_directdebit_CA"
	this.Type = &type_
	return &this
}

// NewEftDetailsWithDefaults instantiates a new EftDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEftDetailsWithDefaults() *EftDetails {
	this := EftDetails{}
	var type_ string = "eft_directdebit_CA"
	this.Type = &type_
	return &this
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *EftDetails) GetBankAccountNumber() string {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftDetails) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *EftDetails) HasBankAccountNumber() bool {
	if o != nil && !common.IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *EftDetails) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetBankCode returns the BankCode field value if set, zero value otherwise.
func (o *EftDetails) GetBankCode() string {
	if o == nil || common.IsNil(o.BankCode) {
		var ret string
		return ret
	}
	return *o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftDetails) GetBankCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankCode) {
		return nil, false
	}
	return o.BankCode, true
}

// HasBankCode returns a boolean if a field has been set.
func (o *EftDetails) HasBankCode() bool {
	if o != nil && !common.IsNil(o.BankCode) {
		return true
	}

	return false
}

// SetBankCode gets a reference to the given string and assigns it to the BankCode field.
func (o *EftDetails) SetBankCode(v string) {
	o.BankCode = &v
}

// GetBankLocationId returns the BankLocationId field value if set, zero value otherwise.
func (o *EftDetails) GetBankLocationId() string {
	if o == nil || common.IsNil(o.BankLocationId) {
		var ret string
		return ret
	}
	return *o.BankLocationId
}

// GetBankLocationIdOk returns a tuple with the BankLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftDetails) GetBankLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankLocationId) {
		return nil, false
	}
	return o.BankLocationId, true
}

// HasBankLocationId returns a boolean if a field has been set.
func (o *EftDetails) HasBankLocationId() bool {
	if o != nil && !common.IsNil(o.BankLocationId) {
		return true
	}

	return false
}

// SetBankLocationId gets a reference to the given string and assigns it to the BankLocationId field.
func (o *EftDetails) SetBankLocationId(v string) {
	o.BankLocationId = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *EftDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *EftDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *EftDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *EftDetails) GetOwnerName() string {
	if o == nil || common.IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftDetails) GetOwnerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *EftDetails) HasOwnerName() bool {
	if o != nil && !common.IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *EftDetails) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *EftDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *EftDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *EftDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *EftDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *EftDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *EftDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *EftDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EftDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EftDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EftDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EftDetails) SetType(v string) {
	o.Type = &v
}

func (o EftDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EftDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BankAccountNumber) {
		toSerialize["bankAccountNumber"] = o.BankAccountNumber
	}
	if !common.IsNil(o.BankCode) {
		toSerialize["bankCode"] = o.BankCode
	}
	if !common.IsNil(o.BankLocationId) {
		toSerialize["bankLocationId"] = o.BankLocationId
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableEftDetails struct {
	value *EftDetails
	isSet bool
}

func (v NullableEftDetails) Get() *EftDetails {
	return v.value
}

func (v *NullableEftDetails) Set(val *EftDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEftDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEftDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEftDetails(val *EftDetails) *NullableEftDetails {
	return &NullableEftDetails{value: val, isSet: true}
}

func (v NullableEftDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEftDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *EftDetails) isValidType() bool {
	var allowedEnumValues = []string{"eft_directdebit_CA"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
