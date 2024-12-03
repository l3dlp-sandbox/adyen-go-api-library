/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the Permit type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Permit{}

// Permit struct for Permit
type Permit struct {
	// Partner ID (when using the permit-per-partner token sharing model).
	PartnerId *string `json:"partnerId,omitempty"`
	// The profile to apply to this permit (when using the shared permits model).
	ProfileReference *string            `json:"profileReference,omitempty"`
	Restriction      *PermitRestriction `json:"restriction,omitempty"`
	// The key to link permit requests to permit results.
	ResultKey *string `json:"resultKey,omitempty"`
	// The expiry date for this permit.
	ValidTillDate *time.Time `json:"validTillDate,omitempty"`
}

// NewPermit instantiates a new Permit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermit() *Permit {
	this := Permit{}
	return &this
}

// NewPermitWithDefaults instantiates a new Permit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermitWithDefaults() *Permit {
	this := Permit{}
	return &this
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *Permit) GetPartnerId() string {
	if o == nil || common.IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permit) GetPartnerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *Permit) HasPartnerId() bool {
	if o != nil && !common.IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *Permit) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetProfileReference returns the ProfileReference field value if set, zero value otherwise.
func (o *Permit) GetProfileReference() string {
	if o == nil || common.IsNil(o.ProfileReference) {
		var ret string
		return ret
	}
	return *o.ProfileReference
}

// GetProfileReferenceOk returns a tuple with the ProfileReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permit) GetProfileReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProfileReference) {
		return nil, false
	}
	return o.ProfileReference, true
}

// HasProfileReference returns a boolean if a field has been set.
func (o *Permit) HasProfileReference() bool {
	if o != nil && !common.IsNil(o.ProfileReference) {
		return true
	}

	return false
}

// SetProfileReference gets a reference to the given string and assigns it to the ProfileReference field.
func (o *Permit) SetProfileReference(v string) {
	o.ProfileReference = &v
}

// GetRestriction returns the Restriction field value if set, zero value otherwise.
func (o *Permit) GetRestriction() PermitRestriction {
	if o == nil || common.IsNil(o.Restriction) {
		var ret PermitRestriction
		return ret
	}
	return *o.Restriction
}

// GetRestrictionOk returns a tuple with the Restriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permit) GetRestrictionOk() (*PermitRestriction, bool) {
	if o == nil || common.IsNil(o.Restriction) {
		return nil, false
	}
	return o.Restriction, true
}

// HasRestriction returns a boolean if a field has been set.
func (o *Permit) HasRestriction() bool {
	if o != nil && !common.IsNil(o.Restriction) {
		return true
	}

	return false
}

// SetRestriction gets a reference to the given PermitRestriction and assigns it to the Restriction field.
func (o *Permit) SetRestriction(v PermitRestriction) {
	o.Restriction = &v
}

// GetResultKey returns the ResultKey field value if set, zero value otherwise.
func (o *Permit) GetResultKey() string {
	if o == nil || common.IsNil(o.ResultKey) {
		var ret string
		return ret
	}
	return *o.ResultKey
}

// GetResultKeyOk returns a tuple with the ResultKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permit) GetResultKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResultKey) {
		return nil, false
	}
	return o.ResultKey, true
}

// HasResultKey returns a boolean if a field has been set.
func (o *Permit) HasResultKey() bool {
	if o != nil && !common.IsNil(o.ResultKey) {
		return true
	}

	return false
}

// SetResultKey gets a reference to the given string and assigns it to the ResultKey field.
func (o *Permit) SetResultKey(v string) {
	o.ResultKey = &v
}

// GetValidTillDate returns the ValidTillDate field value if set, zero value otherwise.
func (o *Permit) GetValidTillDate() time.Time {
	if o == nil || common.IsNil(o.ValidTillDate) {
		var ret time.Time
		return ret
	}
	return *o.ValidTillDate
}

// GetValidTillDateOk returns a tuple with the ValidTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permit) GetValidTillDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ValidTillDate) {
		return nil, false
	}
	return o.ValidTillDate, true
}

// HasValidTillDate returns a boolean if a field has been set.
func (o *Permit) HasValidTillDate() bool {
	if o != nil && !common.IsNil(o.ValidTillDate) {
		return true
	}

	return false
}

// SetValidTillDate gets a reference to the given time.Time and assigns it to the ValidTillDate field.
func (o *Permit) SetValidTillDate(v time.Time) {
	o.ValidTillDate = &v
}

func (o Permit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Permit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PartnerId) {
		toSerialize["partnerId"] = o.PartnerId
	}
	if !common.IsNil(o.ProfileReference) {
		toSerialize["profileReference"] = o.ProfileReference
	}
	if !common.IsNil(o.Restriction) {
		toSerialize["restriction"] = o.Restriction
	}
	if !common.IsNil(o.ResultKey) {
		toSerialize["resultKey"] = o.ResultKey
	}
	if !common.IsNil(o.ValidTillDate) {
		toSerialize["validTillDate"] = o.ValidTillDate
	}
	return toSerialize, nil
}

type NullablePermit struct {
	value *Permit
	isSet bool
}

func (v NullablePermit) Get() *Permit {
	return v.value
}

func (v *NullablePermit) Set(val *Permit) {
	v.value = val
	v.isSet = true
}

func (v NullablePermit) IsSet() bool {
	return v.isSet
}

func (v *NullablePermit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermit(val *Permit) *NullablePermit {
	return &NullablePermit{value: val, isSet: true}
}

func (v NullablePermit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
