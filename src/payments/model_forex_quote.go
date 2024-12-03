/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v15/src/common"
	"time"
)

// checks if the ForexQuote type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ForexQuote{}

// ForexQuote struct for ForexQuote
type ForexQuote struct {
	// The account name.
	Account *string `json:"account,omitempty"`
	// The account type.
	AccountType *string `json:"accountType,omitempty"`
	BaseAmount *Amount `json:"baseAmount,omitempty"`
	// The base points.
	BasePoints int32 `json:"basePoints"`
	Buy *Amount `json:"buy,omitempty"`
	Interbank *Amount `json:"interbank,omitempty"`
	// The reference assigned to the forex quote request.
	Reference *string `json:"reference,omitempty"`
	Sell *Amount `json:"sell,omitempty"`
	// The signature to validate the integrity.
	Signature *string `json:"signature,omitempty"`
	// The source of the forex quote.
	Source *string `json:"source,omitempty"`
	// The type of forex.
	Type *string `json:"type,omitempty"`
	// The date until which the forex quote is valid.
	ValidTill time.Time `json:"validTill"`
}

// NewForexQuote instantiates a new ForexQuote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForexQuote(basePoints int32, validTill time.Time) *ForexQuote {
	this := ForexQuote{}
	this.BasePoints = basePoints
	this.ValidTill = validTill
	return &this
}

// NewForexQuoteWithDefaults instantiates a new ForexQuote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForexQuoteWithDefaults() *ForexQuote {
	this := ForexQuote{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ForexQuote) GetAccount() string {
	if o == nil || common.IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ForexQuote) HasAccount() bool {
	if o != nil && !common.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *ForexQuote) SetAccount(v string) {
	o.Account = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *ForexQuote) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *ForexQuote) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *ForexQuote) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBaseAmount returns the BaseAmount field value if set, zero value otherwise.
func (o *ForexQuote) GetBaseAmount() Amount {
	if o == nil || common.IsNil(o.BaseAmount) {
		var ret Amount
		return ret
	}
	return *o.BaseAmount
}

// GetBaseAmountOk returns a tuple with the BaseAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetBaseAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.BaseAmount) {
		return nil, false
	}
	return o.BaseAmount, true
}

// HasBaseAmount returns a boolean if a field has been set.
func (o *ForexQuote) HasBaseAmount() bool {
	if o != nil && !common.IsNil(o.BaseAmount) {
		return true
	}

	return false
}

// SetBaseAmount gets a reference to the given Amount and assigns it to the BaseAmount field.
func (o *ForexQuote) SetBaseAmount(v Amount) {
	o.BaseAmount = &v
}

// GetBasePoints returns the BasePoints field value
func (o *ForexQuote) GetBasePoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BasePoints
}

// GetBasePointsOk returns a tuple with the BasePoints field value
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetBasePointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePoints, true
}

// SetBasePoints sets field value
func (o *ForexQuote) SetBasePoints(v int32) {
	o.BasePoints = v
}

// GetBuy returns the Buy field value if set, zero value otherwise.
func (o *ForexQuote) GetBuy() Amount {
	if o == nil || common.IsNil(o.Buy) {
		var ret Amount
		return ret
	}
	return *o.Buy
}

// GetBuyOk returns a tuple with the Buy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetBuyOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Buy) {
		return nil, false
	}
	return o.Buy, true
}

// HasBuy returns a boolean if a field has been set.
func (o *ForexQuote) HasBuy() bool {
	if o != nil && !common.IsNil(o.Buy) {
		return true
	}

	return false
}

// SetBuy gets a reference to the given Amount and assigns it to the Buy field.
func (o *ForexQuote) SetBuy(v Amount) {
	o.Buy = &v
}

// GetInterbank returns the Interbank field value if set, zero value otherwise.
func (o *ForexQuote) GetInterbank() Amount {
	if o == nil || common.IsNil(o.Interbank) {
		var ret Amount
		return ret
	}
	return *o.Interbank
}

// GetInterbankOk returns a tuple with the Interbank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetInterbankOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Interbank) {
		return nil, false
	}
	return o.Interbank, true
}

// HasInterbank returns a boolean if a field has been set.
func (o *ForexQuote) HasInterbank() bool {
	if o != nil && !common.IsNil(o.Interbank) {
		return true
	}

	return false
}

// SetInterbank gets a reference to the given Amount and assigns it to the Interbank field.
func (o *ForexQuote) SetInterbank(v Amount) {
	o.Interbank = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ForexQuote) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ForexQuote) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ForexQuote) SetReference(v string) {
	o.Reference = &v
}

// GetSell returns the Sell field value if set, zero value otherwise.
func (o *ForexQuote) GetSell() Amount {
	if o == nil || common.IsNil(o.Sell) {
		var ret Amount
		return ret
	}
	return *o.Sell
}

// GetSellOk returns a tuple with the Sell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetSellOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Sell) {
		return nil, false
	}
	return o.Sell, true
}

// HasSell returns a boolean if a field has been set.
func (o *ForexQuote) HasSell() bool {
	if o != nil && !common.IsNil(o.Sell) {
		return true
	}

	return false
}

// SetSell gets a reference to the given Amount and assigns it to the Sell field.
func (o *ForexQuote) SetSell(v Amount) {
	o.Sell = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ForexQuote) GetSignature() string {
	if o == nil || common.IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetSignatureOk() (*string, bool) {
	if o == nil || common.IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ForexQuote) HasSignature() bool {
	if o != nil && !common.IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *ForexQuote) SetSignature(v string) {
	o.Signature = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ForexQuote) GetSource() string {
	if o == nil || common.IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ForexQuote) HasSource() bool {
	if o != nil && !common.IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ForexQuote) SetSource(v string) {
	o.Source = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ForexQuote) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ForexQuote) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ForexQuote) SetType(v string) {
	o.Type = &v
}

// GetValidTill returns the ValidTill field value
func (o *ForexQuote) GetValidTill() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidTill
}

// GetValidTillOk returns a tuple with the ValidTill field value
// and a boolean to check if the value has been set.
func (o *ForexQuote) GetValidTillOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidTill, true
}

// SetValidTill sets field value
func (o *ForexQuote) SetValidTill(v time.Time) {
	o.ValidTill = v
}

func (o ForexQuote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForexQuote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !common.IsNil(o.BaseAmount) {
		toSerialize["baseAmount"] = o.BaseAmount
	}
	toSerialize["basePoints"] = o.BasePoints
	if !common.IsNil(o.Buy) {
		toSerialize["buy"] = o.Buy
	}
	if !common.IsNil(o.Interbank) {
		toSerialize["interbank"] = o.Interbank
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Sell) {
		toSerialize["sell"] = o.Sell
	}
	if !common.IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !common.IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["validTill"] = o.ValidTill
	return toSerialize, nil
}

type NullableForexQuote struct {
	value *ForexQuote
	isSet bool
}

func (v NullableForexQuote) Get() *ForexQuote {
	return v.value
}

func (v *NullableForexQuote) Set(val *ForexQuote) {
	v.value = val
	v.isSet = true
}

func (v NullableForexQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableForexQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForexQuote(val *ForexQuote) *NullableForexQuote {
	return &NullableForexQuote{value: val, isSet: true}
}

func (v NullableForexQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForexQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



