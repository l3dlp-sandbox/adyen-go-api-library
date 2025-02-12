/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// UtilityApi service
type UtilityApi common.Service

// All parameters accepted by UtilityApi.GetApplePaySession
type UtilityApiGetApplePaySessionInput struct {
	idempotencyKey               *string
	createApplePaySessionRequest *CreateApplePaySessionRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r UtilityApiGetApplePaySessionInput) IdempotencyKey(idempotencyKey string) UtilityApiGetApplePaySessionInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r UtilityApiGetApplePaySessionInput) CreateApplePaySessionRequest(createApplePaySessionRequest CreateApplePaySessionRequest) UtilityApiGetApplePaySessionInput {
	r.createApplePaySessionRequest = &createApplePaySessionRequest
	return r
}

/*
Prepare a request for GetApplePaySession

@return UtilityApiGetApplePaySessionInput
*/
func (a *UtilityApi) GetApplePaySessionInput() UtilityApiGetApplePaySessionInput {
	return UtilityApiGetApplePaySessionInput{}
}

/*
GetApplePaySession Get an Apple Pay session

You need to use this endpoint if you have an API-only integration with Apple Pay which uses Adyen's Apple Pay certificate.

The endpoint returns the Apple Pay session data which you need to complete the [Apple Pay session validation](https://docs.adyen.com/payment-methods/apple-pay/api-only?tab=adyen-certificate-validation_1#complete-apple-pay-session-validation).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UtilityApiGetApplePaySessionInput - Request parameters, see GetApplePaySessionInput
@return ApplePaySessionResponse, *http.Response, error
*/
func (a *UtilityApi) GetApplePaySession(ctx context.Context, r UtilityApiGetApplePaySessionInput) (ApplePaySessionResponse, *http.Response, error) {
	res := &ApplePaySessionResponse{}
	path := "/applePay/sessions"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createApplePaySessionRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by UtilityApi.OriginKeys
type UtilityApiOriginKeysInput struct {
	idempotencyKey         *string
	checkoutUtilityRequest *CheckoutUtilityRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r UtilityApiOriginKeysInput) IdempotencyKey(idempotencyKey string) UtilityApiOriginKeysInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r UtilityApiOriginKeysInput) CheckoutUtilityRequest(checkoutUtilityRequest CheckoutUtilityRequest) UtilityApiOriginKeysInput {
	r.checkoutUtilityRequest = &checkoutUtilityRequest
	return r
}

/*
Prepare a request for OriginKeys

@return UtilityApiOriginKeysInput

Deprecated
*/
func (a *UtilityApi) OriginKeysInput() UtilityApiOriginKeysInput {
	return UtilityApiOriginKeysInput{}
}

/*
OriginKeys Create originKey values for domains

This operation takes the origin domains and returns a JSON object containing the corresponding origin keys for the domains.
> If you're still using origin key for your Web Drop-in or Components integration, we recommend [switching to client key](https://docs.adyen.com/development-resources/client-side-authentication/migrate-from-origin-key-to-client-key). This allows you to use a single key for all origins, add or remove origins without generating a new key, and detect the card type from the number entered in your payment form.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UtilityApiOriginKeysInput - Request parameters, see OriginKeysInput
@return CheckoutUtilityResponse, *http.Response, error

    Deprecated
*/
func (a *UtilityApi) OriginKeys(ctx context.Context, r UtilityApiOriginKeysInput) (CheckoutUtilityResponse, *http.Response, error) {
	res := &CheckoutUtilityResponse{}
	path := "/originKeys"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.checkoutUtilityRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
