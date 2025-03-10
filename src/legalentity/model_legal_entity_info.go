/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the LegalEntityInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LegalEntityInfo{}

// LegalEntityInfo struct for LegalEntityInfo
type LegalEntityInfo struct {
	// Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Capabilities *map[string]LegalEntityCapability `json:"capabilities,omitempty"`
	// List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories.
	EntityAssociations []LegalEntityAssociation `json:"entityAssociations,omitempty"`
	Individual         *Individual              `json:"individual,omitempty"`
	Organization       *Organization            `json:"organization,omitempty"`
	// Your reference for the legal entity, maximum 150 characters.
	Reference          *string             `json:"reference,omitempty"`
	SoleProprietorship *SoleProprietorship `json:"soleProprietorship,omitempty"`
	Trust              *Trust              `json:"trust,omitempty"`
	// The type of legal entity.  Possible values: **individual**, **organization**, **soleProprietorship**, or **trust**.
	Type                      *string                    `json:"type,omitempty"`
	UnincorporatedPartnership *UnincorporatedPartnership `json:"unincorporatedPartnership,omitempty"`
	// A key-value pair that specifies the verification process for a legal entity. Set to **upfront** for upfront verification for [marketplaces](https://docs.adyen.com/marketplaces/verification-overview/verification-types/#upfront-verification).
	VerificationPlan *string `json:"verificationPlan,omitempty"`
}

// NewLegalEntityInfo instantiates a new LegalEntityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalEntityInfo() *LegalEntityInfo {
	this := LegalEntityInfo{}
	return &this
}

// NewLegalEntityInfoWithDefaults instantiates a new LegalEntityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalEntityInfoWithDefaults() *LegalEntityInfo {
	this := LegalEntityInfo{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetCapabilities() map[string]LegalEntityCapability {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret map[string]LegalEntityCapability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetCapabilitiesOk() (*map[string]LegalEntityCapability, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]LegalEntityCapability and assigns it to the Capabilities field.
func (o *LegalEntityInfo) SetCapabilities(v map[string]LegalEntityCapability) {
	o.Capabilities = &v
}

// GetEntityAssociations returns the EntityAssociations field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetEntityAssociations() []LegalEntityAssociation {
	if o == nil || common.IsNil(o.EntityAssociations) {
		var ret []LegalEntityAssociation
		return ret
	}
	return o.EntityAssociations
}

// GetEntityAssociationsOk returns a tuple with the EntityAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetEntityAssociationsOk() ([]LegalEntityAssociation, bool) {
	if o == nil || common.IsNil(o.EntityAssociations) {
		return nil, false
	}
	return o.EntityAssociations, true
}

// HasEntityAssociations returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasEntityAssociations() bool {
	if o != nil && !common.IsNil(o.EntityAssociations) {
		return true
	}

	return false
}

// SetEntityAssociations gets a reference to the given []LegalEntityAssociation and assigns it to the EntityAssociations field.
func (o *LegalEntityInfo) SetEntityAssociations(v []LegalEntityAssociation) {
	o.EntityAssociations = v
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetIndividual() Individual {
	if o == nil || common.IsNil(o.Individual) {
		var ret Individual
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetIndividualOk() (*Individual, bool) {
	if o == nil || common.IsNil(o.Individual) {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasIndividual() bool {
	if o != nil && !common.IsNil(o.Individual) {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given Individual and assigns it to the Individual field.
func (o *LegalEntityInfo) SetIndividual(v Individual) {
	o.Individual = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetOrganization() Organization {
	if o == nil || common.IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetOrganizationOk() (*Organization, bool) {
	if o == nil || common.IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasOrganization() bool {
	if o != nil && !common.IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *LegalEntityInfo) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LegalEntityInfo) SetReference(v string) {
	o.Reference = &v
}

// GetSoleProprietorship returns the SoleProprietorship field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetSoleProprietorship() SoleProprietorship {
	if o == nil || common.IsNil(o.SoleProprietorship) {
		var ret SoleProprietorship
		return ret
	}
	return *o.SoleProprietorship
}

// GetSoleProprietorshipOk returns a tuple with the SoleProprietorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetSoleProprietorshipOk() (*SoleProprietorship, bool) {
	if o == nil || common.IsNil(o.SoleProprietorship) {
		return nil, false
	}
	return o.SoleProprietorship, true
}

// HasSoleProprietorship returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasSoleProprietorship() bool {
	if o != nil && !common.IsNil(o.SoleProprietorship) {
		return true
	}

	return false
}

// SetSoleProprietorship gets a reference to the given SoleProprietorship and assigns it to the SoleProprietorship field.
func (o *LegalEntityInfo) SetSoleProprietorship(v SoleProprietorship) {
	o.SoleProprietorship = &v
}

// GetTrust returns the Trust field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetTrust() Trust {
	if o == nil || common.IsNil(o.Trust) {
		var ret Trust
		return ret
	}
	return *o.Trust
}

// GetTrustOk returns a tuple with the Trust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetTrustOk() (*Trust, bool) {
	if o == nil || common.IsNil(o.Trust) {
		return nil, false
	}
	return o.Trust, true
}

// HasTrust returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasTrust() bool {
	if o != nil && !common.IsNil(o.Trust) {
		return true
	}

	return false
}

// SetTrust gets a reference to the given Trust and assigns it to the Trust field.
func (o *LegalEntityInfo) SetTrust(v Trust) {
	o.Trust = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LegalEntityInfo) SetType(v string) {
	o.Type = &v
}

// GetUnincorporatedPartnership returns the UnincorporatedPartnership field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetUnincorporatedPartnership() UnincorporatedPartnership {
	if o == nil || common.IsNil(o.UnincorporatedPartnership) {
		var ret UnincorporatedPartnership
		return ret
	}
	return *o.UnincorporatedPartnership
}

// GetUnincorporatedPartnershipOk returns a tuple with the UnincorporatedPartnership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetUnincorporatedPartnershipOk() (*UnincorporatedPartnership, bool) {
	if o == nil || common.IsNil(o.UnincorporatedPartnership) {
		return nil, false
	}
	return o.UnincorporatedPartnership, true
}

// HasUnincorporatedPartnership returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasUnincorporatedPartnership() bool {
	if o != nil && !common.IsNil(o.UnincorporatedPartnership) {
		return true
	}

	return false
}

// SetUnincorporatedPartnership gets a reference to the given UnincorporatedPartnership and assigns it to the UnincorporatedPartnership field.
func (o *LegalEntityInfo) SetUnincorporatedPartnership(v UnincorporatedPartnership) {
	o.UnincorporatedPartnership = &v
}

// GetVerificationPlan returns the VerificationPlan field value if set, zero value otherwise.
func (o *LegalEntityInfo) GetVerificationPlan() string {
	if o == nil || common.IsNil(o.VerificationPlan) {
		var ret string
		return ret
	}
	return *o.VerificationPlan
}

// GetVerificationPlanOk returns a tuple with the VerificationPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfo) GetVerificationPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationPlan) {
		return nil, false
	}
	return o.VerificationPlan, true
}

// HasVerificationPlan returns a boolean if a field has been set.
func (o *LegalEntityInfo) HasVerificationPlan() bool {
	if o != nil && !common.IsNil(o.VerificationPlan) {
		return true
	}

	return false
}

// SetVerificationPlan gets a reference to the given string and assigns it to the VerificationPlan field.
func (o *LegalEntityInfo) SetVerificationPlan(v string) {
	o.VerificationPlan = &v
}

func (o LegalEntityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalEntityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !common.IsNil(o.EntityAssociations) {
		toSerialize["entityAssociations"] = o.EntityAssociations
	}
	if !common.IsNil(o.Individual) {
		toSerialize["individual"] = o.Individual
	}
	if !common.IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.SoleProprietorship) {
		toSerialize["soleProprietorship"] = o.SoleProprietorship
	}
	if !common.IsNil(o.Trust) {
		toSerialize["trust"] = o.Trust
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UnincorporatedPartnership) {
		toSerialize["unincorporatedPartnership"] = o.UnincorporatedPartnership
	}
	if !common.IsNil(o.VerificationPlan) {
		toSerialize["verificationPlan"] = o.VerificationPlan
	}
	return toSerialize, nil
}

type NullableLegalEntityInfo struct {
	value *LegalEntityInfo
	isSet bool
}

func (v NullableLegalEntityInfo) Get() *LegalEntityInfo {
	return v.value
}

func (v *NullableLegalEntityInfo) Set(val *LegalEntityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalEntityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalEntityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalEntityInfo(val *LegalEntityInfo) *NullableLegalEntityInfo {
	return &NullableLegalEntityInfo{value: val, isSet: true}
}

func (v NullableLegalEntityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalEntityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *LegalEntityInfo) isValidType() bool {
	var allowedEnumValues = []string{"individual", "organization", "soleProprietorship", "trust", "unincorporatedPartnership"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
