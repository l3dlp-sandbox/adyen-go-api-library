/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// DonationsApi service
type DonationsApi common.Service

// All parameters accepted by DonationsApi.DonationCampaigns
type DonationsApiDonationCampaignsInput struct {
	idempotencyKey           *string
	donationCampaignsRequest *DonationCampaignsRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r DonationsApiDonationCampaignsInput) IdempotencyKey(idempotencyKey string) DonationsApiDonationCampaignsInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r DonationsApiDonationCampaignsInput) DonationCampaignsRequest(donationCampaignsRequest DonationCampaignsRequest) DonationsApiDonationCampaignsInput {
	r.donationCampaignsRequest = &donationCampaignsRequest
	return r
}

/*
Prepare a request for DonationCampaigns

@return DonationsApiDonationCampaignsInput
*/
func (a *DonationsApi) DonationCampaignsInput() DonationsApiDonationCampaignsInput {
	return DonationsApiDonationCampaignsInput{}
}

/*
DonationCampaigns Get a list of donation campaigns.

Queries the available donation campaigns for a donation based on the donation context (like merchant account, currency, and locale). The response contains active donation campaigns.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r DonationsApiDonationCampaignsInput - Request parameters, see DonationCampaignsInput
@return DonationCampaignsResponse, *http.Response, error
*/
func (a *DonationsApi) DonationCampaigns(ctx context.Context, r DonationsApiDonationCampaignsInput) (DonationCampaignsResponse, *http.Response, error) {
	res := &DonationCampaignsResponse{}
	path := "/donationCampaigns"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.donationCampaignsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by DonationsApi.Donations
type DonationsApiDonationsInput struct {
	idempotencyKey         *string
	donationPaymentRequest *DonationPaymentRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r DonationsApiDonationsInput) IdempotencyKey(idempotencyKey string) DonationsApiDonationsInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r DonationsApiDonationsInput) DonationPaymentRequest(donationPaymentRequest DonationPaymentRequest) DonationsApiDonationsInput {
	r.donationPaymentRequest = &donationPaymentRequest
	return r
}

/*
Prepare a request for Donations

@return DonationsApiDonationsInput
*/
func (a *DonationsApi) DonationsInput() DonationsApiDonationsInput {
	return DonationsApiDonationsInput{}
}

/*
Donations Start a transaction for donations

Takes in the donation token generated by the `/payments` request and uses it to make the donation for the donation account specified in the request.

For more information, see [Donations](https://docs.adyen.com/online-payments/donations).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r DonationsApiDonationsInput - Request parameters, see DonationsInput
@return DonationPaymentResponse, *http.Response, error
*/
func (a *DonationsApi) Donations(ctx context.Context, r DonationsApiDonationsInput) (DonationPaymentResponse, *http.Response, error) {
	res := &DonationPaymentResponse{}
	path := "/donations"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.donationPaymentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
