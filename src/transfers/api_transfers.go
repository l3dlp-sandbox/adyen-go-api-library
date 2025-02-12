/*
Transfers API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// TransfersApi service
type TransfersApi common.Service

// All parameters accepted by TransfersApi.TransferFunds
type TransfersApiTransferFundsInput struct {
	transferInfo *TransferInfo
}

func (r TransfersApiTransferFundsInput) TransferInfo(transferInfo TransferInfo) TransfersApiTransferFundsInput {
	r.transferInfo = &transferInfo
	return r
}

/*
Prepare a request for TransferFunds

@return TransfersApiTransferFundsInput
*/
func (a *TransfersApi) TransferFundsInput() TransfersApiTransferFundsInput {
	return TransfersApiTransferFundsInput{}
}

/*
TransferFunds Transfer funds

Starts a request to transfer funds to [balance accounts](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts), [transfer instruments](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments), or third-party bank accounts. Adyen sends the outcome of the transfer request through webhooks.

To use this endpoint, you need an additional role for your API credential and transfers must be enabled for the source balance account. Your Adyen contact will set these up for you.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransfersApiTransferFundsInput - Request parameters, see TransferFundsInput
@return Transfer, *http.Response, error
*/
func (a *TransfersApi) TransferFunds(ctx context.Context, r TransfersApiTransferFundsInput) (Transfer, *http.Response, error) {
	res := &Transfer{}
	path := "/transfers"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.transferInfo,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
