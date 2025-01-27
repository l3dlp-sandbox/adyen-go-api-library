/*
Payments App API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentsapp

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the PaymentsAppDto type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentsAppDto{}

// PaymentsAppDto struct for PaymentsAppDto
type PaymentsAppDto struct {
	// The unique identifier of the Payments App instance.
	InstallationId string `json:"installationId"`
	// The account code associated with the Payments App instance.
	MerchantAccountCode string `json:"merchantAccountCode"`
	// The store code associated with the Payments App instance.
	MerchantStoreCode *string `json:"merchantStoreCode,omitempty"`
	// The status of the Payments App instance.
	Status string `json:"status"`
}

// NewPaymentsAppDto instantiates a new PaymentsAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentsAppDto(installationId string, merchantAccountCode string, status string) *PaymentsAppDto {
	this := PaymentsAppDto{}
	this.InstallationId = installationId
	this.MerchantAccountCode = merchantAccountCode
	this.Status = status
	return &this
}

// NewPaymentsAppDtoWithDefaults instantiates a new PaymentsAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsAppDtoWithDefaults() *PaymentsAppDto {
	this := PaymentsAppDto{}
	return &this
}

// GetInstallationId returns the InstallationId field value
func (o *PaymentsAppDto) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *PaymentsAppDto) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *PaymentsAppDto) SetInstallationId(v string) {
	o.InstallationId = v
}

// GetMerchantAccountCode returns the MerchantAccountCode field value
func (o *PaymentsAppDto) GetMerchantAccountCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountCode
}

// GetMerchantAccountCodeOk returns a tuple with the MerchantAccountCode field value
// and a boolean to check if the value has been set.
func (o *PaymentsAppDto) GetMerchantAccountCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountCode, true
}

// SetMerchantAccountCode sets field value
func (o *PaymentsAppDto) SetMerchantAccountCode(v string) {
	o.MerchantAccountCode = v
}

// GetMerchantStoreCode returns the MerchantStoreCode field value if set, zero value otherwise.
func (o *PaymentsAppDto) GetMerchantStoreCode() string {
	if o == nil || common.IsNil(o.MerchantStoreCode) {
		var ret string
		return ret
	}
	return *o.MerchantStoreCode
}

// GetMerchantStoreCodeOk returns a tuple with the MerchantStoreCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentsAppDto) GetMerchantStoreCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantStoreCode) {
		return nil, false
	}
	return o.MerchantStoreCode, true
}

// HasMerchantStoreCode returns a boolean if a field has been set.
func (o *PaymentsAppDto) HasMerchantStoreCode() bool {
	if o != nil && !common.IsNil(o.MerchantStoreCode) {
		return true
	}

	return false
}

// SetMerchantStoreCode gets a reference to the given string and assigns it to the MerchantStoreCode field.
func (o *PaymentsAppDto) SetMerchantStoreCode(v string) {
	o.MerchantStoreCode = &v
}

// GetStatus returns the Status field value
func (o *PaymentsAppDto) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentsAppDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentsAppDto) SetStatus(v string) {
	o.Status = v
}

func (o PaymentsAppDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentsAppDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["installationId"] = o.InstallationId
	toSerialize["merchantAccountCode"] = o.MerchantAccountCode
	if !common.IsNil(o.MerchantStoreCode) {
		toSerialize["merchantStoreCode"] = o.MerchantStoreCode
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePaymentsAppDto struct {
	value *PaymentsAppDto
	isSet bool
}

func (v NullablePaymentsAppDto) Get() *PaymentsAppDto {
	return v.value
}

func (v *NullablePaymentsAppDto) Set(val *PaymentsAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentsAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentsAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentsAppDto(val *PaymentsAppDto) *NullablePaymentsAppDto {
	return &NullablePaymentsAppDto{value: val, isSet: true}
}

func (v NullablePaymentsAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentsAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
