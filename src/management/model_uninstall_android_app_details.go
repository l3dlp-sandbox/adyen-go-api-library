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

// checks if the UninstallAndroidAppDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UninstallAndroidAppDetails{}

// UninstallAndroidAppDetails struct for UninstallAndroidAppDetails
type UninstallAndroidAppDetails struct {
	// The unique identifier of the app to be uninstalled.
	AppId *string `json:"appId,omitempty"`
	// Type of terminal action: Uninstall an Android app.
	Type *string `json:"type,omitempty"`
}

// NewUninstallAndroidAppDetails instantiates a new UninstallAndroidAppDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUninstallAndroidAppDetails() *UninstallAndroidAppDetails {
	this := UninstallAndroidAppDetails{}
	var type_ string = "UninstallAndroidApp"
	this.Type = &type_
	return &this
}

// NewUninstallAndroidAppDetailsWithDefaults instantiates a new UninstallAndroidAppDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUninstallAndroidAppDetailsWithDefaults() *UninstallAndroidAppDetails {
	this := UninstallAndroidAppDetails{}
	var type_ string = "UninstallAndroidApp"
	this.Type = &type_
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UninstallAndroidAppDetails) GetAppId() string {
	if o == nil || common.IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UninstallAndroidAppDetails) GetAppIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UninstallAndroidAppDetails) HasAppId() bool {
	if o != nil && !common.IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *UninstallAndroidAppDetails) SetAppId(v string) {
	o.AppId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UninstallAndroidAppDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UninstallAndroidAppDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UninstallAndroidAppDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UninstallAndroidAppDetails) SetType(v string) {
	o.Type = &v
}

func (o UninstallAndroidAppDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UninstallAndroidAppDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUninstallAndroidAppDetails struct {
	value *UninstallAndroidAppDetails
	isSet bool
}

func (v NullableUninstallAndroidAppDetails) Get() *UninstallAndroidAppDetails {
	return v.value
}

func (v *NullableUninstallAndroidAppDetails) Set(val *UninstallAndroidAppDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUninstallAndroidAppDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUninstallAndroidAppDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUninstallAndroidAppDetails(val *UninstallAndroidAppDetails) *NullableUninstallAndroidAppDetails {
	return &NullableUninstallAndroidAppDetails{value: val, isSet: true}
}

func (v NullableUninstallAndroidAppDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUninstallAndroidAppDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UninstallAndroidAppDetails) isValidType() bool {
	var allowedEnumValues = []string{"UninstallAndroidApp"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
