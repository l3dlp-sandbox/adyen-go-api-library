/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// checks if the SecureRemoteCommerceCheckoutData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SecureRemoteCommerceCheckoutData{}

// SecureRemoteCommerceCheckoutData struct for SecureRemoteCommerceCheckoutData
type SecureRemoteCommerceCheckoutData struct {
	// The Secure Remote Commerce checkout payload to process the payment with.
	CheckoutPayload *string `json:"checkoutPayload,omitempty"`
	// This is the unique identifier generated by SRC system to track and link SRC messages. Available within the present checkout session (e.g. received in an earlier API response during the present session).
	CorrelationId *string `json:"correlationId,omitempty"`
	// The [card verification code](https://docs.adyen.com/get-started-with-adyen/payment-glossary/#card-security-code-cvc-cvv-cid).
	Cvc *string `json:"cvc,omitempty"`
	// A unique identifier that represents the token associated with a card enrolled. Required for scheme 'mc'.
	DigitalCardId *string `json:"digitalCardId,omitempty"`
	// The Secure Remote Commerce scheme.
	Scheme *string `json:"scheme,omitempty"`
	// A unique identifier that represents the token associated with a card enrolled. Required for scheme 'visa'.
	TokenReference *string `json:"tokenReference,omitempty"`
}

// NewSecureRemoteCommerceCheckoutData instantiates a new SecureRemoteCommerceCheckoutData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecureRemoteCommerceCheckoutData() *SecureRemoteCommerceCheckoutData {
	this := SecureRemoteCommerceCheckoutData{}
	return &this
}

// NewSecureRemoteCommerceCheckoutDataWithDefaults instantiates a new SecureRemoteCommerceCheckoutData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecureRemoteCommerceCheckoutDataWithDefaults() *SecureRemoteCommerceCheckoutData {
	this := SecureRemoteCommerceCheckoutData{}
	return &this
}

// GetCheckoutPayload returns the CheckoutPayload field value if set, zero value otherwise.
func (o *SecureRemoteCommerceCheckoutData) GetCheckoutPayload() string {
	if o == nil || common.IsNil(o.CheckoutPayload) {
		var ret string
		return ret
	}
	return *o.CheckoutPayload
}

// GetCheckoutPayloadOk returns a tuple with the CheckoutPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteCommerceCheckoutData) GetCheckoutPayloadOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutPayload) {
		return nil, false
	}
	return o.CheckoutPayload, true
}

// HasCheckoutPayload returns a boolean if a field has been set.
func (o *SecureRemoteCommerceCheckoutData) HasCheckoutPayload() bool {
	if o != nil && !common.IsNil(o.CheckoutPayload) {
		return true
	}

	return false
}

// SetCheckoutPayload gets a reference to the given string and assigns it to the CheckoutPayload field.
func (o *SecureRemoteCommerceCheckoutData) SetCheckoutPayload(v string) {
	o.CheckoutPayload = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *SecureRemoteCommerceCheckoutData) GetCorrelationId() string {
	if o == nil || common.IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteCommerceCheckoutData) GetCorrelationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *SecureRemoteCommerceCheckoutData) HasCorrelationId() bool {
	if o != nil && !common.IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *SecureRemoteCommerceCheckoutData) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetCvc returns the Cvc field value if set, zero value otherwise.
func (o *SecureRemoteCommerceCheckoutData) GetCvc() string {
	if o == nil || common.IsNil(o.Cvc) {
		var ret string
		return ret
	}
	return *o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteCommerceCheckoutData) GetCvcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cvc) {
		return nil, false
	}
	return o.Cvc, true
}

// HasCvc returns a boolean if a field has been set.
func (o *SecureRemoteCommerceCheckoutData) HasCvc() bool {
	if o != nil && !common.IsNil(o.Cvc) {
		return true
	}

	return false
}

// SetCvc gets a reference to the given string and assigns it to the Cvc field.
func (o *SecureRemoteCommerceCheckoutData) SetCvc(v string) {
	o.Cvc = &v
}

// GetDigitalCardId returns the DigitalCardId field value if set, zero value otherwise.
func (o *SecureRemoteCommerceCheckoutData) GetDigitalCardId() string {
	if o == nil || common.IsNil(o.DigitalCardId) {
		var ret string
		return ret
	}
	return *o.DigitalCardId
}

// GetDigitalCardIdOk returns a tuple with the DigitalCardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteCommerceCheckoutData) GetDigitalCardIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DigitalCardId) {
		return nil, false
	}
	return o.DigitalCardId, true
}

// HasDigitalCardId returns a boolean if a field has been set.
func (o *SecureRemoteCommerceCheckoutData) HasDigitalCardId() bool {
	if o != nil && !common.IsNil(o.DigitalCardId) {
		return true
	}

	return false
}

// SetDigitalCardId gets a reference to the given string and assigns it to the DigitalCardId field.
func (o *SecureRemoteCommerceCheckoutData) SetDigitalCardId(v string) {
	o.DigitalCardId = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *SecureRemoteCommerceCheckoutData) GetScheme() string {
	if o == nil || common.IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteCommerceCheckoutData) GetSchemeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *SecureRemoteCommerceCheckoutData) HasScheme() bool {
	if o != nil && !common.IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *SecureRemoteCommerceCheckoutData) SetScheme(v string) {
	o.Scheme = &v
}

// GetTokenReference returns the TokenReference field value if set, zero value otherwise.
func (o *SecureRemoteCommerceCheckoutData) GetTokenReference() string {
	if o == nil || common.IsNil(o.TokenReference) {
		var ret string
		return ret
	}
	return *o.TokenReference
}

// GetTokenReferenceOk returns a tuple with the TokenReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureRemoteCommerceCheckoutData) GetTokenReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenReference) {
		return nil, false
	}
	return o.TokenReference, true
}

// HasTokenReference returns a boolean if a field has been set.
func (o *SecureRemoteCommerceCheckoutData) HasTokenReference() bool {
	if o != nil && !common.IsNil(o.TokenReference) {
		return true
	}

	return false
}

// SetTokenReference gets a reference to the given string and assigns it to the TokenReference field.
func (o *SecureRemoteCommerceCheckoutData) SetTokenReference(v string) {
	o.TokenReference = &v
}

func (o SecureRemoteCommerceCheckoutData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecureRemoteCommerceCheckoutData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutPayload) {
		toSerialize["checkoutPayload"] = o.CheckoutPayload
	}
	if !common.IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !common.IsNil(o.Cvc) {
		toSerialize["cvc"] = o.Cvc
	}
	if !common.IsNil(o.DigitalCardId) {
		toSerialize["digitalCardId"] = o.DigitalCardId
	}
	if !common.IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !common.IsNil(o.TokenReference) {
		toSerialize["tokenReference"] = o.TokenReference
	}
	return toSerialize, nil
}

type NullableSecureRemoteCommerceCheckoutData struct {
	value *SecureRemoteCommerceCheckoutData
	isSet bool
}

func (v NullableSecureRemoteCommerceCheckoutData) Get() *SecureRemoteCommerceCheckoutData {
	return v.value
}

func (v *NullableSecureRemoteCommerceCheckoutData) Set(val *SecureRemoteCommerceCheckoutData) {
	v.value = val
	v.isSet = true
}

func (v NullableSecureRemoteCommerceCheckoutData) IsSet() bool {
	return v.isSet
}

func (v *NullableSecureRemoteCommerceCheckoutData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecureRemoteCommerceCheckoutData(val *SecureRemoteCommerceCheckoutData) *NullableSecureRemoteCommerceCheckoutData {
	return &NullableSecureRemoteCommerceCheckoutData{value: val, isSet: true}
}

func (v NullableSecureRemoteCommerceCheckoutData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecureRemoteCommerceCheckoutData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SecureRemoteCommerceCheckoutData) isValidScheme() bool {
	var allowedEnumValues = []string{"mc", "visa"}
	for _, allowed := range allowedEnumValues {
		if o.GetScheme() == allowed {
			return true
		}
	}
	return false
}
