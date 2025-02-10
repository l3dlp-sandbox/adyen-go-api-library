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

// PCIQuestionnairesApi service
type PCIQuestionnairesApi common.Service

// All parameters accepted by PCIQuestionnairesApi.CalculatePciStatusOfLegalEntity
type PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput struct {
	id                        string
	calculatePciStatusRequest *CalculatePciStatusRequest
}

func (r PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput) CalculatePciStatusRequest(calculatePciStatusRequest CalculatePciStatusRequest) PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput {
	r.calculatePciStatusRequest = &calculatePciStatusRequest
	return r
}

/*
Prepare a request for CalculatePciStatusOfLegalEntity
@param id The unique identifier of the legal entity to calculate PCI status.
@return PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput
*/
func (a *PCIQuestionnairesApi) CalculatePciStatusOfLegalEntityInput(id string) PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput {
	return PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput{
		id: id,
	}
}

/*
CalculatePciStatusOfLegalEntity Calculate PCI status of a legal entity

Calculate PCI status of a legal entity.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput - Request parameters, see CalculatePciStatusOfLegalEntityInput
@return CalculatePciStatusResponse, *http.Response, error
*/
func (a *PCIQuestionnairesApi) CalculatePciStatusOfLegalEntity(ctx context.Context, r PCIQuestionnairesApiCalculatePciStatusOfLegalEntityInput) (CalculatePciStatusResponse, *http.Response, error) {
	res := &CalculatePciStatusResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires/signingRequired"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.calculatePciStatusRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PCIQuestionnairesApi.GeneratePciQuestionnaire
type PCIQuestionnairesApiGeneratePciQuestionnaireInput struct {
	id                            string
	generatePciDescriptionRequest *GeneratePciDescriptionRequest
}

func (r PCIQuestionnairesApiGeneratePciQuestionnaireInput) GeneratePciDescriptionRequest(generatePciDescriptionRequest GeneratePciDescriptionRequest) PCIQuestionnairesApiGeneratePciQuestionnaireInput {
	r.generatePciDescriptionRequest = &generatePciDescriptionRequest
	return r
}

/*
Prepare a request for GeneratePciQuestionnaire
@param id The unique identifier of the legal entity to get PCI questionnaire information.
@return PCIQuestionnairesApiGeneratePciQuestionnaireInput
*/
func (a *PCIQuestionnairesApi) GeneratePciQuestionnaireInput(id string) PCIQuestionnairesApiGeneratePciQuestionnaireInput {
	return PCIQuestionnairesApiGeneratePciQuestionnaireInput{
		id: id,
	}
}

/*
GeneratePciQuestionnaire Generate PCI questionnaire

Generates the required PCI questionnaires based on the user's [salesChannel](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businessLines__reqParam_salesChannels).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PCIQuestionnairesApiGeneratePciQuestionnaireInput - Request parameters, see GeneratePciQuestionnaireInput
@return GeneratePciDescriptionResponse, *http.Response, error
*/
func (a *PCIQuestionnairesApi) GeneratePciQuestionnaire(ctx context.Context, r PCIQuestionnairesApiGeneratePciQuestionnaireInput) (GeneratePciDescriptionResponse, *http.Response, error) {
	res := &GeneratePciDescriptionResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires/generatePciTemplates"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.generatePciDescriptionRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PCIQuestionnairesApi.GetPciQuestionnaire
type PCIQuestionnairesApiGetPciQuestionnaireInput struct {
	id    string
	pciid string
}

/*
Prepare a request for GetPciQuestionnaire
@param id The legal entity ID of the individual who signed the PCI questionnaire.@param pciid The unique identifier of the signed PCI questionnaire.
@return PCIQuestionnairesApiGetPciQuestionnaireInput
*/
func (a *PCIQuestionnairesApi) GetPciQuestionnaireInput(id string, pciid string) PCIQuestionnairesApiGetPciQuestionnaireInput {
	return PCIQuestionnairesApiGetPciQuestionnaireInput{
		id:    id,
		pciid: pciid,
	}
}

/*
GetPciQuestionnaire Get PCI questionnaire

Returns the signed PCI questionnaire.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PCIQuestionnairesApiGetPciQuestionnaireInput - Request parameters, see GetPciQuestionnaireInput
@return GetPciQuestionnaireResponse, *http.Response, error
*/
func (a *PCIQuestionnairesApi) GetPciQuestionnaire(ctx context.Context, r PCIQuestionnairesApiGetPciQuestionnaireInput) (GetPciQuestionnaireResponse, *http.Response, error) {
	res := &GetPciQuestionnaireResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires/{pciid}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	path = strings.Replace(path, "{"+"pciid"+"}", url.PathEscape(common.ParameterValueToString(r.pciid, "pciid")), -1)
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

// All parameters accepted by PCIQuestionnairesApi.GetPciQuestionnaireDetails
type PCIQuestionnairesApiGetPciQuestionnaireDetailsInput struct {
	id string
}

/*
Prepare a request for GetPciQuestionnaireDetails
@param id The unique identifier of the legal entity to get PCI questionnaire information.
@return PCIQuestionnairesApiGetPciQuestionnaireDetailsInput
*/
func (a *PCIQuestionnairesApi) GetPciQuestionnaireDetailsInput(id string) PCIQuestionnairesApiGetPciQuestionnaireDetailsInput {
	return PCIQuestionnairesApiGetPciQuestionnaireDetailsInput{
		id: id,
	}
}

/*
GetPciQuestionnaireDetails Get PCI questionnaire details

Get a list of signed PCI questionnaires.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PCIQuestionnairesApiGetPciQuestionnaireDetailsInput - Request parameters, see GetPciQuestionnaireDetailsInput
@return GetPciQuestionnaireInfosResponse, *http.Response, error
*/
func (a *PCIQuestionnairesApi) GetPciQuestionnaireDetails(ctx context.Context, r PCIQuestionnairesApiGetPciQuestionnaireDetailsInput) (GetPciQuestionnaireInfosResponse, *http.Response, error) {
	res := &GetPciQuestionnaireInfosResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires"
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

// All parameters accepted by PCIQuestionnairesApi.SignPciQuestionnaire
type PCIQuestionnairesApiSignPciQuestionnaireInput struct {
	id                string
	pciSigningRequest *PciSigningRequest
}

func (r PCIQuestionnairesApiSignPciQuestionnaireInput) PciSigningRequest(pciSigningRequest PciSigningRequest) PCIQuestionnairesApiSignPciQuestionnaireInput {
	r.pciSigningRequest = &pciSigningRequest
	return r
}

/*
Prepare a request for SignPciQuestionnaire
@param id The legal entity ID of the user that has a contractual relationship with your platform.
@return PCIQuestionnairesApiSignPciQuestionnaireInput
*/
func (a *PCIQuestionnairesApi) SignPciQuestionnaireInput(id string) PCIQuestionnairesApiSignPciQuestionnaireInput {
	return PCIQuestionnairesApiSignPciQuestionnaireInput{
		id: id,
	}
}

/*
SignPciQuestionnaire Sign PCI questionnaire

Signs the required PCI questionnaire.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PCIQuestionnairesApiSignPciQuestionnaireInput - Request parameters, see SignPciQuestionnaireInput
@return PciSigningResponse, *http.Response, error
*/
func (a *PCIQuestionnairesApi) SignPciQuestionnaire(ctx context.Context, r PCIQuestionnairesApiSignPciQuestionnaireInput) (PciSigningResponse, *http.Response, error) {
	res := &PciSigningResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires/signPciTemplates"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.pciSigningRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
