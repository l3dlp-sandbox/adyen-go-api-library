/*
Dispute webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputewebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the DisputeEventNotification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DisputeEventNotification{}

// DisputeEventNotification struct for DisputeEventNotification
type DisputeEventNotification struct {
	// The unique Acquirer Reference Number (arn) generated by the card scheme for each capture. You can use the arn to trace the transaction through its lifecycle.
	Arn *string `json:"arn,omitempty"`
	// The unique identifier of the balance platform.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
	// The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Contains information about the dispute.
	Description    *string `json:"description,omitempty"`
	DisputedAmount *Amount `json:"disputedAmount,omitempty"`
	// The ID of the resource.
	Id *string `json:"id,omitempty"`
	// The current status of the dispute.
	Status *string `json:"status,omitempty"`
	// Additional information about the status of the dispute, when available.
	StatusDetail *string `json:"statusDetail,omitempty"`
	// The unique reference of the transaction for which the dispute is requested.
	TransactionId *string `json:"transactionId,omitempty"`
	// The type of dispute raised for the transaction.
	Type *string `json:"type,omitempty"`
}

// NewDisputeEventNotification instantiates a new DisputeEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisputeEventNotification() *DisputeEventNotification {
	this := DisputeEventNotification{}
	return &this
}

// NewDisputeEventNotificationWithDefaults instantiates a new DisputeEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisputeEventNotificationWithDefaults() *DisputeEventNotification {
	this := DisputeEventNotification{}
	return &this
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetArn() string {
	if o == nil || common.IsNil(o.Arn) {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetArnOk() (*string, bool) {
	if o == nil || common.IsNil(o.Arn) {
		return nil, false
	}
	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasArn() bool {
	if o != nil && !common.IsNil(o.Arn) {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *DisputeEventNotification) SetArn(v string) {
	o.Arn = &v
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *DisputeEventNotification) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *DisputeEventNotification) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DisputeEventNotification) SetDescription(v string) {
	o.Description = &v
}

// GetDisputedAmount returns the DisputedAmount field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetDisputedAmount() Amount {
	if o == nil || common.IsNil(o.DisputedAmount) {
		var ret Amount
		return ret
	}
	return *o.DisputedAmount
}

// GetDisputedAmountOk returns a tuple with the DisputedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetDisputedAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.DisputedAmount) {
		return nil, false
	}
	return o.DisputedAmount, true
}

// HasDisputedAmount returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasDisputedAmount() bool {
	if o != nil && !common.IsNil(o.DisputedAmount) {
		return true
	}

	return false
}

// SetDisputedAmount gets a reference to the given Amount and assigns it to the DisputedAmount field.
func (o *DisputeEventNotification) SetDisputedAmount(v Amount) {
	o.DisputedAmount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DisputeEventNotification) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DisputeEventNotification) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetStatusDetail() string {
	if o == nil || common.IsNil(o.StatusDetail) {
		var ret string
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetStatusDetailOk() (*string, bool) {
	if o == nil || common.IsNil(o.StatusDetail) {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasStatusDetail() bool {
	if o != nil && !common.IsNil(o.StatusDetail) {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given string and assigns it to the StatusDetail field.
func (o *DisputeEventNotification) SetStatusDetail(v string) {
	o.StatusDetail = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetTransactionId() string {
	if o == nil || common.IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetTransactionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasTransactionId() bool {
	if o != nil && !common.IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *DisputeEventNotification) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DisputeEventNotification) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeEventNotification) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DisputeEventNotification) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DisputeEventNotification) SetType(v string) {
	o.Type = &v
}

func (o DisputeEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisputeEventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Arn) {
		toSerialize["arn"] = o.Arn
	}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.DisputedAmount) {
		toSerialize["disputedAmount"] = o.DisputedAmount
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StatusDetail) {
		toSerialize["statusDetail"] = o.StatusDetail
	}
	if !common.IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDisputeEventNotification struct {
	value *DisputeEventNotification
	isSet bool
}

func (v NullableDisputeEventNotification) Get() *DisputeEventNotification {
	return v.value
}

func (v *NullableDisputeEventNotification) Set(val *DisputeEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableDisputeEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableDisputeEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisputeEventNotification(val *DisputeEventNotification) *NullableDisputeEventNotification {
	return &NullableDisputeEventNotification{value: val, isSet: true}
}

func (v NullableDisputeEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisputeEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DisputeEventNotification) isValidType() bool {
	var allowedEnumValues = []string{"fraud", "notDelivered"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
