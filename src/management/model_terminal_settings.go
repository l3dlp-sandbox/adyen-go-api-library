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

// checks if the TerminalSettings type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalSettings{}

// TerminalSettings struct for TerminalSettings
type TerminalSettings struct {
	CardholderReceipt *CardholderReceipt `json:"cardholderReceipt,omitempty"`
	Connectivity      *Connectivity      `json:"connectivity,omitempty"`
	// Settings for tipping with or without predefined options to choose from. The maximum number of predefined options is four, or three plus the option to enter a custom tip.
	Gratuities           []Gratuity            `json:"gratuities,omitempty"`
	Hardware             *Hardware             `json:"hardware,omitempty"`
	Localization         *Localization         `json:"localization,omitempty"`
	Nexo                 *Nexo                 `json:"nexo,omitempty"`
	OfflineProcessing    *OfflineProcessing    `json:"offlineProcessing,omitempty"`
	Opi                  *Opi                  `json:"opi,omitempty"`
	Passcodes            *Passcodes            `json:"passcodes,omitempty"`
	PayAtTable           *PayAtTable           `json:"payAtTable,omitempty"`
	Payment              *Payment              `json:"payment,omitempty"`
	ReceiptOptions       *ReceiptOptions       `json:"receiptOptions,omitempty"`
	ReceiptPrinting      *ReceiptPrinting      `json:"receiptPrinting,omitempty"`
	Refunds              *Refunds              `json:"refunds,omitempty"`
	Signature            *Signature            `json:"signature,omitempty"`
	Standalone           *Standalone           `json:"standalone,omitempty"`
	StoreAndForward      *StoreAndForward      `json:"storeAndForward,omitempty"`
	Surcharge            *Surcharge            `json:"surcharge,omitempty"`
	TapToPay             *TapToPay             `json:"tapToPay,omitempty"`
	TerminalInstructions *TerminalInstructions `json:"terminalInstructions,omitempty"`
	Timeouts             *Timeouts             `json:"timeouts,omitempty"`
	WifiProfiles         *WifiProfiles         `json:"wifiProfiles,omitempty"`
}

// NewTerminalSettings instantiates a new TerminalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalSettings() *TerminalSettings {
	this := TerminalSettings{}
	return &this
}

// NewTerminalSettingsWithDefaults instantiates a new TerminalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalSettingsWithDefaults() *TerminalSettings {
	this := TerminalSettings{}
	return &this
}

// GetCardholderReceipt returns the CardholderReceipt field value if set, zero value otherwise.
func (o *TerminalSettings) GetCardholderReceipt() CardholderReceipt {
	if o == nil || common.IsNil(o.CardholderReceipt) {
		var ret CardholderReceipt
		return ret
	}
	return *o.CardholderReceipt
}

// GetCardholderReceiptOk returns a tuple with the CardholderReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetCardholderReceiptOk() (*CardholderReceipt, bool) {
	if o == nil || common.IsNil(o.CardholderReceipt) {
		return nil, false
	}
	return o.CardholderReceipt, true
}

// HasCardholderReceipt returns a boolean if a field has been set.
func (o *TerminalSettings) HasCardholderReceipt() bool {
	if o != nil && !common.IsNil(o.CardholderReceipt) {
		return true
	}

	return false
}

// SetCardholderReceipt gets a reference to the given CardholderReceipt and assigns it to the CardholderReceipt field.
func (o *TerminalSettings) SetCardholderReceipt(v CardholderReceipt) {
	o.CardholderReceipt = &v
}

// GetConnectivity returns the Connectivity field value if set, zero value otherwise.
func (o *TerminalSettings) GetConnectivity() Connectivity {
	if o == nil || common.IsNil(o.Connectivity) {
		var ret Connectivity
		return ret
	}
	return *o.Connectivity
}

// GetConnectivityOk returns a tuple with the Connectivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetConnectivityOk() (*Connectivity, bool) {
	if o == nil || common.IsNil(o.Connectivity) {
		return nil, false
	}
	return o.Connectivity, true
}

// HasConnectivity returns a boolean if a field has been set.
func (o *TerminalSettings) HasConnectivity() bool {
	if o != nil && !common.IsNil(o.Connectivity) {
		return true
	}

	return false
}

// SetConnectivity gets a reference to the given Connectivity and assigns it to the Connectivity field.
func (o *TerminalSettings) SetConnectivity(v Connectivity) {
	o.Connectivity = &v
}

// GetGratuities returns the Gratuities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TerminalSettings) GetGratuities() []Gratuity {
	if o == nil {
		var ret []Gratuity
		return ret
	}
	return o.Gratuities
}

// GetGratuitiesOk returns a tuple with the Gratuities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalSettings) GetGratuitiesOk() ([]Gratuity, bool) {
	if o == nil || common.IsNil(o.Gratuities) {
		return nil, false
	}
	return o.Gratuities, true
}

// HasGratuities returns a boolean if a field has been set.
func (o *TerminalSettings) HasGratuities() bool {
	if o != nil && common.IsNil(o.Gratuities) {
		return true
	}

	return false
}

// SetGratuities gets a reference to the given []Gratuity and assigns it to the Gratuities field.
func (o *TerminalSettings) SetGratuities(v []Gratuity) {
	o.Gratuities = v
}

// GetHardware returns the Hardware field value if set, zero value otherwise.
func (o *TerminalSettings) GetHardware() Hardware {
	if o == nil || common.IsNil(o.Hardware) {
		var ret Hardware
		return ret
	}
	return *o.Hardware
}

// GetHardwareOk returns a tuple with the Hardware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetHardwareOk() (*Hardware, bool) {
	if o == nil || common.IsNil(o.Hardware) {
		return nil, false
	}
	return o.Hardware, true
}

// HasHardware returns a boolean if a field has been set.
func (o *TerminalSettings) HasHardware() bool {
	if o != nil && !common.IsNil(o.Hardware) {
		return true
	}

	return false
}

// SetHardware gets a reference to the given Hardware and assigns it to the Hardware field.
func (o *TerminalSettings) SetHardware(v Hardware) {
	o.Hardware = &v
}

// GetLocalization returns the Localization field value if set, zero value otherwise.
func (o *TerminalSettings) GetLocalization() Localization {
	if o == nil || common.IsNil(o.Localization) {
		var ret Localization
		return ret
	}
	return *o.Localization
}

// GetLocalizationOk returns a tuple with the Localization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetLocalizationOk() (*Localization, bool) {
	if o == nil || common.IsNil(o.Localization) {
		return nil, false
	}
	return o.Localization, true
}

// HasLocalization returns a boolean if a field has been set.
func (o *TerminalSettings) HasLocalization() bool {
	if o != nil && !common.IsNil(o.Localization) {
		return true
	}

	return false
}

// SetLocalization gets a reference to the given Localization and assigns it to the Localization field.
func (o *TerminalSettings) SetLocalization(v Localization) {
	o.Localization = &v
}

// GetNexo returns the Nexo field value if set, zero value otherwise.
func (o *TerminalSettings) GetNexo() Nexo {
	if o == nil || common.IsNil(o.Nexo) {
		var ret Nexo
		return ret
	}
	return *o.Nexo
}

// GetNexoOk returns a tuple with the Nexo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetNexoOk() (*Nexo, bool) {
	if o == nil || common.IsNil(o.Nexo) {
		return nil, false
	}
	return o.Nexo, true
}

// HasNexo returns a boolean if a field has been set.
func (o *TerminalSettings) HasNexo() bool {
	if o != nil && !common.IsNil(o.Nexo) {
		return true
	}

	return false
}

// SetNexo gets a reference to the given Nexo and assigns it to the Nexo field.
func (o *TerminalSettings) SetNexo(v Nexo) {
	o.Nexo = &v
}

// GetOfflineProcessing returns the OfflineProcessing field value if set, zero value otherwise.
func (o *TerminalSettings) GetOfflineProcessing() OfflineProcessing {
	if o == nil || common.IsNil(o.OfflineProcessing) {
		var ret OfflineProcessing
		return ret
	}
	return *o.OfflineProcessing
}

// GetOfflineProcessingOk returns a tuple with the OfflineProcessing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetOfflineProcessingOk() (*OfflineProcessing, bool) {
	if o == nil || common.IsNil(o.OfflineProcessing) {
		return nil, false
	}
	return o.OfflineProcessing, true
}

// HasOfflineProcessing returns a boolean if a field has been set.
func (o *TerminalSettings) HasOfflineProcessing() bool {
	if o != nil && !common.IsNil(o.OfflineProcessing) {
		return true
	}

	return false
}

// SetOfflineProcessing gets a reference to the given OfflineProcessing and assigns it to the OfflineProcessing field.
func (o *TerminalSettings) SetOfflineProcessing(v OfflineProcessing) {
	o.OfflineProcessing = &v
}

// GetOpi returns the Opi field value if set, zero value otherwise.
func (o *TerminalSettings) GetOpi() Opi {
	if o == nil || common.IsNil(o.Opi) {
		var ret Opi
		return ret
	}
	return *o.Opi
}

// GetOpiOk returns a tuple with the Opi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetOpiOk() (*Opi, bool) {
	if o == nil || common.IsNil(o.Opi) {
		return nil, false
	}
	return o.Opi, true
}

// HasOpi returns a boolean if a field has been set.
func (o *TerminalSettings) HasOpi() bool {
	if o != nil && !common.IsNil(o.Opi) {
		return true
	}

	return false
}

// SetOpi gets a reference to the given Opi and assigns it to the Opi field.
func (o *TerminalSettings) SetOpi(v Opi) {
	o.Opi = &v
}

// GetPasscodes returns the Passcodes field value if set, zero value otherwise.
func (o *TerminalSettings) GetPasscodes() Passcodes {
	if o == nil || common.IsNil(o.Passcodes) {
		var ret Passcodes
		return ret
	}
	return *o.Passcodes
}

// GetPasscodesOk returns a tuple with the Passcodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetPasscodesOk() (*Passcodes, bool) {
	if o == nil || common.IsNil(o.Passcodes) {
		return nil, false
	}
	return o.Passcodes, true
}

// HasPasscodes returns a boolean if a field has been set.
func (o *TerminalSettings) HasPasscodes() bool {
	if o != nil && !common.IsNil(o.Passcodes) {
		return true
	}

	return false
}

// SetPasscodes gets a reference to the given Passcodes and assigns it to the Passcodes field.
func (o *TerminalSettings) SetPasscodes(v Passcodes) {
	o.Passcodes = &v
}

// GetPayAtTable returns the PayAtTable field value if set, zero value otherwise.
func (o *TerminalSettings) GetPayAtTable() PayAtTable {
	if o == nil || common.IsNil(o.PayAtTable) {
		var ret PayAtTable
		return ret
	}
	return *o.PayAtTable
}

// GetPayAtTableOk returns a tuple with the PayAtTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetPayAtTableOk() (*PayAtTable, bool) {
	if o == nil || common.IsNil(o.PayAtTable) {
		return nil, false
	}
	return o.PayAtTable, true
}

// HasPayAtTable returns a boolean if a field has been set.
func (o *TerminalSettings) HasPayAtTable() bool {
	if o != nil && !common.IsNil(o.PayAtTable) {
		return true
	}

	return false
}

// SetPayAtTable gets a reference to the given PayAtTable and assigns it to the PayAtTable field.
func (o *TerminalSettings) SetPayAtTable(v PayAtTable) {
	o.PayAtTable = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *TerminalSettings) GetPayment() Payment {
	if o == nil || common.IsNil(o.Payment) {
		var ret Payment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetPaymentOk() (*Payment, bool) {
	if o == nil || common.IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *TerminalSettings) HasPayment() bool {
	if o != nil && !common.IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given Payment and assigns it to the Payment field.
func (o *TerminalSettings) SetPayment(v Payment) {
	o.Payment = &v
}

// GetReceiptOptions returns the ReceiptOptions field value if set, zero value otherwise.
func (o *TerminalSettings) GetReceiptOptions() ReceiptOptions {
	if o == nil || common.IsNil(o.ReceiptOptions) {
		var ret ReceiptOptions
		return ret
	}
	return *o.ReceiptOptions
}

// GetReceiptOptionsOk returns a tuple with the ReceiptOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetReceiptOptionsOk() (*ReceiptOptions, bool) {
	if o == nil || common.IsNil(o.ReceiptOptions) {
		return nil, false
	}
	return o.ReceiptOptions, true
}

// HasReceiptOptions returns a boolean if a field has been set.
func (o *TerminalSettings) HasReceiptOptions() bool {
	if o != nil && !common.IsNil(o.ReceiptOptions) {
		return true
	}

	return false
}

// SetReceiptOptions gets a reference to the given ReceiptOptions and assigns it to the ReceiptOptions field.
func (o *TerminalSettings) SetReceiptOptions(v ReceiptOptions) {
	o.ReceiptOptions = &v
}

// GetReceiptPrinting returns the ReceiptPrinting field value if set, zero value otherwise.
func (o *TerminalSettings) GetReceiptPrinting() ReceiptPrinting {
	if o == nil || common.IsNil(o.ReceiptPrinting) {
		var ret ReceiptPrinting
		return ret
	}
	return *o.ReceiptPrinting
}

// GetReceiptPrintingOk returns a tuple with the ReceiptPrinting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetReceiptPrintingOk() (*ReceiptPrinting, bool) {
	if o == nil || common.IsNil(o.ReceiptPrinting) {
		return nil, false
	}
	return o.ReceiptPrinting, true
}

// HasReceiptPrinting returns a boolean if a field has been set.
func (o *TerminalSettings) HasReceiptPrinting() bool {
	if o != nil && !common.IsNil(o.ReceiptPrinting) {
		return true
	}

	return false
}

// SetReceiptPrinting gets a reference to the given ReceiptPrinting and assigns it to the ReceiptPrinting field.
func (o *TerminalSettings) SetReceiptPrinting(v ReceiptPrinting) {
	o.ReceiptPrinting = &v
}

// GetRefunds returns the Refunds field value if set, zero value otherwise.
func (o *TerminalSettings) GetRefunds() Refunds {
	if o == nil || common.IsNil(o.Refunds) {
		var ret Refunds
		return ret
	}
	return *o.Refunds
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetRefundsOk() (*Refunds, bool) {
	if o == nil || common.IsNil(o.Refunds) {
		return nil, false
	}
	return o.Refunds, true
}

// HasRefunds returns a boolean if a field has been set.
func (o *TerminalSettings) HasRefunds() bool {
	if o != nil && !common.IsNil(o.Refunds) {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given Refunds and assigns it to the Refunds field.
func (o *TerminalSettings) SetRefunds(v Refunds) {
	o.Refunds = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *TerminalSettings) GetSignature() Signature {
	if o == nil || common.IsNil(o.Signature) {
		var ret Signature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetSignatureOk() (*Signature, bool) {
	if o == nil || common.IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *TerminalSettings) HasSignature() bool {
	if o != nil && !common.IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given Signature and assigns it to the Signature field.
func (o *TerminalSettings) SetSignature(v Signature) {
	o.Signature = &v
}

// GetStandalone returns the Standalone field value if set, zero value otherwise.
func (o *TerminalSettings) GetStandalone() Standalone {
	if o == nil || common.IsNil(o.Standalone) {
		var ret Standalone
		return ret
	}
	return *o.Standalone
}

// GetStandaloneOk returns a tuple with the Standalone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetStandaloneOk() (*Standalone, bool) {
	if o == nil || common.IsNil(o.Standalone) {
		return nil, false
	}
	return o.Standalone, true
}

// HasStandalone returns a boolean if a field has been set.
func (o *TerminalSettings) HasStandalone() bool {
	if o != nil && !common.IsNil(o.Standalone) {
		return true
	}

	return false
}

// SetStandalone gets a reference to the given Standalone and assigns it to the Standalone field.
func (o *TerminalSettings) SetStandalone(v Standalone) {
	o.Standalone = &v
}

// GetStoreAndForward returns the StoreAndForward field value if set, zero value otherwise.
func (o *TerminalSettings) GetStoreAndForward() StoreAndForward {
	if o == nil || common.IsNil(o.StoreAndForward) {
		var ret StoreAndForward
		return ret
	}
	return *o.StoreAndForward
}

// GetStoreAndForwardOk returns a tuple with the StoreAndForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetStoreAndForwardOk() (*StoreAndForward, bool) {
	if o == nil || common.IsNil(o.StoreAndForward) {
		return nil, false
	}
	return o.StoreAndForward, true
}

// HasStoreAndForward returns a boolean if a field has been set.
func (o *TerminalSettings) HasStoreAndForward() bool {
	if o != nil && !common.IsNil(o.StoreAndForward) {
		return true
	}

	return false
}

// SetStoreAndForward gets a reference to the given StoreAndForward and assigns it to the StoreAndForward field.
func (o *TerminalSettings) SetStoreAndForward(v StoreAndForward) {
	o.StoreAndForward = &v
}

// GetSurcharge returns the Surcharge field value if set, zero value otherwise.
func (o *TerminalSettings) GetSurcharge() Surcharge {
	if o == nil || common.IsNil(o.Surcharge) {
		var ret Surcharge
		return ret
	}
	return *o.Surcharge
}

// GetSurchargeOk returns a tuple with the Surcharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetSurchargeOk() (*Surcharge, bool) {
	if o == nil || common.IsNil(o.Surcharge) {
		return nil, false
	}
	return o.Surcharge, true
}

// HasSurcharge returns a boolean if a field has been set.
func (o *TerminalSettings) HasSurcharge() bool {
	if o != nil && !common.IsNil(o.Surcharge) {
		return true
	}

	return false
}

// SetSurcharge gets a reference to the given Surcharge and assigns it to the Surcharge field.
func (o *TerminalSettings) SetSurcharge(v Surcharge) {
	o.Surcharge = &v
}

// GetTapToPay returns the TapToPay field value if set, zero value otherwise.
func (o *TerminalSettings) GetTapToPay() TapToPay {
	if o == nil || common.IsNil(o.TapToPay) {
		var ret TapToPay
		return ret
	}
	return *o.TapToPay
}

// GetTapToPayOk returns a tuple with the TapToPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetTapToPayOk() (*TapToPay, bool) {
	if o == nil || common.IsNil(o.TapToPay) {
		return nil, false
	}
	return o.TapToPay, true
}

// HasTapToPay returns a boolean if a field has been set.
func (o *TerminalSettings) HasTapToPay() bool {
	if o != nil && !common.IsNil(o.TapToPay) {
		return true
	}

	return false
}

// SetTapToPay gets a reference to the given TapToPay and assigns it to the TapToPay field.
func (o *TerminalSettings) SetTapToPay(v TapToPay) {
	o.TapToPay = &v
}

// GetTerminalInstructions returns the TerminalInstructions field value if set, zero value otherwise.
func (o *TerminalSettings) GetTerminalInstructions() TerminalInstructions {
	if o == nil || common.IsNil(o.TerminalInstructions) {
		var ret TerminalInstructions
		return ret
	}
	return *o.TerminalInstructions
}

// GetTerminalInstructionsOk returns a tuple with the TerminalInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetTerminalInstructionsOk() (*TerminalInstructions, bool) {
	if o == nil || common.IsNil(o.TerminalInstructions) {
		return nil, false
	}
	return o.TerminalInstructions, true
}

// HasTerminalInstructions returns a boolean if a field has been set.
func (o *TerminalSettings) HasTerminalInstructions() bool {
	if o != nil && !common.IsNil(o.TerminalInstructions) {
		return true
	}

	return false
}

// SetTerminalInstructions gets a reference to the given TerminalInstructions and assigns it to the TerminalInstructions field.
func (o *TerminalSettings) SetTerminalInstructions(v TerminalInstructions) {
	o.TerminalInstructions = &v
}

// GetTimeouts returns the Timeouts field value if set, zero value otherwise.
func (o *TerminalSettings) GetTimeouts() Timeouts {
	if o == nil || common.IsNil(o.Timeouts) {
		var ret Timeouts
		return ret
	}
	return *o.Timeouts
}

// GetTimeoutsOk returns a tuple with the Timeouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetTimeoutsOk() (*Timeouts, bool) {
	if o == nil || common.IsNil(o.Timeouts) {
		return nil, false
	}
	return o.Timeouts, true
}

// HasTimeouts returns a boolean if a field has been set.
func (o *TerminalSettings) HasTimeouts() bool {
	if o != nil && !common.IsNil(o.Timeouts) {
		return true
	}

	return false
}

// SetTimeouts gets a reference to the given Timeouts and assigns it to the Timeouts field.
func (o *TerminalSettings) SetTimeouts(v Timeouts) {
	o.Timeouts = &v
}

// GetWifiProfiles returns the WifiProfiles field value if set, zero value otherwise.
func (o *TerminalSettings) GetWifiProfiles() WifiProfiles {
	if o == nil || common.IsNil(o.WifiProfiles) {
		var ret WifiProfiles
		return ret
	}
	return *o.WifiProfiles
}

// GetWifiProfilesOk returns a tuple with the WifiProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalSettings) GetWifiProfilesOk() (*WifiProfiles, bool) {
	if o == nil || common.IsNil(o.WifiProfiles) {
		return nil, false
	}
	return o.WifiProfiles, true
}

// HasWifiProfiles returns a boolean if a field has been set.
func (o *TerminalSettings) HasWifiProfiles() bool {
	if o != nil && !common.IsNil(o.WifiProfiles) {
		return true
	}

	return false
}

// SetWifiProfiles gets a reference to the given WifiProfiles and assigns it to the WifiProfiles field.
func (o *TerminalSettings) SetWifiProfiles(v WifiProfiles) {
	o.WifiProfiles = &v
}

func (o TerminalSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CardholderReceipt) {
		toSerialize["cardholderReceipt"] = o.CardholderReceipt
	}
	if !common.IsNil(o.Connectivity) {
		toSerialize["connectivity"] = o.Connectivity
	}
	if o.Gratuities != nil {
		toSerialize["gratuities"] = o.Gratuities
	}
	if !common.IsNil(o.Hardware) {
		toSerialize["hardware"] = o.Hardware
	}
	if !common.IsNil(o.Localization) {
		toSerialize["localization"] = o.Localization
	}
	if !common.IsNil(o.Nexo) {
		toSerialize["nexo"] = o.Nexo
	}
	if !common.IsNil(o.OfflineProcessing) {
		toSerialize["offlineProcessing"] = o.OfflineProcessing
	}
	if !common.IsNil(o.Opi) {
		toSerialize["opi"] = o.Opi
	}
	if !common.IsNil(o.Passcodes) {
		toSerialize["passcodes"] = o.Passcodes
	}
	if !common.IsNil(o.PayAtTable) {
		toSerialize["payAtTable"] = o.PayAtTable
	}
	if !common.IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !common.IsNil(o.ReceiptOptions) {
		toSerialize["receiptOptions"] = o.ReceiptOptions
	}
	if !common.IsNil(o.ReceiptPrinting) {
		toSerialize["receiptPrinting"] = o.ReceiptPrinting
	}
	if !common.IsNil(o.Refunds) {
		toSerialize["refunds"] = o.Refunds
	}
	if !common.IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !common.IsNil(o.Standalone) {
		toSerialize["standalone"] = o.Standalone
	}
	if !common.IsNil(o.StoreAndForward) {
		toSerialize["storeAndForward"] = o.StoreAndForward
	}
	if !common.IsNil(o.Surcharge) {
		toSerialize["surcharge"] = o.Surcharge
	}
	if !common.IsNil(o.TapToPay) {
		toSerialize["tapToPay"] = o.TapToPay
	}
	if !common.IsNil(o.TerminalInstructions) {
		toSerialize["terminalInstructions"] = o.TerminalInstructions
	}
	if !common.IsNil(o.Timeouts) {
		toSerialize["timeouts"] = o.Timeouts
	}
	if !common.IsNil(o.WifiProfiles) {
		toSerialize["wifiProfiles"] = o.WifiProfiles
	}
	return toSerialize, nil
}

type NullableTerminalSettings struct {
	value *TerminalSettings
	isSet bool
}

func (v NullableTerminalSettings) Get() *TerminalSettings {
	return v.value
}

func (v *NullableTerminalSettings) Set(val *TerminalSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalSettings(val *TerminalSettings) *NullableTerminalSettings {
	return &NullableTerminalSettings{value: val, isSet: true}
}

func (v NullableTerminalSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
