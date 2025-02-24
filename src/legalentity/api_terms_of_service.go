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

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// TermsOfServiceApi service
type TermsOfServiceApi common.Service

// All parameters accepted by TermsOfServiceApi.AcceptTermsOfService
type TermsOfServiceApiAcceptTermsOfServiceInput struct {
	id                          string
	termsofservicedocumentid    string
	acceptTermsOfServiceRequest *AcceptTermsOfServiceRequest
}

func (r TermsOfServiceApiAcceptTermsOfServiceInput) AcceptTermsOfServiceRequest(acceptTermsOfServiceRequest AcceptTermsOfServiceRequest) TermsOfServiceApiAcceptTermsOfServiceInput {
	r.acceptTermsOfServiceRequest = &acceptTermsOfServiceRequest
	return r
}

/*
Prepare a request for AcceptTermsOfService
@param id The unique identifier of the legal entity.  For sole proprietorships, this is the individual legal entity ID of the owner.  For organizations, this is the ID of the organization.  For legal representatives of individuals, this is the ID of the individual.  @param termsofservicedocumentid The unique identifier of the Terms of Service document.
@return TermsOfServiceApiAcceptTermsOfServiceInput
*/
func (a *TermsOfServiceApi) AcceptTermsOfServiceInput(id string, termsofservicedocumentid string) TermsOfServiceApiAcceptTermsOfServiceInput {
	return TermsOfServiceApiAcceptTermsOfServiceInput{
		id:                       id,
		termsofservicedocumentid: termsofservicedocumentid,
	}
}

/*
AcceptTermsOfService Accept Terms of Service

Accepts Terms of Service.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TermsOfServiceApiAcceptTermsOfServiceInput - Request parameters, see AcceptTermsOfServiceInput
@return AcceptTermsOfServiceResponse, *http.Response, error
*/
func (a *TermsOfServiceApi) AcceptTermsOfService(ctx context.Context, r TermsOfServiceApiAcceptTermsOfServiceInput) (AcceptTermsOfServiceResponse, *http.Response, error) {
	res := &AcceptTermsOfServiceResponse{}
	path := "/legalEntities/{id}/termsOfService/{termsofservicedocumentid}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	path = strings.Replace(path, "{"+"termsofservicedocumentid"+"}", url.PathEscape(common.ParameterValueToString(r.termsofservicedocumentid, "termsofservicedocumentid")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.acceptTermsOfServiceRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by TermsOfServiceApi.GetAcceptedTermsOfServiceDocument
type TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput struct {
	id                                string
	termsofserviceacceptancereference string
	termsOfServiceDocumentFormat      *string
}

// The format of the Terms of Service document. Possible values: **JSON**, **PDF**, or **TXT**
func (r TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput) TermsOfServiceDocumentFormat(termsOfServiceDocumentFormat string) TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput {
	r.termsOfServiceDocumentFormat = &termsOfServiceDocumentFormat
	return r
}

/*
Prepare a request for GetAcceptedTermsOfServiceDocument
@param id The unique identifier of the legal entity. For sole proprietorship, this is the individual legal entity ID of the owner. For organizations, this is the ID of the organization.@param termsofserviceacceptancereference An Adyen-generated reference for the accepted Terms of Service.
@return TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput
*/
func (a *TermsOfServiceApi) GetAcceptedTermsOfServiceDocumentInput(id string, termsofserviceacceptancereference string) TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput {
	return TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput{
		id:                                id,
		termsofserviceacceptancereference: termsofserviceacceptancereference,
	}
}

/*
GetAcceptedTermsOfServiceDocument Get accepted Terms of Service document

Returns the accepted Terms of Service document for a legal entity.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput - Request parameters, see GetAcceptedTermsOfServiceDocumentInput
@return GetAcceptedTermsOfServiceDocumentResponse, *http.Response, error
*/
func (a *TermsOfServiceApi) GetAcceptedTermsOfServiceDocument(ctx context.Context, r TermsOfServiceApiGetAcceptedTermsOfServiceDocumentInput) (GetAcceptedTermsOfServiceDocumentResponse, *http.Response, error) {
	res := &GetAcceptedTermsOfServiceDocumentResponse{}
	path := "/legalEntities/{id}/acceptedTermsOfServiceDocument/{termsofserviceacceptancereference}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	path = strings.Replace(path, "{"+"termsofserviceacceptancereference"+"}", url.PathEscape(common.ParameterValueToString(r.termsofserviceacceptancereference, "termsofserviceacceptancereference")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.termsOfServiceDocumentFormat != nil {
		common.ParameterAddToQuery(queryParams, "termsOfServiceDocumentFormat", r.termsOfServiceDocumentFormat, "")
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

	return *res, httpRes, err
}

// All parameters accepted by TermsOfServiceApi.GetTermsOfServiceDocument
type TermsOfServiceApiGetTermsOfServiceDocumentInput struct {
	id                               string
	getTermsOfServiceDocumentRequest *GetTermsOfServiceDocumentRequest
}

func (r TermsOfServiceApiGetTermsOfServiceDocumentInput) GetTermsOfServiceDocumentRequest(getTermsOfServiceDocumentRequest GetTermsOfServiceDocumentRequest) TermsOfServiceApiGetTermsOfServiceDocumentInput {
	r.getTermsOfServiceDocumentRequest = &getTermsOfServiceDocumentRequest
	return r
}

/*
Prepare a request for GetTermsOfServiceDocument
@param id The unique identifier of the legal entity. For sole proprietorships, this is the individual legal entity ID of the owner. For organizations, this is the ID of the organization.
@return TermsOfServiceApiGetTermsOfServiceDocumentInput
*/
func (a *TermsOfServiceApi) GetTermsOfServiceDocumentInput(id string) TermsOfServiceApiGetTermsOfServiceDocumentInput {
	return TermsOfServiceApiGetTermsOfServiceDocumentInput{
		id: id,
	}
}

/*
GetTermsOfServiceDocument Get Terms of Service document

Returns the Terms of Service document for a legal entity.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TermsOfServiceApiGetTermsOfServiceDocumentInput - Request parameters, see GetTermsOfServiceDocumentInput
@return GetTermsOfServiceDocumentResponse, *http.Response, error
*/
func (a *TermsOfServiceApi) GetTermsOfServiceDocument(ctx context.Context, r TermsOfServiceApiGetTermsOfServiceDocumentInput) (GetTermsOfServiceDocumentResponse, *http.Response, error) {
	res := &GetTermsOfServiceDocumentResponse{}
	path := "/legalEntities/{id}/termsOfService"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.getTermsOfServiceDocumentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by TermsOfServiceApi.GetTermsOfServiceInformationForLegalEntity
type TermsOfServiceApiGetTermsOfServiceInformationForLegalEntityInput struct {
	id string
}

/*
Prepare a request for GetTermsOfServiceInformationForLegalEntity
@param id The unique identifier of the legal entity. For sole proprietorships, this is the individual legal entity ID of the owner. For organizations, this is the ID of the organization.
@return TermsOfServiceApiGetTermsOfServiceInformationForLegalEntityInput
*/
func (a *TermsOfServiceApi) GetTermsOfServiceInformationForLegalEntityInput(id string) TermsOfServiceApiGetTermsOfServiceInformationForLegalEntityInput {
	return TermsOfServiceApiGetTermsOfServiceInformationForLegalEntityInput{
		id: id,
	}
}

/*
GetTermsOfServiceInformationForLegalEntity Get Terms of Service information for a legal entity

Returns Terms of Service information for a legal entity.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TermsOfServiceApiGetTermsOfServiceInformationForLegalEntityInput - Request parameters, see GetTermsOfServiceInformationForLegalEntityInput
@return GetTermsOfServiceAcceptanceInfosResponse, *http.Response, error
*/
func (a *TermsOfServiceApi) GetTermsOfServiceInformationForLegalEntity(ctx context.Context, r TermsOfServiceApiGetTermsOfServiceInformationForLegalEntityInput) (GetTermsOfServiceAcceptanceInfosResponse, *http.Response, error) {
	res := &GetTermsOfServiceAcceptanceInfosResponse{}
	path := "/legalEntities/{id}/termsOfServiceAcceptanceInfos"
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

// All parameters accepted by TermsOfServiceApi.GetTermsOfServiceStatus
type TermsOfServiceApiGetTermsOfServiceStatusInput struct {
	id string
}

/*
Prepare a request for GetTermsOfServiceStatus
@param id The unique identifier of the legal entity. For sole proprietorships, this is the individual legal entity ID of the owner. For organizations, this is the ID of the organization.
@return TermsOfServiceApiGetTermsOfServiceStatusInput
*/
func (a *TermsOfServiceApi) GetTermsOfServiceStatusInput(id string) TermsOfServiceApiGetTermsOfServiceStatusInput {
	return TermsOfServiceApiGetTermsOfServiceStatusInput{
		id: id,
	}
}

/*
GetTermsOfServiceStatus Get Terms of Service status

Returns the required types of Terms of Service that need to be accepted by a legal entity.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TermsOfServiceApiGetTermsOfServiceStatusInput - Request parameters, see GetTermsOfServiceStatusInput
@return CalculateTermsOfServiceStatusResponse, *http.Response, error
*/
func (a *TermsOfServiceApi) GetTermsOfServiceStatus(ctx context.Context, r TermsOfServiceApiGetTermsOfServiceStatusInput) (CalculateTermsOfServiceStatusResponse, *http.Response, error) {
	res := &CalculateTermsOfServiceStatusResponse{}
	path := "/legalEntities/{id}/termsOfServiceStatus"
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
