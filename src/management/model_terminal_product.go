/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the TerminalProduct type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalProduct{}

// TerminalProduct struct for TerminalProduct
type TerminalProduct struct {
	// Information about items included and integration options.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the product.
	Id *string `json:"id,omitempty"`
	// A list of parts included in the terminal package.
	ItemsIncluded []string `json:"itemsIncluded,omitempty"`
	// The descriptive name of the product.
	Name *string `json:"name,omitempty"`
	Price *TerminalProductPrice `json:"price,omitempty"`
}

// NewTerminalProduct instantiates a new TerminalProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalProduct() *TerminalProduct {
	this := TerminalProduct{}
	return &this
}

// NewTerminalProductWithDefaults instantiates a new TerminalProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalProductWithDefaults() *TerminalProduct {
	this := TerminalProduct{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TerminalProduct) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TerminalProduct) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TerminalProduct) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TerminalProduct) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProduct) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TerminalProduct) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TerminalProduct) SetId(v string) {
	o.Id = &v
}

// GetItemsIncluded returns the ItemsIncluded field value if set, zero value otherwise.
func (o *TerminalProduct) GetItemsIncluded() []string {
	if o == nil || common.IsNil(o.ItemsIncluded) {
		var ret []string
		return ret
	}
	return o.ItemsIncluded
}

// GetItemsIncludedOk returns a tuple with the ItemsIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProduct) GetItemsIncludedOk() ([]string, bool) {
	if o == nil || common.IsNil(o.ItemsIncluded) {
		return nil, false
	}
	return o.ItemsIncluded, true
}

// HasItemsIncluded returns a boolean if a field has been set.
func (o *TerminalProduct) HasItemsIncluded() bool {
	if o != nil && !common.IsNil(o.ItemsIncluded) {
		return true
	}

	return false
}

// SetItemsIncluded gets a reference to the given []string and assigns it to the ItemsIncluded field.
func (o *TerminalProduct) SetItemsIncluded(v []string) {
	o.ItemsIncluded = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TerminalProduct) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProduct) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TerminalProduct) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TerminalProduct) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TerminalProduct) GetPrice() TerminalProductPrice {
	if o == nil || common.IsNil(o.Price) {
		var ret TerminalProductPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProduct) GetPriceOk() (*TerminalProductPrice, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TerminalProduct) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given TerminalProductPrice and assigns it to the Price field.
func (o *TerminalProduct) SetPrice(v TerminalProductPrice) {
	o.Price = &v
}

func (o TerminalProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.ItemsIncluded) {
		toSerialize["itemsIncluded"] = o.ItemsIncluded
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableTerminalProduct struct {
	value *TerminalProduct
	isSet bool
}

func (v NullableTerminalProduct) Get() *TerminalProduct {
	return v.value
}

func (v *NullableTerminalProduct) Set(val *TerminalProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalProduct(val *TerminalProduct) *NullableTerminalProduct {
	return &NullableTerminalProduct{value: val, isSet: true}
}

func (v NullableTerminalProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



