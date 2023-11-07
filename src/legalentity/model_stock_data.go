/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the StockData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StockData{}

// StockData struct for StockData
type StockData struct {
	// The four-digit [Market Identifier Code](https://en.wikipedia.org/wiki/Market_Identifier_Code) of the stock market where the organization's stocks are traded.
	MarketIdentifier *string `json:"marketIdentifier,omitempty"`
	// The 12-digit International Securities Identification Number (ISIN) of the company, without dashes (-).
	StockNumber *string `json:"stockNumber,omitempty"`
	// The stock ticker symbol.
	TickerSymbol *string `json:"tickerSymbol,omitempty"`
}

// NewStockData instantiates a new StockData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockData() *StockData {
	this := StockData{}
	return &this
}

// NewStockDataWithDefaults instantiates a new StockData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockDataWithDefaults() *StockData {
	this := StockData{}
	return &this
}

// GetMarketIdentifier returns the MarketIdentifier field value if set, zero value otherwise.
func (o *StockData) GetMarketIdentifier() string {
	if o == nil || common.IsNil(o.MarketIdentifier) {
		var ret string
		return ret
	}
	return *o.MarketIdentifier
}

// GetMarketIdentifierOk returns a tuple with the MarketIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockData) GetMarketIdentifierOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketIdentifier) {
		return nil, false
	}
	return o.MarketIdentifier, true
}

// HasMarketIdentifier returns a boolean if a field has been set.
func (o *StockData) HasMarketIdentifier() bool {
	if o != nil && !common.IsNil(o.MarketIdentifier) {
		return true
	}

	return false
}

// SetMarketIdentifier gets a reference to the given string and assigns it to the MarketIdentifier field.
func (o *StockData) SetMarketIdentifier(v string) {
	o.MarketIdentifier = &v
}

// GetStockNumber returns the StockNumber field value if set, zero value otherwise.
func (o *StockData) GetStockNumber() string {
	if o == nil || common.IsNil(o.StockNumber) {
		var ret string
		return ret
	}
	return *o.StockNumber
}

// GetStockNumberOk returns a tuple with the StockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockData) GetStockNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.StockNumber) {
		return nil, false
	}
	return o.StockNumber, true
}

// HasStockNumber returns a boolean if a field has been set.
func (o *StockData) HasStockNumber() bool {
	if o != nil && !common.IsNil(o.StockNumber) {
		return true
	}

	return false
}

// SetStockNumber gets a reference to the given string and assigns it to the StockNumber field.
func (o *StockData) SetStockNumber(v string) {
	o.StockNumber = &v
}

// GetTickerSymbol returns the TickerSymbol field value if set, zero value otherwise.
func (o *StockData) GetTickerSymbol() string {
	if o == nil || common.IsNil(o.TickerSymbol) {
		var ret string
		return ret
	}
	return *o.TickerSymbol
}

// GetTickerSymbolOk returns a tuple with the TickerSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockData) GetTickerSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.TickerSymbol) {
		return nil, false
	}
	return o.TickerSymbol, true
}

// HasTickerSymbol returns a boolean if a field has been set.
func (o *StockData) HasTickerSymbol() bool {
	if o != nil && !common.IsNil(o.TickerSymbol) {
		return true
	}

	return false
}

// SetTickerSymbol gets a reference to the given string and assigns it to the TickerSymbol field.
func (o *StockData) SetTickerSymbol(v string) {
	o.TickerSymbol = &v
}

func (o StockData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarketIdentifier) {
		toSerialize["marketIdentifier"] = o.MarketIdentifier
	}
	if !common.IsNil(o.StockNumber) {
		toSerialize["stockNumber"] = o.StockNumber
	}
	if !common.IsNil(o.TickerSymbol) {
		toSerialize["tickerSymbol"] = o.TickerSymbol
	}
	return toSerialize, nil
}

type NullableStockData struct {
	value *StockData
	isSet bool
}

func (v NullableStockData) Get() *StockData {
	return v.value
}

func (v *NullableStockData) Set(val *StockData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockData(val *StockData) *NullableStockData {
	return &NullableStockData{value: val, isSet: true}
}

func (v NullableStockData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
