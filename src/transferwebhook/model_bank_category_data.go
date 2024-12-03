/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the BankCategoryData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankCategoryData{}

// BankCategoryData struct for BankCategoryData
type BankCategoryData struct {
	// The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Required for transfers with `category` **bank**.  Possible values:  * **regular**: for normal, low-value transactions.  * **fast**: a faster way to transfer funds, but the fees are higher. Recommended for high-priority, low-value transactions.  * **wire**: the fastest way to transfer funds, but this has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: for instant funds transfers in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: for high-value transfers to a recipient in a different country.  * **internal**: for transfers to an Adyen-issued business bank account (by bank account number/IBAN).
	Priority *string `json:"priority,omitempty"`
	// **bank**
	Type *string `json:"type,omitempty"`
}

// NewBankCategoryData instantiates a new BankCategoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankCategoryData() *BankCategoryData {
	this := BankCategoryData{}
	var type_ string = "bank"
	this.Type = &type_
	return &this
}

// NewBankCategoryDataWithDefaults instantiates a new BankCategoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankCategoryDataWithDefaults() *BankCategoryData {
	this := BankCategoryData{}
	var type_ string = "bank"
	this.Type = &type_
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *BankCategoryData) GetPriority() string {
	if o == nil || common.IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankCategoryData) GetPriorityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *BankCategoryData) HasPriority() bool {
	if o != nil && !common.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *BankCategoryData) SetPriority(v string) {
	o.Priority = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BankCategoryData) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankCategoryData) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BankCategoryData) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BankCategoryData) SetType(v string) {
	o.Type = &v
}

func (o BankCategoryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankCategoryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableBankCategoryData struct {
	value *BankCategoryData
	isSet bool
}

func (v NullableBankCategoryData) Get() *BankCategoryData {
	return v.value
}

func (v *NullableBankCategoryData) Set(val *BankCategoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableBankCategoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableBankCategoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankCategoryData(val *BankCategoryData) *NullableBankCategoryData {
	return &NullableBankCategoryData{value: val, isSet: true}
}

func (v NullableBankCategoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankCategoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *BankCategoryData) isValidPriority() bool {
    var allowedEnumValues = []string{ "crossBorder", "fast", "instant", "internal", "regular", "wire" }
    for _, allowed := range allowedEnumValues {
        if o.GetPriority() == allowed {
            return true
        }
    }
    return false
}
func (o *BankCategoryData) isValidType() bool {
    var allowedEnumValues = []string{ "bank" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

