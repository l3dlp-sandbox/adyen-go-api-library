/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the GetTerminalDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTerminalDetailsResponse{}

// GetTerminalDetailsResponse struct for GetTerminalDetailsResponse
type GetTerminalDetailsResponse struct {
	// The Bluetooth IP address of the terminal.
	BluetoothIp *string `json:"bluetoothIp,omitempty"`
	// The Bluetooth MAC address of the terminal.
	BluetoothMac *string `json:"bluetoothMac,omitempty"`
	// The company account that the terminal is associated with. If this is the only account level shown in the response, the terminal is assigned to the inventory of the company account.
	CompanyAccount string `json:"companyAccount"`
	// The country where the terminal is used.
	Country *string `json:"country,omitempty"`
	// The model name of the terminal.
	DeviceModel *string `json:"deviceModel,omitempty"`
	// Indicates whether assigning IP addresses through a DHCP server is enabled on the terminal.
	DhcpEnabled *bool `json:"dhcpEnabled,omitempty"`
	// The label shown on the status bar of the display. This label (if any) is specified in your Customer Area.
	DisplayLabel *string `json:"displayLabel,omitempty"`
	// The terminal's IP address in your Ethernet network.
	EthernetIp *string `json:"ethernetIp,omitempty"`
	// The terminal's MAC address in your Ethernet network.
	EthernetMac *string `json:"ethernetMac,omitempty"`
	// The software release currently in use on the terminal.
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	// The integrated circuit card identifier (ICCID) of the SIM card in the terminal.
	Iccid *string `json:"iccid,omitempty"`
	// Date and time of the last activity on the terminal. Not included when the last activity was more than 14 days ago.
	LastActivityDateTime *time.Time `json:"lastActivityDateTime,omitempty"`
	// Date and time of the last transaction on the terminal. Not included when the last transaction was more than 14 days ago.
	LastTransactionDateTime *time.Time `json:"lastTransactionDateTime,omitempty"`
	// The Ethernet link negotiation that the terminal uses:   - `auto`: Auto-negotiation  - `100full`: 100 Mbps full duplex
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The merchant account that the terminal is associated with. If the response doesn't contain a `store` the terminal is assigned to this merchant account.
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// Boolean that indicates if the terminal is assigned to the merchant inventory. This is returned when the terminal is assigned to a merchant account.  - If **true**, this indicates that the terminal is in the merchant inventory. This also means that the terminal cannot be boarded.  - If **false**, this indicates that the terminal is assigned to the merchant account as an in-store terminal. This means that the terminal is ready to be boarded, or is already boarded.
	MerchantInventory *bool `json:"merchantInventory,omitempty"`
	// The permanent terminal ID.
	PermanentTerminalId *string `json:"permanentTerminalId,omitempty"`
	// The serial number of the terminal.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// On a terminal that supports 3G or 4G connectivity, indicates the status of the SIM card in the terminal: ACTIVE or INVENTORY.
	SimStatus *string `json:"simStatus,omitempty"`
	// The store code of the store that the terminal is assigned to.
	Store        *string `json:"store,omitempty"`
	StoreDetails *Store  `json:"storeDetails,omitempty"`
	// The unique terminal ID.
	Terminal string `json:"terminal"`
	// The status of the terminal:   - `OnlineToday`, `OnlineLast1Day`, `OnlineLast2Days` etcetera to `OnlineLast7Days`: Indicates when in the past week the terminal was last online.   - `SwitchedOff`: Indicates it was more than a week ago that the terminal was last online.   - `ReAssignToInventoryPending`, `ReAssignToStorePending`, `ReAssignToMerchantInventoryPending`: Indicates the terminal is scheduled to be reassigned.
	TerminalStatus *string `json:"terminalStatus,omitempty"`
	// The terminal's IP address in your Wi-Fi network.
	WifiIp *string `json:"wifiIp,omitempty"`
	// The terminal's MAC address in your Wi-Fi network.
	WifiMac *string `json:"wifiMac,omitempty"`
}

// NewGetTerminalDetailsResponse instantiates a new GetTerminalDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTerminalDetailsResponse(companyAccount string, terminal string) *GetTerminalDetailsResponse {
	this := GetTerminalDetailsResponse{}
	this.CompanyAccount = companyAccount
	this.Terminal = terminal
	return &this
}

// NewGetTerminalDetailsResponseWithDefaults instantiates a new GetTerminalDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTerminalDetailsResponseWithDefaults() *GetTerminalDetailsResponse {
	this := GetTerminalDetailsResponse{}
	return &this
}

// GetBluetoothIp returns the BluetoothIp field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetBluetoothIp() string {
	if o == nil || common.IsNil(o.BluetoothIp) {
		var ret string
		return ret
	}
	return *o.BluetoothIp
}

// GetBluetoothIpOk returns a tuple with the BluetoothIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetBluetoothIpOk() (*string, bool) {
	if o == nil || common.IsNil(o.BluetoothIp) {
		return nil, false
	}
	return o.BluetoothIp, true
}

// HasBluetoothIp returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasBluetoothIp() bool {
	if o != nil && !common.IsNil(o.BluetoothIp) {
		return true
	}

	return false
}

// SetBluetoothIp gets a reference to the given string and assigns it to the BluetoothIp field.
func (o *GetTerminalDetailsResponse) SetBluetoothIp(v string) {
	o.BluetoothIp = &v
}

// GetBluetoothMac returns the BluetoothMac field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetBluetoothMac() string {
	if o == nil || common.IsNil(o.BluetoothMac) {
		var ret string
		return ret
	}
	return *o.BluetoothMac
}

// GetBluetoothMacOk returns a tuple with the BluetoothMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetBluetoothMacOk() (*string, bool) {
	if o == nil || common.IsNil(o.BluetoothMac) {
		return nil, false
	}
	return o.BluetoothMac, true
}

// HasBluetoothMac returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasBluetoothMac() bool {
	if o != nil && !common.IsNil(o.BluetoothMac) {
		return true
	}

	return false
}

// SetBluetoothMac gets a reference to the given string and assigns it to the BluetoothMac field.
func (o *GetTerminalDetailsResponse) SetBluetoothMac(v string) {
	o.BluetoothMac = &v
}

// GetCompanyAccount returns the CompanyAccount field value
func (o *GetTerminalDetailsResponse) GetCompanyAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyAccount
}

// GetCompanyAccountOk returns a tuple with the CompanyAccount field value
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetCompanyAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyAccount, true
}

// SetCompanyAccount sets field value
func (o *GetTerminalDetailsResponse) SetCompanyAccount(v string) {
	o.CompanyAccount = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetCountry() string {
	if o == nil || common.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasCountry() bool {
	if o != nil && !common.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetTerminalDetailsResponse) SetCountry(v string) {
	o.Country = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetDeviceModel() string {
	if o == nil || common.IsNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetDeviceModelOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasDeviceModel() bool {
	if o != nil && !common.IsNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *GetTerminalDetailsResponse) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDhcpEnabled returns the DhcpEnabled field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetDhcpEnabled() bool {
	if o == nil || common.IsNil(o.DhcpEnabled) {
		var ret bool
		return ret
	}
	return *o.DhcpEnabled
}

// GetDhcpEnabledOk returns a tuple with the DhcpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetDhcpEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DhcpEnabled) {
		return nil, false
	}
	return o.DhcpEnabled, true
}

// HasDhcpEnabled returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasDhcpEnabled() bool {
	if o != nil && !common.IsNil(o.DhcpEnabled) {
		return true
	}

	return false
}

// SetDhcpEnabled gets a reference to the given bool and assigns it to the DhcpEnabled field.
func (o *GetTerminalDetailsResponse) SetDhcpEnabled(v bool) {
	o.DhcpEnabled = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetDisplayLabel() string {
	if o == nil || common.IsNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetDisplayLabelOk() (*string, bool) {
	if o == nil || common.IsNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasDisplayLabel() bool {
	if o != nil && !common.IsNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *GetTerminalDetailsResponse) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetEthernetIp returns the EthernetIp field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetEthernetIp() string {
	if o == nil || common.IsNil(o.EthernetIp) {
		var ret string
		return ret
	}
	return *o.EthernetIp
}

// GetEthernetIpOk returns a tuple with the EthernetIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetEthernetIpOk() (*string, bool) {
	if o == nil || common.IsNil(o.EthernetIp) {
		return nil, false
	}
	return o.EthernetIp, true
}

// HasEthernetIp returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasEthernetIp() bool {
	if o != nil && !common.IsNil(o.EthernetIp) {
		return true
	}

	return false
}

// SetEthernetIp gets a reference to the given string and assigns it to the EthernetIp field.
func (o *GetTerminalDetailsResponse) SetEthernetIp(v string) {
	o.EthernetIp = &v
}

// GetEthernetMac returns the EthernetMac field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetEthernetMac() string {
	if o == nil || common.IsNil(o.EthernetMac) {
		var ret string
		return ret
	}
	return *o.EthernetMac
}

// GetEthernetMacOk returns a tuple with the EthernetMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetEthernetMacOk() (*string, bool) {
	if o == nil || common.IsNil(o.EthernetMac) {
		return nil, false
	}
	return o.EthernetMac, true
}

// HasEthernetMac returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasEthernetMac() bool {
	if o != nil && !common.IsNil(o.EthernetMac) {
		return true
	}

	return false
}

// SetEthernetMac gets a reference to the given string and assigns it to the EthernetMac field.
func (o *GetTerminalDetailsResponse) SetEthernetMac(v string) {
	o.EthernetMac = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetFirmwareVersion() string {
	if o == nil || common.IsNil(o.FirmwareVersion) {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirmwareVersion) {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasFirmwareVersion() bool {
	if o != nil && !common.IsNil(o.FirmwareVersion) {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *GetTerminalDetailsResponse) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetIccid returns the Iccid field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetIccid() string {
	if o == nil || common.IsNil(o.Iccid) {
		var ret string
		return ret
	}
	return *o.Iccid
}

// GetIccidOk returns a tuple with the Iccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetIccidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iccid) {
		return nil, false
	}
	return o.Iccid, true
}

// HasIccid returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasIccid() bool {
	if o != nil && !common.IsNil(o.Iccid) {
		return true
	}

	return false
}

// SetIccid gets a reference to the given string and assigns it to the Iccid field.
func (o *GetTerminalDetailsResponse) SetIccid(v string) {
	o.Iccid = &v
}

// GetLastActivityDateTime returns the LastActivityDateTime field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetLastActivityDateTime() time.Time {
	if o == nil || common.IsNil(o.LastActivityDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastActivityDateTime
}

// GetLastActivityDateTimeOk returns a tuple with the LastActivityDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetLastActivityDateTimeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.LastActivityDateTime) {
		return nil, false
	}
	return o.LastActivityDateTime, true
}

// HasLastActivityDateTime returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasLastActivityDateTime() bool {
	if o != nil && !common.IsNil(o.LastActivityDateTime) {
		return true
	}

	return false
}

// SetLastActivityDateTime gets a reference to the given time.Time and assigns it to the LastActivityDateTime field.
func (o *GetTerminalDetailsResponse) SetLastActivityDateTime(v time.Time) {
	o.LastActivityDateTime = &v
}

// GetLastTransactionDateTime returns the LastTransactionDateTime field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetLastTransactionDateTime() time.Time {
	if o == nil || common.IsNil(o.LastTransactionDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastTransactionDateTime
}

// GetLastTransactionDateTimeOk returns a tuple with the LastTransactionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetLastTransactionDateTimeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.LastTransactionDateTime) {
		return nil, false
	}
	return o.LastTransactionDateTime, true
}

// HasLastTransactionDateTime returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasLastTransactionDateTime() bool {
	if o != nil && !common.IsNil(o.LastTransactionDateTime) {
		return true
	}

	return false
}

// SetLastTransactionDateTime gets a reference to the given time.Time and assigns it to the LastTransactionDateTime field.
func (o *GetTerminalDetailsResponse) SetLastTransactionDateTime(v time.Time) {
	o.LastTransactionDateTime = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetLinkNegotiation() string {
	if o == nil || common.IsNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || common.IsNil(o.LinkNegotiation) {
		return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasLinkNegotiation() bool {
	if o != nil && !common.IsNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *GetTerminalDetailsResponse) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *GetTerminalDetailsResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetMerchantInventory returns the MerchantInventory field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetMerchantInventory() bool {
	if o == nil || common.IsNil(o.MerchantInventory) {
		var ret bool
		return ret
	}
	return *o.MerchantInventory
}

// GetMerchantInventoryOk returns a tuple with the MerchantInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetMerchantInventoryOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantInventory) {
		return nil, false
	}
	return o.MerchantInventory, true
}

// HasMerchantInventory returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasMerchantInventory() bool {
	if o != nil && !common.IsNil(o.MerchantInventory) {
		return true
	}

	return false
}

// SetMerchantInventory gets a reference to the given bool and assigns it to the MerchantInventory field.
func (o *GetTerminalDetailsResponse) SetMerchantInventory(v bool) {
	o.MerchantInventory = &v
}

// GetPermanentTerminalId returns the PermanentTerminalId field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetPermanentTerminalId() string {
	if o == nil || common.IsNil(o.PermanentTerminalId) {
		var ret string
		return ret
	}
	return *o.PermanentTerminalId
}

// GetPermanentTerminalIdOk returns a tuple with the PermanentTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetPermanentTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PermanentTerminalId) {
		return nil, false
	}
	return o.PermanentTerminalId, true
}

// HasPermanentTerminalId returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasPermanentTerminalId() bool {
	if o != nil && !common.IsNil(o.PermanentTerminalId) {
		return true
	}

	return false
}

// SetPermanentTerminalId gets a reference to the given string and assigns it to the PermanentTerminalId field.
func (o *GetTerminalDetailsResponse) SetPermanentTerminalId(v string) {
	o.PermanentTerminalId = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetSerialNumber() string {
	if o == nil || common.IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetSerialNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasSerialNumber() bool {
	if o != nil && !common.IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *GetTerminalDetailsResponse) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSimStatus returns the SimStatus field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetSimStatus() string {
	if o == nil || common.IsNil(o.SimStatus) {
		var ret string
		return ret
	}
	return *o.SimStatus
}

// GetSimStatusOk returns a tuple with the SimStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetSimStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.SimStatus) {
		return nil, false
	}
	return o.SimStatus, true
}

// HasSimStatus returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasSimStatus() bool {
	if o != nil && !common.IsNil(o.SimStatus) {
		return true
	}

	return false
}

// SetSimStatus gets a reference to the given string and assigns it to the SimStatus field.
func (o *GetTerminalDetailsResponse) SetSimStatus(v string) {
	o.SimStatus = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *GetTerminalDetailsResponse) SetStore(v string) {
	o.Store = &v
}

// GetStoreDetails returns the StoreDetails field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetStoreDetails() Store {
	if o == nil || common.IsNil(o.StoreDetails) {
		var ret Store
		return ret
	}
	return *o.StoreDetails
}

// GetStoreDetailsOk returns a tuple with the StoreDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetStoreDetailsOk() (*Store, bool) {
	if o == nil || common.IsNil(o.StoreDetails) {
		return nil, false
	}
	return o.StoreDetails, true
}

// HasStoreDetails returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasStoreDetails() bool {
	if o != nil && !common.IsNil(o.StoreDetails) {
		return true
	}

	return false
}

// SetStoreDetails gets a reference to the given Store and assigns it to the StoreDetails field.
func (o *GetTerminalDetailsResponse) SetStoreDetails(v Store) {
	o.StoreDetails = &v
}

// GetTerminal returns the Terminal field value
func (o *GetTerminalDetailsResponse) GetTerminal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetTerminalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terminal, true
}

// SetTerminal sets field value
func (o *GetTerminalDetailsResponse) SetTerminal(v string) {
	o.Terminal = v
}

// GetTerminalStatus returns the TerminalStatus field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetTerminalStatus() string {
	if o == nil || common.IsNil(o.TerminalStatus) {
		var ret string
		return ret
	}
	return *o.TerminalStatus
}

// GetTerminalStatusOk returns a tuple with the TerminalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetTerminalStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.TerminalStatus) {
		return nil, false
	}
	return o.TerminalStatus, true
}

// HasTerminalStatus returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasTerminalStatus() bool {
	if o != nil && !common.IsNil(o.TerminalStatus) {
		return true
	}

	return false
}

// SetTerminalStatus gets a reference to the given string and assigns it to the TerminalStatus field.
func (o *GetTerminalDetailsResponse) SetTerminalStatus(v string) {
	o.TerminalStatus = &v
}

// GetWifiIp returns the WifiIp field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetWifiIp() string {
	if o == nil || common.IsNil(o.WifiIp) {
		var ret string
		return ret
	}
	return *o.WifiIp
}

// GetWifiIpOk returns a tuple with the WifiIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetWifiIpOk() (*string, bool) {
	if o == nil || common.IsNil(o.WifiIp) {
		return nil, false
	}
	return o.WifiIp, true
}

// HasWifiIp returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasWifiIp() bool {
	if o != nil && !common.IsNil(o.WifiIp) {
		return true
	}

	return false
}

// SetWifiIp gets a reference to the given string and assigns it to the WifiIp field.
func (o *GetTerminalDetailsResponse) SetWifiIp(v string) {
	o.WifiIp = &v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *GetTerminalDetailsResponse) GetWifiMac() string {
	if o == nil || common.IsNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsResponse) GetWifiMacOk() (*string, bool) {
	if o == nil || common.IsNil(o.WifiMac) {
		return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *GetTerminalDetailsResponse) HasWifiMac() bool {
	if o != nil && !common.IsNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *GetTerminalDetailsResponse) SetWifiMac(v string) {
	o.WifiMac = &v
}

func (o GetTerminalDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTerminalDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BluetoothIp) {
		toSerialize["bluetoothIp"] = o.BluetoothIp
	}
	if !common.IsNil(o.BluetoothMac) {
		toSerialize["bluetoothMac"] = o.BluetoothMac
	}
	toSerialize["companyAccount"] = o.CompanyAccount
	if !common.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !common.IsNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if !common.IsNil(o.DhcpEnabled) {
		toSerialize["dhcpEnabled"] = o.DhcpEnabled
	}
	if !common.IsNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if !common.IsNil(o.EthernetIp) {
		toSerialize["ethernetIp"] = o.EthernetIp
	}
	if !common.IsNil(o.EthernetMac) {
		toSerialize["ethernetMac"] = o.EthernetMac
	}
	if !common.IsNil(o.FirmwareVersion) {
		toSerialize["firmwareVersion"] = o.FirmwareVersion
	}
	if !common.IsNil(o.Iccid) {
		toSerialize["iccid"] = o.Iccid
	}
	if !common.IsNil(o.LastActivityDateTime) {
		toSerialize["lastActivityDateTime"] = o.LastActivityDateTime
	}
	if !common.IsNil(o.LastTransactionDateTime) {
		toSerialize["lastTransactionDateTime"] = o.LastTransactionDateTime
	}
	if !common.IsNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.MerchantInventory) {
		toSerialize["merchantInventory"] = o.MerchantInventory
	}
	if !common.IsNil(o.PermanentTerminalId) {
		toSerialize["permanentTerminalId"] = o.PermanentTerminalId
	}
	if !common.IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !common.IsNil(o.SimStatus) {
		toSerialize["simStatus"] = o.SimStatus
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	if !common.IsNil(o.StoreDetails) {
		toSerialize["storeDetails"] = o.StoreDetails
	}
	toSerialize["terminal"] = o.Terminal
	if !common.IsNil(o.TerminalStatus) {
		toSerialize["terminalStatus"] = o.TerminalStatus
	}
	if !common.IsNil(o.WifiIp) {
		toSerialize["wifiIp"] = o.WifiIp
	}
	if !common.IsNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	return toSerialize, nil
}

type NullableGetTerminalDetailsResponse struct {
	value *GetTerminalDetailsResponse
	isSet bool
}

func (v NullableGetTerminalDetailsResponse) Get() *GetTerminalDetailsResponse {
	return v.value
}

func (v *NullableGetTerminalDetailsResponse) Set(val *GetTerminalDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTerminalDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTerminalDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTerminalDetailsResponse(val *GetTerminalDetailsResponse) *NullableGetTerminalDetailsResponse {
	return &NullableGetTerminalDetailsResponse{value: val, isSet: true}
}

func (v NullableGetTerminalDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTerminalDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *GetTerminalDetailsResponse) isValidTerminalStatus() bool {
	var allowedEnumValues = []string{"OnlineLast1Day", "OnlineLast2Days", "OnlineLast3Days", "OnlineLast4Days", "OnlineLast5Days", "OnlineLast6Days", "OnlineLast7Days", "OnlineToday", "ReAssignToInventoryPending", "ReAssignToMerchantInventoryPending", "ReAssignToStorePending", "SwitchedOff"}
	for _, allowed := range allowedEnumValues {
		if o.GetTerminalStatus() == allowed {
			return true
		}
	}
	return false
}
