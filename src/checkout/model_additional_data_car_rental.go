/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the AdditionalDataCarRental type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataCarRental{}

// AdditionalDataCarRental struct for AdditionalDataCarRental
type AdditionalDataCarRental struct {
	// The pick-up date. * Date format: `yyyyMMdd`
	CarRentalCheckOutDate *string `json:"carRental.checkOutDate,omitempty"`
	// The customer service phone number of the car rental company. * Format: Alphanumeric * maxLength: 17 * For US and CA numbers must be 10 characters in length * Must not start with a space * Must not contain any special characters such as + or - *Must not be all zeros.
	CarRentalCustomerServiceTollFreeNumber *string `json:"carRental.customerServiceTollFreeNumber,omitempty"`
	// Number of days for which the car is being rented. * Format: Numeric * maxLength: 4 * Must not be all spaces
	CarRentalDaysRented *string `json:"carRental.daysRented,omitempty"`
	// Any fuel charges associated with the rental, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: Numeric * maxLength: 12
	CarRentalFuelCharges *string `json:"carRental.fuelCharges,omitempty"`
	// Any insurance charges associated with the rental, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: Numeric * maxLength: 12 * Must not be all spaces *Must not be all zeros.
	CarRentalInsuranceCharges *string `json:"carRental.insuranceCharges,omitempty"`
	// The city where the car is rented. * Format: Alphanumeric * maxLength: 18 * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalLocationCity *string `json:"carRental.locationCity,omitempty"`
	// The country where the car is rented, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. * Format: Alphanumeric * maxLength: 2
	CarRentalLocationCountry *string `json:"carRental.locationCountry,omitempty"`
	// The state or province where the car is rented. * Format: Alphanumeric * maxLength: 2 * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalLocationStateProvince *string `json:"carRental.locationStateProvince,omitempty"`
	// Indicates if the customer didn't pick up their rental car. * Y - Customer did not pick up their car * N - Not applicable
	CarRentalNoShowIndicator *string `json:"carRental.noShowIndicator,omitempty"`
	// The charge for not returning a car to the original rental location, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * maxLength: 12
	CarRentalOneWayDropOffCharges *string `json:"carRental.oneWayDropOffCharges,omitempty"`
	// The daily rental rate, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: Alphanumeric * maxLength: 12
	CarRentalRate *string `json:"carRental.rate,omitempty"`
	// Specifies whether the given rate is applied daily or weekly. * D - Daily rate * W - Weekly rate
	CarRentalRateIndicator *string `json:"carRental.rateIndicator,omitempty"`
	// The rental agreement number for the car rental. * Format: Alphanumeric * maxLength: 9 * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalRentalAgreementNumber *string `json:"carRental.rentalAgreementNumber,omitempty"`
	// The classification of the rental car. * Format: Alphanumeric * maxLength: 4 * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalRentalClassId *string `json:"carRental.rentalClassId,omitempty"`
	// The name of the person renting the car. * Format: Alphanumeric * maxLength: 26 * If you send more than 26 characters, the name is truncated * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalRenterName *string `json:"carRental.renterName,omitempty"`
	// The city where the car must be returned. * Format: Alphanumeric * maxLength: 18 * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalReturnCity *string `json:"carRental.returnCity,omitempty"`
	// The country where the car must be returned, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. * Format: Alphanumeric * maxLength: 2
	CarRentalReturnCountry *string `json:"carRental.returnCountry,omitempty"`
	// The last date to return the car by. * Date format: `yyyyMMdd` * maxLength: 8
	CarRentalReturnDate *string `json:"carRental.returnDate,omitempty"`
	// The agency code, phone number, or address abbreviation * Format: Alphanumeric * maxLength: 10 * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalReturnLocationId *string `json:"carRental.returnLocationId,omitempty"`
	// The state or province where the car must be returned. * Format: Alphanumeric * maxLength: 3 * Must not start with a space or be all spaces *Must not be all zeros.
	CarRentalReturnStateProvince *string `json:"carRental.returnStateProvince,omitempty"`
	// Indicates if the goods or services were tax-exempt, or if tax was not paid on them.  Values: * Y - Goods or services were tax exempt * N - Tax was not collected
	CarRentalTaxExemptIndicator *string `json:"carRental.taxExemptIndicator,omitempty"`
	// Number of days the car is rented for. This should be included in the auth message. * Format: Numeric * maxLength: 4
	TravelEntertainmentAuthDataDuration *string `json:"travelEntertainmentAuthData.duration,omitempty"`
	// Indicates what market-specific dataset will be submitted or is being submitted. Value should be 'A' for car rental. This should be included in the auth message. * Format: Alphanumeric * maxLength: 1
	TravelEntertainmentAuthDataMarket *string `json:"travelEntertainmentAuthData.market,omitempty"`
}

// NewAdditionalDataCarRental instantiates a new AdditionalDataCarRental object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataCarRental() *AdditionalDataCarRental {
	this := AdditionalDataCarRental{}
	return &this
}

// NewAdditionalDataCarRentalWithDefaults instantiates a new AdditionalDataCarRental object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataCarRentalWithDefaults() *AdditionalDataCarRental {
	this := AdditionalDataCarRental{}
	return &this
}

// GetCarRentalCheckOutDate returns the CarRentalCheckOutDate field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalCheckOutDate() string {
	if o == nil || common.IsNil(o.CarRentalCheckOutDate) {
		var ret string
		return ret
	}
	return *o.CarRentalCheckOutDate
}

// GetCarRentalCheckOutDateOk returns a tuple with the CarRentalCheckOutDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalCheckOutDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalCheckOutDate) {
		return nil, false
	}
	return o.CarRentalCheckOutDate, true
}

// HasCarRentalCheckOutDate returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalCheckOutDate() bool {
	if o != nil && !common.IsNil(o.CarRentalCheckOutDate) {
		return true
	}

	return false
}

// SetCarRentalCheckOutDate gets a reference to the given string and assigns it to the CarRentalCheckOutDate field.
func (o *AdditionalDataCarRental) SetCarRentalCheckOutDate(v string) {
	o.CarRentalCheckOutDate = &v
}

// GetCarRentalCustomerServiceTollFreeNumber returns the CarRentalCustomerServiceTollFreeNumber field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalCustomerServiceTollFreeNumber() string {
	if o == nil || common.IsNil(o.CarRentalCustomerServiceTollFreeNumber) {
		var ret string
		return ret
	}
	return *o.CarRentalCustomerServiceTollFreeNumber
}

// GetCarRentalCustomerServiceTollFreeNumberOk returns a tuple with the CarRentalCustomerServiceTollFreeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalCustomerServiceTollFreeNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalCustomerServiceTollFreeNumber) {
		return nil, false
	}
	return o.CarRentalCustomerServiceTollFreeNumber, true
}

// HasCarRentalCustomerServiceTollFreeNumber returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalCustomerServiceTollFreeNumber() bool {
	if o != nil && !common.IsNil(o.CarRentalCustomerServiceTollFreeNumber) {
		return true
	}

	return false
}

// SetCarRentalCustomerServiceTollFreeNumber gets a reference to the given string and assigns it to the CarRentalCustomerServiceTollFreeNumber field.
func (o *AdditionalDataCarRental) SetCarRentalCustomerServiceTollFreeNumber(v string) {
	o.CarRentalCustomerServiceTollFreeNumber = &v
}

// GetCarRentalDaysRented returns the CarRentalDaysRented field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalDaysRented() string {
	if o == nil || common.IsNil(o.CarRentalDaysRented) {
		var ret string
		return ret
	}
	return *o.CarRentalDaysRented
}

// GetCarRentalDaysRentedOk returns a tuple with the CarRentalDaysRented field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalDaysRentedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalDaysRented) {
		return nil, false
	}
	return o.CarRentalDaysRented, true
}

// HasCarRentalDaysRented returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalDaysRented() bool {
	if o != nil && !common.IsNil(o.CarRentalDaysRented) {
		return true
	}

	return false
}

// SetCarRentalDaysRented gets a reference to the given string and assigns it to the CarRentalDaysRented field.
func (o *AdditionalDataCarRental) SetCarRentalDaysRented(v string) {
	o.CarRentalDaysRented = &v
}

// GetCarRentalFuelCharges returns the CarRentalFuelCharges field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalFuelCharges() string {
	if o == nil || common.IsNil(o.CarRentalFuelCharges) {
		var ret string
		return ret
	}
	return *o.CarRentalFuelCharges
}

// GetCarRentalFuelChargesOk returns a tuple with the CarRentalFuelCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalFuelChargesOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalFuelCharges) {
		return nil, false
	}
	return o.CarRentalFuelCharges, true
}

// HasCarRentalFuelCharges returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalFuelCharges() bool {
	if o != nil && !common.IsNil(o.CarRentalFuelCharges) {
		return true
	}

	return false
}

// SetCarRentalFuelCharges gets a reference to the given string and assigns it to the CarRentalFuelCharges field.
func (o *AdditionalDataCarRental) SetCarRentalFuelCharges(v string) {
	o.CarRentalFuelCharges = &v
}

// GetCarRentalInsuranceCharges returns the CarRentalInsuranceCharges field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalInsuranceCharges() string {
	if o == nil || common.IsNil(o.CarRentalInsuranceCharges) {
		var ret string
		return ret
	}
	return *o.CarRentalInsuranceCharges
}

// GetCarRentalInsuranceChargesOk returns a tuple with the CarRentalInsuranceCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalInsuranceChargesOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalInsuranceCharges) {
		return nil, false
	}
	return o.CarRentalInsuranceCharges, true
}

// HasCarRentalInsuranceCharges returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalInsuranceCharges() bool {
	if o != nil && !common.IsNil(o.CarRentalInsuranceCharges) {
		return true
	}

	return false
}

// SetCarRentalInsuranceCharges gets a reference to the given string and assigns it to the CarRentalInsuranceCharges field.
func (o *AdditionalDataCarRental) SetCarRentalInsuranceCharges(v string) {
	o.CarRentalInsuranceCharges = &v
}

// GetCarRentalLocationCity returns the CarRentalLocationCity field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalLocationCity() string {
	if o == nil || common.IsNil(o.CarRentalLocationCity) {
		var ret string
		return ret
	}
	return *o.CarRentalLocationCity
}

// GetCarRentalLocationCityOk returns a tuple with the CarRentalLocationCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalLocationCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalLocationCity) {
		return nil, false
	}
	return o.CarRentalLocationCity, true
}

// HasCarRentalLocationCity returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalLocationCity() bool {
	if o != nil && !common.IsNil(o.CarRentalLocationCity) {
		return true
	}

	return false
}

// SetCarRentalLocationCity gets a reference to the given string and assigns it to the CarRentalLocationCity field.
func (o *AdditionalDataCarRental) SetCarRentalLocationCity(v string) {
	o.CarRentalLocationCity = &v
}

// GetCarRentalLocationCountry returns the CarRentalLocationCountry field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalLocationCountry() string {
	if o == nil || common.IsNil(o.CarRentalLocationCountry) {
		var ret string
		return ret
	}
	return *o.CarRentalLocationCountry
}

// GetCarRentalLocationCountryOk returns a tuple with the CarRentalLocationCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalLocationCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalLocationCountry) {
		return nil, false
	}
	return o.CarRentalLocationCountry, true
}

// HasCarRentalLocationCountry returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalLocationCountry() bool {
	if o != nil && !common.IsNil(o.CarRentalLocationCountry) {
		return true
	}

	return false
}

// SetCarRentalLocationCountry gets a reference to the given string and assigns it to the CarRentalLocationCountry field.
func (o *AdditionalDataCarRental) SetCarRentalLocationCountry(v string) {
	o.CarRentalLocationCountry = &v
}

// GetCarRentalLocationStateProvince returns the CarRentalLocationStateProvince field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalLocationStateProvince() string {
	if o == nil || common.IsNil(o.CarRentalLocationStateProvince) {
		var ret string
		return ret
	}
	return *o.CarRentalLocationStateProvince
}

// GetCarRentalLocationStateProvinceOk returns a tuple with the CarRentalLocationStateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalLocationStateProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalLocationStateProvince) {
		return nil, false
	}
	return o.CarRentalLocationStateProvince, true
}

// HasCarRentalLocationStateProvince returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalLocationStateProvince() bool {
	if o != nil && !common.IsNil(o.CarRentalLocationStateProvince) {
		return true
	}

	return false
}

// SetCarRentalLocationStateProvince gets a reference to the given string and assigns it to the CarRentalLocationStateProvince field.
func (o *AdditionalDataCarRental) SetCarRentalLocationStateProvince(v string) {
	o.CarRentalLocationStateProvince = &v
}

// GetCarRentalNoShowIndicator returns the CarRentalNoShowIndicator field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalNoShowIndicator() string {
	if o == nil || common.IsNil(o.CarRentalNoShowIndicator) {
		var ret string
		return ret
	}
	return *o.CarRentalNoShowIndicator
}

// GetCarRentalNoShowIndicatorOk returns a tuple with the CarRentalNoShowIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalNoShowIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalNoShowIndicator) {
		return nil, false
	}
	return o.CarRentalNoShowIndicator, true
}

// HasCarRentalNoShowIndicator returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalNoShowIndicator() bool {
	if o != nil && !common.IsNil(o.CarRentalNoShowIndicator) {
		return true
	}

	return false
}

// SetCarRentalNoShowIndicator gets a reference to the given string and assigns it to the CarRentalNoShowIndicator field.
func (o *AdditionalDataCarRental) SetCarRentalNoShowIndicator(v string) {
	o.CarRentalNoShowIndicator = &v
}

// GetCarRentalOneWayDropOffCharges returns the CarRentalOneWayDropOffCharges field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalOneWayDropOffCharges() string {
	if o == nil || common.IsNil(o.CarRentalOneWayDropOffCharges) {
		var ret string
		return ret
	}
	return *o.CarRentalOneWayDropOffCharges
}

// GetCarRentalOneWayDropOffChargesOk returns a tuple with the CarRentalOneWayDropOffCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalOneWayDropOffChargesOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalOneWayDropOffCharges) {
		return nil, false
	}
	return o.CarRentalOneWayDropOffCharges, true
}

// HasCarRentalOneWayDropOffCharges returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalOneWayDropOffCharges() bool {
	if o != nil && !common.IsNil(o.CarRentalOneWayDropOffCharges) {
		return true
	}

	return false
}

// SetCarRentalOneWayDropOffCharges gets a reference to the given string and assigns it to the CarRentalOneWayDropOffCharges field.
func (o *AdditionalDataCarRental) SetCarRentalOneWayDropOffCharges(v string) {
	o.CarRentalOneWayDropOffCharges = &v
}

// GetCarRentalRate returns the CarRentalRate field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalRate() string {
	if o == nil || common.IsNil(o.CarRentalRate) {
		var ret string
		return ret
	}
	return *o.CarRentalRate
}

// GetCarRentalRateOk returns a tuple with the CarRentalRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalRate) {
		return nil, false
	}
	return o.CarRentalRate, true
}

// HasCarRentalRate returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalRate() bool {
	if o != nil && !common.IsNil(o.CarRentalRate) {
		return true
	}

	return false
}

// SetCarRentalRate gets a reference to the given string and assigns it to the CarRentalRate field.
func (o *AdditionalDataCarRental) SetCarRentalRate(v string) {
	o.CarRentalRate = &v
}

// GetCarRentalRateIndicator returns the CarRentalRateIndicator field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalRateIndicator() string {
	if o == nil || common.IsNil(o.CarRentalRateIndicator) {
		var ret string
		return ret
	}
	return *o.CarRentalRateIndicator
}

// GetCarRentalRateIndicatorOk returns a tuple with the CarRentalRateIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalRateIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalRateIndicator) {
		return nil, false
	}
	return o.CarRentalRateIndicator, true
}

// HasCarRentalRateIndicator returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalRateIndicator() bool {
	if o != nil && !common.IsNil(o.CarRentalRateIndicator) {
		return true
	}

	return false
}

// SetCarRentalRateIndicator gets a reference to the given string and assigns it to the CarRentalRateIndicator field.
func (o *AdditionalDataCarRental) SetCarRentalRateIndicator(v string) {
	o.CarRentalRateIndicator = &v
}

// GetCarRentalRentalAgreementNumber returns the CarRentalRentalAgreementNumber field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalRentalAgreementNumber() string {
	if o == nil || common.IsNil(o.CarRentalRentalAgreementNumber) {
		var ret string
		return ret
	}
	return *o.CarRentalRentalAgreementNumber
}

// GetCarRentalRentalAgreementNumberOk returns a tuple with the CarRentalRentalAgreementNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalRentalAgreementNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalRentalAgreementNumber) {
		return nil, false
	}
	return o.CarRentalRentalAgreementNumber, true
}

// HasCarRentalRentalAgreementNumber returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalRentalAgreementNumber() bool {
	if o != nil && !common.IsNil(o.CarRentalRentalAgreementNumber) {
		return true
	}

	return false
}

// SetCarRentalRentalAgreementNumber gets a reference to the given string and assigns it to the CarRentalRentalAgreementNumber field.
func (o *AdditionalDataCarRental) SetCarRentalRentalAgreementNumber(v string) {
	o.CarRentalRentalAgreementNumber = &v
}

// GetCarRentalRentalClassId returns the CarRentalRentalClassId field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalRentalClassId() string {
	if o == nil || common.IsNil(o.CarRentalRentalClassId) {
		var ret string
		return ret
	}
	return *o.CarRentalRentalClassId
}

// GetCarRentalRentalClassIdOk returns a tuple with the CarRentalRentalClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalRentalClassIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalRentalClassId) {
		return nil, false
	}
	return o.CarRentalRentalClassId, true
}

// HasCarRentalRentalClassId returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalRentalClassId() bool {
	if o != nil && !common.IsNil(o.CarRentalRentalClassId) {
		return true
	}

	return false
}

// SetCarRentalRentalClassId gets a reference to the given string and assigns it to the CarRentalRentalClassId field.
func (o *AdditionalDataCarRental) SetCarRentalRentalClassId(v string) {
	o.CarRentalRentalClassId = &v
}

// GetCarRentalRenterName returns the CarRentalRenterName field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalRenterName() string {
	if o == nil || common.IsNil(o.CarRentalRenterName) {
		var ret string
		return ret
	}
	return *o.CarRentalRenterName
}

// GetCarRentalRenterNameOk returns a tuple with the CarRentalRenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalRenterNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalRenterName) {
		return nil, false
	}
	return o.CarRentalRenterName, true
}

// HasCarRentalRenterName returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalRenterName() bool {
	if o != nil && !common.IsNil(o.CarRentalRenterName) {
		return true
	}

	return false
}

// SetCarRentalRenterName gets a reference to the given string and assigns it to the CarRentalRenterName field.
func (o *AdditionalDataCarRental) SetCarRentalRenterName(v string) {
	o.CarRentalRenterName = &v
}

// GetCarRentalReturnCity returns the CarRentalReturnCity field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalReturnCity() string {
	if o == nil || common.IsNil(o.CarRentalReturnCity) {
		var ret string
		return ret
	}
	return *o.CarRentalReturnCity
}

// GetCarRentalReturnCityOk returns a tuple with the CarRentalReturnCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalReturnCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalReturnCity) {
		return nil, false
	}
	return o.CarRentalReturnCity, true
}

// HasCarRentalReturnCity returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalReturnCity() bool {
	if o != nil && !common.IsNil(o.CarRentalReturnCity) {
		return true
	}

	return false
}

// SetCarRentalReturnCity gets a reference to the given string and assigns it to the CarRentalReturnCity field.
func (o *AdditionalDataCarRental) SetCarRentalReturnCity(v string) {
	o.CarRentalReturnCity = &v
}

// GetCarRentalReturnCountry returns the CarRentalReturnCountry field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalReturnCountry() string {
	if o == nil || common.IsNil(o.CarRentalReturnCountry) {
		var ret string
		return ret
	}
	return *o.CarRentalReturnCountry
}

// GetCarRentalReturnCountryOk returns a tuple with the CarRentalReturnCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalReturnCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalReturnCountry) {
		return nil, false
	}
	return o.CarRentalReturnCountry, true
}

// HasCarRentalReturnCountry returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalReturnCountry() bool {
	if o != nil && !common.IsNil(o.CarRentalReturnCountry) {
		return true
	}

	return false
}

// SetCarRentalReturnCountry gets a reference to the given string and assigns it to the CarRentalReturnCountry field.
func (o *AdditionalDataCarRental) SetCarRentalReturnCountry(v string) {
	o.CarRentalReturnCountry = &v
}

// GetCarRentalReturnDate returns the CarRentalReturnDate field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalReturnDate() string {
	if o == nil || common.IsNil(o.CarRentalReturnDate) {
		var ret string
		return ret
	}
	return *o.CarRentalReturnDate
}

// GetCarRentalReturnDateOk returns a tuple with the CarRentalReturnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalReturnDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalReturnDate) {
		return nil, false
	}
	return o.CarRentalReturnDate, true
}

// HasCarRentalReturnDate returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalReturnDate() bool {
	if o != nil && !common.IsNil(o.CarRentalReturnDate) {
		return true
	}

	return false
}

// SetCarRentalReturnDate gets a reference to the given string and assigns it to the CarRentalReturnDate field.
func (o *AdditionalDataCarRental) SetCarRentalReturnDate(v string) {
	o.CarRentalReturnDate = &v
}

// GetCarRentalReturnLocationId returns the CarRentalReturnLocationId field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalReturnLocationId() string {
	if o == nil || common.IsNil(o.CarRentalReturnLocationId) {
		var ret string
		return ret
	}
	return *o.CarRentalReturnLocationId
}

// GetCarRentalReturnLocationIdOk returns a tuple with the CarRentalReturnLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalReturnLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalReturnLocationId) {
		return nil, false
	}
	return o.CarRentalReturnLocationId, true
}

// HasCarRentalReturnLocationId returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalReturnLocationId() bool {
	if o != nil && !common.IsNil(o.CarRentalReturnLocationId) {
		return true
	}

	return false
}

// SetCarRentalReturnLocationId gets a reference to the given string and assigns it to the CarRentalReturnLocationId field.
func (o *AdditionalDataCarRental) SetCarRentalReturnLocationId(v string) {
	o.CarRentalReturnLocationId = &v
}

// GetCarRentalReturnStateProvince returns the CarRentalReturnStateProvince field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalReturnStateProvince() string {
	if o == nil || common.IsNil(o.CarRentalReturnStateProvince) {
		var ret string
		return ret
	}
	return *o.CarRentalReturnStateProvince
}

// GetCarRentalReturnStateProvinceOk returns a tuple with the CarRentalReturnStateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalReturnStateProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalReturnStateProvince) {
		return nil, false
	}
	return o.CarRentalReturnStateProvince, true
}

// HasCarRentalReturnStateProvince returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalReturnStateProvince() bool {
	if o != nil && !common.IsNil(o.CarRentalReturnStateProvince) {
		return true
	}

	return false
}

// SetCarRentalReturnStateProvince gets a reference to the given string and assigns it to the CarRentalReturnStateProvince field.
func (o *AdditionalDataCarRental) SetCarRentalReturnStateProvince(v string) {
	o.CarRentalReturnStateProvince = &v
}

// GetCarRentalTaxExemptIndicator returns the CarRentalTaxExemptIndicator field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetCarRentalTaxExemptIndicator() string {
	if o == nil || common.IsNil(o.CarRentalTaxExemptIndicator) {
		var ret string
		return ret
	}
	return *o.CarRentalTaxExemptIndicator
}

// GetCarRentalTaxExemptIndicatorOk returns a tuple with the CarRentalTaxExemptIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetCarRentalTaxExemptIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.CarRentalTaxExemptIndicator) {
		return nil, false
	}
	return o.CarRentalTaxExemptIndicator, true
}

// HasCarRentalTaxExemptIndicator returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasCarRentalTaxExemptIndicator() bool {
	if o != nil && !common.IsNil(o.CarRentalTaxExemptIndicator) {
		return true
	}

	return false
}

// SetCarRentalTaxExemptIndicator gets a reference to the given string and assigns it to the CarRentalTaxExemptIndicator field.
func (o *AdditionalDataCarRental) SetCarRentalTaxExemptIndicator(v string) {
	o.CarRentalTaxExemptIndicator = &v
}

// GetTravelEntertainmentAuthDataDuration returns the TravelEntertainmentAuthDataDuration field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataDuration() string {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		var ret string
		return ret
	}
	return *o.TravelEntertainmentAuthDataDuration
}

// GetTravelEntertainmentAuthDataDurationOk returns a tuple with the TravelEntertainmentAuthDataDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataDurationOk() (*string, bool) {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		return nil, false
	}
	return o.TravelEntertainmentAuthDataDuration, true
}

// HasTravelEntertainmentAuthDataDuration returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasTravelEntertainmentAuthDataDuration() bool {
	if o != nil && !common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		return true
	}

	return false
}

// SetTravelEntertainmentAuthDataDuration gets a reference to the given string and assigns it to the TravelEntertainmentAuthDataDuration field.
func (o *AdditionalDataCarRental) SetTravelEntertainmentAuthDataDuration(v string) {
	o.TravelEntertainmentAuthDataDuration = &v
}

// GetTravelEntertainmentAuthDataMarket returns the TravelEntertainmentAuthDataMarket field value if set, zero value otherwise.
func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataMarket() string {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		var ret string
		return ret
	}
	return *o.TravelEntertainmentAuthDataMarket
}

// GetTravelEntertainmentAuthDataMarketOk returns a tuple with the TravelEntertainmentAuthDataMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataMarketOk() (*string, bool) {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		return nil, false
	}
	return o.TravelEntertainmentAuthDataMarket, true
}

// HasTravelEntertainmentAuthDataMarket returns a boolean if a field has been set.
func (o *AdditionalDataCarRental) HasTravelEntertainmentAuthDataMarket() bool {
	if o != nil && !common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		return true
	}

	return false
}

// SetTravelEntertainmentAuthDataMarket gets a reference to the given string and assigns it to the TravelEntertainmentAuthDataMarket field.
func (o *AdditionalDataCarRental) SetTravelEntertainmentAuthDataMarket(v string) {
	o.TravelEntertainmentAuthDataMarket = &v
}

func (o AdditionalDataCarRental) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataCarRental) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CarRentalCheckOutDate) {
		toSerialize["carRental.checkOutDate"] = o.CarRentalCheckOutDate
	}
	if !common.IsNil(o.CarRentalCustomerServiceTollFreeNumber) {
		toSerialize["carRental.customerServiceTollFreeNumber"] = o.CarRentalCustomerServiceTollFreeNumber
	}
	if !common.IsNil(o.CarRentalDaysRented) {
		toSerialize["carRental.daysRented"] = o.CarRentalDaysRented
	}
	if !common.IsNil(o.CarRentalFuelCharges) {
		toSerialize["carRental.fuelCharges"] = o.CarRentalFuelCharges
	}
	if !common.IsNil(o.CarRentalInsuranceCharges) {
		toSerialize["carRental.insuranceCharges"] = o.CarRentalInsuranceCharges
	}
	if !common.IsNil(o.CarRentalLocationCity) {
		toSerialize["carRental.locationCity"] = o.CarRentalLocationCity
	}
	if !common.IsNil(o.CarRentalLocationCountry) {
		toSerialize["carRental.locationCountry"] = o.CarRentalLocationCountry
	}
	if !common.IsNil(o.CarRentalLocationStateProvince) {
		toSerialize["carRental.locationStateProvince"] = o.CarRentalLocationStateProvince
	}
	if !common.IsNil(o.CarRentalNoShowIndicator) {
		toSerialize["carRental.noShowIndicator"] = o.CarRentalNoShowIndicator
	}
	if !common.IsNil(o.CarRentalOneWayDropOffCharges) {
		toSerialize["carRental.oneWayDropOffCharges"] = o.CarRentalOneWayDropOffCharges
	}
	if !common.IsNil(o.CarRentalRate) {
		toSerialize["carRental.rate"] = o.CarRentalRate
	}
	if !common.IsNil(o.CarRentalRateIndicator) {
		toSerialize["carRental.rateIndicator"] = o.CarRentalRateIndicator
	}
	if !common.IsNil(o.CarRentalRentalAgreementNumber) {
		toSerialize["carRental.rentalAgreementNumber"] = o.CarRentalRentalAgreementNumber
	}
	if !common.IsNil(o.CarRentalRentalClassId) {
		toSerialize["carRental.rentalClassId"] = o.CarRentalRentalClassId
	}
	if !common.IsNil(o.CarRentalRenterName) {
		toSerialize["carRental.renterName"] = o.CarRentalRenterName
	}
	if !common.IsNil(o.CarRentalReturnCity) {
		toSerialize["carRental.returnCity"] = o.CarRentalReturnCity
	}
	if !common.IsNil(o.CarRentalReturnCountry) {
		toSerialize["carRental.returnCountry"] = o.CarRentalReturnCountry
	}
	if !common.IsNil(o.CarRentalReturnDate) {
		toSerialize["carRental.returnDate"] = o.CarRentalReturnDate
	}
	if !common.IsNil(o.CarRentalReturnLocationId) {
		toSerialize["carRental.returnLocationId"] = o.CarRentalReturnLocationId
	}
	if !common.IsNil(o.CarRentalReturnStateProvince) {
		toSerialize["carRental.returnStateProvince"] = o.CarRentalReturnStateProvince
	}
	if !common.IsNil(o.CarRentalTaxExemptIndicator) {
		toSerialize["carRental.taxExemptIndicator"] = o.CarRentalTaxExemptIndicator
	}
	if !common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		toSerialize["travelEntertainmentAuthData.duration"] = o.TravelEntertainmentAuthDataDuration
	}
	if !common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		toSerialize["travelEntertainmentAuthData.market"] = o.TravelEntertainmentAuthDataMarket
	}
	return toSerialize, nil
}

type NullableAdditionalDataCarRental struct {
	value *AdditionalDataCarRental
	isSet bool
}

func (v NullableAdditionalDataCarRental) Get() *AdditionalDataCarRental {
	return v.value
}

func (v *NullableAdditionalDataCarRental) Set(val *AdditionalDataCarRental) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataCarRental) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataCarRental) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataCarRental(val *AdditionalDataCarRental) *NullableAdditionalDataCarRental {
	return &NullableAdditionalDataCarRental{value: val, isSet: true}
}

func (v NullableAdditionalDataCarRental) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataCarRental) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
