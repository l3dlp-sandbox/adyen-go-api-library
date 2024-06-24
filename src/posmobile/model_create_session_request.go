/*
POS Mobile API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posmobile

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the CreateSessionRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateSessionRequest{}

// CreateSessionRequest struct for CreateSessionRequest
type CreateSessionRequest struct {
	// The unique identifier of your merchant account.
	MerchantAccount string `json:"merchantAccount"`
	// The setup token provided by the POS Mobile SDK.  - When using the Android POS Mobile SDK, obtain the token through the `AuthenticationService.authenticate(setupToken)` callback of `AuthenticationService`.  - When using the iOS POS Mobile SDK, obtain the token through the `PaymentServiceDelegate.register(with:)` callback of `PaymentServiceDelegate`.
	SetupToken string `json:"setupToken"`
	// The unique identifier of the store that you want to process transactions for.
	Store *string `json:"store,omitempty"`
}

// NewCreateSessionRequest instantiates a new CreateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSessionRequest(merchantAccount string, setupToken string) *CreateSessionRequest {
	this := CreateSessionRequest{}
	this.MerchantAccount = merchantAccount
	this.SetupToken = setupToken
	return &this
}

// NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSessionRequestWithDefaults() *CreateSessionRequest {
	this := CreateSessionRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CreateSessionRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CreateSessionRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetSetupToken returns the SetupToken field value
func (o *CreateSessionRequest) GetSetupToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SetupToken
}

// GetSetupTokenOk returns a tuple with the SetupToken field value
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetSetupTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetupToken, true
}

// SetSetupToken sets field value
func (o *CreateSessionRequest) SetSetupToken(v string) {
	o.SetupToken = v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *CreateSessionRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *CreateSessionRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *CreateSessionRequest) SetStore(v string) {
	o.Store = &v
}

func (o CreateSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["setupToken"] = o.SetupToken
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	return toSerialize, nil
}

type NullableCreateSessionRequest struct {
	value *CreateSessionRequest
	isSet bool
}

func (v NullableCreateSessionRequest) Get() *CreateSessionRequest {
	return v.value
}

func (v *NullableCreateSessionRequest) Set(val *CreateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSessionRequest(val *CreateSessionRequest) *NullableCreateSessionRequest {
	return &NullableCreateSessionRequest{value: val, isSet: true}
}

func (v NullableCreateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
