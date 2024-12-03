/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v16/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.Get3dsAvailability
type GeneralApiGet3dsAvailabilityInput struct {
	threeDSAvailabilityRequest *ThreeDSAvailabilityRequest
}

func (r GeneralApiGet3dsAvailabilityInput) ThreeDSAvailabilityRequest(threeDSAvailabilityRequest ThreeDSAvailabilityRequest) GeneralApiGet3dsAvailabilityInput {
	r.threeDSAvailabilityRequest = &threeDSAvailabilityRequest
	return r
}

/*
Prepare a request for Get3dsAvailability

@return GeneralApiGet3dsAvailabilityInput
*/
func (a *GeneralApi) Get3dsAvailabilityInput() GeneralApiGet3dsAvailabilityInput {
	return GeneralApiGet3dsAvailabilityInput{}
}

/*
Get3dsAvailability Check if 3D Secure is available

Verifies whether 3D Secure is available for the specified BIN or card brand. For 3D Secure 2, this endpoint also returns device fingerprinting keys.

For more information, refer to [3D Secure 2](https://docs.adyen.com/online-payments/3d-secure/native-3ds2).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiGet3dsAvailabilityInput - Request parameters, see Get3dsAvailabilityInput
@return ThreeDSAvailabilityResponse, *http.Response, error
*/
func (a *GeneralApi) Get3dsAvailability(ctx context.Context, r GeneralApiGet3dsAvailabilityInput) (ThreeDSAvailabilityResponse, *http.Response, error) {
	res := &ThreeDSAvailabilityResponse{}
	path := "/get3dsAvailability"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.threeDSAvailabilityRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.GetCostEstimate
type GeneralApiGetCostEstimateInput struct {
	costEstimateRequest *CostEstimateRequest
}

func (r GeneralApiGetCostEstimateInput) CostEstimateRequest(costEstimateRequest CostEstimateRequest) GeneralApiGetCostEstimateInput {
	r.costEstimateRequest = &costEstimateRequest
	return r
}

/*
Prepare a request for GetCostEstimate

@return GeneralApiGetCostEstimateInput
*/
func (a *GeneralApi) GetCostEstimateInput() GeneralApiGetCostEstimateInput {
	return GeneralApiGetCostEstimateInput{}
}

/*
GetCostEstimate Get a fees cost estimate

>This API is available only for merchants operating in Australia, the EU, and the UK.

Use the Adyen Cost Estimation API to pre-calculate interchange and scheme fee costs. Knowing these costs prior actual payment authorisation gives you an opportunity to charge those costs to the cardholder, if necessary.

To retrieve this information, make the call to the `/getCostEstimate` endpoint. The response to this call contains the amount of the interchange and scheme fees charged by the network for this transaction, and also which surcharging policy is possible (based on current regulations).

> Since not all information is known in advance (for example, if the cardholder will successfully authenticate via 3D Secure or if you also plan to provide additional Level 2/3 data), the returned amounts are based on a set of assumption criteria you define in the `assumptions` parameter.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiGetCostEstimateInput - Request parameters, see GetCostEstimateInput
@return CostEstimateResponse, *http.Response, error
*/
func (a *GeneralApi) GetCostEstimate(ctx context.Context, r GeneralApiGetCostEstimateInput) (CostEstimateResponse, *http.Response, error) {
	res := &CostEstimateResponse{}
	path := "/getCostEstimate"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.costEstimateRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
