/*
Payments App API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentsapp

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the BoardingTokenResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BoardingTokenResponse{}

// BoardingTokenResponse struct for BoardingTokenResponse
type BoardingTokenResponse struct {
	// The boarding token that allows the Payments App to board.
	BoardingToken string `json:"boardingToken"`
	// The unique identifier of the Payments App instance.
	InstallationId string `json:"installationId"`
}

// NewBoardingTokenResponse instantiates a new BoardingTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardingTokenResponse(boardingToken string, installationId string) *BoardingTokenResponse {
	this := BoardingTokenResponse{}
	this.BoardingToken = boardingToken
	this.InstallationId = installationId
	return &this
}

// NewBoardingTokenResponseWithDefaults instantiates a new BoardingTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardingTokenResponseWithDefaults() *BoardingTokenResponse {
	this := BoardingTokenResponse{}
	return &this
}

// GetBoardingToken returns the BoardingToken field value
func (o *BoardingTokenResponse) GetBoardingToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoardingToken
}

// GetBoardingTokenOk returns a tuple with the BoardingToken field value
// and a boolean to check if the value has been set.
func (o *BoardingTokenResponse) GetBoardingTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoardingToken, true
}

// SetBoardingToken sets field value
func (o *BoardingTokenResponse) SetBoardingToken(v string) {
	o.BoardingToken = v
}

// GetInstallationId returns the InstallationId field value
func (o *BoardingTokenResponse) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *BoardingTokenResponse) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *BoardingTokenResponse) SetInstallationId(v string) {
	o.InstallationId = v
}

func (o BoardingTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoardingTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boardingToken"] = o.BoardingToken
	toSerialize["installationId"] = o.InstallationId
	return toSerialize, nil
}

type NullableBoardingTokenResponse struct {
	value *BoardingTokenResponse
	isSet bool
}

func (v NullableBoardingTokenResponse) Get() *BoardingTokenResponse {
	return v.value
}

func (v *NullableBoardingTokenResponse) Set(val *BoardingTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardingTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardingTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardingTokenResponse(val *BoardingTokenResponse) *NullableBoardingTokenResponse {
	return &NullableBoardingTokenResponse{value: val, isSet: true}
}

func (v NullableBoardingTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardingTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
