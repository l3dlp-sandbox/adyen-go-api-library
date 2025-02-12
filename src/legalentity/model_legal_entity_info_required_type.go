/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the LegalEntityInfoRequiredType type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LegalEntityInfoRequiredType{}

// LegalEntityInfoRequiredType struct for LegalEntityInfoRequiredType
type LegalEntityInfoRequiredType struct {
	// Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Capabilities *map[string]LegalEntityCapability `json:"capabilities,omitempty"`
	// List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories.
	EntityAssociations []LegalEntityAssociation `json:"entityAssociations,omitempty"`
	Individual         *Individual              `json:"individual,omitempty"`
	Organization       *Organization            `json:"organization,omitempty"`
	// Your reference for the legal entity, maximum 150 characters.
	Reference          *string             `json:"reference,omitempty"`
	SoleProprietorship *SoleProprietorship `json:"soleProprietorship,omitempty"`
	// The type of legal entity.   Possible values: **individual**, **organization**, or **soleProprietorship**.
	Type string `json:"type"`
}

// NewLegalEntityInfoRequiredType instantiates a new LegalEntityInfoRequiredType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalEntityInfoRequiredType(type_ string) *LegalEntityInfoRequiredType {
	this := LegalEntityInfoRequiredType{}
	this.Type = type_
	return &this
}

// NewLegalEntityInfoRequiredTypeWithDefaults instantiates a new LegalEntityInfoRequiredType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalEntityInfoRequiredTypeWithDefaults() *LegalEntityInfoRequiredType {
	this := LegalEntityInfoRequiredType{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *LegalEntityInfoRequiredType) GetCapabilities() map[string]LegalEntityCapability {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret map[string]LegalEntityCapability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfoRequiredType) GetCapabilitiesOk() (*map[string]LegalEntityCapability, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *LegalEntityInfoRequiredType) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]LegalEntityCapability and assigns it to the Capabilities field.
func (o *LegalEntityInfoRequiredType) SetCapabilities(v map[string]LegalEntityCapability) {
	o.Capabilities = &v
}

// GetEntityAssociations returns the EntityAssociations field value if set, zero value otherwise.
func (o *LegalEntityInfoRequiredType) GetEntityAssociations() []LegalEntityAssociation {
	if o == nil || common.IsNil(o.EntityAssociations) {
		var ret []LegalEntityAssociation
		return ret
	}
	return o.EntityAssociations
}

// GetEntityAssociationsOk returns a tuple with the EntityAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfoRequiredType) GetEntityAssociationsOk() ([]LegalEntityAssociation, bool) {
	if o == nil || common.IsNil(o.EntityAssociations) {
		return nil, false
	}
	return o.EntityAssociations, true
}

// HasEntityAssociations returns a boolean if a field has been set.
func (o *LegalEntityInfoRequiredType) HasEntityAssociations() bool {
	if o != nil && !common.IsNil(o.EntityAssociations) {
		return true
	}

	return false
}

// SetEntityAssociations gets a reference to the given []LegalEntityAssociation and assigns it to the EntityAssociations field.
func (o *LegalEntityInfoRequiredType) SetEntityAssociations(v []LegalEntityAssociation) {
	o.EntityAssociations = v
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *LegalEntityInfoRequiredType) GetIndividual() Individual {
	if o == nil || common.IsNil(o.Individual) {
		var ret Individual
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfoRequiredType) GetIndividualOk() (*Individual, bool) {
	if o == nil || common.IsNil(o.Individual) {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *LegalEntityInfoRequiredType) HasIndividual() bool {
	if o != nil && !common.IsNil(o.Individual) {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given Individual and assigns it to the Individual field.
func (o *LegalEntityInfoRequiredType) SetIndividual(v Individual) {
	o.Individual = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *LegalEntityInfoRequiredType) GetOrganization() Organization {
	if o == nil || common.IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfoRequiredType) GetOrganizationOk() (*Organization, bool) {
	if o == nil || common.IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *LegalEntityInfoRequiredType) HasOrganization() bool {
	if o != nil && !common.IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *LegalEntityInfoRequiredType) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LegalEntityInfoRequiredType) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfoRequiredType) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LegalEntityInfoRequiredType) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LegalEntityInfoRequiredType) SetReference(v string) {
	o.Reference = &v
}

// GetSoleProprietorship returns the SoleProprietorship field value if set, zero value otherwise.
func (o *LegalEntityInfoRequiredType) GetSoleProprietorship() SoleProprietorship {
	if o == nil || common.IsNil(o.SoleProprietorship) {
		var ret SoleProprietorship
		return ret
	}
	return *o.SoleProprietorship
}

// GetSoleProprietorshipOk returns a tuple with the SoleProprietorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityInfoRequiredType) GetSoleProprietorshipOk() (*SoleProprietorship, bool) {
	if o == nil || common.IsNil(o.SoleProprietorship) {
		return nil, false
	}
	return o.SoleProprietorship, true
}

// HasSoleProprietorship returns a boolean if a field has been set.
func (o *LegalEntityInfoRequiredType) HasSoleProprietorship() bool {
	if o != nil && !common.IsNil(o.SoleProprietorship) {
		return true
	}

	return false
}

// SetSoleProprietorship gets a reference to the given SoleProprietorship and assigns it to the SoleProprietorship field.
func (o *LegalEntityInfoRequiredType) SetSoleProprietorship(v SoleProprietorship) {
	o.SoleProprietorship = &v
}

// GetType returns the Type field value
func (o *LegalEntityInfoRequiredType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LegalEntityInfoRequiredType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LegalEntityInfoRequiredType) SetType(v string) {
	o.Type = v
}

func (o LegalEntityInfoRequiredType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalEntityInfoRequiredType) ToMap() (map[string]interface{}, error) {
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
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableLegalEntityInfoRequiredType struct {
	value *LegalEntityInfoRequiredType
	isSet bool
}

func (v NullableLegalEntityInfoRequiredType) Get() *LegalEntityInfoRequiredType {
	return v.value
}

func (v *NullableLegalEntityInfoRequiredType) Set(val *LegalEntityInfoRequiredType) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalEntityInfoRequiredType) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalEntityInfoRequiredType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalEntityInfoRequiredType(val *LegalEntityInfoRequiredType) *NullableLegalEntityInfoRequiredType {
	return &NullableLegalEntityInfoRequiredType{value: val, isSet: true}
}

func (v NullableLegalEntityInfoRequiredType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalEntityInfoRequiredType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *LegalEntityInfoRequiredType) isValidType() bool {
	var allowedEnumValues = []string{"individual", "organization", "soleProprietorship", "trust", "unincorporatedPartnership"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
