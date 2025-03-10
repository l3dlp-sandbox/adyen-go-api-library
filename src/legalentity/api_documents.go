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

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// DocumentsApi service
type DocumentsApi common.Service

// All parameters accepted by DocumentsApi.DeleteDocument
type DocumentsApiDeleteDocumentInput struct {
	id string
}

/*
Prepare a request for DeleteDocument
@param id The unique identifier of the document to be deleted.
@return DocumentsApiDeleteDocumentInput
*/
func (a *DocumentsApi) DeleteDocumentInput(id string) DocumentsApiDeleteDocumentInput {
	return DocumentsApiDeleteDocumentInput{
		id: id,
	}
}

/*
DeleteDocument Delete a document

Deletes a document.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r DocumentsApiDeleteDocumentInput - Request parameters, see DeleteDocumentInput
@return *http.Response, error
*/
func (a *DocumentsApi) DeleteDocument(ctx context.Context, r DocumentsApiDeleteDocumentInput) (*http.Response, error) {
	var res interface{}
	path := "/documents/{id}"
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

// All parameters accepted by DocumentsApi.GetDocument
type DocumentsApiGetDocumentInput struct {
	id          string
	skipContent *bool
}

// Do not load document content while fetching the document.
func (r DocumentsApiGetDocumentInput) SkipContent(skipContent bool) DocumentsApiGetDocumentInput {
	r.skipContent = &skipContent
	return r
}

/*
Prepare a request for GetDocument
@param id The unique identifier of the document.
@return DocumentsApiGetDocumentInput
*/
func (a *DocumentsApi) GetDocumentInput(id string) DocumentsApiGetDocumentInput {
	return DocumentsApiGetDocumentInput{
		id: id,
	}
}

/*
GetDocument Get a document

Returns a document.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r DocumentsApiGetDocumentInput - Request parameters, see GetDocumentInput
@return Document, *http.Response, error
*/
func (a *DocumentsApi) GetDocument(ctx context.Context, r DocumentsApiGetDocumentInput) (Document, *http.Response, error) {
	res := &Document{}
	path := "/documents/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.skipContent != nil {
		common.ParameterAddToQuery(queryParams, "skipContent", r.skipContent, "")
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

// All parameters accepted by DocumentsApi.UpdateDocument
type DocumentsApiUpdateDocumentInput struct {
	id                         string
	xRequestedVerificationCode *string
	document                   *Document
}

// Use the requested verification code 0_0001 to resolve any suberrors associated with the document. Requested verification codes can only be used in your test environment.
func (r DocumentsApiUpdateDocumentInput) XRequestedVerificationCode(xRequestedVerificationCode string) DocumentsApiUpdateDocumentInput {
	r.xRequestedVerificationCode = &xRequestedVerificationCode
	return r
}

func (r DocumentsApiUpdateDocumentInput) Document(document Document) DocumentsApiUpdateDocumentInput {
	r.document = &document
	return r
}

/*
Prepare a request for UpdateDocument
@param id The unique identifier of the document to be updated.
@return DocumentsApiUpdateDocumentInput
*/
func (a *DocumentsApi) UpdateDocumentInput(id string) DocumentsApiUpdateDocumentInput {
	return DocumentsApiUpdateDocumentInput{
		id: id,
	}
}

/*
UpdateDocument Update a document

Updates a document.

 >You can upload a maximum of 15 pages for photo IDs.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r DocumentsApiUpdateDocumentInput - Request parameters, see UpdateDocumentInput
@return Document, *http.Response, error
*/
func (a *DocumentsApi) UpdateDocument(ctx context.Context, r DocumentsApiUpdateDocumentInput) (Document, *http.Response, error) {
	res := &Document{}
	path := "/documents/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.xRequestedVerificationCode != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "x-requested-verification-code", r.xRequestedVerificationCode, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.document,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by DocumentsApi.UploadDocumentForVerificationChecks
type DocumentsApiUploadDocumentForVerificationChecksInput struct {
	xRequestedVerificationCode *string
	document                   *Document
}

// Use a suberror code as your requested verification code. You can include one code at a time in your request header. Requested verification codes can only be used in your test environment.
func (r DocumentsApiUploadDocumentForVerificationChecksInput) XRequestedVerificationCode(xRequestedVerificationCode string) DocumentsApiUploadDocumentForVerificationChecksInput {
	r.xRequestedVerificationCode = &xRequestedVerificationCode
	return r
}

func (r DocumentsApiUploadDocumentForVerificationChecksInput) Document(document Document) DocumentsApiUploadDocumentForVerificationChecksInput {
	r.document = &document
	return r
}

/*
Prepare a request for UploadDocumentForVerificationChecks

@return DocumentsApiUploadDocumentForVerificationChecksInput
*/
func (a *DocumentsApi) UploadDocumentForVerificationChecksInput() DocumentsApiUploadDocumentForVerificationChecksInput {
	return DocumentsApiUploadDocumentForVerificationChecksInput{}
}

/*
UploadDocumentForVerificationChecks Upload a document for verification checks

Uploads a document for verification checks.

 Adyen uses the information from the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities) to run automated verification checks. If these checks fail, you will be notified to provide additional documents.

 You should only upload documents when Adyen requests additional information for the legal entity.

 >You can upload a maximum of 15 pages for photo IDs.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r DocumentsApiUploadDocumentForVerificationChecksInput - Request parameters, see UploadDocumentForVerificationChecksInput
@return Document, *http.Response, error
*/
func (a *DocumentsApi) UploadDocumentForVerificationChecks(ctx context.Context, r DocumentsApiUploadDocumentForVerificationChecksInput) (Document, *http.Response, error) {
	res := &Document{}
	path := "/documents"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.xRequestedVerificationCode != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "x-requested-verification-code", r.xRequestedVerificationCode, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.document,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
