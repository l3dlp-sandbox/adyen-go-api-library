/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the ShippingLocation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ShippingLocation{}

// ShippingLocation struct for ShippingLocation
type ShippingLocation struct {
	Address *Address `json:"address,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	// The unique identifier of the shipping location, for use as `shippingLocationId` when creating an order.
	Id *string `json:"id,omitempty"`
	// The unique name of the shipping location.
	Name *string `json:"name,omitempty"`
}

// NewShippingLocation instantiates a new ShippingLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingLocation() *ShippingLocation {
	this := ShippingLocation{}
	return &this
}

// NewShippingLocationWithDefaults instantiates a new ShippingLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingLocationWithDefaults() *ShippingLocation {
	this := ShippingLocation{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ShippingLocation) GetAddress() Address {
	if o == nil || common.IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingLocation) GetAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ShippingLocation) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ShippingLocation) SetAddress(v Address) {
	o.Address = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *ShippingLocation) GetContact() Contact {
	if o == nil || common.IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingLocation) GetContactOk() (*Contact, bool) {
	if o == nil || common.IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *ShippingLocation) HasContact() bool {
	if o != nil && !common.IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *ShippingLocation) SetContact(v Contact) {
	o.Contact = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShippingLocation) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingLocation) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShippingLocation) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ShippingLocation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShippingLocation) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingLocation) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShippingLocation) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShippingLocation) SetName(v string) {
	o.Name = &v
}

func (o ShippingLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableShippingLocation struct {
	value *ShippingLocation
	isSet bool
}

func (v NullableShippingLocation) Get() *ShippingLocation {
	return v.value
}

func (v *NullableShippingLocation) Set(val *ShippingLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingLocation(val *ShippingLocation) *NullableShippingLocation {
	return &NullableShippingLocation{value: val, isSet: true}
}

func (v NullableShippingLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
