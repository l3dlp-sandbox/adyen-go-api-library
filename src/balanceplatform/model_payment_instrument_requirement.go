/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the PaymentInstrumentRequirement type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentInstrumentRequirement{}

// PaymentInstrumentRequirement struct for PaymentInstrumentRequirement
type PaymentInstrumentRequirement struct {
	// Specifies the requirements for the payment instrument that need to be included in the request for a particular route.
	Description *string `json:"description,omitempty"`
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the payment instrument is issued. For example, **NL** or **US**.
	IssuingCountryCode *string `json:"issuingCountryCode,omitempty"`
	// The two-character ISO-3166-1 alpha-2 country code list for payment instruments.
	IssuingCountryCodes []string `json:"issuingCountryCodes,omitempty"`
	// Specifies if the requirement only applies to transfers to another balance platform.
	OnlyForCrossBalancePlatform *bool `json:"onlyForCrossBalancePlatform,omitempty"`
	// The type of the payment instrument. For example, \"BankAccount\" or \"Card\".
	PaymentInstrumentType *string `json:"paymentInstrumentType,omitempty"`
	// **paymentInstrumentRequirement**
	Type string `json:"type"`
}

// NewPaymentInstrumentRequirement instantiates a new PaymentInstrumentRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInstrumentRequirement(type_ string) *PaymentInstrumentRequirement {
	this := PaymentInstrumentRequirement{}
	this.Type = type_
	return &this
}

// NewPaymentInstrumentRequirementWithDefaults instantiates a new PaymentInstrumentRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInstrumentRequirementWithDefaults() *PaymentInstrumentRequirement {
	this := PaymentInstrumentRequirement{}
	var type_ string = "paymentInstrumentRequirement"
	this.Type = type_
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentInstrumentRequirement) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRequirement) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentInstrumentRequirement) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentInstrumentRequirement) SetDescription(v string) {
	o.Description = &v
}

// GetIssuingCountryCode returns the IssuingCountryCode field value if set, zero value otherwise.
func (o *PaymentInstrumentRequirement) GetIssuingCountryCode() string {
	if o == nil || common.IsNil(o.IssuingCountryCode) {
		var ret string
		return ret
	}
	return *o.IssuingCountryCode
}

// GetIssuingCountryCodeOk returns a tuple with the IssuingCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRequirement) GetIssuingCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuingCountryCode) {
		return nil, false
	}
	return o.IssuingCountryCode, true
}

// HasIssuingCountryCode returns a boolean if a field has been set.
func (o *PaymentInstrumentRequirement) HasIssuingCountryCode() bool {
	if o != nil && !common.IsNil(o.IssuingCountryCode) {
		return true
	}

	return false
}

// SetIssuingCountryCode gets a reference to the given string and assigns it to the IssuingCountryCode field.
func (o *PaymentInstrumentRequirement) SetIssuingCountryCode(v string) {
	o.IssuingCountryCode = &v
}

// GetIssuingCountryCodes returns the IssuingCountryCodes field value if set, zero value otherwise.
func (o *PaymentInstrumentRequirement) GetIssuingCountryCodes() []string {
	if o == nil || common.IsNil(o.IssuingCountryCodes) {
		var ret []string
		return ret
	}
	return o.IssuingCountryCodes
}

// GetIssuingCountryCodesOk returns a tuple with the IssuingCountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRequirement) GetIssuingCountryCodesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IssuingCountryCodes) {
		return nil, false
	}
	return o.IssuingCountryCodes, true
}

// HasIssuingCountryCodes returns a boolean if a field has been set.
func (o *PaymentInstrumentRequirement) HasIssuingCountryCodes() bool {
	if o != nil && !common.IsNil(o.IssuingCountryCodes) {
		return true
	}

	return false
}

// SetIssuingCountryCodes gets a reference to the given []string and assigns it to the IssuingCountryCodes field.
func (o *PaymentInstrumentRequirement) SetIssuingCountryCodes(v []string) {
	o.IssuingCountryCodes = v
}

// GetOnlyForCrossBalancePlatform returns the OnlyForCrossBalancePlatform field value if set, zero value otherwise.
func (o *PaymentInstrumentRequirement) GetOnlyForCrossBalancePlatform() bool {
	if o == nil || common.IsNil(o.OnlyForCrossBalancePlatform) {
		var ret bool
		return ret
	}
	return *o.OnlyForCrossBalancePlatform
}

// GetOnlyForCrossBalancePlatformOk returns a tuple with the OnlyForCrossBalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRequirement) GetOnlyForCrossBalancePlatformOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OnlyForCrossBalancePlatform) {
		return nil, false
	}
	return o.OnlyForCrossBalancePlatform, true
}

// HasOnlyForCrossBalancePlatform returns a boolean if a field has been set.
func (o *PaymentInstrumentRequirement) HasOnlyForCrossBalancePlatform() bool {
	if o != nil && !common.IsNil(o.OnlyForCrossBalancePlatform) {
		return true
	}

	return false
}

// SetOnlyForCrossBalancePlatform gets a reference to the given bool and assigns it to the OnlyForCrossBalancePlatform field.
func (o *PaymentInstrumentRequirement) SetOnlyForCrossBalancePlatform(v bool) {
	o.OnlyForCrossBalancePlatform = &v
}

// GetPaymentInstrumentType returns the PaymentInstrumentType field value if set, zero value otherwise.
func (o *PaymentInstrumentRequirement) GetPaymentInstrumentType() string {
	if o == nil || common.IsNil(o.PaymentInstrumentType) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentType
}

// GetPaymentInstrumentTypeOk returns a tuple with the PaymentInstrumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRequirement) GetPaymentInstrumentTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentInstrumentType) {
		return nil, false
	}
	return o.PaymentInstrumentType, true
}

// HasPaymentInstrumentType returns a boolean if a field has been set.
func (o *PaymentInstrumentRequirement) HasPaymentInstrumentType() bool {
	if o != nil && !common.IsNil(o.PaymentInstrumentType) {
		return true
	}

	return false
}

// SetPaymentInstrumentType gets a reference to the given string and assigns it to the PaymentInstrumentType field.
func (o *PaymentInstrumentRequirement) SetPaymentInstrumentType(v string) {
	o.PaymentInstrumentType = &v
}

// GetType returns the Type field value
func (o *PaymentInstrumentRequirement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRequirement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentInstrumentRequirement) SetType(v string) {
	o.Type = v
}

func (o PaymentInstrumentRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInstrumentRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.IssuingCountryCode) {
		toSerialize["issuingCountryCode"] = o.IssuingCountryCode
	}
	if !common.IsNil(o.IssuingCountryCodes) {
		toSerialize["issuingCountryCodes"] = o.IssuingCountryCodes
	}
	if !common.IsNil(o.OnlyForCrossBalancePlatform) {
		toSerialize["onlyForCrossBalancePlatform"] = o.OnlyForCrossBalancePlatform
	}
	if !common.IsNil(o.PaymentInstrumentType) {
		toSerialize["paymentInstrumentType"] = o.PaymentInstrumentType
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePaymentInstrumentRequirement struct {
	value *PaymentInstrumentRequirement
	isSet bool
}

func (v NullablePaymentInstrumentRequirement) Get() *PaymentInstrumentRequirement {
	return v.value
}

func (v *NullablePaymentInstrumentRequirement) Set(val *PaymentInstrumentRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrumentRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrumentRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrumentRequirement(val *PaymentInstrumentRequirement) *NullablePaymentInstrumentRequirement {
	return &NullablePaymentInstrumentRequirement{value: val, isSet: true}
}

func (v NullablePaymentInstrumentRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrumentRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentInstrumentRequirement) isValidPaymentInstrumentType() bool {
	var allowedEnumValues = []string{"BankAccount", "Card"}
	for _, allowed := range allowedEnumValues {
		if o.GetPaymentInstrumentType() == allowed {
			return true
		}
	}
	return false
}
func (o *PaymentInstrumentRequirement) isValidType() bool {
	var allowedEnumValues = []string{"paymentInstrumentRequirement"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
