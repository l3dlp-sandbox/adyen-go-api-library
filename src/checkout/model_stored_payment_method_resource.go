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

// checks if the StoredPaymentMethodResource type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredPaymentMethodResource{}

// StoredPaymentMethodResource struct for StoredPaymentMethodResource
type StoredPaymentMethodResource struct {
	// The brand of the card.
	Brand *string `json:"brand,omitempty"`
	// The month the card expires.
	ExpiryMonth *string `json:"expiryMonth,omitempty"`
	// The last two digits of the year the card expires. For example, **22** for the year 2022.
	ExpiryYear *string `json:"expiryYear,omitempty"`
	// The response code returned by an external system (for example after a provisioning operation).
	ExternalResponseCode *string `json:"externalResponseCode,omitempty"`
	// The token reference of a linked token in an external system (for example a network token reference).
	ExternalTokenReference *string `json:"externalTokenReference,omitempty"`
	// The unique payment method code.
	HolderName *string `json:"holderName,omitempty"`
	// The IBAN of the bank account.
	Iban *string `json:"iban,omitempty"`
	// A unique identifier of this stored payment method.
	Id *string `json:"id,omitempty"`
	// The name of the issuer of token or card.
	IssuerName *string `json:"issuerName,omitempty"`
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
	// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference *string `json:"shopperReference,omitempty"`
	// Defines a recurring payment type. Allowed values: * `Subscription` – A transaction for a fixed or variable amount, which follows a fixed schedule. * `CardOnFile` – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * `UnscheduledCardOnFile` – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder's balance drops below a certain amount.
	SupportedRecurringProcessingModels []string `json:"supportedRecurringProcessingModels,omitempty"`
	// The type of payment method.
	Type *string `json:"type,omitempty"`
}

// NewStoredPaymentMethodResource instantiates a new StoredPaymentMethodResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredPaymentMethodResource() *StoredPaymentMethodResource {
	this := StoredPaymentMethodResource{}
	return &this
}

// NewStoredPaymentMethodResourceWithDefaults instantiates a new StoredPaymentMethodResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredPaymentMethodResourceWithDefaults() *StoredPaymentMethodResource {
	this := StoredPaymentMethodResource{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *StoredPaymentMethodResource) SetBrand(v string) {
	o.Brand = &v
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetExpiryMonth() string {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetExpiryMonthOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasExpiryMonth() bool {
	if o != nil && !common.IsNil(o.ExpiryMonth) {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *StoredPaymentMethodResource) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetExpiryYear() string {
	if o == nil || common.IsNil(o.ExpiryYear) {
		var ret string
		return ret
	}
	return *o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetExpiryYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryYear) {
		return nil, false
	}
	return o.ExpiryYear, true
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasExpiryYear() bool {
	if o != nil && !common.IsNil(o.ExpiryYear) {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given string and assigns it to the ExpiryYear field.
func (o *StoredPaymentMethodResource) SetExpiryYear(v string) {
	o.ExpiryYear = &v
}

// GetExternalResponseCode returns the ExternalResponseCode field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetExternalResponseCode() string {
	if o == nil || common.IsNil(o.ExternalResponseCode) {
		var ret string
		return ret
	}
	return *o.ExternalResponseCode
}

// GetExternalResponseCodeOk returns a tuple with the ExternalResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetExternalResponseCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExternalResponseCode) {
		return nil, false
	}
	return o.ExternalResponseCode, true
}

// HasExternalResponseCode returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasExternalResponseCode() bool {
	if o != nil && !common.IsNil(o.ExternalResponseCode) {
		return true
	}

	return false
}

// SetExternalResponseCode gets a reference to the given string and assigns it to the ExternalResponseCode field.
func (o *StoredPaymentMethodResource) SetExternalResponseCode(v string) {
	o.ExternalResponseCode = &v
}

// GetExternalTokenReference returns the ExternalTokenReference field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetExternalTokenReference() string {
	if o == nil || common.IsNil(o.ExternalTokenReference) {
		var ret string
		return ret
	}
	return *o.ExternalTokenReference
}

// GetExternalTokenReferenceOk returns a tuple with the ExternalTokenReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetExternalTokenReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExternalTokenReference) {
		return nil, false
	}
	return o.ExternalTokenReference, true
}

// HasExternalTokenReference returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasExternalTokenReference() bool {
	if o != nil && !common.IsNil(o.ExternalTokenReference) {
		return true
	}

	return false
}

// SetExternalTokenReference gets a reference to the given string and assigns it to the ExternalTokenReference field.
func (o *StoredPaymentMethodResource) SetExternalTokenReference(v string) {
	o.ExternalTokenReference = &v
}

// GetHolderName returns the HolderName field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetHolderName() string {
	if o == nil || common.IsNil(o.HolderName) {
		var ret string
		return ret
	}
	return *o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetHolderNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.HolderName) {
		return nil, false
	}
	return o.HolderName, true
}

// HasHolderName returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasHolderName() bool {
	if o != nil && !common.IsNil(o.HolderName) {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given string and assigns it to the HolderName field.
func (o *StoredPaymentMethodResource) SetHolderName(v string) {
	o.HolderName = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetIban() string {
	if o == nil || common.IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetIbanOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasIban() bool {
	if o != nil && !common.IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *StoredPaymentMethodResource) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StoredPaymentMethodResource) SetId(v string) {
	o.Id = &v
}

// GetIssuerName returns the IssuerName field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetIssuerName() string {
	if o == nil || common.IsNil(o.IssuerName) {
		var ret string
		return ret
	}
	return *o.IssuerName
}

// GetIssuerNameOk returns a tuple with the IssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetIssuerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerName) {
		return nil, false
	}
	return o.IssuerName, true
}

// HasIssuerName returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasIssuerName() bool {
	if o != nil && !common.IsNil(o.IssuerName) {
		return true
	}

	return false
}

// SetIssuerName gets a reference to the given string and assigns it to the IssuerName field.
func (o *StoredPaymentMethodResource) SetIssuerName(v string) {
	o.IssuerName = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetLastFour() string {
	if o == nil || common.IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetLastFourOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasLastFour() bool {
	if o != nil && !common.IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *StoredPaymentMethodResource) SetLastFour(v string) {
	o.LastFour = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StoredPaymentMethodResource) SetName(v string) {
	o.Name = &v
}

// GetNetworkTxReference returns the NetworkTxReference field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetNetworkTxReference() string {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		var ret string
		return ret
	}
	return *o.NetworkTxReference
}

// GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetNetworkTxReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		return nil, false
	}
	return o.NetworkTxReference, true
}

// HasNetworkTxReference returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasNetworkTxReference() bool {
	if o != nil && !common.IsNil(o.NetworkTxReference) {
		return true
	}

	return false
}

// SetNetworkTxReference gets a reference to the given string and assigns it to the NetworkTxReference field.
func (o *StoredPaymentMethodResource) SetNetworkTxReference(v string) {
	o.NetworkTxReference = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetOwnerName() string {
	if o == nil || common.IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetOwnerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasOwnerName() bool {
	if o != nil && !common.IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *StoredPaymentMethodResource) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *StoredPaymentMethodResource) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *StoredPaymentMethodResource) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetSupportedRecurringProcessingModels returns the SupportedRecurringProcessingModels field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetSupportedRecurringProcessingModels() []string {
	if o == nil || common.IsNil(o.SupportedRecurringProcessingModels) {
		var ret []string
		return ret
	}
	return o.SupportedRecurringProcessingModels
}

// GetSupportedRecurringProcessingModelsOk returns a tuple with the SupportedRecurringProcessingModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetSupportedRecurringProcessingModelsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SupportedRecurringProcessingModels) {
		return nil, false
	}
	return o.SupportedRecurringProcessingModels, true
}

// HasSupportedRecurringProcessingModels returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasSupportedRecurringProcessingModels() bool {
	if o != nil && !common.IsNil(o.SupportedRecurringProcessingModels) {
		return true
	}

	return false
}

// SetSupportedRecurringProcessingModels gets a reference to the given []string and assigns it to the SupportedRecurringProcessingModels field.
func (o *StoredPaymentMethodResource) SetSupportedRecurringProcessingModels(v []string) {
	o.SupportedRecurringProcessingModels = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StoredPaymentMethodResource) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredPaymentMethodResource) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StoredPaymentMethodResource) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StoredPaymentMethodResource) SetType(v string) {
	o.Type = &v
}

func (o StoredPaymentMethodResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredPaymentMethodResource) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.ExternalResponseCode) {
		toSerialize["externalResponseCode"] = o.ExternalResponseCode
	}
	if !common.IsNil(o.ExternalTokenReference) {
		toSerialize["externalTokenReference"] = o.ExternalTokenReference
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
	if !common.IsNil(o.IssuerName) {
		toSerialize["issuerName"] = o.IssuerName
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
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	if !common.IsNil(o.SupportedRecurringProcessingModels) {
		toSerialize["supportedRecurringProcessingModels"] = o.SupportedRecurringProcessingModels
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableStoredPaymentMethodResource struct {
	value *StoredPaymentMethodResource
	isSet bool
}

func (v NullableStoredPaymentMethodResource) Get() *StoredPaymentMethodResource {
	return v.value
}

func (v *NullableStoredPaymentMethodResource) Set(val *StoredPaymentMethodResource) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredPaymentMethodResource) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredPaymentMethodResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredPaymentMethodResource(val *StoredPaymentMethodResource) *NullableStoredPaymentMethodResource {
	return &NullableStoredPaymentMethodResource{value: val, isSet: true}
}

func (v NullableStoredPaymentMethodResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredPaymentMethodResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
