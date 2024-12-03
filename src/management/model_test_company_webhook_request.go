/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TestCompanyWebhookRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TestCompanyWebhookRequest{}

// TestCompanyWebhookRequest struct for TestCompanyWebhookRequest
type TestCompanyWebhookRequest struct {
	// List of `merchantId` values for which test webhooks will be sent. The list can have a maximum of 20 `merchantId` values.  If not specified, we send sample notifications to all the merchant accounts that the webhook is configured for. If this is more than 20 merchant accounts, use this list to specify a subset of the merchant accounts for which to send test notifications.
	MerchantIds []string `json:"merchantIds,omitempty"`
	Notification *CustomNotification `json:"notification,omitempty"`
	// List of event codes for which to send test notifications. Only the webhook types below are supported.   Possible values if webhook `type`: **standard**:  * **AUTHORISATION** * **CHARGEBACK_REVERSED** * **ORDER_CLOSED** * **ORDER_OPENED** * **PAIDOUT_REVERSED** * **PAYOUT_THIRDPARTY** * **REFUNDED_REVERSED** * **REFUND_WITH_DATA** * **REPORT_AVAILABLE** * **CUSTOM** - set your custom notification fields in the [`notification`](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/companies/{companyId}/webhooks/{webhookId}/test__reqParam_notification) object.  Possible values if webhook `type`: **banktransfer-notification**:  * **PENDING**  Possible values if webhook `type`: **report-notification**:  * **REPORT_AVAILABLE**  Possible values if webhook `type`: **ideal-notification**:  * **AUTHORISATION**  Possible values if webhook `type`: **pending-notification**:  * **PENDING** 
	Types []string `json:"types,omitempty"`
}

// NewTestCompanyWebhookRequest instantiates a new TestCompanyWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCompanyWebhookRequest() *TestCompanyWebhookRequest {
	this := TestCompanyWebhookRequest{}
	return &this
}

// NewTestCompanyWebhookRequestWithDefaults instantiates a new TestCompanyWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCompanyWebhookRequestWithDefaults() *TestCompanyWebhookRequest {
	this := TestCompanyWebhookRequest{}
	return &this
}

// GetMerchantIds returns the MerchantIds field value if set, zero value otherwise.
func (o *TestCompanyWebhookRequest) GetMerchantIds() []string {
	if o == nil || common.IsNil(o.MerchantIds) {
		var ret []string
		return ret
	}
	return o.MerchantIds
}

// GetMerchantIdsOk returns a tuple with the MerchantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCompanyWebhookRequest) GetMerchantIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.MerchantIds) {
		return nil, false
	}
	return o.MerchantIds, true
}

// HasMerchantIds returns a boolean if a field has been set.
func (o *TestCompanyWebhookRequest) HasMerchantIds() bool {
	if o != nil && !common.IsNil(o.MerchantIds) {
		return true
	}

	return false
}

// SetMerchantIds gets a reference to the given []string and assigns it to the MerchantIds field.
func (o *TestCompanyWebhookRequest) SetMerchantIds(v []string) {
	o.MerchantIds = v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *TestCompanyWebhookRequest) GetNotification() CustomNotification {
	if o == nil || common.IsNil(o.Notification) {
		var ret CustomNotification
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCompanyWebhookRequest) GetNotificationOk() (*CustomNotification, bool) {
	if o == nil || common.IsNil(o.Notification) {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *TestCompanyWebhookRequest) HasNotification() bool {
	if o != nil && !common.IsNil(o.Notification) {
		return true
	}

	return false
}

// SetNotification gets a reference to the given CustomNotification and assigns it to the Notification field.
func (o *TestCompanyWebhookRequest) SetNotification(v CustomNotification) {
	o.Notification = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *TestCompanyWebhookRequest) GetTypes() []string {
	if o == nil || common.IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCompanyWebhookRequest) GetTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *TestCompanyWebhookRequest) HasTypes() bool {
	if o != nil && !common.IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *TestCompanyWebhookRequest) SetTypes(v []string) {
	o.Types = v
}

func (o TestCompanyWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCompanyWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MerchantIds) {
		toSerialize["merchantIds"] = o.MerchantIds
	}
	if !common.IsNil(o.Notification) {
		toSerialize["notification"] = o.Notification
	}
	if !common.IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	return toSerialize, nil
}

type NullableTestCompanyWebhookRequest struct {
	value *TestCompanyWebhookRequest
	isSet bool
}

func (v NullableTestCompanyWebhookRequest) Get() *TestCompanyWebhookRequest {
	return v.value
}

func (v *NullableTestCompanyWebhookRequest) Set(val *TestCompanyWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCompanyWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCompanyWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCompanyWebhookRequest(val *TestCompanyWebhookRequest) *NullableTestCompanyWebhookRequest {
	return &NullableTestCompanyWebhookRequest{value: val, isSet: true}
}

func (v NullableTestCompanyWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCompanyWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



