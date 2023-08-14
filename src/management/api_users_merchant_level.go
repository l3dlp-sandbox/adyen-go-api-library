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

// UsersMerchantLevelApi service
type UsersMerchantLevelApi common.Service

// All parameters accepted by UsersMerchantLevelApi.CreateNewUser
type UsersMerchantLevelApiCreateNewUserInput struct {
	merchantId                string
	createMerchantUserRequest *CreateMerchantUserRequest
}

func (r UsersMerchantLevelApiCreateNewUserInput) CreateMerchantUserRequest(createMerchantUserRequest CreateMerchantUserRequest) UsersMerchantLevelApiCreateNewUserInput {
	r.createMerchantUserRequest = &createMerchantUserRequest
	return r
}

/*
Prepare a request for CreateNewUser
@param merchantId Unique identifier of the merchant.
@return UsersMerchantLevelApiCreateNewUserInput
*/
func (a *UsersMerchantLevelApi) CreateNewUserInput(merchantId string) UsersMerchantLevelApiCreateNewUserInput {
	return UsersMerchantLevelApiCreateNewUserInput{
		merchantId: merchantId,
	}
}

/*
CreateNewUser Create a new user

Creates a user for the `merchantId` specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersMerchantLevelApiCreateNewUserInput - Request parameters, see CreateNewUserInput
@return CreateUserResponse, *http.Response, error
*/
func (a *UsersMerchantLevelApi) CreateNewUser(ctx context.Context, r UsersMerchantLevelApiCreateNewUserInput) (CreateUserResponse, *http.Response, error) {
	res := &CreateUserResponse{}
	path := "/merchants/{merchantId}/users"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createMerchantUserRequest,
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

// All parameters accepted by UsersMerchantLevelApi.GetUserDetails
type UsersMerchantLevelApiGetUserDetailsInput struct {
	merchantId string
	userId     string
}

/*
Prepare a request for GetUserDetails
@param merchantId Unique identifier of the merchant.@param userId Unique identifier of the user.
@return UsersMerchantLevelApiGetUserDetailsInput
*/
func (a *UsersMerchantLevelApi) GetUserDetailsInput(merchantId string, userId string) UsersMerchantLevelApiGetUserDetailsInput {
	return UsersMerchantLevelApiGetUserDetailsInput{
		merchantId: merchantId,
		userId:     userId,
	}
}

/*
GetUserDetails Get user details

Returns user details for the `userId` and the `merchantId` specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersMerchantLevelApiGetUserDetailsInput - Request parameters, see GetUserDetailsInput
@return User, *http.Response, error
*/
func (a *UsersMerchantLevelApi) GetUserDetails(ctx context.Context, r UsersMerchantLevelApiGetUserDetailsInput) (User, *http.Response, error) {
	res := &User{}
	path := "/merchants/{merchantId}/users/{userId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"userId"+"}", url.PathEscape(common.ParameterValueToString(r.userId, "userId")), -1)
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

// All parameters accepted by UsersMerchantLevelApi.ListUsers
type UsersMerchantLevelApiListUsersInput struct {
	merchantId string
	pageNumber *int32
	pageSize   *int32
	username   *string
}

// The number of the page to fetch.
func (r UsersMerchantLevelApiListUsersInput) PageNumber(pageNumber int32) UsersMerchantLevelApiListUsersInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page. Maximum value is **100**. The default is **10** items on a page.
func (r UsersMerchantLevelApiListUsersInput) PageSize(pageSize int32) UsersMerchantLevelApiListUsersInput {
	r.pageSize = &pageSize
	return r
}

// The partial or complete username to select all users that match.
func (r UsersMerchantLevelApiListUsersInput) Username(username string) UsersMerchantLevelApiListUsersInput {
	r.username = &username
	return r
}

/*
Prepare a request for ListUsers
@param merchantId Unique identifier of the merchant.
@return UsersMerchantLevelApiListUsersInput
*/
func (a *UsersMerchantLevelApi) ListUsersInput(merchantId string) UsersMerchantLevelApiListUsersInput {
	return UsersMerchantLevelApiListUsersInput{
		merchantId: merchantId,
	}
}

/*
ListUsers Get a list of users

Returns a list of users associated with the `merchantId` specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersMerchantLevelApiListUsersInput - Request parameters, see ListUsersInput
@return ListMerchantUsersResponse, *http.Response, error
*/
func (a *UsersMerchantLevelApi) ListUsers(ctx context.Context, r UsersMerchantLevelApiListUsersInput) (ListMerchantUsersResponse, *http.Response, error) {
	res := &ListMerchantUsersResponse{}
	path := "/merchants/{merchantId}/users"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	if r.username != nil {
		common.ParameterAddToQuery(queryParams, "username", r.username, "")
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

// All parameters accepted by UsersMerchantLevelApi.UpdateUser
type UsersMerchantLevelApiUpdateUserInput struct {
	merchantId                string
	userId                    string
	updateMerchantUserRequest *UpdateMerchantUserRequest
}

func (r UsersMerchantLevelApiUpdateUserInput) UpdateMerchantUserRequest(updateMerchantUserRequest UpdateMerchantUserRequest) UsersMerchantLevelApiUpdateUserInput {
	r.updateMerchantUserRequest = &updateMerchantUserRequest
	return r
}

/*
Prepare a request for UpdateUser
@param merchantId Unique identifier of the merchant.@param userId Unique identifier of the user.
@return UsersMerchantLevelApiUpdateUserInput
*/
func (a *UsersMerchantLevelApi) UpdateUserInput(merchantId string, userId string) UsersMerchantLevelApiUpdateUserInput {
	return UsersMerchantLevelApiUpdateUserInput{
		merchantId: merchantId,
		userId:     userId,
	}
}

/*
UpdateUser Update a user

Updates user details for the `userId` and the `merchantId` specified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Users read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r UsersMerchantLevelApiUpdateUserInput - Request parameters, see UpdateUserInput
@return User, *http.Response, error
*/
func (a *UsersMerchantLevelApi) UpdateUser(ctx context.Context, r UsersMerchantLevelApiUpdateUserInput) (User, *http.Response, error) {
	res := &User{}
	path := "/merchants/{merchantId}/users/{userId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"userId"+"}", url.PathEscape(common.ParameterValueToString(r.userId, "userId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateMerchantUserRequest,
		res,
		http.MethodPatch,
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