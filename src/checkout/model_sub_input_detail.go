/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the SubInputDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubInputDetail{}

// SubInputDetail struct for SubInputDetail
type SubInputDetail struct {
	// Configuration parameters for the required input.
	Configuration *map[string]string `json:"configuration,omitempty"`
	// In case of a select, the items to choose from.
	Items []Item `json:"items,omitempty"`
	// The value to provide in the result.
	Key *string `json:"key,omitempty"`
	// True if this input is optional to provide.
	Optional *bool `json:"optional,omitempty"`
	// The type of the required input.
	Type *string `json:"type,omitempty"`
	// The value can be pre-filled, if available.
	Value *string `json:"value,omitempty"`
}

// NewSubInputDetail instantiates a new SubInputDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubInputDetail() *SubInputDetail {
	this := SubInputDetail{}
	return &this
}

// NewSubInputDetailWithDefaults instantiates a new SubInputDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubInputDetailWithDefaults() *SubInputDetail {
	this := SubInputDetail{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *SubInputDetail) GetConfiguration() map[string]string {
	if o == nil || common.IsNil(o.Configuration) {
		var ret map[string]string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubInputDetail) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *SubInputDetail) HasConfiguration() bool {
	if o != nil && !common.IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *SubInputDetail) SetConfiguration(v map[string]string) {
	o.Configuration = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SubInputDetail) GetItems() []Item {
	if o == nil || common.IsNil(o.Items) {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubInputDetail) GetItemsOk() ([]Item, bool) {
	if o == nil || common.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SubInputDetail) HasItems() bool {
	if o != nil && !common.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *SubInputDetail) SetItems(v []Item) {
	o.Items = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SubInputDetail) GetKey() string {
	if o == nil || common.IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubInputDetail) GetKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SubInputDetail) HasKey() bool {
	if o != nil && !common.IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SubInputDetail) SetKey(v string) {
	o.Key = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *SubInputDetail) GetOptional() bool {
	if o == nil || common.IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubInputDetail) GetOptionalOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *SubInputDetail) HasOptional() bool {
	if o != nil && !common.IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *SubInputDetail) SetOptional(v bool) {
	o.Optional = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubInputDetail) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubInputDetail) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubInputDetail) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SubInputDetail) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SubInputDetail) GetValue() string {
	if o == nil || common.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubInputDetail) GetValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SubInputDetail) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SubInputDetail) SetValue(v string) {
	o.Value = &v
}

func (o SubInputDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubInputDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !common.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !common.IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !common.IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSubInputDetail struct {
	value *SubInputDetail
	isSet bool
}

func (v NullableSubInputDetail) Get() *SubInputDetail {
	return v.value
}

func (v *NullableSubInputDetail) Set(val *SubInputDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSubInputDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSubInputDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubInputDetail(val *SubInputDetail) *NullableSubInputDetail {
	return &NullableSubInputDetail{value: val, isSet: true}
}

func (v NullableSubInputDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubInputDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



