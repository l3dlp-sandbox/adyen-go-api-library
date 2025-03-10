/*
Adyen Recurring API (deprecated)

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the RecurringDetailsResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecurringDetailsResult{}

// RecurringDetailsResult struct for RecurringDetailsResult
type RecurringDetailsResult struct {
	// The date when the recurring details were created.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Payment details stored for recurring payments.
	Details []RecurringDetailWrapper `json:"details,omitempty"`
	// The most recent email for this shopper (if available).
	LastKnownShopperEmail *string `json:"lastKnownShopperEmail,omitempty"`
	// The reference you use to uniquely identify the shopper (e.g. user ID or account ID).
	ShopperReference *string `json:"shopperReference,omitempty"`
}

// NewRecurringDetailsResult instantiates a new RecurringDetailsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringDetailsResult() *RecurringDetailsResult {
	this := RecurringDetailsResult{}
	return &this
}

// NewRecurringDetailsResultWithDefaults instantiates a new RecurringDetailsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringDetailsResultWithDefaults() *RecurringDetailsResult {
	this := RecurringDetailsResult{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *RecurringDetailsResult) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetailsResult) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *RecurringDetailsResult) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *RecurringDetailsResult) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RecurringDetailsResult) GetDetails() []RecurringDetailWrapper {
	if o == nil || common.IsNil(o.Details) {
		var ret []RecurringDetailWrapper
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetailsResult) GetDetailsOk() ([]RecurringDetailWrapper, bool) {
	if o == nil || common.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RecurringDetailsResult) HasDetails() bool {
	if o != nil && !common.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []RecurringDetailWrapper and assigns it to the Details field.
func (o *RecurringDetailsResult) SetDetails(v []RecurringDetailWrapper) {
	o.Details = v
}

// GetLastKnownShopperEmail returns the LastKnownShopperEmail field value if set, zero value otherwise.
func (o *RecurringDetailsResult) GetLastKnownShopperEmail() string {
	if o == nil || common.IsNil(o.LastKnownShopperEmail) {
		var ret string
		return ret
	}
	return *o.LastKnownShopperEmail
}

// GetLastKnownShopperEmailOk returns a tuple with the LastKnownShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetailsResult) GetLastKnownShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastKnownShopperEmail) {
		return nil, false
	}
	return o.LastKnownShopperEmail, true
}

// HasLastKnownShopperEmail returns a boolean if a field has been set.
func (o *RecurringDetailsResult) HasLastKnownShopperEmail() bool {
	if o != nil && !common.IsNil(o.LastKnownShopperEmail) {
		return true
	}

	return false
}

// SetLastKnownShopperEmail gets a reference to the given string and assigns it to the LastKnownShopperEmail field.
func (o *RecurringDetailsResult) SetLastKnownShopperEmail(v string) {
	o.LastKnownShopperEmail = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *RecurringDetailsResult) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetailsResult) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *RecurringDetailsResult) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *RecurringDetailsResult) SetShopperReference(v string) {
	o.ShopperReference = &v
}

func (o RecurringDetailsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurringDetailsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !common.IsNil(o.LastKnownShopperEmail) {
		toSerialize["lastKnownShopperEmail"] = o.LastKnownShopperEmail
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	return toSerialize, nil
}

type NullableRecurringDetailsResult struct {
	value *RecurringDetailsResult
	isSet bool
}

func (v NullableRecurringDetailsResult) Get() *RecurringDetailsResult {
	return v.value
}

func (v *NullableRecurringDetailsResult) Set(val *RecurringDetailsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringDetailsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringDetailsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringDetailsResult(val *RecurringDetailsResult) *NullableRecurringDetailsResult {
	return &NullableRecurringDetailsResult{value: val, isSet: true}
}

func (v NullableRecurringDetailsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringDetailsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
