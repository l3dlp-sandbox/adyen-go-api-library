/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
	"time"
	"fmt"
)

// TransferDataTracking - The latest tracking information of the transfer.
type TransferDataTracking struct {
	ConfirmationTrackingData *ConfirmationTrackingData
	EstimationTrackingData *EstimationTrackingData
	InternalReviewTrackingData *InternalReviewTrackingData
}

// ConfirmationTrackingDataAsTransferDataTracking is a convenience function that returns ConfirmationTrackingData wrapped in TransferDataTracking
func ConfirmationTrackingDataAsTransferDataTracking(v *ConfirmationTrackingData) TransferDataTracking {
	return TransferDataTracking{
		ConfirmationTrackingData: v,
	}
}

// EstimationTrackingDataAsTransferDataTracking is a convenience function that returns EstimationTrackingData wrapped in TransferDataTracking
func EstimationTrackingDataAsTransferDataTracking(v *EstimationTrackingData) TransferDataTracking {
	return TransferDataTracking{
		EstimationTrackingData: v,
	}
}

// InternalReviewTrackingDataAsTransferDataTracking is a convenience function that returns InternalReviewTrackingData wrapped in TransferDataTracking
func InternalReviewTrackingDataAsTransferDataTracking(v *InternalReviewTrackingData) TransferDataTracking {
	return TransferDataTracking{
		InternalReviewTrackingData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransferDataTracking) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ConfirmationTrackingData
	err = json.Unmarshal(data, &dst.ConfirmationTrackingData)
	if err == nil {
		jsonConfirmationTrackingData, _ := json.Marshal(dst.ConfirmationTrackingData)
		if string(jsonConfirmationTrackingData) == "{}" || !dst.ConfirmationTrackingData.isValidType() { // empty struct
			dst.ConfirmationTrackingData = nil
        } else {
			match++
		}
	} else {
		dst.ConfirmationTrackingData = nil
	}

	// try to unmarshal data into EstimationTrackingData
	err = json.Unmarshal(data, &dst.EstimationTrackingData)
	if err == nil {
		jsonEstimationTrackingData, _ := json.Marshal(dst.EstimationTrackingData)
		if string(jsonEstimationTrackingData) == "{}" || !dst.EstimationTrackingData.isValidType() { // empty struct
			dst.EstimationTrackingData = nil
        } else {
			match++
		}
	} else {
		dst.EstimationTrackingData = nil
	}

	// try to unmarshal data into InternalReviewTrackingData
	err = json.Unmarshal(data, &dst.InternalReviewTrackingData)
	if err == nil {
		jsonInternalReviewTrackingData, _ := json.Marshal(dst.InternalReviewTrackingData)
		if string(jsonInternalReviewTrackingData) == "{}" || !dst.InternalReviewTrackingData.isValidType() { // empty struct
			dst.InternalReviewTrackingData = nil
        } else {
			match++
		}
	} else {
		dst.InternalReviewTrackingData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ConfirmationTrackingData = nil
		dst.EstimationTrackingData = nil
		dst.InternalReviewTrackingData = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransferDataTracking)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransferDataTracking)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransferDataTracking) MarshalJSON() ([]byte, error) {
	if src.ConfirmationTrackingData != nil {
		return json.Marshal(&src.ConfirmationTrackingData)
	}

	if src.EstimationTrackingData != nil {
		return json.Marshal(&src.EstimationTrackingData)
	}

	if src.InternalReviewTrackingData != nil {
		return json.Marshal(&src.InternalReviewTrackingData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransferDataTracking) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ConfirmationTrackingData != nil {
		return obj.ConfirmationTrackingData
	}

	if obj.EstimationTrackingData != nil {
		return obj.EstimationTrackingData
	}

	if obj.InternalReviewTrackingData != nil {
		return obj.InternalReviewTrackingData
	}

	// all schemas are nil
	return nil
}

type NullableTransferDataTracking struct {
	value *TransferDataTracking
	isSet bool
}

func (v NullableTransferDataTracking) Get() *TransferDataTracking {
	return v.value
}

func (v *NullableTransferDataTracking) Set(val *TransferDataTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDataTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDataTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDataTracking(val *TransferDataTracking) *NullableTransferDataTracking {
	return &NullableTransferDataTracking{value: val, isSet: true}
}

func (v NullableTransferDataTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDataTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


