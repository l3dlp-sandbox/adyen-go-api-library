/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the OnboardingLinkSettings type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OnboardingLinkSettings{}

// OnboardingLinkSettings struct for OnboardingLinkSettings
type OnboardingLinkSettings struct {
	AcceptedCountries                        []string `json:"acceptedCountries,omitempty"`
	AllowBankAccountFormatSelection          *bool    `json:"allowBankAccountFormatSelection,omitempty"`
	AllowIntraRegionCrossBorderPayout        *bool    `json:"allowIntraRegionCrossBorderPayout,omitempty"`
	ChangeLegalEntityType                    *bool    `json:"changeLegalEntityType,omitempty"`
	EditPrefilledCountry                     *bool    `json:"editPrefilledCountry,omitempty"`
	HideOnboardingIntroductionIndividual     *bool    `json:"hideOnboardingIntroductionIndividual,omitempty"`
	HideOnboardingIntroductionOrganization   *bool    `json:"hideOnboardingIntroductionOrganization,omitempty"`
	HideOnboardingIntroductionSoleProprietor *bool    `json:"hideOnboardingIntroductionSoleProprietor,omitempty"`
	HideOnboardingIntroductionTrust          *bool    `json:"hideOnboardingIntroductionTrust,omitempty"`
	InstantBankVerification                  *bool    `json:"instantBankVerification,omitempty"`
	RequirePciSignEcomMoto                   *bool    `json:"requirePciSignEcomMoto,omitempty"`
	RequirePciSignEcommerce                  *bool    `json:"requirePciSignEcommerce,omitempty"`
	RequirePciSignPos                        *bool    `json:"requirePciSignPos,omitempty"`
	RequirePciSignPosMoto                    *bool    `json:"requirePciSignPosMoto,omitempty"`
	TransferInstrumentLimit                  *int32   `json:"transferInstrumentLimit,omitempty"`
}

// NewOnboardingLinkSettings instantiates a new OnboardingLinkSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingLinkSettings() *OnboardingLinkSettings {
	this := OnboardingLinkSettings{}
	return &this
}

// NewOnboardingLinkSettingsWithDefaults instantiates a new OnboardingLinkSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingLinkSettingsWithDefaults() *OnboardingLinkSettings {
	this := OnboardingLinkSettings{}
	return &this
}

// GetAcceptedCountries returns the AcceptedCountries field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetAcceptedCountries() []string {
	if o == nil || common.IsNil(o.AcceptedCountries) {
		var ret []string
		return ret
	}
	return o.AcceptedCountries
}

// GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetAcceptedCountriesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AcceptedCountries) {
		return nil, false
	}
	return o.AcceptedCountries, true
}

// HasAcceptedCountries returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasAcceptedCountries() bool {
	if o != nil && !common.IsNil(o.AcceptedCountries) {
		return true
	}

	return false
}

// SetAcceptedCountries gets a reference to the given []string and assigns it to the AcceptedCountries field.
func (o *OnboardingLinkSettings) SetAcceptedCountries(v []string) {
	o.AcceptedCountries = v
}

// GetAllowBankAccountFormatSelection returns the AllowBankAccountFormatSelection field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetAllowBankAccountFormatSelection() bool {
	if o == nil || common.IsNil(o.AllowBankAccountFormatSelection) {
		var ret bool
		return ret
	}
	return *o.AllowBankAccountFormatSelection
}

// GetAllowBankAccountFormatSelectionOk returns a tuple with the AllowBankAccountFormatSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetAllowBankAccountFormatSelectionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AllowBankAccountFormatSelection) {
		return nil, false
	}
	return o.AllowBankAccountFormatSelection, true
}

// HasAllowBankAccountFormatSelection returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasAllowBankAccountFormatSelection() bool {
	if o != nil && !common.IsNil(o.AllowBankAccountFormatSelection) {
		return true
	}

	return false
}

// SetAllowBankAccountFormatSelection gets a reference to the given bool and assigns it to the AllowBankAccountFormatSelection field.
func (o *OnboardingLinkSettings) SetAllowBankAccountFormatSelection(v bool) {
	o.AllowBankAccountFormatSelection = &v
}

// GetAllowIntraRegionCrossBorderPayout returns the AllowIntraRegionCrossBorderPayout field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetAllowIntraRegionCrossBorderPayout() bool {
	if o == nil || common.IsNil(o.AllowIntraRegionCrossBorderPayout) {
		var ret bool
		return ret
	}
	return *o.AllowIntraRegionCrossBorderPayout
}

// GetAllowIntraRegionCrossBorderPayoutOk returns a tuple with the AllowIntraRegionCrossBorderPayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetAllowIntraRegionCrossBorderPayoutOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AllowIntraRegionCrossBorderPayout) {
		return nil, false
	}
	return o.AllowIntraRegionCrossBorderPayout, true
}

// HasAllowIntraRegionCrossBorderPayout returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasAllowIntraRegionCrossBorderPayout() bool {
	if o != nil && !common.IsNil(o.AllowIntraRegionCrossBorderPayout) {
		return true
	}

	return false
}

// SetAllowIntraRegionCrossBorderPayout gets a reference to the given bool and assigns it to the AllowIntraRegionCrossBorderPayout field.
func (o *OnboardingLinkSettings) SetAllowIntraRegionCrossBorderPayout(v bool) {
	o.AllowIntraRegionCrossBorderPayout = &v
}

// GetChangeLegalEntityType returns the ChangeLegalEntityType field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetChangeLegalEntityType() bool {
	if o == nil || common.IsNil(o.ChangeLegalEntityType) {
		var ret bool
		return ret
	}
	return *o.ChangeLegalEntityType
}

// GetChangeLegalEntityTypeOk returns a tuple with the ChangeLegalEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetChangeLegalEntityTypeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ChangeLegalEntityType) {
		return nil, false
	}
	return o.ChangeLegalEntityType, true
}

// HasChangeLegalEntityType returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasChangeLegalEntityType() bool {
	if o != nil && !common.IsNil(o.ChangeLegalEntityType) {
		return true
	}

	return false
}

// SetChangeLegalEntityType gets a reference to the given bool and assigns it to the ChangeLegalEntityType field.
func (o *OnboardingLinkSettings) SetChangeLegalEntityType(v bool) {
	o.ChangeLegalEntityType = &v
}

// GetEditPrefilledCountry returns the EditPrefilledCountry field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetEditPrefilledCountry() bool {
	if o == nil || common.IsNil(o.EditPrefilledCountry) {
		var ret bool
		return ret
	}
	return *o.EditPrefilledCountry
}

// GetEditPrefilledCountryOk returns a tuple with the EditPrefilledCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetEditPrefilledCountryOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EditPrefilledCountry) {
		return nil, false
	}
	return o.EditPrefilledCountry, true
}

// HasEditPrefilledCountry returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasEditPrefilledCountry() bool {
	if o != nil && !common.IsNil(o.EditPrefilledCountry) {
		return true
	}

	return false
}

// SetEditPrefilledCountry gets a reference to the given bool and assigns it to the EditPrefilledCountry field.
func (o *OnboardingLinkSettings) SetEditPrefilledCountry(v bool) {
	o.EditPrefilledCountry = &v
}

// GetHideOnboardingIntroductionIndividual returns the HideOnboardingIntroductionIndividual field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionIndividual() bool {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionIndividual) {
		var ret bool
		return ret
	}
	return *o.HideOnboardingIntroductionIndividual
}

// GetHideOnboardingIntroductionIndividualOk returns a tuple with the HideOnboardingIntroductionIndividual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionIndividualOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionIndividual) {
		return nil, false
	}
	return o.HideOnboardingIntroductionIndividual, true
}

// HasHideOnboardingIntroductionIndividual returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasHideOnboardingIntroductionIndividual() bool {
	if o != nil && !common.IsNil(o.HideOnboardingIntroductionIndividual) {
		return true
	}

	return false
}

// SetHideOnboardingIntroductionIndividual gets a reference to the given bool and assigns it to the HideOnboardingIntroductionIndividual field.
func (o *OnboardingLinkSettings) SetHideOnboardingIntroductionIndividual(v bool) {
	o.HideOnboardingIntroductionIndividual = &v
}

// GetHideOnboardingIntroductionOrganization returns the HideOnboardingIntroductionOrganization field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionOrganization() bool {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionOrganization) {
		var ret bool
		return ret
	}
	return *o.HideOnboardingIntroductionOrganization
}

// GetHideOnboardingIntroductionOrganizationOk returns a tuple with the HideOnboardingIntroductionOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionOrganizationOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionOrganization) {
		return nil, false
	}
	return o.HideOnboardingIntroductionOrganization, true
}

// HasHideOnboardingIntroductionOrganization returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasHideOnboardingIntroductionOrganization() bool {
	if o != nil && !common.IsNil(o.HideOnboardingIntroductionOrganization) {
		return true
	}

	return false
}

// SetHideOnboardingIntroductionOrganization gets a reference to the given bool and assigns it to the HideOnboardingIntroductionOrganization field.
func (o *OnboardingLinkSettings) SetHideOnboardingIntroductionOrganization(v bool) {
	o.HideOnboardingIntroductionOrganization = &v
}

// GetHideOnboardingIntroductionSoleProprietor returns the HideOnboardingIntroductionSoleProprietor field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionSoleProprietor() bool {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionSoleProprietor) {
		var ret bool
		return ret
	}
	return *o.HideOnboardingIntroductionSoleProprietor
}

// GetHideOnboardingIntroductionSoleProprietorOk returns a tuple with the HideOnboardingIntroductionSoleProprietor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionSoleProprietorOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionSoleProprietor) {
		return nil, false
	}
	return o.HideOnboardingIntroductionSoleProprietor, true
}

// HasHideOnboardingIntroductionSoleProprietor returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasHideOnboardingIntroductionSoleProprietor() bool {
	if o != nil && !common.IsNil(o.HideOnboardingIntroductionSoleProprietor) {
		return true
	}

	return false
}

// SetHideOnboardingIntroductionSoleProprietor gets a reference to the given bool and assigns it to the HideOnboardingIntroductionSoleProprietor field.
func (o *OnboardingLinkSettings) SetHideOnboardingIntroductionSoleProprietor(v bool) {
	o.HideOnboardingIntroductionSoleProprietor = &v
}

// GetHideOnboardingIntroductionTrust returns the HideOnboardingIntroductionTrust field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionTrust() bool {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionTrust) {
		var ret bool
		return ret
	}
	return *o.HideOnboardingIntroductionTrust
}

// GetHideOnboardingIntroductionTrustOk returns a tuple with the HideOnboardingIntroductionTrust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetHideOnboardingIntroductionTrustOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HideOnboardingIntroductionTrust) {
		return nil, false
	}
	return o.HideOnboardingIntroductionTrust, true
}

// HasHideOnboardingIntroductionTrust returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasHideOnboardingIntroductionTrust() bool {
	if o != nil && !common.IsNil(o.HideOnboardingIntroductionTrust) {
		return true
	}

	return false
}

// SetHideOnboardingIntroductionTrust gets a reference to the given bool and assigns it to the HideOnboardingIntroductionTrust field.
func (o *OnboardingLinkSettings) SetHideOnboardingIntroductionTrust(v bool) {
	o.HideOnboardingIntroductionTrust = &v
}

// GetInstantBankVerification returns the InstantBankVerification field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetInstantBankVerification() bool {
	if o == nil || common.IsNil(o.InstantBankVerification) {
		var ret bool
		return ret
	}
	return *o.InstantBankVerification
}

// GetInstantBankVerificationOk returns a tuple with the InstantBankVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetInstantBankVerificationOk() (*bool, bool) {
	if o == nil || common.IsNil(o.InstantBankVerification) {
		return nil, false
	}
	return o.InstantBankVerification, true
}

// HasInstantBankVerification returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasInstantBankVerification() bool {
	if o != nil && !common.IsNil(o.InstantBankVerification) {
		return true
	}

	return false
}

// SetInstantBankVerification gets a reference to the given bool and assigns it to the InstantBankVerification field.
func (o *OnboardingLinkSettings) SetInstantBankVerification(v bool) {
	o.InstantBankVerification = &v
}

// GetRequirePciSignEcomMoto returns the RequirePciSignEcomMoto field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetRequirePciSignEcomMoto() bool {
	if o == nil || common.IsNil(o.RequirePciSignEcomMoto) {
		var ret bool
		return ret
	}
	return *o.RequirePciSignEcomMoto
}

// GetRequirePciSignEcomMotoOk returns a tuple with the RequirePciSignEcomMoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetRequirePciSignEcomMotoOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RequirePciSignEcomMoto) {
		return nil, false
	}
	return o.RequirePciSignEcomMoto, true
}

// HasRequirePciSignEcomMoto returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasRequirePciSignEcomMoto() bool {
	if o != nil && !common.IsNil(o.RequirePciSignEcomMoto) {
		return true
	}

	return false
}

// SetRequirePciSignEcomMoto gets a reference to the given bool and assigns it to the RequirePciSignEcomMoto field.
func (o *OnboardingLinkSettings) SetRequirePciSignEcomMoto(v bool) {
	o.RequirePciSignEcomMoto = &v
}

// GetRequirePciSignEcommerce returns the RequirePciSignEcommerce field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetRequirePciSignEcommerce() bool {
	if o == nil || common.IsNil(o.RequirePciSignEcommerce) {
		var ret bool
		return ret
	}
	return *o.RequirePciSignEcommerce
}

// GetRequirePciSignEcommerceOk returns a tuple with the RequirePciSignEcommerce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetRequirePciSignEcommerceOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RequirePciSignEcommerce) {
		return nil, false
	}
	return o.RequirePciSignEcommerce, true
}

// HasRequirePciSignEcommerce returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasRequirePciSignEcommerce() bool {
	if o != nil && !common.IsNil(o.RequirePciSignEcommerce) {
		return true
	}

	return false
}

// SetRequirePciSignEcommerce gets a reference to the given bool and assigns it to the RequirePciSignEcommerce field.
func (o *OnboardingLinkSettings) SetRequirePciSignEcommerce(v bool) {
	o.RequirePciSignEcommerce = &v
}

// GetRequirePciSignPos returns the RequirePciSignPos field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetRequirePciSignPos() bool {
	if o == nil || common.IsNil(o.RequirePciSignPos) {
		var ret bool
		return ret
	}
	return *o.RequirePciSignPos
}

// GetRequirePciSignPosOk returns a tuple with the RequirePciSignPos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetRequirePciSignPosOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RequirePciSignPos) {
		return nil, false
	}
	return o.RequirePciSignPos, true
}

// HasRequirePciSignPos returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasRequirePciSignPos() bool {
	if o != nil && !common.IsNil(o.RequirePciSignPos) {
		return true
	}

	return false
}

// SetRequirePciSignPos gets a reference to the given bool and assigns it to the RequirePciSignPos field.
func (o *OnboardingLinkSettings) SetRequirePciSignPos(v bool) {
	o.RequirePciSignPos = &v
}

// GetRequirePciSignPosMoto returns the RequirePciSignPosMoto field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetRequirePciSignPosMoto() bool {
	if o == nil || common.IsNil(o.RequirePciSignPosMoto) {
		var ret bool
		return ret
	}
	return *o.RequirePciSignPosMoto
}

// GetRequirePciSignPosMotoOk returns a tuple with the RequirePciSignPosMoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetRequirePciSignPosMotoOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RequirePciSignPosMoto) {
		return nil, false
	}
	return o.RequirePciSignPosMoto, true
}

// HasRequirePciSignPosMoto returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasRequirePciSignPosMoto() bool {
	if o != nil && !common.IsNil(o.RequirePciSignPosMoto) {
		return true
	}

	return false
}

// SetRequirePciSignPosMoto gets a reference to the given bool and assigns it to the RequirePciSignPosMoto field.
func (o *OnboardingLinkSettings) SetRequirePciSignPosMoto(v bool) {
	o.RequirePciSignPosMoto = &v
}

// GetTransferInstrumentLimit returns the TransferInstrumentLimit field value if set, zero value otherwise.
func (o *OnboardingLinkSettings) GetTransferInstrumentLimit() int32 {
	if o == nil || common.IsNil(o.TransferInstrumentLimit) {
		var ret int32
		return ret
	}
	return *o.TransferInstrumentLimit
}

// GetTransferInstrumentLimitOk returns a tuple with the TransferInstrumentLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLinkSettings) GetTransferInstrumentLimitOk() (*int32, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentLimit) {
		return nil, false
	}
	return o.TransferInstrumentLimit, true
}

// HasTransferInstrumentLimit returns a boolean if a field has been set.
func (o *OnboardingLinkSettings) HasTransferInstrumentLimit() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentLimit) {
		return true
	}

	return false
}

// SetTransferInstrumentLimit gets a reference to the given int32 and assigns it to the TransferInstrumentLimit field.
func (o *OnboardingLinkSettings) SetTransferInstrumentLimit(v int32) {
	o.TransferInstrumentLimit = &v
}

func (o OnboardingLinkSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingLinkSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcceptedCountries) {
		toSerialize["acceptedCountries"] = o.AcceptedCountries
	}
	if !common.IsNil(o.AllowBankAccountFormatSelection) {
		toSerialize["allowBankAccountFormatSelection"] = o.AllowBankAccountFormatSelection
	}
	if !common.IsNil(o.AllowIntraRegionCrossBorderPayout) {
		toSerialize["allowIntraRegionCrossBorderPayout"] = o.AllowIntraRegionCrossBorderPayout
	}
	if !common.IsNil(o.ChangeLegalEntityType) {
		toSerialize["changeLegalEntityType"] = o.ChangeLegalEntityType
	}
	if !common.IsNil(o.EditPrefilledCountry) {
		toSerialize["editPrefilledCountry"] = o.EditPrefilledCountry
	}
	if !common.IsNil(o.HideOnboardingIntroductionIndividual) {
		toSerialize["hideOnboardingIntroductionIndividual"] = o.HideOnboardingIntroductionIndividual
	}
	if !common.IsNil(o.HideOnboardingIntroductionOrganization) {
		toSerialize["hideOnboardingIntroductionOrganization"] = o.HideOnboardingIntroductionOrganization
	}
	if !common.IsNil(o.HideOnboardingIntroductionSoleProprietor) {
		toSerialize["hideOnboardingIntroductionSoleProprietor"] = o.HideOnboardingIntroductionSoleProprietor
	}
	if !common.IsNil(o.HideOnboardingIntroductionTrust) {
		toSerialize["hideOnboardingIntroductionTrust"] = o.HideOnboardingIntroductionTrust
	}
	if !common.IsNil(o.InstantBankVerification) {
		toSerialize["instantBankVerification"] = o.InstantBankVerification
	}
	if !common.IsNil(o.RequirePciSignEcomMoto) {
		toSerialize["requirePciSignEcomMoto"] = o.RequirePciSignEcomMoto
	}
	if !common.IsNil(o.RequirePciSignEcommerce) {
		toSerialize["requirePciSignEcommerce"] = o.RequirePciSignEcommerce
	}
	if !common.IsNil(o.RequirePciSignPos) {
		toSerialize["requirePciSignPos"] = o.RequirePciSignPos
	}
	if !common.IsNil(o.RequirePciSignPosMoto) {
		toSerialize["requirePciSignPosMoto"] = o.RequirePciSignPosMoto
	}
	if !common.IsNil(o.TransferInstrumentLimit) {
		toSerialize["transferInstrumentLimit"] = o.TransferInstrumentLimit
	}
	return toSerialize, nil
}

type NullableOnboardingLinkSettings struct {
	value *OnboardingLinkSettings
	isSet bool
}

func (v NullableOnboardingLinkSettings) Get() *OnboardingLinkSettings {
	return v.value
}

func (v *NullableOnboardingLinkSettings) Set(val *OnboardingLinkSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingLinkSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingLinkSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingLinkSettings(val *OnboardingLinkSettings) *NullableOnboardingLinkSettings {
	return &NullableOnboardingLinkSettings{value: val, isSet: true}
}

func (v NullableOnboardingLinkSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingLinkSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
