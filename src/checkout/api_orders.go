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

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// OrdersApi service
type OrdersApi common.Service

// All parameters accepted by OrdersApi.CancelOrder
type OrdersApiCancelOrderInput struct {
	idempotencyKey     *string
	cancelOrderRequest *CancelOrderRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r OrdersApiCancelOrderInput) IdempotencyKey(idempotencyKey string) OrdersApiCancelOrderInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r OrdersApiCancelOrderInput) CancelOrderRequest(cancelOrderRequest CancelOrderRequest) OrdersApiCancelOrderInput {
	r.cancelOrderRequest = &cancelOrderRequest
	return r
}

/*
Prepare a request for CancelOrder

@return OrdersApiCancelOrderInput
*/
func (a *OrdersApi) CancelOrderInput() OrdersApiCancelOrderInput {
	return OrdersApiCancelOrderInput{}
}

/*
CancelOrder Cancel an order

Cancels an order. Cancellation of an order results in an automatic rollback of all payments made in the order, either by canceling or refunding the payment, depending on the type of payment method.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r OrdersApiCancelOrderInput - Request parameters, see CancelOrderInput
@return CancelOrderResponse, *http.Response, error
*/
func (a *OrdersApi) CancelOrder(ctx context.Context, r OrdersApiCancelOrderInput) (CancelOrderResponse, *http.Response, error) {
	res := &CancelOrderResponse{}
	path := "/orders/cancel"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.cancelOrderRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by OrdersApi.GetBalanceOfGiftCard
type OrdersApiGetBalanceOfGiftCardInput struct {
	idempotencyKey      *string
	balanceCheckRequest *BalanceCheckRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r OrdersApiGetBalanceOfGiftCardInput) IdempotencyKey(idempotencyKey string) OrdersApiGetBalanceOfGiftCardInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r OrdersApiGetBalanceOfGiftCardInput) BalanceCheckRequest(balanceCheckRequest BalanceCheckRequest) OrdersApiGetBalanceOfGiftCardInput {
	r.balanceCheckRequest = &balanceCheckRequest
	return r
}

/*
Prepare a request for GetBalanceOfGiftCard

@return OrdersApiGetBalanceOfGiftCardInput
*/
func (a *OrdersApi) GetBalanceOfGiftCardInput() OrdersApiGetBalanceOfGiftCardInput {
	return OrdersApiGetBalanceOfGiftCardInput{}
}

/*
GetBalanceOfGiftCard Get the balance of a gift card

Retrieves the balance remaining on a shopper's gift card. To check a gift card's balance, make a POST `/paymentMethods/balance` call and include the gift card's details inside a `paymentMethod` object.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r OrdersApiGetBalanceOfGiftCardInput - Request parameters, see GetBalanceOfGiftCardInput
@return BalanceCheckResponse, *http.Response, error
*/
func (a *OrdersApi) GetBalanceOfGiftCard(ctx context.Context, r OrdersApiGetBalanceOfGiftCardInput) (BalanceCheckResponse, *http.Response, error) {
	res := &BalanceCheckResponse{}
	path := "/paymentMethods/balance"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.balanceCheckRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by OrdersApi.Orders
type OrdersApiOrdersInput struct {
	idempotencyKey     *string
	createOrderRequest *CreateOrderRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r OrdersApiOrdersInput) IdempotencyKey(idempotencyKey string) OrdersApiOrdersInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r OrdersApiOrdersInput) CreateOrderRequest(createOrderRequest CreateOrderRequest) OrdersApiOrdersInput {
	r.createOrderRequest = &createOrderRequest
	return r
}

/*
Prepare a request for Orders

@return OrdersApiOrdersInput
*/
func (a *OrdersApi) OrdersInput() OrdersApiOrdersInput {
	return OrdersApiOrdersInput{}
}

/*
Orders Create an order

Creates an order to be used for partial payments. Make a POST `/orders` call before making a `/payments` call when processing payments with different payment methods.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r OrdersApiOrdersInput - Request parameters, see OrdersInput
@return CreateOrderResponse, *http.Response, error
*/
func (a *OrdersApi) Orders(ctx context.Context, r OrdersApiOrdersInput) (CreateOrderResponse, *http.Response, error) {
	res := &CreateOrderResponse{}
	path := "/orders"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createOrderRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
