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

// MyAPICredentialApi service
type MyAPICredentialApi common.Service

// All parameters accepted by MyAPICredentialApi.AddAllowedOrigin
type MyAPICredentialApiAddAllowedOriginInput struct {
	createAllowedOriginRequest *CreateAllowedOriginRequest
}

func (r MyAPICredentialApiAddAllowedOriginInput) CreateAllowedOriginRequest(createAllowedOriginRequest CreateAllowedOriginRequest) MyAPICredentialApiAddAllowedOriginInput {
	r.createAllowedOriginRequest = &createAllowedOriginRequest
	return r
}

/*
Prepare a request for AddAllowedOrigin

@return MyAPICredentialApiAddAllowedOriginInput
*/
func (a *MyAPICredentialApi) AddAllowedOriginInput() MyAPICredentialApiAddAllowedOriginInput {
	return MyAPICredentialApiAddAllowedOriginInput{}
}

/*
AddAllowedOrigin Add allowed origin

Adds an allowed origin to the list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) of your API credential.
The API key from the request is used to identify the [API credential](https://docs.adyen.com/development-resources/api-credentials).

You can make this request with any of the Management API roles.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r MyAPICredentialApiAddAllowedOriginInput - Request parameters, see AddAllowedOriginInput
@return AllowedOrigin, *http.Response, error
*/
func (a *MyAPICredentialApi) AddAllowedOrigin(ctx context.Context, r MyAPICredentialApiAddAllowedOriginInput) (AllowedOrigin, *http.Response, error) {
	res := &AllowedOrigin{}
	path := "/me/allowedOrigins"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createAllowedOriginRequest,
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

// All parameters accepted by MyAPICredentialApi.GenerateClientKey
type MyAPICredentialApiGenerateClientKeyInput struct {
}

/*
Prepare a request for GenerateClientKey

@return MyAPICredentialApiGenerateClientKeyInput
*/
func (a *MyAPICredentialApi) GenerateClientKeyInput() MyAPICredentialApiGenerateClientKeyInput {
	return MyAPICredentialApiGenerateClientKeyInput{}
}

/*
GenerateClientKey Generate a client key

Generates a new [client key](https://docs.adyen.com/development-resources/client-side-authentication/) used to authenticate requests from your payment environment.
You can use the new client key a few minutes after generating it.
The old client key stops working 24 hours after generating a new one.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r MyAPICredentialApiGenerateClientKeyInput - Request parameters, see GenerateClientKeyInput
@return GenerateClientKeyResponse, *http.Response, error
*/
func (a *MyAPICredentialApi) GenerateClientKey(ctx context.Context, r MyAPICredentialApiGenerateClientKeyInput) (GenerateClientKeyResponse, *http.Response, error) {
	res := &GenerateClientKeyResponse{}
	path := "/me/generateClientKey"
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

// All parameters accepted by MyAPICredentialApi.GetAllowedOriginDetails
type MyAPICredentialApiGetAllowedOriginDetailsInput struct {
	originId string
}

/*
Prepare a request for GetAllowedOriginDetails
@param originId Unique identifier of the allowed origin.
@return MyAPICredentialApiGetAllowedOriginDetailsInput
*/
func (a *MyAPICredentialApi) GetAllowedOriginDetailsInput(originId string) MyAPICredentialApiGetAllowedOriginDetailsInput {
	return MyAPICredentialApiGetAllowedOriginDetailsInput{
		originId: originId,
	}
}

/*
GetAllowedOriginDetails Get allowed origin details

Returns the details of the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) specified in the path.
The API key from the request is used to identify the [API credential](https://docs.adyen.com/development-resources/api-credentials).

You can make this request with any of the Management API roles.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r MyAPICredentialApiGetAllowedOriginDetailsInput - Request parameters, see GetAllowedOriginDetailsInput
@return AllowedOrigin, *http.Response, error
*/
func (a *MyAPICredentialApi) GetAllowedOriginDetails(ctx context.Context, r MyAPICredentialApiGetAllowedOriginDetailsInput) (AllowedOrigin, *http.Response, error) {
	res := &AllowedOrigin{}
	path := "/me/allowedOrigins/{originId}"
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

// All parameters accepted by MyAPICredentialApi.GetAllowedOrigins
type MyAPICredentialApiGetAllowedOriginsInput struct {
}

/*
Prepare a request for GetAllowedOrigins

@return MyAPICredentialApiGetAllowedOriginsInput
*/
func (a *MyAPICredentialApi) GetAllowedOriginsInput() MyAPICredentialApiGetAllowedOriginsInput {
	return MyAPICredentialApiGetAllowedOriginsInput{}
}

/*
GetAllowedOrigins Get allowed origins

Returns the list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) of your [API credential](https://docs.adyen.com/development-resources/api-credentials) based on the API key you used in the request.

You can make this request with any of the Management API roles.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r MyAPICredentialApiGetAllowedOriginsInput - Request parameters, see GetAllowedOriginsInput
@return AllowedOriginsResponse, *http.Response, error
*/
func (a *MyAPICredentialApi) GetAllowedOrigins(ctx context.Context, r MyAPICredentialApiGetAllowedOriginsInput) (AllowedOriginsResponse, *http.Response, error) {
	res := &AllowedOriginsResponse{}
	path := "/me/allowedOrigins"
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

// All parameters accepted by MyAPICredentialApi.GetApiCredentialDetails
type MyAPICredentialApiGetApiCredentialDetailsInput struct {
}

/*
Prepare a request for GetApiCredentialDetails

@return MyAPICredentialApiGetApiCredentialDetailsInput
*/
func (a *MyAPICredentialApi) GetApiCredentialDetailsInput() MyAPICredentialApiGetApiCredentialDetailsInput {
	return MyAPICredentialApiGetApiCredentialDetailsInput{}
}

/*
GetApiCredentialDetails Get API credential details

Returns your [API credential](https://docs.adyen.com/development-resources/api-credentials) details based on the API Key you used in the request.

You can make this request with any of the Management API roles.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r MyAPICredentialApiGetApiCredentialDetailsInput - Request parameters, see GetApiCredentialDetailsInput
@return MeApiCredential, *http.Response, error
*/
func (a *MyAPICredentialApi) GetApiCredentialDetails(ctx context.Context, r MyAPICredentialApiGetApiCredentialDetailsInput) (MeApiCredential, *http.Response, error) {
	res := &MeApiCredential{}
	path := "/me"
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

// All parameters accepted by MyAPICredentialApi.RemoveAllowedOrigin
type MyAPICredentialApiRemoveAllowedOriginInput struct {
	originId string
}

/*
Prepare a request for RemoveAllowedOrigin
@param originId Unique identifier of the allowed origin.
@return MyAPICredentialApiRemoveAllowedOriginInput
*/
func (a *MyAPICredentialApi) RemoveAllowedOriginInput(originId string) MyAPICredentialApiRemoveAllowedOriginInput {
	return MyAPICredentialApiRemoveAllowedOriginInput{
		originId: originId,
	}
}

/*
RemoveAllowedOrigin Remove allowed origin

Removes the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) specified in the path.
The API key from the request is used to identify the [API credential](https://docs.adyen.com/development-resources/api-credentials).

You can make this request with any of the Management API roles.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r MyAPICredentialApiRemoveAllowedOriginInput - Request parameters, see RemoveAllowedOriginInput
@return *http.Response, error
*/
func (a *MyAPICredentialApi) RemoveAllowedOrigin(ctx context.Context, r MyAPICredentialApiRemoveAllowedOriginInput) (*http.Response, error) {
	var res interface{}
	path := "/me/allowedOrigins/{originId}"
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
