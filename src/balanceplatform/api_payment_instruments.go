/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// PaymentInstrumentsApi service
type PaymentInstrumentsApi common.Service

// All parameters accepted by PaymentInstrumentsApi.CreatePaymentInstrument
type PaymentInstrumentsApiCreatePaymentInstrumentInput struct {
	paymentInstrumentInfo *PaymentInstrumentInfo
}

func (r PaymentInstrumentsApiCreatePaymentInstrumentInput) PaymentInstrumentInfo(paymentInstrumentInfo PaymentInstrumentInfo) PaymentInstrumentsApiCreatePaymentInstrumentInput {
	r.paymentInstrumentInfo = &paymentInstrumentInfo
	return r
}

/*
Prepare a request for CreatePaymentInstrument

@return PaymentInstrumentsApiCreatePaymentInstrumentInput
*/
func (a *PaymentInstrumentsApi) CreatePaymentInstrumentInput() PaymentInstrumentsApiCreatePaymentInstrumentInput {
	return PaymentInstrumentsApiCreatePaymentInstrumentInput{}
}

/*
CreatePaymentInstrument Create a payment instrument

Creates a payment instrument to issue a physical card, a virtual card, or a business account to your user.

 For more information, refer to [Issue cards](https://docs.adyen.com/issuing/create-cards) or [Issue business accounts](https://docs.adyen.com/marketplaces-and-platforms/business-accounts).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentInstrumentsApiCreatePaymentInstrumentInput - Request parameters, see CreatePaymentInstrumentInput
@return PaymentInstrument, *http.Response, error
*/
func (a *PaymentInstrumentsApi) CreatePaymentInstrument(ctx context.Context, r PaymentInstrumentsApiCreatePaymentInstrumentInput) (PaymentInstrument, *http.Response, error) {
	res := &PaymentInstrument{}
	path := "/paymentInstruments"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentInstrumentInfo,
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

// All parameters accepted by PaymentInstrumentsApi.GetAllTransactionRulesForPaymentInstrument
type PaymentInstrumentsApiGetAllTransactionRulesForPaymentInstrumentInput struct {
	id string
}

/*
Prepare a request for GetAllTransactionRulesForPaymentInstrument
@param id The unique identifier of the payment instrument.
@return PaymentInstrumentsApiGetAllTransactionRulesForPaymentInstrumentInput
*/
func (a *PaymentInstrumentsApi) GetAllTransactionRulesForPaymentInstrumentInput(id string) PaymentInstrumentsApiGetAllTransactionRulesForPaymentInstrumentInput {
	return PaymentInstrumentsApiGetAllTransactionRulesForPaymentInstrumentInput{
		id: id,
	}
}

/*
GetAllTransactionRulesForPaymentInstrument Get all transaction rules for a payment instrument

Returns a list of transaction rules associated with a payment instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentInstrumentsApiGetAllTransactionRulesForPaymentInstrumentInput - Request parameters, see GetAllTransactionRulesForPaymentInstrumentInput
@return TransactionRulesResponse, *http.Response, error
*/
func (a *PaymentInstrumentsApi) GetAllTransactionRulesForPaymentInstrument(ctx context.Context, r PaymentInstrumentsApiGetAllTransactionRulesForPaymentInstrumentInput) (TransactionRulesResponse, *http.Response, error) {
	res := &TransactionRulesResponse{}
	path := "/paymentInstruments/{id}/transactionRules"
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

// All parameters accepted by PaymentInstrumentsApi.GetPanOfPaymentInstrument
type PaymentInstrumentsApiGetPanOfPaymentInstrumentInput struct {
	id string
}

/*
Prepare a request for GetPanOfPaymentInstrument
@param id The unique identifier of the payment instrument.
@return PaymentInstrumentsApiGetPanOfPaymentInstrumentInput
*/
func (a *PaymentInstrumentsApi) GetPanOfPaymentInstrumentInput(id string) PaymentInstrumentsApiGetPanOfPaymentInstrumentInput {
	return PaymentInstrumentsApiGetPanOfPaymentInstrumentInput{
		id: id,
	}
}

/*
GetPanOfPaymentInstrument Get the PAN of a payment instrument

Returns the primary account number (PAN) of a payment instrument.

To make this request, your API credential must have the following [role](https://docs.adyen.com/issuing/manage-access/api-credentials-web-service#api-permissions):

* Balance Platform BCL PCI role

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentInstrumentsApiGetPanOfPaymentInstrumentInput - Request parameters, see GetPanOfPaymentInstrumentInput
@return PaymentInstrumentRevealInfo, *http.Response, error
*/
func (a *PaymentInstrumentsApi) GetPanOfPaymentInstrument(ctx context.Context, r PaymentInstrumentsApiGetPanOfPaymentInstrumentInput) (PaymentInstrumentRevealInfo, *http.Response, error) {
	res := &PaymentInstrumentRevealInfo{}
	path := "/paymentInstruments/{id}/reveal"
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

// All parameters accepted by PaymentInstrumentsApi.GetPaymentInstrument
type PaymentInstrumentsApiGetPaymentInstrumentInput struct {
	id string
}

/*
Prepare a request for GetPaymentInstrument
@param id The unique identifier of the payment instrument.
@return PaymentInstrumentsApiGetPaymentInstrumentInput
*/
func (a *PaymentInstrumentsApi) GetPaymentInstrumentInput(id string) PaymentInstrumentsApiGetPaymentInstrumentInput {
	return PaymentInstrumentsApiGetPaymentInstrumentInput{
		id: id,
	}
}

/*
GetPaymentInstrument Get a payment instrument

Returns the details of a payment instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentInstrumentsApiGetPaymentInstrumentInput - Request parameters, see GetPaymentInstrumentInput
@return PaymentInstrument, *http.Response, error
*/
func (a *PaymentInstrumentsApi) GetPaymentInstrument(ctx context.Context, r PaymentInstrumentsApiGetPaymentInstrumentInput) (PaymentInstrument, *http.Response, error) {
	res := &PaymentInstrument{}
	path := "/paymentInstruments/{id}"
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

// All parameters accepted by PaymentInstrumentsApi.ListNetworkTokens
type PaymentInstrumentsApiListNetworkTokensInput struct {
	id string
}

/*
Prepare a request for ListNetworkTokens
@param id The unique identifier of the payment instrument.
@return PaymentInstrumentsApiListNetworkTokensInput
*/
func (a *PaymentInstrumentsApi) ListNetworkTokensInput(id string) PaymentInstrumentsApiListNetworkTokensInput {
	return PaymentInstrumentsApiListNetworkTokensInput{
		id: id,
	}
}

/*
ListNetworkTokens List network tokens

List the network tokens connected to a payment instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentInstrumentsApiListNetworkTokensInput - Request parameters, see ListNetworkTokensInput
@return ListNetworkTokensResponse, *http.Response, error
*/
func (a *PaymentInstrumentsApi) ListNetworkTokens(ctx context.Context, r PaymentInstrumentsApiListNetworkTokensInput) (ListNetworkTokensResponse, *http.Response, error) {
	res := &ListNetworkTokensResponse{}
	path := "/paymentInstruments/{id}/networkTokens"
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

// All parameters accepted by PaymentInstrumentsApi.UpdatePaymentInstrument
type PaymentInstrumentsApiUpdatePaymentInstrumentInput struct {
	id                             string
	paymentInstrumentUpdateRequest *PaymentInstrumentUpdateRequest
}

func (r PaymentInstrumentsApiUpdatePaymentInstrumentInput) PaymentInstrumentUpdateRequest(paymentInstrumentUpdateRequest PaymentInstrumentUpdateRequest) PaymentInstrumentsApiUpdatePaymentInstrumentInput {
	r.paymentInstrumentUpdateRequest = &paymentInstrumentUpdateRequest
	return r
}

/*
Prepare a request for UpdatePaymentInstrument
@param id The unique identifier of the payment instrument.
@return PaymentInstrumentsApiUpdatePaymentInstrumentInput
*/
func (a *PaymentInstrumentsApi) UpdatePaymentInstrumentInput(id string) PaymentInstrumentsApiUpdatePaymentInstrumentInput {
	return PaymentInstrumentsApiUpdatePaymentInstrumentInput{
		id: id,
	}
}

/*
UpdatePaymentInstrument Update a payment instrument

Updates a payment instrument. Once a payment instrument is already active, you can only update its status. However, for cards created with **inactive** status, you can still update the balance account associated with the card.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentInstrumentsApiUpdatePaymentInstrumentInput - Request parameters, see UpdatePaymentInstrumentInput
@return UpdatePaymentInstrument, *http.Response, error
*/
func (a *PaymentInstrumentsApi) UpdatePaymentInstrument(ctx context.Context, r PaymentInstrumentsApiUpdatePaymentInstrumentInput) (UpdatePaymentInstrument, *http.Response, error) {
	res := &UpdatePaymentInstrument{}
	path := "/paymentInstruments/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentInstrumentUpdateRequest,
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