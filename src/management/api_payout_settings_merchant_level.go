/*
Management API

API version: 3
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

// PayoutSettingsMerchantLevelApi service
type PayoutSettingsMerchantLevelApi common.Service

// All parameters accepted by PayoutSettingsMerchantLevelApi.AddPayoutSetting
type PayoutSettingsMerchantLevelApiAddPayoutSettingInput struct {
	merchantId            string
	payoutSettingsRequest *PayoutSettingsRequest
}

func (r PayoutSettingsMerchantLevelApiAddPayoutSettingInput) PayoutSettingsRequest(payoutSettingsRequest PayoutSettingsRequest) PayoutSettingsMerchantLevelApiAddPayoutSettingInput {
	r.payoutSettingsRequest = &payoutSettingsRequest
	return r
}

/*
Prepare a request for AddPayoutSetting
@param merchantId The unique identifier of the merchant account.
@return PayoutSettingsMerchantLevelApiAddPayoutSettingInput
*/
func (a *PayoutSettingsMerchantLevelApi) AddPayoutSettingInput(merchantId string) PayoutSettingsMerchantLevelApiAddPayoutSettingInput {
	return PayoutSettingsMerchantLevelApiAddPayoutSettingInput{
		merchantId: merchantId,
	}
}

/*
AddPayoutSetting Add a payout setting

Sends a request to add a payout setting for the merchant account specified in the path. A payout setting links the merchant account to the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments) that contains the details of the payout bank account. Adyen verifies the bank account before allowing and enabling the payout setting.

If you're accepting payments in multiple currencies, you may add multiple payout settings for the merchant account.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):

* Management API—Payout account settings read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PayoutSettingsMerchantLevelApiAddPayoutSettingInput - Request parameters, see AddPayoutSettingInput
@return PayoutSettings, *http.Response, error
*/
func (a *PayoutSettingsMerchantLevelApi) AddPayoutSetting(ctx context.Context, r PayoutSettingsMerchantLevelApiAddPayoutSettingInput) (PayoutSettings, *http.Response, error) {
	res := &PayoutSettings{}
	path := "/merchants/{merchantId}/payoutSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.payoutSettingsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

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

// All parameters accepted by PayoutSettingsMerchantLevelApi.DeletePayoutSetting
type PayoutSettingsMerchantLevelApiDeletePayoutSettingInput struct {
	merchantId       string
	payoutSettingsId string
}

/*
Prepare a request for DeletePayoutSetting
@param merchantId The unique identifier of the merchant account.@param payoutSettingsId The unique identifier of the payout setting.
@return PayoutSettingsMerchantLevelApiDeletePayoutSettingInput
*/
func (a *PayoutSettingsMerchantLevelApi) DeletePayoutSettingInput(merchantId string, payoutSettingsId string) PayoutSettingsMerchantLevelApiDeletePayoutSettingInput {
	return PayoutSettingsMerchantLevelApiDeletePayoutSettingInput{
		merchantId:       merchantId,
		payoutSettingsId: payoutSettingsId,
	}
}

/*
DeletePayoutSetting Delete a payout setting

Deletes the payout setting identified in the path.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):

* Management API—Payout account settings read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PayoutSettingsMerchantLevelApiDeletePayoutSettingInput - Request parameters, see DeletePayoutSettingInput
@return *http.Response, error
*/
func (a *PayoutSettingsMerchantLevelApi) DeletePayoutSetting(ctx context.Context, r PayoutSettingsMerchantLevelApiDeletePayoutSettingInput) (*http.Response, error) {
	var res interface{}
	path := "/merchants/{merchantId}/payoutSettings/{payoutSettingsId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"payoutSettingsId"+"}", url.PathEscape(common.ParameterValueToString(r.payoutSettingsId, "payoutSettingsId")), -1)
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

	if httpRes == nil {
		return httpRes, err
	}

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

// All parameters accepted by PayoutSettingsMerchantLevelApi.GetPayoutSetting
type PayoutSettingsMerchantLevelApiGetPayoutSettingInput struct {
	merchantId       string
	payoutSettingsId string
}

/*
Prepare a request for GetPayoutSetting
@param merchantId The unique identifier of the merchant account.@param payoutSettingsId The unique identifier of the payout setting.
@return PayoutSettingsMerchantLevelApiGetPayoutSettingInput
*/
func (a *PayoutSettingsMerchantLevelApi) GetPayoutSettingInput(merchantId string, payoutSettingsId string) PayoutSettingsMerchantLevelApiGetPayoutSettingInput {
	return PayoutSettingsMerchantLevelApiGetPayoutSettingInput{
		merchantId:       merchantId,
		payoutSettingsId: payoutSettingsId,
	}
}

/*
GetPayoutSetting Get a payout setting

Returns the payout setting identified in the path.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payout account settings read

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PayoutSettingsMerchantLevelApiGetPayoutSettingInput - Request parameters, see GetPayoutSettingInput
@return PayoutSettings, *http.Response, error
*/
func (a *PayoutSettingsMerchantLevelApi) GetPayoutSetting(ctx context.Context, r PayoutSettingsMerchantLevelApiGetPayoutSettingInput) (PayoutSettings, *http.Response, error) {
	res := &PayoutSettings{}
	path := "/merchants/{merchantId}/payoutSettings/{payoutSettingsId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"payoutSettingsId"+"}", url.PathEscape(common.ParameterValueToString(r.payoutSettingsId, "payoutSettingsId")), -1)
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

// All parameters accepted by PayoutSettingsMerchantLevelApi.ListPayoutSettings
type PayoutSettingsMerchantLevelApiListPayoutSettingsInput struct {
	merchantId string
}

/*
Prepare a request for ListPayoutSettings
@param merchantId The unique identifier of the merchant account.
@return PayoutSettingsMerchantLevelApiListPayoutSettingsInput
*/
func (a *PayoutSettingsMerchantLevelApi) ListPayoutSettingsInput(merchantId string) PayoutSettingsMerchantLevelApiListPayoutSettingsInput {
	return PayoutSettingsMerchantLevelApiListPayoutSettingsInput{
		merchantId: merchantId,
	}
}

/*
ListPayoutSettings Get a list of payout settings

Returns the list of payout settings for the merchant account identified in the path.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payout account settings read

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PayoutSettingsMerchantLevelApiListPayoutSettingsInput - Request parameters, see ListPayoutSettingsInput
@return PayoutSettingsResponse, *http.Response, error
*/
func (a *PayoutSettingsMerchantLevelApi) ListPayoutSettings(ctx context.Context, r PayoutSettingsMerchantLevelApiListPayoutSettingsInput) (PayoutSettingsResponse, *http.Response, error) {
	res := &PayoutSettingsResponse{}
	path := "/merchants/{merchantId}/payoutSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
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

// All parameters accepted by PayoutSettingsMerchantLevelApi.UpdatePayoutSetting
type PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput struct {
	merchantId                  string
	payoutSettingsId            string
	updatePayoutSettingsRequest *UpdatePayoutSettingsRequest
}

func (r PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput) UpdatePayoutSettingsRequest(updatePayoutSettingsRequest UpdatePayoutSettingsRequest) PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput {
	r.updatePayoutSettingsRequest = &updatePayoutSettingsRequest
	return r
}

/*
Prepare a request for UpdatePayoutSetting
@param merchantId The unique identifier of the merchant account.@param payoutSettingsId The unique identifier of the payout setting.
@return PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput
*/
func (a *PayoutSettingsMerchantLevelApi) UpdatePayoutSettingInput(merchantId string, payoutSettingsId string) PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput {
	return PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput{
		merchantId:       merchantId,
		payoutSettingsId: payoutSettingsId,
	}
}

/*
UpdatePayoutSetting Update a payout setting

Updates the payout setting identified in the path. You can enable or disable the payout setting.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):

* Management API—Payout account settings read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput - Request parameters, see UpdatePayoutSettingInput
@return PayoutSettings, *http.Response, error
*/
func (a *PayoutSettingsMerchantLevelApi) UpdatePayoutSetting(ctx context.Context, r PayoutSettingsMerchantLevelApiUpdatePayoutSettingInput) (PayoutSettings, *http.Response, error) {
	res := &PayoutSettings{}
	path := "/merchants/{merchantId}/payoutSettings/{payoutSettingsId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"payoutSettingsId"+"}", url.PathEscape(common.ParameterValueToString(r.payoutSettingsId, "payoutSettingsId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updatePayoutSettingsRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

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
