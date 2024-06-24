/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the CheckoutSessionThreeDS2RequestData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutSessionThreeDS2RequestData{}

// CheckoutSessionThreeDS2RequestData struct for CheckoutSessionThreeDS2RequestData
type CheckoutSessionThreeDS2RequestData struct {
	HomePhone   *Phone `json:"homePhone,omitempty"`
	MobilePhone *Phone `json:"mobilePhone,omitempty"`
	// Indicates whether a challenge is requested for this transaction. Possible values: * **01** — No preference * **02** — No challenge requested * **03** — Challenge requested (3DS Requestor preference) * **04** — Challenge requested (Mandate) * **05** — No challenge (transactional risk analysis is already performed) * **06** — Data Only
	ThreeDSRequestorChallengeInd *string `json:"threeDSRequestorChallengeInd,omitempty"`
	WorkPhone                    *Phone  `json:"workPhone,omitempty"`
}

// NewCheckoutSessionThreeDS2RequestData instantiates a new CheckoutSessionThreeDS2RequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionThreeDS2RequestData() *CheckoutSessionThreeDS2RequestData {
	this := CheckoutSessionThreeDS2RequestData{}
	return &this
}

// NewCheckoutSessionThreeDS2RequestDataWithDefaults instantiates a new CheckoutSessionThreeDS2RequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionThreeDS2RequestDataWithDefaults() *CheckoutSessionThreeDS2RequestData {
	this := CheckoutSessionThreeDS2RequestData{}
	return &this
}

// GetHomePhone returns the HomePhone field value if set, zero value otherwise.
func (o *CheckoutSessionThreeDS2RequestData) GetHomePhone() Phone {
	if o == nil || common.IsNil(o.HomePhone) {
		var ret Phone
		return ret
	}
	return *o.HomePhone
}

// GetHomePhoneOk returns a tuple with the HomePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionThreeDS2RequestData) GetHomePhoneOk() (*Phone, bool) {
	if o == nil || common.IsNil(o.HomePhone) {
		return nil, false
	}
	return o.HomePhone, true
}

// HasHomePhone returns a boolean if a field has been set.
func (o *CheckoutSessionThreeDS2RequestData) HasHomePhone() bool {
	if o != nil && !common.IsNil(o.HomePhone) {
		return true
	}

	return false
}

// SetHomePhone gets a reference to the given Phone and assigns it to the HomePhone field.
func (o *CheckoutSessionThreeDS2RequestData) SetHomePhone(v Phone) {
	o.HomePhone = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *CheckoutSessionThreeDS2RequestData) GetMobilePhone() Phone {
	if o == nil || common.IsNil(o.MobilePhone) {
		var ret Phone
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionThreeDS2RequestData) GetMobilePhoneOk() (*Phone, bool) {
	if o == nil || common.IsNil(o.MobilePhone) {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *CheckoutSessionThreeDS2RequestData) HasMobilePhone() bool {
	if o != nil && !common.IsNil(o.MobilePhone) {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given Phone and assigns it to the MobilePhone field.
func (o *CheckoutSessionThreeDS2RequestData) SetMobilePhone(v Phone) {
	o.MobilePhone = &v
}

// GetThreeDSRequestorChallengeInd returns the ThreeDSRequestorChallengeInd field value if set, zero value otherwise.
func (o *CheckoutSessionThreeDS2RequestData) GetThreeDSRequestorChallengeInd() string {
	if o == nil || common.IsNil(o.ThreeDSRequestorChallengeInd) {
		var ret string
		return ret
	}
	return *o.ThreeDSRequestorChallengeInd
}

// GetThreeDSRequestorChallengeIndOk returns a tuple with the ThreeDSRequestorChallengeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionThreeDS2RequestData) GetThreeDSRequestorChallengeIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSRequestorChallengeInd) {
		return nil, false
	}
	return o.ThreeDSRequestorChallengeInd, true
}

// HasThreeDSRequestorChallengeInd returns a boolean if a field has been set.
func (o *CheckoutSessionThreeDS2RequestData) HasThreeDSRequestorChallengeInd() bool {
	if o != nil && !common.IsNil(o.ThreeDSRequestorChallengeInd) {
		return true
	}

	return false
}

// SetThreeDSRequestorChallengeInd gets a reference to the given string and assigns it to the ThreeDSRequestorChallengeInd field.
func (o *CheckoutSessionThreeDS2RequestData) SetThreeDSRequestorChallengeInd(v string) {
	o.ThreeDSRequestorChallengeInd = &v
}

// GetWorkPhone returns the WorkPhone field value if set, zero value otherwise.
func (o *CheckoutSessionThreeDS2RequestData) GetWorkPhone() Phone {
	if o == nil || common.IsNil(o.WorkPhone) {
		var ret Phone
		return ret
	}
	return *o.WorkPhone
}

// GetWorkPhoneOk returns a tuple with the WorkPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionThreeDS2RequestData) GetWorkPhoneOk() (*Phone, bool) {
	if o == nil || common.IsNil(o.WorkPhone) {
		return nil, false
	}
	return o.WorkPhone, true
}

// HasWorkPhone returns a boolean if a field has been set.
func (o *CheckoutSessionThreeDS2RequestData) HasWorkPhone() bool {
	if o != nil && !common.IsNil(o.WorkPhone) {
		return true
	}

	return false
}

// SetWorkPhone gets a reference to the given Phone and assigns it to the WorkPhone field.
func (o *CheckoutSessionThreeDS2RequestData) SetWorkPhone(v Phone) {
	o.WorkPhone = &v
}

func (o CheckoutSessionThreeDS2RequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSessionThreeDS2RequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.HomePhone) {
		toSerialize["homePhone"] = o.HomePhone
	}
	if !common.IsNil(o.MobilePhone) {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if !common.IsNil(o.ThreeDSRequestorChallengeInd) {
		toSerialize["threeDSRequestorChallengeInd"] = o.ThreeDSRequestorChallengeInd
	}
	if !common.IsNil(o.WorkPhone) {
		toSerialize["workPhone"] = o.WorkPhone
	}
	return toSerialize, nil
}

type NullableCheckoutSessionThreeDS2RequestData struct {
	value *CheckoutSessionThreeDS2RequestData
	isSet bool
}

func (v NullableCheckoutSessionThreeDS2RequestData) Get() *CheckoutSessionThreeDS2RequestData {
	return v.value
}

func (v *NullableCheckoutSessionThreeDS2RequestData) Set(val *CheckoutSessionThreeDS2RequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionThreeDS2RequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionThreeDS2RequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionThreeDS2RequestData(val *CheckoutSessionThreeDS2RequestData) *NullableCheckoutSessionThreeDS2RequestData {
	return &NullableCheckoutSessionThreeDS2RequestData{value: val, isSet: true}
}

func (v NullableCheckoutSessionThreeDS2RequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionThreeDS2RequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutSessionThreeDS2RequestData) isValidThreeDSRequestorChallengeInd() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05", "06"}
	for _, allowed := range allowedEnumValues {
		if o.GetThreeDSRequestorChallengeInd() == allowed {
			return true
		}
	}
	return false
}
