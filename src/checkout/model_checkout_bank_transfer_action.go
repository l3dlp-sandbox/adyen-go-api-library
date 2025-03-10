/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the CheckoutBankTransferAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutBankTransferAction{}

// CheckoutBankTransferAction struct for CheckoutBankTransferAction
type CheckoutBankTransferAction struct {
	// The account number of the bank transfer.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The name of the account holder.
	Beneficiary *string `json:"beneficiary,omitempty"`
	// The BIC of the IBAN.
	Bic *string `json:"bic,omitempty"`
	// The url to download payment details with.
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// The IBAN of the bank transfer.
	Iban *string `json:"iban,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// The transfer reference.
	Reference *string `json:"reference,omitempty"`
	// The routing number of the bank transfer.
	RoutingNumber *string `json:"routingNumber,omitempty"`
	// The e-mail of the shopper, included if an e-mail was sent to the shopper.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	// The sort code of the bank transfer.
	SortCode    *string `json:"sortCode,omitempty"`
	TotalAmount *Amount `json:"totalAmount,omitempty"`
	// The type of the action.
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutBankTransferAction instantiates a new CheckoutBankTransferAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutBankTransferAction(type_ string) *CheckoutBankTransferAction {
	this := CheckoutBankTransferAction{}
	this.Type = type_
	return &this
}

// NewCheckoutBankTransferActionWithDefaults instantiates a new CheckoutBankTransferAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutBankTransferActionWithDefaults() *CheckoutBankTransferAction {
	this := CheckoutBankTransferAction{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetAccountNumber() string {
	if o == nil || common.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasAccountNumber() bool {
	if o != nil && !common.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *CheckoutBankTransferAction) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetBeneficiary returns the Beneficiary field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetBeneficiary() string {
	if o == nil || common.IsNil(o.Beneficiary) {
		var ret string
		return ret
	}
	return *o.Beneficiary
}

// GetBeneficiaryOk returns a tuple with the Beneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetBeneficiaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Beneficiary) {
		return nil, false
	}
	return o.Beneficiary, true
}

// HasBeneficiary returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasBeneficiary() bool {
	if o != nil && !common.IsNil(o.Beneficiary) {
		return true
	}

	return false
}

// SetBeneficiary gets a reference to the given string and assigns it to the Beneficiary field.
func (o *CheckoutBankTransferAction) SetBeneficiary(v string) {
	o.Beneficiary = &v
}

// GetBic returns the Bic field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetBic() string {
	if o == nil || common.IsNil(o.Bic) {
		var ret string
		return ret
	}
	return *o.Bic
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetBicOk() (*string, bool) {
	if o == nil || common.IsNil(o.Bic) {
		return nil, false
	}
	return o.Bic, true
}

// HasBic returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasBic() bool {
	if o != nil && !common.IsNil(o.Bic) {
		return true
	}

	return false
}

// SetBic gets a reference to the given string and assigns it to the Bic field.
func (o *CheckoutBankTransferAction) SetBic(v string) {
	o.Bic = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetDownloadUrl() string {
	if o == nil || common.IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetDownloadUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasDownloadUrl() bool {
	if o != nil && !common.IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *CheckoutBankTransferAction) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetIban() string {
	if o == nil || common.IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetIbanOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasIban() bool {
	if o != nil && !common.IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *CheckoutBankTransferAction) SetIban(v string) {
	o.Iban = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutBankTransferAction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CheckoutBankTransferAction) SetReference(v string) {
	o.Reference = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetRoutingNumber() string {
	if o == nil || common.IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetRoutingNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasRoutingNumber() bool {
	if o != nil && !common.IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *CheckoutBankTransferAction) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *CheckoutBankTransferAction) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetSortCode returns the SortCode field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetSortCode() string {
	if o == nil || common.IsNil(o.SortCode) {
		var ret string
		return ret
	}
	return *o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetSortCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SortCode) {
		return nil, false
	}
	return o.SortCode, true
}

// HasSortCode returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasSortCode() bool {
	if o != nil && !common.IsNil(o.SortCode) {
		return true
	}

	return false
}

// SetSortCode gets a reference to the given string and assigns it to the SortCode field.
func (o *CheckoutBankTransferAction) SetSortCode(v string) {
	o.SortCode = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetTotalAmount() Amount {
	if o == nil || common.IsNil(o.TotalAmount) {
		var ret Amount
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetTotalAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasTotalAmount() bool {
	if o != nil && !common.IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given Amount and assigns it to the TotalAmount field.
func (o *CheckoutBankTransferAction) SetTotalAmount(v Amount) {
	o.TotalAmount = &v
}

// GetType returns the Type field value
func (o *CheckoutBankTransferAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutBankTransferAction) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutBankTransferAction) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutBankTransferAction) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutBankTransferAction) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutBankTransferAction) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutBankTransferAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutBankTransferAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !common.IsNil(o.Beneficiary) {
		toSerialize["beneficiary"] = o.Beneficiary
	}
	if !common.IsNil(o.Bic) {
		toSerialize["bic"] = o.Bic
	}
	if !common.IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !common.IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !common.IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.RoutingNumber) {
		toSerialize["routingNumber"] = o.RoutingNumber
	}
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	if !common.IsNil(o.SortCode) {
		toSerialize["sortCode"] = o.SortCode
	}
	if !common.IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutBankTransferAction struct {
	value *CheckoutBankTransferAction
	isSet bool
}

func (v NullableCheckoutBankTransferAction) Get() *CheckoutBankTransferAction {
	return v.value
}

func (v *NullableCheckoutBankTransferAction) Set(val *CheckoutBankTransferAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutBankTransferAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutBankTransferAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutBankTransferAction(val *CheckoutBankTransferAction) *NullableCheckoutBankTransferAction {
	return &NullableCheckoutBankTransferAction{value: val, isSet: true}
}

func (v NullableCheckoutBankTransferAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutBankTransferAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutBankTransferAction) isValidType() bool {
	var allowedEnumValues = []string{"bankTransfer"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
