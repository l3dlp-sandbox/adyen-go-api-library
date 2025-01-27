/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the ExternalPlatform type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExternalPlatform{}

// ExternalPlatform struct for ExternalPlatform
type ExternalPlatform struct {
	// External platform integrator.
	Integrator *string `json:"integrator,omitempty"`
	// Name of the field. For example, Name of External Platform.
	Name *string `json:"name,omitempty"`
	// Version of the field. For example, Version of External Platform.
	Version *string `json:"version,omitempty"`
}

// NewExternalPlatform instantiates a new ExternalPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPlatform() *ExternalPlatform {
	this := ExternalPlatform{}
	return &this
}

// NewExternalPlatformWithDefaults instantiates a new ExternalPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPlatformWithDefaults() *ExternalPlatform {
	this := ExternalPlatform{}
	return &this
}

// GetIntegrator returns the Integrator field value if set, zero value otherwise.
func (o *ExternalPlatform) GetIntegrator() string {
	if o == nil || common.IsNil(o.Integrator) {
		var ret string
		return ret
	}
	return *o.Integrator
}

// GetIntegratorOk returns a tuple with the Integrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPlatform) GetIntegratorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Integrator) {
		return nil, false
	}
	return o.Integrator, true
}

// HasIntegrator returns a boolean if a field has been set.
func (o *ExternalPlatform) HasIntegrator() bool {
	if o != nil && !common.IsNil(o.Integrator) {
		return true
	}

	return false
}

// SetIntegrator gets a reference to the given string and assigns it to the Integrator field.
func (o *ExternalPlatform) SetIntegrator(v string) {
	o.Integrator = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalPlatform) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPlatform) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalPlatform) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalPlatform) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ExternalPlatform) GetVersion() string {
	if o == nil || common.IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPlatform) GetVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ExternalPlatform) HasVersion() bool {
	if o != nil && !common.IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ExternalPlatform) SetVersion(v string) {
	o.Version = &v
}

func (o ExternalPlatform) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalPlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Integrator) {
		toSerialize["integrator"] = o.Integrator
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableExternalPlatform struct {
	value *ExternalPlatform
	isSet bool
}

func (v NullableExternalPlatform) Get() *ExternalPlatform {
	return v.value
}

func (v *NullableExternalPlatform) Set(val *ExternalPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPlatform(val *ExternalPlatform) *NullableExternalPlatform {
	return &NullableExternalPlatform{value: val, isSet: true}
}

func (v NullableExternalPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
