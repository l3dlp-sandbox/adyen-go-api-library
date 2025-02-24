/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the AdditionalDataOpi type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataOpi{}

// AdditionalDataOpi struct for AdditionalDataOpi
type AdditionalDataOpi struct {
	// Optional boolean indicator. Set to **true** if you want an ecommerce transaction to return an `opi.transToken` as additional data in the response.  You can store this Oracle Payment Interface token in your Oracle Opera database. For more information and required settings, see [Oracle Opera](https://docs.adyen.com/plugins/oracle-opera#opi-token-ecommerce).
	OpiIncludeTransToken *string `json:"opi.includeTransToken,omitempty"`
}

// NewAdditionalDataOpi instantiates a new AdditionalDataOpi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataOpi() *AdditionalDataOpi {
	this := AdditionalDataOpi{}
	return &this
}

// NewAdditionalDataOpiWithDefaults instantiates a new AdditionalDataOpi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataOpiWithDefaults() *AdditionalDataOpi {
	this := AdditionalDataOpi{}
	return &this
}

// GetOpiIncludeTransToken returns the OpiIncludeTransToken field value if set, zero value otherwise.
func (o *AdditionalDataOpi) GetOpiIncludeTransToken() string {
	if o == nil || common.IsNil(o.OpiIncludeTransToken) {
		var ret string
		return ret
	}
	return *o.OpiIncludeTransToken
}

// GetOpiIncludeTransTokenOk returns a tuple with the OpiIncludeTransToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpi) GetOpiIncludeTransTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpiIncludeTransToken) {
		return nil, false
	}
	return o.OpiIncludeTransToken, true
}

// HasOpiIncludeTransToken returns a boolean if a field has been set.
func (o *AdditionalDataOpi) HasOpiIncludeTransToken() bool {
	if o != nil && !common.IsNil(o.OpiIncludeTransToken) {
		return true
	}

	return false
}

// SetOpiIncludeTransToken gets a reference to the given string and assigns it to the OpiIncludeTransToken field.
func (o *AdditionalDataOpi) SetOpiIncludeTransToken(v string) {
	o.OpiIncludeTransToken = &v
}

func (o AdditionalDataOpi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataOpi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OpiIncludeTransToken) {
		toSerialize["opi.includeTransToken"] = o.OpiIncludeTransToken
	}
	return toSerialize, nil
}

type NullableAdditionalDataOpi struct {
	value *AdditionalDataOpi
	isSet bool
}

func (v NullableAdditionalDataOpi) Get() *AdditionalDataOpi {
	return v.value
}

func (v *NullableAdditionalDataOpi) Set(val *AdditionalDataOpi) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataOpi) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataOpi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataOpi(val *AdditionalDataOpi) *NullableAdditionalDataOpi {
	return &NullableAdditionalDataOpi{value: val, isSet: true}
}

func (v NullableAdditionalDataOpi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataOpi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
