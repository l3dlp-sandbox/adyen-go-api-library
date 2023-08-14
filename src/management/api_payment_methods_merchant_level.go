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

// PaymentMethodsMerchantLevelApi service
type PaymentMethodsMerchantLevelApi common.Service

// All parameters accepted by PaymentMethodsMerchantLevelApi.AddApplePayDomain
type PaymentMethodsMerchantLevelApiAddApplePayDomainInput struct {
	merchantId      string
	paymentMethodId string
	applePayInfo    *ApplePayInfo
}

func (r PaymentMethodsMerchantLevelApiAddApplePayDomainInput) ApplePayInfo(applePayInfo ApplePayInfo) PaymentMethodsMerchantLevelApiAddApplePayDomainInput {
	r.applePayInfo = &applePayInfo
	return r
}

/*
Prepare a request for AddApplePayDomain
@param merchantId The unique identifier of the merchant account.@param paymentMethodId The unique identifier of the payment method.
@return PaymentMethodsMerchantLevelApiAddApplePayDomainInput
*/
func (a *PaymentMethodsMerchantLevelApi) AddApplePayDomainInput(merchantId string, paymentMethodId string) PaymentMethodsMerchantLevelApiAddApplePayDomainInput {
	return PaymentMethodsMerchantLevelApiAddApplePayDomainInput{
		merchantId:      merchantId,
		paymentMethodId: paymentMethodId,
	}
}

/*
AddApplePayDomain Add an Apple Pay domain

Adds the new domain to the list of Apple Pay domains that are registered with the merchant account and the payment method identified in the path. For more information, see [Apple Pay documentation](https://docs.adyen.com/payment-methods/apple-pay/enable-apple-pay#register-merchant-domain).

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentMethodsMerchantLevelApiAddApplePayDomainInput - Request parameters, see AddApplePayDomainInput
@return *http.Response, error
*/
func (a *PaymentMethodsMerchantLevelApi) AddApplePayDomain(ctx context.Context, r PaymentMethodsMerchantLevelApiAddApplePayDomainInput) (*http.Response, error) {
	var res interface{}
	path := "/merchants/{merchantId}/paymentMethodSettings/{paymentMethodId}/addApplePayDomains"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"paymentMethodId"+"}", url.PathEscape(common.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.applePayInfo,
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

// All parameters accepted by PaymentMethodsMerchantLevelApi.GetAllPaymentMethods
type PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput struct {
	merchantId     string
	storeId        *string
	businessLineId *string
	pageSize       *int32
	pageNumber     *int32
}

// The unique identifier of the store for which to return the payment methods.
func (r PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput) StoreId(storeId string) PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput {
	r.storeId = &storeId
	return r
}

// The unique identifier of the Business Line for which to return the payment methods.
func (r PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput) BusinessLineId(businessLineId string) PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput {
	r.businessLineId = &businessLineId
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput) PageSize(pageSize int32) PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput {
	r.pageSize = &pageSize
	return r
}

// The number of the page to fetch.
func (r PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput) PageNumber(pageNumber int32) PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput {
	r.pageNumber = &pageNumber
	return r
}

/*
Prepare a request for GetAllPaymentMethods
@param merchantId The unique identifier of the merchant account.
@return PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput
*/
func (a *PaymentMethodsMerchantLevelApi) GetAllPaymentMethodsInput(merchantId string) PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput {
	return PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput{
		merchantId: merchantId,
	}
}

/*
GetAllPaymentMethods Get all payment methods

Returns details for all payment methods of the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput - Request parameters, see GetAllPaymentMethodsInput
@return PaymentMethodResponse, *http.Response, error
*/
func (a *PaymentMethodsMerchantLevelApi) GetAllPaymentMethods(ctx context.Context, r PaymentMethodsMerchantLevelApiGetAllPaymentMethodsInput) (PaymentMethodResponse, *http.Response, error) {
	res := &PaymentMethodResponse{}
	path := "/merchants/{merchantId}/paymentMethodSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.storeId != nil {
		common.ParameterAddToQuery(queryParams, "storeId", r.storeId, "")
	}
	if r.businessLineId != nil {
		common.ParameterAddToQuery(queryParams, "businessLineId", r.businessLineId, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
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

// All parameters accepted by PaymentMethodsMerchantLevelApi.GetApplePayDomains
type PaymentMethodsMerchantLevelApiGetApplePayDomainsInput struct {
	merchantId      string
	paymentMethodId string
}

/*
Prepare a request for GetApplePayDomains
@param merchantId The unique identifier of the merchant account.@param paymentMethodId The unique identifier of the payment method.
@return PaymentMethodsMerchantLevelApiGetApplePayDomainsInput
*/
func (a *PaymentMethodsMerchantLevelApi) GetApplePayDomainsInput(merchantId string, paymentMethodId string) PaymentMethodsMerchantLevelApiGetApplePayDomainsInput {
	return PaymentMethodsMerchantLevelApiGetApplePayDomainsInput{
		merchantId:      merchantId,
		paymentMethodId: paymentMethodId,
	}
}

/*
GetApplePayDomains Get Apple Pay domains

Returns all Apple Pay domains that are registered with the merchant account and the payment method identified in the path. For more information, see [Apple Pay documentation](https://docs.adyen.com/payment-methods/apple-pay/enable-apple-pay#register-merchant-domain).

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentMethodsMerchantLevelApiGetApplePayDomainsInput - Request parameters, see GetApplePayDomainsInput
@return ApplePayInfo, *http.Response, error
*/
func (a *PaymentMethodsMerchantLevelApi) GetApplePayDomains(ctx context.Context, r PaymentMethodsMerchantLevelApiGetApplePayDomainsInput) (ApplePayInfo, *http.Response, error) {
	res := &ApplePayInfo{}
	path := "/merchants/{merchantId}/paymentMethodSettings/{paymentMethodId}/getApplePayDomains"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"paymentMethodId"+"}", url.PathEscape(common.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)
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

// All parameters accepted by PaymentMethodsMerchantLevelApi.GetPaymentMethodDetails
type PaymentMethodsMerchantLevelApiGetPaymentMethodDetailsInput struct {
	merchantId      string
	paymentMethodId string
}

/*
Prepare a request for GetPaymentMethodDetails
@param merchantId The unique identifier of the merchant account.@param paymentMethodId The unique identifier of the payment method.
@return PaymentMethodsMerchantLevelApiGetPaymentMethodDetailsInput
*/
func (a *PaymentMethodsMerchantLevelApi) GetPaymentMethodDetailsInput(merchantId string, paymentMethodId string) PaymentMethodsMerchantLevelApiGetPaymentMethodDetailsInput {
	return PaymentMethodsMerchantLevelApiGetPaymentMethodDetailsInput{
		merchantId:      merchantId,
		paymentMethodId: paymentMethodId,
	}
}

/*
GetPaymentMethodDetails Get payment method details

Returns details for the merchant account and the payment method identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentMethodsMerchantLevelApiGetPaymentMethodDetailsInput - Request parameters, see GetPaymentMethodDetailsInput
@return PaymentMethod, *http.Response, error
*/
func (a *PaymentMethodsMerchantLevelApi) GetPaymentMethodDetails(ctx context.Context, r PaymentMethodsMerchantLevelApiGetPaymentMethodDetailsInput) (PaymentMethod, *http.Response, error) {
	res := &PaymentMethod{}
	path := "/merchants/{merchantId}/paymentMethodSettings/{paymentMethodId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"paymentMethodId"+"}", url.PathEscape(common.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)
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

// All parameters accepted by PaymentMethodsMerchantLevelApi.RequestPaymentMethod
type PaymentMethodsMerchantLevelApiRequestPaymentMethodInput struct {
	merchantId             string
	paymentMethodSetupInfo *PaymentMethodSetupInfo
}

func (r PaymentMethodsMerchantLevelApiRequestPaymentMethodInput) PaymentMethodSetupInfo(paymentMethodSetupInfo PaymentMethodSetupInfo) PaymentMethodsMerchantLevelApiRequestPaymentMethodInput {
	r.paymentMethodSetupInfo = &paymentMethodSetupInfo
	return r
}

/*
Prepare a request for RequestPaymentMethod
@param merchantId The unique identifier of the merchant account.
@return PaymentMethodsMerchantLevelApiRequestPaymentMethodInput
*/
func (a *PaymentMethodsMerchantLevelApi) RequestPaymentMethodInput(merchantId string) PaymentMethodsMerchantLevelApiRequestPaymentMethodInput {
	return PaymentMethodsMerchantLevelApiRequestPaymentMethodInput{
		merchantId: merchantId,
	}
}

/*
RequestPaymentMethod Request a payment method

Sends a request to add a new payment method to the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentMethodsMerchantLevelApiRequestPaymentMethodInput - Request parameters, see RequestPaymentMethodInput
@return PaymentMethod, *http.Response, error
*/
func (a *PaymentMethodsMerchantLevelApi) RequestPaymentMethod(ctx context.Context, r PaymentMethodsMerchantLevelApiRequestPaymentMethodInput) (PaymentMethod, *http.Response, error) {
	res := &PaymentMethod{}
	path := "/merchants/{merchantId}/paymentMethodSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentMethodSetupInfo,
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

// All parameters accepted by PaymentMethodsMerchantLevelApi.UpdatePaymentMethod
type PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput struct {
	merchantId              string
	paymentMethodId         string
	updatePaymentMethodInfo *UpdatePaymentMethodInfo
}

func (r PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput) UpdatePaymentMethodInfo(updatePaymentMethodInfo UpdatePaymentMethodInfo) PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput {
	r.updatePaymentMethodInfo = &updatePaymentMethodInfo
	return r
}

/*
Prepare a request for UpdatePaymentMethod
@param merchantId The unique identifier of the merchant account.@param paymentMethodId The unique identifier of the payment method.
@return PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput
*/
func (a *PaymentMethodsMerchantLevelApi) UpdatePaymentMethodInput(merchantId string, paymentMethodId string) PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput {
	return PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput{
		merchantId:      merchantId,
		paymentMethodId: paymentMethodId,
	}
}

/*
UpdatePaymentMethod Update a payment method

Updates payment method details for the merchant account and the payment method identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read and write


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput - Request parameters, see UpdatePaymentMethodInput
@return PaymentMethod, *http.Response, error
*/
func (a *PaymentMethodsMerchantLevelApi) UpdatePaymentMethod(ctx context.Context, r PaymentMethodsMerchantLevelApiUpdatePaymentMethodInput) (PaymentMethod, *http.Response, error) {
	res := &PaymentMethod{}
	path := "/merchants/{merchantId}/paymentMethodSettings/{paymentMethodId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"paymentMethodId"+"}", url.PathEscape(common.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updatePaymentMethodInfo,
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