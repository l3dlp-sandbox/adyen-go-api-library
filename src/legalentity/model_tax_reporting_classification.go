/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the TaxReportingClassification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TaxReportingClassification{}

// TaxReportingClassification struct for TaxReportingClassification
type TaxReportingClassification struct {
	// The organization's business type.  Possible values: **other**, **listedPublicCompany**, **subsidiaryOfListedPublicCompany**, **governmentalOrganization**, **internationalOrganization**, **financialInstitution**.
	BusinessType *string `json:"businessType,omitempty"`
	// The Global Intermediary Identification Number (GIIN) required for FATCA. Only required if the organization is a US financial institution and the `businessType` is **financialInstitution**.
	FinancialInstitutionNumber *string `json:"financialInstitutionNumber,omitempty"`
	// The organization's main source of income. Only required if `businessType` is **other**.  Possible values: **businessOperation**, **realEstateSales**, **investmentInterestOrRoyalty**, **propertyRental**, **other**.
	MainSourceOfIncome *string `json:"mainSourceOfIncome,omitempty"`
	// The tax reporting classification type.  Possible values: **nonFinancialNonReportable**, **financialNonReportable**, **nonFinancialActive**, **nonFinancialPassive**.
	Type *string `json:"type,omitempty"`
}

// NewTaxReportingClassification instantiates a new TaxReportingClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxReportingClassification() *TaxReportingClassification {
	this := TaxReportingClassification{}
	return &this
}

// NewTaxReportingClassificationWithDefaults instantiates a new TaxReportingClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxReportingClassificationWithDefaults() *TaxReportingClassification {
	this := TaxReportingClassification{}
	return &this
}

// GetBusinessType returns the BusinessType field value if set, zero value otherwise.
func (o *TaxReportingClassification) GetBusinessType() string {
	if o == nil || common.IsNil(o.BusinessType) {
		var ret string
		return ret
	}
	return *o.BusinessType
}

// GetBusinessTypeOk returns a tuple with the BusinessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxReportingClassification) GetBusinessTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BusinessType) {
		return nil, false
	}
	return o.BusinessType, true
}

// HasBusinessType returns a boolean if a field has been set.
func (o *TaxReportingClassification) HasBusinessType() bool {
	if o != nil && !common.IsNil(o.BusinessType) {
		return true
	}

	return false
}

// SetBusinessType gets a reference to the given string and assigns it to the BusinessType field.
func (o *TaxReportingClassification) SetBusinessType(v string) {
	o.BusinessType = &v
}

// GetFinancialInstitutionNumber returns the FinancialInstitutionNumber field value if set, zero value otherwise.
func (o *TaxReportingClassification) GetFinancialInstitutionNumber() string {
	if o == nil || common.IsNil(o.FinancialInstitutionNumber) {
		var ret string
		return ret
	}
	return *o.FinancialInstitutionNumber
}

// GetFinancialInstitutionNumberOk returns a tuple with the FinancialInstitutionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxReportingClassification) GetFinancialInstitutionNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.FinancialInstitutionNumber) {
		return nil, false
	}
	return o.FinancialInstitutionNumber, true
}

// HasFinancialInstitutionNumber returns a boolean if a field has been set.
func (o *TaxReportingClassification) HasFinancialInstitutionNumber() bool {
	if o != nil && !common.IsNil(o.FinancialInstitutionNumber) {
		return true
	}

	return false
}

// SetFinancialInstitutionNumber gets a reference to the given string and assigns it to the FinancialInstitutionNumber field.
func (o *TaxReportingClassification) SetFinancialInstitutionNumber(v string) {
	o.FinancialInstitutionNumber = &v
}

// GetMainSourceOfIncome returns the MainSourceOfIncome field value if set, zero value otherwise.
func (o *TaxReportingClassification) GetMainSourceOfIncome() string {
	if o == nil || common.IsNil(o.MainSourceOfIncome) {
		var ret string
		return ret
	}
	return *o.MainSourceOfIncome
}

// GetMainSourceOfIncomeOk returns a tuple with the MainSourceOfIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxReportingClassification) GetMainSourceOfIncomeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MainSourceOfIncome) {
		return nil, false
	}
	return o.MainSourceOfIncome, true
}

// HasMainSourceOfIncome returns a boolean if a field has been set.
func (o *TaxReportingClassification) HasMainSourceOfIncome() bool {
	if o != nil && !common.IsNil(o.MainSourceOfIncome) {
		return true
	}

	return false
}

// SetMainSourceOfIncome gets a reference to the given string and assigns it to the MainSourceOfIncome field.
func (o *TaxReportingClassification) SetMainSourceOfIncome(v string) {
	o.MainSourceOfIncome = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaxReportingClassification) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxReportingClassification) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaxReportingClassification) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaxReportingClassification) SetType(v string) {
	o.Type = &v
}

func (o TaxReportingClassification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxReportingClassification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BusinessType) {
		toSerialize["businessType"] = o.BusinessType
	}
	if !common.IsNil(o.FinancialInstitutionNumber) {
		toSerialize["financialInstitutionNumber"] = o.FinancialInstitutionNumber
	}
	if !common.IsNil(o.MainSourceOfIncome) {
		toSerialize["mainSourceOfIncome"] = o.MainSourceOfIncome
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTaxReportingClassification struct {
	value *TaxReportingClassification
	isSet bool
}

func (v NullableTaxReportingClassification) Get() *TaxReportingClassification {
	return v.value
}

func (v *NullableTaxReportingClassification) Set(val *TaxReportingClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxReportingClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxReportingClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxReportingClassification(val *TaxReportingClassification) *NullableTaxReportingClassification {
	return &NullableTaxReportingClassification{value: val, isSet: true}
}

func (v NullableTaxReportingClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxReportingClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TaxReportingClassification) isValidBusinessType() bool {
	var allowedEnumValues = []string{"other", "listedPublicCompany", "subsidiaryOfListedPublicCompany", "governmentalOrganization", "internationalOrganization", "financialInstitution"}
	for _, allowed := range allowedEnumValues {
		if o.GetBusinessType() == allowed {
			return true
		}
	}
	return false
}
func (o *TaxReportingClassification) isValidMainSourceOfIncome() bool {
	var allowedEnumValues = []string{"businessOperation", "realEstateSales", "investmentInterestOrRoyalty", "propertyRental", "other"}
	for _, allowed := range allowedEnumValues {
		if o.GetMainSourceOfIncome() == allowed {
			return true
		}
	}
	return false
}
func (o *TaxReportingClassification) isValidType() bool {
	var allowedEnumValues = []string{"nonFinancialNonReportable", "financialNonReportable", "nonFinancialActive", "nonFinancialPassive"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
