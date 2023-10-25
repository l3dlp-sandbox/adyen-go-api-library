/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// RecurringApi service
type RecurringApi common.Service

// All parameters accepted by RecurringApi.DeleteTokenForStoredPaymentDetails
type RecurringApiDeleteTokenForStoredPaymentDetailsInput struct {
	storedPaymentMethodId string
	shopperReference      *string
	merchantAccount       *string
}

// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address.
func (r RecurringApiDeleteTokenForStoredPaymentDetailsInput) ShopperReference(shopperReference string) RecurringApiDeleteTokenForStoredPaymentDetailsInput {
	r.shopperReference = &shopperReference
	return r
}

// Your merchant account.
func (r RecurringApiDeleteTokenForStoredPaymentDetailsInput) MerchantAccount(merchantAccount string) RecurringApiDeleteTokenForStoredPaymentDetailsInput {
	r.merchantAccount = &merchantAccount
	return r
}

/*
Prepare a request for DeleteTokenForStoredPaymentDetails
@param storedPaymentMethodId The unique identifier of the token.
@return RecurringApiDeleteTokenForStoredPaymentDetailsInput
*/
func (a *RecurringApi) DeleteTokenForStoredPaymentDetailsInput(storedPaymentMethodId string) RecurringApiDeleteTokenForStoredPaymentDetailsInput {
	return RecurringApiDeleteTokenForStoredPaymentDetailsInput{
		storedPaymentMethodId: storedPaymentMethodId,
	}
}

/*
DeleteTokenForStoredPaymentDetails Delete a token for stored payment details

Deletes the token identified in the path. The token can no longer be used with payment requests.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r RecurringApiDeleteTokenForStoredPaymentDetailsInput - Request parameters, see DeleteTokenForStoredPaymentDetailsInput
@return *http.Response, error
*/
func (a *RecurringApi) DeleteTokenForStoredPaymentDetails(ctx context.Context, r RecurringApiDeleteTokenForStoredPaymentDetailsInput) (*http.Response, error) {
	var res interface{}
	path := "/storedPaymentMethods/{storedPaymentMethodId}"
	path = strings.Replace(path, "{"+"storedPaymentMethodId"+"}", url.PathEscape(common.ParameterValueToString(r.storedPaymentMethodId, "storedPaymentMethodId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.shopperReference != nil {
		common.ParameterAddToQuery(queryParams, "shopperReference", r.shopperReference, "")
	}
	if r.merchantAccount != nil {
		common.ParameterAddToQuery(queryParams, "merchantAccount", r.merchantAccount, "")
	}
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

// All parameters accepted by RecurringApi.GetTokensForStoredPaymentDetails
type RecurringApiGetTokensForStoredPaymentDetailsInput struct {
	shopperReference *string
	merchantAccount  *string
}

// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address.
func (r RecurringApiGetTokensForStoredPaymentDetailsInput) ShopperReference(shopperReference string) RecurringApiGetTokensForStoredPaymentDetailsInput {
	r.shopperReference = &shopperReference
	return r
}

// Your merchant account.
func (r RecurringApiGetTokensForStoredPaymentDetailsInput) MerchantAccount(merchantAccount string) RecurringApiGetTokensForStoredPaymentDetailsInput {
	r.merchantAccount = &merchantAccount
	return r
}

/*
Prepare a request for GetTokensForStoredPaymentDetails

@return RecurringApiGetTokensForStoredPaymentDetailsInput
*/
func (a *RecurringApi) GetTokensForStoredPaymentDetailsInput() RecurringApiGetTokensForStoredPaymentDetailsInput {
	return RecurringApiGetTokensForStoredPaymentDetailsInput{}
}

/*
GetTokensForStoredPaymentDetails Get tokens for stored payment details

Lists the tokens for stored payment details for the shopper identified in the path, if there are any available. The token ID can be used with payment requests for the shopper's payment. A summary of the stored details is included.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r RecurringApiGetTokensForStoredPaymentDetailsInput - Request parameters, see GetTokensForStoredPaymentDetailsInput
@return ListStoredPaymentMethodsResponse, *http.Response, error
*/
func (a *RecurringApi) GetTokensForStoredPaymentDetails(ctx context.Context, r RecurringApiGetTokensForStoredPaymentDetailsInput) (ListStoredPaymentMethodsResponse, *http.Response, error) {
	res := &ListStoredPaymentMethodsResponse{}
	path := "/storedPaymentMethods"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.shopperReference != nil {
		common.ParameterAddToQuery(queryParams, "shopperReference", r.shopperReference, "")
	}
	if r.merchantAccount != nil {
		common.ParameterAddToQuery(queryParams, "merchantAccount", r.merchantAccount, "")
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
