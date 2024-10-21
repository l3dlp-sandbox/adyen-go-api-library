/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the ThreeDSecureData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDSecureData{}

// ThreeDSecureData struct for ThreeDSecureData
type ThreeDSecureData struct {
	// In 3D Secure 1, the authentication response if the shopper was redirected.  In 3D Secure 2, this is the `transStatus` from the challenge result. If the transaction was frictionless, omit this parameter.
	AuthenticationResponse *string `json:"authenticationResponse,omitempty"`
	// The cardholder authentication value (base64 encoded, 20 bytes in a decoded form).
	Cavv *string `json:"cavv,omitempty"`
	// The CAVV algorithm used. Include this only for 3D Secure 1.
	CavvAlgorithm *string `json:"cavvAlgorithm,omitempty"`
	// Indicator informing the Access Control Server (ACS) and the Directory Server (DS) that the authentication has been cancelled. For possible values, refer to [3D Secure API reference](https://docs.adyen.com/online-payments/3d-secure/api-reference#mpidata).
	ChallengeCancel *string `json:"challengeCancel,omitempty"`
	// In 3D Secure 1, this is the enrollment response from the 3D directory server.  In 3D Secure 2, this is the `transStatus` from the `ARes`.
	DirectoryResponse *string `json:"directoryResponse,omitempty"`
	// Supported for 3D Secure 2. The unique transaction identifier assigned by the Directory Server (DS) to identify a single transaction.
	DsTransID *string `json:"dsTransID,omitempty"`
	// The electronic commerce indicator.
	Eci *string `json:"eci,omitempty"`
	// Risk score calculated by Directory Server (DS). Required for Cartes Bancaires integrations.
	RiskScore *string `json:"riskScore,omitempty"`
	// The version of the 3D Secure protocol.
	ThreeDSVersion *string `json:"threeDSVersion,omitempty"`
	// Network token authentication verification value (TAVV). The network token cryptogram.
	TokenAuthenticationVerificationValue *string `json:"tokenAuthenticationVerificationValue,omitempty"`
	// Provides information on why the `transStatus` field has the specified value. For possible values, refer to [our docs](https://docs.adyen.com/online-payments/3d-secure/api-reference#possible-transstatusreason-values).
	TransStatusReason *string `json:"transStatusReason,omitempty"`
	// Supported for 3D Secure 1. The transaction identifier (Base64-encoded, 20 bytes in a decoded form).
	Xid *string `json:"xid,omitempty"`
}

// NewThreeDSecureData instantiates a new ThreeDSecureData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSecureData() *ThreeDSecureData {
	this := ThreeDSecureData{}
	return &this
}

// NewThreeDSecureDataWithDefaults instantiates a new ThreeDSecureData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSecureDataWithDefaults() *ThreeDSecureData {
	this := ThreeDSecureData{}
	return &this
}

// GetAuthenticationResponse returns the AuthenticationResponse field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetAuthenticationResponse() string {
	if o == nil || common.IsNil(o.AuthenticationResponse) {
		var ret string
		return ret
	}
	return *o.AuthenticationResponse
}

// GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetAuthenticationResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthenticationResponse) {
		return nil, false
	}
	return o.AuthenticationResponse, true
}

// HasAuthenticationResponse returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasAuthenticationResponse() bool {
	if o != nil && !common.IsNil(o.AuthenticationResponse) {
		return true
	}

	return false
}

// SetAuthenticationResponse gets a reference to the given string and assigns it to the AuthenticationResponse field.
func (o *ThreeDSecureData) SetAuthenticationResponse(v string) {
	o.AuthenticationResponse = &v
}

// GetCavv returns the Cavv field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetCavv() string {
	if o == nil || common.IsNil(o.Cavv) {
		var ret string
		return ret
	}
	return *o.Cavv
}

// GetCavvOk returns a tuple with the Cavv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetCavvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cavv) {
		return nil, false
	}
	return o.Cavv, true
}

// HasCavv returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasCavv() bool {
	if o != nil && !common.IsNil(o.Cavv) {
		return true
	}

	return false
}

// SetCavv gets a reference to the given string and assigns it to the Cavv field.
func (o *ThreeDSecureData) SetCavv(v string) {
	o.Cavv = &v
}

// GetCavvAlgorithm returns the CavvAlgorithm field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetCavvAlgorithm() string {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		var ret string
		return ret
	}
	return *o.CavvAlgorithm
}

// GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetCavvAlgorithmOk() (*string, bool) {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		return nil, false
	}
	return o.CavvAlgorithm, true
}

// HasCavvAlgorithm returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasCavvAlgorithm() bool {
	if o != nil && !common.IsNil(o.CavvAlgorithm) {
		return true
	}

	return false
}

// SetCavvAlgorithm gets a reference to the given string and assigns it to the CavvAlgorithm field.
func (o *ThreeDSecureData) SetCavvAlgorithm(v string) {
	o.CavvAlgorithm = &v
}

// GetChallengeCancel returns the ChallengeCancel field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetChallengeCancel() string {
	if o == nil || common.IsNil(o.ChallengeCancel) {
		var ret string
		return ret
	}
	return *o.ChallengeCancel
}

// GetChallengeCancelOk returns a tuple with the ChallengeCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetChallengeCancelOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChallengeCancel) {
		return nil, false
	}
	return o.ChallengeCancel, true
}

// HasChallengeCancel returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasChallengeCancel() bool {
	if o != nil && !common.IsNil(o.ChallengeCancel) {
		return true
	}

	return false
}

// SetChallengeCancel gets a reference to the given string and assigns it to the ChallengeCancel field.
func (o *ThreeDSecureData) SetChallengeCancel(v string) {
	o.ChallengeCancel = &v
}

// GetDirectoryResponse returns the DirectoryResponse field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetDirectoryResponse() string {
	if o == nil || common.IsNil(o.DirectoryResponse) {
		var ret string
		return ret
	}
	return *o.DirectoryResponse
}

// GetDirectoryResponseOk returns a tuple with the DirectoryResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetDirectoryResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.DirectoryResponse) {
		return nil, false
	}
	return o.DirectoryResponse, true
}

// HasDirectoryResponse returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasDirectoryResponse() bool {
	if o != nil && !common.IsNil(o.DirectoryResponse) {
		return true
	}

	return false
}

// SetDirectoryResponse gets a reference to the given string and assigns it to the DirectoryResponse field.
func (o *ThreeDSecureData) SetDirectoryResponse(v string) {
	o.DirectoryResponse = &v
}

// GetDsTransID returns the DsTransID field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetDsTransID() string {
	if o == nil || common.IsNil(o.DsTransID) {
		var ret string
		return ret
	}
	return *o.DsTransID
}

// GetDsTransIDOk returns a tuple with the DsTransID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetDsTransIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.DsTransID) {
		return nil, false
	}
	return o.DsTransID, true
}

// HasDsTransID returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasDsTransID() bool {
	if o != nil && !common.IsNil(o.DsTransID) {
		return true
	}

	return false
}

// SetDsTransID gets a reference to the given string and assigns it to the DsTransID field.
func (o *ThreeDSecureData) SetDsTransID(v string) {
	o.DsTransID = &v
}

// GetEci returns the Eci field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetEci() string {
	if o == nil || common.IsNil(o.Eci) {
		var ret string
		return ret
	}
	return *o.Eci
}

// GetEciOk returns a tuple with the Eci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetEciOk() (*string, bool) {
	if o == nil || common.IsNil(o.Eci) {
		return nil, false
	}
	return o.Eci, true
}

// HasEci returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasEci() bool {
	if o != nil && !common.IsNil(o.Eci) {
		return true
	}

	return false
}

// SetEci gets a reference to the given string and assigns it to the Eci field.
func (o *ThreeDSecureData) SetEci(v string) {
	o.Eci = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetRiskScore() string {
	if o == nil || common.IsNil(o.RiskScore) {
		var ret string
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetRiskScoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskScore) {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasRiskScore() bool {
	if o != nil && !common.IsNil(o.RiskScore) {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given string and assigns it to the RiskScore field.
func (o *ThreeDSecureData) SetRiskScore(v string) {
	o.RiskScore = &v
}

// GetThreeDSVersion returns the ThreeDSVersion field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetThreeDSVersion() string {
	if o == nil || common.IsNil(o.ThreeDSVersion) {
		var ret string
		return ret
	}
	return *o.ThreeDSVersion
}

// GetThreeDSVersionOk returns a tuple with the ThreeDSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetThreeDSVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSVersion) {
		return nil, false
	}
	return o.ThreeDSVersion, true
}

// HasThreeDSVersion returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasThreeDSVersion() bool {
	if o != nil && !common.IsNil(o.ThreeDSVersion) {
		return true
	}

	return false
}

// SetThreeDSVersion gets a reference to the given string and assigns it to the ThreeDSVersion field.
func (o *ThreeDSecureData) SetThreeDSVersion(v string) {
	o.ThreeDSVersion = &v
}

// GetTokenAuthenticationVerificationValue returns the TokenAuthenticationVerificationValue field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetTokenAuthenticationVerificationValue() string {
	if o == nil || common.IsNil(o.TokenAuthenticationVerificationValue) {
		var ret string
		return ret
	}
	return *o.TokenAuthenticationVerificationValue
}

// GetTokenAuthenticationVerificationValueOk returns a tuple with the TokenAuthenticationVerificationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetTokenAuthenticationVerificationValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenAuthenticationVerificationValue) {
		return nil, false
	}
	return o.TokenAuthenticationVerificationValue, true
}

// HasTokenAuthenticationVerificationValue returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasTokenAuthenticationVerificationValue() bool {
	if o != nil && !common.IsNil(o.TokenAuthenticationVerificationValue) {
		return true
	}

	return false
}

// SetTokenAuthenticationVerificationValue gets a reference to the given string and assigns it to the TokenAuthenticationVerificationValue field.
func (o *ThreeDSecureData) SetTokenAuthenticationVerificationValue(v string) {
	o.TokenAuthenticationVerificationValue = &v
}

// GetTransStatusReason returns the TransStatusReason field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetTransStatusReason() string {
	if o == nil || common.IsNil(o.TransStatusReason) {
		var ret string
		return ret
	}
	return *o.TransStatusReason
}

// GetTransStatusReasonOk returns a tuple with the TransStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetTransStatusReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransStatusReason) {
		return nil, false
	}
	return o.TransStatusReason, true
}

// HasTransStatusReason returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasTransStatusReason() bool {
	if o != nil && !common.IsNil(o.TransStatusReason) {
		return true
	}

	return false
}

// SetTransStatusReason gets a reference to the given string and assigns it to the TransStatusReason field.
func (o *ThreeDSecureData) SetTransStatusReason(v string) {
	o.TransStatusReason = &v
}

// GetXid returns the Xid field value if set, zero value otherwise.
func (o *ThreeDSecureData) GetXid() string {
	if o == nil || common.IsNil(o.Xid) {
		var ret string
		return ret
	}
	return *o.Xid
}

// GetXidOk returns a tuple with the Xid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureData) GetXidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Xid) {
		return nil, false
	}
	return o.Xid, true
}

// HasXid returns a boolean if a field has been set.
func (o *ThreeDSecureData) HasXid() bool {
	if o != nil && !common.IsNil(o.Xid) {
		return true
	}

	return false
}

// SetXid gets a reference to the given string and assigns it to the Xid field.
func (o *ThreeDSecureData) SetXid(v string) {
	o.Xid = &v
}

func (o ThreeDSecureData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSecureData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthenticationResponse) {
		toSerialize["authenticationResponse"] = o.AuthenticationResponse
	}
	if !common.IsNil(o.Cavv) {
		toSerialize["cavv"] = o.Cavv
	}
	if !common.IsNil(o.CavvAlgorithm) {
		toSerialize["cavvAlgorithm"] = o.CavvAlgorithm
	}
	if !common.IsNil(o.ChallengeCancel) {
		toSerialize["challengeCancel"] = o.ChallengeCancel
	}
	if !common.IsNil(o.DirectoryResponse) {
		toSerialize["directoryResponse"] = o.DirectoryResponse
	}
	if !common.IsNil(o.DsTransID) {
		toSerialize["dsTransID"] = o.DsTransID
	}
	if !common.IsNil(o.Eci) {
		toSerialize["eci"] = o.Eci
	}
	if !common.IsNil(o.RiskScore) {
		toSerialize["riskScore"] = o.RiskScore
	}
	if !common.IsNil(o.ThreeDSVersion) {
		toSerialize["threeDSVersion"] = o.ThreeDSVersion
	}
	if !common.IsNil(o.TokenAuthenticationVerificationValue) {
		toSerialize["tokenAuthenticationVerificationValue"] = o.TokenAuthenticationVerificationValue
	}
	if !common.IsNil(o.TransStatusReason) {
		toSerialize["transStatusReason"] = o.TransStatusReason
	}
	if !common.IsNil(o.Xid) {
		toSerialize["xid"] = o.Xid
	}
	return toSerialize, nil
}

type NullableThreeDSecureData struct {
	value *ThreeDSecureData
	isSet bool
}

func (v NullableThreeDSecureData) Get() *ThreeDSecureData {
	return v.value
}

func (v *NullableThreeDSecureData) Set(val *ThreeDSecureData) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSecureData) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSecureData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSecureData(val *ThreeDSecureData) *NullableThreeDSecureData {
	return &NullableThreeDSecureData{value: val, isSet: true}
}

func (v NullableThreeDSecureData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSecureData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ThreeDSecureData) isValidAuthenticationResponse() bool {
	var allowedEnumValues = []string{"Y", "N", "U", "A"}
	for _, allowed := range allowedEnumValues {
		if o.GetAuthenticationResponse() == allowed {
			return true
		}
	}
	return false
}
func (o *ThreeDSecureData) isValidChallengeCancel() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05", "06", "07"}
	for _, allowed := range allowedEnumValues {
		if o.GetChallengeCancel() == allowed {
			return true
		}
	}
	return false
}
func (o *ThreeDSecureData) isValidDirectoryResponse() bool {
	var allowedEnumValues = []string{"A", "C", "D", "I", "N", "R", "U", "Y"}
	for _, allowed := range allowedEnumValues {
		if o.GetDirectoryResponse() == allowed {
			return true
		}
	}
	return false
}
