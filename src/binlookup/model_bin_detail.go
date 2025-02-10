/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the BinDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BinDetail{}

// BinDetail struct for BinDetail
type BinDetail struct {
	// The country where the card was issued.
	IssuerCountry *string `json:"issuerCountry,omitempty"`
}

// NewBinDetail instantiates a new BinDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinDetail() *BinDetail {
	this := BinDetail{}
	return &this
}

// NewBinDetailWithDefaults instantiates a new BinDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinDetailWithDefaults() *BinDetail {
	this := BinDetail{}
	return &this
}

// GetIssuerCountry returns the IssuerCountry field value if set, zero value otherwise.
func (o *BinDetail) GetIssuerCountry() string {
	if o == nil || common.IsNil(o.IssuerCountry) {
		var ret string
		return ret
	}
	return *o.IssuerCountry
}

// GetIssuerCountryOk returns a tuple with the IssuerCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinDetail) GetIssuerCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerCountry) {
		return nil, false
	}
	return o.IssuerCountry, true
}

// HasIssuerCountry returns a boolean if a field has been set.
func (o *BinDetail) HasIssuerCountry() bool {
	if o != nil && !common.IsNil(o.IssuerCountry) {
		return true
	}

	return false
}

// SetIssuerCountry gets a reference to the given string and assigns it to the IssuerCountry field.
func (o *BinDetail) SetIssuerCountry(v string) {
	o.IssuerCountry = &v
}

func (o BinDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BinDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IssuerCountry) {
		toSerialize["issuerCountry"] = o.IssuerCountry
	}
	return toSerialize, nil
}

type NullableBinDetail struct {
	value *BinDetail
	isSet bool
}

func (v NullableBinDetail) Get() *BinDetail {
	return v.value
}

func (v *NullableBinDetail) Set(val *BinDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBinDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBinDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinDetail(val *BinDetail) *NullableBinDetail {
	return &NullableBinDetail{value: val, isSet: true}
}

func (v NullableBinDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
