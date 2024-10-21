/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v13/src/common"
)

// checks if the CreateOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateOrderResponse{}

// CreateOrderResponse struct for CreateOrderResponse
type CreateOrderResponse struct {
	// Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** > **Developers** > **Additional data**.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Amount         Amount             `json:"amount"`
	// The date that the order will expire.
	ExpiresAt   string       `json:"expiresAt"`
	FraudResult *FraudResult `json:"fraudResult,omitempty"`
	// The encrypted data that will be used by merchant for adding payments to the order.
	OrderData string `json:"orderData"`
	// Adyen's 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// The reference provided by merchant for creating the order.
	Reference *string `json:"reference,omitempty"`
	// If the payment's authorisation is refused or an error occurs during authorisation, this field holds Adyen's mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReason   *string `json:"refusalReason,omitempty"`
	RemainingAmount Amount  `json:"remainingAmount"`
	// The result of the order creation request.  The value is always **Success**.
	ResultCode string `json:"resultCode"`
}

// NewCreateOrderResponse instantiates a new CreateOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderResponse(amount Amount, expiresAt string, orderData string, remainingAmount Amount, resultCode string) *CreateOrderResponse {
	this := CreateOrderResponse{}
	this.Amount = amount
	this.ExpiresAt = expiresAt
	this.OrderData = orderData
	this.RemainingAmount = remainingAmount
	this.ResultCode = resultCode
	return &this
}

// NewCreateOrderResponseWithDefaults instantiates a new CreateOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderResponseWithDefaults() *CreateOrderResponse {
	this := CreateOrderResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *CreateOrderResponse) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *CreateOrderResponse) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *CreateOrderResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAmount returns the Amount field value
func (o *CreateOrderResponse) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateOrderResponse) SetAmount(v Amount) {
	o.Amount = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *CreateOrderResponse) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *CreateOrderResponse) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

// GetFraudResult returns the FraudResult field value if set, zero value otherwise.
func (o *CreateOrderResponse) GetFraudResult() FraudResult {
	if o == nil || common.IsNil(o.FraudResult) {
		var ret FraudResult
		return ret
	}
	return *o.FraudResult
}

// GetFraudResultOk returns a tuple with the FraudResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetFraudResultOk() (*FraudResult, bool) {
	if o == nil || common.IsNil(o.FraudResult) {
		return nil, false
	}
	return o.FraudResult, true
}

// HasFraudResult returns a boolean if a field has been set.
func (o *CreateOrderResponse) HasFraudResult() bool {
	if o != nil && !common.IsNil(o.FraudResult) {
		return true
	}

	return false
}

// SetFraudResult gets a reference to the given FraudResult and assigns it to the FraudResult field.
func (o *CreateOrderResponse) SetFraudResult(v FraudResult) {
	o.FraudResult = &v
}

// GetOrderData returns the OrderData field value
func (o *CreateOrderResponse) GetOrderData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderData
}

// GetOrderDataOk returns a tuple with the OrderData field value
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetOrderDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderData, true
}

// SetOrderData sets field value
func (o *CreateOrderResponse) SetOrderData(v string) {
	o.OrderData = v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *CreateOrderResponse) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *CreateOrderResponse) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *CreateOrderResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CreateOrderResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CreateOrderResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CreateOrderResponse) SetReference(v string) {
	o.Reference = &v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *CreateOrderResponse) GetRefusalReason() string {
	if o == nil || common.IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *CreateOrderResponse) HasRefusalReason() bool {
	if o != nil && !common.IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *CreateOrderResponse) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetRemainingAmount returns the RemainingAmount field value
func (o *CreateOrderResponse) GetRemainingAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.RemainingAmount
}

// GetRemainingAmountOk returns a tuple with the RemainingAmount field value
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetRemainingAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingAmount, true
}

// SetRemainingAmount sets field value
func (o *CreateOrderResponse) SetRemainingAmount(v Amount) {
	o.RemainingAmount = v
}

// GetResultCode returns the ResultCode field value
func (o *CreateOrderResponse) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *CreateOrderResponse) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *CreateOrderResponse) SetResultCode(v string) {
	o.ResultCode = v
}

func (o CreateOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["amount"] = o.Amount
	toSerialize["expiresAt"] = o.ExpiresAt
	if !common.IsNil(o.FraudResult) {
		toSerialize["fraudResult"] = o.FraudResult
	}
	toSerialize["orderData"] = o.OrderData
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	toSerialize["remainingAmount"] = o.RemainingAmount
	toSerialize["resultCode"] = o.ResultCode
	return toSerialize, nil
}

type NullableCreateOrderResponse struct {
	value *CreateOrderResponse
	isSet bool
}

func (v NullableCreateOrderResponse) Get() *CreateOrderResponse {
	return v.value
}

func (v *NullableCreateOrderResponse) Set(val *CreateOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderResponse(val *CreateOrderResponse) *NullableCreateOrderResponse {
	return &NullableCreateOrderResponse{value: val, isSet: true}
}

func (v NullableCreateOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CreateOrderResponse) isValidResultCode() bool {
	var allowedEnumValues = []string{"Success"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}
