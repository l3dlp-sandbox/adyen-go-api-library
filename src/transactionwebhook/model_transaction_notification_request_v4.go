/*
Transaction webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactionwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TransactionNotificationRequestV4 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionNotificationRequestV4{}

// TransactionNotificationRequestV4 struct for TransactionNotificationRequestV4
type TransactionNotificationRequestV4 struct {
	Data Transaction `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of the webhook.
	Type *string `json:"type,omitempty"`
}

// NewTransactionNotificationRequestV4 instantiates a new TransactionNotificationRequestV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionNotificationRequestV4(data Transaction, environment string) *TransactionNotificationRequestV4 {
	this := TransactionNotificationRequestV4{}
	this.Data = data
	this.Environment = environment
	return &this
}

// NewTransactionNotificationRequestV4WithDefaults instantiates a new TransactionNotificationRequestV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionNotificationRequestV4WithDefaults() *TransactionNotificationRequestV4 {
	this := TransactionNotificationRequestV4{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionNotificationRequestV4) GetData() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionNotificationRequestV4) GetDataOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionNotificationRequestV4) SetData(v Transaction) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *TransactionNotificationRequestV4) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *TransactionNotificationRequestV4) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *TransactionNotificationRequestV4) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransactionNotificationRequestV4) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionNotificationRequestV4) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransactionNotificationRequestV4) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransactionNotificationRequestV4) SetType(v string) {
	o.Type = &v
}

func (o TransactionNotificationRequestV4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionNotificationRequestV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTransactionNotificationRequestV4 struct {
	value *TransactionNotificationRequestV4
	isSet bool
}

func (v NullableTransactionNotificationRequestV4) Get() *TransactionNotificationRequestV4 {
	return v.value
}

func (v *NullableTransactionNotificationRequestV4) Set(val *TransactionNotificationRequestV4) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionNotificationRequestV4) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionNotificationRequestV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionNotificationRequestV4(val *TransactionNotificationRequestV4) *NullableTransactionNotificationRequestV4 {
	return &NullableTransactionNotificationRequestV4{value: val, isSet: true}
}

func (v NullableTransactionNotificationRequestV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionNotificationRequestV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransactionNotificationRequestV4) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.transaction.created"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
