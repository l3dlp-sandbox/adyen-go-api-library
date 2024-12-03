/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the IssuedCard type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IssuedCard{}

// IssuedCard struct for IssuedCard
type IssuedCard struct {
	// The authorisation type. For example, **defaultAuthorisation**, **preAuthorisation**, **finalAuthorisation**
	AuthorisationType *string `json:"authorisationType,omitempty"`
	// Indicates the method used for entering the PAN to initiate a transaction.  Possible values: **manual**, **chip**, **magstripe**, **contactless**, **cof**, **ecommerce**, **token**.
	PanEntryMode *string `json:"panEntryMode,omitempty"`
	// Contains information about how the payment was processed. For example, **ecommerce** for online or **pos** for in-person payments.
	ProcessingType *string `json:"processingType,omitempty"`
	RelayedAuthorisationData *RelayedAuthorisationData `json:"relayedAuthorisationData,omitempty"`
	// The identifier of the original payment. This ID is provided by the scheme and can be alphanumeric or numeric, depending on the scheme. The `schemeTraceID` should refer to an original `schemeUniqueTransactionID` provided in an earlier payment (not necessarily processed by Adyen). A `schemeTraceId` is typically available for authorization adjustments or recurring payments.
	SchemeTraceId *string `json:"schemeTraceId,omitempty"`
	// The unique identifier created by the scheme. This ID can be alphanumeric or numeric depending on the scheme.
	SchemeUniqueTransactionId *string `json:"schemeUniqueTransactionId,omitempty"`
	// **issuedCard**
	Type *string `json:"type,omitempty"`
	// The evaluation of the validation facts. See [validation checks](https://docs.adyen.com/issuing/validation-checks) for more information.
	ValidationFacts []TransferNotificationValidationFact `json:"validationFacts,omitempty"`
}

// NewIssuedCard instantiates a new IssuedCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedCard() *IssuedCard {
	this := IssuedCard{}
	var type_ string = "issuedCard"
	this.Type = &type_
	return &this
}

// NewIssuedCardWithDefaults instantiates a new IssuedCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedCardWithDefaults() *IssuedCard {
	this := IssuedCard{}
	var type_ string = "issuedCard"
	this.Type = &type_
	return &this
}

// GetAuthorisationType returns the AuthorisationType field value if set, zero value otherwise.
func (o *IssuedCard) GetAuthorisationType() string {
	if o == nil || common.IsNil(o.AuthorisationType) {
		var ret string
		return ret
	}
	return *o.AuthorisationType
}

// GetAuthorisationTypeOk returns a tuple with the AuthorisationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetAuthorisationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthorisationType) {
		return nil, false
	}
	return o.AuthorisationType, true
}

// HasAuthorisationType returns a boolean if a field has been set.
func (o *IssuedCard) HasAuthorisationType() bool {
	if o != nil && !common.IsNil(o.AuthorisationType) {
		return true
	}

	return false
}

// SetAuthorisationType gets a reference to the given string and assigns it to the AuthorisationType field.
func (o *IssuedCard) SetAuthorisationType(v string) {
	o.AuthorisationType = &v
}

// GetPanEntryMode returns the PanEntryMode field value if set, zero value otherwise.
func (o *IssuedCard) GetPanEntryMode() string {
	if o == nil || common.IsNil(o.PanEntryMode) {
		var ret string
		return ret
	}
	return *o.PanEntryMode
}

// GetPanEntryModeOk returns a tuple with the PanEntryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetPanEntryModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PanEntryMode) {
		return nil, false
	}
	return o.PanEntryMode, true
}

// HasPanEntryMode returns a boolean if a field has been set.
func (o *IssuedCard) HasPanEntryMode() bool {
	if o != nil && !common.IsNil(o.PanEntryMode) {
		return true
	}

	return false
}

// SetPanEntryMode gets a reference to the given string and assigns it to the PanEntryMode field.
func (o *IssuedCard) SetPanEntryMode(v string) {
	o.PanEntryMode = &v
}

// GetProcessingType returns the ProcessingType field value if set, zero value otherwise.
func (o *IssuedCard) GetProcessingType() string {
	if o == nil || common.IsNil(o.ProcessingType) {
		var ret string
		return ret
	}
	return *o.ProcessingType
}

// GetProcessingTypeOk returns a tuple with the ProcessingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetProcessingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProcessingType) {
		return nil, false
	}
	return o.ProcessingType, true
}

// HasProcessingType returns a boolean if a field has been set.
func (o *IssuedCard) HasProcessingType() bool {
	if o != nil && !common.IsNil(o.ProcessingType) {
		return true
	}

	return false
}

// SetProcessingType gets a reference to the given string and assigns it to the ProcessingType field.
func (o *IssuedCard) SetProcessingType(v string) {
	o.ProcessingType = &v
}

// GetRelayedAuthorisationData returns the RelayedAuthorisationData field value if set, zero value otherwise.
func (o *IssuedCard) GetRelayedAuthorisationData() RelayedAuthorisationData {
	if o == nil || common.IsNil(o.RelayedAuthorisationData) {
		var ret RelayedAuthorisationData
		return ret
	}
	return *o.RelayedAuthorisationData
}

// GetRelayedAuthorisationDataOk returns a tuple with the RelayedAuthorisationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetRelayedAuthorisationDataOk() (*RelayedAuthorisationData, bool) {
	if o == nil || common.IsNil(o.RelayedAuthorisationData) {
		return nil, false
	}
	return o.RelayedAuthorisationData, true
}

// HasRelayedAuthorisationData returns a boolean if a field has been set.
func (o *IssuedCard) HasRelayedAuthorisationData() bool {
	if o != nil && !common.IsNil(o.RelayedAuthorisationData) {
		return true
	}

	return false
}

// SetRelayedAuthorisationData gets a reference to the given RelayedAuthorisationData and assigns it to the RelayedAuthorisationData field.
func (o *IssuedCard) SetRelayedAuthorisationData(v RelayedAuthorisationData) {
	o.RelayedAuthorisationData = &v
}

// GetSchemeTraceId returns the SchemeTraceId field value if set, zero value otherwise.
func (o *IssuedCard) GetSchemeTraceId() string {
	if o == nil || common.IsNil(o.SchemeTraceId) {
		var ret string
		return ret
	}
	return *o.SchemeTraceId
}

// GetSchemeTraceIdOk returns a tuple with the SchemeTraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetSchemeTraceIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SchemeTraceId) {
		return nil, false
	}
	return o.SchemeTraceId, true
}

// HasSchemeTraceId returns a boolean if a field has been set.
func (o *IssuedCard) HasSchemeTraceId() bool {
	if o != nil && !common.IsNil(o.SchemeTraceId) {
		return true
	}

	return false
}

// SetSchemeTraceId gets a reference to the given string and assigns it to the SchemeTraceId field.
func (o *IssuedCard) SetSchemeTraceId(v string) {
	o.SchemeTraceId = &v
}

// GetSchemeUniqueTransactionId returns the SchemeUniqueTransactionId field value if set, zero value otherwise.
func (o *IssuedCard) GetSchemeUniqueTransactionId() string {
	if o == nil || common.IsNil(o.SchemeUniqueTransactionId) {
		var ret string
		return ret
	}
	return *o.SchemeUniqueTransactionId
}

// GetSchemeUniqueTransactionIdOk returns a tuple with the SchemeUniqueTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetSchemeUniqueTransactionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SchemeUniqueTransactionId) {
		return nil, false
	}
	return o.SchemeUniqueTransactionId, true
}

// HasSchemeUniqueTransactionId returns a boolean if a field has been set.
func (o *IssuedCard) HasSchemeUniqueTransactionId() bool {
	if o != nil && !common.IsNil(o.SchemeUniqueTransactionId) {
		return true
	}

	return false
}

// SetSchemeUniqueTransactionId gets a reference to the given string and assigns it to the SchemeUniqueTransactionId field.
func (o *IssuedCard) SetSchemeUniqueTransactionId(v string) {
	o.SchemeUniqueTransactionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssuedCard) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssuedCard) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssuedCard) SetType(v string) {
	o.Type = &v
}

// GetValidationFacts returns the ValidationFacts field value if set, zero value otherwise.
func (o *IssuedCard) GetValidationFacts() []TransferNotificationValidationFact {
	if o == nil || common.IsNil(o.ValidationFacts) {
		var ret []TransferNotificationValidationFact
		return ret
	}
	return o.ValidationFacts
}

// GetValidationFactsOk returns a tuple with the ValidationFacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCard) GetValidationFactsOk() ([]TransferNotificationValidationFact, bool) {
	if o == nil || common.IsNil(o.ValidationFacts) {
		return nil, false
	}
	return o.ValidationFacts, true
}

// HasValidationFacts returns a boolean if a field has been set.
func (o *IssuedCard) HasValidationFacts() bool {
	if o != nil && !common.IsNil(o.ValidationFacts) {
		return true
	}

	return false
}

// SetValidationFacts gets a reference to the given []TransferNotificationValidationFact and assigns it to the ValidationFacts field.
func (o *IssuedCard) SetValidationFacts(v []TransferNotificationValidationFact) {
	o.ValidationFacts = v
}

func (o IssuedCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthorisationType) {
		toSerialize["authorisationType"] = o.AuthorisationType
	}
	if !common.IsNil(o.PanEntryMode) {
		toSerialize["panEntryMode"] = o.PanEntryMode
	}
	if !common.IsNil(o.ProcessingType) {
		toSerialize["processingType"] = o.ProcessingType
	}
	if !common.IsNil(o.RelayedAuthorisationData) {
		toSerialize["relayedAuthorisationData"] = o.RelayedAuthorisationData
	}
	if !common.IsNil(o.SchemeTraceId) {
		toSerialize["schemeTraceId"] = o.SchemeTraceId
	}
	if !common.IsNil(o.SchemeUniqueTransactionId) {
		toSerialize["schemeUniqueTransactionId"] = o.SchemeUniqueTransactionId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.ValidationFacts) {
		toSerialize["validationFacts"] = o.ValidationFacts
	}
	return toSerialize, nil
}

type NullableIssuedCard struct {
	value *IssuedCard
	isSet bool
}

func (v NullableIssuedCard) Get() *IssuedCard {
	return v.value
}

func (v *NullableIssuedCard) Set(val *IssuedCard) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedCard) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedCard(val *IssuedCard) *NullableIssuedCard {
	return &NullableIssuedCard{value: val, isSet: true}
}

func (v NullableIssuedCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *IssuedCard) isValidPanEntryMode() bool {
    var allowedEnumValues = []string{ "chip", "cof", "contactless", "ecommerce", "magstripe", "manual", "token" }
    for _, allowed := range allowedEnumValues {
        if o.GetPanEntryMode() == allowed {
            return true
        }
    }
    return false
}
func (o *IssuedCard) isValidProcessingType() bool {
    var allowedEnumValues = []string{ "atmWithdraw", "balanceInquiry", "ecommerce", "moto", "pos", "purchaseWithCashback", "recurring", "token" }
    for _, allowed := range allowedEnumValues {
        if o.GetProcessingType() == allowed {
            return true
        }
    }
    return false
}
func (o *IssuedCard) isValidType() bool {
    var allowedEnumValues = []string{ "issuedCard" }
    for _, allowed := range allowedEnumValues {
        if o.GetType() == allowed {
            return true
        }
    }
    return false
}

