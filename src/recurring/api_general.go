/*
Adyen Recurring API (deprecated)

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.CreatePermit
type GeneralApiCreatePermitInput struct {
	createPermitRequest *CreatePermitRequest
}

func (r GeneralApiCreatePermitInput) CreatePermitRequest(createPermitRequest CreatePermitRequest) GeneralApiCreatePermitInput {
	r.createPermitRequest = &createPermitRequest
	return r
}

/*
Prepare a request for CreatePermit

@return GeneralApiCreatePermitInput

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) CreatePermitInput() GeneralApiCreatePermitInput {
	return GeneralApiCreatePermitInput{}
}

/*
CreatePermit Create new permits linked to a recurring contract.

Create permits for a recurring contract, including support for defining restrictions.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiCreatePermitInput - Request parameters, see CreatePermitInput
@return CreatePermitResult, *http.Response, error

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) CreatePermit(ctx context.Context, r GeneralApiCreatePermitInput) (CreatePermitResult, *http.Response, error) {
	res := &CreatePermitResult{}
	path := "/createPermit"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createPermitRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.Disable
type GeneralApiDisableInput struct {
	disableRequest *DisableRequest
}

func (r GeneralApiDisableInput) DisableRequest(disableRequest DisableRequest) GeneralApiDisableInput {
	r.disableRequest = &disableRequest
	return r
}

/*
Prepare a request for Disable

@return GeneralApiDisableInput

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) DisableInput() GeneralApiDisableInput {
	return GeneralApiDisableInput{}
}

/*
Disable Disable stored payment details

Disables stored payment details to stop charging a shopper with this particular recurring detail ID.

For more information, refer to [Disable stored details](https://docs.adyen.com/classic-integration/recurring-payments/disable-stored-details/).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiDisableInput - Request parameters, see DisableInput
@return DisableResult, *http.Response, error

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) Disable(ctx context.Context, r GeneralApiDisableInput) (DisableResult, *http.Response, error) {
	res := &DisableResult{}
	path := "/disable"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.disableRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.DisablePermit
type GeneralApiDisablePermitInput struct {
	disablePermitRequest *DisablePermitRequest
}

func (r GeneralApiDisablePermitInput) DisablePermitRequest(disablePermitRequest DisablePermitRequest) GeneralApiDisablePermitInput {
	r.disablePermitRequest = &disablePermitRequest
	return r
}

/*
Prepare a request for DisablePermit

@return GeneralApiDisablePermitInput

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) DisablePermitInput() GeneralApiDisablePermitInput {
	return GeneralApiDisablePermitInput{}
}

/*
DisablePermit Disable an existing permit.

Disable a permit that was previously linked to a recurringDetailReference.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiDisablePermitInput - Request parameters, see DisablePermitInput
@return DisablePermitResult, *http.Response, error

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) DisablePermit(ctx context.Context, r GeneralApiDisablePermitInput) (DisablePermitResult, *http.Response, error) {
	res := &DisablePermitResult{}
	path := "/disablePermit"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.disablePermitRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.ListRecurringDetails
type GeneralApiListRecurringDetailsInput struct {
	recurringDetailsRequest *RecurringDetailsRequest
}

func (r GeneralApiListRecurringDetailsInput) RecurringDetailsRequest(recurringDetailsRequest RecurringDetailsRequest) GeneralApiListRecurringDetailsInput {
	r.recurringDetailsRequest = &recurringDetailsRequest
	return r
}

/*
Prepare a request for ListRecurringDetails

@return GeneralApiListRecurringDetailsInput

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) ListRecurringDetailsInput() GeneralApiListRecurringDetailsInput {
	return GeneralApiListRecurringDetailsInput{}
}

/*
ListRecurringDetails Get stored payment details

Lists the stored payment details for a shopper, if there are any available. The recurring detail ID can be used with a regular authorisation request to charge the shopper. A summary of the payment detail is returned for presentation to the shopper.

For more information, refer to [Retrieve stored details](https://docs.adyen.com/classic-integration/recurring-payments/retrieve-stored-details/).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiListRecurringDetailsInput - Request parameters, see ListRecurringDetailsInput
@return RecurringDetailsResult, *http.Response, error

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) ListRecurringDetails(ctx context.Context, r GeneralApiListRecurringDetailsInput) (RecurringDetailsResult, *http.Response, error) {
	res := &RecurringDetailsResult{}
	path := "/listRecurringDetails"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.recurringDetailsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.NotifyShopper
type GeneralApiNotifyShopperInput struct {
	notifyShopperRequest *NotifyShopperRequest
}

func (r GeneralApiNotifyShopperInput) NotifyShopperRequest(notifyShopperRequest NotifyShopperRequest) GeneralApiNotifyShopperInput {
	r.notifyShopperRequest = &notifyShopperRequest
	return r
}

/*
Prepare a request for NotifyShopper

@return GeneralApiNotifyShopperInput

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) NotifyShopperInput() GeneralApiNotifyShopperInput {
	return GeneralApiNotifyShopperInput{}
}

/*
NotifyShopper Ask issuer to notify the shopper

Sends a request to the issuer so they can inform the shopper about the upcoming recurring payment. This endpoint is used only for local acquiring in India. For more information, refer to [Recurring card payments in India](https://docs.adyen.com/payment-methods/cards/cards-recurring-india).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiNotifyShopperInput - Request parameters, see NotifyShopperInput
@return NotifyShopperResult, *http.Response, error

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) NotifyShopper(ctx context.Context, r GeneralApiNotifyShopperInput) (NotifyShopperResult, *http.Response, error) {
	res := &NotifyShopperResult{}
	path := "/notifyShopper"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.notifyShopperRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.ScheduleAccountUpdater
type GeneralApiScheduleAccountUpdaterInput struct {
	scheduleAccountUpdaterRequest *ScheduleAccountUpdaterRequest
}

func (r GeneralApiScheduleAccountUpdaterInput) ScheduleAccountUpdaterRequest(scheduleAccountUpdaterRequest ScheduleAccountUpdaterRequest) GeneralApiScheduleAccountUpdaterInput {
	r.scheduleAccountUpdaterRequest = &scheduleAccountUpdaterRequest
	return r
}

/*
Prepare a request for ScheduleAccountUpdater

@return GeneralApiScheduleAccountUpdaterInput

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) ScheduleAccountUpdaterInput() GeneralApiScheduleAccountUpdaterInput {
	return GeneralApiScheduleAccountUpdaterInput{}
}

/*
ScheduleAccountUpdater Schedule running the Account Updater

When making the API call, you can submit either the credit card information, or the recurring detail reference and the shopper reference:
* If the card information is provided, all the sub-fields for `card` are mandatory.
* If the recurring detail reference is provided, the fields for `shopperReference` and `selectedRecurringDetailReference` are mandatory.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiScheduleAccountUpdaterInput - Request parameters, see ScheduleAccountUpdaterInput
@return ScheduleAccountUpdaterResult, *http.Response, error

Deprecated since Adyen Recurring API (deprecated) v68
*/
func (a *GeneralApi) ScheduleAccountUpdater(ctx context.Context, r GeneralApiScheduleAccountUpdaterInput) (ScheduleAccountUpdaterResult, *http.Response, error) {
	res := &ScheduleAccountUpdaterResult{}
	path := "/scheduleAccountUpdater"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.scheduleAccountUpdaterRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
