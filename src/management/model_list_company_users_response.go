/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the ListCompanyUsersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListCompanyUsersResponse{}

// ListCompanyUsersResponse struct for ListCompanyUsersResponse
type ListCompanyUsersResponse struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// The list of users.
	Data []CompanyUser `json:"data,omitempty"`
	// Total number of items.
	ItemsTotal int32 `json:"itemsTotal"`
	// Total number of pages.
	PagesTotal int32 `json:"pagesTotal"`
}

// NewListCompanyUsersResponse instantiates a new ListCompanyUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCompanyUsersResponse(itemsTotal int32, pagesTotal int32) *ListCompanyUsersResponse {
	this := ListCompanyUsersResponse{}
	this.ItemsTotal = itemsTotal
	this.PagesTotal = pagesTotal
	return &this
}

// NewListCompanyUsersResponseWithDefaults instantiates a new ListCompanyUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCompanyUsersResponseWithDefaults() *ListCompanyUsersResponse {
	this := ListCompanyUsersResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListCompanyUsersResponse) GetLinks() PaginationLinks {
	if o == nil || common.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompanyUsersResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListCompanyUsersResponse) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ListCompanyUsersResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListCompanyUsersResponse) GetData() []CompanyUser {
	if o == nil || common.IsNil(o.Data) {
		var ret []CompanyUser
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompanyUsersResponse) GetDataOk() ([]CompanyUser, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListCompanyUsersResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CompanyUser and assigns it to the Data field.
func (o *ListCompanyUsersResponse) SetData(v []CompanyUser) {
	o.Data = v
}

// GetItemsTotal returns the ItemsTotal field value
func (o *ListCompanyUsersResponse) GetItemsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value
// and a boolean to check if the value has been set.
func (o *ListCompanyUsersResponse) GetItemsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsTotal, true
}

// SetItemsTotal sets field value
func (o *ListCompanyUsersResponse) SetItemsTotal(v int32) {
	o.ItemsTotal = v
}

// GetPagesTotal returns the PagesTotal field value
func (o *ListCompanyUsersResponse) GetPagesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PagesTotal
}

// GetPagesTotalOk returns a tuple with the PagesTotal field value
// and a boolean to check if the value has been set.
func (o *ListCompanyUsersResponse) GetPagesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PagesTotal, true
}

// SetPagesTotal sets field value
func (o *ListCompanyUsersResponse) SetPagesTotal(v int32) {
	o.PagesTotal = v
}

func (o ListCompanyUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCompanyUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["itemsTotal"] = o.ItemsTotal
	toSerialize["pagesTotal"] = o.PagesTotal
	return toSerialize, nil
}

type NullableListCompanyUsersResponse struct {
	value *ListCompanyUsersResponse
	isSet bool
}

func (v NullableListCompanyUsersResponse) Get() *ListCompanyUsersResponse {
	return v.value
}

func (v *NullableListCompanyUsersResponse) Set(val *ListCompanyUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCompanyUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCompanyUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCompanyUsersResponse(val *ListCompanyUsersResponse) *NullableListCompanyUsersResponse {
	return &NullableListCompanyUsersResponse{value: val, isSet: true}
}

func (v NullableListCompanyUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCompanyUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
