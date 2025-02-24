/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the RoutingDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RoutingDetails{}

// RoutingDetails struct for RoutingDetails
type RoutingDetails struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// A code that identifies the problem type.
	ErrorCode *string `json:"errorCode,omitempty"`
	// The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Required for transfers with `category` **bank**.  Possible values:  * **regular**: for normal, low-value transactions.  * **fast**: a faster way to transfer funds, but the fees are higher. Recommended for high-priority, low-value transactions.  * **wire**: the fastest way to transfer funds, but this has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: for instant funds transfers in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: for high-value transfers to a recipient in a different country.  * **internal**: for transfers to an Adyen-issued business bank account (by bank account number/IBAN).
	Priority *string `json:"priority,omitempty"`
	// A short, human-readable summary of the problem type.
	Title *string `json:"title,omitempty"`
}

// NewRoutingDetails instantiates a new RoutingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingDetails() *RoutingDetails {
	this := RoutingDetails{}
	return &this
}

// NewRoutingDetailsWithDefaults instantiates a new RoutingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingDetailsWithDefaults() *RoutingDetails {
	this := RoutingDetails{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *RoutingDetails) GetDetail() string {
	if o == nil || common.IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingDetails) GetDetailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *RoutingDetails) HasDetail() bool {
	if o != nil && !common.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *RoutingDetails) SetDetail(v string) {
	o.Detail = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *RoutingDetails) GetErrorCode() string {
	if o == nil || common.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingDetails) GetErrorCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RoutingDetails) HasErrorCode() bool {
	if o != nil && !common.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *RoutingDetails) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RoutingDetails) GetPriority() string {
	if o == nil || common.IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingDetails) GetPriorityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RoutingDetails) HasPriority() bool {
	if o != nil && !common.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *RoutingDetails) SetPriority(v string) {
	o.Priority = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RoutingDetails) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingDetails) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RoutingDetails) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RoutingDetails) SetTitle(v string) {
	o.Title = &v
}

func (o RoutingDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !common.IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !common.IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !common.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableRoutingDetails struct {
	value *RoutingDetails
	isSet bool
}

func (v NullableRoutingDetails) Get() *RoutingDetails {
	return v.value
}

func (v *NullableRoutingDetails) Set(val *RoutingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingDetails(val *RoutingDetails) *NullableRoutingDetails {
	return &NullableRoutingDetails{value: val, isSet: true}
}

func (v NullableRoutingDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *RoutingDetails) isValidPriority() bool {
	var allowedEnumValues = []string{"crossBorder", "fast", "instant", "internal", "regular", "wire"}
	for _, allowed := range allowedEnumValues {
		if o.GetPriority() == allowed {
			return true
		}
	}
	return false
}
