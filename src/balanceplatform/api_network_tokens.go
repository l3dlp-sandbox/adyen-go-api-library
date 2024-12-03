/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
    "net/http"
    "net/url"
    "strings"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// NetworkTokensApi service
type NetworkTokensApi common.Service

// All parameters accepted by NetworkTokensApi.GetNetworkToken
type NetworkTokensApiGetNetworkTokenInput struct {
	networkTokenId string
}


/*
Prepare a request for GetNetworkToken
@param networkTokenId The unique identifier of the network token.
@return NetworkTokensApiGetNetworkTokenInput
*/
func (a *NetworkTokensApi) GetNetworkTokenInput(networkTokenId string) NetworkTokensApiGetNetworkTokenInput {
	return NetworkTokensApiGetNetworkTokenInput{
		networkTokenId: networkTokenId,
	}
}

/*
GetNetworkToken Get a network token

Returns the details of a network token.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r NetworkTokensApiGetNetworkTokenInput - Request parameters, see GetNetworkTokenInput
@return GetNetworkTokenResponse, *http.Response, error
*/
func (a *NetworkTokensApi) GetNetworkToken(ctx context.Context, r NetworkTokensApiGetNetworkTokenInput) (GetNetworkTokenResponse, *http.Response, error) {
    res := &GetNetworkTokenResponse{}
	path := "/networkTokens/{networkTokenId}"
    path = strings.Replace(path, "{"+"networkTokenId"+"}", url.PathEscape(common.ParameterValueToString(r.networkTokenId, "networkTokenId")), -1)
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


// All parameters accepted by NetworkTokensApi.UpdateNetworkToken
type NetworkTokensApiUpdateNetworkTokenInput struct {
	networkTokenId string
	updateNetworkTokenRequest *UpdateNetworkTokenRequest
}

func (r NetworkTokensApiUpdateNetworkTokenInput) UpdateNetworkTokenRequest(updateNetworkTokenRequest UpdateNetworkTokenRequest) NetworkTokensApiUpdateNetworkTokenInput {
	r.updateNetworkTokenRequest = &updateNetworkTokenRequest
	return r
}


/*
Prepare a request for UpdateNetworkToken
@param networkTokenId The unique identifier of the network token.
@return NetworkTokensApiUpdateNetworkTokenInput
*/
func (a *NetworkTokensApi) UpdateNetworkTokenInput(networkTokenId string) NetworkTokensApiUpdateNetworkTokenInput {
	return NetworkTokensApiUpdateNetworkTokenInput{
		networkTokenId: networkTokenId,
	}
}

/*
UpdateNetworkToken Update a network token

Updates the status of the network token.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r NetworkTokensApiUpdateNetworkTokenInput - Request parameters, see UpdateNetworkTokenInput
@return *http.Response, error
*/
func (a *NetworkTokensApi) UpdateNetworkToken(ctx context.Context, r NetworkTokensApiUpdateNetworkTokenInput) (*http.Response, error) {
    var res interface{}
	path := "/networkTokens/{networkTokenId}"
    path = strings.Replace(path, "{"+"networkTokenId"+"}", url.PathEscape(common.ParameterValueToString(r.networkTokenId, "networkTokenId")), -1)
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        r.updateNetworkTokenRequest,
        res,
        http.MethodPatch,
        a.BasePath()+path,
        queryParams,
        headerParams,
    )

    if httpRes == nil {
        return httpRes, err
    }

    var serviceError common.RestServiceError
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

