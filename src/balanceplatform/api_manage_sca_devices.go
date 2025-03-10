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

	"github.com/adyen/adyen-go-api-library/v20/src/common"
)

// ManageSCADevicesApi service
type ManageSCADevicesApi common.Service

// All parameters accepted by ManageSCADevicesApi.CompleteAssociationBetweenScaDeviceAndResource
type ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput struct {
	deviceId                   string
	associationFinaliseRequest *AssociationFinaliseRequest
}

func (r ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput) AssociationFinaliseRequest(associationFinaliseRequest AssociationFinaliseRequest) ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput {
	r.associationFinaliseRequest = &associationFinaliseRequest
	return r
}

/*
Prepare a request for CompleteAssociationBetweenScaDeviceAndResource
@param deviceId The unique identifier of the SCA device that you are associating with a resource.
@return ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput
*/
func (a *ManageSCADevicesApi) CompleteAssociationBetweenScaDeviceAndResourceInput(deviceId string) ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput {
	return ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput{
		deviceId: deviceId,
	}
}

/*
CompleteAssociationBetweenScaDeviceAndResource Complete an association between an SCA device and a resource

Completes an association between a user's registered SCA device and an Adyen resource. For example, you can associate an SCA device with additional [business accounts](https://docs.adyen.com/platforms/business-accounts/sca/register-devices) or [Adyen-issued cards](https://docs.adyen.com/issuing/3d-secure/oob-auth-sdk/register-devices).

To complete the association, this endpoint validates the authentication data of the registered device.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput - Request parameters, see CompleteAssociationBetweenScaDeviceAndResourceInput
@return AssociationFinaliseResponse, *http.Response, error
*/
func (a *ManageSCADevicesApi) CompleteAssociationBetweenScaDeviceAndResource(ctx context.Context, r ManageSCADevicesApiCompleteAssociationBetweenScaDeviceAndResourceInput) (AssociationFinaliseResponse, *http.Response, error) {
	res := &AssociationFinaliseResponse{}
	path := "/registeredDevices/{deviceId}/associations"
	path = strings.Replace(path, "{"+"deviceId"+"}", url.PathEscape(common.ParameterValueToString(r.deviceId, "deviceId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.associationFinaliseRequest,
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

// All parameters accepted by ManageSCADevicesApi.CompleteRegistrationOfScaDevice
type ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput struct {
	id                 string
	registerSCARequest *RegisterSCARequest
}

func (r ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput) RegisterSCARequest(registerSCARequest RegisterSCARequest) ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput {
	r.registerSCARequest = &registerSCARequest
	return r
}

/*
Prepare a request for CompleteRegistrationOfScaDevice
@param id The unique identifier of the SCA device. You obtain this `id` in the response of a POST&nbsp;[/registeredDevices](https://docs.adyen.com/api-explorer/balanceplatform/2/post/registeredDevices#responses-200-id) request.
@return ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput
*/
func (a *ManageSCADevicesApi) CompleteRegistrationOfScaDeviceInput(id string) ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput {
	return ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput{
		id: id,
	}
}

/*
CompleteRegistrationOfScaDevice Complete the registration of an SCA device

Completes the registration of an SCA device by validating the authentication data of the device. You can register SCA devices for [business accounts](https://docs.adyen.com/platforms/business-accounts/sca) or [Adyen-issued cards](https://docs.adyen.com/issuing/3d-secure/oob-auth-sdk).



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput - Request parameters, see CompleteRegistrationOfScaDeviceInput
@return RegisterSCAFinalResponse, *http.Response, error
*/
func (a *ManageSCADevicesApi) CompleteRegistrationOfScaDevice(ctx context.Context, r ManageSCADevicesApiCompleteRegistrationOfScaDeviceInput) (RegisterSCAFinalResponse, *http.Response, error) {
	res := &RegisterSCAFinalResponse{}
	path := "/registeredDevices/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.registerSCARequest,
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

// All parameters accepted by ManageSCADevicesApi.DeleteRegistrationOfScaDevice
type ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput struct {
	id                  string
	paymentInstrumentId *string
}

// The unique identifier of the payment instrument linked to the SCA device.
func (r ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput) PaymentInstrumentId(paymentInstrumentId string) ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput {
	r.paymentInstrumentId = &paymentInstrumentId
	return r
}

/*
Prepare a request for DeleteRegistrationOfScaDevice
@param id The unique identifier of the SCA device.
@return ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput
*/
func (a *ManageSCADevicesApi) DeleteRegistrationOfScaDeviceInput(id string) ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput {
	return ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput{
		id: id,
	}
}

/*
DeleteRegistrationOfScaDevice Delete a registration of an SCA device

Deletes an SCA device from the list of registered devices of a specific payment instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput - Request parameters, see DeleteRegistrationOfScaDeviceInput
@return *http.Response, error
*/
func (a *ManageSCADevicesApi) DeleteRegistrationOfScaDevice(ctx context.Context, r ManageSCADevicesApiDeleteRegistrationOfScaDeviceInput) (*http.Response, error) {
	var res interface{}
	path := "/registeredDevices/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.paymentInstrumentId != nil {
		common.ParameterAddToQuery(queryParams, "paymentInstrumentId", r.paymentInstrumentId, "")
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

// All parameters accepted by ManageSCADevicesApi.InitiateAssociationBetweenScaDeviceAndResource
type ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput struct {
	deviceId                   string
	associationInitiateRequest *AssociationInitiateRequest
}

func (r ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput) AssociationInitiateRequest(associationInitiateRequest AssociationInitiateRequest) ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput {
	r.associationInitiateRequest = &associationInitiateRequest
	return r
}

/*
Prepare a request for InitiateAssociationBetweenScaDeviceAndResource
@param deviceId The unique identifier of the SCA device that you are associating with a resource.
@return ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput
*/
func (a *ManageSCADevicesApi) InitiateAssociationBetweenScaDeviceAndResourceInput(deviceId string) ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput {
	return ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput{
		deviceId: deviceId,
	}
}

/*
InitiateAssociationBetweenScaDeviceAndResource Initiate an association between an SCA device and a resource

Initiates an association between a user's registered SCA device and an Adyen resource. For example, you can associate an SCA device with additional [business accounts](https://docs.adyen.com/platforms/business-accounts/sca/register-devices) or [Adyen-issued cards](https://docs.adyen.com/issuing/3d-secure/oob-auth-sdk/register-devices).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput - Request parameters, see InitiateAssociationBetweenScaDeviceAndResourceInput
@return AssociationInitiateResponse, *http.Response, error
*/
func (a *ManageSCADevicesApi) InitiateAssociationBetweenScaDeviceAndResource(ctx context.Context, r ManageSCADevicesApiInitiateAssociationBetweenScaDeviceAndResourceInput) (AssociationInitiateResponse, *http.Response, error) {
	res := &AssociationInitiateResponse{}
	path := "/registeredDevices/{deviceId}/associations"
	path = strings.Replace(path, "{"+"deviceId"+"}", url.PathEscape(common.ParameterValueToString(r.deviceId, "deviceId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.associationInitiateRequest,
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

// All parameters accepted by ManageSCADevicesApi.InitiateRegistrationOfScaDevice
type ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput struct {
	registerSCARequest *RegisterSCARequest
}

func (r ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput) RegisterSCARequest(registerSCARequest RegisterSCARequest) ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput {
	r.registerSCARequest = &registerSCARequest
	return r
}

/*
Prepare a request for InitiateRegistrationOfScaDevice

@return ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput
*/
func (a *ManageSCADevicesApi) InitiateRegistrationOfScaDeviceInput() ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput {
	return ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput{}
}

/*
InitiateRegistrationOfScaDevice Initiate the registration of an SCA device

Initiates the registration of a user's device for Strong Customer Authentication (SCA). You can register SCA devices for [business accounts](https://docs.adyen.com/platforms/business-accounts/sca/register-devices) or [Adyen-issued cards](https://docs.adyen.com/issuing/3d-secure/oob-auth-sdk/register-devices).

For a successful request, the device must be eligible for SCA.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput - Request parameters, see InitiateRegistrationOfScaDeviceInput
@return RegisterSCAResponse, *http.Response, error
*/
func (a *ManageSCADevicesApi) InitiateRegistrationOfScaDevice(ctx context.Context, r ManageSCADevicesApiInitiateRegistrationOfScaDeviceInput) (RegisterSCAResponse, *http.Response, error) {
	res := &RegisterSCAResponse{}
	path := "/registeredDevices"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.registerSCARequest,
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

// All parameters accepted by ManageSCADevicesApi.ListRegisteredScaDevices
type ManageSCADevicesApiListRegisteredScaDevicesInput struct {
	paymentInstrumentId *string
	pageNumber          *int32
	pageSize            *int32
}

// The unique identifier of a payment instrument. It limits the returned list to SCA devices associated to this payment instrument.
func (r ManageSCADevicesApiListRegisteredScaDevicesInput) PaymentInstrumentId(paymentInstrumentId string) ManageSCADevicesApiListRegisteredScaDevicesInput {
	r.paymentInstrumentId = &paymentInstrumentId
	return r
}

// The index of the page to retrieve. The index of the first page is 0 (zero).  Default: 0.
func (r ManageSCADevicesApiListRegisteredScaDevicesInput) PageNumber(pageNumber int32) ManageSCADevicesApiListRegisteredScaDevicesInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page.  Default: 20. Maximum: 100.
func (r ManageSCADevicesApiListRegisteredScaDevicesInput) PageSize(pageSize int32) ManageSCADevicesApiListRegisteredScaDevicesInput {
	r.pageSize = &pageSize
	return r
}

/*
Prepare a request for ListRegisteredScaDevices

@return ManageSCADevicesApiListRegisteredScaDevicesInput
*/
func (a *ManageSCADevicesApi) ListRegisteredScaDevicesInput() ManageSCADevicesApiListRegisteredScaDevicesInput {
	return ManageSCADevicesApiListRegisteredScaDevicesInput{}
}

/*
ListRegisteredScaDevices Get a list of registered SCA devices

Get a paginated list of the SCA devices you have currently registered for a specific payment instrument.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageSCADevicesApiListRegisteredScaDevicesInput - Request parameters, see ListRegisteredScaDevicesInput
@return SearchRegisteredDevicesResponse, *http.Response, error
*/
func (a *ManageSCADevicesApi) ListRegisteredScaDevices(ctx context.Context, r ManageSCADevicesApiListRegisteredScaDevicesInput) (SearchRegisteredDevicesResponse, *http.Response, error) {
	res := &SearchRegisteredDevicesResponse{}
	path := "/registeredDevices"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.paymentInstrumentId != nil {
		common.ParameterAddToQuery(queryParams, "paymentInstrumentId", r.paymentInstrumentId, "")
	}
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
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
