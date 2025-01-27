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

// checks if the ListWebhooksResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListWebhooksResponse{}

// ListWebhooksResponse struct for ListWebhooksResponse
type ListWebhooksResponse struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// Reference to the account.
	AccountReference *string `json:"accountReference,omitempty"`
	// The list of webhooks configured for this account.
	Data []Webhook `json:"data,omitempty"`
	// Total number of items.
	ItemsTotal int32 `json:"itemsTotal"`
	// Total number of pages.
	PagesTotal int32 `json:"pagesTotal"`
}

// NewListWebhooksResponse instantiates a new ListWebhooksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhooksResponse(itemsTotal int32, pagesTotal int32) *ListWebhooksResponse {
	this := ListWebhooksResponse{}
	this.ItemsTotal = itemsTotal
	this.PagesTotal = pagesTotal
	return &this
}

// NewListWebhooksResponseWithDefaults instantiates a new ListWebhooksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhooksResponseWithDefaults() *ListWebhooksResponse {
	this := ListWebhooksResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListWebhooksResponse) GetLinks() PaginationLinks {
	if o == nil || common.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhooksResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListWebhooksResponse) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ListWebhooksResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetAccountReference returns the AccountReference field value if set, zero value otherwise.
func (o *ListWebhooksResponse) GetAccountReference() string {
	if o == nil || common.IsNil(o.AccountReference) {
		var ret string
		return ret
	}
	return *o.AccountReference
}

// GetAccountReferenceOk returns a tuple with the AccountReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhooksResponse) GetAccountReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountReference) {
		return nil, false
	}
	return o.AccountReference, true
}

// HasAccountReference returns a boolean if a field has been set.
func (o *ListWebhooksResponse) HasAccountReference() bool {
	if o != nil && !common.IsNil(o.AccountReference) {
		return true
	}

	return false
}

// SetAccountReference gets a reference to the given string and assigns it to the AccountReference field.
func (o *ListWebhooksResponse) SetAccountReference(v string) {
	o.AccountReference = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListWebhooksResponse) GetData() []Webhook {
	if o == nil || common.IsNil(o.Data) {
		var ret []Webhook
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhooksResponse) GetDataOk() ([]Webhook, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListWebhooksResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Webhook and assigns it to the Data field.
func (o *ListWebhooksResponse) SetData(v []Webhook) {
	o.Data = v
}

// GetItemsTotal returns the ItemsTotal field value
func (o *ListWebhooksResponse) GetItemsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value
// and a boolean to check if the value has been set.
func (o *ListWebhooksResponse) GetItemsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsTotal, true
}

// SetItemsTotal sets field value
func (o *ListWebhooksResponse) SetItemsTotal(v int32) {
	o.ItemsTotal = v
}

// GetPagesTotal returns the PagesTotal field value
func (o *ListWebhooksResponse) GetPagesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PagesTotal
}

// GetPagesTotalOk returns a tuple with the PagesTotal field value
// and a boolean to check if the value has been set.
func (o *ListWebhooksResponse) GetPagesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PagesTotal, true
}

// SetPagesTotal sets field value
func (o *ListWebhooksResponse) SetPagesTotal(v int32) {
	o.PagesTotal = v
}

func (o ListWebhooksResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhooksResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !common.IsNil(o.AccountReference) {
		toSerialize["accountReference"] = o.AccountReference
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["itemsTotal"] = o.ItemsTotal
	toSerialize["pagesTotal"] = o.PagesTotal
	return toSerialize, nil
}

type NullableListWebhooksResponse struct {
	value *ListWebhooksResponse
	isSet bool
}

func (v NullableListWebhooksResponse) Get() *ListWebhooksResponse {
	return v.value
}

func (v *NullableListWebhooksResponse) Set(val *ListWebhooksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhooksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhooksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhooksResponse(val *ListWebhooksResponse) *NullableListWebhooksResponse {
	return &NullableListWebhooksResponse{value: val, isSet: true}
}

func (v NullableListWebhooksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhooksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
