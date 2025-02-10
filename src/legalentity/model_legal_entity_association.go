/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the LegalEntityAssociation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LegalEntityAssociation{}

// LegalEntityAssociation struct for LegalEntityAssociation
type LegalEntityAssociation struct {
	// The unique identifier of another legal entity with which the `legalEntityId` is associated. When the `legalEntityId` is associated to legal entities other than the current one, the response returns all the associations.
	AssociatorId *string `json:"associatorId,omitempty"`
	// The legal entity type of associated legal entity.  For example, **organization**, **soleProprietorship** or **individual**.
	EntityType *string `json:"entityType,omitempty"`
	// The individual's job title if the `type` is **uboThroughControl** or **signatory**.
	JobTitle *string `json:"jobTitle,omitempty"`
	// The unique identifier of the associated [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id).
	LegalEntityId string `json:"legalEntityId"`
	// The name of the associated [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id).  - For **individual**, `name.firstName` and `name.lastName`. - For **organization**, `legalName`. - For **soleProprietorship**, `name`.
	Name *string `json:"name,omitempty"`
	// Default value: **false**Indicates if the `type` **director**, **secondaryPartner** or **shareholder** is a nominee. Only applicable to New Zealand.
	Nominee *bool `json:"nominee,omitempty"`
	// The individual's relationship to a legal representative if the `type` is **legalRepresentative**. Possible values: **parent**, **guardian**.
	Relationship *string `json:"relationship,omitempty"`
	// Defines the KYC exemption reason for a settlor associated with a trust. Only applicable to trusts in Australia.  For example, **professionalServiceProvider**, **deceased**, or **contributionBelowThreshold**.
	SettlorExemptionReason []string `json:"settlorExemptionReason,omitempty"`
	// Defines the relationship of the legal entity to the current legal entity.  Possible value for individuals: **legalRepresentative**.  Possible values for organizations: **director**, **signatory**, **trustOwnership**, **uboThroughOwnership**, **uboThroughControl**, or **ultimateParentCompany**.  Possible values for sole proprietorships: **soleProprietorship**.  Possible value for trusts: **trust**.  Possible values for trust members: **definedBeneficiary**, **protector**, **secondaryTrustee**, **settlor**, **uboThroughControl**, or **uboThroughOwnership**.  Possible value for unincorporated partnership: **unincorporatedPartnership**.  Possible values for unincorporated partnership members: **secondaryPartner**, **uboThroughControl**, **uboThroughOwnership**
	Type string `json:"type"`
}

// NewLegalEntityAssociation instantiates a new LegalEntityAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalEntityAssociation(legalEntityId string, type_ string) *LegalEntityAssociation {
	this := LegalEntityAssociation{}
	this.LegalEntityId = legalEntityId
	this.Type = type_
	return &this
}

// NewLegalEntityAssociationWithDefaults instantiates a new LegalEntityAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalEntityAssociationWithDefaults() *LegalEntityAssociation {
	this := LegalEntityAssociation{}
	return &this
}

// GetAssociatorId returns the AssociatorId field value if set, zero value otherwise.
func (o *LegalEntityAssociation) GetAssociatorId() string {
	if o == nil || common.IsNil(o.AssociatorId) {
		var ret string
		return ret
	}
	return *o.AssociatorId
}

// GetAssociatorIdOk returns a tuple with the AssociatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetAssociatorIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssociatorId) {
		return nil, false
	}
	return o.AssociatorId, true
}

// HasAssociatorId returns a boolean if a field has been set.
func (o *LegalEntityAssociation) HasAssociatorId() bool {
	if o != nil && !common.IsNil(o.AssociatorId) {
		return true
	}

	return false
}

// SetAssociatorId gets a reference to the given string and assigns it to the AssociatorId field.
func (o *LegalEntityAssociation) SetAssociatorId(v string) {
	o.AssociatorId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *LegalEntityAssociation) GetEntityType() string {
	if o == nil || common.IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetEntityTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *LegalEntityAssociation) HasEntityType() bool {
	if o != nil && !common.IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *LegalEntityAssociation) SetEntityType(v string) {
	o.EntityType = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *LegalEntityAssociation) GetJobTitle() string {
	if o == nil || common.IsNil(o.JobTitle) {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetJobTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.JobTitle) {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *LegalEntityAssociation) HasJobTitle() bool {
	if o != nil && !common.IsNil(o.JobTitle) {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *LegalEntityAssociation) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLegalEntityId returns the LegalEntityId field value
func (o *LegalEntityAssociation) GetLegalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetLegalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalEntityId, true
}

// SetLegalEntityId sets field value
func (o *LegalEntityAssociation) SetLegalEntityId(v string) {
	o.LegalEntityId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LegalEntityAssociation) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LegalEntityAssociation) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LegalEntityAssociation) SetName(v string) {
	o.Name = &v
}

// GetNominee returns the Nominee field value if set, zero value otherwise.
func (o *LegalEntityAssociation) GetNominee() bool {
	if o == nil || common.IsNil(o.Nominee) {
		var ret bool
		return ret
	}
	return *o.Nominee
}

// GetNomineeOk returns a tuple with the Nominee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetNomineeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Nominee) {
		return nil, false
	}
	return o.Nominee, true
}

// HasNominee returns a boolean if a field has been set.
func (o *LegalEntityAssociation) HasNominee() bool {
	if o != nil && !common.IsNil(o.Nominee) {
		return true
	}

	return false
}

// SetNominee gets a reference to the given bool and assigns it to the Nominee field.
func (o *LegalEntityAssociation) SetNominee(v bool) {
	o.Nominee = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *LegalEntityAssociation) GetRelationship() string {
	if o == nil || common.IsNil(o.Relationship) {
		var ret string
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetRelationshipOk() (*string, bool) {
	if o == nil || common.IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *LegalEntityAssociation) HasRelationship() bool {
	if o != nil && !common.IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given string and assigns it to the Relationship field.
func (o *LegalEntityAssociation) SetRelationship(v string) {
	o.Relationship = &v
}

// GetSettlorExemptionReason returns the SettlorExemptionReason field value if set, zero value otherwise.
func (o *LegalEntityAssociation) GetSettlorExemptionReason() []string {
	if o == nil || common.IsNil(o.SettlorExemptionReason) {
		var ret []string
		return ret
	}
	return o.SettlorExemptionReason
}

// GetSettlorExemptionReasonOk returns a tuple with the SettlorExemptionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetSettlorExemptionReasonOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SettlorExemptionReason) {
		return nil, false
	}
	return o.SettlorExemptionReason, true
}

// HasSettlorExemptionReason returns a boolean if a field has been set.
func (o *LegalEntityAssociation) HasSettlorExemptionReason() bool {
	if o != nil && !common.IsNil(o.SettlorExemptionReason) {
		return true
	}

	return false
}

// SetSettlorExemptionReason gets a reference to the given []string and assigns it to the SettlorExemptionReason field.
func (o *LegalEntityAssociation) SetSettlorExemptionReason(v []string) {
	o.SettlorExemptionReason = v
}

// GetType returns the Type field value
func (o *LegalEntityAssociation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LegalEntityAssociation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LegalEntityAssociation) SetType(v string) {
	o.Type = v
}

func (o LegalEntityAssociation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalEntityAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AssociatorId) {
		toSerialize["associatorId"] = o.AssociatorId
	}
	if !common.IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !common.IsNil(o.JobTitle) {
		toSerialize["jobTitle"] = o.JobTitle
	}
	toSerialize["legalEntityId"] = o.LegalEntityId
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Nominee) {
		toSerialize["nominee"] = o.Nominee
	}
	if !common.IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	if !common.IsNil(o.SettlorExemptionReason) {
		toSerialize["settlorExemptionReason"] = o.SettlorExemptionReason
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableLegalEntityAssociation struct {
	value *LegalEntityAssociation
	isSet bool
}

func (v NullableLegalEntityAssociation) Get() *LegalEntityAssociation {
	return v.value
}

func (v *NullableLegalEntityAssociation) Set(val *LegalEntityAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalEntityAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalEntityAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalEntityAssociation(val *LegalEntityAssociation) *NullableLegalEntityAssociation {
	return &NullableLegalEntityAssociation{value: val, isSet: true}
}

func (v NullableLegalEntityAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalEntityAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *LegalEntityAssociation) isValidType() bool {
	var allowedEnumValues = []string{"definedBeneficiary", "director", "immediateParentCompany", "legalRepresentative", "pciSignatory", "protector", "secondaryPartner", "secondaryTrustee", "settlor", "signatory", "soleProprietorship", "trust", "trustOwnership", "uboThroughControl", "uboThroughOwnership", "ultimateParentCompany", "undefinedBeneficiary", "unincorporatedPartnership"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
