/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the UpdateStoreRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateStoreRequest{}

// UpdateStoreRequest struct for UpdateStoreRequest
type UpdateStoreRequest struct {
	Address *UpdatableAddress `json:"address,omitempty"`
	// The unique identifiers of the [business lines](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businesslines__resParam_id) that the store is associated with.
	BusinessLineIds []string `json:"businessLineIds,omitempty"`
	// The description of the store.
	Description *string `json:"description,omitempty"`
	// When using the Zip payment method: The location ID that Zip has assigned to your store.
	ExternalReferenceId *string                  `json:"externalReferenceId,omitempty"`
	SplitConfiguration  *StoreSplitConfiguration `json:"splitConfiguration,omitempty"`
	// The status of the store. Possible values are:  - **active**: This value is assigned automatically when a store is created.  - **inactive**: The maximum [transaction limits and number of Store-and-Forward transactions](https://docs.adyen.com/point-of-sale/determine-account-structure/configure-features#payment-features) for the store are set to 0. This blocks new transactions, but captures are still possible. - **closed**: The terminals of the store are reassigned to the merchant inventory, so they can't process payments.  You can change the status from **active** to **inactive**, and from **inactive** to **active** or **closed**.  Once **closed**, a store can't be reopened.
	Status *string `json:"status,omitempty"`
}

// NewUpdateStoreRequest instantiates a new UpdateStoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStoreRequest() *UpdateStoreRequest {
	this := UpdateStoreRequest{}
	return &this
}

// NewUpdateStoreRequestWithDefaults instantiates a new UpdateStoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStoreRequestWithDefaults() *UpdateStoreRequest {
	this := UpdateStoreRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateStoreRequest) GetAddress() UpdatableAddress {
	if o == nil || common.IsNil(o.Address) {
		var ret UpdatableAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStoreRequest) GetAddressOk() (*UpdatableAddress, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateStoreRequest) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given UpdatableAddress and assigns it to the Address field.
func (o *UpdateStoreRequest) SetAddress(v UpdatableAddress) {
	o.Address = &v
}

// GetBusinessLineIds returns the BusinessLineIds field value if set, zero value otherwise.
func (o *UpdateStoreRequest) GetBusinessLineIds() []string {
	if o == nil || common.IsNil(o.BusinessLineIds) {
		var ret []string
		return ret
	}
	return o.BusinessLineIds
}

// GetBusinessLineIdsOk returns a tuple with the BusinessLineIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStoreRequest) GetBusinessLineIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.BusinessLineIds) {
		return nil, false
	}
	return o.BusinessLineIds, true
}

// HasBusinessLineIds returns a boolean if a field has been set.
func (o *UpdateStoreRequest) HasBusinessLineIds() bool {
	if o != nil && !common.IsNil(o.BusinessLineIds) {
		return true
	}

	return false
}

// SetBusinessLineIds gets a reference to the given []string and assigns it to the BusinessLineIds field.
func (o *UpdateStoreRequest) SetBusinessLineIds(v []string) {
	o.BusinessLineIds = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateStoreRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStoreRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateStoreRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateStoreRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExternalReferenceId returns the ExternalReferenceId field value if set, zero value otherwise.
func (o *UpdateStoreRequest) GetExternalReferenceId() string {
	if o == nil || common.IsNil(o.ExternalReferenceId) {
		var ret string
		return ret
	}
	return *o.ExternalReferenceId
}

// GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStoreRequest) GetExternalReferenceIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExternalReferenceId) {
		return nil, false
	}
	return o.ExternalReferenceId, true
}

// HasExternalReferenceId returns a boolean if a field has been set.
func (o *UpdateStoreRequest) HasExternalReferenceId() bool {
	if o != nil && !common.IsNil(o.ExternalReferenceId) {
		return true
	}

	return false
}

// SetExternalReferenceId gets a reference to the given string and assigns it to the ExternalReferenceId field.
func (o *UpdateStoreRequest) SetExternalReferenceId(v string) {
	o.ExternalReferenceId = &v
}

// GetSplitConfiguration returns the SplitConfiguration field value if set, zero value otherwise.
func (o *UpdateStoreRequest) GetSplitConfiguration() StoreSplitConfiguration {
	if o == nil || common.IsNil(o.SplitConfiguration) {
		var ret StoreSplitConfiguration
		return ret
	}
	return *o.SplitConfiguration
}

// GetSplitConfigurationOk returns a tuple with the SplitConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStoreRequest) GetSplitConfigurationOk() (*StoreSplitConfiguration, bool) {
	if o == nil || common.IsNil(o.SplitConfiguration) {
		return nil, false
	}
	return o.SplitConfiguration, true
}

// HasSplitConfiguration returns a boolean if a field has been set.
func (o *UpdateStoreRequest) HasSplitConfiguration() bool {
	if o != nil && !common.IsNil(o.SplitConfiguration) {
		return true
	}

	return false
}

// SetSplitConfiguration gets a reference to the given StoreSplitConfiguration and assigns it to the SplitConfiguration field.
func (o *UpdateStoreRequest) SetSplitConfiguration(v StoreSplitConfiguration) {
	o.SplitConfiguration = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateStoreRequest) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStoreRequest) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateStoreRequest) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateStoreRequest) SetStatus(v string) {
	o.Status = &v
}

func (o UpdateStoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.BusinessLineIds) {
		toSerialize["businessLineIds"] = o.BusinessLineIds
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.ExternalReferenceId) {
		toSerialize["externalReferenceId"] = o.ExternalReferenceId
	}
	if !common.IsNil(o.SplitConfiguration) {
		toSerialize["splitConfiguration"] = o.SplitConfiguration
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUpdateStoreRequest struct {
	value *UpdateStoreRequest
	isSet bool
}

func (v NullableUpdateStoreRequest) Get() *UpdateStoreRequest {
	return v.value
}

func (v *NullableUpdateStoreRequest) Set(val *UpdateStoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStoreRequest(val *UpdateStoreRequest) *NullableUpdateStoreRequest {
	return &NullableUpdateStoreRequest{value: val, isSet: true}
}

func (v NullableUpdateStoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpdateStoreRequest) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "closed", "inactive"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}