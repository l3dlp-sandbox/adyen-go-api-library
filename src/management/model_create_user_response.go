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

// checks if the CreateUserResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateUserResponse{}

// CreateUserResponse struct for CreateUserResponse
type CreateUserResponse struct {
	Links *Links `json:"_links,omitempty"`
	// The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user.
	AccountGroups []string `json:"accountGroups,omitempty"`
	// Indicates whether this user is active.
	Active *bool `json:"active,omitempty"`
	// Set of apps available to this user
	Apps []string `json:"apps,omitempty"`
	// The email address of the user.
	Email string `json:"email"`
	// The unique identifier of the user.
	Id   string `json:"id"`
	Name *Name  `json:"name,omitempty"`
	// The list of [roles](https://docs.adyen.com/account/user-roles) for this user.
	Roles []string `json:"roles"`
	// The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**.
	TimeZoneCode string `json:"timeZoneCode"`
	// The username for this user.
	Username string `json:"username"`
}

// NewCreateUserResponse instantiates a new CreateUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserResponse(email string, id string, roles []string, timeZoneCode string, username string) *CreateUserResponse {
	this := CreateUserResponse{}
	this.Email = email
	this.Id = id
	this.Roles = roles
	this.TimeZoneCode = timeZoneCode
	this.Username = username
	return &this
}

// NewCreateUserResponseWithDefaults instantiates a new CreateUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserResponseWithDefaults() *CreateUserResponse {
	this := CreateUserResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateUserResponse) GetLinks() Links {
	if o == nil || common.IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetLinksOk() (*Links, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateUserResponse) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *CreateUserResponse) SetLinks(v Links) {
	o.Links = &v
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *CreateUserResponse) GetAccountGroups() []string {
	if o == nil || common.IsNil(o.AccountGroups) {
		var ret []string
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetAccountGroupsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *CreateUserResponse) HasAccountGroups() bool {
	if o != nil && !common.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []string and assigns it to the AccountGroups field.
func (o *CreateUserResponse) SetAccountGroups(v []string) {
	o.AccountGroups = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateUserResponse) GetActive() bool {
	if o == nil || common.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateUserResponse) HasActive() bool {
	if o != nil && !common.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateUserResponse) SetActive(v bool) {
	o.Active = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *CreateUserResponse) GetApps() []string {
	if o == nil || common.IsNil(o.Apps) {
		var ret []string
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetAppsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *CreateUserResponse) HasApps() bool {
	if o != nil && !common.IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []string and assigns it to the Apps field.
func (o *CreateUserResponse) SetApps(v []string) {
	o.Apps = v
}

// GetEmail returns the Email field value
func (o *CreateUserResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateUserResponse) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *CreateUserResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateUserResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateUserResponse) GetName() Name {
	if o == nil || common.IsNil(o.Name) {
		var ret Name
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateUserResponse) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given Name and assigns it to the Name field.
func (o *CreateUserResponse) SetName(v Name) {
	o.Name = &v
}

// GetRoles returns the Roles field value
func (o *CreateUserResponse) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *CreateUserResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetTimeZoneCode returns the TimeZoneCode field value
func (o *CreateUserResponse) GetTimeZoneCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZoneCode
}

// GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field value
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetTimeZoneCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZoneCode, true
}

// SetTimeZoneCode sets field value
func (o *CreateUserResponse) SetTimeZoneCode(v string) {
	o.TimeZoneCode = v
}

// GetUsername returns the Username field value
func (o *CreateUserResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateUserResponse) SetUsername(v string) {
	o.Username = v
}

func (o CreateUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !common.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if !common.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !common.IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["roles"] = o.Roles
	toSerialize["timeZoneCode"] = o.TimeZoneCode
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableCreateUserResponse struct {
	value *CreateUserResponse
	isSet bool
}

func (v NullableCreateUserResponse) Get() *CreateUserResponse {
	return v.value
}

func (v *NullableCreateUserResponse) Set(val *CreateUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserResponse(val *CreateUserResponse) *NullableCreateUserResponse {
	return &NullableCreateUserResponse{value: val, isSet: true}
}

func (v NullableCreateUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
