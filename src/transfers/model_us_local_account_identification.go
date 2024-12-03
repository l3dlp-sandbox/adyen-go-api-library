/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the USLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &USLocalAccountIdentification{}

// USLocalAccountIdentification struct for USLocalAccountIdentification
type USLocalAccountIdentification struct {
	// The bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The bank account type.  Possible values: **checking** or **savings**. Defaults to **checking**.
	AccountType *string `json:"accountType,omitempty"`
	// The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace.
	RoutingNumber string `json:"routingNumber"`
	// **usLocal**
	Type string `json:"type"`
}

// NewUSLocalAccountIdentification instantiates a new USLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUSLocalAccountIdentification(accountNumber string, routingNumber string, type_ string) *USLocalAccountIdentification {
	this := USLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	var accountType string = "checking"
	this.AccountType = &accountType
	this.RoutingNumber = routingNumber
	this.Type = type_
	return &this
}

// NewUSLocalAccountIdentificationWithDefaults instantiates a new USLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUSLocalAccountIdentificationWithDefaults() *USLocalAccountIdentification {
	this := USLocalAccountIdentification{}
	var accountType string = "checking"
	this.AccountType = &accountType
	var type_ string = "usLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *USLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *USLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *USLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *USLocalAccountIdentification) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *USLocalAccountIdentification) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *USLocalAccountIdentification) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *USLocalAccountIdentification) SetAccountType(v string) {
	o.AccountType = &v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *USLocalAccountIdentification) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *USLocalAccountIdentification) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *USLocalAccountIdentification) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetType returns the Type field value
func (o *USLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *USLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *USLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o USLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o USLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	toSerialize["routingNumber"] = o.RoutingNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableUSLocalAccountIdentification struct {
	value *USLocalAccountIdentification
	isSet bool
}

func (v NullableUSLocalAccountIdentification) Get() *USLocalAccountIdentification {
	return v.value
}

func (v *NullableUSLocalAccountIdentification) Set(val *USLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableUSLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableUSLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUSLocalAccountIdentification(val *USLocalAccountIdentification) *NullableUSLocalAccountIdentification {
	return &NullableUSLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableUSLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUSLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *USLocalAccountIdentification) isValidAccountType() bool {
    var allowedEnumValues = []string{ "checking", "savings" }
    for _, allowed := range allowedEnumValues {
        if o.GetAccountType() == allowed {
            return true
        }
    }
    return false
}
func (o *USLocalAccountIdentification) isValidType() bool {
    var allowedEnumValues = []string{ "usLocal" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

