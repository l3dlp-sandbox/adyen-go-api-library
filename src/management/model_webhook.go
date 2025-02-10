/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// checks if the Webhook type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Webhook{}

// Webhook struct for Webhook
type Webhook struct {
	Links *WebhookLinks `json:"_links,omitempty"`
	// Indicates if expired SSL certificates are accepted. Default value: **false**.
	AcceptsExpiredCertificate *bool `json:"acceptsExpiredCertificate,omitempty"`
	// Indicates if self-signed SSL certificates are accepted. Default value: **false**.
	AcceptsSelfSignedCertificate *bool `json:"acceptsSelfSignedCertificate,omitempty"`
	// Indicates if untrusted SSL certificates are accepted. Default value: **false**.
	AcceptsUntrustedRootCertificate *bool `json:"acceptsUntrustedRootCertificate,omitempty"`
	// Reference to the account the webook is set on.
	AccountReference *string `json:"accountReference,omitempty"`
	// Indicates if the webhook configuration is active. The field must be **true** for you to receive webhooks about events related an account.
	Active             bool                        `json:"active"`
	AdditionalSettings *AdditionalSettingsResponse `json:"additionalSettings,omitempty"`
	// The alias of our SSL certificate. When you receive a notification from us, the alias from the HMAC signature will match this alias.
	CertificateAlias *string `json:"certificateAlias,omitempty"`
	// Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json**
	CommunicationFormat string `json:"communicationFormat"`
	// Your description for this webhook configuration.
	Description *string `json:"description,omitempty"`
	// SSL version to access the public webhook URL specified in the `url` field. Possible values: * **TLSv1.3** * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use `sslVersion`: **TLSv1.2**.
	EncryptionProtocol *string `json:"encryptionProtocol,omitempty"`
	// Shows how merchant accounts are included in company-level webhooks. Possible values: * **includeAccounts** * **excludeAccounts** * **allAccounts**: Includes all merchant accounts, and does not require specifying `filterMerchantAccounts`.
	FilterMerchantAccountType *string `json:"filterMerchantAccountType,omitempty"`
	// A list of merchant account names that are included or excluded from receiving the webhook. Inclusion or exclusion is based on the value defined for `filterMerchantAccountType`.  Required if `filterMerchantAccountType` is either: * **includeAccounts** * **excludeAccounts**  Not needed for `filterMerchantAccountType`: **allAccounts**.
	FilterMerchantAccounts []string `json:"filterMerchantAccounts,omitempty"`
	// Indicates if the webhook configuration has errors that need troubleshooting. If the value is **true**, troubleshoot the configuration using the [testing endpoint](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/companies/{companyId}/webhooks/{webhookid}/test).
	HasError *bool `json:"hasError,omitempty"`
	// Indicates if the webhook is password protected.
	HasPassword *bool `json:"hasPassword,omitempty"`
	// The [checksum](https://en.wikipedia.org/wiki/Key_checksum_value) of the HMAC key generated for this webhook. You can use this value to uniquely identify the HMAC key configured for this webhook.
	HmacKeyCheckValue *string `json:"hmacKeyCheckValue,omitempty"`
	// Unique identifier for this webhook.
	Id *string `json:"id,omitempty"`
	// Network type for Terminal API details webhooks.
	NetworkType *string `json:"networkType,omitempty"`
	// Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if `communicationFormat`: **soap**.
	PopulateSoapActionHeader *bool `json:"populateSoapActionHeader,omitempty"`
	// The type of webhook. Possible values are:  - **standard** - **account-settings-notification** - **banktransfer-notification** - **boletobancario-notification** - **directdebit-notification** - **ach-notification-of-change-notification** - **pending-notification** - **ideal-notification** - **ideal-pending-notification** - **report-notification** - **terminal-api-notification** - **terminal-settings** - **terminal-boarding**  Find out more about [standard notification webhooks](https://docs.adyen.com/development-resources/webhooks/understand-notifications#event-codes) and [other types of notifications](https://docs.adyen.com/development-resources/webhooks/understand-notifications#other-notifications).
	Type string `json:"type"`
	// Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**.
	Url string `json:"url"`
	// Username to access the webhook URL.
	Username *string `json:"username,omitempty"`
}

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook(active bool, communicationFormat string, type_ string, url string) *Webhook {
	this := Webhook{}
	this.Active = active
	this.CommunicationFormat = communicationFormat
	this.Type = type_
	this.Url = url
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Webhook) GetLinks() WebhookLinks {
	if o == nil || common.IsNil(o.Links) {
		var ret WebhookLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetLinksOk() (*WebhookLinks, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Webhook) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WebhookLinks and assigns it to the Links field.
func (o *Webhook) SetLinks(v WebhookLinks) {
	o.Links = &v
}

// GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field value if set, zero value otherwise.
func (o *Webhook) GetAcceptsExpiredCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsExpiredCertificate
}

// GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetAcceptsExpiredCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		return nil, false
	}
	return o.AcceptsExpiredCertificate, true
}

// HasAcceptsExpiredCertificate returns a boolean if a field has been set.
func (o *Webhook) HasAcceptsExpiredCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsExpiredCertificate) {
		return true
	}

	return false
}

// SetAcceptsExpiredCertificate gets a reference to the given bool and assigns it to the AcceptsExpiredCertificate field.
func (o *Webhook) SetAcceptsExpiredCertificate(v bool) {
	o.AcceptsExpiredCertificate = &v
}

// GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field value if set, zero value otherwise.
func (o *Webhook) GetAcceptsSelfSignedCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsSelfSignedCertificate
}

// GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetAcceptsSelfSignedCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		return nil, false
	}
	return o.AcceptsSelfSignedCertificate, true
}

// HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.
func (o *Webhook) HasAcceptsSelfSignedCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsSelfSignedCertificate) {
		return true
	}

	return false
}

// SetAcceptsSelfSignedCertificate gets a reference to the given bool and assigns it to the AcceptsSelfSignedCertificate field.
func (o *Webhook) SetAcceptsSelfSignedCertificate(v bool) {
	o.AcceptsSelfSignedCertificate = &v
}

// GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field value if set, zero value otherwise.
func (o *Webhook) GetAcceptsUntrustedRootCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsUntrustedRootCertificate
}

// GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetAcceptsUntrustedRootCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return nil, false
	}
	return o.AcceptsUntrustedRootCertificate, true
}

// HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.
func (o *Webhook) HasAcceptsUntrustedRootCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return true
	}

	return false
}

// SetAcceptsUntrustedRootCertificate gets a reference to the given bool and assigns it to the AcceptsUntrustedRootCertificate field.
func (o *Webhook) SetAcceptsUntrustedRootCertificate(v bool) {
	o.AcceptsUntrustedRootCertificate = &v
}

// GetAccountReference returns the AccountReference field value if set, zero value otherwise.
func (o *Webhook) GetAccountReference() string {
	if o == nil || common.IsNil(o.AccountReference) {
		var ret string
		return ret
	}
	return *o.AccountReference
}

// GetAccountReferenceOk returns a tuple with the AccountReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetAccountReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountReference) {
		return nil, false
	}
	return o.AccountReference, true
}

// HasAccountReference returns a boolean if a field has been set.
func (o *Webhook) HasAccountReference() bool {
	if o != nil && !common.IsNil(o.AccountReference) {
		return true
	}

	return false
}

// SetAccountReference gets a reference to the given string and assigns it to the AccountReference field.
func (o *Webhook) SetAccountReference(v string) {
	o.AccountReference = &v
}

// GetActive returns the Active field value
func (o *Webhook) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *Webhook) SetActive(v bool) {
	o.Active = v
}

// GetAdditionalSettings returns the AdditionalSettings field value if set, zero value otherwise.
func (o *Webhook) GetAdditionalSettings() AdditionalSettingsResponse {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		var ret AdditionalSettingsResponse
		return ret
	}
	return *o.AdditionalSettings
}

// GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetAdditionalSettingsOk() (*AdditionalSettingsResponse, bool) {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		return nil, false
	}
	return o.AdditionalSettings, true
}

// HasAdditionalSettings returns a boolean if a field has been set.
func (o *Webhook) HasAdditionalSettings() bool {
	if o != nil && !common.IsNil(o.AdditionalSettings) {
		return true
	}

	return false
}

// SetAdditionalSettings gets a reference to the given AdditionalSettingsResponse and assigns it to the AdditionalSettings field.
func (o *Webhook) SetAdditionalSettings(v AdditionalSettingsResponse) {
	o.AdditionalSettings = &v
}

// GetCertificateAlias returns the CertificateAlias field value if set, zero value otherwise.
func (o *Webhook) GetCertificateAlias() string {
	if o == nil || common.IsNil(o.CertificateAlias) {
		var ret string
		return ret
	}
	return *o.CertificateAlias
}

// GetCertificateAliasOk returns a tuple with the CertificateAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCertificateAliasOk() (*string, bool) {
	if o == nil || common.IsNil(o.CertificateAlias) {
		return nil, false
	}
	return o.CertificateAlias, true
}

// HasCertificateAlias returns a boolean if a field has been set.
func (o *Webhook) HasCertificateAlias() bool {
	if o != nil && !common.IsNil(o.CertificateAlias) {
		return true
	}

	return false
}

// SetCertificateAlias gets a reference to the given string and assigns it to the CertificateAlias field.
func (o *Webhook) SetCertificateAlias(v string) {
	o.CertificateAlias = &v
}

// GetCommunicationFormat returns the CommunicationFormat field value
func (o *Webhook) GetCommunicationFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommunicationFormat
}

// GetCommunicationFormatOk returns a tuple with the CommunicationFormat field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetCommunicationFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunicationFormat, true
}

// SetCommunicationFormat sets field value
func (o *Webhook) SetCommunicationFormat(v string) {
	o.CommunicationFormat = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Webhook) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Webhook) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Webhook) SetDescription(v string) {
	o.Description = &v
}

// GetEncryptionProtocol returns the EncryptionProtocol field value if set, zero value otherwise.
func (o *Webhook) GetEncryptionProtocol() string {
	if o == nil || common.IsNil(o.EncryptionProtocol) {
		var ret string
		return ret
	}
	return *o.EncryptionProtocol
}

// GetEncryptionProtocolOk returns a tuple with the EncryptionProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEncryptionProtocolOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptionProtocol) {
		return nil, false
	}
	return o.EncryptionProtocol, true
}

// HasEncryptionProtocol returns a boolean if a field has been set.
func (o *Webhook) HasEncryptionProtocol() bool {
	if o != nil && !common.IsNil(o.EncryptionProtocol) {
		return true
	}

	return false
}

// SetEncryptionProtocol gets a reference to the given string and assigns it to the EncryptionProtocol field.
func (o *Webhook) SetEncryptionProtocol(v string) {
	o.EncryptionProtocol = &v
}

// GetFilterMerchantAccountType returns the FilterMerchantAccountType field value if set, zero value otherwise.
func (o *Webhook) GetFilterMerchantAccountType() string {
	if o == nil || common.IsNil(o.FilterMerchantAccountType) {
		var ret string
		return ret
	}
	return *o.FilterMerchantAccountType
}

// GetFilterMerchantAccountTypeOk returns a tuple with the FilterMerchantAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetFilterMerchantAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterMerchantAccountType) {
		return nil, false
	}
	return o.FilterMerchantAccountType, true
}

// HasFilterMerchantAccountType returns a boolean if a field has been set.
func (o *Webhook) HasFilterMerchantAccountType() bool {
	if o != nil && !common.IsNil(o.FilterMerchantAccountType) {
		return true
	}

	return false
}

// SetFilterMerchantAccountType gets a reference to the given string and assigns it to the FilterMerchantAccountType field.
func (o *Webhook) SetFilterMerchantAccountType(v string) {
	o.FilterMerchantAccountType = &v
}

// GetFilterMerchantAccounts returns the FilterMerchantAccounts field value if set, zero value otherwise.
func (o *Webhook) GetFilterMerchantAccounts() []string {
	if o == nil || common.IsNil(o.FilterMerchantAccounts) {
		var ret []string
		return ret
	}
	return o.FilterMerchantAccounts
}

// GetFilterMerchantAccountsOk returns a tuple with the FilterMerchantAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetFilterMerchantAccountsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.FilterMerchantAccounts) {
		return nil, false
	}
	return o.FilterMerchantAccounts, true
}

// HasFilterMerchantAccounts returns a boolean if a field has been set.
func (o *Webhook) HasFilterMerchantAccounts() bool {
	if o != nil && !common.IsNil(o.FilterMerchantAccounts) {
		return true
	}

	return false
}

// SetFilterMerchantAccounts gets a reference to the given []string and assigns it to the FilterMerchantAccounts field.
func (o *Webhook) SetFilterMerchantAccounts(v []string) {
	o.FilterMerchantAccounts = v
}

// GetHasError returns the HasError field value if set, zero value otherwise.
func (o *Webhook) GetHasError() bool {
	if o == nil || common.IsNil(o.HasError) {
		var ret bool
		return ret
	}
	return *o.HasError
}

// GetHasErrorOk returns a tuple with the HasError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHasErrorOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HasError) {
		return nil, false
	}
	return o.HasError, true
}

// HasHasError returns a boolean if a field has been set.
func (o *Webhook) HasHasError() bool {
	if o != nil && !common.IsNil(o.HasError) {
		return true
	}

	return false
}

// SetHasError gets a reference to the given bool and assigns it to the HasError field.
func (o *Webhook) SetHasError(v bool) {
	o.HasError = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *Webhook) GetHasPassword() bool {
	if o == nil || common.IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHasPasswordOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *Webhook) HasHasPassword() bool {
	if o != nil && !common.IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *Webhook) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetHmacKeyCheckValue returns the HmacKeyCheckValue field value if set, zero value otherwise.
func (o *Webhook) GetHmacKeyCheckValue() string {
	if o == nil || common.IsNil(o.HmacKeyCheckValue) {
		var ret string
		return ret
	}
	return *o.HmacKeyCheckValue
}

// GetHmacKeyCheckValueOk returns a tuple with the HmacKeyCheckValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHmacKeyCheckValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.HmacKeyCheckValue) {
		return nil, false
	}
	return o.HmacKeyCheckValue, true
}

// HasHmacKeyCheckValue returns a boolean if a field has been set.
func (o *Webhook) HasHmacKeyCheckValue() bool {
	if o != nil && !common.IsNil(o.HmacKeyCheckValue) {
		return true
	}

	return false
}

// SetHmacKeyCheckValue gets a reference to the given string and assigns it to the HmacKeyCheckValue field.
func (o *Webhook) SetHmacKeyCheckValue(v string) {
	o.HmacKeyCheckValue = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Webhook) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Webhook) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Webhook) SetId(v string) {
	o.Id = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *Webhook) GetNetworkType() string {
	if o == nil || common.IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetNetworkTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *Webhook) HasNetworkType() bool {
	if o != nil && !common.IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *Webhook) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field value if set, zero value otherwise.
func (o *Webhook) GetPopulateSoapActionHeader() bool {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		var ret bool
		return ret
	}
	return *o.PopulateSoapActionHeader
}

// GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetPopulateSoapActionHeaderOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		return nil, false
	}
	return o.PopulateSoapActionHeader, true
}

// HasPopulateSoapActionHeader returns a boolean if a field has been set.
func (o *Webhook) HasPopulateSoapActionHeader() bool {
	if o != nil && !common.IsNil(o.PopulateSoapActionHeader) {
		return true
	}

	return false
}

// SetPopulateSoapActionHeader gets a reference to the given bool and assigns it to the PopulateSoapActionHeader field.
func (o *Webhook) SetPopulateSoapActionHeader(v bool) {
	o.PopulateSoapActionHeader = &v
}

// GetType returns the Type field value
func (o *Webhook) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Webhook) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *Webhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Webhook) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Webhook) GetUsername() string {
	if o == nil || common.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Webhook) HasUsername() bool {
	if o != nil && !common.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Webhook) SetUsername(v string) {
	o.Username = &v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Webhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !common.IsNil(o.AcceptsExpiredCertificate) {
		toSerialize["acceptsExpiredCertificate"] = o.AcceptsExpiredCertificate
	}
	if !common.IsNil(o.AcceptsSelfSignedCertificate) {
		toSerialize["acceptsSelfSignedCertificate"] = o.AcceptsSelfSignedCertificate
	}
	if !common.IsNil(o.AcceptsUntrustedRootCertificate) {
		toSerialize["acceptsUntrustedRootCertificate"] = o.AcceptsUntrustedRootCertificate
	}
	if !common.IsNil(o.AccountReference) {
		toSerialize["accountReference"] = o.AccountReference
	}
	toSerialize["active"] = o.Active
	if !common.IsNil(o.AdditionalSettings) {
		toSerialize["additionalSettings"] = o.AdditionalSettings
	}
	if !common.IsNil(o.CertificateAlias) {
		toSerialize["certificateAlias"] = o.CertificateAlias
	}
	toSerialize["communicationFormat"] = o.CommunicationFormat
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.EncryptionProtocol) {
		toSerialize["encryptionProtocol"] = o.EncryptionProtocol
	}
	if !common.IsNil(o.FilterMerchantAccountType) {
		toSerialize["filterMerchantAccountType"] = o.FilterMerchantAccountType
	}
	if !common.IsNil(o.FilterMerchantAccounts) {
		toSerialize["filterMerchantAccounts"] = o.FilterMerchantAccounts
	}
	if !common.IsNil(o.HasError) {
		toSerialize["hasError"] = o.HasError
	}
	if !common.IsNil(o.HasPassword) {
		toSerialize["hasPassword"] = o.HasPassword
	}
	if !common.IsNil(o.HmacKeyCheckValue) {
		toSerialize["hmacKeyCheckValue"] = o.HmacKeyCheckValue
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.NetworkType) {
		toSerialize["networkType"] = o.NetworkType
	}
	if !common.IsNil(o.PopulateSoapActionHeader) {
		toSerialize["populateSoapActionHeader"] = o.PopulateSoapActionHeader
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	if !common.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Webhook) isValidCommunicationFormat() bool {
	var allowedEnumValues = []string{"http", "json", "soap"}
	for _, allowed := range allowedEnumValues {
		if o.GetCommunicationFormat() == allowed {
			return true
		}
	}
	return false
}
func (o *Webhook) isValidEncryptionProtocol() bool {
	var allowedEnumValues = []string{"HTTP", "TLSv1.2", "TLSv1.3"}
	for _, allowed := range allowedEnumValues {
		if o.GetEncryptionProtocol() == allowed {
			return true
		}
	}
	return false
}
func (o *Webhook) isValidFilterMerchantAccountType() bool {
	var allowedEnumValues = []string{"allAccounts", "excludeAccounts", "includeAccounts"}
	for _, allowed := range allowedEnumValues {
		if o.GetFilterMerchantAccountType() == allowed {
			return true
		}
	}
	return false
}
func (o *Webhook) isValidNetworkType() bool {
	var allowedEnumValues = []string{"local", "public"}
	for _, allowed := range allowedEnumValues {
		if o.GetNetworkType() == allowed {
			return true
		}
	}
	return false
}
