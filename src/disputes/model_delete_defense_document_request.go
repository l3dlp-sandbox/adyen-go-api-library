/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the DeleteDefenseDocumentRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeleteDefenseDocumentRequest{}

// DeleteDefenseDocumentRequest struct for DeleteDefenseDocumentRequest
type DeleteDefenseDocumentRequest struct {
	// The document type code of the defense document.
	DefenseDocumentType string `json:"defenseDocumentType"`
	// The PSP reference assigned to the dispute.
	DisputePspReference string `json:"disputePspReference"`
	// The merchant account identifier, for which you want to process the dispute transaction.
	MerchantAccountCode string `json:"merchantAccountCode"`
}

// NewDeleteDefenseDocumentRequest instantiates a new DeleteDefenseDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDefenseDocumentRequest(defenseDocumentType string, disputePspReference string, merchantAccountCode string) *DeleteDefenseDocumentRequest {
	this := DeleteDefenseDocumentRequest{}
	this.DefenseDocumentType = defenseDocumentType
	this.DisputePspReference = disputePspReference
	this.MerchantAccountCode = merchantAccountCode
	return &this
}

// NewDeleteDefenseDocumentRequestWithDefaults instantiates a new DeleteDefenseDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDefenseDocumentRequestWithDefaults() *DeleteDefenseDocumentRequest {
	this := DeleteDefenseDocumentRequest{}
	return &this
}

// GetDefenseDocumentType returns the DefenseDocumentType field value
func (o *DeleteDefenseDocumentRequest) GetDefenseDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefenseDocumentType
}

// GetDefenseDocumentTypeOk returns a tuple with the DefenseDocumentType field value
// and a boolean to check if the value has been set.
func (o *DeleteDefenseDocumentRequest) GetDefenseDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefenseDocumentType, true
}

// SetDefenseDocumentType sets field value
func (o *DeleteDefenseDocumentRequest) SetDefenseDocumentType(v string) {
	o.DefenseDocumentType = v
}

// GetDisputePspReference returns the DisputePspReference field value
func (o *DeleteDefenseDocumentRequest) GetDisputePspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisputePspReference
}

// GetDisputePspReferenceOk returns a tuple with the DisputePspReference field value
// and a boolean to check if the value has been set.
func (o *DeleteDefenseDocumentRequest) GetDisputePspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputePspReference, true
}

// SetDisputePspReference sets field value
func (o *DeleteDefenseDocumentRequest) SetDisputePspReference(v string) {
	o.DisputePspReference = v
}

// GetMerchantAccountCode returns the MerchantAccountCode field value
func (o *DeleteDefenseDocumentRequest) GetMerchantAccountCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountCode
}

// GetMerchantAccountCodeOk returns a tuple with the MerchantAccountCode field value
// and a boolean to check if the value has been set.
func (o *DeleteDefenseDocumentRequest) GetMerchantAccountCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountCode, true
}

// SetMerchantAccountCode sets field value
func (o *DeleteDefenseDocumentRequest) SetMerchantAccountCode(v string) {
	o.MerchantAccountCode = v
}

func (o DeleteDefenseDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDefenseDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defenseDocumentType"] = o.DefenseDocumentType
	toSerialize["disputePspReference"] = o.DisputePspReference
	toSerialize["merchantAccountCode"] = o.MerchantAccountCode
	return toSerialize, nil
}

type NullableDeleteDefenseDocumentRequest struct {
	value *DeleteDefenseDocumentRequest
	isSet bool
}

func (v NullableDeleteDefenseDocumentRequest) Get() *DeleteDefenseDocumentRequest {
	return v.value
}

func (v *NullableDeleteDefenseDocumentRequest) Set(val *DeleteDefenseDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDefenseDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDefenseDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDefenseDocumentRequest(val *DeleteDefenseDocumentRequest) *NullableDeleteDefenseDocumentRequest {
	return &NullableDeleteDefenseDocumentRequest{value: val, isSet: true}
}

func (v NullableDeleteDefenseDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDefenseDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



