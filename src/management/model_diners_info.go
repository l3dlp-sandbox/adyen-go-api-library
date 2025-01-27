/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the DinersInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DinersInfo{}

// DinersInfo struct for DinersInfo
type DinersInfo struct {
	// MID (Merchant ID) number. Required for merchants operating in Japan. Format: 14 numeric characters.
	MidNumber *string `json:"midNumber,omitempty"`
	// Indicates whether the JCB Merchant ID is reused from a previously configured JCB payment method. The default value is **false**. For merchants operating in Japan, this field is required and must be set to **true**.
	ReuseMidNumber bool `json:"reuseMidNumber"`
	// Specifies the service level (settlement type) of this payment method. Required for merchants operating in Japan. Possible values:  * **noContract**: Adyen holds the contract with JCB.  * **gatewayContract**: JCB receives the settlement and handles disputes, then pays out to you or your sub-merchant directly.
	ServiceLevel           *string                     `json:"serviceLevel,omitempty"`
	TransactionDescription *TransactionDescriptionInfo `json:"transactionDescription,omitempty"`
}

// NewDinersInfo instantiates a new DinersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDinersInfo(reuseMidNumber bool) *DinersInfo {
	this := DinersInfo{}
	this.ReuseMidNumber = reuseMidNumber
	return &this
}

// NewDinersInfoWithDefaults instantiates a new DinersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDinersInfoWithDefaults() *DinersInfo {
	this := DinersInfo{}
	var reuseMidNumber bool = false
	this.ReuseMidNumber = reuseMidNumber
	return &this
}

// GetMidNumber returns the MidNumber field value if set, zero value otherwise.
func (o *DinersInfo) GetMidNumber() string {
	if o == nil || common.IsNil(o.MidNumber) {
		var ret string
		return ret
	}
	return *o.MidNumber
}

// GetMidNumberOk returns a tuple with the MidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DinersInfo) GetMidNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.MidNumber) {
		return nil, false
	}
	return o.MidNumber, true
}

// HasMidNumber returns a boolean if a field has been set.
func (o *DinersInfo) HasMidNumber() bool {
	if o != nil && !common.IsNil(o.MidNumber) {
		return true
	}

	return false
}

// SetMidNumber gets a reference to the given string and assigns it to the MidNumber field.
func (o *DinersInfo) SetMidNumber(v string) {
	o.MidNumber = &v
}

// GetReuseMidNumber returns the ReuseMidNumber field value
func (o *DinersInfo) GetReuseMidNumber() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReuseMidNumber
}

// GetReuseMidNumberOk returns a tuple with the ReuseMidNumber field value
// and a boolean to check if the value has been set.
func (o *DinersInfo) GetReuseMidNumberOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReuseMidNumber, true
}

// SetReuseMidNumber sets field value
func (o *DinersInfo) SetReuseMidNumber(v bool) {
	o.ReuseMidNumber = v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *DinersInfo) GetServiceLevel() string {
	if o == nil || common.IsNil(o.ServiceLevel) {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DinersInfo) GetServiceLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.ServiceLevel) {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *DinersInfo) HasServiceLevel() bool {
	if o != nil && !common.IsNil(o.ServiceLevel) {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *DinersInfo) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *DinersInfo) GetTransactionDescription() TransactionDescriptionInfo {
	if o == nil || common.IsNil(o.TransactionDescription) {
		var ret TransactionDescriptionInfo
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DinersInfo) GetTransactionDescriptionOk() (*TransactionDescriptionInfo, bool) {
	if o == nil || common.IsNil(o.TransactionDescription) {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *DinersInfo) HasTransactionDescription() bool {
	if o != nil && !common.IsNil(o.TransactionDescription) {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given TransactionDescriptionInfo and assigns it to the TransactionDescription field.
func (o *DinersInfo) SetTransactionDescription(v TransactionDescriptionInfo) {
	o.TransactionDescription = &v
}

func (o DinersInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DinersInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MidNumber) {
		toSerialize["midNumber"] = o.MidNumber
	}
	toSerialize["reuseMidNumber"] = o.ReuseMidNumber
	if !common.IsNil(o.ServiceLevel) {
		toSerialize["serviceLevel"] = o.ServiceLevel
	}
	if !common.IsNil(o.TransactionDescription) {
		toSerialize["transactionDescription"] = o.TransactionDescription
	}
	return toSerialize, nil
}

type NullableDinersInfo struct {
	value *DinersInfo
	isSet bool
}

func (v NullableDinersInfo) Get() *DinersInfo {
	return v.value
}

func (v *NullableDinersInfo) Set(val *DinersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDinersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDinersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDinersInfo(val *DinersInfo) *NullableDinersInfo {
	return &NullableDinersInfo{value: val, isSet: true}
}

func (v NullableDinersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDinersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DinersInfo) isValidServiceLevel() bool {
	var allowedEnumValues = []string{"noContract", "gatewayContract"}
	for _, allowed := range allowedEnumValues {
		if o.GetServiceLevel() == allowed {
			return true
		}
	}
	return false
}
