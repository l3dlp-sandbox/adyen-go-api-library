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

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// ClientKeyCompanyLevelApi service
type ClientKeyCompanyLevelApi common.Service

// All parameters accepted by ClientKeyCompanyLevelApi.GenerateNewClientKey
type ClientKeyCompanyLevelApiGenerateNewClientKeyInput struct {
	companyId       string
	apiCredentialId string
}

/*
Prepare a request for GenerateNewClientKey
@param companyId The unique identifier of the company account.@param apiCredentialId Unique identifier of the API credential.
@return ClientKeyCompanyLevelApiGenerateNewClientKeyInput
*/
func (a *ClientKeyCompanyLevelApi) GenerateNewClientKeyInput(companyId string, apiCredentialId string) ClientKeyCompanyLevelApiGenerateNewClientKeyInput {
	return ClientKeyCompanyLevelApiGenerateNewClientKeyInput{
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
	}
}

/*
GenerateNewClientKey Generate new client key

Returns a new [client key](https://docs.adyen.com/development-resources/client-side-authentication#how-it-works) for the API credential identified in the path. You can use the new client key a few minutes after generating it. The old client key stops working 24 hours after generating a new one.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ClientKeyCompanyLevelApiGenerateNewClientKeyInput - Request parameters, see GenerateNewClientKeyInput
@return GenerateClientKeyResponse, *http.Response, error
*/
func (a *ClientKeyCompanyLevelApi) GenerateNewClientKey(ctx context.Context, r ClientKeyCompanyLevelApiGenerateNewClientKeyInput) (GenerateClientKeyResponse, *http.Response, error) {
	res := &GenerateClientKeyResponse{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/generateClientKey"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
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
