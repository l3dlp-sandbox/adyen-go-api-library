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

	"github.com/adyen/adyen-go-api-library/v18/src/common"
)

// BusinessLinesApi service
type BusinessLinesApi common.Service

// All parameters accepted by BusinessLinesApi.CreateBusinessLine
type BusinessLinesApiCreateBusinessLineInput struct {
	businessLineInfo *BusinessLineInfo
}

func (r BusinessLinesApiCreateBusinessLineInput) BusinessLineInfo(businessLineInfo BusinessLineInfo) BusinessLinesApiCreateBusinessLineInput {
	r.businessLineInfo = &businessLineInfo
	return r
}

/*
Prepare a request for CreateBusinessLine

@return BusinessLinesApiCreateBusinessLineInput
*/
func (a *BusinessLinesApi) CreateBusinessLineInput() BusinessLinesApiCreateBusinessLineInput {
	return BusinessLinesApiCreateBusinessLineInput{}
}

/*
CreateBusinessLine Create a business line

Creates a business line.

This resource contains information about your user's line of business, including their industry and their source of funds. Adyen uses this information to verify your users as required by payment industry regulations.Adyen informs you of the verification results through webhooks or API responses.

You can create a maximum of 3 business lines per legal entity for payment processing.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BusinessLinesApiCreateBusinessLineInput - Request parameters, see CreateBusinessLineInput
@return BusinessLine, *http.Response, error
*/
func (a *BusinessLinesApi) CreateBusinessLine(ctx context.Context, r BusinessLinesApiCreateBusinessLineInput) (BusinessLine, *http.Response, error) {
	res := &BusinessLine{}
	path := "/businessLines"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.businessLineInfo,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by BusinessLinesApi.DeleteBusinessLine
type BusinessLinesApiDeleteBusinessLineInput struct {
	id string
}

/*
Prepare a request for DeleteBusinessLine
@param id The unique identifier of the business line to be deleted.
@return BusinessLinesApiDeleteBusinessLineInput
*/
func (a *BusinessLinesApi) DeleteBusinessLineInput(id string) BusinessLinesApiDeleteBusinessLineInput {
	return BusinessLinesApiDeleteBusinessLineInput{
		id: id,
	}
}

/*
DeleteBusinessLine Delete a business line

Deletes a business line.

 >If you delete a business line linked to a [payment method](https://docs.adyen.com/development-resources/paymentmethodvariant#management-api), it can affect your merchant account's ability to use the [payment method](https://docs.adyen.com/api-explorer/Management/latest/post/merchants/_merchantId_/paymentMethodSettings). The business line is removed from all linked merchant accounts.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BusinessLinesApiDeleteBusinessLineInput - Request parameters, see DeleteBusinessLineInput
@return *http.Response, error
*/
func (a *BusinessLinesApi) DeleteBusinessLine(ctx context.Context, r BusinessLinesApiDeleteBusinessLineInput) (*http.Response, error) {
	var res interface{}
	path := "/businessLines/{id}"
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

// All parameters accepted by BusinessLinesApi.GetBusinessLine
type BusinessLinesApiGetBusinessLineInput struct {
	id string
}

/*
Prepare a request for GetBusinessLine
@param id The unique identifier of the business line.
@return BusinessLinesApiGetBusinessLineInput
*/
func (a *BusinessLinesApi) GetBusinessLineInput(id string) BusinessLinesApiGetBusinessLineInput {
	return BusinessLinesApiGetBusinessLineInput{
		id: id,
	}
}

/*
GetBusinessLine Get a business line

Returns the detail of a business line.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BusinessLinesApiGetBusinessLineInput - Request parameters, see GetBusinessLineInput
@return BusinessLine, *http.Response, error
*/
func (a *BusinessLinesApi) GetBusinessLine(ctx context.Context, r BusinessLinesApiGetBusinessLineInput) (BusinessLine, *http.Response, error) {
	res := &BusinessLine{}
	path := "/businessLines/{id}"
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

// All parameters accepted by BusinessLinesApi.UpdateBusinessLine
type BusinessLinesApiUpdateBusinessLineInput struct {
	id                     string
	businessLineInfoUpdate *BusinessLineInfoUpdate
}

func (r BusinessLinesApiUpdateBusinessLineInput) BusinessLineInfoUpdate(businessLineInfoUpdate BusinessLineInfoUpdate) BusinessLinesApiUpdateBusinessLineInput {
	r.businessLineInfoUpdate = &businessLineInfoUpdate
	return r
}

/*
Prepare a request for UpdateBusinessLine
@param id The unique identifier of the business line.
@return BusinessLinesApiUpdateBusinessLineInput
*/
func (a *BusinessLinesApi) UpdateBusinessLineInput(id string) BusinessLinesApiUpdateBusinessLineInput {
	return BusinessLinesApiUpdateBusinessLineInput{
		id: id,
	}
}

/*
UpdateBusinessLine Update a business line

Updates a business line.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BusinessLinesApiUpdateBusinessLineInput - Request parameters, see UpdateBusinessLineInput
@return BusinessLine, *http.Response, error
*/
func (a *BusinessLinesApi) UpdateBusinessLine(ctx context.Context, r BusinessLinesApiUpdateBusinessLineInput) (BusinessLine, *http.Response, error) {
	res := &BusinessLine{}
	path := "/businessLines/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.businessLineInfoUpdate,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
