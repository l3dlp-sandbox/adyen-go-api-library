/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the ShopperStatement type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ShopperStatement{}

// ShopperStatement struct for ShopperStatement
type ShopperStatement struct {
	// The name you want to be shown on the shopper's bank or credit card statement. Can't be all numbers. If a shopper statement is present, this field is required.
	DoingBusinessAsName *string `json:"doingBusinessAsName,omitempty"`
	// The type of shopperstatement you want to use: fixed, append or dynamic
	Type *string `json:"type,omitempty"`
}

// NewShopperStatement instantiates a new ShopperStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopperStatement() *ShopperStatement {
	this := ShopperStatement{}
	var type_ string = "dynamic"
	this.Type = &type_
	return &this
}

// NewShopperStatementWithDefaults instantiates a new ShopperStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopperStatementWithDefaults() *ShopperStatement {
	this := ShopperStatement{}
	var type_ string = "dynamic"
	this.Type = &type_
	return &this
}

// GetDoingBusinessAsName returns the DoingBusinessAsName field value if set, zero value otherwise.
func (o *ShopperStatement) GetDoingBusinessAsName() string {
	if o == nil || common.IsNil(o.DoingBusinessAsName) {
		var ret string
		return ret
	}
	return *o.DoingBusinessAsName
}

// GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperStatement) GetDoingBusinessAsNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.DoingBusinessAsName) {
		return nil, false
	}
	return o.DoingBusinessAsName, true
}

// HasDoingBusinessAsName returns a boolean if a field has been set.
func (o *ShopperStatement) HasDoingBusinessAsName() bool {
	if o != nil && !common.IsNil(o.DoingBusinessAsName) {
		return true
	}

	return false
}

// SetDoingBusinessAsName gets a reference to the given string and assigns it to the DoingBusinessAsName field.
func (o *ShopperStatement) SetDoingBusinessAsName(v string) {
	o.DoingBusinessAsName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ShopperStatement) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperStatement) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ShopperStatement) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ShopperStatement) SetType(v string) {
	o.Type = &v
}

func (o ShopperStatement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopperStatement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DoingBusinessAsName) {
		toSerialize["doingBusinessAsName"] = o.DoingBusinessAsName
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableShopperStatement struct {
	value *ShopperStatement
	isSet bool
}

func (v NullableShopperStatement) Get() *ShopperStatement {
	return v.value
}

func (v *NullableShopperStatement) Set(val *ShopperStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableShopperStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableShopperStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopperStatement(val *ShopperStatement) *NullableShopperStatement {
	return &NullableShopperStatement{value: val, isSet: true}
}

func (v NullableShopperStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopperStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ShopperStatement) isValidType() bool {
	var allowedEnumValues = []string{"append", "dynamic", "fixed"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}