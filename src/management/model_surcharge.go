/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Surcharge type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Surcharge{}

// Surcharge struct for Surcharge
type Surcharge struct {
	// Show the surcharge details on the terminal, so the shopper can confirm.
	AskConfirmation *bool `json:"askConfirmation,omitempty"`
	// Surcharge fees or percentages for specific cards, funding sources (credit or debit), and currencies.
	Configurations []Configuration `json:"configurations,omitempty"`
}

// NewSurcharge instantiates a new Surcharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurcharge() *Surcharge {
	this := Surcharge{}
	return &this
}

// NewSurchargeWithDefaults instantiates a new Surcharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurchargeWithDefaults() *Surcharge {
	this := Surcharge{}
	return &this
}

// GetAskConfirmation returns the AskConfirmation field value if set, zero value otherwise.
func (o *Surcharge) GetAskConfirmation() bool {
	if o == nil || common.IsNil(o.AskConfirmation) {
		var ret bool
		return ret
	}
	return *o.AskConfirmation
}

// GetAskConfirmationOk returns a tuple with the AskConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surcharge) GetAskConfirmationOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AskConfirmation) {
		return nil, false
	}
	return o.AskConfirmation, true
}

// HasAskConfirmation returns a boolean if a field has been set.
func (o *Surcharge) HasAskConfirmation() bool {
	if o != nil && !common.IsNil(o.AskConfirmation) {
		return true
	}

	return false
}

// SetAskConfirmation gets a reference to the given bool and assigns it to the AskConfirmation field.
func (o *Surcharge) SetAskConfirmation(v bool) {
	o.AskConfirmation = &v
}

// GetConfigurations returns the Configurations field value if set, zero value otherwise.
func (o *Surcharge) GetConfigurations() []Configuration {
	if o == nil || common.IsNil(o.Configurations) {
		var ret []Configuration
		return ret
	}
	return o.Configurations
}

// GetConfigurationsOk returns a tuple with the Configurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surcharge) GetConfigurationsOk() ([]Configuration, bool) {
	if o == nil || common.IsNil(o.Configurations) {
		return nil, false
	}
	return o.Configurations, true
}

// HasConfigurations returns a boolean if a field has been set.
func (o *Surcharge) HasConfigurations() bool {
	if o != nil && !common.IsNil(o.Configurations) {
		return true
	}

	return false
}

// SetConfigurations gets a reference to the given []Configuration and assigns it to the Configurations field.
func (o *Surcharge) SetConfigurations(v []Configuration) {
	o.Configurations = v
}

func (o Surcharge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Surcharge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AskConfirmation) {
		toSerialize["askConfirmation"] = o.AskConfirmation
	}
	if !common.IsNil(o.Configurations) {
		toSerialize["configurations"] = o.Configurations
	}
	return toSerialize, nil
}

type NullableSurcharge struct {
	value *Surcharge
	isSet bool
}

func (v NullableSurcharge) Get() *Surcharge {
	return v.value
}

func (v *NullableSurcharge) Set(val *Surcharge) {
	v.value = val
	v.isSet = true
}

func (v NullableSurcharge) IsSet() bool {
	return v.isSet
}

func (v *NullableSurcharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurcharge(val *Surcharge) *NullableSurcharge {
	return &NullableSurcharge{value: val, isSet: true}
}

func (v NullableSurcharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurcharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



