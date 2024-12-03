/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the ModificationResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModificationResult{}

// ModificationResult struct for ModificationResult
type ModificationResult struct {
	// This field contains additional data, which may be returned in a particular modification response.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// Adyen's 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference string `json:"pspReference"`
	// Indicates if the modification request has been received for processing.
	Response string `json:"response"`
}

// NewModificationResult instantiates a new ModificationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModificationResult(pspReference string, response string) *ModificationResult {
	this := ModificationResult{}
	this.PspReference = pspReference
	this.Response = response
	return &this
}

// NewModificationResultWithDefaults instantiates a new ModificationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModificationResultWithDefaults() *ModificationResult {
	this := ModificationResult{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ModificationResult) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModificationResult) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ModificationResult) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ModificationResult) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetPspReference returns the PspReference field value
func (o *ModificationResult) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *ModificationResult) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *ModificationResult) SetPspReference(v string) {
	o.PspReference = v
}

// GetResponse returns the Response field value
func (o *ModificationResult) GetResponse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *ModificationResult) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *ModificationResult) SetResponse(v string) {
	o.Response = v
}

func (o ModificationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModificationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["pspReference"] = o.PspReference
	toSerialize["response"] = o.Response
	return toSerialize, nil
}

type NullableModificationResult struct {
	value *ModificationResult
	isSet bool
}

func (v NullableModificationResult) Get() *ModificationResult {
	return v.value
}

func (v *NullableModificationResult) Set(val *ModificationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModificationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModificationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModificationResult(val *ModificationResult) *NullableModificationResult {
	return &NullableModificationResult{value: val, isSet: true}
}

func (v NullableModificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModificationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *ModificationResult) isValidResponse() bool {
    var allowedEnumValues = []string{ "[capture-received]", "[cancel-received]", "[refund-received]", "[cancelOrRefund-received]", "[adjustAuthorisation-received]", "[donation-received]", "[technical-cancel-received]", "[voidPendingRefund-received]", "Authorised" }
    for _, allowed := range allowedEnumValues {
        if o.GetResponse() == allowed {
            return true
        }
    }
    return false
}

