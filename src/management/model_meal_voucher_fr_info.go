/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the MealVoucherFRInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MealVoucherFRInfo{}

// MealVoucherFRInfo struct for MealVoucherFRInfo
type MealVoucherFRInfo struct {
	// Meal Voucher conecsId. Format: digits only
	ConecsId string `json:"conecsId"`
	// Meal Voucher siret. Format: 14 digits.
	Siret string `json:"siret"`
	// The list of additional payment methods. Allowed values: **mealVoucher_FR_edenred**, **mealVoucher_FR_groupeup**, **mealVoucher_FR_natixis**, **mealVoucher_FR_sodexo**.
	SubTypes []string `json:"subTypes"`
}

// NewMealVoucherFRInfo instantiates a new MealVoucherFRInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMealVoucherFRInfo(conecsId string, siret string, subTypes []string) *MealVoucherFRInfo {
	this := MealVoucherFRInfo{}
	this.ConecsId = conecsId
	this.Siret = siret
	this.SubTypes = subTypes
	return &this
}

// NewMealVoucherFRInfoWithDefaults instantiates a new MealVoucherFRInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMealVoucherFRInfoWithDefaults() *MealVoucherFRInfo {
	this := MealVoucherFRInfo{}
	return &this
}

// GetConecsId returns the ConecsId field value
func (o *MealVoucherFRInfo) GetConecsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConecsId
}

// GetConecsIdOk returns a tuple with the ConecsId field value
// and a boolean to check if the value has been set.
func (o *MealVoucherFRInfo) GetConecsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConecsId, true
}

// SetConecsId sets field value
func (o *MealVoucherFRInfo) SetConecsId(v string) {
	o.ConecsId = v
}

// GetSiret returns the Siret field value
func (o *MealVoucherFRInfo) GetSiret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Siret
}

// GetSiretOk returns a tuple with the Siret field value
// and a boolean to check if the value has been set.
func (o *MealVoucherFRInfo) GetSiretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Siret, true
}

// SetSiret sets field value
func (o *MealVoucherFRInfo) SetSiret(v string) {
	o.Siret = v
}

// GetSubTypes returns the SubTypes field value
func (o *MealVoucherFRInfo) GetSubTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubTypes
}

// GetSubTypesOk returns a tuple with the SubTypes field value
// and a boolean to check if the value has been set.
func (o *MealVoucherFRInfo) GetSubTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubTypes, true
}

// SetSubTypes sets field value
func (o *MealVoucherFRInfo) SetSubTypes(v []string) {
	o.SubTypes = v
}

func (o MealVoucherFRInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MealVoucherFRInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conecsId"] = o.ConecsId
	toSerialize["siret"] = o.Siret
	toSerialize["subTypes"] = o.SubTypes
	return toSerialize, nil
}

type NullableMealVoucherFRInfo struct {
	value *MealVoucherFRInfo
	isSet bool
}

func (v NullableMealVoucherFRInfo) Get() *MealVoucherFRInfo {
	return v.value
}

func (v *NullableMealVoucherFRInfo) Set(val *MealVoucherFRInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMealVoucherFRInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMealVoucherFRInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMealVoucherFRInfo(val *MealVoucherFRInfo) *NullableMealVoucherFRInfo {
	return &NullableMealVoucherFRInfo{value: val, isSet: true}
}

func (v NullableMealVoucherFRInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMealVoucherFRInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
