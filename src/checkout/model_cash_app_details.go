/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the CashAppDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CashAppDetails{}

// CashAppDetails struct for CashAppDetails
type CashAppDetails struct {
	// Cash App issued cashtag for recurring payment
	Cashtag *string `json:"cashtag,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// Cash App issued customerId for recurring payment
	CustomerId *string `json:"customerId,omitempty"`
	// Cash App issued grantId for one time payment
	GrantId *string `json:"grantId,omitempty"`
	// Cash App issued onFileGrantId for recurring payment
	OnFileGrantId *string `json:"onFileGrantId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// Cash App request id
	RequestId *string `json:"requestId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// The payment method subtype.
	Subtype *string `json:"subtype,omitempty"`
	// cashapp
	Type *string `json:"type,omitempty"`
}

// NewCashAppDetails instantiates a new CashAppDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashAppDetails() *CashAppDetails {
	this := CashAppDetails{}
	var type_ string = "cashapp"
	this.Type = &type_
	return &this
}

// NewCashAppDetailsWithDefaults instantiates a new CashAppDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashAppDetailsWithDefaults() *CashAppDetails {
	this := CashAppDetails{}
	var type_ string = "cashapp"
	this.Type = &type_
	return &this
}

// GetCashtag returns the Cashtag field value if set, zero value otherwise.
func (o *CashAppDetails) GetCashtag() string {
	if o == nil || common.IsNil(o.Cashtag) {
		var ret string
		return ret
	}
	return *o.Cashtag
}

// GetCashtagOk returns a tuple with the Cashtag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetCashtagOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cashtag) {
		return nil, false
	}
	return o.Cashtag, true
}

// HasCashtag returns a boolean if a field has been set.
func (o *CashAppDetails) HasCashtag() bool {
	if o != nil && !common.IsNil(o.Cashtag) {
		return true
	}

	return false
}

// SetCashtag gets a reference to the given string and assigns it to the Cashtag field.
func (o *CashAppDetails) SetCashtag(v string) {
	o.Cashtag = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *CashAppDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *CashAppDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *CashAppDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CashAppDetails) GetCustomerId() string {
	if o == nil || common.IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetCustomerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CashAppDetails) HasCustomerId() bool {
	if o != nil && !common.IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CashAppDetails) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *CashAppDetails) GetGrantId() string {
	if o == nil || common.IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetGrantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *CashAppDetails) HasGrantId() bool {
	if o != nil && !common.IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *CashAppDetails) SetGrantId(v string) {
	o.GrantId = &v
}

// GetOnFileGrantId returns the OnFileGrantId field value if set, zero value otherwise.
func (o *CashAppDetails) GetOnFileGrantId() string {
	if o == nil || common.IsNil(o.OnFileGrantId) {
		var ret string
		return ret
	}
	return *o.OnFileGrantId
}

// GetOnFileGrantIdOk returns a tuple with the OnFileGrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetOnFileGrantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OnFileGrantId) {
		return nil, false
	}
	return o.OnFileGrantId, true
}

// HasOnFileGrantId returns a boolean if a field has been set.
func (o *CashAppDetails) HasOnFileGrantId() bool {
	if o != nil && !common.IsNil(o.OnFileGrantId) {
		return true
	}

	return false
}

// SetOnFileGrantId gets a reference to the given string and assigns it to the OnFileGrantId field.
func (o *CashAppDetails) SetOnFileGrantId(v string) {
	o.OnFileGrantId = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated
func (o *CashAppDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CashAppDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *CashAppDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated
func (o *CashAppDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CashAppDetails) GetRequestId() string {
	if o == nil || common.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CashAppDetails) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CashAppDetails) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *CashAppDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *CashAppDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *CashAppDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *CashAppDetails) GetSubtype() string {
	if o == nil || common.IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetSubtypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *CashAppDetails) HasSubtype() bool {
	if o != nil && !common.IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *CashAppDetails) SetSubtype(v string) {
	o.Subtype = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CashAppDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAppDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CashAppDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CashAppDetails) SetType(v string) {
	o.Type = &v
}

func (o CashAppDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashAppDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Cashtag) {
		toSerialize["cashtag"] = o.Cashtag
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !common.IsNil(o.GrantId) {
		toSerialize["grantId"] = o.GrantId
	}
	if !common.IsNil(o.OnFileGrantId) {
		toSerialize["onFileGrantId"] = o.OnFileGrantId
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCashAppDetails struct {
	value *CashAppDetails
	isSet bool
}

func (v NullableCashAppDetails) Get() *CashAppDetails {
	return v.value
}

func (v *NullableCashAppDetails) Set(val *CashAppDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCashAppDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCashAppDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashAppDetails(val *CashAppDetails) *NullableCashAppDetails {
	return &NullableCashAppDetails{value: val, isSet: true}
}

func (v NullableCashAppDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashAppDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CashAppDetails) isValidType() bool {
	var allowedEnumValues = []string{"cashapp"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
