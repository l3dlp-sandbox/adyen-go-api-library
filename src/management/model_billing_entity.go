/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the BillingEntity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BillingEntity{}

// BillingEntity struct for BillingEntity
type BillingEntity struct {
	Address *Address `json:"address,omitempty"`
	// The email address of the billing entity.
	Email *string `json:"email,omitempty"`
	// The unique identifier of the billing entity, for use as `billingEntityId` when creating an order.
	Id *string `json:"id,omitempty"`
	// The unique name of the billing entity.
	Name *string `json:"name,omitempty"`
	// The tax number of the billing entity.
	TaxId *string `json:"taxId,omitempty"`
}

// NewBillingEntity instantiates a new BillingEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEntity() *BillingEntity {
	this := BillingEntity{}
	return &this
}

// NewBillingEntityWithDefaults instantiates a new BillingEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEntityWithDefaults() *BillingEntity {
	this := BillingEntity{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *BillingEntity) GetAddress() Address {
	if o == nil || common.IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEntity) GetAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *BillingEntity) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *BillingEntity) SetAddress(v Address) {
	o.Address = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BillingEntity) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEntity) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BillingEntity) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BillingEntity) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingEntity) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEntity) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingEntity) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingEntity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingEntity) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEntity) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingEntity) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingEntity) SetName(v string) {
	o.Name = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *BillingEntity) GetTaxId() string {
	if o == nil || common.IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEntity) GetTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *BillingEntity) HasTaxId() bool {
	if o != nil && !common.IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *BillingEntity) SetTaxId(v string) {
	o.TaxId = &v
}

func (o BillingEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.TaxId) {
		toSerialize["taxId"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableBillingEntity struct {
	value *BillingEntity
	isSet bool
}

func (v NullableBillingEntity) Get() *BillingEntity {
	return v.value
}

func (v *NullableBillingEntity) Set(val *BillingEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingEntity(val *BillingEntity) *NullableBillingEntity {
	return &NullableBillingEntity{value: val, isSet: true}
}

func (v NullableBillingEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
