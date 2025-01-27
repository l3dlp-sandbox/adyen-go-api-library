/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the DonationCampaignsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DonationCampaignsResponse{}

// DonationCampaignsResponse struct for DonationCampaignsResponse
type DonationCampaignsResponse struct {
	// List of active donation campaigns for your merchant account.
	DonationCampaigns []DonationCampaign `json:"donationCampaigns,omitempty"`
}

// NewDonationCampaignsResponse instantiates a new DonationCampaignsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDonationCampaignsResponse() *DonationCampaignsResponse {
	this := DonationCampaignsResponse{}
	return &this
}

// NewDonationCampaignsResponseWithDefaults instantiates a new DonationCampaignsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDonationCampaignsResponseWithDefaults() *DonationCampaignsResponse {
	this := DonationCampaignsResponse{}
	return &this
}

// GetDonationCampaigns returns the DonationCampaigns field value if set, zero value otherwise.
func (o *DonationCampaignsResponse) GetDonationCampaigns() []DonationCampaign {
	if o == nil || common.IsNil(o.DonationCampaigns) {
		var ret []DonationCampaign
		return ret
	}
	return o.DonationCampaigns
}

// GetDonationCampaignsOk returns a tuple with the DonationCampaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaignsResponse) GetDonationCampaignsOk() ([]DonationCampaign, bool) {
	if o == nil || common.IsNil(o.DonationCampaigns) {
		return nil, false
	}
	return o.DonationCampaigns, true
}

// HasDonationCampaigns returns a boolean if a field has been set.
func (o *DonationCampaignsResponse) HasDonationCampaigns() bool {
	if o != nil && !common.IsNil(o.DonationCampaigns) {
		return true
	}

	return false
}

// SetDonationCampaigns gets a reference to the given []DonationCampaign and assigns it to the DonationCampaigns field.
func (o *DonationCampaignsResponse) SetDonationCampaigns(v []DonationCampaign) {
	o.DonationCampaigns = v
}

func (o DonationCampaignsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DonationCampaignsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DonationCampaigns) {
		toSerialize["donationCampaigns"] = o.DonationCampaigns
	}
	return toSerialize, nil
}

type NullableDonationCampaignsResponse struct {
	value *DonationCampaignsResponse
	isSet bool
}

func (v NullableDonationCampaignsResponse) Get() *DonationCampaignsResponse {
	return v.value
}

func (v *NullableDonationCampaignsResponse) Set(val *DonationCampaignsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDonationCampaignsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDonationCampaignsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonationCampaignsResponse(val *DonationCampaignsResponse) *NullableDonationCampaignsResponse {
	return &NullableDonationCampaignsResponse{value: val, isSet: true}
}

func (v NullableDonationCampaignsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonationCampaignsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
