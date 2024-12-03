/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the DefenseDocumentType type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DefenseDocumentType{}

// DefenseDocumentType struct for DefenseDocumentType
type DefenseDocumentType struct {
	// When **true**, you've successfully uploaded this type of defense document. When **false**, you haven't uploaded this defense document type.
	Available bool `json:"available"`
	// The document type code of the defense document.
	DefenseDocumentTypeCode string `json:"defenseDocumentTypeCode"`
	// Indicates to what extent the defense document is required in the defense process.  Possible values:   * **Required**: You must supply the document.   * **OneOrMore**: You must supply at least one of the documents with this label.  * **Optional**: You can choose to supply the document.  * **AlternativeRequired**: You must supply a generic defense document. To enable this functionality, contact our Support Team. When enabled, you can supply a generic defense document for all schemes.
	RequirementLevel string `json:"requirementLevel"`
}

// NewDefenseDocumentType instantiates a new DefenseDocumentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefenseDocumentType(available bool, defenseDocumentTypeCode string, requirementLevel string) *DefenseDocumentType {
	this := DefenseDocumentType{}
	this.Available = available
	this.DefenseDocumentTypeCode = defenseDocumentTypeCode
	this.RequirementLevel = requirementLevel
	return &this
}

// NewDefenseDocumentTypeWithDefaults instantiates a new DefenseDocumentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefenseDocumentTypeWithDefaults() *DefenseDocumentType {
	this := DefenseDocumentType{}
	return &this
}

// GetAvailable returns the Available field value
func (o *DefenseDocumentType) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *DefenseDocumentType) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *DefenseDocumentType) SetAvailable(v bool) {
	o.Available = v
}

// GetDefenseDocumentTypeCode returns the DefenseDocumentTypeCode field value
func (o *DefenseDocumentType) GetDefenseDocumentTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefenseDocumentTypeCode
}

// GetDefenseDocumentTypeCodeOk returns a tuple with the DefenseDocumentTypeCode field value
// and a boolean to check if the value has been set.
func (o *DefenseDocumentType) GetDefenseDocumentTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenseDocumentTypeCode, true
}

// SetDefenseDocumentTypeCode sets field value
func (o *DefenseDocumentType) SetDefenseDocumentTypeCode(v string) {
	o.DefenseDocumentTypeCode = v
}

// GetRequirementLevel returns the RequirementLevel field value
func (o *DefenseDocumentType) GetRequirementLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequirementLevel
}

// GetRequirementLevelOk returns a tuple with the RequirementLevel field value
// and a boolean to check if the value has been set.
func (o *DefenseDocumentType) GetRequirementLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequirementLevel, true
}

// SetRequirementLevel sets field value
func (o *DefenseDocumentType) SetRequirementLevel(v string) {
	o.RequirementLevel = v
}

func (o DefenseDocumentType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefenseDocumentType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["available"] = o.Available
	toSerialize["defenseDocumentTypeCode"] = o.DefenseDocumentTypeCode
	toSerialize["requirementLevel"] = o.RequirementLevel
	return toSerialize, nil
}

type NullableDefenseDocumentType struct {
	value *DefenseDocumentType
	isSet bool
}

func (v NullableDefenseDocumentType) Get() *DefenseDocumentType {
	return v.value
}

func (v *NullableDefenseDocumentType) Set(val *DefenseDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableDefenseDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableDefenseDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefenseDocumentType(val *DefenseDocumentType) *NullableDefenseDocumentType {
	return &NullableDefenseDocumentType{value: val, isSet: true}
}

func (v NullableDefenseDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefenseDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



