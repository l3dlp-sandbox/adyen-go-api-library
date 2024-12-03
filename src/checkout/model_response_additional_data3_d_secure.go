/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the ResponseAdditionalData3DSecure type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalData3DSecure{}

// ResponseAdditionalData3DSecure struct for ResponseAdditionalData3DSecure
type ResponseAdditionalData3DSecure struct {
	// Information provided by the issuer to the cardholder. If this field is present, you need to display this information to the cardholder. 
	CardHolderInfo *string `json:"cardHolderInfo,omitempty"`
	// The Cardholder Authentication Verification Value (CAVV) for the 3D Secure authentication session, as a Base64-encoded 20-byte array.
	Cavv *string `json:"cavv,omitempty"`
	// The CAVV algorithm used.
	CavvAlgorithm *string `json:"cavvAlgorithm,omitempty"`
	// Shows the [exemption type](https://docs.adyen.com/payments-fundamentals/psd2-sca-compliance-and-implementation-guide#specifypreferenceinyourapirequest) that Adyen requested for the payment.   Possible values: * **lowValue**  * **secureCorporate**  * **trustedBeneficiary**  * **transactionRiskAnalysis** 
	ScaExemptionRequested *string `json:"scaExemptionRequested,omitempty"`
	// Indicates whether a card is enrolled for 3D Secure 2.
	Threeds2CardEnrolled *bool `json:"threeds2.cardEnrolled,omitempty"`
}

// NewResponseAdditionalData3DSecure instantiates a new ResponseAdditionalData3DSecure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalData3DSecure() *ResponseAdditionalData3DSecure {
	this := ResponseAdditionalData3DSecure{}
	return &this
}

// NewResponseAdditionalData3DSecureWithDefaults instantiates a new ResponseAdditionalData3DSecure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalData3DSecureWithDefaults() *ResponseAdditionalData3DSecure {
	this := ResponseAdditionalData3DSecure{}
	return &this
}

// GetCardHolderInfo returns the CardHolderInfo field value if set, zero value otherwise.
func (o *ResponseAdditionalData3DSecure) GetCardHolderInfo() string {
	if o == nil || common.IsNil(o.CardHolderInfo) {
		var ret string
		return ret
	}
	return *o.CardHolderInfo
}

// GetCardHolderInfoOk returns a tuple with the CardHolderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalData3DSecure) GetCardHolderInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardHolderInfo) {
		return nil, false
	}
	return o.CardHolderInfo, true
}

// HasCardHolderInfo returns a boolean if a field has been set.
func (o *ResponseAdditionalData3DSecure) HasCardHolderInfo() bool {
	if o != nil && !common.IsNil(o.CardHolderInfo) {
		return true
	}

	return false
}

// SetCardHolderInfo gets a reference to the given string and assigns it to the CardHolderInfo field.
func (o *ResponseAdditionalData3DSecure) SetCardHolderInfo(v string) {
	o.CardHolderInfo = &v
}

// GetCavv returns the Cavv field value if set, zero value otherwise.
func (o *ResponseAdditionalData3DSecure) GetCavv() string {
	if o == nil || common.IsNil(o.Cavv) {
		var ret string
		return ret
	}
	return *o.Cavv
}

// GetCavvOk returns a tuple with the Cavv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalData3DSecure) GetCavvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cavv) {
		return nil, false
	}
	return o.Cavv, true
}

// HasCavv returns a boolean if a field has been set.
func (o *ResponseAdditionalData3DSecure) HasCavv() bool {
	if o != nil && !common.IsNil(o.Cavv) {
		return true
	}

	return false
}

// SetCavv gets a reference to the given string and assigns it to the Cavv field.
func (o *ResponseAdditionalData3DSecure) SetCavv(v string) {
	o.Cavv = &v
}

// GetCavvAlgorithm returns the CavvAlgorithm field value if set, zero value otherwise.
func (o *ResponseAdditionalData3DSecure) GetCavvAlgorithm() string {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		var ret string
		return ret
	}
	return *o.CavvAlgorithm
}

// GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalData3DSecure) GetCavvAlgorithmOk() (*string, bool) {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		return nil, false
	}
	return o.CavvAlgorithm, true
}

// HasCavvAlgorithm returns a boolean if a field has been set.
func (o *ResponseAdditionalData3DSecure) HasCavvAlgorithm() bool {
	if o != nil && !common.IsNil(o.CavvAlgorithm) {
		return true
	}

	return false
}

// SetCavvAlgorithm gets a reference to the given string and assigns it to the CavvAlgorithm field.
func (o *ResponseAdditionalData3DSecure) SetCavvAlgorithm(v string) {
	o.CavvAlgorithm = &v
}

// GetScaExemptionRequested returns the ScaExemptionRequested field value if set, zero value otherwise.
func (o *ResponseAdditionalData3DSecure) GetScaExemptionRequested() string {
	if o == nil || common.IsNil(o.ScaExemptionRequested) {
		var ret string
		return ret
	}
	return *o.ScaExemptionRequested
}

// GetScaExemptionRequestedOk returns a tuple with the ScaExemptionRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalData3DSecure) GetScaExemptionRequestedOk() (*string, bool) {
	if o == nil || common.IsNil(o.ScaExemptionRequested) {
		return nil, false
	}
	return o.ScaExemptionRequested, true
}

// HasScaExemptionRequested returns a boolean if a field has been set.
func (o *ResponseAdditionalData3DSecure) HasScaExemptionRequested() bool {
	if o != nil && !common.IsNil(o.ScaExemptionRequested) {
		return true
	}

	return false
}

// SetScaExemptionRequested gets a reference to the given string and assigns it to the ScaExemptionRequested field.
func (o *ResponseAdditionalData3DSecure) SetScaExemptionRequested(v string) {
	o.ScaExemptionRequested = &v
}

// GetThreeds2CardEnrolled returns the Threeds2CardEnrolled field value if set, zero value otherwise.
func (o *ResponseAdditionalData3DSecure) GetThreeds2CardEnrolled() bool {
	if o == nil || common.IsNil(o.Threeds2CardEnrolled) {
		var ret bool
		return ret
	}
	return *o.Threeds2CardEnrolled
}

// GetThreeds2CardEnrolledOk returns a tuple with the Threeds2CardEnrolled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalData3DSecure) GetThreeds2CardEnrolledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Threeds2CardEnrolled) {
		return nil, false
	}
	return o.Threeds2CardEnrolled, true
}

// HasThreeds2CardEnrolled returns a boolean if a field has been set.
func (o *ResponseAdditionalData3DSecure) HasThreeds2CardEnrolled() bool {
	if o != nil && !common.IsNil(o.Threeds2CardEnrolled) {
		return true
	}

	return false
}

// SetThreeds2CardEnrolled gets a reference to the given bool and assigns it to the Threeds2CardEnrolled field.
func (o *ResponseAdditionalData3DSecure) SetThreeds2CardEnrolled(v bool) {
	o.Threeds2CardEnrolled = &v
}

func (o ResponseAdditionalData3DSecure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalData3DSecure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CardHolderInfo) {
		toSerialize["cardHolderInfo"] = o.CardHolderInfo
	}
	if !common.IsNil(o.Cavv) {
		toSerialize["cavv"] = o.Cavv
	}
	if !common.IsNil(o.CavvAlgorithm) {
		toSerialize["cavvAlgorithm"] = o.CavvAlgorithm
	}
	if !common.IsNil(o.ScaExemptionRequested) {
		toSerialize["scaExemptionRequested"] = o.ScaExemptionRequested
	}
	if !common.IsNil(o.Threeds2CardEnrolled) {
		toSerialize["threeds2.cardEnrolled"] = o.Threeds2CardEnrolled
	}
	return toSerialize, nil
}

type NullableResponseAdditionalData3DSecure struct {
	value *ResponseAdditionalData3DSecure
	isSet bool
}

func (v NullableResponseAdditionalData3DSecure) Get() *ResponseAdditionalData3DSecure {
	return v.value
}

func (v *NullableResponseAdditionalData3DSecure) Set(val *ResponseAdditionalData3DSecure) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalData3DSecure) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalData3DSecure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalData3DSecure(val *ResponseAdditionalData3DSecure) *NullableResponseAdditionalData3DSecure {
	return &NullableResponseAdditionalData3DSecure{value: val, isSet: true}
}

func (v NullableResponseAdditionalData3DSecure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalData3DSecure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



