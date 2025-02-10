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

// checks if the LegalEntityCapability type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LegalEntityCapability{}

// LegalEntityCapability struct for LegalEntityCapability
type LegalEntityCapability struct {
	// Indicates whether the capability is allowed. Adyen sets this to **true** if the verification is successful.
	Allowed *bool `json:"allowed,omitempty"`
	// The capability level that is allowed for the legal entity.  Possible values: **notApplicable**, **low**, **medium**, **high**.
	AllowedLevel    *string             `json:"allowedLevel,omitempty"`
	AllowedSettings *CapabilitySettings `json:"allowedSettings,omitempty"`
	// Indicates whether the capability is requested. To check whether the legal entity is permitted to use the capability, refer to the `allowed` field.
	Requested *bool `json:"requested,omitempty"`
	// The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring.  Possible values: **notApplicable**, **low**, **medium**, **high**.
	RequestedLevel    *string             `json:"requestedLevel,omitempty"`
	RequestedSettings *CapabilitySettings `json:"requestedSettings,omitempty"`
	// The capability status of transfer instruments associated with the legal entity.
	TransferInstruments []SupportingEntityCapability `json:"transferInstruments,omitempty"`
	// The status of the verification checks for the capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the `errors` array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
	VerificationStatus *string `json:"verificationStatus,omitempty"`
}

// NewLegalEntityCapability instantiates a new LegalEntityCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalEntityCapability() *LegalEntityCapability {
	this := LegalEntityCapability{}
	return &this
}

// NewLegalEntityCapabilityWithDefaults instantiates a new LegalEntityCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalEntityCapabilityWithDefaults() *LegalEntityCapability {
	this := LegalEntityCapability{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetAllowed() bool {
	if o == nil || common.IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasAllowed() bool {
	if o != nil && !common.IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *LegalEntityCapability) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetAllowedLevel returns the AllowedLevel field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetAllowedLevel() string {
	if o == nil || common.IsNil(o.AllowedLevel) {
		var ret string
		return ret
	}
	return *o.AllowedLevel
}

// GetAllowedLevelOk returns a tuple with the AllowedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetAllowedLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.AllowedLevel) {
		return nil, false
	}
	return o.AllowedLevel, true
}

// HasAllowedLevel returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasAllowedLevel() bool {
	if o != nil && !common.IsNil(o.AllowedLevel) {
		return true
	}

	return false
}

// SetAllowedLevel gets a reference to the given string and assigns it to the AllowedLevel field.
func (o *LegalEntityCapability) SetAllowedLevel(v string) {
	o.AllowedLevel = &v
}

// GetAllowedSettings returns the AllowedSettings field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetAllowedSettings() CapabilitySettings {
	if o == nil || common.IsNil(o.AllowedSettings) {
		var ret CapabilitySettings
		return ret
	}
	return *o.AllowedSettings
}

// GetAllowedSettingsOk returns a tuple with the AllowedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetAllowedSettingsOk() (*CapabilitySettings, bool) {
	if o == nil || common.IsNil(o.AllowedSettings) {
		return nil, false
	}
	return o.AllowedSettings, true
}

// HasAllowedSettings returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasAllowedSettings() bool {
	if o != nil && !common.IsNil(o.AllowedSettings) {
		return true
	}

	return false
}

// SetAllowedSettings gets a reference to the given CapabilitySettings and assigns it to the AllowedSettings field.
func (o *LegalEntityCapability) SetAllowedSettings(v CapabilitySettings) {
	o.AllowedSettings = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetRequested() bool {
	if o == nil || common.IsNil(o.Requested) {
		var ret bool
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetRequestedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasRequested() bool {
	if o != nil && !common.IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given bool and assigns it to the Requested field.
func (o *LegalEntityCapability) SetRequested(v bool) {
	o.Requested = &v
}

// GetRequestedLevel returns the RequestedLevel field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetRequestedLevel() string {
	if o == nil || common.IsNil(o.RequestedLevel) {
		var ret string
		return ret
	}
	return *o.RequestedLevel
}

// GetRequestedLevelOk returns a tuple with the RequestedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetRequestedLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestedLevel) {
		return nil, false
	}
	return o.RequestedLevel, true
}

// HasRequestedLevel returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasRequestedLevel() bool {
	if o != nil && !common.IsNil(o.RequestedLevel) {
		return true
	}

	return false
}

// SetRequestedLevel gets a reference to the given string and assigns it to the RequestedLevel field.
func (o *LegalEntityCapability) SetRequestedLevel(v string) {
	o.RequestedLevel = &v
}

// GetRequestedSettings returns the RequestedSettings field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetRequestedSettings() CapabilitySettings {
	if o == nil || common.IsNil(o.RequestedSettings) {
		var ret CapabilitySettings
		return ret
	}
	return *o.RequestedSettings
}

// GetRequestedSettingsOk returns a tuple with the RequestedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetRequestedSettingsOk() (*CapabilitySettings, bool) {
	if o == nil || common.IsNil(o.RequestedSettings) {
		return nil, false
	}
	return o.RequestedSettings, true
}

// HasRequestedSettings returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasRequestedSettings() bool {
	if o != nil && !common.IsNil(o.RequestedSettings) {
		return true
	}

	return false
}

// SetRequestedSettings gets a reference to the given CapabilitySettings and assigns it to the RequestedSettings field.
func (o *LegalEntityCapability) SetRequestedSettings(v CapabilitySettings) {
	o.RequestedSettings = &v
}

// GetTransferInstruments returns the TransferInstruments field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetTransferInstruments() []SupportingEntityCapability {
	if o == nil || common.IsNil(o.TransferInstruments) {
		var ret []SupportingEntityCapability
		return ret
	}
	return o.TransferInstruments
}

// GetTransferInstrumentsOk returns a tuple with the TransferInstruments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetTransferInstrumentsOk() ([]SupportingEntityCapability, bool) {
	if o == nil || common.IsNil(o.TransferInstruments) {
		return nil, false
	}
	return o.TransferInstruments, true
}

// HasTransferInstruments returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasTransferInstruments() bool {
	if o != nil && !common.IsNil(o.TransferInstruments) {
		return true
	}

	return false
}

// SetTransferInstruments gets a reference to the given []SupportingEntityCapability and assigns it to the TransferInstruments field.
func (o *LegalEntityCapability) SetTransferInstruments(v []SupportingEntityCapability) {
	o.TransferInstruments = v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *LegalEntityCapability) GetVerificationStatus() string {
	if o == nil || common.IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntityCapability) GetVerificationStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *LegalEntityCapability) HasVerificationStatus() bool {
	if o != nil && !common.IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *LegalEntityCapability) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o LegalEntityCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalEntityCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !common.IsNil(o.AllowedLevel) {
		toSerialize["allowedLevel"] = o.AllowedLevel
	}
	if !common.IsNil(o.AllowedSettings) {
		toSerialize["allowedSettings"] = o.AllowedSettings
	}
	if !common.IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !common.IsNil(o.RequestedLevel) {
		toSerialize["requestedLevel"] = o.RequestedLevel
	}
	if !common.IsNil(o.RequestedSettings) {
		toSerialize["requestedSettings"] = o.RequestedSettings
	}
	if !common.IsNil(o.TransferInstruments) {
		toSerialize["transferInstruments"] = o.TransferInstruments
	}
	if !common.IsNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullableLegalEntityCapability struct {
	value *LegalEntityCapability
	isSet bool
}

func (v NullableLegalEntityCapability) Get() *LegalEntityCapability {
	return v.value
}

func (v *NullableLegalEntityCapability) Set(val *LegalEntityCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalEntityCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalEntityCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalEntityCapability(val *LegalEntityCapability) *NullableLegalEntityCapability {
	return &NullableLegalEntityCapability{value: val, isSet: true}
}

func (v NullableLegalEntityCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalEntityCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *LegalEntityCapability) isValidAllowedLevel() bool {
	var allowedEnumValues = []string{"high", "low", "medium", "notApplicable"}
	for _, allowed := range allowedEnumValues {
		if o.GetAllowedLevel() == allowed {
			return true
		}
	}
	return false
}
func (o *LegalEntityCapability) isValidRequestedLevel() bool {
	var allowedEnumValues = []string{"high", "low", "medium", "notApplicable"}
	for _, allowed := range allowedEnumValues {
		if o.GetRequestedLevel() == allowed {
			return true
		}
	}
	return false
}
