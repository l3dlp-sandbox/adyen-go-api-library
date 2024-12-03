/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the PaymentDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentDetailsResponse{}

// PaymentDetailsResponse struct for PaymentDetailsResponse
type PaymentDetailsResponse struct {
	// Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** > **Developers** > **Additional data**.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Amount *Amount `json:"amount,omitempty"`
	// Donation Token containing payment details for Adyen Giving.
	DonationToken *string `json:"donationToken,omitempty"`
	FraudResult *FraudResult `json:"fraudResult,omitempty"`
	// The reference used during the /payments request.
	MerchantReference *string `json:"merchantReference,omitempty"`
	Order *CheckoutOrderResponse `json:"order,omitempty"`
	PaymentMethod *ResponsePaymentMethod `json:"paymentMethod,omitempty"`
	// Adyen's 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// If the payment's authorisation is refused or an error occurs during authorisation, this field holds Adyen's mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReason *string `json:"refusalReason,omitempty"`
	// Code that specifies the refusal reason. For more information, see [Authorisation refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReasonCode *string `json:"refusalReasonCode,omitempty"`
	// The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the `refusalReason` field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper's device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **PartiallyAuthorised** – The payment has been authorised for a partial amount. This happens for card payments when the merchant supports Partial Authorisations and the cardholder has insufficient funds. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the `refusalReason` field. This is a final state.
	ResultCode *string `json:"resultCode,omitempty"`
	// The shopperLocale.
	ShopperLocale *string `json:"shopperLocale,omitempty"`
	ThreeDS2ResponseData *ThreeDS2ResponseData `json:"threeDS2ResponseData,omitempty"`
	ThreeDS2Result *ThreeDS2Result `json:"threeDS2Result,omitempty"`
	// When non-empty, contains a value that you must submit to the `/payments/details` endpoint as `paymentData`.
	ThreeDSPaymentData *string `json:"threeDSPaymentData,omitempty"`
}

// NewPaymentDetailsResponse instantiates a new PaymentDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDetailsResponse() *PaymentDetailsResponse {
	this := PaymentDetailsResponse{}
	return &this
}

// NewPaymentDetailsResponseWithDefaults instantiates a new PaymentDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDetailsResponseWithDefaults() *PaymentDetailsResponse {
	this := PaymentDetailsResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *PaymentDetailsResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *PaymentDetailsResponse) SetAmount(v Amount) {
	o.Amount = &v
}

// GetDonationToken returns the DonationToken field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetDonationToken() string {
	if o == nil || common.IsNil(o.DonationToken) {
		var ret string
		return ret
	}
	return *o.DonationToken
}

// GetDonationTokenOk returns a tuple with the DonationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetDonationTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.DonationToken) {
		return nil, false
	}
	return o.DonationToken, true
}

// HasDonationToken returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasDonationToken() bool {
	if o != nil && !common.IsNil(o.DonationToken) {
		return true
	}

	return false
}

// SetDonationToken gets a reference to the given string and assigns it to the DonationToken field.
func (o *PaymentDetailsResponse) SetDonationToken(v string) {
	o.DonationToken = &v
}

// GetFraudResult returns the FraudResult field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetFraudResult() FraudResult {
	if o == nil || common.IsNil(o.FraudResult) {
		var ret FraudResult
		return ret
	}
	return *o.FraudResult
}

// GetFraudResultOk returns a tuple with the FraudResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetFraudResultOk() (*FraudResult, bool) {
	if o == nil || common.IsNil(o.FraudResult) {
		return nil, false
	}
	return o.FraudResult, true
}

// HasFraudResult returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasFraudResult() bool {
	if o != nil && !common.IsNil(o.FraudResult) {
		return true
	}

	return false
}

// SetFraudResult gets a reference to the given FraudResult and assigns it to the FraudResult field.
func (o *PaymentDetailsResponse) SetFraudResult(v FraudResult) {
	o.FraudResult = &v
}

// GetMerchantReference returns the MerchantReference field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetMerchantReference() string {
	if o == nil || common.IsNil(o.MerchantReference) {
		var ret string
		return ret
	}
	return *o.MerchantReference
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantReference) {
		return nil, false
	}
	return o.MerchantReference, true
}

// HasMerchantReference returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasMerchantReference() bool {
	if o != nil && !common.IsNil(o.MerchantReference) {
		return true
	}

	return false
}

// SetMerchantReference gets a reference to the given string and assigns it to the MerchantReference field.
func (o *PaymentDetailsResponse) SetMerchantReference(v string) {
	o.MerchantReference = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetOrder() CheckoutOrderResponse {
	if o == nil || common.IsNil(o.Order) {
		var ret CheckoutOrderResponse
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetOrderOk() (*CheckoutOrderResponse, bool) {
	if o == nil || common.IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasOrder() bool {
	if o != nil && !common.IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given CheckoutOrderResponse and assigns it to the Order field.
func (o *PaymentDetailsResponse) SetOrder(v CheckoutOrderResponse) {
	o.Order = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetPaymentMethod() ResponsePaymentMethod {
	if o == nil || common.IsNil(o.PaymentMethod) {
		var ret ResponsePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetPaymentMethodOk() (*ResponsePaymentMethod, bool) {
	if o == nil || common.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasPaymentMethod() bool {
	if o != nil && !common.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given ResponsePaymentMethod and assigns it to the PaymentMethod field.
func (o *PaymentDetailsResponse) SetPaymentMethod(v ResponsePaymentMethod) {
	o.PaymentMethod = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *PaymentDetailsResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetRefusalReason() string {
	if o == nil || common.IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasRefusalReason() bool {
	if o != nil && !common.IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *PaymentDetailsResponse) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetRefusalReasonCode returns the RefusalReasonCode field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetRefusalReasonCode() string {
	if o == nil || common.IsNil(o.RefusalReasonCode) {
		var ret string
		return ret
	}
	return *o.RefusalReasonCode
}

// GetRefusalReasonCodeOk returns a tuple with the RefusalReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetRefusalReasonCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReasonCode) {
		return nil, false
	}
	return o.RefusalReasonCode, true
}

// HasRefusalReasonCode returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasRefusalReasonCode() bool {
	if o != nil && !common.IsNil(o.RefusalReasonCode) {
		return true
	}

	return false
}

// SetRefusalReasonCode gets a reference to the given string and assigns it to the RefusalReasonCode field.
func (o *PaymentDetailsResponse) SetRefusalReasonCode(v string) {
	o.RefusalReasonCode = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetResultCode() string {
	if o == nil || common.IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasResultCode() bool {
	if o != nil && !common.IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *PaymentDetailsResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetShopperLocale returns the ShopperLocale field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetShopperLocale() string {
	if o == nil || common.IsNil(o.ShopperLocale) {
		var ret string
		return ret
	}
	return *o.ShopperLocale
}

// GetShopperLocaleOk returns a tuple with the ShopperLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetShopperLocaleOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperLocale) {
		return nil, false
	}
	return o.ShopperLocale, true
}

// HasShopperLocale returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasShopperLocale() bool {
	if o != nil && !common.IsNil(o.ShopperLocale) {
		return true
	}

	return false
}

// SetShopperLocale gets a reference to the given string and assigns it to the ShopperLocale field.
func (o *PaymentDetailsResponse) SetShopperLocale(v string) {
	o.ShopperLocale = &v
}

// GetThreeDS2ResponseData returns the ThreeDS2ResponseData field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetThreeDS2ResponseData() ThreeDS2ResponseData {
	if o == nil || common.IsNil(o.ThreeDS2ResponseData) {
		var ret ThreeDS2ResponseData
		return ret
	}
	return *o.ThreeDS2ResponseData
}

// GetThreeDS2ResponseDataOk returns a tuple with the ThreeDS2ResponseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetThreeDS2ResponseDataOk() (*ThreeDS2ResponseData, bool) {
	if o == nil || common.IsNil(o.ThreeDS2ResponseData) {
		return nil, false
	}
	return o.ThreeDS2ResponseData, true
}

// HasThreeDS2ResponseData returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasThreeDS2ResponseData() bool {
	if o != nil && !common.IsNil(o.ThreeDS2ResponseData) {
		return true
	}

	return false
}

// SetThreeDS2ResponseData gets a reference to the given ThreeDS2ResponseData and assigns it to the ThreeDS2ResponseData field.
func (o *PaymentDetailsResponse) SetThreeDS2ResponseData(v ThreeDS2ResponseData) {
	o.ThreeDS2ResponseData = &v
}

// GetThreeDS2Result returns the ThreeDS2Result field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetThreeDS2Result() ThreeDS2Result {
	if o == nil || common.IsNil(o.ThreeDS2Result) {
		var ret ThreeDS2Result
		return ret
	}
	return *o.ThreeDS2Result
}

// GetThreeDS2ResultOk returns a tuple with the ThreeDS2Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetThreeDS2ResultOk() (*ThreeDS2Result, bool) {
	if o == nil || common.IsNil(o.ThreeDS2Result) {
		return nil, false
	}
	return o.ThreeDS2Result, true
}

// HasThreeDS2Result returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasThreeDS2Result() bool {
	if o != nil && !common.IsNil(o.ThreeDS2Result) {
		return true
	}

	return false
}

// SetThreeDS2Result gets a reference to the given ThreeDS2Result and assigns it to the ThreeDS2Result field.
func (o *PaymentDetailsResponse) SetThreeDS2Result(v ThreeDS2Result) {
	o.ThreeDS2Result = &v
}

// GetThreeDSPaymentData returns the ThreeDSPaymentData field value if set, zero value otherwise.
func (o *PaymentDetailsResponse) GetThreeDSPaymentData() string {
	if o == nil || common.IsNil(o.ThreeDSPaymentData) {
		var ret string
		return ret
	}
	return *o.ThreeDSPaymentData
}

// GetThreeDSPaymentDataOk returns a tuple with the ThreeDSPaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsResponse) GetThreeDSPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSPaymentData) {
		return nil, false
	}
	return o.ThreeDSPaymentData, true
}

// HasThreeDSPaymentData returns a boolean if a field has been set.
func (o *PaymentDetailsResponse) HasThreeDSPaymentData() bool {
	if o != nil && !common.IsNil(o.ThreeDSPaymentData) {
		return true
	}

	return false
}

// SetThreeDSPaymentData gets a reference to the given string and assigns it to the ThreeDSPaymentData field.
func (o *PaymentDetailsResponse) SetThreeDSPaymentData(v string) {
	o.ThreeDSPaymentData = &v
}

func (o PaymentDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.DonationToken) {
		toSerialize["donationToken"] = o.DonationToken
	}
	if !common.IsNil(o.FraudResult) {
		toSerialize["fraudResult"] = o.FraudResult
	}
	if !common.IsNil(o.MerchantReference) {
		toSerialize["merchantReference"] = o.MerchantReference
	}
	if !common.IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !common.IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	if !common.IsNil(o.RefusalReasonCode) {
		toSerialize["refusalReasonCode"] = o.RefusalReasonCode
	}
	if !common.IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !common.IsNil(o.ShopperLocale) {
		toSerialize["shopperLocale"] = o.ShopperLocale
	}
	if !common.IsNil(o.ThreeDS2ResponseData) {
		toSerialize["threeDS2ResponseData"] = o.ThreeDS2ResponseData
	}
	if !common.IsNil(o.ThreeDS2Result) {
		toSerialize["threeDS2Result"] = o.ThreeDS2Result
	}
	if !common.IsNil(o.ThreeDSPaymentData) {
		toSerialize["threeDSPaymentData"] = o.ThreeDSPaymentData
	}
	return toSerialize, nil
}

type NullablePaymentDetailsResponse struct {
	value *PaymentDetailsResponse
	isSet bool
}

func (v NullablePaymentDetailsResponse) Get() *PaymentDetailsResponse {
	return v.value
}

func (v *NullablePaymentDetailsResponse) Set(val *PaymentDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDetailsResponse(val *PaymentDetailsResponse) *NullablePaymentDetailsResponse {
	return &NullablePaymentDetailsResponse{value: val, isSet: true}
}

func (v NullablePaymentDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *PaymentDetailsResponse) isValidResultCode() bool {
    var allowedEnumValues = []string{ "AuthenticationFinished", "AuthenticationNotRequired", "Authorised", "Cancelled", "ChallengeShopper", "Error", "IdentifyShopper", "PartiallyAuthorised", "Pending", "PresentToShopper", "Received", "RedirectShopper", "Refused", "Success" }
    for _, allowed := range allowedEnumValues {
        if o.GetResultCode() == allowed {
            return true
        }
    }
    return false
}

