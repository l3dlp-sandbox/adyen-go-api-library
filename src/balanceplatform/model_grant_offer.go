/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// checks if the GrantOffer type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GrantOffer{}

// GrantOffer struct for GrantOffer
type GrantOffer struct {
	// The identifier of the account holder to which the grant is offered.
	AccountHolderId string  `json:"accountHolderId"`
	Amount          *Amount `json:"amount,omitempty"`
	// The contract type of the grant offer. Possible value: **cashAdvance**, **loan**.
	ContractType *string `json:"contractType,omitempty"`
	// The end date of the grant offer validity period.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Fee       *Fee       `json:"fee,omitempty"`
	// The unique identifier of the grant offer.
	Id        *string    `json:"id,omitempty"`
	Repayment *Repayment `json:"repayment,omitempty"`
	// The starting date of the grant offer validity period.
	StartsAt *time.Time `json:"startsAt,omitempty"`
}

// NewGrantOffer instantiates a new GrantOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantOffer(accountHolderId string) *GrantOffer {
	this := GrantOffer{}
	this.AccountHolderId = accountHolderId
	return &this
}

// NewGrantOfferWithDefaults instantiates a new GrantOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantOfferWithDefaults() *GrantOffer {
	this := GrantOffer{}
	return &this
}

// GetAccountHolderId returns the AccountHolderId field value
func (o *GrantOffer) GetAccountHolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountHolderId
}

// GetAccountHolderIdOk returns a tuple with the AccountHolderId field value
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetAccountHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolderId, true
}

// SetAccountHolderId sets field value
func (o *GrantOffer) SetAccountHolderId(v string) {
	o.AccountHolderId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GrantOffer) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GrantOffer) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *GrantOffer) SetAmount(v Amount) {
	o.Amount = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *GrantOffer) GetContractType() string {
	if o == nil || common.IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetContractTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *GrantOffer) HasContractType() bool {
	if o != nil && !common.IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *GrantOffer) SetContractType(v string) {
	o.ContractType = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GrantOffer) GetExpiresAt() time.Time {
	if o == nil || common.IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GrantOffer) HasExpiresAt() bool {
	if o != nil && !common.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *GrantOffer) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GrantOffer) GetFee() Fee {
	if o == nil || common.IsNil(o.Fee) {
		var ret Fee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetFeeOk() (*Fee, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GrantOffer) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given Fee and assigns it to the Fee field.
func (o *GrantOffer) SetFee(v Fee) {
	o.Fee = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GrantOffer) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GrantOffer) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GrantOffer) SetId(v string) {
	o.Id = &v
}

// GetRepayment returns the Repayment field value if set, zero value otherwise.
func (o *GrantOffer) GetRepayment() Repayment {
	if o == nil || common.IsNil(o.Repayment) {
		var ret Repayment
		return ret
	}
	return *o.Repayment
}

// GetRepaymentOk returns a tuple with the Repayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetRepaymentOk() (*Repayment, bool) {
	if o == nil || common.IsNil(o.Repayment) {
		return nil, false
	}
	return o.Repayment, true
}

// HasRepayment returns a boolean if a field has been set.
func (o *GrantOffer) HasRepayment() bool {
	if o != nil && !common.IsNil(o.Repayment) {
		return true
	}

	return false
}

// SetRepayment gets a reference to the given Repayment and assigns it to the Repayment field.
func (o *GrantOffer) SetRepayment(v Repayment) {
	o.Repayment = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *GrantOffer) GetStartsAt() time.Time {
	if o == nil || common.IsNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantOffer) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GrantOffer) HasStartsAt() bool {
	if o != nil && !common.IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *GrantOffer) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

func (o GrantOffer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountHolderId"] = o.AccountHolderId
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !common.IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Repayment) {
		toSerialize["repayment"] = o.Repayment
	}
	if !common.IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	return toSerialize, nil
}

type NullableGrantOffer struct {
	value *GrantOffer
	isSet bool
}

func (v NullableGrantOffer) Get() *GrantOffer {
	return v.value
}

func (v *NullableGrantOffer) Set(val *GrantOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantOffer(val *GrantOffer) *NullableGrantOffer {
	return &NullableGrantOffer{value: val, isSet: true}
}

func (v NullableGrantOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *GrantOffer) isValidContractType() bool {
	var allowedEnumValues = []string{"cashAdvance", "loan"}
	for _, allowed := range allowedEnumValues {
		if o.GetContractType() == allowed {
			return true
		}
	}
	return false
}
