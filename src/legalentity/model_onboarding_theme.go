/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the OnboardingTheme type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OnboardingTheme{}

// OnboardingTheme struct for OnboardingTheme
type OnboardingTheme struct {
	// The creation date of the theme.
	CreatedAt time.Time `json:"createdAt"`
	// The description of the theme.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the theme.
	Id string `json:"id"`
	// The properties of the theme.
	Properties map[string]string `json:"properties"`
	// The date when the theme was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewOnboardingTheme instantiates a new OnboardingTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingTheme(createdAt time.Time, id string, properties map[string]string) *OnboardingTheme {
	this := OnboardingTheme{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Properties = properties
	return &this
}

// NewOnboardingThemeWithDefaults instantiates a new OnboardingTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingThemeWithDefaults() *OnboardingTheme {
	this := OnboardingTheme{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *OnboardingTheme) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OnboardingTheme) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OnboardingTheme) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OnboardingTheme) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingTheme) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OnboardingTheme) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OnboardingTheme) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *OnboardingTheme) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OnboardingTheme) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OnboardingTheme) SetId(v string) {
	o.Id = v
}

// GetProperties returns the Properties field value
func (o *OnboardingTheme) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *OnboardingTheme) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *OnboardingTheme) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OnboardingTheme) GetUpdatedAt() time.Time {
	if o == nil || common.IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingTheme) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OnboardingTheme) HasUpdatedAt() bool {
	if o != nil && !common.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OnboardingTheme) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o OnboardingTheme) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["properties"] = o.Properties
	if !common.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableOnboardingTheme struct {
	value *OnboardingTheme
	isSet bool
}

func (v NullableOnboardingTheme) Get() *OnboardingTheme {
	return v.value
}

func (v *NullableOnboardingTheme) Set(val *OnboardingTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingTheme(val *OnboardingTheme) *NullableOnboardingTheme {
	return &NullableOnboardingTheme{value: val, isSet: true}
}

func (v NullableOnboardingTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
