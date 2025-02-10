/*
Management API

API version: 3
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

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// AccountCompanyLevelApi service
type AccountCompanyLevelApi common.Service

// All parameters accepted by AccountCompanyLevelApi.GetCompanyAccount
type AccountCompanyLevelApiGetCompanyAccountInput struct {
	companyId string
}

/*
Prepare a request for GetCompanyAccount
@param companyId The unique identifier of the company account.
@return AccountCompanyLevelApiGetCompanyAccountInput
*/
func (a *AccountCompanyLevelApi) GetCompanyAccountInput(companyId string) AccountCompanyLevelApiGetCompanyAccountInput {
	return AccountCompanyLevelApiGetCompanyAccountInput{
		companyId: companyId,
	}
}

/*
GetCompanyAccount Get a company account

Returns the company account specified in the path. Your API credential must have access to the company account.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Account read

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountCompanyLevelApiGetCompanyAccountInput - Request parameters, see GetCompanyAccountInput
@return Company, *http.Response, error
*/
func (a *AccountCompanyLevelApi) GetCompanyAccount(ctx context.Context, r AccountCompanyLevelApiGetCompanyAccountInput) (Company, *http.Response, error) {
	res := &Company{}
	path := "/companies/{companyId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
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

	if httpRes == nil {
		return *res, httpRes, err
	}

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

// All parameters accepted by AccountCompanyLevelApi.ListCompanyAccounts
type AccountCompanyLevelApiListCompanyAccountsInput struct {
	pageNumber *int32
	pageSize   *int32
}

// The number of the page to fetch.
func (r AccountCompanyLevelApiListCompanyAccountsInput) PageNumber(pageNumber int32) AccountCompanyLevelApiListCompanyAccountsInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r AccountCompanyLevelApiListCompanyAccountsInput) PageSize(pageSize int32) AccountCompanyLevelApiListCompanyAccountsInput {
	r.pageSize = &pageSize
	return r
}

/*
Prepare a request for ListCompanyAccounts

@return AccountCompanyLevelApiListCompanyAccountsInput
*/
func (a *AccountCompanyLevelApi) ListCompanyAccountsInput() AccountCompanyLevelApiListCompanyAccountsInput {
	return AccountCompanyLevelApiListCompanyAccountsInput{}
}

/*
ListCompanyAccounts Get a list of company accounts

Returns the list of company accounts that your API credential has access to. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):

* Management API—Account read

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountCompanyLevelApiListCompanyAccountsInput - Request parameters, see ListCompanyAccountsInput
@return ListCompanyResponse, *http.Response, error
*/
func (a *AccountCompanyLevelApi) ListCompanyAccounts(ctx context.Context, r AccountCompanyLevelApiListCompanyAccountsInput) (ListCompanyResponse, *http.Response, error) {
	res := &ListCompanyResponse{}
	path := "/companies"
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

	if httpRes == nil {
		return *res, httpRes, err
	}

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

// All parameters accepted by AccountCompanyLevelApi.ListMerchantAccounts
type AccountCompanyLevelApiListMerchantAccountsInput struct {
	companyId  string
	pageNumber *int32
	pageSize   *int32
}

// The number of the page to fetch.
func (r AccountCompanyLevelApiListMerchantAccountsInput) PageNumber(pageNumber int32) AccountCompanyLevelApiListMerchantAccountsInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r AccountCompanyLevelApiListMerchantAccountsInput) PageSize(pageSize int32) AccountCompanyLevelApiListMerchantAccountsInput {
	r.pageSize = &pageSize
	return r
}

/*
Prepare a request for ListMerchantAccounts
@param companyId The unique identifier of the company account.
@return AccountCompanyLevelApiListMerchantAccountsInput
*/
func (a *AccountCompanyLevelApi) ListMerchantAccountsInput(companyId string) AccountCompanyLevelApiListMerchantAccountsInput {
	return AccountCompanyLevelApiListMerchantAccountsInput{
		companyId: companyId,
	}
}

/*
ListMerchantAccounts Get a list of merchant accounts

Returns the list of merchant accounts under the company account specified in the path. The list only includes merchant accounts that your API credential has access to. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Account read

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountCompanyLevelApiListMerchantAccountsInput - Request parameters, see ListMerchantAccountsInput
@return ListMerchantResponse, *http.Response, error
*/
func (a *AccountCompanyLevelApi) ListMerchantAccounts(ctx context.Context, r AccountCompanyLevelApiListMerchantAccountsInput) (ListMerchantResponse, *http.Response, error) {
	res := &ListMerchantResponse{}
	path := "/companies/{companyId}/merchants"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
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

	if httpRes == nil {
		return *res, httpRes, err
	}

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
