/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the VerificationError type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VerificationError{}

// VerificationError struct for VerificationError
type VerificationError struct {
	// Contains the capabilities that the verification error applies to.
	Capabilities []string `json:"capabilities,omitempty"`
	// The verification error code.
	Code *string `json:"code,omitempty"`
	// A description of the error.
	Message *string `json:"message,omitempty"`
	// Contains the actions that you can take to resolve the verification error.
	RemediatingActions []RemediatingAction `json:"remediatingActions,omitempty"`
	// Contains more granular information about the verification error.
	SubErrors []VerificationErrorRecursive `json:"subErrors,omitempty"`
	// The type of error.   Possible values: **invalidInput**, **dataMissing**.
	Type *string `json:"type,omitempty"`
}

// NewVerificationError instantiates a new VerificationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationError() *VerificationError {
	this := VerificationError{}
	return &this
}

// NewVerificationErrorWithDefaults instantiates a new VerificationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationErrorWithDefaults() *VerificationError {
	this := VerificationError{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *VerificationError) GetCapabilities() []string {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationError) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *VerificationError) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *VerificationError) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VerificationError) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationError) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VerificationError) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VerificationError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VerificationError) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationError) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VerificationError) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *VerificationError) SetMessage(v string) {
	o.Message = &v
}

// GetRemediatingActions returns the RemediatingActions field value if set, zero value otherwise.
func (o *VerificationError) GetRemediatingActions() []RemediatingAction {
	if o == nil || common.IsNil(o.RemediatingActions) {
		var ret []RemediatingAction
		return ret
	}
	return o.RemediatingActions
}

// GetRemediatingActionsOk returns a tuple with the RemediatingActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationError) GetRemediatingActionsOk() ([]RemediatingAction, bool) {
	if o == nil || common.IsNil(o.RemediatingActions) {
		return nil, false
	}
	return o.RemediatingActions, true
}

// HasRemediatingActions returns a boolean if a field has been set.
func (o *VerificationError) HasRemediatingActions() bool {
	if o != nil && !common.IsNil(o.RemediatingActions) {
		return true
	}

	return false
}

// SetRemediatingActions gets a reference to the given []RemediatingAction and assigns it to the RemediatingActions field.
func (o *VerificationError) SetRemediatingActions(v []RemediatingAction) {
	o.RemediatingActions = v
}

// GetSubErrors returns the SubErrors field value if set, zero value otherwise.
func (o *VerificationError) GetSubErrors() []VerificationErrorRecursive {
	if o == nil || common.IsNil(o.SubErrors) {
		var ret []VerificationErrorRecursive
		return ret
	}
	return o.SubErrors
}

// GetSubErrorsOk returns a tuple with the SubErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationError) GetSubErrorsOk() ([]VerificationErrorRecursive, bool) {
	if o == nil || common.IsNil(o.SubErrors) {
		return nil, false
	}
	return o.SubErrors, true
}

// HasSubErrors returns a boolean if a field has been set.
func (o *VerificationError) HasSubErrors() bool {
	if o != nil && !common.IsNil(o.SubErrors) {
		return true
	}

	return false
}

// SetSubErrors gets a reference to the given []VerificationErrorRecursive and assigns it to the SubErrors field.
func (o *VerificationError) SetSubErrors(v []VerificationErrorRecursive) {
	o.SubErrors = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VerificationError) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationError) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VerificationError) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VerificationError) SetType(v string) {
	o.Type = &v
}

func (o VerificationError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !common.IsNil(o.RemediatingActions) {
		toSerialize["remediatingActions"] = o.RemediatingActions
	}
	if !common.IsNil(o.SubErrors) {
		toSerialize["subErrors"] = o.SubErrors
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableVerificationError struct {
	value *VerificationError
	isSet bool
}

func (v NullableVerificationError) Get() *VerificationError {
	return v.value
}

func (v *NullableVerificationError) Set(val *VerificationError) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationError) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationError(val *VerificationError) *NullableVerificationError {
	return &NullableVerificationError{value: val, isSet: true}
}

func (v NullableVerificationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *VerificationError) isValidType() bool {
    var allowedEnumValues = []string{ "dataMissing", "invalidInput", "pendingStatus" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

