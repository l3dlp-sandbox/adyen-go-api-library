/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the VisaCheckoutDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VisaCheckoutDetails{}

// VisaCheckoutDetails struct for VisaCheckoutDetails
type VisaCheckoutDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// **visacheckout**
	Type *string `json:"type,omitempty"`
	// The Visa Click to Pay Call ID value. When your shopper selects a payment and/or a shipping address from Visa Click to Pay, you will receive a Visa Click to Pay Call ID.
	VisaCheckoutCallId string `json:"visaCheckoutCallId"`
}

// NewVisaCheckoutDetails instantiates a new VisaCheckoutDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisaCheckoutDetails(visaCheckoutCallId string) *VisaCheckoutDetails {
	this := VisaCheckoutDetails{}
	var type_ string = "visacheckout"
	this.Type = &type_
	this.VisaCheckoutCallId = visaCheckoutCallId
	return &this
}

// NewVisaCheckoutDetailsWithDefaults instantiates a new VisaCheckoutDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisaCheckoutDetailsWithDefaults() *VisaCheckoutDetails {
	this := VisaCheckoutDetails{}
	var type_ string = "visacheckout"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *VisaCheckoutDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisaCheckoutDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *VisaCheckoutDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *VisaCheckoutDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *VisaCheckoutDetails) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisaCheckoutDetails) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *VisaCheckoutDetails) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *VisaCheckoutDetails) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VisaCheckoutDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisaCheckoutDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VisaCheckoutDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VisaCheckoutDetails) SetType(v string) {
	o.Type = &v
}

// GetVisaCheckoutCallId returns the VisaCheckoutCallId field value
func (o *VisaCheckoutDetails) GetVisaCheckoutCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisaCheckoutCallId
}

// GetVisaCheckoutCallIdOk returns a tuple with the VisaCheckoutCallId field value
// and a boolean to check if the value has been set.
func (o *VisaCheckoutDetails) GetVisaCheckoutCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisaCheckoutCallId, true
}

// SetVisaCheckoutCallId sets field value
func (o *VisaCheckoutDetails) SetVisaCheckoutCallId(v string) {
	o.VisaCheckoutCallId = v
}

func (o VisaCheckoutDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisaCheckoutDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["visaCheckoutCallId"] = o.VisaCheckoutCallId
	return toSerialize, nil
}

type NullableVisaCheckoutDetails struct {
	value *VisaCheckoutDetails
	isSet bool
}

func (v NullableVisaCheckoutDetails) Get() *VisaCheckoutDetails {
	return v.value
}

func (v *NullableVisaCheckoutDetails) Set(val *VisaCheckoutDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableVisaCheckoutDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableVisaCheckoutDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisaCheckoutDetails(val *VisaCheckoutDetails) *NullableVisaCheckoutDetails {
	return &NullableVisaCheckoutDetails{value: val, isSet: true}
}

func (v NullableVisaCheckoutDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisaCheckoutDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *VisaCheckoutDetails) isValidFundingSource() bool {
	var allowedEnumValues = []string{"debit"}
	for _, allowed := range allowedEnumValues {
		if o.GetFundingSource() == allowed {
			return true
		}
	}
	return false
}
func (o *VisaCheckoutDetails) isValidType() bool {
	var allowedEnumValues = []string{"visacheckout"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
