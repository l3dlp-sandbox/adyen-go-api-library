/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the Connectivity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Connectivity{}

// Connectivity struct for Connectivity
type Connectivity struct {
	// Indicates the status of the SIM card in the payment terminal. Can be updated and received only at terminal level, and only for models that support cellular connectivity.  Possible values: * **ACTIVATED**: the SIM card is activated. Cellular connectivity may still need to be enabled on the terminal itself, in the **Network** settings. * **INVENTORY**: the SIM card is not activated. The terminal can't use cellular connectivity.
	SimcardStatus *string `json:"simcardStatus,omitempty"`
}

// NewConnectivity instantiates a new Connectivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectivity() *Connectivity {
	this := Connectivity{}
	return &this
}

// NewConnectivityWithDefaults instantiates a new Connectivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectivityWithDefaults() *Connectivity {
	this := Connectivity{}
	return &this
}

// GetSimcardStatus returns the SimcardStatus field value if set, zero value otherwise.
func (o *Connectivity) GetSimcardStatus() string {
	if o == nil || common.IsNil(o.SimcardStatus) {
		var ret string
		return ret
	}
	return *o.SimcardStatus
}

// GetSimcardStatusOk returns a tuple with the SimcardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connectivity) GetSimcardStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.SimcardStatus) {
		return nil, false
	}
	return o.SimcardStatus, true
}

// HasSimcardStatus returns a boolean if a field has been set.
func (o *Connectivity) HasSimcardStatus() bool {
	if o != nil && !common.IsNil(o.SimcardStatus) {
		return true
	}

	return false
}

// SetSimcardStatus gets a reference to the given string and assigns it to the SimcardStatus field.
func (o *Connectivity) SetSimcardStatus(v string) {
	o.SimcardStatus = &v
}

func (o Connectivity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connectivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SimcardStatus) {
		toSerialize["simcardStatus"] = o.SimcardStatus
	}
	return toSerialize, nil
}

type NullableConnectivity struct {
	value *Connectivity
	isSet bool
}

func (v NullableConnectivity) Get() *Connectivity {
	return v.value
}

func (v *NullableConnectivity) Set(val *Connectivity) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectivity) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectivity(val *Connectivity) *NullableConnectivity {
	return &NullableConnectivity{value: val, isSet: true}
}

func (v NullableConnectivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Connectivity) isValidSimcardStatus() bool {
	var allowedEnumValues = []string{"ACTIVATED", "INVENTORY"}
	for _, allowed := range allowedEnumValues {
		if o.GetSimcardStatus() == allowed {
			return true
		}
	}
	return false
}
