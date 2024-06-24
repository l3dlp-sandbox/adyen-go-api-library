/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the InputDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InputDetail{}

// InputDetail struct for InputDetail
type InputDetail struct {
	// Configuration parameters for the required input.
	Configuration *map[string]string `json:"configuration,omitempty"`
	// Input details can also be provided recursively.
	Details []SubInputDetail `json:"details,omitempty"`
	// Input details can also be provided recursively (deprecated).
	// Deprecated
	InputDetails []SubInputDetail `json:"inputDetails,omitempty"`
	// In case of a select, the URL from which to query the items.
	ItemSearchUrl *string `json:"itemSearchUrl,omitempty"`
	// In case of a select, the items to choose from.
	Items []Item `json:"items,omitempty"`
	// The value to provide in the result.
	Key *string `json:"key,omitempty"`
	// True if this input value is optional.
	Optional *bool `json:"optional,omitempty"`
	// The type of the required input.
	Type *string `json:"type,omitempty"`
	// The value can be pre-filled, if available.
	Value *string `json:"value,omitempty"`
}

// NewInputDetail instantiates a new InputDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputDetail() *InputDetail {
	this := InputDetail{}
	return &this
}

// NewInputDetailWithDefaults instantiates a new InputDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputDetailWithDefaults() *InputDetail {
	this := InputDetail{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *InputDetail) GetConfiguration() map[string]string {
	if o == nil || common.IsNil(o.Configuration) {
		var ret map[string]string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *InputDetail) HasConfiguration() bool {
	if o != nil && !common.IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *InputDetail) SetConfiguration(v map[string]string) {
	o.Configuration = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InputDetail) GetDetails() []SubInputDetail {
	if o == nil || common.IsNil(o.Details) {
		var ret []SubInputDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetDetailsOk() ([]SubInputDetail, bool) {
	if o == nil || common.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InputDetail) HasDetails() bool {
	if o != nil && !common.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []SubInputDetail and assigns it to the Details field.
func (o *InputDetail) SetDetails(v []SubInputDetail) {
	o.Details = v
}

// GetInputDetails returns the InputDetails field value if set, zero value otherwise.
// Deprecated
func (o *InputDetail) GetInputDetails() []SubInputDetail {
	if o == nil || common.IsNil(o.InputDetails) {
		var ret []SubInputDetail
		return ret
	}
	return o.InputDetails
}

// GetInputDetailsOk returns a tuple with the InputDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *InputDetail) GetInputDetailsOk() ([]SubInputDetail, bool) {
	if o == nil || common.IsNil(o.InputDetails) {
		return nil, false
	}
	return o.InputDetails, true
}

// HasInputDetails returns a boolean if a field has been set.
func (o *InputDetail) HasInputDetails() bool {
	if o != nil && !common.IsNil(o.InputDetails) {
		return true
	}

	return false
}

// SetInputDetails gets a reference to the given []SubInputDetail and assigns it to the InputDetails field.
// Deprecated
func (o *InputDetail) SetInputDetails(v []SubInputDetail) {
	o.InputDetails = v
}

// GetItemSearchUrl returns the ItemSearchUrl field value if set, zero value otherwise.
func (o *InputDetail) GetItemSearchUrl() string {
	if o == nil || common.IsNil(o.ItemSearchUrl) {
		var ret string
		return ret
	}
	return *o.ItemSearchUrl
}

// GetItemSearchUrlOk returns a tuple with the ItemSearchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetItemSearchUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.ItemSearchUrl) {
		return nil, false
	}
	return o.ItemSearchUrl, true
}

// HasItemSearchUrl returns a boolean if a field has been set.
func (o *InputDetail) HasItemSearchUrl() bool {
	if o != nil && !common.IsNil(o.ItemSearchUrl) {
		return true
	}

	return false
}

// SetItemSearchUrl gets a reference to the given string and assigns it to the ItemSearchUrl field.
func (o *InputDetail) SetItemSearchUrl(v string) {
	o.ItemSearchUrl = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InputDetail) GetItems() []Item {
	if o == nil || common.IsNil(o.Items) {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetItemsOk() ([]Item, bool) {
	if o == nil || common.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InputDetail) HasItems() bool {
	if o != nil && !common.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *InputDetail) SetItems(v []Item) {
	o.Items = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InputDetail) GetKey() string {
	if o == nil || common.IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InputDetail) HasKey() bool {
	if o != nil && !common.IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InputDetail) SetKey(v string) {
	o.Key = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *InputDetail) GetOptional() bool {
	if o == nil || common.IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetOptionalOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *InputDetail) HasOptional() bool {
	if o != nil && !common.IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *InputDetail) SetOptional(v bool) {
	o.Optional = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InputDetail) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InputDetail) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InputDetail) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InputDetail) GetValue() string {
	if o == nil || common.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputDetail) GetValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InputDetail) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InputDetail) SetValue(v string) {
	o.Value = &v
}

func (o InputDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !common.IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !common.IsNil(o.InputDetails) {
		toSerialize["inputDetails"] = o.InputDetails
	}
	if !common.IsNil(o.ItemSearchUrl) {
		toSerialize["itemSearchUrl"] = o.ItemSearchUrl
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

type NullableInputDetail struct {
	value *InputDetail
	isSet bool
}

func (v NullableInputDetail) Get() *InputDetail {
	return v.value
}

func (v *NullableInputDetail) Set(val *InputDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableInputDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableInputDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputDetail(val *InputDetail) *NullableInputDetail {
	return &NullableInputDetail{value: val, isSet: true}
}

func (v NullableInputDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
