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

// checks if the TransferServiceRestServiceError type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferServiceRestServiceError{}

// TransferServiceRestServiceError struct for TransferServiceRestServiceError
type TransferServiceRestServiceError struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`
	// A code that identifies the problem type.
	ErrorCode string `json:"errorCode"`
	// A unique URI that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`
	// Detailed explanation of each validation error, when applicable.
	InvalidFields []InvalidField `json:"invalidFields,omitempty"`
	// A unique reference for the request, essentially the same as `pspReference`.
	RequestId *string                `json:"requestId,omitempty"`
	Response  map[string]interface{} `json:"response,omitempty"`
	// Detailed explanation of each attempt to route the transfer with the priorities from the request.
	RoutingDetails []RoutingDetails `json:"routingDetails,omitempty"`
	// The HTTP status code.
	Status int32 `json:"status"`
	// A short, human-readable summary of the problem type.
	Title string `json:"title"`
	// A URI that identifies the problem type, pointing to human-readable documentation on this problem type.
	Type string `json:"type"`
}

// NewTransferServiceRestServiceError instantiates a new TransferServiceRestServiceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferServiceRestServiceError(detail string, errorCode string, status int32, title string, type_ string) *TransferServiceRestServiceError {
	this := TransferServiceRestServiceError{}
	this.Detail = detail
	this.ErrorCode = errorCode
	this.Status = status
	this.Title = title
	this.Type = type_
	return &this
}

// NewTransferServiceRestServiceErrorWithDefaults instantiates a new TransferServiceRestServiceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferServiceRestServiceErrorWithDefaults() *TransferServiceRestServiceError {
	this := TransferServiceRestServiceError{}
	return &this
}

// GetDetail returns the Detail field value
func (o *TransferServiceRestServiceError) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *TransferServiceRestServiceError) SetDetail(v string) {
	o.Detail = v
}

// GetErrorCode returns the ErrorCode field value
func (o *TransferServiceRestServiceError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *TransferServiceRestServiceError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *TransferServiceRestServiceError) GetInstance() string {
	if o == nil || common.IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetInstanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *TransferServiceRestServiceError) HasInstance() bool {
	if o != nil && !common.IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *TransferServiceRestServiceError) SetInstance(v string) {
	o.Instance = &v
}

// GetInvalidFields returns the InvalidFields field value if set, zero value otherwise.
func (o *TransferServiceRestServiceError) GetInvalidFields() []InvalidField {
	if o == nil || common.IsNil(o.InvalidFields) {
		var ret []InvalidField
		return ret
	}
	return o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetInvalidFieldsOk() ([]InvalidField, bool) {
	if o == nil || common.IsNil(o.InvalidFields) {
		return nil, false
	}
	return o.InvalidFields, true
}

// HasInvalidFields returns a boolean if a field has been set.
func (o *TransferServiceRestServiceError) HasInvalidFields() bool {
	if o != nil && !common.IsNil(o.InvalidFields) {
		return true
	}

	return false
}

// SetInvalidFields gets a reference to the given []InvalidField and assigns it to the InvalidFields field.
func (o *TransferServiceRestServiceError) SetInvalidFields(v []InvalidField) {
	o.InvalidFields = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *TransferServiceRestServiceError) GetRequestId() string {
	if o == nil || common.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *TransferServiceRestServiceError) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *TransferServiceRestServiceError) SetRequestId(v string) {
	o.RequestId = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *TransferServiceRestServiceError) GetResponse() map[string]interface{} {
	if o == nil || common.IsNil(o.Response) {
		var ret map[string]interface{}
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetResponseOk() (map[string]interface{}, bool) {
	if o == nil || common.IsNil(o.Response) {
		return map[string]interface{}{}, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *TransferServiceRestServiceError) HasResponse() bool {
	if o != nil && !common.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given map[string]interface{} and assigns it to the Response field.
func (o *TransferServiceRestServiceError) SetResponse(v map[string]interface{}) {
	o.Response = v
}

// GetRoutingDetails returns the RoutingDetails field value if set, zero value otherwise.
func (o *TransferServiceRestServiceError) GetRoutingDetails() []RoutingDetails {
	if o == nil || common.IsNil(o.RoutingDetails) {
		var ret []RoutingDetails
		return ret
	}
	return o.RoutingDetails
}

// GetRoutingDetailsOk returns a tuple with the RoutingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetRoutingDetailsOk() ([]RoutingDetails, bool) {
	if o == nil || common.IsNil(o.RoutingDetails) {
		return nil, false
	}
	return o.RoutingDetails, true
}

// HasRoutingDetails returns a boolean if a field has been set.
func (o *TransferServiceRestServiceError) HasRoutingDetails() bool {
	if o != nil && !common.IsNil(o.RoutingDetails) {
		return true
	}

	return false
}

// SetRoutingDetails gets a reference to the given []RoutingDetails and assigns it to the RoutingDetails field.
func (o *TransferServiceRestServiceError) SetRoutingDetails(v []RoutingDetails) {
	o.RoutingDetails = v
}

// GetStatus returns the Status field value
func (o *TransferServiceRestServiceError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferServiceRestServiceError) SetStatus(v int32) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *TransferServiceRestServiceError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TransferServiceRestServiceError) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *TransferServiceRestServiceError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferServiceRestServiceError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferServiceRestServiceError) SetType(v string) {
	o.Type = v
}

func (o TransferServiceRestServiceError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferServiceRestServiceError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	toSerialize["errorCode"] = o.ErrorCode
	if !common.IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !common.IsNil(o.InvalidFields) {
		toSerialize["invalidFields"] = o.InvalidFields
	}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !common.IsNil(o.RoutingDetails) {
		toSerialize["routingDetails"] = o.RoutingDetails
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransferServiceRestServiceError struct {
	value *TransferServiceRestServiceError
	isSet bool
}

func (v NullableTransferServiceRestServiceError) Get() *TransferServiceRestServiceError {
	return v.value
}

func (v *NullableTransferServiceRestServiceError) Set(val *TransferServiceRestServiceError) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferServiceRestServiceError) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferServiceRestServiceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferServiceRestServiceError(val *TransferServiceRestServiceError) *NullableTransferServiceRestServiceError {
	return &NullableTransferServiceRestServiceError{value: val, isSet: true}
}

func (v NullableTransferServiceRestServiceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferServiceRestServiceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
