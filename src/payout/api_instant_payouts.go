/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// InstantPayoutsApi service
type InstantPayoutsApi common.Service

// All parameters accepted by InstantPayoutsApi.Payout
type InstantPayoutsApiPayoutInput struct {
	payoutRequest *PayoutRequest
}

func (r InstantPayoutsApiPayoutInput) PayoutRequest(payoutRequest PayoutRequest) InstantPayoutsApiPayoutInput {
	r.payoutRequest = &payoutRequest
	return r
}

/*
Prepare a request for Payout

@return InstantPayoutsApiPayoutInput
*/
func (a *InstantPayoutsApi) PayoutInput() InstantPayoutsApiPayoutInput {
	return InstantPayoutsApiPayoutInput{}
}

/*
Payout Make an instant card payout

With this call, you can pay out to your customers, and funds will be made available within 30 minutes on the cardholder's bank account (this is dependent on whether the issuer supports this functionality). Instant card payouts are only supported for Visa and Mastercard cards.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r InstantPayoutsApiPayoutInput - Request parameters, see PayoutInput
@return PayoutResponse, *http.Response, error
*/
func (a *InstantPayoutsApi) Payout(ctx context.Context, r InstantPayoutsApiPayoutInput) (PayoutResponse, *http.Response, error) {
	res := &PayoutResponse{}
	path := "/payout"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.payoutRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}