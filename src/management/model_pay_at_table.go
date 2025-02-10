/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the PayAtTable type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayAtTable{}

// PayAtTable struct for PayAtTable
type PayAtTable struct {
	// Allowed authentication methods: Magswipe, Manual Entry.
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
	// Enable Pay at table.
	EnablePayAtTable *bool `json:"enablePayAtTable,omitempty"`
	// Sets the allowed payment instrument for Pay at table transactions.  Can be: **cash** or **card**. If not set, the terminal presents both options.
	PaymentInstrument common.NullableString `json:"paymentInstrument,omitempty"`
}

// NewPayAtTable instantiates a new PayAtTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayAtTable() *PayAtTable {
	this := PayAtTable{}
	return &this
}

// NewPayAtTableWithDefaults instantiates a new PayAtTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayAtTableWithDefaults() *PayAtTable {
	this := PayAtTable{}
	return &this
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *PayAtTable) GetAuthenticationMethod() string {
	if o == nil || common.IsNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayAtTable) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *PayAtTable) HasAuthenticationMethod() bool {
	if o != nil && !common.IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *PayAtTable) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

// GetEnablePayAtTable returns the EnablePayAtTable field value if set, zero value otherwise.
func (o *PayAtTable) GetEnablePayAtTable() bool {
	if o == nil || common.IsNil(o.EnablePayAtTable) {
		var ret bool
		return ret
	}
	return *o.EnablePayAtTable
}

// GetEnablePayAtTableOk returns a tuple with the EnablePayAtTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayAtTable) GetEnablePayAtTableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnablePayAtTable) {
		return nil, false
	}
	return o.EnablePayAtTable, true
}

// HasEnablePayAtTable returns a boolean if a field has been set.
func (o *PayAtTable) HasEnablePayAtTable() bool {
	if o != nil && !common.IsNil(o.EnablePayAtTable) {
		return true
	}

	return false
}

// SetEnablePayAtTable gets a reference to the given bool and assigns it to the EnablePayAtTable field.
func (o *PayAtTable) SetEnablePayAtTable(v bool) {
	o.EnablePayAtTable = &v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayAtTable) GetPaymentInstrument() string {
	if o == nil || common.IsNil(o.PaymentInstrument.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentInstrument.Get()
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayAtTable) GetPaymentInstrumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentInstrument.Get(), o.PaymentInstrument.IsSet()
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *PayAtTable) HasPaymentInstrument() bool {
	if o != nil && o.PaymentInstrument.IsSet() {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given NullableString and assigns it to the PaymentInstrument field.
func (o *PayAtTable) SetPaymentInstrument(v string) {
	o.PaymentInstrument.Set(&v)
}

// SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil
func (o *PayAtTable) SetPaymentInstrumentNil() {
	o.PaymentInstrument.Set(nil)
}

// UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
func (o *PayAtTable) UnsetPaymentInstrument() {
	o.PaymentInstrument.Unset()
}

func (o PayAtTable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayAtTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	if !common.IsNil(o.EnablePayAtTable) {
		toSerialize["enablePayAtTable"] = o.EnablePayAtTable
	}
	if o.PaymentInstrument.IsSet() {
		toSerialize["paymentInstrument"] = o.PaymentInstrument.Get()
	}
	return toSerialize, nil
}

type NullablePayAtTable struct {
	value *PayAtTable
	isSet bool
}

func (v NullablePayAtTable) Get() *PayAtTable {
	return v.value
}

func (v *NullablePayAtTable) Set(val *PayAtTable) {
	v.value = val
	v.isSet = true
}

func (v NullablePayAtTable) IsSet() bool {
	return v.isSet
}

func (v *NullablePayAtTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayAtTable(val *PayAtTable) *NullablePayAtTable {
	return &NullablePayAtTable{value: val, isSet: true}
}

func (v NullablePayAtTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayAtTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PayAtTable) isValidAuthenticationMethod() bool {
	var allowedEnumValues = []string{"MAGSWIPE", "MKE"}
	for _, allowed := range allowedEnumValues {
		if o.GetAuthenticationMethod() == allowed {
			return true
		}
	}
	return false
}
func (o *PayAtTable) isValidPaymentInstrument() bool {
	var allowedEnumValues = []string{"Cash", "Card"}
	for _, allowed := range allowedEnumValues {
		if o.GetPaymentInstrument() == allowed {
			return true
		}
	}
	return false
}
