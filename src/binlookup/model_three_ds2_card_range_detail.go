/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the ThreeDS2CardRangeDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDS2CardRangeDetail{}

// ThreeDS2CardRangeDetail struct for ThreeDS2CardRangeDetail
type ThreeDS2CardRangeDetail struct {
	// Provides additional information to the 3DS Server. Possible values: - 01 (Authentication is available at ACS) - 02 (Attempts supported by ACS or DS) - 03 (Decoupled authentication supported) - 04 (Whitelisting supported)
	AcsInfoInd []string `json:"acsInfoInd,omitempty"`
	// Card brand.
	BrandCode *string `json:"brandCode,omitempty"`
	// BIN end range.
	EndRange *string `json:"endRange,omitempty"`
	// BIN start range.
	StartRange *string `json:"startRange,omitempty"`
	// Supported 3D Secure protocol versions
	ThreeDS2Versions []string `json:"threeDS2Versions,omitempty"`
	// In a 3D Secure 2 browser-based flow, this is the URL where you should send the device fingerprint to.
	ThreeDSMethodURL *string `json:"threeDSMethodURL,omitempty"`
}

// NewThreeDS2CardRangeDetail instantiates a new ThreeDS2CardRangeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDS2CardRangeDetail() *ThreeDS2CardRangeDetail {
	this := ThreeDS2CardRangeDetail{}
	return &this
}

// NewThreeDS2CardRangeDetailWithDefaults instantiates a new ThreeDS2CardRangeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDS2CardRangeDetailWithDefaults() *ThreeDS2CardRangeDetail {
	this := ThreeDS2CardRangeDetail{}
	return &this
}

// GetAcsInfoInd returns the AcsInfoInd field value if set, zero value otherwise.
func (o *ThreeDS2CardRangeDetail) GetAcsInfoInd() []string {
	if o == nil || common.IsNil(o.AcsInfoInd) {
		var ret []string
		return ret
	}
	return o.AcsInfoInd
}

// GetAcsInfoIndOk returns a tuple with the AcsInfoInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2CardRangeDetail) GetAcsInfoIndOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AcsInfoInd) {
		return nil, false
	}
	return o.AcsInfoInd, true
}

// HasAcsInfoInd returns a boolean if a field has been set.
func (o *ThreeDS2CardRangeDetail) HasAcsInfoInd() bool {
	if o != nil && !common.IsNil(o.AcsInfoInd) {
		return true
	}

	return false
}

// SetAcsInfoInd gets a reference to the given []string and assigns it to the AcsInfoInd field.
func (o *ThreeDS2CardRangeDetail) SetAcsInfoInd(v []string) {
	o.AcsInfoInd = v
}

// GetBrandCode returns the BrandCode field value if set, zero value otherwise.
func (o *ThreeDS2CardRangeDetail) GetBrandCode() string {
	if o == nil || common.IsNil(o.BrandCode) {
		var ret string
		return ret
	}
	return *o.BrandCode
}

// GetBrandCodeOk returns a tuple with the BrandCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2CardRangeDetail) GetBrandCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BrandCode) {
		return nil, false
	}
	return o.BrandCode, true
}

// HasBrandCode returns a boolean if a field has been set.
func (o *ThreeDS2CardRangeDetail) HasBrandCode() bool {
	if o != nil && !common.IsNil(o.BrandCode) {
		return true
	}

	return false
}

// SetBrandCode gets a reference to the given string and assigns it to the BrandCode field.
func (o *ThreeDS2CardRangeDetail) SetBrandCode(v string) {
	o.BrandCode = &v
}

// GetEndRange returns the EndRange field value if set, zero value otherwise.
func (o *ThreeDS2CardRangeDetail) GetEndRange() string {
	if o == nil || common.IsNil(o.EndRange) {
		var ret string
		return ret
	}
	return *o.EndRange
}

// GetEndRangeOk returns a tuple with the EndRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2CardRangeDetail) GetEndRangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EndRange) {
		return nil, false
	}
	return o.EndRange, true
}

// HasEndRange returns a boolean if a field has been set.
func (o *ThreeDS2CardRangeDetail) HasEndRange() bool {
	if o != nil && !common.IsNil(o.EndRange) {
		return true
	}

	return false
}

// SetEndRange gets a reference to the given string and assigns it to the EndRange field.
func (o *ThreeDS2CardRangeDetail) SetEndRange(v string) {
	o.EndRange = &v
}

// GetStartRange returns the StartRange field value if set, zero value otherwise.
func (o *ThreeDS2CardRangeDetail) GetStartRange() string {
	if o == nil || common.IsNil(o.StartRange) {
		var ret string
		return ret
	}
	return *o.StartRange
}

// GetStartRangeOk returns a tuple with the StartRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2CardRangeDetail) GetStartRangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StartRange) {
		return nil, false
	}
	return o.StartRange, true
}

// HasStartRange returns a boolean if a field has been set.
func (o *ThreeDS2CardRangeDetail) HasStartRange() bool {
	if o != nil && !common.IsNil(o.StartRange) {
		return true
	}

	return false
}

// SetStartRange gets a reference to the given string and assigns it to the StartRange field.
func (o *ThreeDS2CardRangeDetail) SetStartRange(v string) {
	o.StartRange = &v
}

// GetThreeDS2Versions returns the ThreeDS2Versions field value if set, zero value otherwise.
func (o *ThreeDS2CardRangeDetail) GetThreeDS2Versions() []string {
	if o == nil || common.IsNil(o.ThreeDS2Versions) {
		var ret []string
		return ret
	}
	return o.ThreeDS2Versions
}

// GetThreeDS2VersionsOk returns a tuple with the ThreeDS2Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2CardRangeDetail) GetThreeDS2VersionsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.ThreeDS2Versions) {
		return nil, false
	}
	return o.ThreeDS2Versions, true
}

// HasThreeDS2Versions returns a boolean if a field has been set.
func (o *ThreeDS2CardRangeDetail) HasThreeDS2Versions() bool {
	if o != nil && !common.IsNil(o.ThreeDS2Versions) {
		return true
	}

	return false
}

// SetThreeDS2Versions gets a reference to the given []string and assigns it to the ThreeDS2Versions field.
func (o *ThreeDS2CardRangeDetail) SetThreeDS2Versions(v []string) {
	o.ThreeDS2Versions = v
}

// GetThreeDSMethodURL returns the ThreeDSMethodURL field value if set, zero value otherwise.
func (o *ThreeDS2CardRangeDetail) GetThreeDSMethodURL() string {
	if o == nil || common.IsNil(o.ThreeDSMethodURL) {
		var ret string
		return ret
	}
	return *o.ThreeDSMethodURL
}

// GetThreeDSMethodURLOk returns a tuple with the ThreeDSMethodURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2CardRangeDetail) GetThreeDSMethodURLOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSMethodURL) {
		return nil, false
	}
	return o.ThreeDSMethodURL, true
}

// HasThreeDSMethodURL returns a boolean if a field has been set.
func (o *ThreeDS2CardRangeDetail) HasThreeDSMethodURL() bool {
	if o != nil && !common.IsNil(o.ThreeDSMethodURL) {
		return true
	}

	return false
}

// SetThreeDSMethodURL gets a reference to the given string and assigns it to the ThreeDSMethodURL field.
func (o *ThreeDS2CardRangeDetail) SetThreeDSMethodURL(v string) {
	o.ThreeDSMethodURL = &v
}

func (o ThreeDS2CardRangeDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDS2CardRangeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcsInfoInd) {
		toSerialize["acsInfoInd"] = o.AcsInfoInd
	}
	if !common.IsNil(o.BrandCode) {
		toSerialize["brandCode"] = o.BrandCode
	}
	if !common.IsNil(o.EndRange) {
		toSerialize["endRange"] = o.EndRange
	}
	if !common.IsNil(o.StartRange) {
		toSerialize["startRange"] = o.StartRange
	}
	if !common.IsNil(o.ThreeDS2Versions) {
		toSerialize["threeDS2Versions"] = o.ThreeDS2Versions
	}
	if !common.IsNil(o.ThreeDSMethodURL) {
		toSerialize["threeDSMethodURL"] = o.ThreeDSMethodURL
	}
	return toSerialize, nil
}

type NullableThreeDS2CardRangeDetail struct {
	value *ThreeDS2CardRangeDetail
	isSet bool
}

func (v NullableThreeDS2CardRangeDetail) Get() *ThreeDS2CardRangeDetail {
	return v.value
}

func (v *NullableThreeDS2CardRangeDetail) Set(val *ThreeDS2CardRangeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDS2CardRangeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDS2CardRangeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDS2CardRangeDetail(val *ThreeDS2CardRangeDetail) *NullableThreeDS2CardRangeDetail {
	return &NullableThreeDS2CardRangeDetail{value: val, isSet: true}
}

func (v NullableThreeDS2CardRangeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDS2CardRangeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
