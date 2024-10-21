/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the TerminalSettingsData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalSettingsData{}

// TerminalSettingsData struct for TerminalSettingsData
type TerminalSettingsData struct {
	// The unique identifier of the company account.
	CompanyId *string `json:"companyId,omitempty"`
	// The unique identifier of the merchant account.
	MerchantId *string `json:"merchantId,omitempty"`
	// The unique identifier of the store.
	StoreId *string `json:"storeId,omitempty"`
	// The unique identifier of the terminal.
	TerminalId *string `json:"terminalId,omitempty"`
	// Indicates whether the terminal settings were updated using the Customer Area or the Management API.
	UpdateSource string `json:"updateSource"`
	// The user that updated the terminal settings. Can be Adyen or your API credential username.
	User string `json:"user"`
}

// NewTerminalSettingsData instantiates a new TerminalSettingsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalSettingsData(updateSource string, user string) *TerminalSettingsData {
	this := TerminalSettingsData{}
	this.UpdateSource = updateSource
	this.User = user
	return &this
}

// NewTerminalSettingsDataWithDefaults instantiates a new TerminalSettingsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalSettingsDataWithDefaults() *TerminalSettingsData {
	this := TerminalSettingsData{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *TerminalSettingsData) GetCompanyId() string {
	if o == nil || common.IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettingsData) GetCompanyIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *TerminalSettingsData) HasCompanyId() bool {
	if o != nil && !common.IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *TerminalSettingsData) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *TerminalSettingsData) GetMerchantId() string {
	if o == nil || common.IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettingsData) GetMerchantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *TerminalSettingsData) HasMerchantId() bool {
	if o != nil && !common.IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *TerminalSettingsData) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *TerminalSettingsData) GetStoreId() string {
	if o == nil || common.IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettingsData) GetStoreIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *TerminalSettingsData) HasStoreId() bool {
	if o != nil && !common.IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *TerminalSettingsData) SetStoreId(v string) {
	o.StoreId = &v
}

// GetTerminalId returns the TerminalId field value if set, zero value otherwise.
func (o *TerminalSettingsData) GetTerminalId() string {
	if o == nil || common.IsNil(o.TerminalId) {
		var ret string
		return ret
	}
	return *o.TerminalId
}

// GetTerminalIdOk returns a tuple with the TerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettingsData) GetTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TerminalId) {
		return nil, false
	}
	return o.TerminalId, true
}

// HasTerminalId returns a boolean if a field has been set.
func (o *TerminalSettingsData) HasTerminalId() bool {
	if o != nil && !common.IsNil(o.TerminalId) {
		return true
	}

	return false
}

// SetTerminalId gets a reference to the given string and assigns it to the TerminalId field.
func (o *TerminalSettingsData) SetTerminalId(v string) {
	o.TerminalId = &v
}

// GetUpdateSource returns the UpdateSource field value
func (o *TerminalSettingsData) GetUpdateSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateSource
}

// GetUpdateSourceOk returns a tuple with the UpdateSource field value
// and a boolean to check if the value has been set.
func (o *TerminalSettingsData) GetUpdateSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateSource, true
}

// SetUpdateSource sets field value
func (o *TerminalSettingsData) SetUpdateSource(v string) {
	o.UpdateSource = v
}

// GetUser returns the User field value
func (o *TerminalSettingsData) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TerminalSettingsData) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TerminalSettingsData) SetUser(v string) {
	o.User = v
}

func (o TerminalSettingsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalSettingsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CompanyId) {
		toSerialize["companyId"] = o.CompanyId
	}
	if !common.IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !common.IsNil(o.StoreId) {
		toSerialize["storeId"] = o.StoreId
	}
	if !common.IsNil(o.TerminalId) {
		toSerialize["terminalId"] = o.TerminalId
	}
	toSerialize["updateSource"] = o.UpdateSource
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableTerminalSettingsData struct {
	value *TerminalSettingsData
	isSet bool
}

func (v NullableTerminalSettingsData) Get() *TerminalSettingsData {
	return v.value
}

func (v *NullableTerminalSettingsData) Set(val *TerminalSettingsData) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalSettingsData) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalSettingsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalSettingsData(val *TerminalSettingsData) *NullableTerminalSettingsData {
	return &NullableTerminalSettingsData{value: val, isSet: true}
}

func (v NullableTerminalSettingsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalSettingsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TerminalSettingsData) isValidUpdateSource() bool {
	var allowedEnumValues = []string{"Customer Area", "Management Api"}
	for _, allowed := range allowedEnumValues {
		if o.GetUpdateSource() == allowed {
			return true
		}
	}
	return false
}
