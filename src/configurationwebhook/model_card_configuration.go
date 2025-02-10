/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the CardConfiguration type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardConfiguration{}

// CardConfiguration struct for CardConfiguration
type CardConfiguration struct {
	// Overrides the activation label design ID defined in the `configurationProfileId`. The activation label is attached to the card and contains the activation instructions.
	Activation *string `json:"activation,omitempty"`
	// Your app's URL, if you want to activate cards through your app. For example, **my-app://ref1236a7d**. A QR code is created based on this URL, and is included in the carrier. Before you use this field, reach out to your Adyen contact to set up the QR code process.   Maximum length: 255 characters.
	ActivationUrl *string      `json:"activationUrl,omitempty"`
	BulkAddress   *BulkAddress `json:"bulkAddress,omitempty"`
	// The ID of the card image. This is the image that will be printed on the full front of the card.
	CardImageId *string `json:"cardImageId,omitempty"`
	// Overrides the carrier design ID defined in the `configurationProfileId`. The carrier is the letter or packaging to which the card is attached.
	Carrier *string `json:"carrier,omitempty"`
	// The ID of the carrier image. This is the image that will printed on the letter to which the card is attached.
	CarrierImageId *string `json:"carrierImageId,omitempty"`
	// The ID of the card configuration profile that contains the settings of the card. For example, the envelope and PIN mailer designs or the logistics company handling the shipment. All the settings in the profile are applied to the card, unless you provide other fields to override them.  For example, send the `shipmentMethod` to override the logistics company defined in the card configuration profile.
	ConfigurationProfileId string `json:"configurationProfileId"`
	// The three-letter [ISO-4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the card. For example, **EUR**.
	Currency *string `json:"currency,omitempty"`
	// Overrides the envelope design ID defined in the `configurationProfileId`.
	Envelope *string `json:"envelope,omitempty"`
	// Overrides the insert design ID defined in the `configurationProfileId`. An insert is any additional material, such as marketing materials, that are shipped together with the card.
	Insert *string `json:"insert,omitempty"`
	// The two-letter [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code of the card. For example, **en**.
	Language *string `json:"language,omitempty"`
	// The ID of the logo image. This is the image that will be printed on the partial front of the card, such as a logo on the upper right corner.
	LogoImageId *string `json:"logoImageId,omitempty"`
	// Overrides the PIN mailer design ID defined in the `configurationProfileId`. The PIN mailer is the letter on which the PIN is printed.
	PinMailer *string `json:"pinMailer,omitempty"`
	// Overrides the logistics company defined in the `configurationProfileId`.
	ShipmentMethod *string `json:"shipmentMethod,omitempty"`
}

// NewCardConfiguration instantiates a new CardConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardConfiguration(configurationProfileId string) *CardConfiguration {
	this := CardConfiguration{}
	this.ConfigurationProfileId = configurationProfileId
	return &this
}

// NewCardConfigurationWithDefaults instantiates a new CardConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardConfigurationWithDefaults() *CardConfiguration {
	this := CardConfiguration{}
	return &this
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *CardConfiguration) GetActivation() string {
	if o == nil || common.IsNil(o.Activation) {
		var ret string
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetActivationOk() (*string, bool) {
	if o == nil || common.IsNil(o.Activation) {
		return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *CardConfiguration) HasActivation() bool {
	if o != nil && !common.IsNil(o.Activation) {
		return true
	}

	return false
}

// SetActivation gets a reference to the given string and assigns it to the Activation field.
func (o *CardConfiguration) SetActivation(v string) {
	o.Activation = &v
}

// GetActivationUrl returns the ActivationUrl field value if set, zero value otherwise.
func (o *CardConfiguration) GetActivationUrl() string {
	if o == nil || common.IsNil(o.ActivationUrl) {
		var ret string
		return ret
	}
	return *o.ActivationUrl
}

// GetActivationUrlOk returns a tuple with the ActivationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetActivationUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActivationUrl) {
		return nil, false
	}
	return o.ActivationUrl, true
}

// HasActivationUrl returns a boolean if a field has been set.
func (o *CardConfiguration) HasActivationUrl() bool {
	if o != nil && !common.IsNil(o.ActivationUrl) {
		return true
	}

	return false
}

// SetActivationUrl gets a reference to the given string and assigns it to the ActivationUrl field.
func (o *CardConfiguration) SetActivationUrl(v string) {
	o.ActivationUrl = &v
}

// GetBulkAddress returns the BulkAddress field value if set, zero value otherwise.
func (o *CardConfiguration) GetBulkAddress() BulkAddress {
	if o == nil || common.IsNil(o.BulkAddress) {
		var ret BulkAddress
		return ret
	}
	return *o.BulkAddress
}

// GetBulkAddressOk returns a tuple with the BulkAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetBulkAddressOk() (*BulkAddress, bool) {
	if o == nil || common.IsNil(o.BulkAddress) {
		return nil, false
	}
	return o.BulkAddress, true
}

// HasBulkAddress returns a boolean if a field has been set.
func (o *CardConfiguration) HasBulkAddress() bool {
	if o != nil && !common.IsNil(o.BulkAddress) {
		return true
	}

	return false
}

// SetBulkAddress gets a reference to the given BulkAddress and assigns it to the BulkAddress field.
func (o *CardConfiguration) SetBulkAddress(v BulkAddress) {
	o.BulkAddress = &v
}

// GetCardImageId returns the CardImageId field value if set, zero value otherwise.
func (o *CardConfiguration) GetCardImageId() string {
	if o == nil || common.IsNil(o.CardImageId) {
		var ret string
		return ret
	}
	return *o.CardImageId
}

// GetCardImageIdOk returns a tuple with the CardImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetCardImageIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardImageId) {
		return nil, false
	}
	return o.CardImageId, true
}

// HasCardImageId returns a boolean if a field has been set.
func (o *CardConfiguration) HasCardImageId() bool {
	if o != nil && !common.IsNil(o.CardImageId) {
		return true
	}

	return false
}

// SetCardImageId gets a reference to the given string and assigns it to the CardImageId field.
func (o *CardConfiguration) SetCardImageId(v string) {
	o.CardImageId = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *CardConfiguration) GetCarrier() string {
	if o == nil || common.IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetCarrierOk() (*string, bool) {
	if o == nil || common.IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *CardConfiguration) HasCarrier() bool {
	if o != nil && !common.IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *CardConfiguration) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCarrierImageId returns the CarrierImageId field value if set, zero value otherwise.
func (o *CardConfiguration) GetCarrierImageId() string {
	if o == nil || common.IsNil(o.CarrierImageId) {
		var ret string
		return ret
	}
	return *o.CarrierImageId
}

// GetCarrierImageIdOk returns a tuple with the CarrierImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetCarrierImageIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarrierImageId) {
		return nil, false
	}
	return o.CarrierImageId, true
}

// HasCarrierImageId returns a boolean if a field has been set.
func (o *CardConfiguration) HasCarrierImageId() bool {
	if o != nil && !common.IsNil(o.CarrierImageId) {
		return true
	}

	return false
}

// SetCarrierImageId gets a reference to the given string and assigns it to the CarrierImageId field.
func (o *CardConfiguration) SetCarrierImageId(v string) {
	o.CarrierImageId = &v
}

// GetConfigurationProfileId returns the ConfigurationProfileId field value
func (o *CardConfiguration) GetConfigurationProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationProfileId
}

// GetConfigurationProfileIdOk returns a tuple with the ConfigurationProfileId field value
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetConfigurationProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationProfileId, true
}

// SetConfigurationProfileId sets field value
func (o *CardConfiguration) SetConfigurationProfileId(v string) {
	o.ConfigurationProfileId = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CardConfiguration) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CardConfiguration) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CardConfiguration) SetCurrency(v string) {
	o.Currency = &v
}

// GetEnvelope returns the Envelope field value if set, zero value otherwise.
func (o *CardConfiguration) GetEnvelope() string {
	if o == nil || common.IsNil(o.Envelope) {
		var ret string
		return ret
	}
	return *o.Envelope
}

// GetEnvelopeOk returns a tuple with the Envelope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetEnvelopeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Envelope) {
		return nil, false
	}
	return o.Envelope, true
}

// HasEnvelope returns a boolean if a field has been set.
func (o *CardConfiguration) HasEnvelope() bool {
	if o != nil && !common.IsNil(o.Envelope) {
		return true
	}

	return false
}

// SetEnvelope gets a reference to the given string and assigns it to the Envelope field.
func (o *CardConfiguration) SetEnvelope(v string) {
	o.Envelope = &v
}

// GetInsert returns the Insert field value if set, zero value otherwise.
func (o *CardConfiguration) GetInsert() string {
	if o == nil || common.IsNil(o.Insert) {
		var ret string
		return ret
	}
	return *o.Insert
}

// GetInsertOk returns a tuple with the Insert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetInsertOk() (*string, bool) {
	if o == nil || common.IsNil(o.Insert) {
		return nil, false
	}
	return o.Insert, true
}

// HasInsert returns a boolean if a field has been set.
func (o *CardConfiguration) HasInsert() bool {
	if o != nil && !common.IsNil(o.Insert) {
		return true
	}

	return false
}

// SetInsert gets a reference to the given string and assigns it to the Insert field.
func (o *CardConfiguration) SetInsert(v string) {
	o.Insert = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *CardConfiguration) GetLanguage() string {
	if o == nil || common.IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetLanguageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *CardConfiguration) HasLanguage() bool {
	if o != nil && !common.IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *CardConfiguration) SetLanguage(v string) {
	o.Language = &v
}

// GetLogoImageId returns the LogoImageId field value if set, zero value otherwise.
func (o *CardConfiguration) GetLogoImageId() string {
	if o == nil || common.IsNil(o.LogoImageId) {
		var ret string
		return ret
	}
	return *o.LogoImageId
}

// GetLogoImageIdOk returns a tuple with the LogoImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetLogoImageIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LogoImageId) {
		return nil, false
	}
	return o.LogoImageId, true
}

// HasLogoImageId returns a boolean if a field has been set.
func (o *CardConfiguration) HasLogoImageId() bool {
	if o != nil && !common.IsNil(o.LogoImageId) {
		return true
	}

	return false
}

// SetLogoImageId gets a reference to the given string and assigns it to the LogoImageId field.
func (o *CardConfiguration) SetLogoImageId(v string) {
	o.LogoImageId = &v
}

// GetPinMailer returns the PinMailer field value if set, zero value otherwise.
func (o *CardConfiguration) GetPinMailer() string {
	if o == nil || common.IsNil(o.PinMailer) {
		var ret string
		return ret
	}
	return *o.PinMailer
}

// GetPinMailerOk returns a tuple with the PinMailer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetPinMailerOk() (*string, bool) {
	if o == nil || common.IsNil(o.PinMailer) {
		return nil, false
	}
	return o.PinMailer, true
}

// HasPinMailer returns a boolean if a field has been set.
func (o *CardConfiguration) HasPinMailer() bool {
	if o != nil && !common.IsNil(o.PinMailer) {
		return true
	}

	return false
}

// SetPinMailer gets a reference to the given string and assigns it to the PinMailer field.
func (o *CardConfiguration) SetPinMailer(v string) {
	o.PinMailer = &v
}

// GetShipmentMethod returns the ShipmentMethod field value if set, zero value otherwise.
func (o *CardConfiguration) GetShipmentMethod() string {
	if o == nil || common.IsNil(o.ShipmentMethod) {
		var ret string
		return ret
	}
	return *o.ShipmentMethod
}

// GetShipmentMethodOk returns a tuple with the ShipmentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardConfiguration) GetShipmentMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShipmentMethod) {
		return nil, false
	}
	return o.ShipmentMethod, true
}

// HasShipmentMethod returns a boolean if a field has been set.
func (o *CardConfiguration) HasShipmentMethod() bool {
	if o != nil && !common.IsNil(o.ShipmentMethod) {
		return true
	}

	return false
}

// SetShipmentMethod gets a reference to the given string and assigns it to the ShipmentMethod field.
func (o *CardConfiguration) SetShipmentMethod(v string) {
	o.ShipmentMethod = &v
}

func (o CardConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Activation) {
		toSerialize["activation"] = o.Activation
	}
	if !common.IsNil(o.ActivationUrl) {
		toSerialize["activationUrl"] = o.ActivationUrl
	}
	if !common.IsNil(o.BulkAddress) {
		toSerialize["bulkAddress"] = o.BulkAddress
	}
	if !common.IsNil(o.CardImageId) {
		toSerialize["cardImageId"] = o.CardImageId
	}
	if !common.IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !common.IsNil(o.CarrierImageId) {
		toSerialize["carrierImageId"] = o.CarrierImageId
	}
	toSerialize["configurationProfileId"] = o.ConfigurationProfileId
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.Envelope) {
		toSerialize["envelope"] = o.Envelope
	}
	if !common.IsNil(o.Insert) {
		toSerialize["insert"] = o.Insert
	}
	if !common.IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !common.IsNil(o.LogoImageId) {
		toSerialize["logoImageId"] = o.LogoImageId
	}
	if !common.IsNil(o.PinMailer) {
		toSerialize["pinMailer"] = o.PinMailer
	}
	if !common.IsNil(o.ShipmentMethod) {
		toSerialize["shipmentMethod"] = o.ShipmentMethod
	}
	return toSerialize, nil
}

type NullableCardConfiguration struct {
	value *CardConfiguration
	isSet bool
}

func (v NullableCardConfiguration) Get() *CardConfiguration {
	return v.value
}

func (v *NullableCardConfiguration) Set(val *CardConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCardConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCardConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardConfiguration(val *CardConfiguration) *NullableCardConfiguration {
	return &NullableCardConfiguration{value: val, isSet: true}
}

func (v NullableCardConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
