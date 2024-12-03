/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the DeviceInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeviceInfo{}

// DeviceInfo struct for DeviceInfo
type DeviceInfo struct {
	// The technology used to capture the card details.
	CardCaptureTechnology *string `json:"cardCaptureTechnology,omitempty"`
	// The name of the device.
	DeviceName *string `json:"deviceName,omitempty"`
	// The form factor of the device to be provisioned.
	FormFactor *string `json:"formFactor,omitempty"`
	// The IMEI number of the device being provisioned.
	Imei *string `json:"imei,omitempty"`
	// The 2-digit device type provided on the ISO messages that the token is being provisioned to.
	IsoDeviceType *string `json:"isoDeviceType,omitempty"`
	// The MSISDN of the device being provisioned.
	Msisdn *string `json:"msisdn,omitempty"`
	// The name of the device operating system.
	OsName *string `json:"osName,omitempty"`
	// The version of the device operating system.
	OsVersion *string `json:"osVersion,omitempty"`
	// Different types of payments supported for the network token.
	PaymentTypes []string `json:"paymentTypes,omitempty"`
	// The serial number of the device.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// The architecture or technology used for network token storage.
	StorageTechnology *string `json:"storageTechnology,omitempty"`
}

// NewDeviceInfo instantiates a new DeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInfo() *DeviceInfo {
	this := DeviceInfo{}
	return &this
}

// NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInfoWithDefaults() *DeviceInfo {
	this := DeviceInfo{}
	return &this
}

// GetCardCaptureTechnology returns the CardCaptureTechnology field value if set, zero value otherwise.
func (o *DeviceInfo) GetCardCaptureTechnology() string {
	if o == nil || common.IsNil(o.CardCaptureTechnology) {
		var ret string
		return ret
	}
	return *o.CardCaptureTechnology
}

// GetCardCaptureTechnologyOk returns a tuple with the CardCaptureTechnology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetCardCaptureTechnologyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardCaptureTechnology) {
		return nil, false
	}
	return o.CardCaptureTechnology, true
}

// HasCardCaptureTechnology returns a boolean if a field has been set.
func (o *DeviceInfo) HasCardCaptureTechnology() bool {
	if o != nil && !common.IsNil(o.CardCaptureTechnology) {
		return true
	}

	return false
}

// SetCardCaptureTechnology gets a reference to the given string and assigns it to the CardCaptureTechnology field.
func (o *DeviceInfo) SetCardCaptureTechnology(v string) {
	o.CardCaptureTechnology = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *DeviceInfo) GetDeviceName() string {
	if o == nil || common.IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetDeviceNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *DeviceInfo) HasDeviceName() bool {
	if o != nil && !common.IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *DeviceInfo) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise.
func (o *DeviceInfo) GetFormFactor() string {
	if o == nil || common.IsNil(o.FormFactor) {
		var ret string
		return ret
	}
	return *o.FormFactor
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetFormFactorOk() (*string, bool) {
	if o == nil || common.IsNil(o.FormFactor) {
		return nil, false
	}
	return o.FormFactor, true
}

// HasFormFactor returns a boolean if a field has been set.
func (o *DeviceInfo) HasFormFactor() bool {
	if o != nil && !common.IsNil(o.FormFactor) {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given string and assigns it to the FormFactor field.
func (o *DeviceInfo) SetFormFactor(v string) {
	o.FormFactor = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *DeviceInfo) GetImei() string {
	if o == nil || common.IsNil(o.Imei) {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetImeiOk() (*string, bool) {
	if o == nil || common.IsNil(o.Imei) {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *DeviceInfo) HasImei() bool {
	if o != nil && !common.IsNil(o.Imei) {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *DeviceInfo) SetImei(v string) {
	o.Imei = &v
}

// GetIsoDeviceType returns the IsoDeviceType field value if set, zero value otherwise.
func (o *DeviceInfo) GetIsoDeviceType() string {
	if o == nil || common.IsNil(o.IsoDeviceType) {
		var ret string
		return ret
	}
	return *o.IsoDeviceType
}

// GetIsoDeviceTypeOk returns a tuple with the IsoDeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetIsoDeviceTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsoDeviceType) {
		return nil, false
	}
	return o.IsoDeviceType, true
}

// HasIsoDeviceType returns a boolean if a field has been set.
func (o *DeviceInfo) HasIsoDeviceType() bool {
	if o != nil && !common.IsNil(o.IsoDeviceType) {
		return true
	}

	return false
}

// SetIsoDeviceType gets a reference to the given string and assigns it to the IsoDeviceType field.
func (o *DeviceInfo) SetIsoDeviceType(v string) {
	o.IsoDeviceType = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *DeviceInfo) GetMsisdn() string {
	if o == nil || common.IsNil(o.Msisdn) {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetMsisdnOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msisdn) {
		return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *DeviceInfo) HasMsisdn() bool {
	if o != nil && !common.IsNil(o.Msisdn) {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *DeviceInfo) SetMsisdn(v string) {
	o.Msisdn = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *DeviceInfo) GetOsName() string {
	if o == nil || common.IsNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetOsNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OsName) {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *DeviceInfo) HasOsName() bool {
	if o != nil && !common.IsNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *DeviceInfo) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceInfo) GetOsVersion() string {
	if o == nil || common.IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetOsVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceInfo) HasOsVersion() bool {
	if o != nil && !common.IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *DeviceInfo) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPaymentTypes returns the PaymentTypes field value if set, zero value otherwise.
func (o *DeviceInfo) GetPaymentTypes() []string {
	if o == nil || common.IsNil(o.PaymentTypes) {
		var ret []string
		return ret
	}
	return o.PaymentTypes
}

// GetPaymentTypesOk returns a tuple with the PaymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetPaymentTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.PaymentTypes) {
		return nil, false
	}
	return o.PaymentTypes, true
}

// HasPaymentTypes returns a boolean if a field has been set.
func (o *DeviceInfo) HasPaymentTypes() bool {
	if o != nil && !common.IsNil(o.PaymentTypes) {
		return true
	}

	return false
}

// SetPaymentTypes gets a reference to the given []string and assigns it to the PaymentTypes field.
func (o *DeviceInfo) SetPaymentTypes(v []string) {
	o.PaymentTypes = v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DeviceInfo) GetSerialNumber() string {
	if o == nil || common.IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetSerialNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DeviceInfo) HasSerialNumber() bool {
	if o != nil && !common.IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *DeviceInfo) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStorageTechnology returns the StorageTechnology field value if set, zero value otherwise.
func (o *DeviceInfo) GetStorageTechnology() string {
	if o == nil || common.IsNil(o.StorageTechnology) {
		var ret string
		return ret
	}
	return *o.StorageTechnology
}

// GetStorageTechnologyOk returns a tuple with the StorageTechnology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetStorageTechnologyOk() (*string, bool) {
	if o == nil || common.IsNil(o.StorageTechnology) {
		return nil, false
	}
	return o.StorageTechnology, true
}

// HasStorageTechnology returns a boolean if a field has been set.
func (o *DeviceInfo) HasStorageTechnology() bool {
	if o != nil && !common.IsNil(o.StorageTechnology) {
		return true
	}

	return false
}

// SetStorageTechnology gets a reference to the given string and assigns it to the StorageTechnology field.
func (o *DeviceInfo) SetStorageTechnology(v string) {
	o.StorageTechnology = &v
}

func (o DeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CardCaptureTechnology) {
		toSerialize["cardCaptureTechnology"] = o.CardCaptureTechnology
	}
	if !common.IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !common.IsNil(o.FormFactor) {
		toSerialize["formFactor"] = o.FormFactor
	}
	if !common.IsNil(o.Imei) {
		toSerialize["imei"] = o.Imei
	}
	if !common.IsNil(o.IsoDeviceType) {
		toSerialize["isoDeviceType"] = o.IsoDeviceType
	}
	if !common.IsNil(o.Msisdn) {
		toSerialize["msisdn"] = o.Msisdn
	}
	if !common.IsNil(o.OsName) {
		toSerialize["osName"] = o.OsName
	}
	if !common.IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !common.IsNil(o.PaymentTypes) {
		toSerialize["paymentTypes"] = o.PaymentTypes
	}
	if !common.IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !common.IsNil(o.StorageTechnology) {
		toSerialize["storageTechnology"] = o.StorageTechnology
	}
	return toSerialize, nil
}

type NullableDeviceInfo struct {
	value *DeviceInfo
	isSet bool
}

func (v NullableDeviceInfo) Get() *DeviceInfo {
	return v.value
}

func (v *NullableDeviceInfo) Set(val *DeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInfo(val *DeviceInfo) *NullableDeviceInfo {
	return &NullableDeviceInfo{value: val, isSet: true}
}

func (v NullableDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



