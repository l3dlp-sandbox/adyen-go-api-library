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

// checks if the TerminalOrderRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalOrderRequest{}

// TerminalOrderRequest struct for TerminalOrderRequest
type TerminalOrderRequest struct {
	// The identification of the billing entity to use for the order.    > When ordering products in Brazil, you do not need to include the `billingEntityId` in the request.
	BillingEntityId *string `json:"billingEntityId,omitempty"`
	// The merchant-defined purchase order reference.
	CustomerOrderReference *string `json:"customerOrderReference,omitempty"`
	// The products included in the order.
	Items []OrderItem `json:"items,omitempty"`
	// Type of order
	OrderType *string `json:"orderType,omitempty"`
	// The identification of the shipping location to use for the order.
	ShippingLocationId *string `json:"shippingLocationId,omitempty"`
	// The tax number of the billing entity.
	TaxId *string `json:"taxId,omitempty"`
}

// NewTerminalOrderRequest instantiates a new TerminalOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalOrderRequest() *TerminalOrderRequest {
	this := TerminalOrderRequest{}
	return &this
}

// NewTerminalOrderRequestWithDefaults instantiates a new TerminalOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalOrderRequestWithDefaults() *TerminalOrderRequest {
	this := TerminalOrderRequest{}
	return &this
}

// GetBillingEntityId returns the BillingEntityId field value if set, zero value otherwise.
func (o *TerminalOrderRequest) GetBillingEntityId() string {
	if o == nil || common.IsNil(o.BillingEntityId) {
		var ret string
		return ret
	}
	return *o.BillingEntityId
}

// GetBillingEntityIdOk returns a tuple with the BillingEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrderRequest) GetBillingEntityIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingEntityId) {
		return nil, false
	}
	return o.BillingEntityId, true
}

// HasBillingEntityId returns a boolean if a field has been set.
func (o *TerminalOrderRequest) HasBillingEntityId() bool {
	if o != nil && !common.IsNil(o.BillingEntityId) {
		return true
	}

	return false
}

// SetBillingEntityId gets a reference to the given string and assigns it to the BillingEntityId field.
func (o *TerminalOrderRequest) SetBillingEntityId(v string) {
	o.BillingEntityId = &v
}

// GetCustomerOrderReference returns the CustomerOrderReference field value if set, zero value otherwise.
func (o *TerminalOrderRequest) GetCustomerOrderReference() string {
	if o == nil || common.IsNil(o.CustomerOrderReference) {
		var ret string
		return ret
	}
	return *o.CustomerOrderReference
}

// GetCustomerOrderReferenceOk returns a tuple with the CustomerOrderReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrderRequest) GetCustomerOrderReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CustomerOrderReference) {
		return nil, false
	}
	return o.CustomerOrderReference, true
}

// HasCustomerOrderReference returns a boolean if a field has been set.
func (o *TerminalOrderRequest) HasCustomerOrderReference() bool {
	if o != nil && !common.IsNil(o.CustomerOrderReference) {
		return true
	}

	return false
}

// SetCustomerOrderReference gets a reference to the given string and assigns it to the CustomerOrderReference field.
func (o *TerminalOrderRequest) SetCustomerOrderReference(v string) {
	o.CustomerOrderReference = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TerminalOrderRequest) GetItems() []OrderItem {
	if o == nil || common.IsNil(o.Items) {
		var ret []OrderItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrderRequest) GetItemsOk() ([]OrderItem, bool) {
	if o == nil || common.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TerminalOrderRequest) HasItems() bool {
	if o != nil && !common.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderItem and assigns it to the Items field.
func (o *TerminalOrderRequest) SetItems(v []OrderItem) {
	o.Items = v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *TerminalOrderRequest) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrderRequest) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *TerminalOrderRequest) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *TerminalOrderRequest) SetOrderType(v string) {
	o.OrderType = &v
}

// GetShippingLocationId returns the ShippingLocationId field value if set, zero value otherwise.
func (o *TerminalOrderRequest) GetShippingLocationId() string {
	if o == nil || common.IsNil(o.ShippingLocationId) {
		var ret string
		return ret
	}
	return *o.ShippingLocationId
}

// GetShippingLocationIdOk returns a tuple with the ShippingLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrderRequest) GetShippingLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShippingLocationId) {
		return nil, false
	}
	return o.ShippingLocationId, true
}

// HasShippingLocationId returns a boolean if a field has been set.
func (o *TerminalOrderRequest) HasShippingLocationId() bool {
	if o != nil && !common.IsNil(o.ShippingLocationId) {
		return true
	}

	return false
}

// SetShippingLocationId gets a reference to the given string and assigns it to the ShippingLocationId field.
func (o *TerminalOrderRequest) SetShippingLocationId(v string) {
	o.ShippingLocationId = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *TerminalOrderRequest) GetTaxId() string {
	if o == nil || common.IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrderRequest) GetTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *TerminalOrderRequest) HasTaxId() bool {
	if o != nil && !common.IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *TerminalOrderRequest) SetTaxId(v string) {
	o.TaxId = &v
}

func (o TerminalOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BillingEntityId) {
		toSerialize["billingEntityId"] = o.BillingEntityId
	}
	if !common.IsNil(o.CustomerOrderReference) {
		toSerialize["customerOrderReference"] = o.CustomerOrderReference
	}
	if !common.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.ShippingLocationId) {
		toSerialize["shippingLocationId"] = o.ShippingLocationId
	}
	if !common.IsNil(o.TaxId) {
		toSerialize["taxId"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableTerminalOrderRequest struct {
	value *TerminalOrderRequest
	isSet bool
}

func (v NullableTerminalOrderRequest) Get() *TerminalOrderRequest {
	return v.value
}

func (v *NullableTerminalOrderRequest) Set(val *TerminalOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalOrderRequest(val *TerminalOrderRequest) *NullableTerminalOrderRequest {
	return &NullableTerminalOrderRequest{value: val, isSet: true}
}

func (v NullableTerminalOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



