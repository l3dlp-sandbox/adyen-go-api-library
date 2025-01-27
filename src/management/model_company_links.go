/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the CompanyLinks type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CompanyLinks{}

// CompanyLinks struct for CompanyLinks
type CompanyLinks struct {
	ApiCredentials *LinksElement `json:"apiCredentials,omitempty"`
	Self           LinksElement  `json:"self"`
	Users          *LinksElement `json:"users,omitempty"`
	Webhooks       *LinksElement `json:"webhooks,omitempty"`
}

// NewCompanyLinks instantiates a new CompanyLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyLinks(self LinksElement) *CompanyLinks {
	this := CompanyLinks{}
	this.Self = self
	return &this
}

// NewCompanyLinksWithDefaults instantiates a new CompanyLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyLinksWithDefaults() *CompanyLinks {
	this := CompanyLinks{}
	return &this
}

// GetApiCredentials returns the ApiCredentials field value if set, zero value otherwise.
func (o *CompanyLinks) GetApiCredentials() LinksElement {
	if o == nil || common.IsNil(o.ApiCredentials) {
		var ret LinksElement
		return ret
	}
	return *o.ApiCredentials
}

// GetApiCredentialsOk returns a tuple with the ApiCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLinks) GetApiCredentialsOk() (*LinksElement, bool) {
	if o == nil || common.IsNil(o.ApiCredentials) {
		return nil, false
	}
	return o.ApiCredentials, true
}

// HasApiCredentials returns a boolean if a field has been set.
func (o *CompanyLinks) HasApiCredentials() bool {
	if o != nil && !common.IsNil(o.ApiCredentials) {
		return true
	}

	return false
}

// SetApiCredentials gets a reference to the given LinksElement and assigns it to the ApiCredentials field.
func (o *CompanyLinks) SetApiCredentials(v LinksElement) {
	o.ApiCredentials = &v
}

// GetSelf returns the Self field value
func (o *CompanyLinks) GetSelf() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *CompanyLinks) GetSelfOk() (*LinksElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *CompanyLinks) SetSelf(v LinksElement) {
	o.Self = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *CompanyLinks) GetUsers() LinksElement {
	if o == nil || common.IsNil(o.Users) {
		var ret LinksElement
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLinks) GetUsersOk() (*LinksElement, bool) {
	if o == nil || common.IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CompanyLinks) HasUsers() bool {
	if o != nil && !common.IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given LinksElement and assigns it to the Users field.
func (o *CompanyLinks) SetUsers(v LinksElement) {
	o.Users = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *CompanyLinks) GetWebhooks() LinksElement {
	if o == nil || common.IsNil(o.Webhooks) {
		var ret LinksElement
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyLinks) GetWebhooksOk() (*LinksElement, bool) {
	if o == nil || common.IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *CompanyLinks) HasWebhooks() bool {
	if o != nil && !common.IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given LinksElement and assigns it to the Webhooks field.
func (o *CompanyLinks) SetWebhooks(v LinksElement) {
	o.Webhooks = &v
}

func (o CompanyLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApiCredentials) {
		toSerialize["apiCredentials"] = o.ApiCredentials
	}
	toSerialize["self"] = o.Self
	if !common.IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !common.IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	return toSerialize, nil
}

type NullableCompanyLinks struct {
	value *CompanyLinks
	isSet bool
}

func (v NullableCompanyLinks) Get() *CompanyLinks {
	return v.value
}

func (v *NullableCompanyLinks) Set(val *CompanyLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyLinks(val *CompanyLinks) *NullableCompanyLinks {
	return &NullableCompanyLinks{value: val, isSet: true}
}

func (v NullableCompanyLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
