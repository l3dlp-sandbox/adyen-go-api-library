/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the TestOutput type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TestOutput{}

// TestOutput struct for TestOutput
type TestOutput struct {
	// Unique identifier of the merchant account that the notification is about.
	MerchantId *string `json:"merchantId,omitempty"`
	// The response your server returned for the test webhook.  Your server must respond with **[accepted]** for the test webhook to be successful (`data.status`: **success**). Find out more about [accepting notifications](https://docs.adyen.com/development-resources/webhooks#accept-notifications)  You can use the value of this field together with the [`responseCode`](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-responseCode) value to troubleshoot unsuccessful test webhooks.
	Output *string `json:"output,omitempty"`
	// The [body of the notification webhook](https://docs.adyen.com/development-resources/webhooks/understand-notifications#notification-structure) that was sent to your server.
	RequestSent *string `json:"requestSent,omitempty"`
	// The HTTP response code for your server's response to the test webhook.  You can use the value of this field together with the the [`output`](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-output) field value to troubleshoot failed test webhooks.
	ResponseCode *string `json:"responseCode,omitempty"`
	// The time between sending the test webhook and receiving the response from your server. You can use it as an indication of how long your server takes to process a webhook notification. Measured in milliseconds, for example **304 ms**.
	ResponseTime *string `json:"responseTime,omitempty"`
	// The status of the test request. Possible values are: * **success**, if `data.output`: **[accepted]** and `data.responseCode`: **200**. * **failed**, in all other cases.  You can use the value of the [`output`](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-output) field together with the [`responseCode`](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/merchants/{merchantId}/webhooks/{id}/test__resParam_data-responseCode) value to troubleshoot failed test webhooks.
	Status string `json:"status"`
}

// NewTestOutput instantiates a new TestOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestOutput(status string) *TestOutput {
	this := TestOutput{}
	this.Status = status
	return &this
}

// NewTestOutputWithDefaults instantiates a new TestOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestOutputWithDefaults() *TestOutput {
	this := TestOutput{}
	return &this
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *TestOutput) GetMerchantId() string {
	if o == nil || common.IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOutput) GetMerchantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *TestOutput) HasMerchantId() bool {
	if o != nil && !common.IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *TestOutput) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *TestOutput) GetOutput() string {
	if o == nil || common.IsNil(o.Output) {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOutput) GetOutputOk() (*string, bool) {
	if o == nil || common.IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *TestOutput) HasOutput() bool {
	if o != nil && !common.IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *TestOutput) SetOutput(v string) {
	o.Output = &v
}

// GetRequestSent returns the RequestSent field value if set, zero value otherwise.
func (o *TestOutput) GetRequestSent() string {
	if o == nil || common.IsNil(o.RequestSent) {
		var ret string
		return ret
	}
	return *o.RequestSent
}

// GetRequestSentOk returns a tuple with the RequestSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOutput) GetRequestSentOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestSent) {
		return nil, false
	}
	return o.RequestSent, true
}

// HasRequestSent returns a boolean if a field has been set.
func (o *TestOutput) HasRequestSent() bool {
	if o != nil && !common.IsNil(o.RequestSent) {
		return true
	}

	return false
}

// SetRequestSent gets a reference to the given string and assigns it to the RequestSent field.
func (o *TestOutput) SetRequestSent(v string) {
	o.RequestSent = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *TestOutput) GetResponseCode() string {
	if o == nil || common.IsNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOutput) GetResponseCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *TestOutput) HasResponseCode() bool {
	if o != nil && !common.IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *TestOutput) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *TestOutput) GetResponseTime() string {
	if o == nil || common.IsNil(o.ResponseTime) {
		var ret string
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOutput) GetResponseTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *TestOutput) HasResponseTime() bool {
	if o != nil && !common.IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given string and assigns it to the ResponseTime field.
func (o *TestOutput) SetResponseTime(v string) {
	o.ResponseTime = &v
}

// GetStatus returns the Status field value
func (o *TestOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestOutput) SetStatus(v string) {
	o.Status = v
}

func (o TestOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !common.IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !common.IsNil(o.RequestSent) {
		toSerialize["requestSent"] = o.RequestSent
	}
	if !common.IsNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !common.IsNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableTestOutput struct {
	value *TestOutput
	isSet bool
}

func (v NullableTestOutput) Get() *TestOutput {
	return v.value
}

func (v *NullableTestOutput) Set(val *TestOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTestOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTestOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestOutput(val *TestOutput) *NullableTestOutput {
	return &NullableTestOutput{value: val, isSet: true}
}

func (v NullableTestOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
