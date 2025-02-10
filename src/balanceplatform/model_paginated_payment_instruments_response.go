/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the PaginatedPaymentInstrumentsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaginatedPaymentInstrumentsResponse{}

// PaginatedPaymentInstrumentsResponse struct for PaginatedPaymentInstrumentsResponse
type PaginatedPaymentInstrumentsResponse struct {
	// Indicates whether there are more items on the next page.
	HasNext bool `json:"hasNext"`
	// Indicates whether there are more items on the previous page.
	HasPrevious bool `json:"hasPrevious"`
	// List of payment instruments associated with the balance account.
	PaymentInstruments []PaymentInstrument `json:"paymentInstruments"`
}

// NewPaginatedPaymentInstrumentsResponse instantiates a new PaginatedPaymentInstrumentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPaymentInstrumentsResponse(hasNext bool, hasPrevious bool, paymentInstruments []PaymentInstrument) *PaginatedPaymentInstrumentsResponse {
	this := PaginatedPaymentInstrumentsResponse{}
	this.HasNext = hasNext
	this.HasPrevious = hasPrevious
	this.PaymentInstruments = paymentInstruments
	return &this
}

// NewPaginatedPaymentInstrumentsResponseWithDefaults instantiates a new PaginatedPaymentInstrumentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPaymentInstrumentsResponseWithDefaults() *PaginatedPaymentInstrumentsResponse {
	this := PaginatedPaymentInstrumentsResponse{}
	return &this
}

// GetHasNext returns the HasNext field value
func (o *PaginatedPaymentInstrumentsResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *PaginatedPaymentInstrumentsResponse) GetHasNextOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *PaginatedPaymentInstrumentsResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrevious returns the HasPrevious field value
func (o *PaginatedPaymentInstrumentsResponse) GetHasPrevious() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrevious
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value
// and a boolean to check if the value has been set.
func (o *PaginatedPaymentInstrumentsResponse) GetHasPreviousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPrevious, true
}

// SetHasPrevious sets field value
func (o *PaginatedPaymentInstrumentsResponse) SetHasPrevious(v bool) {
	o.HasPrevious = v
}

// GetPaymentInstruments returns the PaymentInstruments field value
func (o *PaginatedPaymentInstrumentsResponse) GetPaymentInstruments() []PaymentInstrument {
	if o == nil {
		var ret []PaymentInstrument
		return ret
	}

	return o.PaymentInstruments
}

// GetPaymentInstrumentsOk returns a tuple with the PaymentInstruments field value
// and a boolean to check if the value has been set.
func (o *PaginatedPaymentInstrumentsResponse) GetPaymentInstrumentsOk() ([]PaymentInstrument, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentInstruments, true
}

// SetPaymentInstruments sets field value
func (o *PaginatedPaymentInstrumentsResponse) SetPaymentInstruments(v []PaymentInstrument) {
	o.PaymentInstruments = v
}

func (o PaginatedPaymentInstrumentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedPaymentInstrumentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasNext"] = o.HasNext
	toSerialize["hasPrevious"] = o.HasPrevious
	toSerialize["paymentInstruments"] = o.PaymentInstruments
	return toSerialize, nil
}

type NullablePaginatedPaymentInstrumentsResponse struct {
	value *PaginatedPaymentInstrumentsResponse
	isSet bool
}

func (v NullablePaginatedPaymentInstrumentsResponse) Get() *PaginatedPaymentInstrumentsResponse {
	return v.value
}

func (v *NullablePaginatedPaymentInstrumentsResponse) Set(val *PaginatedPaymentInstrumentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPaymentInstrumentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPaymentInstrumentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPaymentInstrumentsResponse(val *PaginatedPaymentInstrumentsResponse) *NullablePaginatedPaymentInstrumentsResponse {
	return &NullablePaginatedPaymentInstrumentsResponse{value: val, isSet: true}
}

func (v NullablePaginatedPaymentInstrumentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPaymentInstrumentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
