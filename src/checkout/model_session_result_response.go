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

// checks if the SessionResultResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SessionResultResponse{}

// SessionResultResponse struct for SessionResultResponse
type SessionResultResponse struct {
	// A unique identifier of the session.
	Id *string `json:"id,omitempty"`
	// The status of the session. The status included in the response doesn't get updated. Don't make the request again to check for payment status updates.  Possible values:           * **completed** – The shopper completed the payment. This means that the payment was authorized.          * **paymentPending** – The shopper is in the process of making the payment. This applies to payment methods with an asynchronous flow.          * **refused** – The session has been refused, due to too many refused payment attempts. Shoppers can no longer complete the payment with this session.          * **canceled** – The shopper canceled the payment.          * **active** – The session is still active and can be paid.          * **expired** – The session expired (default: 1 hour after session creation). Shoppers can no longer complete the payment with this session.
	Status *string `json:"status,omitempty"`
}

// NewSessionResultResponse instantiates a new SessionResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionResultResponse() *SessionResultResponse {
	this := SessionResultResponse{}
	return &this
}

// NewSessionResultResponseWithDefaults instantiates a new SessionResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionResultResponseWithDefaults() *SessionResultResponse {
	this := SessionResultResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SessionResultResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResultResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SessionResultResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SessionResultResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SessionResultResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResultResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SessionResultResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SessionResultResponse) SetStatus(v string) {
	o.Status = &v
}

func (o SessionResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSessionResultResponse struct {
	value *SessionResultResponse
	isSet bool
}

func (v NullableSessionResultResponse) Get() *SessionResultResponse {
	return v.value
}

func (v *NullableSessionResultResponse) Set(val *SessionResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionResultResponse(val *SessionResultResponse) *NullableSessionResultResponse {
	return &NullableSessionResultResponse{value: val, isSet: true}
}

func (v NullableSessionResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SessionResultResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "canceled", "completed", "expired", "paymentPending", "refused"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
