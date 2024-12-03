/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// checks if the TerminalConnectivityEthernet type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalConnectivityEthernet{}

// TerminalConnectivityEthernet struct for TerminalConnectivityEthernet
type TerminalConnectivityEthernet struct {
	// The terminal's ethernet IP address.
	IpAddress *string `json:"ipAddress,omitempty"`
	// The ethernet link negotiation that the terminal uses.
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The terminal's ethernet MAC address.
	MacAddress *string `json:"macAddress,omitempty"`
}

// NewTerminalConnectivityEthernet instantiates a new TerminalConnectivityEthernet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalConnectivityEthernet() *TerminalConnectivityEthernet {
	this := TerminalConnectivityEthernet{}
	return &this
}

// NewTerminalConnectivityEthernetWithDefaults instantiates a new TerminalConnectivityEthernet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalConnectivityEthernetWithDefaults() *TerminalConnectivityEthernet {
	this := TerminalConnectivityEthernet{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *TerminalConnectivityEthernet) GetIpAddress() string {
	if o == nil || common.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivityEthernet) GetIpAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *TerminalConnectivityEthernet) HasIpAddress() bool {
	if o != nil && !common.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *TerminalConnectivityEthernet) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *TerminalConnectivityEthernet) GetLinkNegotiation() string {
	if o == nil || common.IsNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivityEthernet) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || common.IsNil(o.LinkNegotiation) {
		return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *TerminalConnectivityEthernet) HasLinkNegotiation() bool {
	if o != nil && !common.IsNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *TerminalConnectivityEthernet) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *TerminalConnectivityEthernet) GetMacAddress() string {
	if o == nil || common.IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivityEthernet) GetMacAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *TerminalConnectivityEthernet) HasMacAddress() bool {
	if o != nil && !common.IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *TerminalConnectivityEthernet) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o TerminalConnectivityEthernet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalConnectivityEthernet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !common.IsNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	if !common.IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	return toSerialize, nil
}

type NullableTerminalConnectivityEthernet struct {
	value *TerminalConnectivityEthernet
	isSet bool
}

func (v NullableTerminalConnectivityEthernet) Get() *TerminalConnectivityEthernet {
	return v.value
}

func (v *NullableTerminalConnectivityEthernet) Set(val *TerminalConnectivityEthernet) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalConnectivityEthernet) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalConnectivityEthernet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalConnectivityEthernet(val *TerminalConnectivityEthernet) *NullableTerminalConnectivityEthernet {
	return &NullableTerminalConnectivityEthernet{value: val, isSet: true}
}

func (v NullableTerminalConnectivityEthernet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalConnectivityEthernet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



