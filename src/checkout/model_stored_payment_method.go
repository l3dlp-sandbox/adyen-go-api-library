/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the StoredPaymentMethod type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredPaymentMethod{}

// StoredPaymentMethod struct for StoredPaymentMethod
type StoredPaymentMethod struct {
	// The brand of the card.
	Brand *string `json:"brand,omitempty"`
	// The month the card expires.
	ExpiryMonth *string `json:"expiryMonth,omitempty"`
	// The last two digits of the year the card expires. For example, **22** for the year 2022.
	ExpiryYear *string `json:"expiryYear,omitempty"`
	// The unique payment method code.
	HolderName *string `json:"holderName,omitempty"`
	// The IBAN of the bank account.
	Iban *string `json:"iban,omitempty"`
	// A unique identifier of this stored payment method.
	Id *string `json:"id,omitempty"`
	// The last four digits of the PAN.
	LastFour *string `json:"lastFour,omitempty"`
	// The display name of the stored payment method.
	Name *string `json:"name,omitempty"`
	// Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID.
	NetworkTxReference *string `json:"networkTxReference,omitempty"`
	// The name of the bank account holder.
	OwnerName *string `json:"ownerName,omitempty"`
	// The shopper’s email address.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	// The supported recurring processing models for this stored payment method.
	SupportedRecurringProcessingModels []string `json:"supportedRecurringProcessingModels,omitempty"`
	// The supported shopper interactions for this stored payment method.
	SupportedShopperInteractions []string `json:"supportedShopperInteractions,omitempty"`
	// The type of payment method.
	Type *string `json:"type,omitempty"`
}

// NewStoredPaymentMethod instantiates a new StoredPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredPaymentMethod() *StoredPaymentMethod {
	this := StoredPaymentMethod{}
	return &this
}

// NewStoredPaymentMethodWithDefaults instantiates a new StoredPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredPaymentMethodWithDefaults() *StoredPaymentMethod {
	this := StoredPaymentMethod{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *StoredPaymentMethod) SetBrand(v string) {
	o.Brand = &v
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetExpiryMonth() string {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetExpiryMonthOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasExpiryMonth() bool {
	if o != nil && !common.IsNil(o.ExpiryMonth) {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *StoredPaymentMethod) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetExpiryYear() string {
	if o == nil || common.IsNil(o.ExpiryYear) {
		var ret string
		return ret
	}
	return *o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetExpiryYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryYear) {
		return nil, false
	}
	return o.ExpiryYear, true
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasExpiryYear() bool {
	if o != nil && !common.IsNil(o.ExpiryYear) {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given string and assigns it to the ExpiryYear field.
func (o *StoredPaymentMethod) SetExpiryYear(v string) {
	o.ExpiryYear = &v
}

// GetHolderName returns the HolderName field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetHolderName() string {
	if o == nil || common.IsNil(o.HolderName) {
		var ret string
		return ret
	}
	return *o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetHolderNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.HolderName) {
		return nil, false
	}
	return o.HolderName, true
}

// HasHolderName returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasHolderName() bool {
	if o != nil && !common.IsNil(o.HolderName) {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given string and assigns it to the HolderName field.
func (o *StoredPaymentMethod) SetHolderName(v string) {
	o.HolderName = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetIban() string {
	if o == nil || common.IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetIbanOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasIban() bool {
	if o != nil && !common.IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *StoredPaymentMethod) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StoredPaymentMethod) SetId(v string) {
	o.Id = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetLastFour() string {
	if o == nil || common.IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetLastFourOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasLastFour() bool {
	if o != nil && !common.IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *StoredPaymentMethod) SetLastFour(v string) {
	o.LastFour = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StoredPaymentMethod) SetName(v string) {
	o.Name = &v
}

// GetNetworkTxReference returns the NetworkTxReference field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetNetworkTxReference() string {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		var ret string
		return ret
	}
	return *o.NetworkTxReference
}

// GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetNetworkTxReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		return nil, false
	}
	return o.NetworkTxReference, true
}

// HasNetworkTxReference returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasNetworkTxReference() bool {
	if o != nil && !common.IsNil(o.NetworkTxReference) {
		return true
	}

	return false
}

// SetNetworkTxReference gets a reference to the given string and assigns it to the NetworkTxReference field.
func (o *StoredPaymentMethod) SetNetworkTxReference(v string) {
	o.NetworkTxReference = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetOwnerName() string {
	if o == nil || common.IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetOwnerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasOwnerName() bool {
	if o != nil && !common.IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *StoredPaymentMethod) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *StoredPaymentMethod) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetSupportedRecurringProcessingModels returns the SupportedRecurringProcessingModels field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetSupportedRecurringProcessingModels() []string {
	if o == nil || common.IsNil(o.SupportedRecurringProcessingModels) {
		var ret []string
		return ret
	}
	return o.SupportedRecurringProcessingModels
}

// GetSupportedRecurringProcessingModelsOk returns a tuple with the SupportedRecurringProcessingModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetSupportedRecurringProcessingModelsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SupportedRecurringProcessingModels) {
		return nil, false
	}
	return o.SupportedRecurringProcessingModels, true
}

// HasSupportedRecurringProcessingModels returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasSupportedRecurringProcessingModels() bool {
	if o != nil && !common.IsNil(o.SupportedRecurringProcessingModels) {
		return true
	}

	return false
}

// SetSupportedRecurringProcessingModels gets a reference to the given []string and assigns it to the SupportedRecurringProcessingModels field.
func (o *StoredPaymentMethod) SetSupportedRecurringProcessingModels(v []string) {
	o.SupportedRecurringProcessingModels = v
}

// GetSupportedShopperInteractions returns the SupportedShopperInteractions field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetSupportedShopperInteractions() []string {
	if o == nil || common.IsNil(o.SupportedShopperInteractions) {
		var ret []string
		return ret
	}
	return o.SupportedShopperInteractions
}

// GetSupportedShopperInteractionsOk returns a tuple with the SupportedShopperInteractions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetSupportedShopperInteractionsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SupportedShopperInteractions) {
		return nil, false
	}
	return o.SupportedShopperInteractions, true
}

// HasSupportedShopperInteractions returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasSupportedShopperInteractions() bool {
	if o != nil && !common.IsNil(o.SupportedShopperInteractions) {
		return true
	}

	return false
}

// SetSupportedShopperInteractions gets a reference to the given []string and assigns it to the SupportedShopperInteractions field.
func (o *StoredPaymentMethod) SetSupportedShopperInteractions(v []string) {
	o.SupportedShopperInteractions = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StoredPaymentMethod) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StoredPaymentMethod) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StoredPaymentMethod) SetType(v string) {
	o.Type = &v
}

func (o StoredPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !common.IsNil(o.ExpiryMonth) {
		toSerialize["expiryMonth"] = o.ExpiryMonth
	}
	if !common.IsNil(o.ExpiryYear) {
		toSerialize["expiryYear"] = o.ExpiryYear
	}
	if !common.IsNil(o.HolderName) {
		toSerialize["holderName"] = o.HolderName
	}
	if !common.IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.LastFour) {
		toSerialize["lastFour"] = o.LastFour
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.NetworkTxReference) {
		toSerialize["networkTxReference"] = o.NetworkTxReference
	}
	if !common.IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	if !common.IsNil(o.SupportedRecurringProcessingModels) {
		toSerialize["supportedRecurringProcessingModels"] = o.SupportedRecurringProcessingModels
	}
	if !common.IsNil(o.SupportedShopperInteractions) {
		toSerialize["supportedShopperInteractions"] = o.SupportedShopperInteractions
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableStoredPaymentMethod struct {
	value *StoredPaymentMethod
	isSet bool
}

func (v NullableStoredPaymentMethod) Get() *StoredPaymentMethod {
	return v.value
}

func (v *NullableStoredPaymentMethod) Set(val *StoredPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredPaymentMethod(val *StoredPaymentMethod) *NullableStoredPaymentMethod {
	return &NullableStoredPaymentMethod{value: val, isSet: true}
}

func (v NullableStoredPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
