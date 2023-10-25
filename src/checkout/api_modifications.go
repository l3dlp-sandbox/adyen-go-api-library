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
	"strings"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// ModificationsApi service
type ModificationsApi common.Service

// All parameters accepted by ModificationsApi.CancelAuthorisedPayment
type ModificationsApiCancelAuthorisedPaymentInput struct {
	idempotencyKey                 *string
	standalonePaymentCancelRequest *StandalonePaymentCancelRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r ModificationsApiCancelAuthorisedPaymentInput) IdempotencyKey(idempotencyKey string) ModificationsApiCancelAuthorisedPaymentInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ModificationsApiCancelAuthorisedPaymentInput) StandalonePaymentCancelRequest(standalonePaymentCancelRequest StandalonePaymentCancelRequest) ModificationsApiCancelAuthorisedPaymentInput {
	r.standalonePaymentCancelRequest = &standalonePaymentCancelRequest
	return r
}

/*
Prepare a request for CancelAuthorisedPayment

@return ModificationsApiCancelAuthorisedPaymentInput
*/
func (a *ModificationsApi) CancelAuthorisedPaymentInput() ModificationsApiCancelAuthorisedPaymentInput {
	return ModificationsApiCancelAuthorisedPaymentInput{}
}

/*
CancelAuthorisedPayment Cancel an authorised payment

Cancels the authorisation on a payment that has not yet been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/captures), and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**TECHNICAL_CANCEL** webhook](https://docs.adyen.com/online-payments/cancel#cancellation-webhook).

If you want to cancel a payment using the [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference), use the [`/payments/{paymentPspReference}/cancels`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/cancels) endpoint instead.

If you want to cancel a payment but are not sure whether it has been captured, use the [`/payments/{paymentPspReference}/reversals`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/reversals) endpoint instead.

For more information, refer to [Cancel](https://docs.adyen.com/online-payments/cancel).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiCancelAuthorisedPaymentInput - Request parameters, see CancelAuthorisedPaymentInput
@return StandalonePaymentCancelResponse, *http.Response, error
*/
func (a *ModificationsApi) CancelAuthorisedPayment(ctx context.Context, r ModificationsApiCancelAuthorisedPaymentInput) (StandalonePaymentCancelResponse, *http.Response, error) {
	res := &StandalonePaymentCancelResponse{}
	path := "/cancels"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.standalonePaymentCancelRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.CancelAuthorisedPaymentByPspReference
type ModificationsApiCancelAuthorisedPaymentByPspReferenceInput struct {
	paymentPspReference  string
	idempotencyKey       *string
	paymentCancelRequest *PaymentCancelRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r ModificationsApiCancelAuthorisedPaymentByPspReferenceInput) IdempotencyKey(idempotencyKey string) ModificationsApiCancelAuthorisedPaymentByPspReferenceInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ModificationsApiCancelAuthorisedPaymentByPspReferenceInput) PaymentCancelRequest(paymentCancelRequest PaymentCancelRequest) ModificationsApiCancelAuthorisedPaymentByPspReferenceInput {
	r.paymentCancelRequest = &paymentCancelRequest
	return r
}

/*
Prepare a request for CancelAuthorisedPaymentByPspReference
@param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to cancel.
@return ModificationsApiCancelAuthorisedPaymentByPspReferenceInput
*/
func (a *ModificationsApi) CancelAuthorisedPaymentByPspReferenceInput(paymentPspReference string) ModificationsApiCancelAuthorisedPaymentByPspReferenceInput {
	return ModificationsApiCancelAuthorisedPaymentByPspReferenceInput{
		paymentPspReference: paymentPspReference,
	}
}

/*
CancelAuthorisedPaymentByPspReference Cancel an authorised payment

Cancels the authorisation on a payment that has not yet been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/paymentPspReference/captures), and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**CANCELLATION** webhook](https://docs.adyen.com/online-payments/cancel#cancellation-webhook).

If you want to cancel a payment but don't have the [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference), use the [`/cancels`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/cancels) endpoint instead.

If you want to cancel a payment but are not sure whether it has been captured, use the [`/payments/{paymentPspReference}/reversals`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/reversals) endpoint instead.

For more information, refer to [Cancel](https://docs.adyen.com/online-payments/cancel).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiCancelAuthorisedPaymentByPspReferenceInput - Request parameters, see CancelAuthorisedPaymentByPspReferenceInput
@return PaymentCancelResponse, *http.Response, error
*/
func (a *ModificationsApi) CancelAuthorisedPaymentByPspReference(ctx context.Context, r ModificationsApiCancelAuthorisedPaymentByPspReferenceInput) (PaymentCancelResponse, *http.Response, error) {
	res := &PaymentCancelResponse{}
	path := "/payments/{paymentPspReference}/cancels"
	path = strings.Replace(path, "{"+"paymentPspReference"+"}", url.PathEscape(common.ParameterValueToString(r.paymentPspReference, "paymentPspReference")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentCancelRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.CaptureAuthorisedPayment
type ModificationsApiCaptureAuthorisedPaymentInput struct {
	paymentPspReference   string
	idempotencyKey        *string
	paymentCaptureRequest *PaymentCaptureRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r ModificationsApiCaptureAuthorisedPaymentInput) IdempotencyKey(idempotencyKey string) ModificationsApiCaptureAuthorisedPaymentInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ModificationsApiCaptureAuthorisedPaymentInput) PaymentCaptureRequest(paymentCaptureRequest PaymentCaptureRequest) ModificationsApiCaptureAuthorisedPaymentInput {
	r.paymentCaptureRequest = &paymentCaptureRequest
	return r
}

/*
Prepare a request for CaptureAuthorisedPayment
@param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to capture.
@return ModificationsApiCaptureAuthorisedPaymentInput
*/
func (a *ModificationsApi) CaptureAuthorisedPaymentInput(paymentPspReference string) ModificationsApiCaptureAuthorisedPaymentInput {
	return ModificationsApiCaptureAuthorisedPaymentInput{
		paymentPspReference: paymentPspReference,
	}
}

/*
CaptureAuthorisedPayment Capture an authorised payment

 Captures an authorised payment and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**CAPTURE** webhook](https://docs.adyen.com/online-payments/capture#capture-notification).

You can capture either the full authorised amount or a part of the authorised amount. By default, any unclaimed amount after a partial capture gets cancelled. This does not apply if you enabled multiple partial captures on your account and the payment method supports multiple partial captures.

[Automatic capture](https://docs.adyen.com/online-payments/capture#automatic-capture) is the default setting for most payment methods. In these cases, you don't need to make capture requests. However, making capture requests for payments that are captured automatically does not result in double charges.

For more information, refer to [Capture](https://docs.adyen.com/online-payments/capture).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiCaptureAuthorisedPaymentInput - Request parameters, see CaptureAuthorisedPaymentInput
@return PaymentCaptureResponse, *http.Response, error
*/
func (a *ModificationsApi) CaptureAuthorisedPayment(ctx context.Context, r ModificationsApiCaptureAuthorisedPaymentInput) (PaymentCaptureResponse, *http.Response, error) {
	res := &PaymentCaptureResponse{}
	path := "/payments/{paymentPspReference}/captures"
	path = strings.Replace(path, "{"+"paymentPspReference"+"}", url.PathEscape(common.ParameterValueToString(r.paymentPspReference, "paymentPspReference")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentCaptureRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.RefundCapturedPayment
type ModificationsApiRefundCapturedPaymentInput struct {
	paymentPspReference  string
	idempotencyKey       *string
	paymentRefundRequest *PaymentRefundRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r ModificationsApiRefundCapturedPaymentInput) IdempotencyKey(idempotencyKey string) ModificationsApiRefundCapturedPaymentInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ModificationsApiRefundCapturedPaymentInput) PaymentRefundRequest(paymentRefundRequest PaymentRefundRequest) ModificationsApiRefundCapturedPaymentInput {
	r.paymentRefundRequest = &paymentRefundRequest
	return r
}

/*
Prepare a request for RefundCapturedPayment
@param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to refund.
@return ModificationsApiRefundCapturedPaymentInput
*/
func (a *ModificationsApi) RefundCapturedPaymentInput(paymentPspReference string) ModificationsApiRefundCapturedPaymentInput {
	return ModificationsApiRefundCapturedPaymentInput{
		paymentPspReference: paymentPspReference,
	}
}

/*
RefundCapturedPayment Refund a captured payment

Refunds a payment that has been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/captures), and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**REFUND** webhook](https://docs.adyen.com/online-payments/refund#refund-webhook).

You can refund either the full captured amount or a part of the captured amount. You can also perform multiple partial refunds, as long as their sum doesn't exceed the captured amount.

> Some payment methods do not support partial refunds. To learn if a payment method supports partial refunds, refer to the payment method page such as [cards](https://docs.adyen.com/payment-methods/cards#supported-cards), [iDEAL](https://docs.adyen.com/payment-methods/ideal), or [Klarna](https://docs.adyen.com/payment-methods/klarna).

If you want to refund a payment but are not sure whether it has been captured, use the [`/payments/{paymentPspReference}/reversals`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/reversals) endpoint instead.

For more information, refer to [Refund](https://docs.adyen.com/online-payments/refund).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiRefundCapturedPaymentInput - Request parameters, see RefundCapturedPaymentInput
@return PaymentRefundResponse, *http.Response, error
*/
func (a *ModificationsApi) RefundCapturedPayment(ctx context.Context, r ModificationsApiRefundCapturedPaymentInput) (PaymentRefundResponse, *http.Response, error) {
	res := &PaymentRefundResponse{}
	path := "/payments/{paymentPspReference}/refunds"
	path = strings.Replace(path, "{"+"paymentPspReference"+"}", url.PathEscape(common.ParameterValueToString(r.paymentPspReference, "paymentPspReference")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentRefundRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.RefundOrCancelPayment
type ModificationsApiRefundOrCancelPaymentInput struct {
	paymentPspReference    string
	idempotencyKey         *string
	paymentReversalRequest *PaymentReversalRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r ModificationsApiRefundOrCancelPaymentInput) IdempotencyKey(idempotencyKey string) ModificationsApiRefundOrCancelPaymentInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ModificationsApiRefundOrCancelPaymentInput) PaymentReversalRequest(paymentReversalRequest PaymentReversalRequest) ModificationsApiRefundOrCancelPaymentInput {
	r.paymentReversalRequest = &paymentReversalRequest
	return r
}

/*
Prepare a request for RefundOrCancelPayment
@param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to reverse.
@return ModificationsApiRefundOrCancelPaymentInput
*/
func (a *ModificationsApi) RefundOrCancelPaymentInput(paymentPspReference string) ModificationsApiRefundOrCancelPaymentInput {
	return ModificationsApiRefundOrCancelPaymentInput{
		paymentPspReference: paymentPspReference,
	}
}

/*
RefundOrCancelPayment Refund or cancel a payment

[Refunds](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/refunds) a payment if it has already been captured, and [cancels](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/cancels) a payment if it has not yet been captured. Returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**CANCEL_OR_REFUND** webhook](https://docs.adyen.com/online-payments/reverse#cancel-or-refund-webhook).

The reversed amount is always the full payment amount.
> Do not use this request for payments that involve multiple partial captures.

For more information, refer to [Reversal](https://docs.adyen.com/online-payments/reversal).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiRefundOrCancelPaymentInput - Request parameters, see RefundOrCancelPaymentInput
@return PaymentReversalResponse, *http.Response, error
*/
func (a *ModificationsApi) RefundOrCancelPayment(ctx context.Context, r ModificationsApiRefundOrCancelPaymentInput) (PaymentReversalResponse, *http.Response, error) {
	res := &PaymentReversalResponse{}
	path := "/payments/{paymentPspReference}/reversals"
	path = strings.Replace(path, "{"+"paymentPspReference"+"}", url.PathEscape(common.ParameterValueToString(r.paymentPspReference, "paymentPspReference")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentReversalRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by ModificationsApi.UpdateAuthorisedAmount
type ModificationsApiUpdateAuthorisedAmountInput struct {
	paymentPspReference        string
	idempotencyKey             *string
	paymentAmountUpdateRequest *PaymentAmountUpdateRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r ModificationsApiUpdateAuthorisedAmountInput) IdempotencyKey(idempotencyKey string) ModificationsApiUpdateAuthorisedAmountInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ModificationsApiUpdateAuthorisedAmountInput) PaymentAmountUpdateRequest(paymentAmountUpdateRequest PaymentAmountUpdateRequest) ModificationsApiUpdateAuthorisedAmountInput {
	r.paymentAmountUpdateRequest = &paymentAmountUpdateRequest
	return r
}

/*
Prepare a request for UpdateAuthorisedAmount
@param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment.
@return ModificationsApiUpdateAuthorisedAmountInput
*/
func (a *ModificationsApi) UpdateAuthorisedAmountInput(paymentPspReference string) ModificationsApiUpdateAuthorisedAmountInput {
	return ModificationsApiUpdateAuthorisedAmountInput{
		paymentPspReference: paymentPspReference,
	}
}

/*
UpdateAuthorisedAmount Update an authorised amount

Increases or decreases the authorised payment amount and returns a unique reference for this request. You get the outcome of the request asynchronously, in an [**AUTHORISATION_ADJUSTMENT** webhook](https://docs.adyen.com/development-resources/webhooks/understand-notifications#event-codes).

You can only update authorised amounts that have not yet been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/captures).

The amount you specify in the request is the updated amount, which is larger or smaller than the initial authorised amount.

For more information, refer to [Authorisation adjustment](https://docs.adyen.com/online-payments/adjust-authorisation#use-cases).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ModificationsApiUpdateAuthorisedAmountInput - Request parameters, see UpdateAuthorisedAmountInput
@return PaymentAmountUpdateResponse, *http.Response, error
*/
func (a *ModificationsApi) UpdateAuthorisedAmount(ctx context.Context, r ModificationsApiUpdateAuthorisedAmountInput) (PaymentAmountUpdateResponse, *http.Response, error) {
	res := &PaymentAmountUpdateResponse{}
	path := "/payments/{paymentPspReference}/amountUpdates"
	path = strings.Replace(path, "{"+"paymentPspReference"+"}", url.PathEscape(common.ParameterValueToString(r.paymentPspReference, "paymentPspReference")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentAmountUpdateRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
