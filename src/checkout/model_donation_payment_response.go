/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the DonationPaymentResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DonationPaymentResponse{}

// DonationPaymentResponse struct for DonationPaymentResponse
type DonationPaymentResponse struct {
	Amount *Amount `json:"amount,omitempty"`
	// The Adyen account name of your charity. We will provide you with this account name once your chosen charity has been [onboarded](https://docs.adyen.com/online-payments/donations#onboarding).
	DonationAccount *string `json:"donationAccount,omitempty"`
	// Your unique resource identifier.
	Id *string `json:"id,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount *string          `json:"merchantAccount,omitempty"`
	Payment         *PaymentResponse `json:"payment,omitempty"`
	// The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\"). Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// The status of the donation transaction.  Possible values: * **completed** * **pending** * **refused**
	Status *string `json:"status,omitempty"`
}

// NewDonationPaymentResponse instantiates a new DonationPaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDonationPaymentResponse() *DonationPaymentResponse {
	this := DonationPaymentResponse{}
	return &this
}

// NewDonationPaymentResponseWithDefaults instantiates a new DonationPaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDonationPaymentResponseWithDefaults() *DonationPaymentResponse {
	this := DonationPaymentResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DonationPaymentResponse) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationPaymentResponse) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DonationPaymentResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *DonationPaymentResponse) SetAmount(v Amount) {
	o.Amount = &v
}

// GetDonationAccount returns the DonationAccount field value if set, zero value otherwise.
func (o *DonationPaymentResponse) GetDonationAccount() string {
	if o == nil || common.IsNil(o.DonationAccount) {
		var ret string
		return ret
	}
	return *o.DonationAccount
}

// GetDonationAccountOk returns a tuple with the DonationAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationPaymentResponse) GetDonationAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.DonationAccount) {
		return nil, false
	}
	return o.DonationAccount, true
}

// HasDonationAccount returns a boolean if a field has been set.
func (o *DonationPaymentResponse) HasDonationAccount() bool {
	if o != nil && !common.IsNil(o.DonationAccount) {
		return true
	}

	return false
}

// SetDonationAccount gets a reference to the given string and assigns it to the DonationAccount field.
func (o *DonationPaymentResponse) SetDonationAccount(v string) {
	o.DonationAccount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DonationPaymentResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationPaymentResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DonationPaymentResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DonationPaymentResponse) SetId(v string) {
	o.Id = &v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *DonationPaymentResponse) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationPaymentResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *DonationPaymentResponse) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *DonationPaymentResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *DonationPaymentResponse) GetPayment() PaymentResponse {
	if o == nil || common.IsNil(o.Payment) {
		var ret PaymentResponse
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationPaymentResponse) GetPaymentOk() (*PaymentResponse, bool) {
	if o == nil || common.IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *DonationPaymentResponse) HasPayment() bool {
	if o != nil && !common.IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given PaymentResponse and assigns it to the Payment field.
func (o *DonationPaymentResponse) SetPayment(v PaymentResponse) {
	o.Payment = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *DonationPaymentResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationPaymentResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *DonationPaymentResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *DonationPaymentResponse) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DonationPaymentResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationPaymentResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DonationPaymentResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DonationPaymentResponse) SetStatus(v string) {
	o.Status = &v
}

func (o DonationPaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DonationPaymentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.DonationAccount) {
		toSerialize["donationAccount"] = o.DonationAccount
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableDonationPaymentResponse struct {
	value *DonationPaymentResponse
	isSet bool
}

func (v NullableDonationPaymentResponse) Get() *DonationPaymentResponse {
	return v.value
}

func (v *NullableDonationPaymentResponse) Set(val *DonationPaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDonationPaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDonationPaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonationPaymentResponse(val *DonationPaymentResponse) *NullableDonationPaymentResponse {
	return &NullableDonationPaymentResponse{value: val, isSet: true}
}

func (v NullableDonationPaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonationPaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DonationPaymentResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"completed", "pending", "refused"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
