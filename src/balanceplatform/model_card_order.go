/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the CardOrder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardOrder{}

// CardOrder struct for CardOrder
type CardOrder struct {
	// The date when the card order is created.
	BeginDate *time.Time `json:"beginDate,omitempty"`
	// The unique identifier of the card manufacturer profile.
	CardManufacturingProfileId *string `json:"cardManufacturingProfileId,omitempty"`
	// The date when the card order processing ends.
	ClosedDate *time.Time `json:"closedDate,omitempty"`
	// The date when you manually closed the card order.  Card orders are automatically closed by the end of the day it was created. If you manually closed it beforehand, the closing date is shown as the `endDate`.
	EndDate *time.Time `json:"endDate,omitempty"`
	// The unique identifier of the card order.
	Id *string `json:"id,omitempty"`
	// The date when the card order processing begins.
	LockDate *time.Time `json:"lockDate,omitempty"`
	// The service center.
	ServiceCenter *string `json:"serviceCenter,omitempty"`
	// The status of the card order.  Possible values: **Open**, **Closed**.
	Status *string `json:"status,omitempty"`
}

// NewCardOrder instantiates a new CardOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardOrder() *CardOrder {
	this := CardOrder{}
	return &this
}

// NewCardOrderWithDefaults instantiates a new CardOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardOrderWithDefaults() *CardOrder {
	this := CardOrder{}
	return &this
}

// GetBeginDate returns the BeginDate field value if set, zero value otherwise.
func (o *CardOrder) GetBeginDate() time.Time {
	if o == nil || common.IsNil(o.BeginDate) {
		var ret time.Time
		return ret
	}
	return *o.BeginDate
}

// GetBeginDateOk returns a tuple with the BeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetBeginDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.BeginDate) {
		return nil, false
	}
	return o.BeginDate, true
}

// HasBeginDate returns a boolean if a field has been set.
func (o *CardOrder) HasBeginDate() bool {
	if o != nil && !common.IsNil(o.BeginDate) {
		return true
	}

	return false
}

// SetBeginDate gets a reference to the given time.Time and assigns it to the BeginDate field.
func (o *CardOrder) SetBeginDate(v time.Time) {
	o.BeginDate = &v
}

// GetCardManufacturingProfileId returns the CardManufacturingProfileId field value if set, zero value otherwise.
func (o *CardOrder) GetCardManufacturingProfileId() string {
	if o == nil || common.IsNil(o.CardManufacturingProfileId) {
		var ret string
		return ret
	}
	return *o.CardManufacturingProfileId
}

// GetCardManufacturingProfileIdOk returns a tuple with the CardManufacturingProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetCardManufacturingProfileIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardManufacturingProfileId) {
		return nil, false
	}
	return o.CardManufacturingProfileId, true
}

// HasCardManufacturingProfileId returns a boolean if a field has been set.
func (o *CardOrder) HasCardManufacturingProfileId() bool {
	if o != nil && !common.IsNil(o.CardManufacturingProfileId) {
		return true
	}

	return false
}

// SetCardManufacturingProfileId gets a reference to the given string and assigns it to the CardManufacturingProfileId field.
func (o *CardOrder) SetCardManufacturingProfileId(v string) {
	o.CardManufacturingProfileId = &v
}

// GetClosedDate returns the ClosedDate field value if set, zero value otherwise.
func (o *CardOrder) GetClosedDate() time.Time {
	if o == nil || common.IsNil(o.ClosedDate) {
		var ret time.Time
		return ret
	}
	return *o.ClosedDate
}

// GetClosedDateOk returns a tuple with the ClosedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetClosedDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ClosedDate) {
		return nil, false
	}
	return o.ClosedDate, true
}

// HasClosedDate returns a boolean if a field has been set.
func (o *CardOrder) HasClosedDate() bool {
	if o != nil && !common.IsNil(o.ClosedDate) {
		return true
	}

	return false
}

// SetClosedDate gets a reference to the given time.Time and assigns it to the ClosedDate field.
func (o *CardOrder) SetClosedDate(v time.Time) {
	o.ClosedDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CardOrder) GetEndDate() time.Time {
	if o == nil || common.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetEndDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CardOrder) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CardOrder) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CardOrder) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CardOrder) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CardOrder) SetId(v string) {
	o.Id = &v
}

// GetLockDate returns the LockDate field value if set, zero value otherwise.
func (o *CardOrder) GetLockDate() time.Time {
	if o == nil || common.IsNil(o.LockDate) {
		var ret time.Time
		return ret
	}
	return *o.LockDate
}

// GetLockDateOk returns a tuple with the LockDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetLockDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.LockDate) {
		return nil, false
	}
	return o.LockDate, true
}

// HasLockDate returns a boolean if a field has been set.
func (o *CardOrder) HasLockDate() bool {
	if o != nil && !common.IsNil(o.LockDate) {
		return true
	}

	return false
}

// SetLockDate gets a reference to the given time.Time and assigns it to the LockDate field.
func (o *CardOrder) SetLockDate(v time.Time) {
	o.LockDate = &v
}

// GetServiceCenter returns the ServiceCenter field value if set, zero value otherwise.
func (o *CardOrder) GetServiceCenter() string {
	if o == nil || common.IsNil(o.ServiceCenter) {
		var ret string
		return ret
	}
	return *o.ServiceCenter
}

// GetServiceCenterOk returns a tuple with the ServiceCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetServiceCenterOk() (*string, bool) {
	if o == nil || common.IsNil(o.ServiceCenter) {
		return nil, false
	}
	return o.ServiceCenter, true
}

// HasServiceCenter returns a boolean if a field has been set.
func (o *CardOrder) HasServiceCenter() bool {
	if o != nil && !common.IsNil(o.ServiceCenter) {
		return true
	}

	return false
}

// SetServiceCenter gets a reference to the given string and assigns it to the ServiceCenter field.
func (o *CardOrder) SetServiceCenter(v string) {
	o.ServiceCenter = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CardOrder) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrder) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CardOrder) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CardOrder) SetStatus(v string) {
	o.Status = &v
}

func (o CardOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BeginDate) {
		toSerialize["beginDate"] = o.BeginDate
	}
	if !common.IsNil(o.CardManufacturingProfileId) {
		toSerialize["cardManufacturingProfileId"] = o.CardManufacturingProfileId
	}
	if !common.IsNil(o.ClosedDate) {
		toSerialize["closedDate"] = o.ClosedDate
	}
	if !common.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.LockDate) {
		toSerialize["lockDate"] = o.LockDate
	}
	if !common.IsNil(o.ServiceCenter) {
		toSerialize["serviceCenter"] = o.ServiceCenter
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCardOrder struct {
	value *CardOrder
	isSet bool
}

func (v NullableCardOrder) Get() *CardOrder {
	return v.value
}

func (v *NullableCardOrder) Set(val *CardOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCardOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCardOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardOrder(val *CardOrder) *NullableCardOrder {
	return &NullableCardOrder{value: val, isSet: true}
}

func (v NullableCardOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CardOrder) isValidStatus() bool {
	var allowedEnumValues = []string{"closed", "open"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
