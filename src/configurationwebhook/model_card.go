/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Card type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Card{}

// Card struct for Card
type Card struct {
	Authentication *Authentication `json:"authentication,omitempty"`
	// The bank identification number (BIN) of the card number.
	Bin *string `json:"bin,omitempty"`
	// The brand of the physical or the virtual card. Possible values: **visa**, **mc**.
	Brand string `json:"brand"`
	// The brand variant of the physical or the virtual card. For example, **visadebit** or **mcprepaid**. >Reach out to your Adyen contact to get the values relevant for your integration.
	BrandVariant string `json:"brandVariant"`
	// The name of the cardholder.  Maximum length: 26 characters.
	CardholderName string `json:"cardholderName"`
	Configuration *CardConfiguration `json:"configuration,omitempty"`
	// The CVC2 value of the card. > The CVC2 is not sent by default. This is only returned in the `POST` response for single-use virtual cards.
	Cvc *string `json:"cvc,omitempty"`
	DeliveryContact *DeliveryContact `json:"deliveryContact,omitempty"`
	Expiration *Expiry `json:"expiration,omitempty"`
	// The form factor of the card. Possible values: **virtual**, **physical**.
	FormFactor string `json:"formFactor"`
	// Last last four digits of the card number.
	LastFour *string `json:"lastFour,omitempty"`
	// The primary account number (PAN) of the card. > The PAN is masked by default and returned only for single-use virtual cards.
	Number string `json:"number"`
	// Allocates a specific product range for either a physical or a virtual card. Possible values: **fullySupported**, **secureCorporate**. >Reach out to your Adyen contact to get the values relevant for your integration.
	ThreeDSecure *string `json:"threeDSecure,omitempty"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard(brand string, brandVariant string, cardholderName string, formFactor string, number string) *Card {
	this := Card{}
	this.Brand = brand
	this.BrandVariant = brandVariant
	this.CardholderName = cardholderName
	this.FormFactor = formFactor
	this.Number = number
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *Card) GetAuthentication() Authentication {
	if o == nil || common.IsNil(o.Authentication) {
		var ret Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetAuthenticationOk() (*Authentication, bool) {
	if o == nil || common.IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *Card) HasAuthentication() bool {
	if o != nil && !common.IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given Authentication and assigns it to the Authentication field.
func (o *Card) SetAuthentication(v Authentication) {
	o.Authentication = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *Card) GetBin() string {
	if o == nil || common.IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *Card) HasBin() bool {
	if o != nil && !common.IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *Card) SetBin(v string) {
	o.Bin = &v
}

// GetBrand returns the Brand field value
func (o *Card) GetBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
func (o *Card) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brand, true
}

// SetBrand sets field value
func (o *Card) SetBrand(v string) {
	o.Brand = v
}

// GetBrandVariant returns the BrandVariant field value
func (o *Card) GetBrandVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandVariant
}

// GetBrandVariantOk returns a tuple with the BrandVariant field value
// and a boolean to check if the value has been set.
func (o *Card) GetBrandVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandVariant, true
}

// SetBrandVariant sets field value
func (o *Card) SetBrandVariant(v string) {
	o.BrandVariant = v
}

// GetCardholderName returns the CardholderName field value
func (o *Card) GetCardholderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value
// and a boolean to check if the value has been set.
func (o *Card) GetCardholderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardholderName, true
}

// SetCardholderName sets field value
func (o *Card) SetCardholderName(v string) {
	o.CardholderName = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *Card) GetConfiguration() CardConfiguration {
	if o == nil || common.IsNil(o.Configuration) {
		var ret CardConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetConfigurationOk() (*CardConfiguration, bool) {
	if o == nil || common.IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Card) HasConfiguration() bool {
	if o != nil && !common.IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given CardConfiguration and assigns it to the Configuration field.
func (o *Card) SetConfiguration(v CardConfiguration) {
	o.Configuration = &v
}

// GetCvc returns the Cvc field value if set, zero value otherwise.
func (o *Card) GetCvc() string {
	if o == nil || common.IsNil(o.Cvc) {
		var ret string
		return ret
	}
	return *o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCvcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cvc) {
		return nil, false
	}
	return o.Cvc, true
}

// HasCvc returns a boolean if a field has been set.
func (o *Card) HasCvc() bool {
	if o != nil && !common.IsNil(o.Cvc) {
		return true
	}

	return false
}

// SetCvc gets a reference to the given string and assigns it to the Cvc field.
func (o *Card) SetCvc(v string) {
	o.Cvc = &v
}

// GetDeliveryContact returns the DeliveryContact field value if set, zero value otherwise.
func (o *Card) GetDeliveryContact() DeliveryContact {
	if o == nil || common.IsNil(o.DeliveryContact) {
		var ret DeliveryContact
		return ret
	}
	return *o.DeliveryContact
}

// GetDeliveryContactOk returns a tuple with the DeliveryContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetDeliveryContactOk() (*DeliveryContact, bool) {
	if o == nil || common.IsNil(o.DeliveryContact) {
		return nil, false
	}
	return o.DeliveryContact, true
}

// HasDeliveryContact returns a boolean if a field has been set.
func (o *Card) HasDeliveryContact() bool {
	if o != nil && !common.IsNil(o.DeliveryContact) {
		return true
	}

	return false
}

// SetDeliveryContact gets a reference to the given DeliveryContact and assigns it to the DeliveryContact field.
func (o *Card) SetDeliveryContact(v DeliveryContact) {
	o.DeliveryContact = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *Card) GetExpiration() Expiry {
	if o == nil || common.IsNil(o.Expiration) {
		var ret Expiry
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpirationOk() (*Expiry, bool) {
	if o == nil || common.IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *Card) HasExpiration() bool {
	if o != nil && !common.IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given Expiry and assigns it to the Expiration field.
func (o *Card) SetExpiration(v Expiry) {
	o.Expiration = &v
}

// GetFormFactor returns the FormFactor field value
func (o *Card) GetFormFactor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormFactor
}

// GetFormFactorOk returns a tuple with the FormFactor field value
// and a boolean to check if the value has been set.
func (o *Card) GetFormFactorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormFactor, true
}

// SetFormFactor sets field value
func (o *Card) SetFormFactor(v string) {
	o.FormFactor = v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *Card) GetLastFour() string {
	if o == nil || common.IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetLastFourOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *Card) HasLastFour() bool {
	if o != nil && !common.IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *Card) SetLastFour(v string) {
	o.LastFour = &v
}

// GetNumber returns the Number field value
func (o *Card) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Card) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Card) SetNumber(v string) {
	o.Number = v
}

// GetThreeDSecure returns the ThreeDSecure field value if set, zero value otherwise.
func (o *Card) GetThreeDSecure() string {
	if o == nil || common.IsNil(o.ThreeDSecure) {
		var ret string
		return ret
	}
	return *o.ThreeDSecure
}

// GetThreeDSecureOk returns a tuple with the ThreeDSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetThreeDSecureOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSecure) {
		return nil, false
	}
	return o.ThreeDSecure, true
}

// HasThreeDSecure returns a boolean if a field has been set.
func (o *Card) HasThreeDSecure() bool {
	if o != nil && !common.IsNil(o.ThreeDSecure) {
		return true
	}

	return false
}

// SetThreeDSecure gets a reference to the given string and assigns it to the ThreeDSecure field.
func (o *Card) SetThreeDSecure(v string) {
	o.ThreeDSecure = &v
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Card) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !common.IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	toSerialize["brand"] = o.Brand
	toSerialize["brandVariant"] = o.BrandVariant
	toSerialize["cardholderName"] = o.CardholderName
	if !common.IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !common.IsNil(o.Cvc) {
		toSerialize["cvc"] = o.Cvc
	}
	if !common.IsNil(o.DeliveryContact) {
		toSerialize["deliveryContact"] = o.DeliveryContact
	}
	if !common.IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	toSerialize["formFactor"] = o.FormFactor
	if !common.IsNil(o.LastFour) {
		toSerialize["lastFour"] = o.LastFour
	}
	toSerialize["number"] = o.Number
	if !common.IsNil(o.ThreeDSecure) {
		toSerialize["threeDSecure"] = o.ThreeDSecure
	}
	return toSerialize, nil
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


func (o *Card) isValidFormFactor() bool {
    var allowedEnumValues = []string{ "physical", "unknown", "virtual" }
    for _, allowed := range allowedEnumValues {
        if o.GetFormFactor() == allowed {
            return true
        }
    }
    return false
}

