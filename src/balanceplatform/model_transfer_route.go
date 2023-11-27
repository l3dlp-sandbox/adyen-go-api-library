/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the TransferRoute type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferRoute{}

// TransferRoute struct for TransferRoute
type TransferRoute struct {
	//  The type of transfer.   Possible values:    - **bank**: Transfer to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.
	Category *string `json:"category,omitempty"`
	// The two-character ISO-3166-1 alpha-2 country code of the counterparty. For example, **US** or **NL**.
	Country *string `json:"country,omitempty"`
	// The three-character ISO currency code of transfer. For example, **USD** or **EUR**.
	Currency *string `json:"currency,omitempty"`
	// The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Possible values:  * **regular**: For normal, low-value transactions.  * **fast**: Faster way to transfer funds but has higher fees. Recommended for high-priority, low-value transactions.  * **wire**: Fastest way to transfer funds but has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: Instant way to transfer funds in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: High-value transfer to a recipient in a different country.  * **internal**: Transfer to an Adyen-issued business bank account (by bank account number/IBAN).
	Priority     *string                    `json:"priority,omitempty"`
	Requirements *TransferRouteRequirements `json:"requirements,omitempty"`
}

// NewTransferRoute instantiates a new TransferRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRoute() *TransferRoute {
	this := TransferRoute{}
	return &this
}

// NewTransferRouteWithDefaults instantiates a new TransferRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRouteWithDefaults() *TransferRoute {
	this := TransferRoute{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *TransferRoute) GetCategory() string {
	if o == nil || common.IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRoute) GetCategoryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *TransferRoute) HasCategory() bool {
	if o != nil && !common.IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *TransferRoute) SetCategory(v string) {
	o.Category = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TransferRoute) GetCountry() string {
	if o == nil || common.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRoute) GetCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TransferRoute) HasCountry() bool {
	if o != nil && !common.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TransferRoute) SetCountry(v string) {
	o.Country = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TransferRoute) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRoute) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransferRoute) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TransferRoute) SetCurrency(v string) {
	o.Currency = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TransferRoute) GetPriority() string {
	if o == nil || common.IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRoute) GetPriorityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TransferRoute) HasPriority() bool {
	if o != nil && !common.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *TransferRoute) SetPriority(v string) {
	o.Priority = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *TransferRoute) GetRequirements() TransferRouteRequirements {
	if o == nil || common.IsNil(o.Requirements) {
		var ret TransferRouteRequirements
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRoute) GetRequirementsOk() (*TransferRouteRequirements, bool) {
	if o == nil || common.IsNil(o.Requirements) {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *TransferRoute) HasRequirements() bool {
	if o != nil && !common.IsNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given TransferRouteRequirements and assigns it to the Requirements field.
func (o *TransferRoute) SetRequirements(v TransferRouteRequirements) {
	o.Requirements = &v
}

func (o TransferRoute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !common.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !common.IsNil(o.Requirements) {
		toSerialize["requirements"] = o.Requirements
	}
	return toSerialize, nil
}

type NullableTransferRoute struct {
	value *TransferRoute
	isSet bool
}

func (v NullableTransferRoute) Get() *TransferRoute {
	return v.value
}

func (v *NullableTransferRoute) Set(val *TransferRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRoute(val *TransferRoute) *NullableTransferRoute {
	return &NullableTransferRoute{value: val, isSet: true}
}

func (v NullableTransferRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferRoute) isValidCategory() bool {
	var allowedEnumValues = []string{"bank", "card", "grants", "internal", "issuedCard", "migration", "platformPayment", "topUp", "upgrade"}
	for _, allowed := range allowedEnumValues {
		if o.GetCategory() == allowed {
			return true
		}
	}
	return false
}
func (o *TransferRoute) isValidPriority() bool {
	var allowedEnumValues = []string{"crossBorder", "fast", "instant", "internal", "regular", "wire"}
	for _, allowed := range allowedEnumValues {
		if o.GetPriority() == allowed {
			return true
		}
	}
	return false
}
