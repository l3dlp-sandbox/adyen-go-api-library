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

// checks if the PayWithGoogleDonations type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayWithGoogleDonations{}

// PayWithGoogleDonations struct for PayWithGoogleDonations
type PayWithGoogleDonations struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// The `token` that you obtained from the [Google Pay API](https://developers.google.com/pay/api/web/reference/response-objects#PaymentData) `PaymentData` response.
	GooglePayToken string `json:"googlePayToken"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// Required for mobile integrations. Version of the 3D Secure 2 mobile SDK.
	ThreeDS2SdkVersion *string `json:"threeDS2SdkVersion,omitempty"`
	// **paywithgoogle**
	Type *string `json:"type,omitempty"`
}

// NewPayWithGoogleDonations instantiates a new PayWithGoogleDonations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayWithGoogleDonations(googlePayToken string) *PayWithGoogleDonations {
	this := PayWithGoogleDonations{}
	this.GooglePayToken = googlePayToken
	var type_ string = "paywithgoogle"
	this.Type = &type_
	return &this
}

// NewPayWithGoogleDonationsWithDefaults instantiates a new PayWithGoogleDonations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayWithGoogleDonationsWithDefaults() *PayWithGoogleDonations {
	this := PayWithGoogleDonations{}
	var type_ string = "paywithgoogle"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PayWithGoogleDonations) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayWithGoogleDonations) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PayWithGoogleDonations) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PayWithGoogleDonations) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *PayWithGoogleDonations) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayWithGoogleDonations) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *PayWithGoogleDonations) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *PayWithGoogleDonations) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetGooglePayToken returns the GooglePayToken field value
func (o *PayWithGoogleDonations) GetGooglePayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GooglePayToken
}

// GetGooglePayTokenOk returns a tuple with the GooglePayToken field value
// and a boolean to check if the value has been set.
func (o *PayWithGoogleDonations) GetGooglePayTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GooglePayToken, true
}

// SetGooglePayToken sets field value
func (o *PayWithGoogleDonations) SetGooglePayToken(v string) {
	o.GooglePayToken = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PayWithGoogleDonations) GetRecurringDetailReference() string {
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
func (o *PayWithGoogleDonations) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *PayWithGoogleDonations) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PayWithGoogleDonations) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *PayWithGoogleDonations) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayWithGoogleDonations) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *PayWithGoogleDonations) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *PayWithGoogleDonations) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetThreeDS2SdkVersion returns the ThreeDS2SdkVersion field value if set, zero value otherwise.
func (o *PayWithGoogleDonations) GetThreeDS2SdkVersion() string {
	if o == nil || common.IsNil(o.ThreeDS2SdkVersion) {
		var ret string
		return ret
	}
	return *o.ThreeDS2SdkVersion
}

// GetThreeDS2SdkVersionOk returns a tuple with the ThreeDS2SdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayWithGoogleDonations) GetThreeDS2SdkVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDS2SdkVersion) {
		return nil, false
	}
	return o.ThreeDS2SdkVersion, true
}

// HasThreeDS2SdkVersion returns a boolean if a field has been set.
func (o *PayWithGoogleDonations) HasThreeDS2SdkVersion() bool {
	if o != nil && !common.IsNil(o.ThreeDS2SdkVersion) {
		return true
	}

	return false
}

// SetThreeDS2SdkVersion gets a reference to the given string and assigns it to the ThreeDS2SdkVersion field.
func (o *PayWithGoogleDonations) SetThreeDS2SdkVersion(v string) {
	o.ThreeDS2SdkVersion = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PayWithGoogleDonations) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayWithGoogleDonations) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PayWithGoogleDonations) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PayWithGoogleDonations) SetType(v string) {
	o.Type = &v
}

func (o PayWithGoogleDonations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayWithGoogleDonations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	toSerialize["googlePayToken"] = o.GooglePayToken
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.ThreeDS2SdkVersion) {
		toSerialize["threeDS2SdkVersion"] = o.ThreeDS2SdkVersion
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePayWithGoogleDonations struct {
	value *PayWithGoogleDonations
	isSet bool
}

func (v NullablePayWithGoogleDonations) Get() *PayWithGoogleDonations {
	return v.value
}

func (v *NullablePayWithGoogleDonations) Set(val *PayWithGoogleDonations) {
	v.value = val
	v.isSet = true
}

func (v NullablePayWithGoogleDonations) IsSet() bool {
	return v.isSet
}

func (v *NullablePayWithGoogleDonations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayWithGoogleDonations(val *PayWithGoogleDonations) *NullablePayWithGoogleDonations {
	return &NullablePayWithGoogleDonations{value: val, isSet: true}
}

func (v NullablePayWithGoogleDonations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayWithGoogleDonations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PayWithGoogleDonations) isValidFundingSource() bool {
	var allowedEnumValues = []string{"credit", "debit"}
	for _, allowed := range allowedEnumValues {
		if o.GetFundingSource() == allowed {
			return true
		}
	}
	return false
}
func (o *PayWithGoogleDonations) isValidType() bool {
	var allowedEnumValues = []string{"paywithgoogle"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
