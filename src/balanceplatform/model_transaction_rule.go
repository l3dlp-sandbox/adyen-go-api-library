/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// checks if the TransactionRule type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRule{}

// TransactionRule struct for TransactionRule
type TransactionRule struct {
	// The level at which data must be accumulated, used in rules with `type` **velocity** or **maxUsage**. The level must be the [same or lower in hierarchy](https://docs.adyen.com/issuing/transaction-rules#accumulate-data) than the `entityKey`.  If not provided, by default, the rule will accumulate data at the **paymentInstrument** level.  Possible values: **paymentInstrument**, **paymentInstrumentGroup**, **balanceAccount**, **accountHolder**, **balancePlatform**.
	AggregationLevel *string `json:"aggregationLevel,omitempty"`
	// Your description for the transaction rule.
	Description string `json:"description"`
	// The date when the rule will stop being evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.  If not provided, the rule will be evaluated until the rule status is set to **inactive**.
	EndDate   *string                  `json:"endDate,omitempty"`
	EntityKey TransactionRuleEntityKey `json:"entityKey"`
	// The unique identifier of the transaction rule.
	Id       *string                 `json:"id,omitempty"`
	Interval TransactionRuleInterval `json:"interval"`
	// The [outcome](https://docs.adyen.com/issuing/transaction-rules#outcome) that will be applied when a transaction meets the conditions of the rule. If not provided, by default, this is set to **hardBlock**.  Possible values:   * **hardBlock**: the transaction is declined.  * **scoreBased**: the transaction is assigned the `score` you specified. Adyen calculates the total score and if it exceeds 100, the transaction is declined.
	OutcomeType *string `json:"outcomeType,omitempty"`
	// Your reference for the transaction rule.
	Reference string `json:"reference"`
	// Indicates the type of request to which the rule applies. If not provided, by default, this is set to **authorization**.  Possible values: **authorization**, **authentication**, **tokenization**, **bankTransfer**.
	RequestType      *string                     `json:"requestType,omitempty"`
	RuleRestrictions TransactionRuleRestrictions `json:"ruleRestrictions"`
	// A positive or negative score applied to the transaction if it meets the conditions of the rule. Required when `outcomeType` is **scoreBased**.  The value must be between **-100** and **100**.
	Score *int32 `json:"score,omitempty"`
	// The date when the rule will start to be evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.  If not provided when creating a transaction rule, the `startDate` is set to the date when the rule status is set to **active**.
	StartDate *string `json:"startDate,omitempty"`
	// The status of the transaction rule. If you provide a `startDate` in the request, the rule is automatically created  with an **active** status.   Possible values: **active**, **inactive**.
	Status *string `json:"status,omitempty"`
	// The [type of rule](https://docs.adyen.com/issuing/transaction-rules#rule-types), which defines if a rule blocks transactions based on individual characteristics or accumulates data.  Possible values:  * **blockList**: decline a transaction when the conditions are met.  * **maxUsage**: add the amount or number of transactions for the lifetime of a payment instrument, and then decline a transaction when the specified limits are met.  * **velocity**: add the amount or number of transactions based on a specified time interval, and then decline a transaction when the specified limits are met.
	Type string `json:"type"`
}

// NewTransactionRule instantiates a new TransactionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRule(description string, entityKey TransactionRuleEntityKey, interval TransactionRuleInterval, reference string, ruleRestrictions TransactionRuleRestrictions, type_ string) *TransactionRule {
	this := TransactionRule{}
	this.Description = description
	this.EntityKey = entityKey
	this.Interval = interval
	this.Reference = reference
	this.RuleRestrictions = ruleRestrictions
	this.Type = type_
	return &this
}

// NewTransactionRuleWithDefaults instantiates a new TransactionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleWithDefaults() *TransactionRule {
	this := TransactionRule{}
	return &this
}

// GetAggregationLevel returns the AggregationLevel field value if set, zero value otherwise.
func (o *TransactionRule) GetAggregationLevel() string {
	if o == nil || common.IsNil(o.AggregationLevel) {
		var ret string
		return ret
	}
	return *o.AggregationLevel
}

// GetAggregationLevelOk returns a tuple with the AggregationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetAggregationLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.AggregationLevel) {
		return nil, false
	}
	return o.AggregationLevel, true
}

// HasAggregationLevel returns a boolean if a field has been set.
func (o *TransactionRule) HasAggregationLevel() bool {
	if o != nil && !common.IsNil(o.AggregationLevel) {
		return true
	}

	return false
}

// SetAggregationLevel gets a reference to the given string and assigns it to the AggregationLevel field.
func (o *TransactionRule) SetAggregationLevel(v string) {
	o.AggregationLevel = &v
}

// GetDescription returns the Description field value
func (o *TransactionRule) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransactionRule) SetDescription(v string) {
	o.Description = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TransactionRule) GetEndDate() string {
	if o == nil || common.IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetEndDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TransactionRule) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *TransactionRule) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEntityKey returns the EntityKey field value
func (o *TransactionRule) GetEntityKey() TransactionRuleEntityKey {
	if o == nil {
		var ret TransactionRuleEntityKey
		return ret
	}

	return o.EntityKey
}

// GetEntityKeyOk returns a tuple with the EntityKey field value
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetEntityKeyOk() (*TransactionRuleEntityKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityKey, true
}

// SetEntityKey sets field value
func (o *TransactionRule) SetEntityKey(v TransactionRuleEntityKey) {
	o.EntityKey = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionRule) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionRule) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionRule) SetId(v string) {
	o.Id = &v
}

// GetInterval returns the Interval field value
func (o *TransactionRule) GetInterval() TransactionRuleInterval {
	if o == nil {
		var ret TransactionRuleInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetIntervalOk() (*TransactionRuleInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *TransactionRule) SetInterval(v TransactionRuleInterval) {
	o.Interval = v
}

// GetOutcomeType returns the OutcomeType field value if set, zero value otherwise.
func (o *TransactionRule) GetOutcomeType() string {
	if o == nil || common.IsNil(o.OutcomeType) {
		var ret string
		return ret
	}
	return *o.OutcomeType
}

// GetOutcomeTypeOk returns a tuple with the OutcomeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetOutcomeTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OutcomeType) {
		return nil, false
	}
	return o.OutcomeType, true
}

// HasOutcomeType returns a boolean if a field has been set.
func (o *TransactionRule) HasOutcomeType() bool {
	if o != nil && !common.IsNil(o.OutcomeType) {
		return true
	}

	return false
}

// SetOutcomeType gets a reference to the given string and assigns it to the OutcomeType field.
func (o *TransactionRule) SetOutcomeType(v string) {
	o.OutcomeType = &v
}

// GetReference returns the Reference field value
func (o *TransactionRule) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *TransactionRule) SetReference(v string) {
	o.Reference = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *TransactionRule) GetRequestType() string {
	if o == nil || common.IsNil(o.RequestType) {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetRequestTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *TransactionRule) HasRequestType() bool {
	if o != nil && !common.IsNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *TransactionRule) SetRequestType(v string) {
	o.RequestType = &v
}

// GetRuleRestrictions returns the RuleRestrictions field value
func (o *TransactionRule) GetRuleRestrictions() TransactionRuleRestrictions {
	if o == nil {
		var ret TransactionRuleRestrictions
		return ret
	}

	return o.RuleRestrictions
}

// GetRuleRestrictionsOk returns a tuple with the RuleRestrictions field value
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetRuleRestrictionsOk() (*TransactionRuleRestrictions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleRestrictions, true
}

// SetRuleRestrictions sets field value
func (o *TransactionRule) SetRuleRestrictions(v TransactionRuleRestrictions) {
	o.RuleRestrictions = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *TransactionRule) GetScore() int32 {
	if o == nil || common.IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetScoreOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *TransactionRule) HasScore() bool {
	if o != nil && !common.IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *TransactionRule) SetScore(v int32) {
	o.Score = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TransactionRule) GetStartDate() string {
	if o == nil || common.IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetStartDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TransactionRule) HasStartDate() bool {
	if o != nil && !common.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TransactionRule) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionRule) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionRule) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransactionRule) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value
func (o *TransactionRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionRule) SetType(v string) {
	o.Type = v
}

func (o TransactionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AggregationLevel) {
		toSerialize["aggregationLevel"] = o.AggregationLevel
	}
	toSerialize["description"] = o.Description
	if !common.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["entityKey"] = o.EntityKey
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["interval"] = o.Interval
	if !common.IsNil(o.OutcomeType) {
		toSerialize["outcomeType"] = o.OutcomeType
	}
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.RequestType) {
		toSerialize["requestType"] = o.RequestType
	}
	toSerialize["ruleRestrictions"] = o.RuleRestrictions
	if !common.IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !common.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransactionRule struct {
	value *TransactionRule
	isSet bool
}

func (v NullableTransactionRule) Get() *TransactionRule {
	return v.value
}

func (v *NullableTransactionRule) Set(val *TransactionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRule(val *TransactionRule) *NullableTransactionRule {
	return &NullableTransactionRule{value: val, isSet: true}
}

func (v NullableTransactionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransactionRule) isValidOutcomeType() bool {
	var allowedEnumValues = []string{"enforceSCA", "hardBlock", "scoreBased", "timedBlock"}
	for _, allowed := range allowedEnumValues {
		if o.GetOutcomeType() == allowed {
			return true
		}
	}
	return false
}
func (o *TransactionRule) isValidRequestType() bool {
	var allowedEnumValues = []string{"authentication", "authorization", "bankTransfer", "tokenization"}
	for _, allowed := range allowedEnumValues {
		if o.GetRequestType() == allowed {
			return true
		}
	}
	return false
}
func (o *TransactionRule) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "inactive"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *TransactionRule) isValidType() bool {
	var allowedEnumValues = []string{"allowList", "blockList", "maxUsage", "velocity"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
