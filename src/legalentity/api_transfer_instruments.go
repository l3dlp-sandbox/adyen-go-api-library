/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// TransferInstrumentsApi service
type TransferInstrumentsApi common.Service

// All parameters accepted by TransferInstrumentsApi.CreateTransferInstrument
type TransferInstrumentsApiCreateTransferInstrumentInput struct {
	xRequestedVerificationCode *string
	transferInstrumentInfo     *TransferInstrumentInfo
}

// Use a suberror code as your requested verification code. You can include one code at a time in your request header. Requested verification codes can only be used in your test environment.
func (r TransferInstrumentsApiCreateTransferInstrumentInput) XRequestedVerificationCode(xRequestedVerificationCode string) TransferInstrumentsApiCreateTransferInstrumentInput {
	r.xRequestedVerificationCode = &xRequestedVerificationCode
	return r
}

func (r TransferInstrumentsApiCreateTransferInstrumentInput) TransferInstrumentInfo(transferInstrumentInfo TransferInstrumentInfo) TransferInstrumentsApiCreateTransferInstrumentInput {
	r.transferInstrumentInfo = &transferInstrumentInfo
	return r
}

/*
Prepare a request for CreateTransferInstrument

@return TransferInstrumentsApiCreateTransferInstrumentInput
*/
func (a *TransferInstrumentsApi) CreateTransferInstrumentInput() TransferInstrumentsApiCreateTransferInstrumentInput {
	return TransferInstrumentsApiCreateTransferInstrumentInput{}
}

/*
CreateTransferInstrument Create a transfer instrument

Creates a transfer instrument.

A transfer instrument is a bank account that a legal entity owns. Adyen performs verification checks on the transfer instrument as required by payment industry regulations. We inform you of the verification results through webhooks or API responses.

When the transfer instrument passes the verification checks, you can start sending funds from the balance platform to the transfer instrument (such as payouts).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransferInstrumentsApiCreateTransferInstrumentInput - Request parameters, see CreateTransferInstrumentInput
@return TransferInstrument, *http.Response, error
*/
func (a *TransferInstrumentsApi) CreateTransferInstrument(ctx context.Context, r TransferInstrumentsApiCreateTransferInstrumentInput) (TransferInstrument, *http.Response, error) {
	res := &TransferInstrument{}
	path := "/transferInstruments"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.xRequestedVerificationCode != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "x-requested-verification-code", r.xRequestedVerificationCode, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.transferInstrumentInfo,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by TransferInstrumentsApi.DeleteTransferInstrument
type TransferInstrumentsApiDeleteTransferInstrumentInput struct {
	id string
}

/*
Prepare a request for DeleteTransferInstrument
@param id The unique identifier of the transfer instrument to be deleted.
@return TransferInstrumentsApiDeleteTransferInstrumentInput
*/
func (a *TransferInstrumentsApi) DeleteTransferInstrumentInput(id string) TransferInstrumentsApiDeleteTransferInstrumentInput {
	return TransferInstrumentsApiDeleteTransferInstrumentInput{
		id: id,
	}
}

/*
DeleteTransferInstrument Delete a transfer instrument

Deletes a transfer instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransferInstrumentsApiDeleteTransferInstrumentInput - Request parameters, see DeleteTransferInstrumentInput
@return *http.Response, error
*/
func (a *TransferInstrumentsApi) DeleteTransferInstrument(ctx context.Context, r TransferInstrumentsApiDeleteTransferInstrumentInput) (*http.Response, error) {
	var res interface{}
	path := "/transferInstruments/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodDelete,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return httpRes, err
}

// All parameters accepted by TransferInstrumentsApi.GetTransferInstrument
type TransferInstrumentsApiGetTransferInstrumentInput struct {
	id string
}

/*
Prepare a request for GetTransferInstrument
@param id The unique identifier of the transfer instrument.
@return TransferInstrumentsApiGetTransferInstrumentInput
*/
func (a *TransferInstrumentsApi) GetTransferInstrumentInput(id string) TransferInstrumentsApiGetTransferInstrumentInput {
	return TransferInstrumentsApiGetTransferInstrumentInput{
		id: id,
	}
}

/*
GetTransferInstrument Get a transfer instrument

Returns the details of a transfer instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransferInstrumentsApiGetTransferInstrumentInput - Request parameters, see GetTransferInstrumentInput
@return TransferInstrument, *http.Response, error
*/
func (a *TransferInstrumentsApi) GetTransferInstrument(ctx context.Context, r TransferInstrumentsApiGetTransferInstrumentInput) (TransferInstrument, *http.Response, error) {
	res := &TransferInstrument{}
	path := "/transferInstruments/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by TransferInstrumentsApi.UpdateTransferInstrument
type TransferInstrumentsApiUpdateTransferInstrumentInput struct {
	id                         string
	xRequestedVerificationCode *string
	transferInstrumentInfo     *TransferInstrumentInfo
}

// Use the requested verification code 0_0001 to resolve any suberrors associated with the transfer instrument. Requested verification codes can only be used in your test environment.
func (r TransferInstrumentsApiUpdateTransferInstrumentInput) XRequestedVerificationCode(xRequestedVerificationCode string) TransferInstrumentsApiUpdateTransferInstrumentInput {
	r.xRequestedVerificationCode = &xRequestedVerificationCode
	return r
}

func (r TransferInstrumentsApiUpdateTransferInstrumentInput) TransferInstrumentInfo(transferInstrumentInfo TransferInstrumentInfo) TransferInstrumentsApiUpdateTransferInstrumentInput {
	r.transferInstrumentInfo = &transferInstrumentInfo
	return r
}

/*
Prepare a request for UpdateTransferInstrument
@param id The unique identifier of the transfer instrument.
@return TransferInstrumentsApiUpdateTransferInstrumentInput
*/
func (a *TransferInstrumentsApi) UpdateTransferInstrumentInput(id string) TransferInstrumentsApiUpdateTransferInstrumentInput {
	return TransferInstrumentsApiUpdateTransferInstrumentInput{
		id: id,
	}
}

/*
UpdateTransferInstrument Update a transfer instrument

Updates a transfer instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransferInstrumentsApiUpdateTransferInstrumentInput - Request parameters, see UpdateTransferInstrumentInput
@return TransferInstrument, *http.Response, error
*/
func (a *TransferInstrumentsApi) UpdateTransferInstrument(ctx context.Context, r TransferInstrumentsApiUpdateTransferInstrumentInput) (TransferInstrument, *http.Response, error) {
	res := &TransferInstrument{}
	path := "/transferInstruments/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.xRequestedVerificationCode != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "x-requested-verification-code", r.xRequestedVerificationCode, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.transferInstrumentInfo,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
