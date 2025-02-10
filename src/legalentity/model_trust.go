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

// checks if the Trust type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Trust{}

// Trust struct for Trust
type Trust struct {
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the governing country.
	CountryOfGoverningLaw string `json:"countryOfGoverningLaw"`
	// The date when the legal arrangement was incorporated in YYYY-MM-DD format.
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	// A short description about the trust. Only applicable for charitable trusts in New Zealand.
	Description *string `json:"description,omitempty"`
	// The registered name, if different from the `name`.
	DoingBusinessAs *string `json:"doingBusinessAs,omitempty"`
	// The legal name.
	Name                     string   `json:"name"`
	PrincipalPlaceOfBusiness *Address `json:"principalPlaceOfBusiness,omitempty"`
	RegisteredAddress        Address  `json:"registeredAddress"`
	// The registration number.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// The tax information of the entity.
	TaxInformation []TaxInformation `json:"taxInformation,omitempty"`
	// Type of trust.  See possible values for trusts in [Australia](https://docs.adyen.com/platforms/verification-requirements/?tab=trust_3_4#trust-types-in-australia) and [New Zealand](https://docs.adyen.com/platforms/verification-requirements/?tab=trust_3_4#trust-types-in-new-zealand).
	Type string `json:"type"`
	// The undefined beneficiary information of the entity.
	UndefinedBeneficiaryInfo []UndefinedBeneficiary `json:"undefinedBeneficiaryInfo,omitempty"`
	// The reason for not providing a VAT number.  Possible values: **industryExemption**, **belowTaxThreshold**.
	VatAbsenceReason *string `json:"vatAbsenceReason,omitempty"`
	// The VAT number.
	VatNumber *string `json:"vatNumber,omitempty"`
}

// NewTrust instantiates a new Trust object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrust(countryOfGoverningLaw string, name string, registeredAddress Address, type_ string) *Trust {
	this := Trust{}
	this.CountryOfGoverningLaw = countryOfGoverningLaw
	this.Name = name
	this.RegisteredAddress = registeredAddress
	this.Type = type_
	return &this
}

// NewTrustWithDefaults instantiates a new Trust object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustWithDefaults() *Trust {
	this := Trust{}
	return &this
}

// GetCountryOfGoverningLaw returns the CountryOfGoverningLaw field value
func (o *Trust) GetCountryOfGoverningLaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryOfGoverningLaw
}

// GetCountryOfGoverningLawOk returns a tuple with the CountryOfGoverningLaw field value
// and a boolean to check if the value has been set.
func (o *Trust) GetCountryOfGoverningLawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryOfGoverningLaw, true
}

// SetCountryOfGoverningLaw sets field value
func (o *Trust) SetCountryOfGoverningLaw(v string) {
	o.CountryOfGoverningLaw = v
}

// GetDateOfIncorporation returns the DateOfIncorporation field value if set, zero value otherwise.
func (o *Trust) GetDateOfIncorporation() string {
	if o == nil || common.IsNil(o.DateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.DateOfIncorporation
}

// GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetDateOfIncorporationOk() (*string, bool) {
	if o == nil || common.IsNil(o.DateOfIncorporation) {
		return nil, false
	}
	return o.DateOfIncorporation, true
}

// HasDateOfIncorporation returns a boolean if a field has been set.
func (o *Trust) HasDateOfIncorporation() bool {
	if o != nil && !common.IsNil(o.DateOfIncorporation) {
		return true
	}

	return false
}

// SetDateOfIncorporation gets a reference to the given string and assigns it to the DateOfIncorporation field.
func (o *Trust) SetDateOfIncorporation(v string) {
	o.DateOfIncorporation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Trust) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Trust) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Trust) SetDescription(v string) {
	o.Description = &v
}

// GetDoingBusinessAs returns the DoingBusinessAs field value if set, zero value otherwise.
func (o *Trust) GetDoingBusinessAs() string {
	if o == nil || common.IsNil(o.DoingBusinessAs) {
		var ret string
		return ret
	}
	return *o.DoingBusinessAs
}

// GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetDoingBusinessAsOk() (*string, bool) {
	if o == nil || common.IsNil(o.DoingBusinessAs) {
		return nil, false
	}
	return o.DoingBusinessAs, true
}

// HasDoingBusinessAs returns a boolean if a field has been set.
func (o *Trust) HasDoingBusinessAs() bool {
	if o != nil && !common.IsNil(o.DoingBusinessAs) {
		return true
	}

	return false
}

// SetDoingBusinessAs gets a reference to the given string and assigns it to the DoingBusinessAs field.
func (o *Trust) SetDoingBusinessAs(v string) {
	o.DoingBusinessAs = &v
}

// GetName returns the Name field value
func (o *Trust) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Trust) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Trust) SetName(v string) {
	o.Name = v
}

// GetPrincipalPlaceOfBusiness returns the PrincipalPlaceOfBusiness field value if set, zero value otherwise.
func (o *Trust) GetPrincipalPlaceOfBusiness() Address {
	if o == nil || common.IsNil(o.PrincipalPlaceOfBusiness) {
		var ret Address
		return ret
	}
	return *o.PrincipalPlaceOfBusiness
}

// GetPrincipalPlaceOfBusinessOk returns a tuple with the PrincipalPlaceOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetPrincipalPlaceOfBusinessOk() (*Address, bool) {
	if o == nil || common.IsNil(o.PrincipalPlaceOfBusiness) {
		return nil, false
	}
	return o.PrincipalPlaceOfBusiness, true
}

// HasPrincipalPlaceOfBusiness returns a boolean if a field has been set.
func (o *Trust) HasPrincipalPlaceOfBusiness() bool {
	if o != nil && !common.IsNil(o.PrincipalPlaceOfBusiness) {
		return true
	}

	return false
}

// SetPrincipalPlaceOfBusiness gets a reference to the given Address and assigns it to the PrincipalPlaceOfBusiness field.
func (o *Trust) SetPrincipalPlaceOfBusiness(v Address) {
	o.PrincipalPlaceOfBusiness = &v
}

// GetRegisteredAddress returns the RegisteredAddress field value
func (o *Trust) GetRegisteredAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.RegisteredAddress
}

// GetRegisteredAddressOk returns a tuple with the RegisteredAddress field value
// and a boolean to check if the value has been set.
func (o *Trust) GetRegisteredAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegisteredAddress, true
}

// SetRegisteredAddress sets field value
func (o *Trust) SetRegisteredAddress(v Address) {
	o.RegisteredAddress = v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *Trust) GetRegistrationNumber() string {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *Trust) HasRegistrationNumber() bool {
	if o != nil && !common.IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *Trust) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetTaxInformation returns the TaxInformation field value if set, zero value otherwise.
func (o *Trust) GetTaxInformation() []TaxInformation {
	if o == nil || common.IsNil(o.TaxInformation) {
		var ret []TaxInformation
		return ret
	}
	return o.TaxInformation
}

// GetTaxInformationOk returns a tuple with the TaxInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetTaxInformationOk() ([]TaxInformation, bool) {
	if o == nil || common.IsNil(o.TaxInformation) {
		return nil, false
	}
	return o.TaxInformation, true
}

// HasTaxInformation returns a boolean if a field has been set.
func (o *Trust) HasTaxInformation() bool {
	if o != nil && !common.IsNil(o.TaxInformation) {
		return true
	}

	return false
}

// SetTaxInformation gets a reference to the given []TaxInformation and assigns it to the TaxInformation field.
func (o *Trust) SetTaxInformation(v []TaxInformation) {
	o.TaxInformation = v
}

// GetType returns the Type field value
func (o *Trust) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Trust) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Trust) SetType(v string) {
	o.Type = v
}

// GetUndefinedBeneficiaryInfo returns the UndefinedBeneficiaryInfo field value if set, zero value otherwise.
func (o *Trust) GetUndefinedBeneficiaryInfo() []UndefinedBeneficiary {
	if o == nil || common.IsNil(o.UndefinedBeneficiaryInfo) {
		var ret []UndefinedBeneficiary
		return ret
	}
	return o.UndefinedBeneficiaryInfo
}

// GetUndefinedBeneficiaryInfoOk returns a tuple with the UndefinedBeneficiaryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetUndefinedBeneficiaryInfoOk() ([]UndefinedBeneficiary, bool) {
	if o == nil || common.IsNil(o.UndefinedBeneficiaryInfo) {
		return nil, false
	}
	return o.UndefinedBeneficiaryInfo, true
}

// HasUndefinedBeneficiaryInfo returns a boolean if a field has been set.
func (o *Trust) HasUndefinedBeneficiaryInfo() bool {
	if o != nil && !common.IsNil(o.UndefinedBeneficiaryInfo) {
		return true
	}

	return false
}

// SetUndefinedBeneficiaryInfo gets a reference to the given []UndefinedBeneficiary and assigns it to the UndefinedBeneficiaryInfo field.
func (o *Trust) SetUndefinedBeneficiaryInfo(v []UndefinedBeneficiary) {
	o.UndefinedBeneficiaryInfo = v
}

// GetVatAbsenceReason returns the VatAbsenceReason field value if set, zero value otherwise.
func (o *Trust) GetVatAbsenceReason() string {
	if o == nil || common.IsNil(o.VatAbsenceReason) {
		var ret string
		return ret
	}
	return *o.VatAbsenceReason
}

// GetVatAbsenceReasonOk returns a tuple with the VatAbsenceReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetVatAbsenceReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.VatAbsenceReason) {
		return nil, false
	}
	return o.VatAbsenceReason, true
}

// HasVatAbsenceReason returns a boolean if a field has been set.
func (o *Trust) HasVatAbsenceReason() bool {
	if o != nil && !common.IsNil(o.VatAbsenceReason) {
		return true
	}

	return false
}

// SetVatAbsenceReason gets a reference to the given string and assigns it to the VatAbsenceReason field.
func (o *Trust) SetVatAbsenceReason(v string) {
	o.VatAbsenceReason = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *Trust) GetVatNumber() string {
	if o == nil || common.IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trust) GetVatNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Trust) HasVatNumber() bool {
	if o != nil && !common.IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *Trust) SetVatNumber(v string) {
	o.VatNumber = &v
}

func (o Trust) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trust) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryOfGoverningLaw"] = o.CountryOfGoverningLaw
	if !common.IsNil(o.DateOfIncorporation) {
		toSerialize["dateOfIncorporation"] = o.DateOfIncorporation
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.DoingBusinessAs) {
		toSerialize["doingBusinessAs"] = o.DoingBusinessAs
	}
	toSerialize["name"] = o.Name
	if !common.IsNil(o.PrincipalPlaceOfBusiness) {
		toSerialize["principalPlaceOfBusiness"] = o.PrincipalPlaceOfBusiness
	}
	toSerialize["registeredAddress"] = o.RegisteredAddress
	if !common.IsNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if !common.IsNil(o.TaxInformation) {
		toSerialize["taxInformation"] = o.TaxInformation
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.UndefinedBeneficiaryInfo) {
		toSerialize["undefinedBeneficiaryInfo"] = o.UndefinedBeneficiaryInfo
	}
	if !common.IsNil(o.VatAbsenceReason) {
		toSerialize["vatAbsenceReason"] = o.VatAbsenceReason
	}
	if !common.IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	return toSerialize, nil
}

type NullableTrust struct {
	value *Trust
	isSet bool
}

func (v NullableTrust) Get() *Trust {
	return v.value
}

func (v *NullableTrust) Set(val *Trust) {
	v.value = val
	v.isSet = true
}

func (v NullableTrust) IsSet() bool {
	return v.isSet
}

func (v *NullableTrust) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrust(val *Trust) *NullableTrust {
	return &NullableTrust{value: val, isSet: true}
}

func (v NullableTrust) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrust) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Trust) isValidType() bool {
	var allowedEnumValues = []string{"businessTrust", "cashManagementTrust", "charitableTrust", "corporateUnitTrust", "deceasedEstate", "discretionaryTrust", "discretionaryInvestmentTrust", "discretionaryServicesManagementTrust", "discretionaryTradingTrust", "familyTrust", "firstHomeSaverAccountsTrust", "fixedTrust", "fixedUnitTrust", "hybridTrust", "listedPublicUnitTrust", "otherTrust", "pooledSuperannuationTrust", "publicTradingTrust", "unlistedPublicUnitTrust"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
func (o *Trust) isValidVatAbsenceReason() bool {
	var allowedEnumValues = []string{"industryExemption", "belowTaxThreshold"}
	for _, allowed := range allowedEnumValues {
		if o.GetVatAbsenceReason() == allowed {
			return true
		}
	}
	return false
}
