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

	"github.com/adyen/adyen-go-api-library/v10/src/common"
)

// AllowedOriginsMerchantLevelApi service
type AllowedOriginsMerchantLevelApi common.Service

// All parameters accepted by AllowedOriginsMerchantLevelApi.CreateAllowedOrigin
type AllowedOriginsMerchantLevelApiCreateAllowedOriginInput struct {
	merchantId      string
	apiCredentialId string
	allowedOrigin   *AllowedOrigin
}

func (r AllowedOriginsMerchantLevelApiCreateAllowedOriginInput) AllowedOrigin(allowedOrigin AllowedOrigin) AllowedOriginsMerchantLevelApiCreateAllowedOriginInput {
	r.allowedOrigin = &allowedOrigin
	return r
}

/*
Prepare a request for CreateAllowedOrigin
@param merchantId The unique identifier of the merchant account.@param apiCredentialId Unique identifier of the API credential.
@return AllowedOriginsMerchantLevelApiCreateAllowedOriginInput
*/
func (a *AllowedOriginsMerchantLevelApi) CreateAllowedOriginInput(merchantId string, apiCredentialId string) AllowedOriginsMerchantLevelApiCreateAllowedOriginInput {
	return AllowedOriginsMerchantLevelApiCreateAllowedOriginInput{
		merchantId:      merchantId,
		apiCredentialId: apiCredentialId,
	}
}

/*
CreateAllowedOrigin Create an allowed origin

Adds a new [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) to the API credential's list of allowed origins.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsMerchantLevelApiCreateAllowedOriginInput - Request parameters, see CreateAllowedOriginInput
@return AllowedOrigin, *http.Response, error
*/
func (a *AllowedOriginsMerchantLevelApi) CreateAllowedOrigin(ctx context.Context, r AllowedOriginsMerchantLevelApiCreateAllowedOriginInput) (AllowedOrigin, *http.Response, error) {
	res := &AllowedOrigin{}
	path := "/merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.allowedOrigin,
		res,
		http.MethodPost,
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

// All parameters accepted by AllowedOriginsMerchantLevelApi.DeleteAllowedOrigin
type AllowedOriginsMerchantLevelApiDeleteAllowedOriginInput struct {
	merchantId      string
	apiCredentialId string
	originId        string
}

/*
Prepare a request for DeleteAllowedOrigin
@param merchantId The unique identifier of the merchant account.@param apiCredentialId Unique identifier of the API credential.@param originId Unique identifier of the allowed origin.
@return AllowedOriginsMerchantLevelApiDeleteAllowedOriginInput
*/
func (a *AllowedOriginsMerchantLevelApi) DeleteAllowedOriginInput(merchantId string, apiCredentialId string, originId string) AllowedOriginsMerchantLevelApiDeleteAllowedOriginInput {
	return AllowedOriginsMerchantLevelApiDeleteAllowedOriginInput{
		merchantId:      merchantId,
		apiCredentialId: apiCredentialId,
		originId:        originId,
	}
}

/*
DeleteAllowedOrigin Delete an allowed origin

Removes the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path. As soon as an allowed origin is removed, we no longer accept client-side requests from that domain.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsMerchantLevelApiDeleteAllowedOriginInput - Request parameters, see DeleteAllowedOriginInput
@return *http.Response, error
*/
func (a *AllowedOriginsMerchantLevelApi) DeleteAllowedOrigin(ctx context.Context, r AllowedOriginsMerchantLevelApiDeleteAllowedOriginInput) (*http.Response, error) {
	var res interface{}
	path := "/merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	path = strings.Replace(path, "{"+"originId"+"}", url.PathEscape(common.ParameterValueToString(r.originId, "originId")), -1)
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

	if httpRes == nil {
		return httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	return httpRes, err
}

// All parameters accepted by AllowedOriginsMerchantLevelApi.GetAllowedOrigin
type AllowedOriginsMerchantLevelApiGetAllowedOriginInput struct {
	merchantId      string
	apiCredentialId string
	originId        string
}

/*
Prepare a request for GetAllowedOrigin
@param merchantId The unique identifier of the merchant account.@param apiCredentialId Unique identifier of the API credential.@param originId Unique identifier of the allowed origin.
@return AllowedOriginsMerchantLevelApiGetAllowedOriginInput
*/
func (a *AllowedOriginsMerchantLevelApi) GetAllowedOriginInput(merchantId string, apiCredentialId string, originId string) AllowedOriginsMerchantLevelApiGetAllowedOriginInput {
	return AllowedOriginsMerchantLevelApiGetAllowedOriginInput{
		merchantId:      merchantId,
		apiCredentialId: apiCredentialId,
		originId:        originId,
	}
}

/*
GetAllowedOrigin Get an allowed origin

Returns the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsMerchantLevelApiGetAllowedOriginInput - Request parameters, see GetAllowedOriginInput
@return AllowedOrigin, *http.Response, error
*/
func (a *AllowedOriginsMerchantLevelApi) GetAllowedOrigin(ctx context.Context, r AllowedOriginsMerchantLevelApiGetAllowedOriginInput) (AllowedOrigin, *http.Response, error) {
	res := &AllowedOrigin{}
	path := "/merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	path = strings.Replace(path, "{"+"originId"+"}", url.PathEscape(common.ParameterValueToString(r.originId, "originId")), -1)
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

// All parameters accepted by AllowedOriginsMerchantLevelApi.ListAllowedOrigins
type AllowedOriginsMerchantLevelApiListAllowedOriginsInput struct {
	merchantId      string
	apiCredentialId string
}

/*
Prepare a request for ListAllowedOrigins
@param merchantId The unique identifier of the merchant account.@param apiCredentialId Unique identifier of the API credential.
@return AllowedOriginsMerchantLevelApiListAllowedOriginsInput
*/
func (a *AllowedOriginsMerchantLevelApi) ListAllowedOriginsInput(merchantId string, apiCredentialId string) AllowedOriginsMerchantLevelApiListAllowedOriginsInput {
	return AllowedOriginsMerchantLevelApiListAllowedOriginsInput{
		merchantId:      merchantId,
		apiCredentialId: apiCredentialId,
	}
}

/*
ListAllowedOrigins Get a list of allowed origins

Returns the list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsMerchantLevelApiListAllowedOriginsInput - Request parameters, see ListAllowedOriginsInput
@return AllowedOriginsResponse, *http.Response, error
*/
func (a *AllowedOriginsMerchantLevelApi) ListAllowedOrigins(ctx context.Context, r AllowedOriginsMerchantLevelApiListAllowedOriginsInput) (AllowedOriginsResponse, *http.Response, error) {
	res := &AllowedOriginsResponse{}
	path := "/merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
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
