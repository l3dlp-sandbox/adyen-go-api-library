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

// checks if the UpdateCompanyUserRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateCompanyUserRequest{}

// UpdateCompanyUserRequest struct for UpdateCompanyUserRequest
type UpdateCompanyUserRequest struct {
	// The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user.
	AccountGroups []string `json:"accountGroups,omitempty"`
	// Indicates whether this user is active.
	Active *bool `json:"active,omitempty"`
	// The list of [merchant accounts](https://docs.adyen.com/account/account-structure#merchant-accounts) to associate the user with.
	AssociatedMerchantAccounts []string `json:"associatedMerchantAccounts,omitempty"`
	// The email address of the user.
	Email *string `json:"email,omitempty"`
	// The requested login method for the user. To use SSO, you must already have SSO configured with Adyen before creating the user.  Possible values: **Username & account**, **Email**, or **SSO**
	LoginMethod *string `json:"loginMethod,omitempty"`
	Name        *Name2  `json:"name,omitempty"`
	// The list of [roles](https://docs.adyen.com/account/user-roles) for this user.
	Roles []string `json:"roles,omitempty"`
	// The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**.
	TimeZoneCode *string `json:"timeZoneCode,omitempty"`
}

// NewUpdateCompanyUserRequest instantiates a new UpdateCompanyUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCompanyUserRequest() *UpdateCompanyUserRequest {
	this := UpdateCompanyUserRequest{}
	return &this
}

// NewUpdateCompanyUserRequestWithDefaults instantiates a new UpdateCompanyUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCompanyUserRequestWithDefaults() *UpdateCompanyUserRequest {
	this := UpdateCompanyUserRequest{}
	return &this
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetAccountGroups() []string {
	if o == nil || common.IsNil(o.AccountGroups) {
		var ret []string
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetAccountGroupsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasAccountGroups() bool {
	if o != nil && !common.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []string and assigns it to the AccountGroups field.
func (o *UpdateCompanyUserRequest) SetAccountGroups(v []string) {
	o.AccountGroups = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetActive() bool {
	if o == nil || common.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasActive() bool {
	if o != nil && !common.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateCompanyUserRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetAssociatedMerchantAccounts() []string {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		var ret []string
		return ret
	}
	return o.AssociatedMerchantAccounts
}

// GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetAssociatedMerchantAccountsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		return nil, false
	}
	return o.AssociatedMerchantAccounts, true
}

// HasAssociatedMerchantAccounts returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasAssociatedMerchantAccounts() bool {
	if o != nil && !common.IsNil(o.AssociatedMerchantAccounts) {
		return true
	}

	return false
}

// SetAssociatedMerchantAccounts gets a reference to the given []string and assigns it to the AssociatedMerchantAccounts field.
func (o *UpdateCompanyUserRequest) SetAssociatedMerchantAccounts(v []string) {
	o.AssociatedMerchantAccounts = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateCompanyUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetLoginMethod returns the LoginMethod field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetLoginMethod() string {
	if o == nil || common.IsNil(o.LoginMethod) {
		var ret string
		return ret
	}
	return *o.LoginMethod
}

// GetLoginMethodOk returns a tuple with the LoginMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetLoginMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoginMethod) {
		return nil, false
	}
	return o.LoginMethod, true
}

// HasLoginMethod returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasLoginMethod() bool {
	if o != nil && !common.IsNil(o.LoginMethod) {
		return true
	}

	return false
}

// SetLoginMethod gets a reference to the given string and assigns it to the LoginMethod field.
func (o *UpdateCompanyUserRequest) SetLoginMethod(v string) {
	o.LoginMethod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetName() Name2 {
	if o == nil || common.IsNil(o.Name) {
		var ret Name2
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetNameOk() (*Name2, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given Name2 and assigns it to the Name field.
func (o *UpdateCompanyUserRequest) SetName(v Name2) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UpdateCompanyUserRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetTimeZoneCode returns the TimeZoneCode field value if set, zero value otherwise.
func (o *UpdateCompanyUserRequest) GetTimeZoneCode() string {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		var ret string
		return ret
	}
	return *o.TimeZoneCode
}

// GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyUserRequest) GetTimeZoneCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		return nil, false
	}
	return o.TimeZoneCode, true
}

// HasTimeZoneCode returns a boolean if a field has been set.
func (o *UpdateCompanyUserRequest) HasTimeZoneCode() bool {
	if o != nil && !common.IsNil(o.TimeZoneCode) {
		return true
	}

	return false
}

// SetTimeZoneCode gets a reference to the given string and assigns it to the TimeZoneCode field.
func (o *UpdateCompanyUserRequest) SetTimeZoneCode(v string) {
	o.TimeZoneCode = &v
}

func (o UpdateCompanyUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCompanyUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if !common.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !common.IsNil(o.AssociatedMerchantAccounts) {
		toSerialize["associatedMerchantAccounts"] = o.AssociatedMerchantAccounts
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.LoginMethod) {
		toSerialize["loginMethod"] = o.LoginMethod
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !common.IsNil(o.TimeZoneCode) {
		toSerialize["timeZoneCode"] = o.TimeZoneCode
	}
	return toSerialize, nil
}

type NullableUpdateCompanyUserRequest struct {
	value *UpdateCompanyUserRequest
	isSet bool
}

func (v NullableUpdateCompanyUserRequest) Get() *UpdateCompanyUserRequest {
	return v.value
}

func (v *NullableUpdateCompanyUserRequest) Set(val *UpdateCompanyUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCompanyUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCompanyUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCompanyUserRequest(val *UpdateCompanyUserRequest) *NullableUpdateCompanyUserRequest {
	return &NullableUpdateCompanyUserRequest{value: val, isSet: true}
}

func (v NullableUpdateCompanyUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCompanyUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
