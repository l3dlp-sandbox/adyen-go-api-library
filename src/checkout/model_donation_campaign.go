/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the DonationCampaign type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DonationCampaign{}

// DonationCampaign struct for DonationCampaign
type DonationCampaign struct {
	Amounts *Amounts `json:"amounts,omitempty"`
	// The URL for the banner of the nonprofit or campaign.
	BannerUrl *string `json:"bannerUrl,omitempty"`
	// The name of the donation campaign..
	CampaignName *string `json:"campaignName,omitempty"`
	// The cause of the nonprofit.
	CauseName *string   `json:"causeName,omitempty"`
	Donation  *Donation `json:"donation,omitempty"`
	// The unique campaign ID of the donation campaign.
	Id *string `json:"id,omitempty"`
	// The URL for the logo of the nonprofit.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// The description of the nonprofit.
	NonprofitDescription *string `json:"nonprofitDescription,omitempty"`
	// The name of the nonprofit organization that receives the donation.
	NonprofitName *string `json:"nonprofitName,omitempty"`
	// The website URL of the nonprofit.
	NonprofitUrl *string `json:"nonprofitUrl,omitempty"`
	// The URL of the terms and conditions page of the nonprofit and the campaign.
	TermsAndConditionsUrl *string `json:"termsAndConditionsUrl,omitempty"`
}

// NewDonationCampaign instantiates a new DonationCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDonationCampaign() *DonationCampaign {
	this := DonationCampaign{}
	return &this
}

// NewDonationCampaignWithDefaults instantiates a new DonationCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDonationCampaignWithDefaults() *DonationCampaign {
	this := DonationCampaign{}
	return &this
}

// GetAmounts returns the Amounts field value if set, zero value otherwise.
func (o *DonationCampaign) GetAmounts() Amounts {
	if o == nil || common.IsNil(o.Amounts) {
		var ret Amounts
		return ret
	}
	return *o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetAmountsOk() (*Amounts, bool) {
	if o == nil || common.IsNil(o.Amounts) {
		return nil, false
	}
	return o.Amounts, true
}

// HasAmounts returns a boolean if a field has been set.
func (o *DonationCampaign) HasAmounts() bool {
	if o != nil && !common.IsNil(o.Amounts) {
		return true
	}

	return false
}

// SetAmounts gets a reference to the given Amounts and assigns it to the Amounts field.
func (o *DonationCampaign) SetAmounts(v Amounts) {
	o.Amounts = &v
}

// GetBannerUrl returns the BannerUrl field value if set, zero value otherwise.
func (o *DonationCampaign) GetBannerUrl() string {
	if o == nil || common.IsNil(o.BannerUrl) {
		var ret string
		return ret
	}
	return *o.BannerUrl
}

// GetBannerUrlOk returns a tuple with the BannerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetBannerUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.BannerUrl) {
		return nil, false
	}
	return o.BannerUrl, true
}

// HasBannerUrl returns a boolean if a field has been set.
func (o *DonationCampaign) HasBannerUrl() bool {
	if o != nil && !common.IsNil(o.BannerUrl) {
		return true
	}

	return false
}

// SetBannerUrl gets a reference to the given string and assigns it to the BannerUrl field.
func (o *DonationCampaign) SetBannerUrl(v string) {
	o.BannerUrl = &v
}

// GetCampaignName returns the CampaignName field value if set, zero value otherwise.
func (o *DonationCampaign) GetCampaignName() string {
	if o == nil || common.IsNil(o.CampaignName) {
		var ret string
		return ret
	}
	return *o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetCampaignNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CampaignName) {
		return nil, false
	}
	return o.CampaignName, true
}

// HasCampaignName returns a boolean if a field has been set.
func (o *DonationCampaign) HasCampaignName() bool {
	if o != nil && !common.IsNil(o.CampaignName) {
		return true
	}

	return false
}

// SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.
func (o *DonationCampaign) SetCampaignName(v string) {
	o.CampaignName = &v
}

// GetCauseName returns the CauseName field value if set, zero value otherwise.
func (o *DonationCampaign) GetCauseName() string {
	if o == nil || common.IsNil(o.CauseName) {
		var ret string
		return ret
	}
	return *o.CauseName
}

// GetCauseNameOk returns a tuple with the CauseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetCauseNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CauseName) {
		return nil, false
	}
	return o.CauseName, true
}

// HasCauseName returns a boolean if a field has been set.
func (o *DonationCampaign) HasCauseName() bool {
	if o != nil && !common.IsNil(o.CauseName) {
		return true
	}

	return false
}

// SetCauseName gets a reference to the given string and assigns it to the CauseName field.
func (o *DonationCampaign) SetCauseName(v string) {
	o.CauseName = &v
}

// GetDonation returns the Donation field value if set, zero value otherwise.
func (o *DonationCampaign) GetDonation() Donation {
	if o == nil || common.IsNil(o.Donation) {
		var ret Donation
		return ret
	}
	return *o.Donation
}

// GetDonationOk returns a tuple with the Donation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetDonationOk() (*Donation, bool) {
	if o == nil || common.IsNil(o.Donation) {
		return nil, false
	}
	return o.Donation, true
}

// HasDonation returns a boolean if a field has been set.
func (o *DonationCampaign) HasDonation() bool {
	if o != nil && !common.IsNil(o.Donation) {
		return true
	}

	return false
}

// SetDonation gets a reference to the given Donation and assigns it to the Donation field.
func (o *DonationCampaign) SetDonation(v Donation) {
	o.Donation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DonationCampaign) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DonationCampaign) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DonationCampaign) SetId(v string) {
	o.Id = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *DonationCampaign) GetLogoUrl() string {
	if o == nil || common.IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetLogoUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *DonationCampaign) HasLogoUrl() bool {
	if o != nil && !common.IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *DonationCampaign) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetNonprofitDescription returns the NonprofitDescription field value if set, zero value otherwise.
func (o *DonationCampaign) GetNonprofitDescription() string {
	if o == nil || common.IsNil(o.NonprofitDescription) {
		var ret string
		return ret
	}
	return *o.NonprofitDescription
}

// GetNonprofitDescriptionOk returns a tuple with the NonprofitDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetNonprofitDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.NonprofitDescription) {
		return nil, false
	}
	return o.NonprofitDescription, true
}

// HasNonprofitDescription returns a boolean if a field has been set.
func (o *DonationCampaign) HasNonprofitDescription() bool {
	if o != nil && !common.IsNil(o.NonprofitDescription) {
		return true
	}

	return false
}

// SetNonprofitDescription gets a reference to the given string and assigns it to the NonprofitDescription field.
func (o *DonationCampaign) SetNonprofitDescription(v string) {
	o.NonprofitDescription = &v
}

// GetNonprofitName returns the NonprofitName field value if set, zero value otherwise.
func (o *DonationCampaign) GetNonprofitName() string {
	if o == nil || common.IsNil(o.NonprofitName) {
		var ret string
		return ret
	}
	return *o.NonprofitName
}

// GetNonprofitNameOk returns a tuple with the NonprofitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetNonprofitNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.NonprofitName) {
		return nil, false
	}
	return o.NonprofitName, true
}

// HasNonprofitName returns a boolean if a field has been set.
func (o *DonationCampaign) HasNonprofitName() bool {
	if o != nil && !common.IsNil(o.NonprofitName) {
		return true
	}

	return false
}

// SetNonprofitName gets a reference to the given string and assigns it to the NonprofitName field.
func (o *DonationCampaign) SetNonprofitName(v string) {
	o.NonprofitName = &v
}

// GetNonprofitUrl returns the NonprofitUrl field value if set, zero value otherwise.
func (o *DonationCampaign) GetNonprofitUrl() string {
	if o == nil || common.IsNil(o.NonprofitUrl) {
		var ret string
		return ret
	}
	return *o.NonprofitUrl
}

// GetNonprofitUrlOk returns a tuple with the NonprofitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetNonprofitUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.NonprofitUrl) {
		return nil, false
	}
	return o.NonprofitUrl, true
}

// HasNonprofitUrl returns a boolean if a field has been set.
func (o *DonationCampaign) HasNonprofitUrl() bool {
	if o != nil && !common.IsNil(o.NonprofitUrl) {
		return true
	}

	return false
}

// SetNonprofitUrl gets a reference to the given string and assigns it to the NonprofitUrl field.
func (o *DonationCampaign) SetNonprofitUrl(v string) {
	o.NonprofitUrl = &v
}

// GetTermsAndConditionsUrl returns the TermsAndConditionsUrl field value if set, zero value otherwise.
func (o *DonationCampaign) GetTermsAndConditionsUrl() string {
	if o == nil || common.IsNil(o.TermsAndConditionsUrl) {
		var ret string
		return ret
	}
	return *o.TermsAndConditionsUrl
}

// GetTermsAndConditionsUrlOk returns a tuple with the TermsAndConditionsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonationCampaign) GetTermsAndConditionsUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TermsAndConditionsUrl) {
		return nil, false
	}
	return o.TermsAndConditionsUrl, true
}

// HasTermsAndConditionsUrl returns a boolean if a field has been set.
func (o *DonationCampaign) HasTermsAndConditionsUrl() bool {
	if o != nil && !common.IsNil(o.TermsAndConditionsUrl) {
		return true
	}

	return false
}

// SetTermsAndConditionsUrl gets a reference to the given string and assigns it to the TermsAndConditionsUrl field.
func (o *DonationCampaign) SetTermsAndConditionsUrl(v string) {
	o.TermsAndConditionsUrl = &v
}

func (o DonationCampaign) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DonationCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amounts) {
		toSerialize["amounts"] = o.Amounts
	}
	if !common.IsNil(o.BannerUrl) {
		toSerialize["bannerUrl"] = o.BannerUrl
	}
	if !common.IsNil(o.CampaignName) {
		toSerialize["campaignName"] = o.CampaignName
	}
	if !common.IsNil(o.CauseName) {
		toSerialize["causeName"] = o.CauseName
	}
	if !common.IsNil(o.Donation) {
		toSerialize["donation"] = o.Donation
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.LogoUrl) {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if !common.IsNil(o.NonprofitDescription) {
		toSerialize["nonprofitDescription"] = o.NonprofitDescription
	}
	if !common.IsNil(o.NonprofitName) {
		toSerialize["nonprofitName"] = o.NonprofitName
	}
	if !common.IsNil(o.NonprofitUrl) {
		toSerialize["nonprofitUrl"] = o.NonprofitUrl
	}
	if !common.IsNil(o.TermsAndConditionsUrl) {
		toSerialize["termsAndConditionsUrl"] = o.TermsAndConditionsUrl
	}
	return toSerialize, nil
}

type NullableDonationCampaign struct {
	value *DonationCampaign
	isSet bool
}

func (v NullableDonationCampaign) Get() *DonationCampaign {
	return v.value
}

func (v *NullableDonationCampaign) Set(val *DonationCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableDonationCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableDonationCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonationCampaign(val *DonationCampaign) *NullableDonationCampaign {
	return &NullableDonationCampaign{value: val, isSet: true}
}

func (v NullableDonationCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonationCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
