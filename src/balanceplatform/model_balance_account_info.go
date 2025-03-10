/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the BalanceAccountInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceAccountInfo{}

// BalanceAccountInfo struct for BalanceAccountInfo
type BalanceAccountInfo struct {
	// The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account.
	AccountHolderId string `json:"accountHolderId"`
	// The default three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance account. The default value is **EUR**. > After a balance account is created, you cannot change its default currency.
	DefaultCurrencyCode *string `json:"defaultCurrencyCode,omitempty"`
	// A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder.
	Description *string `json:"description,omitempty"`
	// A set of key and value pairs for general use. The keys do not have specific names and may be used for storing miscellaneous data as desired. > Note that during an update of metadata, the omission of existing key-value pairs will result in the deletion of those key-value pairs.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The unique identifier of the account of the migrated account holder in the classic integration.
	MigratedAccountCode          *string                       `json:"migratedAccountCode,omitempty"`
	PlatformPaymentConfiguration *PlatformPaymentConfiguration `json:"platformPaymentConfiguration,omitempty"`
	// Your reference for the balance account, maximum 150 characters.
	Reference *string `json:"reference,omitempty"`
	// The time zone of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewBalanceAccountInfo instantiates a new BalanceAccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceAccountInfo(accountHolderId string) *BalanceAccountInfo {
	this := BalanceAccountInfo{}
	this.AccountHolderId = accountHolderId
	return &this
}

// NewBalanceAccountInfoWithDefaults instantiates a new BalanceAccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceAccountInfoWithDefaults() *BalanceAccountInfo {
	this := BalanceAccountInfo{}
	return &this
}

// GetAccountHolderId returns the AccountHolderId field value
func (o *BalanceAccountInfo) GetAccountHolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountHolderId
}

// GetAccountHolderIdOk returns a tuple with the AccountHolderId field value
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetAccountHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolderId, true
}

// SetAccountHolderId sets field value
func (o *BalanceAccountInfo) SetAccountHolderId(v string) {
	o.AccountHolderId = v
}

// GetDefaultCurrencyCode returns the DefaultCurrencyCode field value if set, zero value otherwise.
func (o *BalanceAccountInfo) GetDefaultCurrencyCode() string {
	if o == nil || common.IsNil(o.DefaultCurrencyCode) {
		var ret string
		return ret
	}
	return *o.DefaultCurrencyCode
}

// GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetDefaultCurrencyCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.DefaultCurrencyCode) {
		return nil, false
	}
	return o.DefaultCurrencyCode, true
}

// HasDefaultCurrencyCode returns a boolean if a field has been set.
func (o *BalanceAccountInfo) HasDefaultCurrencyCode() bool {
	if o != nil && !common.IsNil(o.DefaultCurrencyCode) {
		return true
	}

	return false
}

// SetDefaultCurrencyCode gets a reference to the given string and assigns it to the DefaultCurrencyCode field.
func (o *BalanceAccountInfo) SetDefaultCurrencyCode(v string) {
	o.DefaultCurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BalanceAccountInfo) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BalanceAccountInfo) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BalanceAccountInfo) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BalanceAccountInfo) GetMetadata() map[string]string {
	if o == nil || common.IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BalanceAccountInfo) HasMetadata() bool {
	if o != nil && !common.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *BalanceAccountInfo) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMigratedAccountCode returns the MigratedAccountCode field value if set, zero value otherwise.
func (o *BalanceAccountInfo) GetMigratedAccountCode() string {
	if o == nil || common.IsNil(o.MigratedAccountCode) {
		var ret string
		return ret
	}
	return *o.MigratedAccountCode
}

// GetMigratedAccountCodeOk returns a tuple with the MigratedAccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetMigratedAccountCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MigratedAccountCode) {
		return nil, false
	}
	return o.MigratedAccountCode, true
}

// HasMigratedAccountCode returns a boolean if a field has been set.
func (o *BalanceAccountInfo) HasMigratedAccountCode() bool {
	if o != nil && !common.IsNil(o.MigratedAccountCode) {
		return true
	}

	return false
}

// SetMigratedAccountCode gets a reference to the given string and assigns it to the MigratedAccountCode field.
func (o *BalanceAccountInfo) SetMigratedAccountCode(v string) {
	o.MigratedAccountCode = &v
}

// GetPlatformPaymentConfiguration returns the PlatformPaymentConfiguration field value if set, zero value otherwise.
func (o *BalanceAccountInfo) GetPlatformPaymentConfiguration() PlatformPaymentConfiguration {
	if o == nil || common.IsNil(o.PlatformPaymentConfiguration) {
		var ret PlatformPaymentConfiguration
		return ret
	}
	return *o.PlatformPaymentConfiguration
}

// GetPlatformPaymentConfigurationOk returns a tuple with the PlatformPaymentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetPlatformPaymentConfigurationOk() (*PlatformPaymentConfiguration, bool) {
	if o == nil || common.IsNil(o.PlatformPaymentConfiguration) {
		return nil, false
	}
	return o.PlatformPaymentConfiguration, true
}

// HasPlatformPaymentConfiguration returns a boolean if a field has been set.
func (o *BalanceAccountInfo) HasPlatformPaymentConfiguration() bool {
	if o != nil && !common.IsNil(o.PlatformPaymentConfiguration) {
		return true
	}

	return false
}

// SetPlatformPaymentConfiguration gets a reference to the given PlatformPaymentConfiguration and assigns it to the PlatformPaymentConfiguration field.
func (o *BalanceAccountInfo) SetPlatformPaymentConfiguration(v PlatformPaymentConfiguration) {
	o.PlatformPaymentConfiguration = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *BalanceAccountInfo) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *BalanceAccountInfo) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *BalanceAccountInfo) SetReference(v string) {
	o.Reference = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *BalanceAccountInfo) GetTimeZone() string {
	if o == nil || common.IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountInfo) GetTimeZoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *BalanceAccountInfo) HasTimeZone() bool {
	if o != nil && !common.IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *BalanceAccountInfo) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o BalanceAccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceAccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountHolderId"] = o.AccountHolderId
	if !common.IsNil(o.DefaultCurrencyCode) {
		toSerialize["defaultCurrencyCode"] = o.DefaultCurrencyCode
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !common.IsNil(o.MigratedAccountCode) {
		toSerialize["migratedAccountCode"] = o.MigratedAccountCode
	}
	if !common.IsNil(o.PlatformPaymentConfiguration) {
		toSerialize["platformPaymentConfiguration"] = o.PlatformPaymentConfiguration
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableBalanceAccountInfo struct {
	value *BalanceAccountInfo
	isSet bool
}

func (v NullableBalanceAccountInfo) Get() *BalanceAccountInfo {
	return v.value
}

func (v *NullableBalanceAccountInfo) Set(val *BalanceAccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceAccountInfo(val *BalanceAccountInfo) *NullableBalanceAccountInfo {
	return &NullableBalanceAccountInfo{value: val, isSet: true}
}

func (v NullableBalanceAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
