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

// LegalEntitiesApi service
type LegalEntitiesApi common.Service

// All parameters accepted by LegalEntitiesApi.CheckLegalEntitysVerificationErrors
type LegalEntitiesApiCheckLegalEntitysVerificationErrorsInput struct {
	id string
}

/*
Prepare a request for CheckLegalEntitysVerificationErrors
@param id The unique identifier of the legal entity.
@return LegalEntitiesApiCheckLegalEntitysVerificationErrorsInput
*/
func (a *LegalEntitiesApi) CheckLegalEntitysVerificationErrorsInput(id string) LegalEntitiesApiCheckLegalEntitysVerificationErrorsInput {
	return LegalEntitiesApiCheckLegalEntitysVerificationErrorsInput{
		id: id,
	}
}

/*
CheckLegalEntitysVerificationErrors Check a legal entity's verification errors

Returns the verification errors for a legal entity and its supporting entities.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r LegalEntitiesApiCheckLegalEntitysVerificationErrorsInput - Request parameters, see CheckLegalEntitysVerificationErrorsInput
@return VerificationErrors, *http.Response, error
*/
func (a *LegalEntitiesApi) CheckLegalEntitysVerificationErrors(ctx context.Context, r LegalEntitiesApiCheckLegalEntitysVerificationErrorsInput) (VerificationErrors, *http.Response, error) {
	res := &VerificationErrors{}
	path := "/legalEntities/{id}/checkVerificationErrors"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
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

	return *res, httpRes, err
}

// All parameters accepted by LegalEntitiesApi.ConfirmDataReview
type LegalEntitiesApiConfirmDataReviewInput struct {
	id string
}

/*
Prepare a request for ConfirmDataReview
@param id The unique identifier of the legal entity.
@return LegalEntitiesApiConfirmDataReviewInput
*/
func (a *LegalEntitiesApi) ConfirmDataReviewInput(id string) LegalEntitiesApiConfirmDataReviewInput {
	return LegalEntitiesApiConfirmDataReviewInput{
		id: id,
	}
}

/*
ConfirmDataReview Confirm data review

Confirms that your user has reviewed the data for the legal entity specified in the path. Call this endpoint to inform Adyen that your user reviewed and verified that the data is up-to-date. The endpoint returns the timestamp of when Adyen received the request.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r LegalEntitiesApiConfirmDataReviewInput - Request parameters, see ConfirmDataReviewInput
@return DataReviewConfirmationResponse, *http.Response, error
*/
func (a *LegalEntitiesApi) ConfirmDataReview(ctx context.Context, r LegalEntitiesApiConfirmDataReviewInput) (DataReviewConfirmationResponse, *http.Response, error) {
	res := &DataReviewConfirmationResponse{}
	path := "/legalEntities/{id}/confirmDataReview"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
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

	return *res, httpRes, err
}

// All parameters accepted by LegalEntitiesApi.CreateLegalEntity
type LegalEntitiesApiCreateLegalEntityInput struct {
	xRequestedVerificationCode  *string
	legalEntityInfoRequiredType *LegalEntityInfoRequiredType
}

// Use a suberror code as your requested verification code. You can include one code at a time in your request header. Requested verification codes can only be used in your test environment.
func (r LegalEntitiesApiCreateLegalEntityInput) XRequestedVerificationCode(xRequestedVerificationCode string) LegalEntitiesApiCreateLegalEntityInput {
	r.xRequestedVerificationCode = &xRequestedVerificationCode
	return r
}

func (r LegalEntitiesApiCreateLegalEntityInput) LegalEntityInfoRequiredType(legalEntityInfoRequiredType LegalEntityInfoRequiredType) LegalEntitiesApiCreateLegalEntityInput {
	r.legalEntityInfoRequiredType = &legalEntityInfoRequiredType
	return r
}

/*
Prepare a request for CreateLegalEntity

@return LegalEntitiesApiCreateLegalEntityInput
*/
func (a *LegalEntitiesApi) CreateLegalEntityInput() LegalEntitiesApiCreateLegalEntityInput {
	return LegalEntitiesApiCreateLegalEntityInput{}
}

/*
CreateLegalEntity Create a legal entity

Creates a legal entity.

This resource contains information about the user that will be onboarded in your platform. Adyen uses this information to perform verification checks as required by payment industry regulations. Adyen informs you of the verification results through webhooks or API responses.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r LegalEntitiesApiCreateLegalEntityInput - Request parameters, see CreateLegalEntityInput
@return LegalEntity, *http.Response, error
*/
func (a *LegalEntitiesApi) CreateLegalEntity(ctx context.Context, r LegalEntitiesApiCreateLegalEntityInput) (LegalEntity, *http.Response, error) {
	res := &LegalEntity{}
	path := "/legalEntities"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.xRequestedVerificationCode != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "x-requested-verification-code", r.xRequestedVerificationCode, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.legalEntityInfoRequiredType,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by LegalEntitiesApi.GetAllBusinessLinesUnderLegalEntity
type LegalEntitiesApiGetAllBusinessLinesUnderLegalEntityInput struct {
	id string
}

/*
Prepare a request for GetAllBusinessLinesUnderLegalEntity
@param id The unique identifier of the legal entity.
@return LegalEntitiesApiGetAllBusinessLinesUnderLegalEntityInput
*/
func (a *LegalEntitiesApi) GetAllBusinessLinesUnderLegalEntityInput(id string) LegalEntitiesApiGetAllBusinessLinesUnderLegalEntityInput {
	return LegalEntitiesApiGetAllBusinessLinesUnderLegalEntityInput{
		id: id,
	}
}

/*
GetAllBusinessLinesUnderLegalEntity Get all business lines under a legal entity

Returns the business lines owned by a legal entity.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r LegalEntitiesApiGetAllBusinessLinesUnderLegalEntityInput - Request parameters, see GetAllBusinessLinesUnderLegalEntityInput
@return BusinessLines, *http.Response, error
*/
func (a *LegalEntitiesApi) GetAllBusinessLinesUnderLegalEntity(ctx context.Context, r LegalEntitiesApiGetAllBusinessLinesUnderLegalEntityInput) (BusinessLines, *http.Response, error) {
	res := &BusinessLines{}
	path := "/legalEntities/{id}/businessLines"
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

// All parameters accepted by LegalEntitiesApi.GetLegalEntity
type LegalEntitiesApiGetLegalEntityInput struct {
	id string
}

/*
Prepare a request for GetLegalEntity
@param id The unique identifier of the legal entity.
@return LegalEntitiesApiGetLegalEntityInput
*/
func (a *LegalEntitiesApi) GetLegalEntityInput(id string) LegalEntitiesApiGetLegalEntityInput {
	return LegalEntitiesApiGetLegalEntityInput{
		id: id,
	}
}

/*
GetLegalEntity Get a legal entity

Returns a legal entity.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r LegalEntitiesApiGetLegalEntityInput - Request parameters, see GetLegalEntityInput
@return LegalEntity, *http.Response, error
*/
func (a *LegalEntitiesApi) GetLegalEntity(ctx context.Context, r LegalEntitiesApiGetLegalEntityInput) (LegalEntity, *http.Response, error) {
	res := &LegalEntity{}
	path := "/legalEntities/{id}"
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

// All parameters accepted by LegalEntitiesApi.UpdateLegalEntity
type LegalEntitiesApiUpdateLegalEntityInput struct {
	id                         string
	xRequestedVerificationCode *string
	legalEntityInfo            *LegalEntityInfo
}

// Use the requested verification code 0_0001 to resolve any suberrors associated with the legal entity. Requested verification codes can only be used in your test environment.
func (r LegalEntitiesApiUpdateLegalEntityInput) XRequestedVerificationCode(xRequestedVerificationCode string) LegalEntitiesApiUpdateLegalEntityInput {
	r.xRequestedVerificationCode = &xRequestedVerificationCode
	return r
}

func (r LegalEntitiesApiUpdateLegalEntityInput) LegalEntityInfo(legalEntityInfo LegalEntityInfo) LegalEntitiesApiUpdateLegalEntityInput {
	r.legalEntityInfo = &legalEntityInfo
	return r
}

/*
Prepare a request for UpdateLegalEntity
@param id The unique identifier of the legal entity.
@return LegalEntitiesApiUpdateLegalEntityInput
*/
func (a *LegalEntitiesApi) UpdateLegalEntityInput(id string) LegalEntitiesApiUpdateLegalEntityInput {
	return LegalEntitiesApiUpdateLegalEntityInput{
		id: id,
	}
}

/*
UpdateLegalEntity Update a legal entity

Updates a legal entity.

 >To change the legal entity type, include only the new `type` in your request. To update the `entityAssociations` array, you need to replace the entire array. For example, if the array has 3 entries and you want to remove 1 entry, you need to PATCH the resource with the remaining 2 entries.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r LegalEntitiesApiUpdateLegalEntityInput - Request parameters, see UpdateLegalEntityInput
@return LegalEntity, *http.Response, error
*/
func (a *LegalEntitiesApi) UpdateLegalEntity(ctx context.Context, r LegalEntitiesApiUpdateLegalEntityInput) (LegalEntity, *http.Response, error) {
	res := &LegalEntity{}
	path := "/legalEntities/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.xRequestedVerificationCode != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "x-requested-verification-code", r.xRequestedVerificationCode, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.legalEntityInfo,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
