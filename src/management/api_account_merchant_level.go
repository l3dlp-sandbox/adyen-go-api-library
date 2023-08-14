/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// AccountMerchantLevelApi service
type AccountMerchantLevelApi common.Service

// All parameters accepted by AccountMerchantLevelApi.CreateMerchantAccount
type AccountMerchantLevelApiCreateMerchantAccountInput struct {
	createMerchantRequest *CreateMerchantRequest
}

func (r AccountMerchantLevelApiCreateMerchantAccountInput) CreateMerchantRequest(createMerchantRequest CreateMerchantRequest) AccountMerchantLevelApiCreateMerchantAccountInput {
	r.createMerchantRequest = &createMerchantRequest
	return r
}

/*
Prepare a request for CreateMerchantAccount

@return AccountMerchantLevelApiCreateMerchantAccountInput
*/
func (a *AccountMerchantLevelApi) CreateMerchantAccountInput() AccountMerchantLevelApiCreateMerchantAccountInput {
	return AccountMerchantLevelApiCreateMerchantAccountInput{}
}

/*
CreateMerchantAccount Create a merchant account

Creates a merchant account for the company account specified in the request.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Accounts read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountMerchantLevelApiCreateMerchantAccountInput - Request parameters, see CreateMerchantAccountInput
@return CreateMerchantResponse, *http.Response, error
*/
func (a *AccountMerchantLevelApi) CreateMerchantAccount(ctx context.Context, r AccountMerchantLevelApiCreateMerchantAccountInput) (CreateMerchantResponse, *http.Response, error) {
	res := &CreateMerchantResponse{}
	path := "/merchants"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createMerchantRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

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

// All parameters accepted by AccountMerchantLevelApi.GetMerchantAccount
type AccountMerchantLevelApiGetMerchantAccountInput struct {
	merchantId string
}

/*
Prepare a request for GetMerchantAccount
@param merchantId The unique identifier of the merchant account.
@return AccountMerchantLevelApiGetMerchantAccountInput
*/
func (a *AccountMerchantLevelApi) GetMerchantAccountInput(merchantId string) AccountMerchantLevelApiGetMerchantAccountInput {
	return AccountMerchantLevelApiGetMerchantAccountInput{
		merchantId: merchantId,
	}
}

/*
GetMerchantAccount Get a merchant account

Returns the merchant account specified in the path. Your API credential must have access to the merchant account.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Account read

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountMerchantLevelApiGetMerchantAccountInput - Request parameters, see GetMerchantAccountInput
@return Merchant, *http.Response, error
*/
func (a *AccountMerchantLevelApi) GetMerchantAccount(ctx context.Context, r AccountMerchantLevelApiGetMerchantAccountInput) (Merchant, *http.Response, error) {
	res := &Merchant{}
	path := "/merchants/{merchantId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
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

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

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

// All parameters accepted by AccountMerchantLevelApi.ListMerchantAccounts
type AccountMerchantLevelApiListMerchantAccountsInput struct {
	pageNumber *int32
	pageSize   *int32
}

// The number of the page to fetch.
func (r AccountMerchantLevelApiListMerchantAccountsInput) PageNumber(pageNumber int32) AccountMerchantLevelApiListMerchantAccountsInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r AccountMerchantLevelApiListMerchantAccountsInput) PageSize(pageSize int32) AccountMerchantLevelApiListMerchantAccountsInput {
	r.pageSize = &pageSize
	return r
}

/*
Prepare a request for ListMerchantAccounts

@return AccountMerchantLevelApiListMerchantAccountsInput
*/
func (a *AccountMerchantLevelApi) ListMerchantAccountsInput() AccountMerchantLevelApiListMerchantAccountsInput {
	return AccountMerchantLevelApiListMerchantAccountsInput{}
}

/*
ListMerchantAccounts Get a list of merchant accounts

Returns the list of merchant accounts that your API credential has access to. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Account read

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountMerchantLevelApiListMerchantAccountsInput - Request parameters, see ListMerchantAccountsInput
@return ListMerchantResponse, *http.Response, error
*/
func (a *AccountMerchantLevelApi) ListMerchantAccounts(ctx context.Context, r AccountMerchantLevelApiListMerchantAccountsInput) (ListMerchantResponse, *http.Response, error) {
	res := &ListMerchantResponse{}
	path := "/merchants"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
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

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

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

// All parameters accepted by AccountMerchantLevelApi.RequestToActivateMerchantAccount
type AccountMerchantLevelApiRequestToActivateMerchantAccountInput struct {
	merchantId string
}

/*
Prepare a request for RequestToActivateMerchantAccount
@param merchantId The unique identifier of the merchant account.
@return AccountMerchantLevelApiRequestToActivateMerchantAccountInput
*/
func (a *AccountMerchantLevelApi) RequestToActivateMerchantAccountInput(merchantId string) AccountMerchantLevelApiRequestToActivateMerchantAccountInput {
	return AccountMerchantLevelApiRequestToActivateMerchantAccountInput{
		merchantId: merchantId,
	}
}

/*
RequestToActivateMerchantAccount Request to activate a merchant account

Sends a request to activate the merchant account identified in the path.

You get the result of the activation asychronously through a [`merchant.updated`](https://docs.adyen.com/api-explorer/ManagementNotification/latest/post/merchant.updated) webhook. Once the merchant account is activated, you can start using it to accept payments and payouts.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Accounts read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountMerchantLevelApiRequestToActivateMerchantAccountInput - Request parameters, see RequestToActivateMerchantAccountInput
@return RequestActivationResponse, *http.Response, error
*/
func (a *AccountMerchantLevelApi) RequestToActivateMerchantAccount(ctx context.Context, r AccountMerchantLevelApiRequestToActivateMerchantAccountInput) (RequestActivationResponse, *http.Response, error) {
	res := &RequestActivationResponse{}
	path := "/merchants/{merchantId}/activate"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

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